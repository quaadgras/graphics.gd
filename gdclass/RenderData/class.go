package RenderData

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Abstract render data object, exists for the duration of rendering a single viewport.
[b]Note:[/b] This is an internal rendering server object, do not instantiate this from script.

*/
type Go [1]classdb.RenderData

/*
Returns the [RenderSceneBuffers] object managing the scene buffers for rendering this viewport.
*/
func (self Go) GetRenderSceneBuffers() gdclass.RenderSceneBuffers {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.RenderSceneBuffers(class(self).GetRenderSceneBuffers(gc))
}

/*
Returns the [RenderSceneData] object managing this frames scene data.
*/
func (self Go) GetRenderSceneData() gdclass.RenderSceneData {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.RenderSceneData(class(self).GetRenderSceneData(gc))
}

/*
Returns the [RID] of the environments object in the [RenderingServer] being used to render this viewport.
*/
func (self Go) GetEnvironment() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetEnvironment())
}

/*
Returns the [RID] of the camera attributes object in the [RenderingServer] being used to render this viewport.
*/
func (self Go) GetCameraAttributes() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetCameraAttributes())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RenderData
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RenderData"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Returns the [RenderSceneBuffers] object managing the scene buffers for rendering this viewport.
*/
//go:nosplit
func (self class) GetRenderSceneBuffers(ctx gd.Lifetime) gdclass.RenderSceneBuffers {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderData.Bind_get_render_scene_buffers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.RenderSceneBuffers
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the [RenderSceneData] object managing this frames scene data.
*/
//go:nosplit
func (self class) GetRenderSceneData(ctx gd.Lifetime) gdclass.RenderSceneData {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderData.Bind_get_render_scene_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.RenderSceneData
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the [RID] of the environments object in the [RenderingServer] being used to render this viewport.
*/
//go:nosplit
func (self class) GetEnvironment() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderData.Bind_get_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [RID] of the camera attributes object in the [RenderingServer] being used to render this viewport.
*/
//go:nosplit
func (self class) GetCameraAttributes() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderData.Bind_get_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRenderData() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRenderData() Go { return *((*Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("RenderData", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
