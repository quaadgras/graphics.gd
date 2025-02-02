package classdb

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"unsafe"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	EditorPluginClass "graphics.gd/classdb/EditorPlugin"
	EngineClass "graphics.gd/classdb/Engine"
	NodeClass "graphics.gd/classdb/Node"
	ScriptClass "graphics.gd/classdb/Script"
	ScriptLanguageClass "graphics.gd/classdb/ScriptLanguage"
	ShaderMaterialClass "graphics.gd/classdb/ShaderMaterial"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Path"
	"graphics.gd/variant/String"

	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
)

// Tool can be embedded inside a struct to make it run in the editor.
type Tool interface{ tool() }

// Extension can be embedded inside of a struct to represent a new Extension type.
// The extended class will be available by calling the [Extension.Super] method.
type Extension[T Class, S gd.IsClass] struct {
	gd.Class[T, S]
}

// NameOf returns the defined name for the given [Extension]-embedding type.
func NameOf(T Class) string {
	return nameOf(reflect.TypeOf(T))
}

func NameFor[T Class]() string {
	return nameOf(reflect.TypeFor[T]())
}

func (class Extension[T, S]) super() S {
	return class.Super()
}

func (class *Extension[T, S]) Super() S {
	class.AsObject()
	return *class.Class.Super()
}

func (class Extension[T, S]) getObject() [1]gd.Object {
	return *(*[1]gd.Object)(unsafe.Pointer(class.Class.Super()))
}

func (class *Extension[T, S]) setObject(obj [1]gd.Object) {
	*(*[1]gd.Object)(unsafe.Pointer(class.Class.Super())) = obj
}

func (class Extension[T, S]) superType() reflect.Type {
	return reflect.TypeFor[S]()
}

type isClass interface {
	gd.IsClass

	getObject() [1]gd.Object
	setObject([1]gd.Object)

	superType() reflect.Type
}

type Class interface {
	superType() reflect.Type
	getObject() [1]gd.Object
	Virtual(string) reflect.Value
}

func (class *Extension[T, S]) AsObject() [1]gd.Object {
	obj := class.getObject()
	if obj == ([1]gd.Object{}) {
		impl, ok := registered.Load(reflect.TypeFor[T]())
		if !ok {
			Register[T]()
			impl, ok = registered.Load(reflect.TypeFor[T]())
		}
		if ok {
			instancer := impl.(*classImplementation)
			obj = instancer.createInstance(reflect.NewAt(reflect.TypeFor[T](), unsafe.Pointer(class)))
			class.setObject(obj)
		}
	}
	pointers.Get(obj[0])
	return obj
}

var registered sync.Map

/*
Register registers a struct available for use inside Godot
extending the given 'Parent' Godot class. The 'Struct' type must
be a named struct that embeds a [Extension] field specifying the
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

If the Struct extends [EditorPluginClass] then it will be added to
the editor as a plugin.

If the Struct extends [MainLoopClass] or [SceneTree] then it will be
used as the main loop for the application.

If the Struct implements an OnRegister(Lifetime) method, it will
be called on a temporary instance when the class is registered.
*/
func Register[T Class]() {
	register := func() {
		var classType = reflect.TypeFor[T]()
		var base = classType
		for base.Field(0).Anonymous {
			if base.Field(0).Name == "Class" {
				break
			}
			base = base.Field(0).Type
		}
		if !base.Implements(reflect.TypeFor[Class]()) {
			panic("gdextension.RegisterClass: Class type must embed a gd.Extension field as the first field")
		}
		var superType = ([1]T{})[0].superType()
		if classType.Kind() != reflect.Struct || classType.Name() == "" {
			panic("gdextension.RegisterClass: Class type must be a named struct")
		}
		var rename = nameOf(classType) // support 'gd' tag for renaming the class within Godot.
		var tool = false
		var super = reflect.New(superType).Elem().Interface()
		switch super.(type) {
		case interface{ AsScript() ScriptClass.Instance },
			interface {
				AsEditorPlugin() EditorPluginClass.Instance
			},
			interface {
				AsScriptLanguage() ScriptLanguageClass.Instance
			}:
			tool = true
		}
		switch any(([1]T{})[0]).(type) {
		case Tool:
			tool = true
		}
		var reference T
		var className = pointers.Pin(gd.NewStringName(rename))
		var superName = pointers.Pin(gd.NewStringName(nameOf(superType)))

		var impl = &classImplementation{
			Name:           className,
			Super:          superName,
			Type:           classType,
			Tool:           tool,
			VirtualMethods: reference.Virtual,
		}
		registered.Store(classType, impl)

		gd.Global.ClassDB.RegisterClass(gd.Global.ExtensionToken, className, superName, impl)
		registerClassInformation(className, rename, nameOf(superType), classType)
		gd.RegisterCleanup(func() {
			gd.Global.ClassDB.UnregisterClass(gd.Global.ExtensionToken, className)
			registered.Delete(classType)
			className.Free()
			superName.Free()
		})
		switch super.(type) {
		case interface {
			AsShaderMaterial() ShaderMaterialClass.Instance
		}:
		default:
			registerSignals(className, classType)
			registerMethods(className, classType)
		}

		if registrator, ok := any(reference).(interface{ OnRegister() }); ok {
			registrator.OnRegister()
		}
	}
	if gd.Linked {
		register()
	} else {
		gd.StartupFunctions = append(gd.StartupFunctions, register)
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

func registerClassInformation(className gd.StringName, classNameString string, inherits string, rtype reflect.Type) {
	var class xmlDocumentation
	class.Name = classNameString
	class.Inherits = inherits
	class.Version = "4.0"
	extractDocTag := func(tag reflect.StructTag) string {
		_, docs, _ := strings.Cut(string(tag), "\n")
		docs = strings.Replace(docs, "\t", "", -1)
		return strings.TrimSpace(docs)
	}
	if rtype.Field(0).Anonymous {
		docs := extractDocTag(rtype.Field(0).Tag)
		brief, whole, _ := strings.Cut(docs, "\n\n")
		if brief != "" {
			brief = classNameString + " " + brief
		}
		class.BriefDescription = brief
		class.Description = whole
	}
	for i := 1; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		if !field.IsExported() || field.Anonymous || field.Name == "Object" {
			continue
		}
		if _, ok := field.Type.MethodByName("AsNode"); ok || field.Type.Kind() == reflect.Chan {
			continue
		}
		name := String.ToSnakeCase(field.Name)
		if field.Tag.Get("gd") != "" {
			name = field.Tag.Get("gd")
		}
		if reflect.PointerTo(field.Type).Implements(reflect.TypeOf([0]gd.IsSignal{}).Elem()) {
			var signal xmlSignal
			name, _, _ = strings.Cut(name, "(")
			signal.Name = name
			signal.Description = extractDocTag(field.Tag)
			class.Signals = append(class.Signals, signal)
			continue
		}
		ptype, ok := propertyOf(className, field)
		if ok {
			var member xmlMember
			member.Name = name
			member.Description = extractDocTag(field.Tag)
			member.Type = ptype.Type.String()
			class.Members = append(class.Members, member)
			gd.Global.ClassDB.RegisterClassProperty(gd.Global.ExtensionToken, className, ptype, gd.NewStringName(""), gd.NewStringName(""))
		}
	}
	Callable.Defer(Callable.New(func() {
		if EngineClass.IsEditorHint() {
			docs, _ := xml.Marshal(class)
			gd.Global.EditorHelp.Load(docs)
		}
	}))
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

func (class classImplementation) CreateInstance() [1]gd.Object {
	return class.createInstance(reflect.Value{})
}

func (class classImplementation) createInstance(reuse reflect.Value) [1]gd.Object {
	var super = gd.Global.ClassDB.ConstructObject(class.Super)
	super = [1]gd.Object{pointers.Pin(pointers.Lay(super[0]))}
	instance := class.reloadInstance(reuse, super)
	gd.Global.Object.SetInstance(super, class.Name, instance)
	gd.Global.Object.SetInstanceBinding(super, gd.Global.ExtensionToken, nil, nil)
	instance.OnCreate()
	return super
}

func (class classImplementation) reloadInstance(value reflect.Value, super [1]gd.Object) gd.ObjectInterface {
	if !value.IsValid() {
		value = reflect.New(class.Type)
	}
	extensionClass := value.Interface().(isClass)
	extensionClass.setObject(super)

	gd.ExtensionInstances.Store(pointers.Get(super[0])[0], extensionClass)

	value = value.Elem()

	// TODO cache this check
	var signals []signalChan
	var chSignals []signalChan
	for i := 0; i < value.NumField(); i++ {
		var field = value.Type().Field(i)
		if !field.IsExported() || field.Name == "Object" {
			continue
		}
		var (
			rvalue = value.Field(i).Addr()
		)
		name := field.Name
		if tag := field.Tag.Get("gd"); tag != "" {
			name = tag
		}
		name, _, _ = strings.Cut(name, "(")
		// Signal fields need to have their values injected into the field, so that they can be used (emitted).
		if reflect.PointerTo(field.Type).Implements(reflect.TypeOf([0]gd.IsSignal{}).Elem()) {
			signal := pointers.Pin(gd.NewSignalOf(super, gd.NewStringName(name)))
			gd.SetSignal(rvalue.Interface().(gd.IsSignal), signal)
			signals = append(signals, signalChan{
				signal: signal,
			})
		}
		if field.Type.Kind() == reflect.Chan && field.Type.ChanDir() == reflect.SendDir {
			signal := pointers.Pin(gd.NewSignalOf(super, gd.NewStringName(name)))
			ch := reflect.MakeChan(reflect.ChanOf(reflect.BothDir, field.Type.Elem()), 0)
			rvalue.Elem().Set(ch)
			signals = append(signals, signalChan{
				signal: signal,
				rvalue: ch,
			})
			chSignals = append(chSignals, signalChan{
				signal: signal,
				rvalue: ch,
			})
		}
	}
	if len(signals) > 0 {
		go manageSignals(super[0].AsObject()[0].GetInstanceId(), chSignals)
	}
	return &instanceImplementation{
		object:   pointers.Get(super[0])[0],
		Value:    value.Addr().Interface().(isClass),
		signals:  signals,
		isEditor: !class.Tool && EngineClass.IsEditorHint(),
	}
}

func (class classImplementation) GetVirtual(name gd.StringName) any {
	if !class.Tool && EngineClass.IsEditorHint() {
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
		atype := method.Type.In(i)
		btype := vtype.In(i)
		if atype != btype && !(atype.ConvertibleTo(btype) && atype.Kind() == btype.Kind()) {
			panic(fmt.Sprintf("gdextension.RegisterClass: Method %s.%s does not match %s.%s\nis %s want %s", class.Type.Name(), GoName, virtual.Type().Name(), name, method.Type, vtype))
		}
	}
	var copy = reflect.New(method.Type)
	copy.Elem().Set(method.Func)
	var fn = reflect.NewAt(vtype, copy.UnsafePointer()).Elem()
	return virtual.Call([]reflect.Value{fn})[0].Interface()
}

type instanceImplementation struct {
	object  uint64
	Value   isClass
	signals []signalChan

	// FIXME use a bitfield for these booleans.
	isEditor bool
}

var lastGC int

func (instance *instanceImplementation) OnCreate() {
	if impl, ok := instance.Value.(interface {
		OnCreate()
	}); ok {
		impl.OnCreate()
	}
}

func (instance *instanceImplementation) Notification(what int32, reversed bool) {
	if what == 13 { // NOTIFICATION_READY
		instance.ready()
	}
	if !instance.isEditor {
		switch notify := instance.Value.(type) {
		case interface{ Notification(gd.NotificationType) }:
			notify.Notification(gd.NotificationType(what))
		default:
		}
	}
}

func (instance *instanceImplementation) ToString() (gd.String, bool) {
	switch onfree := instance.Value.(type) {
	case interface{ ToString() string }:
		return gd.NewString(onfree.ToString()), true
	}
	return gd.String{}, false
}

func (instance *instanceImplementation) Reference()   {}
func (instance *instanceImplementation) Unreference() {}

func (instance *instanceImplementation) CallVirtual(name gd.StringName, virtual any, args gd.Address, back gd.Address) {
	virtual.(gd.ExtensionClassCallVirtualFunc)(instance.Value, args, back)
}

func (instance *instanceImplementation) GetRID() gd.RID {
	return 0
}

func (instance *instanceImplementation) Free() {
	for _, signal := range instance.signals {
		if signal.rvalue.IsValid() {
			signal.rvalue.Close()
		}
		signal.signal.Free()
	}
	rvalue := reflect.ValueOf(instance.Value).Elem()
	for i := range rvalue.NumField() {
		field := rvalue.Type().Field(i)
		if !field.IsExported() || field.Name == "Object" {
			continue
		}
		type isNode interface {
			AsNode() NodeClass.Instance
		}
		nodeType := reflect.TypeOf([0]isNode{}).Elem()
		if field.Type.Implements(nodeType) || reflect.PointerTo(field.Type).Implements(nodeType) {
			continue
		}
		if field.Type.Implements(reflect.TypeFor[interface{ Free() }]()) {
			rvalue.Field(i).Interface().(interface{ Free() }).Free()
		}
		if field.Type.Kind() == reflect.Array && field.Type.Len() == 1 && field.Type.Elem().Implements(reflect.TypeFor[interface{ Free() }]()) {
			rvalue.Field(i).Index(0).Interface().(interface{ Free() }).Free()
		}
	}
	gd.ExtensionInstances.Delete(instance.object)
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
	parent, ok := As[NodeClass.Instance](Object.Instance(instance.Value.getObject()))
	if !ok {
		return
	}

	var rvalue = reflect.ValueOf(instance.Value).Elem()
	for i := 0; i < rvalue.NumField(); i++ {
		field := rvalue.Type().Field(i)
		if !field.IsExported() || field.Name == "Class" {
			continue
		}
		instance.assertChild(rvalue.Field(i).Addr().Interface(), field, parent, parent)
	}
	if !instance.isEditor {
		switch ready := instance.Value.(type) {
		case interface{ Ready() }:
			ready.Ready()
		}
	}
}

func (instance *instanceImplementation) assertChild(value any, field reflect.StructField, parent, owner Node) {
	type isNode interface {
		AsNode() NodeClass.Instance
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
	path := Path.ToNode(String.New(name))
	if !NodeClass.Advanced(parent).HasNode(path) {
		child := gd.Global.ClassDB.ConstructObject(gd.NewStringName(nameOf(field.Type)))
		defer pointers.End(child[0])
		native, ok := gd.ExtensionInstances.Load(pointers.Get(child[0])[0])
		if ok {
			rvalue.Elem().Set(reflect.ValueOf(native))
			class = native.(isNode)
		} else {
			type isUnsafe interface {
				UnsafePointer() unsafe.Pointer
			}
			*(*gd.Object)(class.(isUnsafe).UnsafePointer()) = pointers.Raw[gd.Object](pointers.Get(child[0]))
		}
		var mode NodeClass.InternalMode = NodeClass.InternalModeDisabled
		if !field.IsExported() {
			mode = NodeClass.InternalModeFront
		}
		NodeClass.Advanced(class.AsNode()).SetName(String.New(field.Name))
		NodeClass.Advanced(parent).AddChild(class.AsNode(), true, mode)
		NodeClass.Advanced(class.AsNode()).SetOwner(owner)
		return
	}
	var node = NodeClass.Advanced(parent).GetNode(path)
	if name := node[0].AsObject()[0].GetClass().String(); name != nameOf(field.Type) {
		panic(fmt.Sprintf("gd.Register: Node %s.%s is not of type %s (%s)", rvalue.Type().Name(), field.Name, field.Type.Name(), name))
	}
	ref, native := gd.ExtensionInstances.Load(pointers.Get(node[0])[0])
	if native {
		rvalue.Elem().Set(reflect.ValueOf(ref))
		pointers.End(node[0])
	} else {
		type isUnsafe interface {
			UnsafePointer() unsafe.Pointer
		}
		*(*gd.Object)(class.(isUnsafe).UnsafePointer()) = pointers.Raw[gd.Object](pointers.Get(node[0]))
		pointers.End(node[0])
	}
}
