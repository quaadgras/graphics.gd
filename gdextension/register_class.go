package gdextension

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"unsafe"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"grow.graphics/gd/objects"
	"grow.graphics/gd/objects/EditorPlugin"
	"grow.graphics/gd/objects/Engine"
	"grow.graphics/gd/objects/MainLoop"
	"grow.graphics/gd/objects/Node"
	"grow.graphics/gd/objects/ProjectSettings"
	"grow.graphics/gd/objects/Script"
	"grow.graphics/gd/objects/ScriptLanguage"
	"grow.graphics/gd/variant/StringName"

	gd "grow.graphics/gd/internal"
	classdb "grow.graphics/gd/internal/classdb"
	"grow.graphics/gd/internal/pointers"
)

// Tool can be embedded inside a struct to make it run in the editor.
type Tool interface{ tool() }

// Class can be embedded inside of a struct to represent a new Class type.
// The extended class will be available by calling the [Class.Super] method.
type Class[T any, S gd.IsClass] struct {
	gd.Class[T, S]
}

func (class *Class[T, S]) getObject() gd.Object {
	return *(*gd.Object)(unsafe.Pointer(class.Super()))
}

func (class *Class[T, S]) setObject(obj gd.Object) {
	*(*gd.Object)(unsafe.Pointer(class.Super())) = obj
}

type isClass interface {
	getObject() gd.Object
	setObject(gd.Object)
}

func (class *Class[T, S]) AsObject() gd.Object {
	obj := class.getObject()
	if obj == (gd.Object{}) {
		impl, ok := registered.Load(reflect.TypeFor[T]())
		if ok {
			instancer := impl.(*classImplementation)
			obj = instancer.CreateInstance()
			class.setObject(obj)
		}
	}
	return obj
}

var registered sync.Map

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
func Register[Struct gd.Extends[Parent], Parent gd.IsClass]() {
	var classType = reflect.TypeFor[Struct]()
	var superType = reflect.TypeFor[Parent]()
	if classType.Kind() != reflect.Struct || classType.Name() == "" {
		panic("gdextension.RegisterClass: Class type must be a named struct")
	}
	var rename = classNameOf(classType) // support 'gd' tag for renaming the class within Godot.
	var tool = false
	var super Parent
	switch any(super).(type) {
	case interface{ AsScript() Script.Instance },
		interface{ AsEditorPlugin() EditorPlugin.Instance },
		interface {
			AsScriptLanguage() ScriptLanguage.Instance
		}:
		tool = true
	}
	var reference Struct
	var className = pointers.Pin(StringName.New(rename))
	var superName = pointers.Pin(StringName.New(classNameOf(superType)))

	var impl = &classImplementation{
		Name:           className,
		Super:          superName,
		Type:           classType,
		Tool:           tool,
		VirtualMethods: reference.Virtual,
	}
	registered.Store(classType, impl)

	gd.Global.ClassDB.RegisterClass(gd.Global.ExtensionToken, className, superName, impl)
	gd.RegisterCleanup(func() {
		gd.Global.ClassDB.UnregisterClass(gd.Global.ExtensionToken, className)
		registered.Delete(classType)
		className.Free()
		superName.Free()
	})
	registerSignals(className, classType)
	registerMethods(className, classType)
	if _, ok := any(super).(interface {
		AsMainLoop() MainLoop.Instance
	}); ok {
		main_loop_type := gd.NewString("application/run/main_loop_type")
		ProjectSettings.Advanced().SetInitialValue(main_loop_type, gd.NewVariant(className))
		ProjectSettings.Advanced().SetSetting(main_loop_type, gd.NewVariant(className))
	}
	if registrator, ok := any(reference).(interface{ OnRegister() }); ok {
		registrator.OnRegister()
	}
}

func classNameOf(rtype reflect.Type) string {
	if rtype.Kind() == reflect.Ptr || rtype.Kind() == reflect.Array {
		return classNameOf(rtype.Elem())
	}
	if rtype.Kind() == reflect.Struct && rtype.NumField() > 0 {
		if rtype.Field(0).Anonymous {
			if rename, ok := rtype.Field(0).Tag.Lookup("gd"); ok {
				return rename
			}
			if rtype.Name() == "" {
				return classNameOf(rtype.Field(0).Type)
			}
		}
		return strings.TrimPrefix(rtype.Name(), "class")
	}
	return ""
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
	Name  gd.StringName
	Super gd.StringName

	Tool bool

	Type reflect.Type

	VirtualMethods func(string) reflect.Value
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

func (class classImplementation) CreateInstance() gd.Object {
	var super = gd.Global.ClassDB.ConstructObject(class.Super)
	instance := class.reloadInstance(super)
	gd.Global.Object.SetInstance(super, class.Name, instance)
	gd.Global.Object.SetInstanceBinding(super, gd.Global.ExtensionToken, nil, nil)
	instance.OnCreate()
	return super
}

func (class classImplementation) ReloadInstance(super gd.Object) gd.ObjectInterface {
	return class.reloadInstance(super)
}

func (class classImplementation) reloadInstance(super gd.Object) gd.ObjectInterface {
	var value = reflect.New(class.Type)
	extensionClass := value.Interface().(isClass)
	extensionClass.setObject(super)

	gd.ExtensionInstances.Store(pointers.Get(super)[0], extensionClass)

	value = value.Elem()

	// TODO cache this check
	var signals []signalChan
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
		if field.Type.Kind() == reflect.Chan && field.Type.ChanDir() == reflect.SendDir {
			signal := gd.NewSignalOf(super, gd.NewStringName(name))
			ch := reflect.MakeChan(reflect.ChanOf(reflect.BothDir, field.Type.Elem()), 0)
			rvalue.Elem().Set(ch)
			signals = append(signals, signalChan{
				signal: signal,
				rvalue: ch,
			})
		}
	}
	if len(signals) > 0 {
		go manageSignals(super.AsObject().GetInstanceId(), signals)
	}
	return &instanceImplementation{
		object:   pointers.Get(super)[0],
		Value:    value.Addr().Interface().(isClass),
		signals:  signals,
		isEditor: !class.Tool && Engine.IsEditorHint(),
	}
}

func (class classImplementation) GetVirtual(name gd.StringName) any {
	if !class.Tool && Engine.IsEditorHint() {
		return nil
	}
	var virtual = class.VirtualMethods(name.String())
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
	return virtual.Call([]reflect.Value{fn})[0].Interface()
}

type instanceImplementation struct {
	object   uintptr
	Value    isClass
	signals  []signalChan
	isEditor bool
}

var lastGC int

func (instance *instanceImplementation) setupForCall() func() {
	if frame := Engine.GetFramesDrawn(); frame > lastGC {
		pointers.MarkAndSweep()
	}
	return func() {}
}

func (instance *instanceImplementation) OnCreate() {
	if impl, ok := instance.Value.(interface {
		OnCreate()
	}); ok {
		defer instance.setupForCall()()
		impl.OnCreate()
	}
}

func (instance *instanceImplementation) Set(name gd.StringName, value gd.Variant) bool {
	if impl, ok := instance.Value.(interface {
		Set(gd.StringName, gd.Variant) gd.Bool
	}); ok {
		defer instance.setupForCall()()
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
	val := value.Interface()
	converted := reflect.ValueOf(val)
	if !converted.IsValid() {
		return false
	}
	if converted.Type().ConvertibleTo(field.Type()) {
		converted = converted.Convert(field.Type())
	}
	if !converted.Type().AssignableTo(field.Type()) {
		switch field.Type().Kind() {
		case reflect.String:
			s, ok := converted.Interface().(gd.String)
			if ok {
				field.SetString(s.String())
				return true
			}
		}
		return false
	}
	if obj, ok := val.(gd.IsClass); ok {
		ref, ok := gd.As[gd.RefCounted](obj.AsObject())
		if ok {
			ref.Reference()
		}
		//	field.Addr().Interface().(gd.PointerToClass).SetPointer(mmm.Pin[gd.Pointer](instance.Value.GetKeepAlive().Lifetime, instance.Value.GetKeepAlive().API, mmm.End(obj.AsObject().AsPointer())))
	} else {
		field.Set(converted)
	}
	if impl, ok := instance.Value.(interface {
		OnSet(gd.StringName, gd.Variant)
	}); ok {
		defer instance.setupForCall()()
		impl.OnSet(name, value)
	}
	return true
}

func (instance *instanceImplementation) Get(name gd.StringName) (gd.Variant, bool) {
	if impl, ok := instance.Value.(interface {
		Get(gd.StringName) gd.Variant
	}); ok {
		defer instance.setupForCall()()
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
	return gd.NewVariant(field.Interface()), true
}

func (instance *instanceImplementation) GetPropertyList() []gd.PropertyInfo {
	if impl, ok := instance.Value.(interface {
		GetPropertyList() []gd.PropertyInfo
	}); ok {
		defer instance.setupForCall()()
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
		list = append(list, propertyOf(field))
	}
	return list
}

func (instance *instanceImplementation) PropertyCanRevert(name gd.StringName) bool {
	if impl, ok := instance.Value.(interface {
		PropertyCanRevert(gd.StringName) gd.Bool
	}); ok {
		defer instance.setupForCall()()
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
func (instance *instanceImplementation) PropertyGetRevert(name gd.StringName) (gd.Variant, bool) {
	if impl, ok := instance.Value.(interface {
		PropertyGetRevert(gd.StringName) (gd.Variant, bool)
	}); ok {
		defer instance.setupForCall()()
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
	return gd.NewVariant(value.Elem().Interface()), true
}

func (instance *instanceImplementation) ValidateProperty(name gd.StringName, info *gd.PropertyInfo) bool {
	defer instance.setupForCall()()
	switch validate := instance.Value.(type) {
	case interface {
		ValidateProperty(gd.StringName, *gd.PropertyInfo) gd.Bool
	}:
		return bool(validate.ValidateProperty(name, info))
	}
	return true
}

func (instance *instanceImplementation) Notification(what int32, reversed bool) {
	if what == 13 { // NOTIFICATION_READY
		instance.ready()
	}
	if !instance.isEditor {
		defer instance.setupForCall()()
		switch notify := instance.Value.(type) {
		case interface{ Notification(gd.NotificationType) }:
			notify.Notification(gd.NotificationType(what))
		default:
		}
	}
}

func (instance *instanceImplementation) ToString() (gd.String, bool) {
	defer instance.setupForCall()()
	switch onfree := instance.Value.(type) {
	case interface{ ToString() gd.String }:
		return onfree.ToString(), true
	}
	return gd.String{}, false
}

func (instance *instanceImplementation) Reference()   {}
func (instance *instanceImplementation) Unreference() {}

func (instance *instanceImplementation) CallVirtual(name gd.StringName, virtual any, args gd.UnsafeArgs, back gd.UnsafeBack) {
	defer instance.setupForCall()()
	virtual.(gd.ExtensionClassCallVirtualFunc)(instance.Value, args, back)
}

func (instance *instanceImplementation) GetRID() gd.RID {
	return 0
}

func (instance *instanceImplementation) Free() {
	for _, signal := range instance.signals {
		signal.rvalue.Close()
	}
	gd.ExtensionInstances.Delete(instance.object)
	defer instance.setupForCall()()
	switch onfree := instance.Value.(type) {
	case interface{ OnFree() }:
		onfree.OnFree()
	}
}

// ready is responsible for asserting the scene tree for struct members that implement
// Super().AsNode() and asserting that these nodes are added as children to the Super.
//
// TODO this could be partially pre-compiled for a given [Register] type and cached in
// order to avoid any use of reflection at instantiation time.
func (instance *instanceImplementation) ready() {
	var parent objects.Node = objects.Node{classdb.Node(instance.Value.getObject())}

	var rvalue = reflect.ValueOf(instance.Value).Elem()
	for i := 0; i < rvalue.NumField(); i++ {
		field := rvalue.Type().Field(i)
		if !field.IsExported() || field.Name == "Class" {
			continue
		}
		instance.assertChild(rvalue.Field(i).Addr().Interface(), field, parent, parent)
	}
	if !instance.isEditor {
		defer instance.setupForCall()()
		switch ready := instance.Value.(type) {
		case interface{ Ready() }:
			ready.Ready()
		}
	}
}

func (instance *instanceImplementation) assertChild(value any, field reflect.StructField, parent, owner objects.Node) {
	type isNode interface {
		UnsafePointer() unsafe.Pointer
		AsObject() gd.Object
		AsNode() Node.Instance
	}
	nodeType := reflect.TypeOf([0]isNode{}).Elem()
	if !field.Type.Implements(nodeType) && !reflect.PointerTo(field.Type).Implements(nodeType) {
		return
	}
	var (
		rvalue = reflect.ValueOf(value)
	)
	if rvalue.Elem().Kind() == reflect.Pointer {
		rvalue.Elem().Set(reflect.New(rvalue.Elem().Type().Elem()))
		value = rvalue.Elem().Interface()
	}
	class := value.(isNode)
	if rvalue.Elem().Kind() == reflect.Struct {
		defer func() {
			rvalue := rvalue.Elem()
			for i := 0; i < rvalue.NumField(); i++ {
				field := rvalue.Type().Field(i)
				if !field.IsExported() || field.Name == "Class" || field.Anonymous {
					continue
				}
				instance.assertChild(rvalue.Field(i).Addr().Interface(), field, class.AsNode(), owner)
			}
		}()
	}
	name := field.Name
	if tag := field.Tag.Get("gd"); tag != "" {
		name = tag
	}
	path := gd.NewString(name).NodePath()
	if !Node.Advanced(parent).HasNode(path) {
		child := gd.Global.ClassDB.ConstructObject(gd.NewStringName(classNameOf(field.Type)))
		native, ok := gd.ExtensionInstances.Load(pointers.Get(child)[0])
		if ok {
			rvalue.Elem().Set(reflect.ValueOf(native))
			class = native.(isNode)
		}
		var mode Node.InternalMode = Node.InternalModeDisabled
		if !field.IsExported() {
			mode = Node.InternalModeFront
		}
		Node.Advanced(class.AsNode()).SetName(gd.NewString(field.Name))
		var adding objects.Node = objects.Node{classdb.Node(class.AsObject())}
		Node.Advanced(parent).AddChild(adding, true, mode)
		Node.Advanced(class.AsNode()).SetOwner(owner)
		return
	}
	var node = Node.Advanced(parent).GetNode(path)
	if name := node[0].AsObject().GetClass().String(); name != classNameOf(field.Type) {
		panic(fmt.Sprintf("gd.Register: Node %s.%s is not of type %s (%s)", rvalue.Type().Name(), field.Name, field.Type.Name(), name))
	}
	ref, native := gd.ExtensionInstances.Load(pointers.Get(node[0])[0])
	if native {
		rvalue.Elem().Set(reflect.ValueOf(ref))
		pointers.End(node[0])
	} else {
		*(*gd.Object)(class.UnsafePointer()) = pointers.New[gd.Object](pointers.Get(node[0]))
		pointers.End(node[0])
	}
}
