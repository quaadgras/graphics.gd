package RenderSceneDataRD

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/RenderSceneData"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Object holds scene data related to rendering a single frame of a viewport.
[b]Note:[/b] This is an internal rendering server object, do not instantiate this from script.

*/
type Simple [1]classdb.RenderSceneDataRD
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RenderSceneDataRD
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsRenderSceneDataRD() Expert { return self[0].AsRenderSceneDataRD() }


//go:nosplit
func (self Simple) AsRenderSceneDataRD() Simple { return self[0].AsRenderSceneDataRD() }


//go:nosplit
func (self class) AsRenderSceneData() RenderSceneData.Expert { return self[0].AsRenderSceneData() }


//go:nosplit
func (self Simple) AsRenderSceneData() RenderSceneData.Simple { return self[0].AsRenderSceneData() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("RenderSceneDataRD", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
