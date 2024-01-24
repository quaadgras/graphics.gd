//go:build !generate

package gd

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	"runtime.link/mmm"
)

/*
RegisterClass registers a struct available for use inside Godot
extending the given 'Parent' Godot class. The 'Struct' type must
be a named struct that embeds a [Class] field specifying the
parent class to extend.

	type MyClass struct {
		Class[MyClass, Node2D]
	}

Use this in a main or init function to register your Go structs
and they will become available within the Godot engine for use
in the editor and/or within scripts.

All exported fields and methods will be exposed to Godot, so
take caution when embedding types, as their fields and methods
will be promoted.

If the Struct extends [EditorPlugin] then it will be added to
the editor as a plugin.
*/
func Register[Struct Extends[Parent], Parent IsClass](godot Context) {
	var classType = reflect.TypeOf([0]Struct{}).Elem()
	var superType = reflect.TypeOf([0]Parent{}).Elem()
	if classType.Kind() != reflect.Struct || classType.Name() == "" {
		panic("gdextension.RegisterClass: Class type must be a named struct")
	}

	var className = godot.StringName(classType.Name())
	var superName = godot.StringName(strings.TrimPrefix(superType.Name(), "class"))

	godot.API.ClassDB.RegisterClass(godot.API.ExtensionToken, className, superName,
		&classImplementation{
			Name:  className,
			Super: superName,
			Godot: godot.API,
			Type:  classType,
			Value: reflect.New(classType).Interface().(IsClass),
		})

	registerSignals(godot, className, classType)
	registerMethods(godot, className, classType)
}

func convertName(fnName string) string {
	if fnName == "seek" {
		return "SeekTo"
	}
	if fnName == "type_string" {
		return "TypeToString"
	}
	fnName = strings.ToLower(fnName)
	joins := []string{}
	split := strings.Split(fnName, "_")
	for _, word := range split {
		joins = append(joins, strings.Title(word))
	}
	return strings.Join(joins, "")
}

// injectDependenciesInto the given freshly allocated value, this
// function makes sure any [Object] types are instantiated and
// that any referenced singletons are injected.
func injectDependenciesInto(godot Context, value reflect.Value) {
	if value.CanAddr() && value.Kind() != reflect.Struct {
		panic("gdextension.injectDependenciesInto: value must be an addressable struct")
	}
	localCtx := NewContext(godot.API)
	defer localCtx.End()

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)

		// support private fields.
		fieldValue := reflect.NewAt(field.Type, unsafe.Add(value.Addr().UnsafePointer(), field.Offset)).Interface()

		container, ok := fieldValue.(PointerToClass)
		if ok && container.Pointer() == 0 && field.Type.Implements(reflect.TypeOf([0]Singleton{}).Elem()) {
			var name = localCtx.StringName(strings.TrimPrefix(field.Type.Name(), "class"))
			singleton := godot.API.Object.GetSingleton(localCtx, name)
			container.SetPointer(mmm.Let[Pointer](godot.Lifetime, godot.API, singleton.Pointer()))
		}
	}
}

type classImplementation struct {
	Name  StringName
	Super StringName

	Godot *API
	Type  reflect.Type

	Value IsClass
}

var _ ClassInterface = classImplementation{}

func (class classImplementation) IsVirtual() bool {
	return false
}

func (class classImplementation) IsAbstract() bool {
	if class.Type.Kind() == reflect.Interface {
		return true
	}
	return false
}

func (class classImplementation) IsExposed() bool {
	return true
}

func (class classImplementation) CreateInstance() Object {
	ctx := NewContext(class.Godot)
	var super = class.Godot.ClassDB.ConstructObject(ctx, class.Super)
	super.SetPointer(mmm.Let[Pointer](ctx.Lifetime, ctx.API, mmm.End(super.AsPointer())))
	var value = reflect.New(class.Type)
	injectDependenciesInto(ctx, value.Elem())
	value.Interface().(PointerToClass).SetPointer(
		mmm.Let[Pointer](ctx.Lifetime, ctx.API, super.Pointer()))
	class.Godot.Object.SetInstance(super, class.Name, &instanceImplementation{
		object:  mmm.Get(super.AsPointer()),
		Context: ctx,
		Value:   value.Interface(),
	})
	class.Godot.instances[mmm.Get(super.AsPointer())] = value.Interface()
	return super
}

func (class classImplementation) GetVirtual(name StringName) any {
	ctx := NewContext(class.Godot)
	defer ctx.End()

	var Engine Engine
	Engine.SetPointer(class.Godot.Object.GetSingleton(ctx, ctx.StringName("Engine")).AsPointer())
	if Engine.IsEditorHint() {
		return nil
	}

	var virtual = VirtualByName(class.Value, name.String())
	if !virtual.IsValid() {
		return nil
	}
	var vtype = virtual.Type().In(0)
	GoName := convertName(name.String())

	method, ok := reflect.PtrTo(class.Type).MethodByName(GoName)
	if !ok {
		return nil
	}
	if method.Type.NumIn() != vtype.NumIn() {
		panic(fmt.Sprintf("gdextension.RegisterClass: Method %s.%s does not match %s.%s", class.Type.Name(), GoName, virtual.Type().Name(), name))
	}
	for i := 1; i < method.Type.NumIn(); i++ {
		if method.Type.In(i) != vtype.In(i) {
			panic(fmt.Sprintf("gdextension.RegisterClass: Method %s.%s does not match %s.%s", class.Type.Name(), GoName, virtual.Type().Name(), name))
		}
	}
	var copy = reflect.New(method.Type)
	copy.Elem().Set(method.Func)
	var fn = reflect.NewAt(vtype, copy.UnsafePointer()).Elem()
	return virtual.Call([]reflect.Value{fn, reflect.ValueOf(class.Godot)})[0].Interface()
}

type instanceImplementation struct {
	object  uintptr
	Context Context
	Value   any
}

func (instance *instanceImplementation) Set(name StringName, value Variant) bool {
	rvalue := reflect.ValueOf(instance.Value).Elem()
	field := rvalue.FieldByName(name.String())
	if !field.IsValid() {
		return false
	}
	field.Set(reflect.ValueOf(value.Interface(instance.Context)))
	return true
}

func (instance *instanceImplementation) Get(name StringName) (Variant, bool) {
	rvalue := reflect.ValueOf(instance.Value).Elem()
	field := rvalue.FieldByName(name.String())
	if !field.IsValid() {
		return Variant{}, false
	}
	return instance.Context.Variant(field.Interface()), true
}

func (instance *instanceImplementation) GetPropertyList(godot Context) []PropertyInfo {
	rtype := reflect.ValueOf(instance.Value).Elem().Type()
	var list = make([]PropertyInfo, 0, rtype.NumField())
	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		if !field.IsExported() || field.Anonymous {
			continue
		}
		var name = field.Name
		tag, ok := field.Tag.Lookup("gd")
		if ok {
			name = tag
		}
		list = append(list, PropertyInfo{
			Type:       variantTypeOf(field.Type),
			Name:       godot.StringName(name),
			ClassName:  godot.StringName(classNameOf(field.Type)),
			Hint:       0,
			HintString: godot.String(""),
			Usage:      2 | 4,
		})
	}
	return list
}

func (instance *instanceImplementation) PropertyCanRevert(name StringName) bool { return false }
func (instance *instanceImplementation) PropertyGetRevert(name StringName) Variant {
	return Variant{}
}
func (instance *instanceImplementation) ValidateProperty(name StringName, info PropertyInfo) bool {
	return true
}

func (instance *instanceImplementation) Notification(int32, bool) {}

func (instance *instanceImplementation) ToString() (String, bool) {
	return String{}, false
}

func (instance *instanceImplementation) Reference()   {}
func (instance *instanceImplementation) Unreference() {}

func (instance *instanceImplementation) CallVirtual(name StringName, virtual any, args UnsafeArgs, back UnsafeBack) {
	virtual.(ExtensionClassCallVirtualFunc)(instance.Value, args, back)
}

func (instance *instanceImplementation) GetRID() RID {
	return 0
}

func (instance *instanceImplementation) Free() {
	delete(instance.Context.API.instances, instance.object)
	instance.Context.End()
}
