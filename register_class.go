//go:build !generate

package gd

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	gd "grow.graphics/gd/internal"
	"runtime.link/api/call"
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
func Register[Struct gd.Extends[Parent], Parent gd.IsClass](godot Context) {
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
			Value: reflect.New(classType).Interface().(gd.IsClass),
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
//
// for node types, that implement Super().AsNode(), this function
// will assert that any member nodes are children of the node on
// ready.
func injectDependenciesInto(godot Context, value reflect.Value) {
	if value.CanAddr() && value.Kind() != reflect.Struct {
		panic("gdextension.injectDependenciesInto: value must be an addressable struct")
	}
	localCtx := gd.NewContext(godot.API)
	defer localCtx.End()

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)

		// support private fields.
		fieldValue := reflect.NewAt(field.Type, unsafe.Add(value.Addr().UnsafePointer(), field.Offset)).Interface()

		container, ok := fieldValue.(gd.PointerToClass)
		if ok && mmm.Get(container.AsPointer()) == 0 && field.Type.Implements(reflect.TypeOf([0]gd.Singleton{}).Elem()) {
			var name = localCtx.StringName(strings.TrimPrefix(field.Type.Name(), "class"))
			singleton := godot.API.Object.GetSingleton(localCtx, name)
			container.SetPointer(mmm.Let[gd.Pointer](godot.Lifetime, godot.API, mmm.Get(singleton.AsPointer())))
		}
	}
}

type classImplementation struct {
	Name  StringName
	Super StringName

	Godot *gd.API
	Type  reflect.Type

	Value gd.IsClass
}

var _ gd.ClassInterface = classImplementation{}

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
	ctx := gd.NewContext(class.Godot)
	var super = class.Godot.ClassDB.ConstructObject(ctx, class.Super)
	super.SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, mmm.End(super.AsPointer())))
	var value = reflect.New(class.Type)
	injectDependenciesInto(ctx, value.Elem())
	value.Interface().(gd.PointerToClass).SetPointer(
		mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, mmm.Get(super.AsPointer())))
	class.Godot.Object.SetInstance(super, class.Name, &instanceImplementation{
		object:  mmm.Get(super.AsPointer()),
		Context: ctx,
		Value:   value.Interface(),
	})
	class.Godot.Instances[mmm.Get(super.AsPointer())] = value.Interface()
	return super
}

func (class classImplementation) GetVirtual(name StringName) any {
	ctx := gd.NewContext(class.Godot)
	defer ctx.End()

	var Engine Object
	Engine.SetPointer(class.Godot.Object.GetSingleton(ctx, ctx.StringName("Engine")).AsPointer())

	is_editor_hint := func() bool {
		var selfPtr = Engine.AsPointer()
		var frame = call.New()
		var r_ret = call.Ret[bool](frame)
		mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_is_editor_hint, Engine.AsObject(), frame.Array(0), r_ret.Uintptr())
		var ret = r_ret.Get()
		frame.Free()
		return ret
	}
	if is_editor_hint() {
		return nil
	}

	var virtual = gd.VirtualByName(class.Value, name.String())
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

func (instance *instanceImplementation) Set(name StringName, value gd.Variant) bool {
	rvalue := reflect.ValueOf(instance.Value).Elem()
	field := rvalue.FieldByName(name.String())
	if !field.IsValid() {
		return false
	}
	field.Set(reflect.ValueOf(value.Interface(instance.Context)))
	return true
}

func (instance *instanceImplementation) Get(name StringName) (gd.Variant, bool) {
	rvalue := reflect.ValueOf(instance.Value).Elem()
	field := rvalue.FieldByName(name.String())
	if !field.IsValid() {
		return gd.Variant{}, false
	}
	return instance.Context.Variant(field.Interface()), true
}

func (instance *instanceImplementation) GetPropertyList(godot Context) []gd.PropertyInfo {
	rtype := reflect.ValueOf(instance.Value).Elem().Type()
	var list = make([]gd.PropertyInfo, 0, rtype.NumField())
	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		if !field.IsExported() || field.Anonymous {
			continue
		}
		if _, ok := field.Type.MethodByName("AsNode"); ok {
			continue
		}
		var name = field.Name
		tag, ok := field.Tag.Lookup("gd")
		if ok {
			name = tag
		}
		list = append(list, gd.PropertyInfo{
			Type:       variantTypeOf(field.Type),
			Name:       godot.StringName(name),
			ClassName:  godot.StringName(classNameOf(field.Type)),
			Hint:       PropertyHintResourceType,
			HintString: godot.String(classNameOf(field.Type)),
			Usage:      PropertyUsageStorage | PropertyUsageEditor,
		})
	}
	return list
}

func (instance *instanceImplementation) PropertyCanRevert(name StringName) bool { return false }
func (instance *instanceImplementation) PropertyGetRevert(name StringName) gd.Variant {
	return gd.Variant{}
}
func (instance *instanceImplementation) ValidateProperty(name StringName, info gd.PropertyInfo) bool {
	return true
}

func (instance *instanceImplementation) Notification(what int32, reversed bool) {
	if what == 13 { // NOTIFICATION_READY
		instance.ready()
	}
}

func (instance *instanceImplementation) ToString() (String, bool) {
	return String{}, false
}

func (instance *instanceImplementation) Reference()   {}
func (instance *instanceImplementation) Unreference() {}

func (instance *instanceImplementation) CallVirtual(name StringName, virtual any, args gd.UnsafeArgs, back gd.UnsafeBack) {
	virtual.(gd.ExtensionClassCallVirtualFunc)(instance.Value, args, back)
}

func (instance *instanceImplementation) GetRID() RID {
	return 0
}

func (instance *instanceImplementation) Free() {
	delete(instance.Context.API.Instances, instance.object)
	instance.Context.End()
}

// ready is responsible for asserting the scene tree for struct members that implement
// Super().AsNode() and asserting that these nodes are added as children to the Super.
//
// TODO this could be pre-compiled for a given [Register] type and cached in order to
// avoid any use of reflection at instantiation time.
func (instance *instanceImplementation) ready() {
	tmp := gd.NewContext(instance.Context.API)
	defer tmp.End()

	var parent Node
	parent.SetPointer(mmm.Let[gd.Pointer](tmp.Lifetime, instance.Context.API, instance.object))

	var rvalue = reflect.ValueOf(instance.Value).Elem()
	for i := 0; i < rvalue.NumField(); i++ {
		var field = rvalue.Type().Field(i)
		var value = reflect.NewAt(field.Type, unsafe.Add(rvalue.Field(i).Addr().UnsafePointer(), field.Offset)).Interface() // support unexported fields.
		if field.Name == "Class" {
			continue
		}
		class, ok := value.(interface {
			gd.PointerToClass

			AsNode() Node
		})
		if !ok {
			continue
		}
		path := tmp.String(field.Name).NodePath(tmp)
		if !parent.HasNode(path) {
			child := instance.Context.API.ClassDB.ConstructObject(instance.Context, tmp.StringName(classNameOf(field.Type)))
			class.SetPointer(mmm.Let[gd.Pointer](instance.Context.Lifetime, tmp.API, mmm.Get(child.AsPointer())))
			child.SetPointer(mmm.Let[gd.Pointer](instance.Context.Lifetime, tmp.API, mmm.End(child.AsPointer())))
			var mode NodeInternalMode
			if !field.IsExported() {
				mode = NodeInternalModeFront
			}
			class.AsNode().SetName(tmp.String(field.Name))
			parent.AddChild(class.AsNode(), false, mode)
			class.AsNode().SetOwner(parent)
			continue
		}
		var node = parent.GetNode(tmp, path)
		if name := node.AsObject().GetClass(tmp).String(); name != classNameOf(field.Type) {
			panic(fmt.Sprintf("gd.Register: Node %s.%s is not of type %s (%s)", rvalue.Type().Name(), field.Name, field.Type.Name(), name))
		}
		class.SetPointer(mmm.Let[gd.Pointer](instance.Context.Lifetime, tmp.API, mmm.End(node.AsPointer())))
	}
}
