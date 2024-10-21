package GLTFCamera

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Represents a camera as defined by the base GLTF spec.

*/
type Simple [1]classdb.GLTFCamera
func (self Simple) FromNode(camera_node [1]classdb.Camera3D) [1]classdb.GLTFCamera {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.GLTFCamera(Expert(self).FromNode(gc, camera_node))
}
func (self Simple) ToNode() [1]classdb.Camera3D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Camera3D(Expert(self).ToNode(gc))
}
func (self Simple) FromDictionary(dictionary gd.Dictionary) [1]classdb.GLTFCamera {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.GLTFCamera(Expert(self).FromDictionary(gc, dictionary))
}
func (self Simple) ToDictionary() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).ToDictionary(gc))
}
func (self Simple) GetPerspective() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetPerspective())
}
func (self Simple) SetPerspective(perspective bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPerspective(perspective)
}
func (self Simple) GetFov() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFov()))
}
func (self Simple) SetFov(fov float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFov(gd.Float(fov))
}
func (self Simple) GetSizeMag() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSizeMag()))
}
func (self Simple) SetSizeMag(size_mag float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSizeMag(gd.Float(size_mag))
}
func (self Simple) GetDepthFar() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDepthFar()))
}
func (self Simple) SetDepthFar(zdepth_far float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDepthFar(gd.Float(zdepth_far))
}
func (self Simple) GetDepthNear() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDepthNear()))
}
func (self Simple) SetDepthNear(zdepth_near float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDepthNear(gd.Float(zdepth_near))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.GLTFCamera
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Create a new GLTFCamera instance from the given Godot [Camera3D] node.
*/
//go:nosplit
func (self class) FromNode(ctx gd.Lifetime, camera_node object.Camera3D) object.GLTFCamera {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(camera_node[0].AsPointer())[0])
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.GLTFCamera.Bind_from_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.GLTFCamera
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Converts this GLTFCamera instance into a Godot [Camera3D] node.
*/
//go:nosplit
func (self class) ToNode(ctx gd.Lifetime) object.Camera3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFCamera.Bind_to_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Camera3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates a new GLTFCamera instance by parsing the given [Dictionary].
*/
//go:nosplit
func (self class) FromDictionary(ctx gd.Lifetime, dictionary gd.Dictionary) object.GLTFCamera {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(dictionary))
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.GLTFCamera.Bind_from_dictionary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.GLTFCamera
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Serializes this GLTFCamera instance into a [Dictionary].
*/
//go:nosplit
func (self class) ToDictionary(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFCamera.Bind_to_dictionary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetPerspective() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFCamera.Bind_get_perspective, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPerspective(perspective bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, perspective)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFCamera.Bind_set_perspective, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFov() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFCamera.Bind_get_fov, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFov(fov gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fov)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFCamera.Bind_set_fov, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSizeMag() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFCamera.Bind_get_size_mag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSizeMag(size_mag gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size_mag)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFCamera.Bind_set_size_mag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDepthFar() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFCamera.Bind_get_depth_far, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDepthFar(zdepth_far gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, zdepth_far)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFCamera.Bind_set_depth_far, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDepthNear() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFCamera.Bind_get_depth_near, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDepthNear(zdepth_near gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, zdepth_near)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFCamera.Bind_set_depth_near, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsGLTFCamera() Expert { return self[0].AsGLTFCamera() }


//go:nosplit
func (self Simple) AsGLTFCamera() Simple { return self[0].AsGLTFCamera() }


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
func init() {classdb.Register("GLTFCamera", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
