package RenderSceneData

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
Abstract scene data object, exists for the duration of rendering a single viewport.
[b]Note:[/b] This is an internal rendering server object, do not instantiate this from script.
*/
type Instance [1]classdb.RenderSceneData

/*
Returns the camera transform used to render this frame.
[b]Note:[/b] If more than one view is rendered, this will return a centered transform.
*/
func (self Instance) GetCamTransform() gd.Transform3D {
	return gd.Transform3D(class(self).GetCamTransform())
}

/*
Returns the camera projection used to render this frame.
[b]Note:[/b] If more than one view is rendered, this will return a combined projection.
*/
func (self Instance) GetCamProjection() gd.Projection {
	return gd.Projection(class(self).GetCamProjection())
}

/*
Returns the number of views being rendered.
*/
func (self Instance) GetViewCount() int {
	return int(int(class(self).GetViewCount()))
}

/*
Returns the eye offset per view used to render this frame. This is the offset between our camera transform and the eye transform.
*/
func (self Instance) GetViewEyeOffset(view int) gd.Vector3 {
	return gd.Vector3(class(self).GetViewEyeOffset(gd.Int(view)))
}

/*
Returns the view projection per view used to render this frame.
[b]Note:[/b] If a single view is rendered, this returns the camera projection. If more than one view is rendered, this will return a projection for the given view including the eye offset.
*/
func (self Instance) GetViewProjection(view int) gd.Projection {
	return gd.Projection(class(self).GetViewProjection(gd.Int(view)))
}

/*
Return the [RID] of the uniform buffer containing the scene data as a UBO.
*/
func (self Instance) GetUniformBuffer() gd.RID {
	return gd.RID(class(self).GetUniformBuffer())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RenderSceneData

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RenderSceneData"))
	return Instance{classdb.RenderSceneData(object)}
}

/*
Returns the camera transform used to render this frame.
[b]Note:[/b] If more than one view is rendered, this will return a centered transform.
*/
//go:nosplit
func (self class) GetCamTransform() gd.Transform3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneData.Bind_get_cam_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Projection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneData.Bind_get_cam_projection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of views being rendered.
*/
//go:nosplit
func (self class) GetViewCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneData.Bind_get_view_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the eye offset per view used to render this frame. This is the offset between our camera transform and the eye transform.
*/
//go:nosplit
func (self class) GetViewEyeOffset(view gd.Int) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, view)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneData.Bind_get_view_eye_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, view)
	var r_ret = callframe.Ret[gd.Projection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneData.Bind_get_view_projection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Return the [RID] of the uniform buffer containing the scene data as a UBO.
*/
//go:nosplit
func (self class) GetUniformBuffer() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneData.Bind_get_uniform_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRenderSceneData() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRenderSceneData() Instance { return *((*Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("RenderSceneData", func(ptr gd.Object) any { return classdb.RenderSceneData(ptr) })
}
