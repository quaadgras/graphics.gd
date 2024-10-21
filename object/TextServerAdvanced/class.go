package TextServerAdvanced

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/TextServerExtension"
import "grow.graphics/gd/object/TextServer"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
An implementation of [TextServer] that uses HarfBuzz, ICU and SIL Graphite to support BiDi, complex text layouts and contextual OpenType features. This is Godot's default primary [TextServer] interface.

*/
type Simple [1]classdb.TextServerAdvanced
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.TextServerAdvanced
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsTextServerAdvanced() Expert { return self[0].AsTextServerAdvanced() }


//go:nosplit
func (self Simple) AsTextServerAdvanced() Simple { return self[0].AsTextServerAdvanced() }


//go:nosplit
func (self class) AsTextServerExtension() TextServerExtension.Expert { return self[0].AsTextServerExtension() }


//go:nosplit
func (self Simple) AsTextServerExtension() TextServerExtension.Simple { return self[0].AsTextServerExtension() }


//go:nosplit
func (self class) AsTextServer() TextServer.Expert { return self[0].AsTextServer() }


//go:nosplit
func (self Simple) AsTextServer() TextServer.Simple { return self[0].AsTextServer() }


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
func init() {classdb.Register("TextServerAdvanced", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
