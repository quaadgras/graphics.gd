package GDScript

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Script"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A script implemented in the GDScript programming language, saved with the [code].gd[/code] extension. The script extends the functionality of all objects that instantiate it.
Calling [method new] creates a new instance of the script. [method Object.set_script] extends an existing object, if that object's class matches one of the script's base classes.
If you are looking for GDScript's built-in functions, see [@GDScript] instead.

*/
type Go [1]classdb.GDScript
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GDScript
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("GDScript"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsGDScript() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGDScript() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsScript() Script.GD { return *((*Script.GD)(unsafe.Pointer(&self))) }
func (self Go) AsScript() Script.Go { return *((*Script.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsScript(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsScript(), name)
	}
}
func init() {classdb.Register("GDScript", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
