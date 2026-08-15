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
	"sync/atomic"
	"unsafe"
	"weak"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"graphics.gd/classdb/EditorInterface"
	"graphics.gd/classdb/EditorPlugin"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/MainLoop"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/Script"
	"graphics.gd/classdb/ScriptLanguage"
	"graphics.gd/classdb/ShaderMaterial"

	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Path"
	"graphics.gd/variant/RefCounted"
	"graphics.gd/variant/Signal"
	"graphics.gd/variant/String"

	gd "graphics.gd/internal"
	"graphics.gd/internal/docgen"
	"graphics.gd/internal/gdclass"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdreference"
	"graphics.gd/internal/notifyfilter"
	"graphics.gd/internal/pointers"
	"graphics.gd/internal/ring"
	"graphics.gd/internal/threadsafe"
)

var classes threadsafe.Handles[*classImplementation, gdextension.ExtensionClassID]

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
	if registrationDisabled {
		return
	}
	var superType = gdclass.SuperType(([1]T{})[0])
	var super = reflect.New(superType).Elem().Interface()
	var classType = reflect.TypeFor[T]()
	if handlesNotifications(classType) {
		notifyfilter.Want()
	}

	// Extensions of editor-only classes can only be registered once the engine
	// has registered its own editor classes (at the editor initialization
	// level) and their virtual methods only ever run inside the editor, so
	// they are implicitly tool classes.
	var editorOnly = extendsEditorClass(superType)

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
		var tool = editorOnly
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
			tab:            classTab(classType),
			VirtualMethods: reference.Virtual,
			Constructor: func() reflect.Value {
				return reflect.New(classType)
			},
			virtuals: new(threadsafe.Map[gdextension.StringName, resolvedVirtual]),
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
		gdextension.Host.ClassDB.Register.Class(pointers.Get(className), pointers.Get(superName), classes.New(impl), false, false, true, false, pointers.Get(iconString))

		gd.RegisterCleanup(func() {
			gdextension.Host.ClassDB.Register.Removal(pointers.Get(className))
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
				gdextension.Host.Editor.AddPlugin(pointers.Get(className))
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

	var deferToEditor = editorOnly
	switch super.(type) {
	case interface{ AsScript() Script.Instance },
		interface {
			AsEditorPlugin() EditorPlugin.Instance
		},
		interface {
			AsScriptLanguage() ScriptLanguage.Instance
		}:
		deferToEditor = true
	}
	if deferToEditor {
		if gd.LinkedEditor {
			maybeDefer(register)
		} else {
			gd.EditorStartupFunctions = append(gd.EditorStartupFunctions, func() {
				maybeDefer(register)
			})
		}
	} else {
		if gd.Linked {
			maybeDefer(register)
		} else {
			gd.StartupFunctions = append(gd.StartupFunctions, func() {
				maybeDefer(register)
			})
		}
	}
}

// extendsEditorClass reports whether the given super type descends from an
// engine class that is only available inside the editor. Such classes are
// registered by the engine at the editor initialization level, so extensions
// of them cannot be registered any earlier than that. Unlike
// [findEngineClass], this works before any parent extension classes have
// been registered, by walking the Go types alone.
func extendsEditorClass(superType reflect.Type) bool {
	currentType := superType
	for {
		iface, ok := reflect.New(currentType).Elem().Interface().(gdclass.Interface)
		if !ok {
			iface, ok = reflect.New(currentType).Interface().(gdclass.Interface)
		}
		if !ok {
			// Not an extension class, so it's a built-in Godot class.
			return gdclass.EditorClasses[nameOf(currentType)]
		}
		parentType := gdclass.SuperType(iface)
		if parentType == nil || parentType == currentType {
			return gdclass.EditorClasses[nameOf(currentType)]
		}
		currentType = parentType
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
			path := pointers.New[gd.String](gdextension.Host.Library.Location())
			data, err := os.Open(filepath.Join(filepath.Dir(path.String()), "library_documentation.xml"))
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
			gdextension.Host.Editor.AddDocumentation(string(docs))
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

	// tab is the static itab of (*Type, gdclass.Pointer), resolved once at
	// registration so virtual dispatch can rebuild the receiver interface
	// from (tab, instance word) without touching the instance record. Nil
	// on builds where the interface layout is unknown (iface_portable.go).
	tab unsafe.Pointer

	VirtualMethods func(string) reflect.Value
	Constructor    func() reflect.Value

	Singletons []reflect.StructField

	// virtuals caches GetVirtual results keyed by the engine's interned
	// method name. The engine resolves virtuals per instance, but the
	// answer depends only on the class and the name: without the cache
	// every node entering the tree pays a name conversion and a String
	// round-trip per virtual method queried. The result is boxed so a
	// cached "no such virtual" is distinguishable from a missing entry.
	virtuals *threadsafe.Map[gdextension.StringName, resolvedVirtual]
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
	// A custom constructor may have already instantiated the engine side of
	// this value: calling any method on the object before returning it does
	// that (through Extension.createObject). Registering the value again
	// would orphan the first registration while its dispatch word is still
	// pinned — leaking the pin (a fatal runtime.Pinner finalizer panic once
	// the orphan is collected) and an engine object — so hand back the
	// existing instance instead.
	if raw := gdreference.GetObject(*super); raw != 0 {
		if id := gdextension.Host.Objects.Extension.Fetch(raw); id != 0 {
			if existing := instances.Get(id); existing != nil {
				if add_root {
					s, _ := reflect.TypeAssert[gdclass.Pointer](value)
					existing.setStrong(s)
				}
				return [1]gdreference.Object{*super}
			}
		}
	}
	// Objects.Make is classdb_construct_object3, which already establishes the
	// initial reference (refcount=1) for RefCounted built-in ancestors. This also
	// satisfies the create_instance3 contract (Godot expects the creation func to
	// return RefCounted with refcount=1), so we must NOT InitRef again here.
	gdreference.PinObject(super, gdextension.Host.Objects.Make(pointers.Get(class.EngineClass)))
	instance := class.reloadInstance(value, super)
	id := instances.New(instance, value)
	gdextension.Host.Objects.Extension.Setup(gdreference.GetObject(*super), pointers.Get(class.Name), id)
	if keepalive := compile_keepalive(reflect.PointerTo(class.Type)); keepalive != nil {
		roots.Insert(value, keepalive)
	}
	if add_root {
		s, _ := reflect.TypeAssert[gdclass.Pointer](value)
		instance.setStrong(s)
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
			rvalue.Interface().(Signal.Pointer).SetAny(Signal.Via(gd.WrapSignal(signal)))
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
	impl := &instanceImplementation{
		object:     gdreference.GetObject(*super),
		Type:       class.Type,
		weak:       weak.Make(super),
		signals:    signals,
		isEditor:   !class.Tool && Engine.IsEditorHint(),
		isMainLoop: class.isMainLoop,
	}
	impl.engineMemory = allocateFields(value.Type(), value.Addr().UnsafePointer())
	return impl
}

// GetVirtual resolves the implementation of a virtual method by its
// engine name. The engine asks per instance, but the answer depends only
// on the class and the (interned, so pointer-stable) name, so it is
// resolved once and cached: without the cache every node entering the
// tree pays name conversion and String round-trips per virtual queried.
type resolvedVirtual struct{ value any }

func (class classImplementation) GetVirtual(name gd.StringName) any {
	key := gdextension.StringName(pointers.Get(name))
	if cached, ok := class.virtuals.Lookup(key); ok {
		return cached.value
	}
	virtual := class.getVirtual(name)
	class.virtuals.Insert(key, resolvedVirtual{value: virtual})
	return virtual
}

// tickOf resolves the direct call target behind [pinnedVirtualFunc.tick], or
// nil for any virtual that does not qualify. Only _process and
// _physics_process are worth specialising — they are the two virtuals a scene
// calls once per node per frame — and only when the Go method has exactly the
// shape the shortcut can call: one Float.X argument, no results. Anything
// else, including a method the engine knows but Go declares differently, is
// left to the generic wrapper.
func (class classImplementation) tickOf(name string) func(unsafe.Pointer, Float.X) {
	if tickDisabled {
		return nil
	}
	switch name {
	case "_process", "_physics_process":
	default:
		return nil
	}
	// The shortcut hands the engine's instance word straight to the method as
	// its receiver, which is only the address of the user's struct on builds
	// where classTab resolved an itab (see instanceID). Where it did not, the
	// word is an opaque counter and only the instances table can turn it back
	// into a receiver, so the generic path has to take the call.
	if class.tab == nil {
		return nil
	}
	if !class.Tool && class.InEditor {
		return nil
	}
	method, ok := reflect.PointerTo(class.Type).MethodByName(convertName(name))
	if !ok {
		return nil
	}
	if method.Type.NumIn() != 2 || method.Type.NumOut() != 0 || method.Type.In(1) != reflect.TypeFor[Float.X]() {
		return nil
	}
	// Reinterpret the method as a call whose receiver is a raw pointer: a Go
	// func value is one word wide whatever its signature, and the instance
	// word the engine hands back on every callback IS the address of the
	// user's struct (see instanceID), which is what the receiver expects.
	// This is the same reinterpretation getVirtual makes to fit a method to
	// the engine's virtual signature.
	held := reflect.New(method.Type)
	held.Elem().Set(method.Func)
	return *(*func(unsafe.Pointer, Float.X))(held.UnsafePointer())
}

func (class classImplementation) getVirtual(name gd.StringName) any {
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
	object gdextension.Object
	Type   reflect.Type
	// strong roots the wrapper against Go's GC while the engine owns
	// references. Accessed via strongInterface/setStrong: ownership
	// transfers (ExtensionInstanceGoOnly) write it from user goroutines
	// while main-thread callbacks read it, and an interface value cannot
	// be read atomically without the box.
	strong  atomic.Pointer[gdclass.Pointer]
	weak    weak.Pointer[gdreference.Object]
	cleanup runtime.Cleanup
	signals []signalChan

	// engineMemory holds the allocations behind the class's
	// [Engine.Allocated] fields, handed back when the engine frees the
	// instance (instanceTable.Del).
	engineMemory []gdextension.Pointer

	// itab caches the interface type-word for (*Type, gdclass.Pointer). It is
	// static per class, so once resolved every subsequent [Interface] call can
	// rebuild the interface from (itab, live weak pointer) without reflection.
	// Storing only the static itab — never the data pointer — keeps the weak
	// reference weak (it does not pin the Go object against GC). Only used on
	// builds where the interface layout is known, see iface_gc.go.
	itab unsafe.Pointer

	// FIXME use a bitfield for these booleans.
	isEditor, isMainLoop, freed bool

	// pinner pins the user's struct while the engine holds its address as
	// the instance's dispatch word (see instanceID in iface_gc.go). Held
	// from creation until the engine frees the instance, released early if
	// Go takes sole ownership of the object (ExtensionInstanceGoOnly).
	// Unused on portable builds, where the dispatch word is opaque.
	pinner runtime.Pinner
}

var lastGC int

// strongInterface reads the strong root; nil when the wrapper is weakly
// held (Go owns the object).
func (instance *instanceImplementation) strongInterface() gdclass.Pointer {
	if p := instance.strong.Load(); p != nil {
		return *p
	}
	return nil
}

// setStrong publishes (or clears, with nil) the strong root.
func (instance *instanceImplementation) setStrong(iface gdclass.Pointer) {
	if iface == nil {
		instance.strong.Store(nil)
		return
	}
	instance.strong.Store(&iface)
}

func (instance *instanceImplementation) Interface() (gdclass.Pointer, bool) {
	if s := instance.strongInterface(); s != nil {
		return s, true
	}
	ptr := instance.weak.Value()
	if ptr == nil {
		return nil, false
	}
	// Fast path: rebuild the interface from the cached static itab and the
	// live pointer, avoiding reflect on every virtual-method dispatch.
	if iface, ok := instance.cachedInterface(unsafe.Pointer(ptr)); ok {
		return iface, true
	}
	iface, ok := reflect.TypeAssert[gdclass.Pointer](reflect.NewAt(instance.Type, unsafe.Pointer(ptr)))
	if !ok {
		return nil, false
	}
	instance.cacheInterface(iface)
	return iface, true
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

// handlesNotifications reports whether *T implements any of the Notification
// interfaces dispatched by [instanceImplementation.Notification]'s type switch.
// Classes that don't are served by the engine-side filter in gd.c, which drops
// per-frame process-tick notifications before they cross into Go; registering
// a class that does implement one disables that filter (see notifyfilter).
func handlesNotifications(rtype reflect.Type) bool {
	ptr := reflect.PointerTo(rtype)
	return ptr.Implements(reflect.TypeFor[interface{ Notification(gd.NotificationType) }]()) ||
		ptr.Implements(reflect.TypeFor[interface{ Notification(Object.Notification) }]()) ||
		ptr.Implements(reflect.TypeFor[interface {
			Notification(Object.Notification, bool)
		}]()) ||
		ptr.Implements(reflect.TypeFor[interface{ Notification(int, bool) }]())
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
		for _, handler := range gd.CrashHandlers {
			handler()
		}
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
				gdextension.Host.Objects.Unsafe.Free(gdreference.GetObject(gdreference.Object(ref)))
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

// compositeEmbed reports whether t is a struct type with an
// anonymous embedded pointer to a Godot-class wrapper. Returns the
// embedded field on success. See [assertCompositeChild] for the
// pattern this enables.
func compositeEmbed(t, nodeType reflect.Type) (reflect.StructField, bool) {
	if t.Kind() != reflect.Struct {
		return reflect.StructField{}, false
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.Anonymous {
			continue
		}
		// Anonymous Extension[T] embeds (which is how every
		// graphics.gd-registered class declares its parent) live
		// inside Struct, not Pointer — they're stamped values, not
		// pointers. Only treat anonymous *Pointer* embeds as the
		// composite marker, so we don't mistakenly route every
		// graphics.gd class through the composite path.
		if f.Type.Kind() == reflect.Pointer && f.Type.Implements(nodeType) {
			return f, true
		}
	}
	return reflect.StructField{}, false
}

// assertCompositeChild handles the composite-node-binding case
// flagged by [compositeEmbed]: the scene node identified by
// field.Name (or its `gd:"..."` tag) is bound to the anonymous
// embedded pointer of value's struct type, and the struct's
// remaining named fields are then resolved recursively as children
// of that node.
//
// Tag handling and the "create if missing" fallback mirror the
// non-composite path below — we just route the assignment to the
// embedded pointer field instead of the struct as a whole.
func (instance *instanceImplementation) assertCompositeChild(value any, field, embed reflect.StructField, parent, owner [1]gdclass.Node) {
	type isNode interface {
		AsNode() Node.Instance
	}
	rvalue := reflect.ValueOf(value)
	name := field.Name
	if tag := field.Tag.Get("gd"); tag != "" {
		if tag == "-" {
			return
		}
		name = tag
	}
	path := Path.ToNode(String.New(name))
	if !Node.Advanced(parent).HasNode(path) {
		Engine.RaiseWarning("classdb: composite-binding field " + field.Name +
			" expects scene child " + name + " but none was found; skipping")
		return
	}
	node := Node.Instance(Node.Advanced(parent).GetNode(path))
	native := gd.ExtensionInstanceLookup(gdreference.GetObject(gdclass.GetNode(node[0])[0]))
	if native == nil {
		return
	}
	nativeValue := reflect.ValueOf(native)
	embedField := rvalue.Elem().FieldByIndex(embed.Index)
	if nativeValue.Type() != embedField.Type() {
		Engine.RaiseWarning("classdb: composite-binding field " + field.Name +
			" has embedded " + embedField.Type().String() +
			" but scene node " + name + " is " + nativeValue.Type().String())
		return
	}
	embedField.Set(nativeValue)
	// Recurse into the struct's remaining named fields, treating
	// them as children of the just-bound node. We deliberately skip
	// other anonymous embeds — they're either the one we just bound
	// or aren't part of this binding pattern.
	for i := 0; i < field.Type.NumField(); i++ {
		sf := field.Type.Field(i)
		if sf.Anonymous {
			continue
		}
		if _, hasTag := sf.Tag.Lookup("gd"); !hasTag && !sf.IsExported() {
			continue
		}
		ptr := rvalue.Elem().FieldByIndex(sf.Index).Addr().Interface()
		instance.assertChild(ptr, sf, node, owner, Node.InternalModeDisabled)
	}
	gdreference.EndObject(gdclass.GetNode(node[0])[0])
}

func (instance *instanceImplementation) assertChild(value any, field reflect.StructField, parent, owner [1]gdclass.Node, internal Node.InternalMode) {
	type isNode interface {
		AsNode() Node.Instance
	}
	var (
		rvalue = reflect.ValueOf(value)
	)
	nodeType := reflect.TypeFor[isNode]()

	// "Composite" node binding: a struct field whose type contains
	// an anonymous embedded pointer to a Godot class. Lets users
	// write declarations like
	//
	//     Toolbar struct {
	//         *Triangle           // bound to the scene node
	//         Settings TextureButton.Instance  // child of Triangle
	//         Undo     TextureButton.Instance
	//     }
	//
	// where the scene tree has a Triangle named "Toolbar" with
	// TextureButton children. Without this branch the field would
	// trip the standard assignment path below — node's *Triangle
	// type vs the wrapper struct type don't match — and panic with
	// "reflect.Set: value of type *Triangle is not assignable to
	// type struct {...}".
	if embed, ok := compositeEmbed(field.Type, nodeType); ok {
		instance.assertCompositeChild(value, field, embed, parent, owner)
		return
	}
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
			child := [1]gdreference.Object{gdreference.OwnObject(gdextension.Host.Objects.Make(pointers.Get(gd.NewStringName(nameOf(field.Type)))), gd.Free)}
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
		child := [1]gdreference.Object{gdreference.OwnObject(gdextension.Host.Objects.Make(pointers.Get(gd.NewStringName(nameOf(field.Type)))), gd.Free)}
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
