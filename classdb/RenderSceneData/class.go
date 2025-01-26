// Package RenderSceneData provides methods for working with RenderSceneData object instances.
package RenderSceneData

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Projection"
import "graphics.gd/variant/Vector3"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any

/*
Abstract scene data object, exists for the duration of rendering a single viewport.
[b]Note:[/b] This is an internal rendering server object, do not instantiate this from script.
*/
type Instance [1]gdclass.RenderSceneData

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRenderSceneData() Instance
}

/*
Returns the camera transform used to render this frame.
[b]Note:[/b] If more than one view is rendered, this will return a centered transform.
*/
func (self Instance) GetCamTransform() Transform3D.BasisOrigin { //gd:RenderSceneData.get_cam_transform
	return Transform3D.BasisOrigin(class(self).GetCamTransform())
}

/*
Returns the camera projection used to render this frame.
[b]Note:[/b] If more than one view is rendered, this will return a combined projection.
*/
func (self Instance) GetCamProjection() Projection.XYZW { //gd:RenderSceneData.get_cam_projection
	return Projection.XYZW(class(self).GetCamProjection())
}

/*
Returns the number of views being rendered.
*/
func (self Instance) GetViewCount() int { //gd:RenderSceneData.get_view_count
	return int(int(class(self).GetViewCount()))
}

/*
Returns the eye offset per view used to render this frame. This is the offset between our camera transform and the eye transform.
*/
func (self Instance) GetViewEyeOffset(view int) Vector3.XYZ { //gd:RenderSceneData.get_view_eye_offset
	return Vector3.XYZ(class(self).GetViewEyeOffset(gd.Int(view)))
}

/*
Returns the view projection per view used to render this frame.
[b]Note:[/b] If a single view is rendered, this returns the camera projection. If more than one view is rendered, this will return a projection for the given view including the eye offset.
*/
func (self Instance) GetViewProjection(view int) Projection.XYZW { //gd:RenderSceneData.get_view_projection
	return Projection.XYZW(class(self).GetViewProjection(gd.Int(view)))
}

/*
Return the [RID] of the uniform buffer containing the scene data as a UBO.
*/
func (self Instance) GetUniformBuffer() RID.UniformBuffer { //gd:RenderSceneData.get_uniform_buffer
	return RID.UniformBuffer(class(self).GetUniformBuffer())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RenderSceneData

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RenderSceneData"))
	casted := Instance{*(*gdclass.RenderSceneData)(unsafe.Pointer(&object))}
	return casted
}

/*
Returns the camera transform used to render this frame.
[b]Note:[/b] If more than one view is rendered, this will return a centered transform.
*/
//go:nosplit
func (self class) GetCamTransform() gd.Transform3D { //gd:RenderSceneData.get_cam_transform
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneData.Bind_get_cam_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the camera projection used to render this frame.
[b]Note:[/b] If more than one view is rendered, this will return a combined projection.
*/
//go:nosplit
func (self class) GetCamProjection() gd.Projection { //gd:RenderSceneData.get_cam_projection
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Projection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneData.Bind_get_cam_projection, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of views being rendered.
*/
//go:nosplit
func (self class) GetViewCount() gd.Int { //gd:RenderSceneData.get_view_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneData.Bind_get_view_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the eye offset per view used to render this frame. This is the offset between our camera transform and the eye transform.
*/
//go:nosplit
func (self class) GetViewEyeOffset(view gd.Int) gd.Vector3 { //gd:RenderSceneData.get_view_eye_offset
	var frame = callframe.New()
	callframe.Arg(frame, view)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneData.Bind_get_view_eye_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the view projection per view used to render this frame.
[b]Note:[/b] If a single view is rendered, this returns the camera projection. If more than one view is rendered, this will return a projection for the given view including the eye offset.
*/
//go:nosplit
func (self class) GetViewProjection(view gd.Int) gd.Projection { //gd:RenderSceneData.get_view_projection
	var frame = callframe.New()
	callframe.Arg(frame, view)
	var r_ret = callframe.Ret[gd.Projection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneData.Bind_get_view_projection, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Return the [RID] of the uniform buffer containing the scene data as a UBO.
*/
//go:nosplit
func (self class) GetUniformBuffer() gd.RID { //gd:RenderSceneData.get_uniform_buffer
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneData.Bind_get_uniform_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRenderSceneData() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRenderSceneData() Instance { return *((*Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Instance(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("RenderSceneData", func(ptr gd.Object) any {
		return [1]gdclass.RenderSceneData{*(*gdclass.RenderSceneData)(unsafe.Pointer(&ptr))}
	})
}
