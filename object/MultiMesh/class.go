package MultiMesh

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
MultiMesh provides low-level mesh instancing. Drawing thousands of [MeshInstance3D] nodes can be slow, since each object is submitted to the GPU then drawn individually.
MultiMesh is much faster as it can draw thousands of instances with a single draw call, resulting in less API overhead.
As a drawback, if the instances are too far away from each other, performance may be reduced as every single instance will always render (they are spatially indexed as one, for the whole object).
Since instances may have any behavior, the AABB used for visibility must be provided by the user.
[b]Note:[/b] A MultiMesh is a single object, therefore the same maximum lights per object restriction applies. This means, that once the maximum lights are consumed by one or more instances, the rest of the MultiMesh instances will [b]not[/b] receive any lighting.
[b]Note:[/b] Blend Shapes will be ignored if used in a MultiMesh.

*/
type Simple [1]classdb.MultiMesh
func (self Simple) SetMesh(mesh [1]classdb.Mesh) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMesh(mesh)
}
func (self Simple) GetMesh() [1]classdb.Mesh {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Mesh(Expert(self).GetMesh(gc))
}
func (self Simple) SetUseColors(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseColors(enable)
}
func (self Simple) IsUsingColors() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsUsingColors())
}
func (self Simple) SetUseCustomData(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseCustomData(enable)
}
func (self Simple) IsUsingCustomData() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsUsingCustomData())
}
func (self Simple) SetTransformFormat(format classdb.MultiMeshTransformFormat) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTransformFormat(format)
}
func (self Simple) GetTransformFormat() classdb.MultiMeshTransformFormat {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.MultiMeshTransformFormat(Expert(self).GetTransformFormat())
}
func (self Simple) SetInstanceCount(count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInstanceCount(gd.Int(count))
}
func (self Simple) GetInstanceCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetInstanceCount()))
}
func (self Simple) SetVisibleInstanceCount(count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVisibleInstanceCount(gd.Int(count))
}
func (self Simple) GetVisibleInstanceCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetVisibleInstanceCount()))
}
func (self Simple) SetInstanceTransform(instance int, transform gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInstanceTransform(gd.Int(instance), transform)
}
func (self Simple) SetInstanceTransform2d(instance int, transform gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInstanceTransform2d(gd.Int(instance), transform)
}
func (self Simple) GetInstanceTransform(instance int) gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).GetInstanceTransform(gd.Int(instance)))
}
func (self Simple) GetInstanceTransform2d(instance int) gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(Expert(self).GetInstanceTransform2d(gd.Int(instance)))
}
func (self Simple) SetInstanceColor(instance int, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInstanceColor(gd.Int(instance), color)
}
func (self Simple) GetInstanceColor(instance int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetInstanceColor(gd.Int(instance)))
}
func (self Simple) SetInstanceCustomData(instance int, custom_data gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInstanceCustomData(gd.Int(instance), custom_data)
}
func (self Simple) GetInstanceCustomData(instance int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetInstanceCustomData(gd.Int(instance)))
}
func (self Simple) SetCustomAabb(aabb gd.AABB) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCustomAabb(aabb)
}
func (self Simple) GetCustomAabb() gd.AABB {
	gc := gd.GarbageCollector(); _ = gc
	return gd.AABB(Expert(self).GetCustomAabb())
}
func (self Simple) GetAabb() gd.AABB {
	gc := gd.GarbageCollector(); _ = gc
	return gd.AABB(Expert(self).GetAabb())
}
func (self Simple) GetBuffer() gd.PackedFloat32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedFloat32Array(Expert(self).GetBuffer(gc))
}
func (self Simple) SetBuffer(buffer gd.PackedFloat32Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBuffer(buffer)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.MultiMesh
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetMesh(mesh object.Mesh)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(mesh[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_set_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMesh(ctx gd.Lifetime) object.Mesh {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_get_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Mesh
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseColors(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_set_use_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingColors() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_is_using_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseCustomData(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_set_use_custom_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingCustomData() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_is_using_custom_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTransformFormat(format classdb.MultiMeshTransformFormat)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, format)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_set_transform_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTransformFormat() classdb.MultiMeshTransformFormat {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.MultiMeshTransformFormat](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_get_transform_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInstanceCount(count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_set_instance_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInstanceCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_get_instance_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisibleInstanceCount(count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_set_visible_instance_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVisibleInstanceCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_get_visible_instance_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [Transform3D] for a specific instance.
*/
//go:nosplit
func (self class) SetInstanceTransform(instance gd.Int, transform gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_set_instance_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [Transform2D] for a specific instance.
*/
//go:nosplit
func (self class) SetInstanceTransform2d(instance gd.Int, transform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_set_instance_transform_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Transform3D] of a specific instance.
*/
//go:nosplit
func (self class) GetInstanceTransform(instance gd.Int) gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_get_instance_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Transform2D] of a specific instance.
*/
//go:nosplit
func (self class) GetInstanceTransform2d(instance gd.Int) gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_get_instance_transform_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the color of a specific instance by [i]multiplying[/i] the mesh's existing vertex colors. This allows for different color tinting per instance.
[b]Note:[/b] Each component is stored in 32 bits in the Forward+ and Mobile rendering methods, but is packed into 16 bits in the Compatibility rendering method.
For the color to take effect, ensure that [member use_colors] is [code]true[/code] on the [MultiMesh] and [member BaseMaterial3D.vertex_color_use_as_albedo] is [code]true[/code] on the material. If you intend to set an absolute color instead of tinting, make sure the material's albedo color is set to pure white ([code]Color(1, 1, 1)[/code]).
*/
//go:nosplit
func (self class) SetInstanceColor(instance gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_set_instance_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets a specific instance's color multiplier.
*/
//go:nosplit
func (self class) GetInstanceColor(instance gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_get_instance_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets custom data for a specific instance. [param custom_data] is a [Color] type only to contain 4 floating-point numbers.
[b]Note:[/b] Each number is stored in 32 bits in the Forward+ and Mobile rendering methods, but is packed into 16 bits in the Compatibility rendering method.
For the custom data to be used, ensure that [member use_custom_data] is [code]true[/code].
This custom instance data has to be manually accessed in your custom shader using [code]INSTANCE_CUSTOM[/code].
*/
//go:nosplit
func (self class) SetInstanceCustomData(instance gd.Int, custom_data gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, custom_data)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_set_instance_custom_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the custom data that has been set for a specific instance.
*/
//go:nosplit
func (self class) GetInstanceCustomData(instance gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_get_instance_custom_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCustomAabb(aabb gd.AABB)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, aabb)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_set_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCustomAabb() gd.AABB {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_get_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the visibility axis-aligned bounding box in local space.
*/
//go:nosplit
func (self class) GetAabb() gd.AABB {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_get_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetBuffer(ctx gd.Lifetime) gd.PackedFloat32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_get_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedFloat32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBuffer(buffer gd.PackedFloat32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(buffer))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiMesh.Bind_set_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsMultiMesh() Expert { return self[0].AsMultiMesh() }


//go:nosplit
func (self Simple) AsMultiMesh() Simple { return self[0].AsMultiMesh() }


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

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("MultiMesh", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type TransformFormat = classdb.MultiMeshTransformFormat

const (
/*Use this when using 2D transforms.*/
	Transform2d TransformFormat = 0
/*Use this when using 3D transforms.*/
	Transform3d TransformFormat = 1
)
