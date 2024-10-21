package RenderSceneData

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
Abstract scene data object, exists for the duration of rendering a single viewport.
[b]Note:[/b] This is an internal rendering server object, do not instantiate this from script.

*/
type Simple [1]classdb.RenderSceneData
func (self Simple) GetCamTransform() gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).GetCamTransform())
}
func (self Simple) GetCamProjection() gd.Projection {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Projection(Expert(self).GetCamProjection())
}
func (self Simple) GetViewCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetViewCount()))
}
func (self Simple) GetViewEyeOffset(view int) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetViewEyeOffset(gd.Int(view)))
}
func (self Simple) GetViewProjection(view int) gd.Projection {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Projection(Expert(self).GetViewProjection(gd.Int(view)))
}
func (self Simple) GetUniformBuffer() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetUniformBuffer())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RenderSceneData
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the camera transform used to render this frame.
[b]Note:[/b] If more than one view is rendered, this will return a centered transform.
*/
//go:nosplit
func (self class) GetCamTransform() gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneData.Bind_get_cam_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the camera projection used to render this frame.
[b]Note:[/b] If more than one view is rendered, this will return a combined projection.
*/
//go:nosplit
func (self class) GetCamProjection() gd.Projection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Projection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneData.Bind_get_cam_projection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of views being rendered.
*/
//go:nosplit
func (self class) GetViewCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneData.Bind_get_view_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the eye offset per view used to render this frame. This is the offset between our camera transform and the eye transform.
*/
//go:nosplit
func (self class) GetViewEyeOffset(view gd.Int) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, view)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneData.Bind_get_view_eye_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the view projection per view used to render this frame.
[b]Note:[/b] If a single view is rendered, this returns the camera projection. If more than one view is rendered, this will return a projection for the given view including the eye offset.
*/
//go:nosplit
func (self class) GetViewProjection(view gd.Int) gd.Projection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, view)
	var r_ret = callframe.Ret[gd.Projection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneData.Bind_get_view_projection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Return the [RID] of the uniform buffer containing the scene data as a UBO.
*/
//go:nosplit
func (self class) GetUniformBuffer() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneData.Bind_get_uniform_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsRenderSceneData() Expert { return self[0].AsRenderSceneData() }


//go:nosplit
func (self Simple) AsRenderSceneData() Simple { return self[0].AsRenderSceneData() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("RenderSceneData", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
