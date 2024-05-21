//go:build !generate

package gd

import (
	"fmt"
	"reflect"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	gd "grow.graphics/gd/internal"
	"runtime.link/mmm"
)

type onFree mmm.Pointer[func(), onFree, [0]uintptr]

func (cb onFree) Free() {
	(*mmm.API(cb))()
	mmm.End(cb)
}

/*
RegisterClass registers a struct available for use inside Godot
extending the given 'Parent' Godot class. The 'Struct' type must
be a named struct that embeds a [Class] field specifying the
parent class to extend.

	type MyClass struct {
		Class[MyClass, Node2D] `gd:"MyClass"`
	}

The tag can be adjusted in order to change the name of the class
within Godot.

Use this in a main or init function to register your Go structs
and they will become available within the Godot engine for use
in the editor and/or within scripts.

All exported fields and methods will be exposed to Godot, so
take caution when embedding types, as their fields and methods
will be promoted.

If the Struct extends [EditorPlugin] then it will be added to
the editor as a plugin.

If the Struct extends [MainLoop] or [SceneTree] then it will be
used as the main loop for the application.
*/
func Register[Struct gd.Extends[Parent], Parent gd.IsClass](godot Context) {
	var classType = reflect.TypeOf([0]Struct{}).Elem()
	var superType = reflect.TypeOf([0]Parent{}).Elem()
	if classType.Kind() != reflect.Struct || classType.Name() == "" {
		panic("gdextension.RegisterClass: Class type must be a named struct")
	}

	// Support 'gd' tag for renaming the class within Godot.
	var rename = classType.Name()
	for i := 0; i < classType.NumField(); i++ {
		field := classType.Field(i)
		if field.Name == "Class" {
			if val := field.Tag.Get("gd"); val != "" {
				rename = val
			}
		}
	}

	var className = godot.StringName(rename)
	var superName = godot.StringName(strings.TrimPrefix(superType.Name(), "class"))

	godot.API.ClassDB.RegisterClass(godot.API.ExtensionToken, className, superName,
		&classImplementation{
			Name:  className,
			Super: superName,
			Godot: godot.API,
			Type:  classType,
			Value: reflect.New(classType).Interface().(gd.IsClass),
		})
	unregister := func() {
		godot.API.ClassDB.UnregisterClass(godot.API.ExtensionToken, className)
	}
	mmm.Pin[onFree](godot.Lifetime, &unregister, [0]uintptr{})

	registerSignals(godot, className, classType)
	registerMethods(godot, className, classType)

	if superType.Implements(reflect.TypeOf([0]interface {
		AsMainLoop() MainLoop
	}{}).Elem()) {
		main_loop_type := godot.String("application/run/main_loop_type")
		ProjectSettings(godot).SetInitialValue(main_loop_type, godot.Variant(className))
		ProjectSettings(godot).SetSetting(main_loop_type, godot.Variant(className))
	}
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
		joins = append(joins, cases.Title(language.English).String(word))
	}
	return strings.Join(joins, "")
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
	return class.Type.Kind() == reflect.Interface
}

func (class classImplementation) IsExposed() bool {
	return true // TODO return false if the Go type is not exported.
}

func (class classImplementation) CreateInstance() Object {
	ctx := gd.NewContext(class.Godot)
	var super = class.Godot.ClassDB.ConstructObject(ctx, class.Super)
	super.SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, mmm.End(super.AsPointer())))
	instance := class.reloadInstance(ctx, super)
	class.Godot.Object.SetInstance(super, class.Name, instance)
	class.Godot.Object.SetInstanceBinding(super, ctx.API.ExtensionToken, nil, nil)
	return super
}

func (class classImplementation) ReloadInstance(super Object) gd.ObjectInterface {
	return class.reloadInstance(gd.NewContext(class.Godot), super)
}

func (class classImplementation) reloadInstance(ctx gd.Context, super Object) gd.ObjectInterface {
	var value = reflect.New(class.Type)
	value.Interface().(gd.PointerToClass).SetPointer(
		mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, mmm.Get(super.AsPointer())))
	class.Godot.Instances[mmm.Get(super.AsPointer())[0]] = value.Interface()

	value = value.Elem()

	// TODO cache this check
	for i := 0; i < value.NumField(); i++ {
		var field = value.Type().Field(i)
		if !field.IsExported() || field.Name == "Class" {
			continue
		}
		var (
			rvalue = value.Field(i).Addr()
		)
		name := field.Name
		if tag := field.Tag.Get("gd"); tag != "" {
			name = tag
		}
		// Signal fields need to have their values injected into the field, so that they can be used (emitted).
		if setter, ok := rvalue.Interface().(isSignal); ok {
			signal := ctx.SignalOf(ctx, super, ctx.StringName(name))
			scoped := mmm.Let[gd.Signal](ctx.Lifetime, ctx.API, mmm.End(signal))
			setter.setSignal(scoped)
			emit := rvalue.Elem().FieldByName("Emit")
			fnType := emit.Type()
			emit.Set(reflect.MakeFunc(fnType, func(args []reflect.Value) (results []reflect.Value) {
				tmp := gd.NewContext(ctx.API)
				var variants = make([]gd.Variant, 0, len(args))
				for _, arg := range args {
					variants = append(variants, tmp.Variant(arg.Interface()))
				}
				scoped.Emit(variants...)
				return nil
			}))
		}
	}
	return &instanceImplementation{
		object:  mmm.Get(super.AsPointer())[0],
		Context: ctx,
		Value:   value.Addr().Interface(),
	}
}

func (class classImplementation) GetVirtual(name StringName) any {
	ctx := gd.NewContext(class.Godot)
	defer ctx.End()

	if Engine(ctx).IsEditorHint() {
		return nil
	}

	var virtual = gd.VirtualByName(class.Value, name.String())
	if !virtual.IsValid() {
		return nil
	}
	var vtype = virtual.Type().In(0)
	GoName := convertName(name.String())
	if GoName == "Ready" {
		return nil // special case, as we override this method for all node types, so that we can assert the scene tree.
	}
	method, ok := reflect.PointerTo(class.Type).MethodByName(GoName)
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
	tmp := gd.NewContext(instance.Context.API)
	defer tmp.End()
	val := value.Interface(tmp)
	converted := reflect.ValueOf(val)
	if !converted.IsValid() {
		return false
	}
	if converted.Type().ConvertibleTo(field.Type()) {
		converted = converted.Convert(field.Type())
	}
	if !converted.Type().AssignableTo(field.Type()) {
		if field.Type().Implements(reflect.TypeOf([0]gd.IsArray{}).Elem()) {
			method, ok := field.Type().MethodByName("Typed")
			if ok {
				arr, isArray := val.(gd.Array)
				if !isArray {
					return false
				}
				arr = mmm.Pin[gd.Array](instance.Context.Lifetime, instance.Context.API, mmm.End(arr))
				array := reflect.New(method.Type.In(0)).Elem()
				array.Set(reflect.ValueOf(arr).Convert(method.Type.In(0)))
				field.Set(array)
				return true
			}
		}
		return false
	}
	if obj, ok := val.(gd.IsClass); ok {
		ref, ok := gd.As[gd.RefCounted](instance.Context, obj.AsObject())
		if ok {
			ref.Reference()
		}
		field.Addr().Interface().(gd.PointerToClass).SetPointer(mmm.Pin[gd.Pointer](instance.Context.Lifetime, instance.Context.API, mmm.End(obj.AsPointer())))
		return true
	}
	field.Set(converted)
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
		list = append(list, propertyOf(godot, field))
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
	if !Engine(instance.Context).IsEditorHint() {
		impl, ok := instance.Value.(interface {
			Notification(gd.Context, gd.NotificationType)
		})
		if ok {
			tmp := gd.NewContext(instance.Context.API)
			defer tmp.End()
			impl.Notification(tmp, gd.NotificationType(what))
		}
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
// TODO this could be partially pre-compiled for a given [Register] type and cached in
// order to avoid any use of reflection at instantiation time.
func (instance *instanceImplementation) ready() {
	tmp := gd.NewContext(instance.Context.API)
	defer tmp.End()

	var parent Node
	parent.SetPointer(mmm.Let[gd.Pointer](tmp.Lifetime, instance.Context.API, [2]uintptr{instance.object}))

	var rvalue = reflect.ValueOf(instance.Value).Elem()
	for i := 0; i < rvalue.NumField(); i++ {
		field := rvalue.Type().Field(i)
		if !field.IsExported() || field.Name == "Class" {
			continue
		}
		instance.assertChild(tmp, rvalue.Field(i).Addr().Interface(), field, parent)
	}
	if !Engine(tmp).IsEditorHint() {
		impl, ok := instance.Value.(interface {
			Ready(Context)
		})
		if ok {
			impl.Ready(tmp)
		}
	}
}

func (instance *instanceImplementation) assertChild(tmp Context, value any, field reflect.StructField, parent Node) {
	var (
		rvalue = reflect.ValueOf(value)
	)
	if rvalue.Elem().Kind() == reflect.Pointer {
		rvalue.Elem().Set(reflect.New(rvalue.Elem().Type().Elem()))
		value = rvalue.Elem().Interface()
	}
	if rvalue.Elem().Kind() == reflect.Struct {
		defer func() {
			rvalue := rvalue.Elem()
			for i := 0; i < rvalue.NumField(); i++ {
				field := rvalue.Type().Field(i)
				if !field.IsExported() || field.Name == "Class" {
					continue
				}
				instance.assertChild(tmp, rvalue.Field(i).Addr().Interface(), field, parent)
			}
		}()
	}
	type isNode interface {
		gd.PointerToClass

		AsNode() Node
	}
	class, ok := value.(isNode)
	if !ok {
		return
	}
	name := field.Name
	if tag := field.Tag.Get("gd"); tag != "" {
		name = tag
	}
	path := tmp.String(name).NodePath(tmp)
	if !parent.HasNode(path) {
		child := instance.Context.API.ClassDB.ConstructObject(instance.Context, tmp.StringName(classNameOf(field.Type)))
		native, ok := instance.Context.API.Instances[mmm.Get(child.AsPointer())[0]]
		if ok {
			rvalue.Elem().Set(reflect.ValueOf(native))
			class = native.(isNode)
		} else {
			class.SetPointer(mmm.Let[gd.Pointer](instance.Context.Lifetime, tmp.API, mmm.Get(child.AsPointer())))
		}
		child.SetPointer(mmm.Let[gd.Pointer](instance.Context.Lifetime, tmp.API, mmm.End(child.AsPointer())))
		var mode NodeInternalMode
		if !field.IsExported() {
			mode = NodeInternalModeFront
		}
		class.AsNode().SetName(tmp.String(field.Name))
		var adding Node
		adding.SetPointer(mmm.Pin[gd.Pointer](tmp.Lifetime, tmp.API, class.AsPointer().Pointer()))
		parent.AddChild(adding, false, mode)
		class.AsNode().SetOwner(parent)
		return
	}
	var node = parent.GetNode(tmp, path)
	if name := node.AsObject().GetClass(tmp).String(); name != classNameOf(field.Type) {
		panic(fmt.Sprintf("gd.Register: Node %s.%s is not of type %s (%s)", rvalue.Type().Name(), field.Name, field.Type.Name(), name))
	}
	ref, native := tmp.API.Instances[mmm.Get(node.AsPointer())[0]]
	if native {
		rvalue.Elem().Set(reflect.ValueOf(ref))
		mmm.End(node.AsPointer())
	} else {
		class.SetPointer(mmm.Let[gd.Pointer](instance.Context.Lifetime, tmp.API, mmm.End(node.AsPointer())))
	}
}
