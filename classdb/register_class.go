package classdb

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"maps"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	EditorInterfaceClass "graphics.gd/classdb/EditorInterface"
	EditorPluginClass "graphics.gd/classdb/EditorPlugin"
	EngineClass "graphics.gd/classdb/Engine"
	NodeClass "graphics.gd/classdb/Node"
	ScriptClass "graphics.gd/classdb/Script"
	ScriptLanguageClass "graphics.gd/classdb/ScriptLanguage"
	ShaderMaterialClass "graphics.gd/classdb/ShaderMaterial"

	"graphics.gd/variant/Object"
	"graphics.gd/variant/Path"
	"graphics.gd/variant/RefCounted"
	"graphics.gd/variant/Signal"
	"graphics.gd/variant/String"

	gd "graphics.gd/internal"
	"graphics.gd/internal/docgen"
	"graphics.gd/internal/gdclass"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
)

// Tool can be embedded inside a struct to make it run in the editor.
type Tool interface{ tool() }

// Deprecated: use a classdb package Extension instead, ie. Node.Extension[MyClass]
type Extension[T Class, S gd.IsClass] = gdclass.Extension[T, S]

// NameOf returns the defined name for the given [Extension]-embedding type.
func NameOf(T Class) string {
	return nameOf(reflect.TypeOf(T))
}

func NameFor[T Class]() string {
	return nameOf(reflect.TypeFor[T]())
}

type Class = gdclass.Interface

/*
Register registers a struct available for use inside The Engine as
an object (or class) by extending the given 'Parent' Engine class.
The 'Struct' type must be a named struct with the first field
embedding [Extension] referring to itself and specifying the
parent class to extend.

	type MyClass struct {
		Class[MyClass, Node2D] `gd:"MyClass"`
	}

The tag can be adjusted in order to change the name of the class
within The Engine.

Use this in a main or init function to register your Go structs
and they will become available within The Engine for use in the
editor and/or within scripts. Call this before loading the scene.

All exported fields and methods will be exposed to The Engine, so
take caution when embedding types, as their fields and methods
will be promoted. They will be exported as snake_case by default,
for fields, the exported name can be adjusted with the 'gd' tag.

The following struct tags can be used to adjust the behavior of
class members within The Engine:

  - range can be used to specify the range hint of the member.
  - group can be used to group members together in the editor.

This function accepts a variable number of additional arguments,
they may either be func, map[string]any (where each any is a func),
map[string]string or map[string]int, these arguments can be used to
register static methods, rename existing methods, add symbol documentation
or to define constants respectively. As a special case, if a function is
passed which name begins with 'New' and accepts no arguments, returning T,
then it will be registered as the constructor for the class when it is
instantiated from within The Engine.

If the Struct extends [EditorPluginClass] then it will be added
to the editor as a plugin.
*/
func Register[T Class](exports ...any) {
	var superType = gdclass.SuperType(([1]T{})[0])
	var super = reflect.New(superType).Elem().Interface()
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

		if classType.Kind() != reflect.Struct || classType.Name() == "" {
			panic("gdextension.RegisterClass: Class type must be a named struct")
		}
		var rename = nameOf(classType) // support 'gd' tag for renaming the class within Godot.
		var tool = false
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

		var refCounted bool
		switch super.(type) {
		case interface{ AsRefCounted() [1]gd.RefCounted }:
			refCounted = true // FIXME I think this can be unsafely overridden by the user, so we should check if the type is actually a RefCounted type.
		}

		var impl = &classImplementation{
			Name:           className,
			Super:          superName,
			Type:           classType,
			Tool:           tool,
			RefCounted:     refCounted,
			VirtualMethods: reference.Virtual,
			Constructor: func() reflect.Value {
				return reflect.New(classType)
			},
		}
		gdclass.Registered.Store(classType, impl)

		gdextension.Host.ClassDB.Register.Class(pointers.Get(className), pointers.Get(superName), gdextension.ExtensionClassID(cgoNewHandle(impl)), false, false, true, false, gdextension.String{})

		gd.RegisterCleanup(func() {
			gdextension.Host.ClassDB.Register.Removal(pointers.Get(className))
			gdclass.Registered.Delete(classType)
			className.Free()
			superName.Free()
		})
		var (
			documentation = make(map[string]string)
		)
		var method_renames = make(map[uintptr]string)
		for _, export := range exports {
			switch export := export.(type) {
			case Trampoline[T]:

			case map[string]string:
				maps.Copy(documentation, export)
			case map[string]int:
				for name, value := range export {
					gdextension.Host.ClassDB.Register.Constant(
						pointers.Get(className),
						pointers.Get(gd.NewStringName("")),
						pointers.Get(gd.NewStringName(name)),
						int64(value),
						false,
					)
				}
			case map[string]any:
				for name, fn := range export {
					if reflect.TypeOf(fn).Kind() != reflect.Func {
						panic(fmt.Sprintf("gdextension.RegisterClass: invalid map elem type %T (expected function)", fn))
					}
					rvalue := reflect.ValueOf(fn)
					pc := rvalue.Pointer()
					fname := runtime.FuncForPC(pc).Name()
					if strings.Count(path.Base(fname), ".") > 1 {
						method_renames[pc] = name
					} else {
						registerStaticMethod(className, name, reflect.ValueOf(fn))
					}
				}
			default:
				rvalue := reflect.ValueOf(export)
				switch rvalue.Kind() {
				case reflect.Func:
					pc := rvalue.Pointer()
					fname := runtime.FuncForPC(pc).Name()
					name := fname
					i := String.FindLast(name, ".")
					name = name[i+1:]
					if String.HasPrefix(name, "New") && rvalue.Type().NumIn() == 0 && rvalue.Type().NumOut() == 1 && rvalue.Type().Out(0) == reflect.PointerTo(classType) {
						impl.Constructor = func() reflect.Value {
							return rvalue.Call(nil)[0]
						}
					} else if strings.Count(path.Base(fname), ".") > 1 {
						method_renames[pc] = name
					} else {
						registerStaticMethod(className, String.ToSnakeCase(name), rvalue)
					}
				default:
					panic(fmt.Sprintf("gdextension.RegisterClass: invalid argument type %T (expected function or map)", export))
				}
			}
		}
		switch super.(type) {
		case interface {
			AsShaderMaterial() ShaderMaterialClass.Instance
		}:
		default:
			registerClassInformation(className, rename, nameOf(superType), classType, documentation, method_renames)
			registerSignals(className, classType)
			registerMethods(className, classType, method_renames)
		}
		if registrator, ok := any(reference).(interface{ OnRegister() }); ok {
			registrator.OnRegister()
		}
		if EngineClass.IsEditorHint() {
			switch super.(type) {
			case EditorPluginClass.Any:
				gdextension.Host.Editor.AddPlugin(pointers.Get(className))
			}
		}
	}
	switch super.(type) {
	case interface{ AsScript() ScriptClass.Instance },
		interface {
			AsEditorPlugin() EditorPluginClass.Instance
		},
		interface {
			AsScriptLanguage() ScriptLanguageClass.Instance
		}:
		gd.EditorStartupFunctions = append(gd.EditorStartupFunctions, register)
	default:
		if gd.Linked {
			register()
		} else {
			gd.StartupFunctions = append(gd.StartupFunctions, register)
		}
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
	for word := range strings.SplitSeq(fnName, "_") {
		joins = append(joins, cases.Title(language.English).String(word))
	}
	return strings.Join(joins, "")
}

var preloaded_documentation = make(map[string]docgen.Class)

func init() {
	gd.StartupFunctions = append(gd.StartupFunctions, func() {
		if EngineClass.IsEditorHint() {
			path := pointers.New[gd.String](gdextension.Host.Library.Location())
			data, err := os.Open(filepath.Join(filepath.Dir(path.String()), "library_documentation.xml"))
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					return
				}

				EngineClass.Raise(err)
			}
			var dec = xml.NewDecoder(data)
			var docs docgen.XML
			for {
				if err := dec.Decode(&docs); err != nil {
					if err == io.EOF {
						break
					}
					EngineClass.Raise(fmt.Errorf("failed to unmarshal library documentation: %w", err))
					break
				}
			}
			for _, class := range docs {
				preloaded_documentation[class.Name] = class
			}
		}
	})
}

func registerClassInformation(className gd.StringName, classNameString string, inherits string, rtype reflect.Type, docs map[string]string, method_renames map[uintptr]string) {
	var class = preloaded_documentation[classNameString]
	class.Name = classNameString
	class.Inherits = inherits
	class.Version = "4.0"
	extractDocTag := func(tag reflect.StructTag) string {
		_, docs, _ := strings.Cut(string(tag), "\n")
		docs = strings.Replace(docs, "\t", "", -1)
		return strings.TrimSpace(docs)
	}
	extractDoc := func(docs string) string {
		docs = strings.Replace(docs, "\t", "", -1)
		return strings.TrimSpace(docs)
	}
	if rtype.Field(0).Anonymous {
		docs := extractDocTag(rtype.Field(0).Tag)
		brief, whole, _ := strings.Cut(docs, "\n\n")
		if brief != "" {
			brief = classNameString + " " + brief
		}
		if docs != "" {
			class.BriefDescription = brief
			class.Description = whole
		}
	}
	ungroupedFields := make([]reflect.StructField, 0)
	groupedFields := map[string][]reflect.StructField{}
	for _, field := range reflect.VisibleFields(rtype) {
		groupName := field.Tag.Get("group")
		if groupName == "" {
			ungroupedFields = append(ungroupedFields, field)
			continue
		}
		groupedFields[groupName] = append(groupedFields[groupName], field)
	}
	registerField := func(field reflect.StructField) {
		if !field.IsExported() || field.Anonymous || field.Name == "Object" {
			return
		}
		if _, ok := field.Type.MethodByName("AsNode"); ok || field.Type.Kind() == reflect.Chan {
			return
		}
		name := String.ToSnakeCase(field.Name)
		if tag := field.Tag.Get("gd"); tag != "" {
			if tag == "-" {
				return
			}
			name = tag
		}
		if reflect.PointerTo(field.Type).Implements(reflect.TypeFor[Signal.Pointer]()) {
			var signal docgen.Signal
			name, _, _ = strings.Cut(name, "(")
			signal.Name = name
			signal.Description = extractDocTag(field.Tag)
			if docs, ok := docs[name]; ok {
				signal.Description = extractDoc(docs)
			}
			class.Signals = append(class.Signals, signal)
			return
		}
		var ptype gdextension.PropertyList
		ptype = gdextension.Host.ClassDB.PropertyList.Make(1)
		if propertyOf(className, field, ptype) {
			var exists bool
			var member = new(docgen.Member)
			for i := range class.Members {
				if class.Members[i].Name == name {
					member = &class.Members[i]
					exists = true
					break
				}
			}
			member.Name = name
			if doctag := extractDocTag(field.Tag); doctag != "" {
				member.Description = doctag
			}
			if member.Description != "" {
				member.Description = member.Name + " " + member.Description
			}
			if docs, ok := docs[member.Name]; ok {
				member.Description = extractDoc(docs)
			}
			member.Type = gdextension.Host.ClassDB.PropertyList.Info.Type(ptype).String()
			if !exists {
				class.Members = append(class.Members, *member)
			}
			gdextension.Host.ClassDB.Register.Property(pointers.Get(className), ptype, pointers.Get(gd.NewStringName("")), pointers.Get(gd.NewStringName("")))
		}
		gdextension.Host.ClassDB.PropertyList.Free(ptype)
	}
	for _, field := range ungroupedFields {
		registerField(field)
	}
	for groupName, fields := range groupedFields {
		gdextension.Host.ClassDB.Register.PropertyGroup(
			pointers.Get(className),
			pointers.Get(gd.NewString(groupName)),
			pointers.Get(gd.NewString("")),
		)
		for _, field := range fields {
			registerField(field)
		}
	}
	rtype = reflect.PointerTo(rtype)
	for i := range rtype.NumMethod() {
		name := String.ToSnakeCase(rtype.Method(i).Name)
		if rename, ok := method_renames[rtype.Method(i).Func.Pointer()]; ok {
			name = rename
		}
		if _, ok := docs[name]; !ok {
			continue
		}
		var exists bool
		var method = new(docgen.Method)
		for i := range class.Methods {
			if class.Methods[i].Name == name {
				method = &class.Methods[i]
				exists = true
				break
			}
		}
		method.Name = name
		if docs := extractDoc(docs[name]); docs != "" {
			method.Description = docs
		}
		if !exists {
			class.Methods = append(class.Methods, *method)
		}
	}
	gd.NewCallable(func() {
		if EngineClass.IsEditorHint() {
			docs, _ := xml.Marshal(class)
			gdextension.Host.Editor.AddDocumentation(string(docs))
		}
	}).CallDeferred()
}

type classImplementation struct {
	Name  gd.StringName
	Super gd.StringName

	Tool       bool
	RefCounted bool

	Type reflect.Type

	VirtualMethods func(string) reflect.Value
	Constructor    func() reflect.Value
}

func (class classImplementation) IsVirtual() bool {
	return false
}

func (class classImplementation) IsAbstract() bool {
	return class.Type.Kind() == reflect.Interface
}

func (class classImplementation) IsExposed() bool {
	return true // TODO return false if the Go type is not exported.
}

func (class classImplementation) CreateInstance(notify_postinitialize bool) [1]gd.Object {
	return class.CreateInstanceFrom(class.Constructor(), notify_postinitialize)
}

func (class classImplementation) CreateInstanceFrom(value reflect.Value, notify_postinitialize bool) [1]gd.Object {
	var super = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gdextension.Host.Objects.Make(pointers.Get(class.Super)))})}
	if class.RefCounted {
		gd.RefCounted(super[0]).InitRef()
	}
	super = [1]gd.Object{pointers.Pin(super[0])}
	instance := class.reloadInstance(value, super)
	gdextension.Host.Objects.Extension.Setup(gdextension.Object(pointers.Get(super[0])[0]), pointers.Get(class.Name), gdextension.ExtensionInstanceID(cgoNewHandle(instance)))
	if notify_postinitialize {
		super[0].Notification(0, false)
	}
	instance.OnCreate(value)
	return super
}

func (class classImplementation) CreateGoInstanceFrom(value reflect.Value, notify_postinitialize bool) [1]gd.Object {
	var super = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gdextension.Host.Objects.Make(pointers.Get(class.Super)))})}
	if class.RefCounted {
		gd.RefCounted(super[0]).InitRef()
	}
	instance := class.reloadInstance(value, super)
	gdextension.Host.Objects.Extension.Setup(gdextension.Object(pointers.Get(super[0])[0]), pointers.Get(class.Name), gdextension.ExtensionInstanceID(cgoNewHandle(instance)))
	if notify_postinitialize {
		super[0].Notification(0, false)
	}
	instance.OnCreate(value)
	return super
}

func (class classImplementation) reloadInstance(value reflect.Value, super [1]gd.Object) *instanceImplementation {
	extensionClass := value.Interface().(gdclass.Pointer)
	gdclass.SetObject(extensionClass, super)

	value = value.Elem()

	// TODO cache this check
	var signals []signalChan
	var chSignals []signalChan
	for _, field := range reflect.VisibleFields(value.Type()) {
		if !field.IsExported() || field.Name == "Object" {
			continue
		}
		var (
			rvalue = value.FieldByIndex(field.Index).Addr()
		)
		name := field.Name
		if tag := field.Tag.Get("gd"); tag != "" {
			name = tag
		}
		name, _, _ = strings.Cut(name, "(")
		// Signal fields need to have their values injected into the field, so that they can be used (emitted).
		if reflect.PointerTo(field.Type).Implements(reflect.TypeFor[Signal.Pointer]()) {
			signal := pointers.Pin(gd.NewSignalOf(super, gd.NewStringName(name)))
			rvalue.Interface().(Signal.Pointer).SetAny(Signal.Via(gd.SignalProxy{}, pointers.Pack(signal)))
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
		go manageSignals(Object.Instance(super[0].AsObject()).ID(), chSignals)
	}
	return &instanceImplementation{
		object:   pointers.Get(super[0])[0],
		Value:    value.Addr().Interface().(gdclass.Pointer),
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
	if method.Type.NumOut() != vtype.NumOut() {
		panic(fmt.Sprintf("gdextension.RegisterClass: Method %s.%s does not match %s.%s\nis %s want %s", class.Type.Name(), GoName, virtual.Type().Name(), name, method.Type, vtype))
	}
	if method.Type.NumOut() > 0 {
		atype := method.Type.Out(0)
		btype := vtype.Out(0)
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
	Value   gdclass.Pointer
	signals []signalChan

	// FIXME use a bitfield for these booleans.
	isEditor, freed bool
}

var lastGC int

func (instance *instanceImplementation) OnCreate(value reflect.Value) {
	if impl, ok := instance.Value.(interface {
		OnCreate()
	}); ok {
		impl.OnCreate()
	}
	if impl, ok := instance.Value.(interface {
		OnCreate(value reflect.Value)
	}); ok {
		impl.OnCreate(value)
	}
	if impl, ok := instance.Value.(interface {
		Init()
	}); ok {
		impl.Init()
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
		case interface{ Notification(int, bool) }:
			notify.Notification(int(what), reversed)
		default:
		}
	}
}

func (instance *instanceImplementation) ToString() (gd.String, bool) {
	switch onfree := instance.Value.(type) {
	case interface{ ToString() string }:
		return gd.NewString(onfree.ToString()), true
	case interface{ String() string }:
		return gd.NewString(onfree.String()), true
	}
	return gd.String{}, false
}

func (instance *instanceImplementation) Reference() {

}
func (instance *instanceImplementation) Unreference() bool {
	return false
}

func (instance *instanceImplementation) CallVirtual(virtual any, args, back gdextension.Pointer) {
	virtual.(gd.ExtensionClassCallVirtualFunc)(instance.Value, args, back)
}

func (instance *instanceImplementation) GetRID() gd.RID {
	return 0
}

func (instance *instanceImplementation) Free() {
	if instance.freed {
		return
	}
	for _, signal := range instance.signals {
		if signal.rvalue.IsValid() {
			signal.rvalue.Close()
		}
		signal.signal.Free()
	}
	rvalue := reflect.ValueOf(instance.Value).Elem()
	for _, field := range reflect.VisibleFields(rvalue.Type()) {
		if !field.IsExported() || field.Name == "Extension" || rvalue.FieldByIndex(field.Index).IsZero() {
			continue
		}
		type isNode interface {
			AsNode() NodeClass.Instance
		}
		nodeType := reflect.TypeOf([0]isNode{}).Elem()
		if field.Type.Implements(nodeType) || reflect.PointerTo(field.Type).Implements(nodeType) {
			continue
		}
		// we need to unreference any pinned resources (pinned here means that the engine `set` them).
		if field.Type.Implements(reflect.TypeFor[RefCounted.Any]()) {
			ref := rvalue.FieldByIndex(field.Index).Interface().(RefCounted.Any).AsRefCounted()[0]
			raw, kind := pointers.Ask(gd.Object(ref))
			if kind == pointers.Pinned {
				if ref.Unreference() {
					gdextension.Host.Objects.Unsafe.Free(gdextension.Object(raw[0]))
				}
			}
		}
	}
	switch onfree := instance.Value.(type) {
	case interface{ OnFree() }:
		onfree.OnFree()
	}
	instance.freed = true
}

// ready is responsible for asserting the scene tree for struct members that implement
// Super().AsNode() and asserting that these nodes are added as children to the Super.
//
// TODO this could be partially pre-compiled for a given [Register] type and cached in
// order to avoid any use of reflection at instantiation time.
func (instance *instanceImplementation) ready() {
	parent, ok := Object.As[NodeClass.Instance](Object.Instance(gdclass.GetObject(instance.Value)))
	if !ok {
		return
	}
	var rvalue = reflect.ValueOf(instance.Value).Elem()
	for _, field := range reflect.VisibleFields(rvalue.Type()) {
		if !field.IsExported() || field.Name == "Extension" {
			continue
		}
		instance.assertChild(rvalue.FieldByIndex(field.Index).Addr().Interface(), field, parent, parent)
	}
	if !instance.isEditor {
		switch ready := instance.Value.(type) {
		case interface{ Ready() }:
			ready.Ready()
		}
	}
}

func (instance *instanceImplementation) assertChild(value any, field reflect.StructField, parent, owner [1]gdclass.Node) {
	type isNode interface {
		AsNode() NodeClass.Instance
	}
	var (
		rvalue = reflect.ValueOf(value)
	)
	nodeType := reflect.TypeOf([0]isNode{}).Elem()
	if !field.Type.Implements(nodeType) && !reflect.PointerTo(field.Type).Implements(nodeType) {
		if field.Type.Kind() == reflect.Struct {
			for _, field := range reflect.VisibleFields(field.Type) {
				if field.IsExported() {
					instance.assertChild(rvalue.Elem().FieldByIndex(field.Index).Addr().Interface(), field, parent, owner)
				}
			}
		}
		return
	}
	if field.Anonymous {
		return
	}
	if rvalue.Elem().Kind() == reflect.Pointer {
		rvalue.Elem().Set(reflect.New(rvalue.Elem().Type().Elem()))
		value = rvalue.Elem().Interface()
	}
	class := value.(isNode)
	if rvalue.Elem().Kind() == reflect.Struct {
		defer func() {
			rvalue := rvalue.Elem()
			for _, field := range reflect.VisibleFields(rvalue.Type()) {
				if !field.IsExported() || field.Name == "Class" || field.Anonymous {
					continue
				}
				instance.assertChild(rvalue.FieldByIndex(field.Index).Addr().Interface(), field, class.AsNode(), owner)
			}
		}()
	}
	name := field.Name
	if tag := field.Tag.Get("gd"); tag != "" {
		name = tag
	}
	path := Path.ToNode(String.New(name))
	if !NodeClass.Advanced(parent).HasNode(path) {
		child := [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gdextension.Host.Objects.Make(pointers.Get(gd.NewStringName(nameOf(field.Type)))))})}
		child[0].Notification(0, false)
		defer pointers.End(child[0])
		native := gd.ExtensionInstanceLookup(gdextension.Object(pointers.Get(child[0])[0]))
		if native != nil {
			rvalue.Elem().Set(reflect.ValueOf(native))
			class = native.(isNode)
		} else {
			class.(gd.IsClassCastable).SetObject([1]gd.Object{pointers.Raw[gd.Object](pointers.Get(child[0]))})
		}
		var mode NodeClass.InternalMode = NodeClass.InternalModeDisabled
		if !field.IsExported() {
			mode = NodeClass.InternalModeFront
		}
		NodeClass.Advanced(class.AsNode()).SetName(String.New(field.Name))
		NodeClass.Advanced(parent).AddChild(class.AsNode(), true, mode)
		if EngineClass.IsEditorHint() {
			NodeClass.Advanced(class.AsNode()).SetOwner(EditorInterfaceClass.GetEditedSceneRoot())
		}
		return
	}
	var node = NodeClass.Advanced(parent).GetNode(path)
	native := gd.ExtensionInstanceLookup(gdextension.Object(pointers.Get(node[0])[0]))
	if native != nil {
		if reflect.ValueOf(native).Type() != rvalue.Elem().Type() {
			fmt.Printf("gd.Register: Node %s.%s is not of type %s (%s)", rvalue.Type().Name(), field.Name, field.Type.Name(), name)
			panic(fmt.Sprintf("gd.Register: Node %s.%s is not of type %s (%s)", rvalue.Type().Name(), field.Name, field.Type.Name(), name))
		}
		rvalue.Elem().Set(reflect.ValueOf(native))
		pointers.End(node[0])
	} else {
		if !class.(gd.IsClassCastable).SetObject([1]gd.Object{pointers.Raw[gd.Object](pointers.Get(node[0]))}) {
			fmt.Printf("gd.Register: Node %s.%s is not of type %s (%s)", rvalue.Type().Name(), field.Name, field.Type.Name(), name)
			panic(fmt.Sprintf("gd.Register: Node %s.%s is not of type %s (%s)", rvalue.Type().Name(), field.Name, field.Type.Name(), name))
		}
		pointers.End(node[0])
	}
}
