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

If the Struct implements an OnRegister(Lifetime) method, it will
be called on a temporary instance when the class is registered.
*/
func Register[Struct gd.Extends[Parent], Parent gd.IsClass](godot Lifetime) {
	var classType = reflect.TypeOf([0]Struct{}).Elem()
	var superType = reflect.TypeOf([0]Parent{}).Elem()
	if classType.Kind() != reflect.Struct || classType.Name() == "" {
		panic("gdextension.RegisterClass: Class type must be a named struct")
	}

	// Support 'gd' tag for renaming the class within Godot.
	var rename = classNameOf(classType)

	var tool = false
	switch {
	case superType.Implements(reflect.TypeOf([0]interface{ AsScript() Script }{}).Elem()),
		superType.Implements(reflect.TypeOf([0]interface{ AsEditorPlugin() EditorPlugin }{}).Elem()),
		superType.Implements(reflect.TypeOf([0]interface{ AsScriptLanguage() ScriptLanguage }{}).Elem()):
		tool = true
	}

	var className = godot.StringName(rename)
	var superName = godot.StringName(strings.TrimPrefix(superType.Name(), "class"))

	godot.API.ClassDB.RegisterClass(godot.API.ExtensionToken, className, superName,
		&classImplementation{
			Name:  className,
			Super: superName,
			Godot: godot.API,
			Type:  classType,
			Tool:  tool,
			Value: reflect.New(classType).Interface().(gd.ExtensionClass),
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

	type onRegister interface {
		OnRegister(godot Lifetime)
	}
	if reflect.PointerTo(classType).Implements(reflect.TypeOf([0]onRegister{}).Elem()) {
		impl := reflect.New(classType).Interface().(onRegister)
		impl.OnRegister(godot)
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

	Tool bool

	Godot *gd.API
	Type  reflect.Type

	Value gd.ExtensionClass
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
	ctx := gd.NewLifetime(class.Godot)
	var super = class.Godot.ClassDB.ConstructObject(ctx, class.Super)
	super.SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, mmm.End(super.AsPointer())))
	instance := class.reloadInstance(ctx, super)
	class.Godot.Object.SetInstance(super, class.Name, instance)
	class.Godot.Object.SetInstanceBinding(super, ctx.API.ExtensionToken, nil, nil)
	instance.OnCreate()
	return super
}

func (class classImplementation) ReloadInstance(super Object) gd.ObjectInterface {
	return class.reloadInstance(gd.NewLifetime(class.Godot), super)
}

func (class classImplementation) reloadInstance(ctx gd.Lifetime, super Object) gd.ObjectInterface {
	var value = reflect.New(class.Type)
	extensionClass := value.Interface().(gd.ExtensionClass)

	extensionClass.SetPointer(
		mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, mmm.Get(super.AsPointer())))
	extensionClass.SetKeepAlive(ctx)

	class.Godot.Instances[mmm.Get(super.AsPointer())[0]] = extensionClass

	value = value.Elem()

	// TODO cache this check
	var signals []signalChan
	var thread = gd.NewLifetime(ctx.API)
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
				tmp := gd.NewLifetime(ctx.API)
				defer tmp.End()
				var variants = make([]gd.Variant, 0, len(args))
				for _, arg := range args {
					variants = append(variants, tmp.Variant(arg.Interface()))
				}
				scoped.Emit(variants...)
				return nil
			}))
		}
		if field.Type.Kind() == reflect.Chan && field.Type.ChanDir() == reflect.SendDir {
			signal := ctx.SignalOf(ctx, super, ctx.StringName(name))
			scoped := mmm.Let[gd.Signal](thread.Lifetime, thread.API, mmm.End(signal))
			ch := reflect.MakeChan(reflect.ChanOf(reflect.BothDir, field.Type.Elem()), 0)
			rvalue.Elem().Set(ch)
			signals = append(signals, signalChan{
				signal: scoped,
				rvalue: ch,
			})
		}
	}
	if len(signals) > 0 {
		go manageSignals(thread, super.AsObject().GetInstanceId(), signals)
	}
	return &instanceImplementation{
		object:  mmm.Get(super.AsPointer())[0],
		Value:   value.Addr().Interface().(gd.ExtensionClass),
		signals: signals,
	}
}

func (class classImplementation) GetVirtual(name StringName) any {
	ctx := gd.NewLifetime(class.Godot)
	defer ctx.End()

	if !class.Tool && Engine(ctx).IsEditorHint() {
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
	// legacy support for [gd.Context] parameter, which is now optional (and discouraged)
	// we generate a very slow function for compatibility reasons.
	if method.Type.NumIn() > 1 {
		if method.Type.In(1) == reflect.TypeOf(ctx) {
			var args = []reflect.Type{method.Type.In(0)}
			for i := 2; i < method.Type.NumIn(); i++ {
				args = append(args, method.Type.In(i))
			}
			var rets = []reflect.Type{}
			for i := 0; i < method.Type.NumOut(); i++ {
				rets = append(rets, method.Type.Out(i))
			}
			method.Type = reflect.FuncOf(args, rets, method.Type.IsVariadic())
			var old = method.Func
			method.Func = reflect.MakeFunc(method.Type, func(args []reflect.Value) (results []reflect.Value) {
				instance := args[0].Interface().(gd.ExtensionClass)
				godot := reflect.ValueOf(instance.GetKeepAlive())
				args = append([]reflect.Value{args[0], godot}, args[1:]...)
				return old.Call(args)
			})
		}
	}
	if method.Type.NumIn() != vtype.NumIn() {
		panic(fmt.Sprintf("gdextension.RegisterClass: Method %s.%s does not match %s.%s\nis %s want %s", class.Type.Name(), GoName, virtual.Type().Name(), name, method.Type, vtype))
	}
	for i := 1; i < method.Type.NumIn(); i++ {
		if method.Type.In(i) != vtype.In(i) {
			panic(fmt.Sprintf("gdextension.RegisterClass: Method %s.%s does not match %s.%s\nis %s want %s", class.Type.Name(), GoName, virtual.Type().Name(), name, method.Type, vtype))
		}
	}
	var copy = reflect.New(method.Type)
	copy.Elem().Set(method.Func)
	var fn = reflect.NewAt(vtype, copy.UnsafePointer()).Elem()
	return virtual.Call([]reflect.Value{fn, reflect.ValueOf(class.Godot)})[0].Interface()
}

type instanceImplementation struct {
	object  uintptr
	Value   gd.ExtensionClass
	signals []signalChan
}

func (instance instanceImplementation) OnCreate() {
	tmp := NewLifetime(instance.Value.GetKeepAlive())
	defer tmp.End()
	if impl, ok := instance.Value.(interface {
		OnCreate()
	}); ok {
		instance.Value.SetTemporary(tmp)
		impl.OnCreate()
	}
}

func (instance *instanceImplementation) Set(name StringName, value gd.Variant) bool {
	if impl, ok := instance.Value.(interface {
		Set(gd.StringName, gd.Variant) gd.Bool
	}); ok {
		tmp := NewLifetime(instance.Value.GetKeepAlive())
		defer tmp.End()
		instance.Value.SetTemporary(tmp)
		ok := bool(impl.Set(name, value))
		if ok {
			if impl, ok := instance.Value.(interface {
				OnSet(gd.StringName, gd.Variant)
			}); ok {
				impl.OnSet(name, value)
			}
		}
		return ok
	}
	sname := name.String()
	rvalue := reflect.ValueOf(instance.Value).Elem()
	field := rvalue.FieldByName(sname)
	if !field.IsValid() {
		for i := 0; i < rvalue.NumField(); i++ {
			if rvalue.Type().Field(i).Tag.Get("gd") == sname {
				field = rvalue.Field(i)
				break
			}
		}
		if !field.IsValid() {
			return false
		}
	}
	if !field.CanSet() {
		return false
	}
	val := value.Interface(instance.Value.GetKeepAlive())
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
				arr = mmm.Pin[gd.Array](instance.Value.GetKeepAlive().Lifetime, instance.Value.GetKeepAlive().API, mmm.End(arr))
				array := reflect.New(method.Type.In(0)).Elem()
				array.Set(reflect.ValueOf(arr).Convert(method.Type.In(0)))
				field.Set(array)
				return true
			}
		}
		return false
	}
	if obj, ok := val.(gd.IsClass); ok {
		ref, ok := gd.As[gd.RefCounted](instance.Value.GetKeepAlive(), obj.AsObject())
		if ok {
			ref.Reference()
		}
		field.Addr().Interface().(gd.PointerToClass).SetPointer(mmm.Pin[gd.Pointer](instance.Value.GetKeepAlive().Lifetime, instance.Value.GetKeepAlive().API, mmm.End(obj.AsPointer())))
	} else {
		field.Set(converted)
	}
	if impl, ok := instance.Value.(interface {
		OnSet(gd.StringName, gd.Variant)
	}); ok {
		tmp := NewLifetime(instance.Value.GetKeepAlive())
		defer tmp.End()
		instance.Value.SetTemporary(tmp)
		impl.OnSet(name, value)
	}
	return true
}

func (instance *instanceImplementation) Get(name StringName) (gd.Variant, bool) {
	if impl, ok := instance.Value.(interface {
		Get(StringName) gd.Variant
	}); ok {
		tmp := NewLifetime(instance.Value.GetKeepAlive())
		defer tmp.End()
		instance.Value.SetTemporary(tmp)
		return impl.Get(name), true
	}
	sname := name.String()
	rvalue := reflect.ValueOf(instance.Value).Elem()
	field := rvalue.FieldByName(sname)
	if !field.IsValid() {
		for i := 0; i < rvalue.NumField(); i++ {
			rfield := rvalue.Type().Field(i)
			if !rfield.Anonymous && rfield.Tag.Get("gd") == sname {
				field = rvalue.Field(i)
				break
			}
		}
		if !field.IsValid() {
			return gd.Variant{}, false
		}
	}
	if reflect.PointerTo(field.Type()).Implements(reflect.TypeOf([0]isSignal{}).Elem()) {
		return gd.Variant{}, false
	}
	return instance.Value.GetKeepAlive().Variant(field.Interface()), true
}

func (instance *instanceImplementation) GetPropertyList(godot Lifetime) []gd.PropertyInfo {
	if impl, ok := instance.Value.(interface {
		GetPropertyList() []gd.PropertyInfo
	}); ok {
		instance.Value.SetTemporary(godot)
		return impl.GetPropertyList()
	}
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

func (instance *instanceImplementation) PropertyCanRevert(name StringName) bool {
	if impl, ok := instance.Value.(interface {
		PropertyCanRevert(gd.StringName) gd.Bool
	}); ok {
		tmp := NewLifetime(instance.Value.GetKeepAlive())
		defer tmp.End()
		instance.Value.SetTemporary(tmp)
		return bool(impl.PropertyCanRevert(name))
	}
	sname := name.String()
	rtype := reflect.TypeOf(instance.Value).Elem()
	field, ok := rtype.FieldByName(sname)
	if !ok {
		for i := 0; i < rtype.NumField(); i++ {
			if rtype.Field(i).Tag.Get("gd") == sname {
				field = rtype.Field(i)
				ok = true
				break
			}
		}
	}
	if !ok {
		return false
	}
	_, ok = field.Tag.Lookup("default")
	return ok
}
func (instance *instanceImplementation) PropertyGetRevert(godot Lifetime, name StringName) (gd.Variant, bool) {
	tmp := NewLifetime(instance.Value.GetKeepAlive())
	defer tmp.End()
	if impl, ok := instance.Value.(interface {
		PropertyGetRevert(gd.StringName) (gd.Variant, bool)
	}); ok {
		instance.Value.SetTemporary(tmp)
		return impl.PropertyGetRevert(name)
	}
	sname := name.String()
	rtype := reflect.TypeOf(instance.Value).Elem()
	field, ok := rtype.FieldByName(sname)
	if !ok {
		for i := 0; i < rtype.NumField(); i++ {
			if rtype.Field(i).Tag.Get("gd") == sname {
				field = rtype.Field(i)
				ok = true
				break
			}
		}
	}
	if !ok {
		return gd.Variant{}, false
	}
	var value = reflect.New(field.Type)
	if tag := field.Tag.Get("default"); tag != "" {
		_, err := fmt.Sscanf(tag, "%v", value.Interface())
		if err != nil {
			return gd.Variant{}, false
		}
	}
	return godot.Variant(value.Elem().Interface()), true
}

func (instance *instanceImplementation) ValidateProperty(name StringName, info *gd.PropertyInfo) bool {
	tmp := NewLifetime(instance.Value.GetKeepAlive())
	defer tmp.End()
	instance.Value.SetTemporary(tmp)
	switch validate := instance.Value.(type) {
	case interface {
		ValidateProperty(gd.StringName, *gd.PropertyInfo) gd.Bool
	}:
		return bool(validate.ValidateProperty(name, info))
	case interface {
		ValidateProperty(gd.Lifetime, gd.StringName, *gd.PropertyInfo) gd.Bool
	}:
		return bool(validate.ValidateProperty(tmp, name, info))
	}
	return true
}

func (instance *instanceImplementation) Notification(what int32, reversed bool) {
	if what == 13 { // NOTIFICATION_READY
		instance.ready()
	}
	if !Engine(instance.Value.GetKeepAlive()).IsEditorHint() {
		tmp := NewLifetime(instance.Value.GetKeepAlive())
		defer tmp.End()
		instance.Value.SetTemporary(tmp)
		switch notify := instance.Value.(type) {
		case interface{ Notification(gd.NotificationType) }:
			notify.Notification(gd.NotificationType(what))
		case interface {
			Notification(gd.Lifetime, gd.NotificationType)
		}:
			notify.Notification(tmp, gd.NotificationType(what))
		}
	}
}

func (instance *instanceImplementation) ToString() (String, bool) {
	tmp := NewLifetime(instance.Value.GetKeepAlive())
	defer tmp.End()
	instance.Value.SetTemporary(tmp)
	switch onfree := instance.Value.(type) {
	case interface{ ToString() gd.String }:
		return onfree.ToString(), true
	case interface{ ToString(gd.Lifetime) gd.String }:
		return onfree.ToString(tmp), true
	}
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
	for _, signal := range instance.signals {
		signal.rvalue.Close()
	}
	delete(instance.Value.GetKeepAlive().API.Instances, instance.object)
	instance.Value.SetTemporary(instance.Value.GetKeepAlive())
	switch onfree := instance.Value.(type) {
	case interface{ OnFree() }:
		onfree.OnFree()
	case interface{ OnFree(gd.Lifetime) }:
		onfree.OnFree(instance.Value.GetKeepAlive())
	}
	instance.Value.GetKeepAlive().End()
}

// ready is responsible for asserting the scene tree for struct members that implement
// Super().AsNode() and asserting that these nodes are added as children to the Super.
//
// TODO this could be partially pre-compiled for a given [Register] type and cached in
// order to avoid any use of reflection at instantiation time.
func (instance *instanceImplementation) ready() {
	tmp := NewLifetime(instance.Value.GetKeepAlive())
	defer tmp.End()

	var parent Node
	parent.SetPointer(mmm.Let[gd.Pointer](tmp.Lifetime, instance.Value.GetKeepAlive().API, [2]uintptr{instance.object}))

	var rvalue = reflect.ValueOf(instance.Value).Elem()
	for i := 0; i < rvalue.NumField(); i++ {
		field := rvalue.Type().Field(i)
		if !field.IsExported() || field.Name == "Class" {
			continue
		}
		instance.assertChild(tmp, rvalue.Field(i).Addr().Interface(), field, parent, parent)
	}
	if !Engine(tmp).IsEditorHint() {
		instance.Value.SetTemporary(tmp)
		switch ready := instance.Value.(type) {
		case interface{ Ready() }:
			ready.Ready()
		case interface{ Ready(gd.Lifetime) }:
			ready.Ready(tmp)
		}
	}
}

func (instance *instanceImplementation) assertChild(tmp Lifetime, value any, field reflect.StructField, parent, owner Node) {
	var (
		rvalue = reflect.ValueOf(value)
	)
	if rvalue.Elem().Kind() == reflect.Pointer {
		rvalue.Elem().Set(reflect.New(rvalue.Elem().Type().Elem()))
		value = rvalue.Elem().Interface()
	}
	type isNode interface {
		gd.PointerToClass

		AsNode() Node
	}
	class, ok := value.(isNode)
	if !ok {
		return
	}
	if rvalue.Elem().Kind() == reflect.Struct {
		defer func() {
			rvalue := rvalue.Elem()
			for i := 0; i < rvalue.NumField(); i++ {
				field := rvalue.Type().Field(i)
				if !field.IsExported() || field.Name == "Class" || field.Anonymous {
					continue
				}
				instance.assertChild(tmp, rvalue.Field(i).Addr().Interface(), field, class.AsNode(), owner)
			}
		}()
	}
	name := field.Name
	if tag := field.Tag.Get("gd"); tag != "" {
		name = tag
	}
	path := tmp.String(name).NodePath(tmp)
	if !parent.HasNode(path) {
		child := instance.Value.GetKeepAlive().API.ClassDB.ConstructObject(instance.Value.GetKeepAlive(), tmp.StringName(classNameOf(field.Type)))
		native, ok := instance.Value.GetKeepAlive().API.Instances[mmm.Get(child.AsPointer())[0]]
		if ok {
			rvalue.Elem().Set(reflect.ValueOf(native))
			class = native.(isNode)
		} else {
			class.SetPointer(mmm.Let[gd.Pointer](instance.Value.GetKeepAlive().Lifetime, tmp.API, mmm.Get(child.AsPointer())))
		}
		child.SetPointer(mmm.Let[gd.Pointer](instance.Value.GetKeepAlive().Lifetime, tmp.API, mmm.End(child.AsPointer())))
		var mode NodeInternalMode = NodeInternalModeDisabled
		if !field.IsExported() {
			mode = NodeInternalModeFront
		}
		class.AsNode().SetName(tmp.String(field.Name))
		var adding Node
		adding.SetPointer(mmm.Pin[gd.Pointer](tmp.Lifetime, tmp.API, class.AsPointer().Pointer()))
		parent.AddChild(adding, true, mode)
		class.AsNode().SetOwner(owner)
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
		class.SetPointer(mmm.Let[gd.Pointer](instance.Value.GetKeepAlive().Lifetime, tmp.API, mmm.End(node.AsPointer())))
	}
}
