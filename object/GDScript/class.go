package GDScript

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Script"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A script implemented in the GDScript programming language, saved with the [code].gd[/code] extension. The script extends the functionality of all objects that instantiate it.
Calling [method new] creates a new instance of the script. [method Object.set_script] extends an existing object, if that object's class matches one of the script's base classes.
If you are looking for GDScript's built-in functions, see [@GDScript] instead.

*/
type Simple [1]classdb.GDScript
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.GDScript
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsGDScript() Expert { return self[0].AsGDScript() }


//go:nosplit
func (self Simple) AsGDScript() Simple { return self[0].AsGDScript() }


//go:nosplit
func (self class) AsScript() Script.Expert { return self[0].AsScript() }


//go:nosplit
func (self Simple) AsScript() Script.Simple { return self[0].AsScript() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("GDScript", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
