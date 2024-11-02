package RenderSceneDataRD

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/RenderSceneData"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Object holds scene data related to rendering a single frame of a viewport.
[b]Note:[/b] This is an internal rendering server object, do not instantiate this from script.
*/
type Instance [1]classdb.RenderSceneDataRD

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RenderSceneDataRD

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RenderSceneDataRD"))
	return Instance{classdb.RenderSceneDataRD(object)}
}

func (self class) AsRenderSceneDataRD() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRenderSceneDataRD() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRenderSceneData() RenderSceneData.Advanced {
	return *((*RenderSceneData.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRenderSceneData() RenderSceneData.Instance {
	return *((*RenderSceneData.Instance)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRenderSceneData(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRenderSceneData(), name)
	}
}
func init() {
	classdb.Register("RenderSceneDataRD", func(ptr gd.Object) any { return classdb.RenderSceneDataRD(ptr) })
}
