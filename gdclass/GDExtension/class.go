package GDExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
The [GDExtension] resource type represents a [url=https://en.wikipedia.org/wiki/Shared_library]shared library[/url] which can expand the functionality of the engine. The [GDExtensionManager] singleton is responsible for loading, reloading, and unloading [GDExtension] resources.
[b]Note:[/b] GDExtension itself is not a scripting language and has no relation to [GDScript] resources.
*/
type Instance [1]classdb.GDExtension

/*
Returns [code]true[/code] if this extension's library has been opened.
*/
func (self Instance) IsLibraryOpen() bool {
	return bool(class(self).IsLibraryOpen())
}

/*
Returns the lowest level required for this extension to be properly initialized (see the [enum InitializationLevel] enum).
*/
func (self Instance) GetMinimumLibraryInitializationLevel() gd.GDExtensionInitializationLevel {
	return gd.GDExtensionInitializationLevel(class(self).GetMinimumLibraryInitializationLevel())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.GDExtension

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GDExtension"))
	return Instance{classdb.GDExtension(object)}
}

/*
Returns [code]true[/code] if this extension's library has been opened.
*/
//go:nosplit
func (self class) IsLibraryOpen() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GDExtension.Bind_is_library_open, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the lowest level required for this extension to be properly initialized (see the [enum InitializationLevel] enum).
*/
//go:nosplit
func (self class) GetMinimumLibraryInitializationLevel() gd.GDExtensionInitializationLevel {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.GDExtensionInitializationLevel](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GDExtension.Bind_get_minimum_library_initialization_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsGDExtension() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGDExtension() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	classdb.Register("GDExtension", func(ptr gd.Object) any { return classdb.GDExtension(ptr) })
}

type InitializationLevel = classdb.GDExtensionInitializationLevel

const (
	/*The library is initialized at the same time as the core features of the engine.*/
	InitializationLevelCore InitializationLevel = 0
	/*The library is initialized at the same time as the engine's servers (such as [RenderingServer] or [PhysicsServer3D]).*/
	InitializationLevelServers InitializationLevel = 1
	/*The library is initialized at the same time as the engine's scene-related classes.*/
	InitializationLevelScene InitializationLevel = 2
	/*The library is initialized at the same time as the engine's editor classes. Only happens when loading the GDExtension in the editor.*/
	InitializationLevelEditor InitializationLevel = 3
)
