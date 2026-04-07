package classdb

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"iter"
	"maps"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"
	"unsafe"
	"weak"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"graphics.gd/classdb/EditorInterface"
	"graphics.gd/classdb/EditorPlugin"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/GDExtension"
	"graphics.gd/classdb/MainLoop"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/Script"
	"graphics.gd/classdb/ScriptLanguage"
	"graphics.gd/classdb/ShaderMaterial"

	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Path"
	"graphics.gd/variant/RefCounted"
	"graphics.gd/variant/Signal"
	"graphics.gd/variant/String"

	gdunsafe "graphics.gd"
	gd "graphics.gd/internal"
	"graphics.gd/internal/docgen"
	"graphics.gd/internal/gdclass"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdreference"
	"graphics.gd/internal/pointers"
	"graphics.gd/internal/ring"
	"graphics.gd/internal/threadsafe"
)

// pendingRegistrations holds registration functions that are waiting for their parent
// extension class to be registered first. Key is the parent's reflect.Type.
var pendingRegistrations = make(map[reflect.Type][]func())

// registeredTypes tracks which extension class types have been registered.
var registeredTypes = make(map[reflect.Type]bool)

// processPendingRegistrations checks if any deferred registrations can now proceed
// because their parent class has been registered.
func processPendingRegistrations(registeredType reflect.Type) {
	if pending, ok := pendingRegistrations[registeredType]; ok {
		delete(pendingRegistrations, registeredType)
		for _, register := range pending {
			register()
		}
	}
}

// Tool can be embedded inside a struct to make it run in the editor.
type Tool interface{ tool() }

// NameOf returns the defined name for the given [Extension]-embedding type.
func NameOf(T Class) string {
	return nameOf(reflect.TypeOf(T))
}

func NameFor[T Class]() string {
	return nameOf(reflect.TypeFor[T]())
}

type Class = gdclass.Interface

var singletons threadsafe.Map[reflect.Type, reflect.Value]

func init() {
	gd.RegisterCleanup(func() {
		for _, value := range singletons.Iter() {
			switch singleton := value.Interface().(type) {
			case Node.Any:
				continue
			case Object.Any:
				gd.ObjectFree(singleton.AsObject()[0])
			}
		}
	})
}

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

If the Struct extends [EditorPlugin] then it will be added
to the editor as a plugin.
*/
func Register[T Class](exports ...any) {
	var superType = gdclass.SuperType(([1]T{})[0])
	var super = reflect.New(superType).Elem().Interface()
	var classType = reflect.TypeFor[T]()

	var underlyingType = gdclass.GoType(([1]T{})[0])
	var trivialExtension = classType.Size() == underlyingType.Size() && classType.NumField() == 1 && classType.Field(0).Type == underlyingType
	if !trivialExtension && classType != underlyingType && !classType.ConvertibleTo(underlyingType) {
		// FIXME enable this as a strict safety check at some point.
		Engine.RaiseWarning("classdb.Register: embedded Extension type must match the registered type\nSee https://the.graphics.gd/guide/classdb/register/#inheritance")
	}

	register := func() {
		// Mark this type as registered and process any pending child registrations
		registeredTypes[classType] = true
		defer processPendingRegistrations(classType)

		compile_keepalive(reflect.PointerTo(classType))
		var base = classType
		var tags reflect.StructTag = base.Field(0).Tag
		var embedded_name string
		for base.Field(0).Anonymous {
			if base.Field(0).Name == "Class" {
				break
			}
			base = base.Field(0).Type
			if embedded_name == "" {
				embedded_name = classType.Field(0).Name
			}
		}
		if !base.Implements(reflect.TypeFor[Class]()) {
			panic("classdb.Register: Class type must embed an Extension[T] as the first field")
		}
		if classType.Kind() != reflect.Struct || classType.Name() == "" {
			panic("classdb.Register: Class type must be a named struct")
		}
		var rename = nameOf(classType) // support 'gd' tag for renaming the class within Godot.
		if embedded_name == "Singleton" {
			rename = "GoSingleton" + rename
		}
		var tool = false
		switch super.(type) {
		case interface{ AsScript() Script.Instance },
			interface {
				AsEditorPlugin() EditorPlugin.Instance
			},
			interface {
				AsScriptLanguage() ScriptLanguage.Instance
			}:
			tool = true
		}
		var isMainLoop bool
		switch super.(type) {
		case interface{ AsMainLoop() MainLoop.Instance }:
			isMainLoop = true
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

		// Find the engine class (first non-extension ancestor).
		// This is safe because the deferred registration system ensures
		// parent extension classes are registered before their children.
		engineClass := findEngineClass(superType)

		var impl = &classImplementation{
			Name:           className,
			Super:          superName,
			SuperType:      superType,
			EngineClass:    engineClass,
			Type:           classType,
			Tool:           tool,
			RefCounted:     refCounted,
			isMainLoop:     isMainLoop,
			InEditor:       Engine.IsEditorHint(),
			VirtualMethods: reference.Virtual,
			Constructor: func() reflect.Value {
				return reflect.New(classType)
			},
		}
		for _, field := range reflect.VisibleFields(classType) {
			if field.Type.Kind() == reflect.Pointer && field.Type.Elem().Kind() == reflect.Struct && field.Type.Elem().NumField() > 0 {
				check := field.Type.Elem().Field(0)
				if check.Name == "Singleton" && check.Anonymous && check.Type.Implements(reflect.TypeFor[Class]()) {
					impl.Singletons = append(impl.Singletons, field)
				}
			}
		}
		gdclass.Registered.Store(classType, impl)

		var iconString gd.String
		if icon, ok := tags.Lookup("icon"); ok {
			iconString = gd.NewString(icon)
		}
		gdunsafe.RegisterClass(gdunsafe.StringName(pointers.Get(className)[0]), gdunsafe.StringName(pointers.Get(superName)[0]), &classWrapper{impl: impl}, false, false, true, false, gdunsafe.String(pointers.Get(iconString)[0]))

		gd.RegisterCleanup(func() {
			gdunsafe.RegisterRemoval(gdunsafe.StringName(pointers.Get(className)[0]))
			gdclass.Registered.Delete(classType)
			className.Free()
			superName.Free()
			engineClass.Free()
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
					gdunsafe.RegisterConstant(
						gdunsafe.StringName(pointers.Get(className)[0]),
						gdunsafe.StringName(pointers.Get(gd.NewStringName(""))[0]),
						gdunsafe.StringName(pointers.Get(gd.NewStringName(name))[0]),
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
			AsShaderMaterial() ShaderMaterial.Instance
		}:
		default:
			registerClassInformation(className, rename, nameOf(superType), classType, documentation, method_renames)
			registerSignals(className, classType)
			registerMethods(className, classType, method_renames)
		}
		if registrator, ok := any(reference).(interface{ OnRegister() }); ok {
			registrator.OnRegister()
		}
		if Engine.IsEditorHint() {
			switch super.(type) {
			case EditorPlugin.Any:
				gdunsafe.EditorAddPlugin(gdunsafe.StringName(pointers.Get(className)[0]))
			}
		}
		if embedded_name == "Singleton" {
			construct := impl.Constructor()
			singleton, _ := reflect.TypeAssert[Object.Any](construct)
			singletons.Insert(classType, construct)
			Engine.RegisterSingleton(strings.TrimPrefix(rename, "GoSingleton"), singleton.AsObject())
			if node, ok := singleton.(Node.Any); ok {
				Callable.Defer(Callable.New(func() {
					ptrs := gdreference.GetObject(node.AsNode().AsObject()[0])
					SceneTree.Add(Node.Instance{gdclass.NewNode(gdreference.OwnObject(ptrs, gd.Free))})
				}))
			}
		}
	}
	// Check if the parent type is an extension class that hasn't been registered yet.
	// If so, defer this registration until after the parent is registered.
	// This is needed because when extension classes inherit from other extension classes,
	// we need to know the full inheritance chain to find the underlying engine class.
	isParentExtensionClass := superType.Implements(reflect.TypeFor[gdclass.Interface]()) ||
		reflect.PointerTo(superType).Implements(reflect.TypeFor[gdclass.Interface]())

	maybeDefer := func(doRegister func()) {
		if isParentExtensionClass && !registeredTypes[superType] {
			// Parent extension class not registered yet, defer this registration
			pendingRegistrations[superType] = append(pendingRegistrations[superType], doRegister)
		} else {
			doRegister()
		}
	}

	switch super.(type) {
	case interface{ AsScript() Script.Instance },
		interface {
			AsEditorPlugin() EditorPlugin.Instance
		},
		interface {
			AsScriptLanguage() ScriptLanguage.Instance
		}:
		gd.EditorStartupFunctions = append(gd.EditorStartupFunctions, func() {
			maybeDefer(register)
		})
	default:
		if gd.Linked {
			maybeDefer(register)
		} else {
			gd.StartupFunctions = append(gd.StartupFunctions, func() {
				maybeDefer(register)
			})
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
		if Engine.IsEditorHint() {
			path := GDExtension.LibraryPath()
			data, err := os.Open(filepath.Join(filepath.Dir(path), "library_documentation.xml"))
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					return
				}

				Engine.Raise(err)
			}
			var dec = xml.NewDecoder(data)
			var docs docgen.XML
			for {
				if err := dec.Decode(&docs); err != nil {
					if err == io.EOF {
						break
					}
					Engine.Raise(fmt.Errorf("failed to unmarshal library documentation: %w", err))
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
		if _, ok := field.Type.MethodByName("AsNode"); ok {
			return
		}
		name := String.ToSnakeCase(field.Name)
		if tag := field.Tag.Get("gd"); tag != "" {
			if tag == "-" {
				return
			}
			name = tag
		}
		if (field.Type.Kind() == reflect.Chan && field.Type.ChanDir() == reflect.SendDir) || reflect.PointerTo(field.Type).Implements(reflect.TypeFor[Signal.Pointer]()) {
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
		var list = gdunsafe.MakePropertyList(1)
		if ptype, ok := propertyOf(className, field, list); ok {
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
			member.Type = ptype.String()
			if !exists {
				class.Members = append(class.Members, *member)
			}
			gdunsafe.RegisterProperty(gdunsafe.StringName(pointers.Get(className)[0]), list, gdunsafe.StringName(pointers.Get(gd.NewStringName(""))[0]), gdunsafe.StringName(pointers.Get(gd.NewStringName(""))[0]))
		}
		list.Free()
	}
	for _, field := range ungroupedFields {
		registerField(field)
	}
	for groupName, fields := range groupedFields {
		gdunsafe.RegisterPropertyGroup(
			gdunsafe.StringName(pointers.Get(className)[0]),
			gdunsafe.String(pointers.Get(gd.NewString(groupName))[0]),
			gdunsafe.String(pointers.Get(gd.NewString(""))[0]),
		)
		for _, field := range fields {
			registerField(field)
		}
	}
	rtype = reflect.PointerTo(rtype)
	for method := range rtype.Methods() {
		name := String.ToSnakeCase(method.Name)
		if rename, ok := method_renames[method.Func.Pointer()]; ok {
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
		if Engine.IsEditorHint() {
			docs, _ := xml.Marshal(class)
			gdunsafe.EditorAddDocumentation(string(docs))
		}
	}).CallDeferred()
}

type classImplementation struct {
	Name        gd.StringName
	Super       gd.StringName
	SuperType   reflect.Type
	EngineClass gd.StringName // The first non-extension ancestor class, cached at registration time

	Tool       bool
	RefCounted bool
	isMainLoop bool

	InEditor bool

	Type reflect.Type

	VirtualMethods func(string) reflect.Value
	Constructor    func() reflect.Value

	Singletons []reflect.StructField
}

// findEngineClass walks up the inheritance chain and returns the name of the
// first class that is not a Go extension class. This must be called after
// parent extension classes are registered (ensured by deferred registration).
func findEngineClass(superType reflect.Type) gd.StringName {
	currentType := superType
	for {
		// Check if this type is a registered extension class
		if !registeredTypes[currentType] {
			// This is not an extension class, so it's a built-in Godot class
			return pointers.Pin(gd.NewStringName(nameOf(currentType)))
		}
		// It's an extension class, get its parent type
		parentType := gdclass.SuperType(reflect.New(currentType).Elem().Interface().(gdclass.Interface))
		if parentType == nil || parentType == currentType {
			// Safety check to prevent infinite loop
			return pointers.Pin(gd.NewStringName(nameOf(currentType)))
		}
		currentType = parentType
	}
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

func (class classImplementation) CreateInstance(notify_postinitialize bool) [1]gdreference.Object {
	return class.CreateInstanceFrom(class.Constructor(), notify_postinitialize, true)
}

func (class classImplementation) CreateInstanceFrom(value reflect.Value, notify_postinitialize bool, add_root bool) [1]gdreference.Object {
	// Use EngineClass (the first built-in ancestor) for construction, not Super.
	// This prevents the issue where extension classes inheriting from other extension
	// classes would trigger multiple calls to set_instance_binding on the same object.
	var super *gdreference.Object = (*gdreference.Object)(value.UnsafePointer())
	gdreference.PinObject(super, gdextension.Object(gdunsafe.MakeObject(gdunsafe.StringName(pointers.Get(class.EngineClass)[0]))))
	if class.RefCounted {
		gd.RefCounted(*super).InitRef()
	}
	instance := class.reloadInstance(value, super)
	gdunsafe.Object(gdreference.GetObject(*super)).ExtensionSetup(gdunsafe.StringName(pointers.Get(class.Name)[0]), &instanceWrapper{impl: instance})
	if keepalive := compile_keepalive(reflect.PointerTo(class.Type)); keepalive != nil {
		roots.Insert(value, keepalive)
	}
	if add_root {
		instance.strong, _ = reflect.TypeAssert[gdclass.Pointer](value)
	}
	instance.cleanup = runtime.AddCleanup(super, func(raw gdextension.Object) {
		Callable.Defer(Callable.New(func() {
			gd.Free(raw)
		}))
	}, gdreference.GetObject(*super))
	for _, field := range class.Singletons {
		if singleton, ok := singletons.Lookup(field.Type.Elem()); ok {
			value.Elem().FieldByIndex(field.Index).Set(singleton)
		}
	}
	if notify_postinitialize {
		gd.ObjectNotification(*super, 0, false)
	}
	instance.OnCreate(value)
	return [1]gdreference.Object{*super}
}

func (class classImplementation) reloadInstance(value reflect.Value, super *gdreference.Object) *instanceImplementation {
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
		name := String.ToSnakeCase(field.Name)
		if tag := field.Tag.Get("gd"); tag != "" {
			name = tag
		}
		name, _, _ = strings.Cut(name, "(")
		// Signal fields need to have their values injected into the field, so that they can be used (emitted).
		if reflect.PointerTo(field.Type).Implements(reflect.TypeFor[Signal.Pointer]()) {
			signal := pointers.Pin(gd.NewSignalOf([1]gdreference.Object{*super}, gd.NewStringName(name)))
			rvalue.Interface().(Signal.Pointer).SetAny(Signal.Via(gd.SignalProxy{}, pointers.Pack(signal)))
			signals = append(signals, signalChan{
				signal: signal,
			})
		}
		if field.Type.Kind() == reflect.Chan && field.Type.ChanDir() == reflect.SendDir {
			signal := pointers.Pin(gd.NewSignalOf([1]gdreference.Object{*super}, gd.NewStringName(name)))
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
		go manageSignals(Object.Instance{*super}.ID(), chSignals)
	}
	return &instanceImplementation{
		object:     gdreference.GetObject(*super),
		Type:       class.Type,
		weak:       weak.Make(super),
		signals:    signals,
		isEditor:   !class.Tool && Engine.IsEditorHint(),
		isMainLoop: class.isMainLoop,
	}
}

func (class classImplementation) GetVirtual(name gd.StringName) any {
	if !class.Tool && class.InEditor {
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
	object  gdextension.Object
	Type    reflect.Type
	strong  gdclass.Pointer
	weak    weak.Pointer[gdreference.Object]
	cleanup runtime.Cleanup
	signals []signalChan

	// FIXME use a bitfield for these booleans.
	isEditor, isMainLoop, freed bool
}

var lastGC int

func (instance *instanceImplementation) Interface() (gdclass.Pointer, bool) {
	if instance.strong != nil {
		return instance.strong, true
	}
	ptr := instance.weak.Value()
	if ptr == nil {
		return nil, false
	}
	return reflect.TypeAssert[gdclass.Pointer](reflect.NewAt(instance.Type, unsafe.Pointer(ptr)))
}

func (instance *instanceImplementation) OnCreate(value reflect.Value) {
	val, ok := instance.Interface()
	if !ok {
		return
	}
	if impl, ok := val.(interface {
		OnCreate()
	}); ok {
		impl.OnCreate()
	}
	if impl, ok := val.(interface {
		OnCreate(value reflect.Value)
	}); ok {
		impl.OnCreate(value)
	}
	if impl, ok := val.(interface {
		Init()
	}); ok {
		impl.Init()
	}
}

func (instance *instanceImplementation) Notification(what Object.Notification, reversed bool) {
	val, ok := instance.Interface()
	if !ok {
		return
	}
	if what == Node.NotificationReady {
		instance.ready()
	}
	if instance.isMainLoop && what == MainLoop.NotificationCrash {
		if idx := ring.CrashIndex; idx != 0xFFFFFFFF {
			e := &ring.Main.Entries[idx]
			if e.PC != 0 {
				fn := runtime.FuncForPC(e.PC)
				if fn != nil {
					file, line := fn.FileLine(e.PC)
					fmt.Fprintf(os.Stderr,
						"crash in ring buffer flush at entry %d: %s (%s:%d)\n",
						idx, fn.Name(), file, line)
				}
			}
		}
		debug.PrintStack()
	}
	if !instance.isEditor {
		switch notify := val.(type) {
		case interface{ Notification(gd.NotificationType) }:
			notify.Notification(gd.NotificationType(what))
		case interface{ Notification(Object.Notification) }:
			notify.Notification(what)
		case interface {
			Notification(Object.Notification, bool)
		}:
			notify.Notification(what, reversed)
		case interface{ Notification(int, bool) }:
			notify.Notification(int(what), reversed)
		default:
		}
	}
}

func (instance *instanceImplementation) ToString() (gd.String, bool) {
	val, ok := instance.Interface()
	if !ok {
		return gd.String{}, false
	}
	switch onfree := val.(type) {
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

func (instance *instanceImplementation) CallVirtual(virtual gd.ExtensionClassCallVirtualFunc, args, back gdextension.Pointer) {
	val, ok := instance.Interface()
	if !ok {
		return
	}
	virtual(val, args, back)
}

func (instance *instanceImplementation) GetRID() gd.RID {
	return 0
}

func (instance *instanceImplementation) Free() {
	if instance.freed {
		return
	}
	val, ok := instance.Interface()
	if !ok {
		return
	}
	instance.cleanup.Stop()
	gdreference.EndObject(val.AsObject()[0])
	roots.Remove(reflect.ValueOf(val))
	for _, signal := range instance.signals {
		if signal.rvalue.IsValid() {
			signal.rvalue.Close()
		}
		signal.signal.Free()
	}
	rvalue := reflect.ValueOf(val).Elem()
	for _, field := range reflect.VisibleFields(rvalue.Type()) {
		if !field.IsExported() || field.Name == "Extension" || rvalue.FieldByIndex(field.Index).IsZero() {
			continue
		}
		type isNode interface {
			AsNode() Node.Instance
		}
		nodeType := reflect.TypeFor[isNode]()
		if field.Type.Implements(nodeType) || reflect.PointerTo(field.Type).Implements(nodeType) {
			continue
		}
		// we need to unreference any pinned resources (pinned here means that the engine `set` them).
		if field.Type.Implements(reflect.TypeFor[RefCounted.Any]()) {
			ref := rvalue.FieldByIndex(field.Index).Interface().(RefCounted.Any).AsRefCounted()[0]
			if ref.Unreference() {
				gdreference.EndObject(gdreference.Object(ref))
				gdunsafe.Object(gdreference.GetObject(gdreference.Object(ref))).Free()
			}
		}
	}
	switch onfree := val.(type) {
	case interface{ OnFree() }:
		onfree.OnFree()
	}
	instance.freed = true
}

func flatFieldsOf(rtype reflect.Type) iter.Seq[reflect.StructField] {
	if rtype.Kind() != reflect.Struct {
		return func(yield func(reflect.StructField) bool) {}
	}
	return func(yield func(reflect.StructField) bool) {
		for field := range rtype.Fields() {
			if field.Name == "Extension" || field.Name == "Singleton" {
				continue
			}
			if field.Anonymous && field.Type.Kind() == reflect.Struct {
				for child := range flatFieldsOf(field.Type) {
					child.Offset += field.Offset
					child.Index = append(field.Index, child.Index...)
					if !yield(child) {
						break
					}
				}
				continue
			}
			if !yield(field) {
				break
			}
		}
	}
}

// ready is responsible for asserting the scene tree for struct members that implement
// Super().AsNode() and asserting that these nodes are added as children to the Super.
//
// TODO this could be partially pre-compiled for a given [Register] type and cached in
// order to avoid any use of reflection at instantiation time.
func (instance *instanceImplementation) ready() {
	val, ok := instance.Interface()
	if !ok {
		return
	}
	parent, ok := Object.As[Node.Instance](Object.Instance(gdclass.GetObjectFromInterface(val)))
	if !ok {
		return
	}
	var rvalue = reflect.ValueOf(val).Elem()
	var front = true
	for field := range flatFieldsOf(rvalue.Type()) {
		if _, hasTag := field.Tag.Lookup("gd"); !hasTag && !field.IsExported() {
			continue
		}
		if field.Type.Kind() == reflect.Pointer {
			if _, ok := singletons.Lookup(field.Type.Elem()); ok {
				continue
			}
		}
		var internal = Node.InternalModeDisabled
		var pointer any
		if field.IsExported() {
			pointer = rvalue.FieldByIndex(field.Index).Addr().Interface()
			front = false
		} else {
			pointer = reflect.NewAt(field.Type, unsafe.Add(rvalue.Addr().UnsafePointer(), field.Offset)).Interface()
			if front {
				internal = Node.InternalModeFront
			} else {
				internal = Node.InternalModeBack
			}
		}
		instance.assertChild(pointer, field, parent, parent, internal)
	}
	if !instance.isEditor {
		val, ok := instance.Interface()
		if !ok {
			return
		}
		switch ready := val.(type) {
		case interface{ Ready() }:
			ready.Ready()
		}
	}
}

func (instance *instanceImplementation) assertChild(value any, field reflect.StructField, parent, owner [1]gdclass.Node, internal Node.InternalMode) {
	type isNode interface {
		AsNode() Node.Instance
	}
	var (
		rvalue = reflect.ValueOf(value)
	)
	nodeType := reflect.TypeFor[isNode]()
	if !field.Type.Implements(nodeType) && !reflect.PointerTo(field.Type).Implements(nodeType) {
		if field.Type.Kind() == reflect.Struct {
			var front = true
			for field := range flatFieldsOf(field.Type) {
				if _, hasTag := field.Tag.Lookup("gd"); !hasTag && !field.IsExported() {
					continue
				}
				var pointer any
				var internal = Node.InternalModeDisabled
				if field.IsExported() {
					pointer = rvalue.Elem().FieldByIndex(field.Index).Addr().Interface()
					front = false
				} else {
					pointer = reflect.NewAt(field.Type, unsafe.Add(rvalue.UnsafePointer(), field.Offset)).Interface()
					if front {
						internal = Node.InternalModeFront
					} else {
						internal = Node.InternalModeBack
					}
				}
				instance.assertChild(pointer, field, parent, owner, internal)
			}
		}
		return
	}
	if field.Anonymous {
		return
	}
	var not_initialised = rvalue.Elem().IsZero()
	if rvalue.Elem().Kind() == reflect.Pointer {
		if rvalue.Elem().IsNil() {
			rvalue.Elem().Set(reflect.New(rvalue.Elem().Type().Elem()))
			not_initialised = true
		}
		value = rvalue.Elem().Interface()
	}
	class := value.(isNode)
	if rvalue.Elem().Kind() == reflect.Struct {
		defer func() {
			var front = true
			for field := range flatFieldsOf(rvalue.Elem().Type()) {
				if _, hasTag := field.Tag.Lookup("gd"); !hasTag && !field.IsExported() {
					continue
				}
				var pointer any
				var internal = Node.InternalModeDisabled
				if field.IsExported() {
					pointer = rvalue.Elem().FieldByIndex(field.Index).Addr().Interface()
					front = false
				} else {
					pointer = reflect.NewAt(field.Type, unsafe.Add(rvalue.UnsafePointer(), field.Offset)).Interface()
					if front {
						internal = Node.InternalModeFront
					} else {
						internal = Node.InternalModeBack
					}
				}
				instance.assertChild(pointer, field, class.AsNode(), owner, internal)
			}
		}()
	}
	name := field.Name
	if tag := field.Tag.Get("gd"); tag != "" {
		if tag == "-" {
			return
		}
		name = tag
	}
	path := Path.ToNode(String.New(name))
	if !Node.Advanced(parent).HasNode(path) {
		if not_initialised {
			child := [1]gdreference.Object{gdreference.OwnObject(gdextension.Object(gdunsafe.MakeObject(gdunsafe.StringName(pointers.Get(gd.NewStringName(nameOf(field.Type)))[0]))), gd.Free)}
			gd.ObjectNotification(child[0], 0, false)
			defer gdreference.EndObject(child[0])
			native := gd.ExtensionInstanceLookup(gdreference.GetObject(child[0]))
			if native != nil {
				rvalue.Elem().Set(reflect.ValueOf(native))
				class = native.(isNode)
			} else {
				class.(gd.IsClassCastable).SetObject([1]gdreference.Object{gdreference.RawObject(gdreference.GetObject(child[0]))})
			}
		}
		var mode Node.InternalMode = Node.InternalModeDisabled | internal
		if !field.IsExported() {
			mode = Node.InternalModeFront
		}
		Node.Advanced(class.AsNode()).SetName(String.Name(String.New(field.Name)))
		Node.Advanced(parent).AddChild(class.AsNode(), true, mode)
		if Engine.IsEditorHint() {
			Node.Advanced(class.AsNode()).SetOwner(EditorInterface.GetEditedSceneRoot())
		}
		return
	}
	var node = Node.Instance(Node.Advanced(parent).GetNode(path))
	native := gd.ExtensionInstanceLookup(gdreference.GetObject(gdclass.GetNode(node[0])[0]))
	if native != nil {
		if reflect.ValueOf(native).Type() == rvalue.Elem().Type() {
			rvalue.Elem().Set(reflect.ValueOf(native))
			gdreference.EndObject(gdclass.GetNode(node[0])[0])
			return
		}
	} else {
		castable, ok := class.(gd.IsClassCastable)
		if ok && castable.SetObject([1]gdreference.Object{gdreference.RawObject(gdreference.GetObject(gdclass.GetNode(node[0])[0]))}) {
			gdreference.EndObject(gdclass.GetNode(node[0])[0])
			return
		}
	}
	// Node exists but has the wrong type, replace it with the correct type.
	Engine.RaiseWarning("graphics.gd DeclarativeChildren[" + nameOf(instance.Type) + "]: converting " + string(Node.Advanced(parent).GetPath().String()) + "/" + field.Name +
		" into " + nameOf(field.Type) + " (previously " + Object.Instance(node.AsObject()).ClassName() + ")")
	if not_initialised {
		child := [1]gdreference.Object{gdreference.OwnObject(gdextension.Object(gdunsafe.MakeObject(gdunsafe.StringName(pointers.Get(gd.NewStringName(nameOf(field.Type)))[0]))), gd.Free)}
		gd.ObjectNotification(child[0], 0, false)
		defer gdreference.EndObject(child[0])
		native := gd.ExtensionInstanceLookup(gdreference.GetObject(child[0]))
		if native != nil {
			rvalue.Elem().Set(reflect.ValueOf(native))
			class = native.(isNode)
		} else {
			class.(gd.IsClassCastable).SetObject([1]gdreference.Object{gdreference.RawObject(gdreference.GetObject(child[0]))})
		}
	}
	Node.Advanced(class.AsNode()).SetName(String.Name(String.New(field.Name)))
	// Copy compatible storable properties from the old node to the replacement.
	newProps := make(map[string]struct{})
	for _, p := range Object.GetPropertyList(class.AsNode()) {
		if p.Usage&int(PropertyUsageStorage) != 0 {
			newProps[p.Name] = struct{}{}
		}
	}
	for _, p := range Object.GetPropertyList(node) {
		if p.Usage&int(PropertyUsageStorage) == 0 {
			continue
		}
		if _, ok := newProps[p.Name]; !ok {
			continue
		}
		val := Object.Get(node, p.Name)
		if val == nil {
			continue
		}
		Object.Set(class.AsNode(), p.Name, val)
	}
	Node.Advanced(node).ReplaceBy(class.AsNode(), true)
	Node.Advanced(node).QueueFree()
	if Engine.IsEditorHint() {
		Node.Advanced(class.AsNode()).SetOwner(EditorInterface.GetEditedSceneRoot())
	}
}
