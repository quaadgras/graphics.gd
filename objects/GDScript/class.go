package GDScript

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Script"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A script implemented in the GDScript programming language, saved with the [code].gd[/code] extension. The script extends the functionality of all objects that instantiate it.
Calling [method new] creates a new instance of the script. [method Object.set_script] extends an existing object, if that object's class matches one of the script's base classes.
If you are looking for GDScript's built-in functions, see [@GDScript] instead.
*/
type Instance [1]classdb.GDScript

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.GDScript

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GDScript"))
	return Instance{classdb.GDScript(object)}
}

func (self class) AsGDScript() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGDScript() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsScript() Script.Advanced    { return *((*Script.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsScript() Script.Instance { return *((*Script.Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsScript(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsScript(), name)
	}
}
func init() { classdb.Register("GDScript", func(ptr gd.Object) any { return classdb.GDScript(ptr) }) }
