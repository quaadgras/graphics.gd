package RenderData

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Abstract render data object, exists for the duration of rendering a single viewport.
[b]Note:[/b] This is an internal rendering server object, do not instantiate this from script.

*/
type Simple [1]classdb.RenderData
func (self Simple) GetRenderSceneBuffers() [1]classdb.RenderSceneBuffers {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.RenderSceneBuffers(Expert(self).GetRenderSceneBuffers(gc))
}
func (self Simple) GetRenderSceneData() [1]classdb.RenderSceneData {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.RenderSceneData(Expert(self).GetRenderSceneData(gc))
}
func (self Simple) GetEnvironment() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetEnvironment())
}
func (self Simple) GetCameraAttributes() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetCameraAttributes())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RenderData
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the [RenderSceneBuffers] object managing the scene buffers for rendering this viewport.
*/
//go:nosplit
func (self class) GetRenderSceneBuffers(ctx gd.Lifetime) object.RenderSceneBuffers {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderData.Bind_get_render_scene_buffers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.RenderSceneBuffers
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the [RenderSceneData] object managing this frames scene data.
*/
//go:nosplit
func (self class) GetRenderSceneData(ctx gd.Lifetime) object.RenderSceneData {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderData.Bind_get_render_scene_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.RenderSceneData
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

//go:nosplit
func (self class) AsRenderData() Expert { return self[0].AsRenderData() }


//go:nosplit
func (self Simple) AsRenderData() Simple { return self[0].AsRenderData() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("RenderData", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
