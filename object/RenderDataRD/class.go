package RenderDataRD

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/RenderData"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This object manages all render data for the rendering device based renderers.
[b]Note:[/b] This is an internal rendering server object only exposed for GDExtension plugins.

*/
type Simple [1]classdb.RenderDataRD
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RenderDataRD
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsRenderDataRD() Expert { return self[0].AsRenderDataRD() }


//go:nosplit
func (self Simple) AsRenderDataRD() Simple { return self[0].AsRenderDataRD() }


//go:nosplit
func (self class) AsRenderData() RenderData.Expert { return self[0].AsRenderData() }


//go:nosplit
func (self Simple) AsRenderData() RenderData.Simple { return self[0].AsRenderData() }

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
func init() {classdb.Register("RenderDataRD", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
