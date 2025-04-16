// Package MultiMesh provides methods for working with MultiMesh object instances.
package MultiMesh

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/AABB"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Transform2D"
import "graphics.gd/variant/Transform3D"

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
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
MultiMesh provides low-level mesh instancing. Drawing thousands of [MeshInstance3D] nodes can be slow, since each object is submitted to the GPU then drawn individually.
MultiMesh is much faster as it can draw thousands of instances with a single draw call, resulting in less API overhead.
As a drawback, if the instances are too far away from each other, performance may be reduced as every single instance will always render (they are spatially indexed as one, for the whole object).
Since instances may have any behavior, the AABB used for visibility must be provided by the user.
[b]Note:[/b] A MultiMesh is a single object, therefore the same maximum lights per object restriction applies. This means, that once the maximum lights are consumed by one or more instances, the rest of the MultiMesh instances will [b]not[/b] receive any lighting.
[b]Note:[/b] Blend Shapes will be ignored if used in a MultiMesh.
*/
type Instance [1]gdclass.MultiMesh

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsMultiMesh() Instance
}

/*
Sets the [Transform3D] for a specific instance.
*/
func (self Instance) SetInstanceTransform(instance int, transform Transform3D.BasisOrigin) { //gd:MultiMesh.set_instance_transform
	Advanced(self).SetInstanceTransform(int64(instance), Transform3D.BasisOrigin(transform))
}

/*
Sets the [Transform2D] for a specific instance.
*/
func (self Instance) SetInstanceTransform2d(instance int, transform Transform2D.OriginXY) { //gd:MultiMesh.set_instance_transform_2d
	Advanced(self).SetInstanceTransform2d(int64(instance), Transform2D.OriginXY(transform))
}

/*
Returns the [Transform3D] of a specific instance.
*/
func (self Instance) GetInstanceTransform(instance int) Transform3D.BasisOrigin { //gd:MultiMesh.get_instance_transform
	return Transform3D.BasisOrigin(Advanced(self).GetInstanceTransform(int64(instance)))
}

/*
Returns the [Transform2D] of a specific instance.
*/
func (self Instance) GetInstanceTransform2d(instance int) Transform2D.OriginXY { //gd:MultiMesh.get_instance_transform_2d
	return Transform2D.OriginXY(Advanced(self).GetInstanceTransform2d(int64(instance)))
}

/*
Sets the color of a specific instance by [i]multiplying[/i] the mesh's existing vertex colors. This allows for different color tinting per instance.
[b]Note:[/b] Each component is stored in 32 bits in the Forward+ and Mobile rendering methods, but is packed into 16 bits in the Compatibility rendering method.
For the color to take effect, ensure that [member use_colors] is [code]true[/code] on the [MultiMesh] and [member BaseMaterial3D.vertex_color_use_as_albedo] is [code]true[/code] on the material. If you intend to set an absolute color instead of tinting, make sure the material's albedo color is set to pure white ([code]Color(1, 1, 1)[/code]).
*/
func (self Instance) SetInstanceColor(instance int, color Color.RGBA) { //gd:MultiMesh.set_instance_color
	Advanced(self).SetInstanceColor(int64(instance), Color.RGBA(color))
}

/*
Gets a specific instance's color multiplier.
*/
func (self Instance) GetInstanceColor(instance int) Color.RGBA { //gd:MultiMesh.get_instance_color
	return Color.RGBA(Advanced(self).GetInstanceColor(int64(instance)))
}

/*
Sets custom data for a specific instance. [param custom_data] is a [Color] type only to contain 4 floating-point numbers.
[b]Note:[/b] Each number is stored in 32 bits in the Forward+ and Mobile rendering methods, but is packed into 16 bits in the Compatibility rendering method.
For the custom data to be used, ensure that [member use_custom_data] is [code]true[/code].
This custom instance data has to be manually accessed in your custom shader using [code]INSTANCE_CUSTOM[/code].
*/
func (self Instance) SetInstanceCustomData(instance int, custom_data Color.RGBA) { //gd:MultiMesh.set_instance_custom_data
	Advanced(self).SetInstanceCustomData(int64(instance), Color.RGBA(custom_data))
}

/*
Returns the custom data that has been set for a specific instance.
*/
func (self Instance) GetInstanceCustomData(instance int) Color.RGBA { //gd:MultiMesh.get_instance_custom_data
	return Color.RGBA(Advanced(self).GetInstanceCustomData(int64(instance)))
}

/*
When using [i]physics interpolation[/i], this function allows you to prevent interpolation on an instance in the current physics tick.
This allows you to move instances instantaneously, and should usually be used when initially placing an instance such as a bullet to prevent graphical glitches.
*/
func (self Instance) ResetInstancePhysicsInterpolation(instance int) { //gd:MultiMesh.reset_instance_physics_interpolation
	Advanced(self).ResetInstancePhysicsInterpolation(int64(instance))
}

/*
Returns the visibility axis-aligned bounding box in local space.
*/
func (self Instance) GetAabb() AABB.PositionSize { //gd:MultiMesh.get_aabb
	return AABB.PositionSize(Advanced(self).GetAabb())
}

/*
An alternative to setting the [member buffer] property, which can be used with [i]physics interpolation[/i]. This method takes two arrays, and can set the data for the current and previous tick in one go. The renderer will automatically interpolate the data at each frame.
This is useful for situations where the order of instances may change from physics tick to tick, such as particle systems.
When the order of instances is coherent, the simpler alternative of setting [member buffer] can still be used with interpolation.
*/
func (self Instance) SetBufferInterpolated(buffer_curr []float32, buffer_prev []float32) { //gd:MultiMesh.set_buffer_interpolated
	Advanced(self).SetBufferInterpolated(Packed.New(buffer_curr...), Packed.New(buffer_prev...))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.MultiMesh

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MultiMesh"))
	casted := Instance{*(*gdclass.MultiMesh)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) TransformFormat() gdclass.MultiMeshTransformFormat {
	return gdclass.MultiMeshTransformFormat(class(self).GetTransformFormat())
}

func (self Instance) SetTransformFormat(value gdclass.MultiMeshTransformFormat) {
	class(self).SetTransformFormat(value)
}

func (self Instance) UseColors() bool {
	return bool(class(self).IsUsingColors())
}

func (self Instance) SetUseColors(value bool) {
	class(self).SetUseColors(value)
}

func (self Instance) UseCustomData() bool {
	return bool(class(self).IsUsingCustomData())
}

func (self Instance) SetUseCustomData(value bool) {
	class(self).SetUseCustomData(value)
}

func (self Instance) CustomAabb() AABB.PositionSize {
	return AABB.PositionSize(class(self).GetCustomAabb())
}

func (self Instance) SetCustomAabb(value AABB.PositionSize) {
	class(self).SetCustomAabb(AABB.PositionSize(value))
}

func (self Instance) InstanceCount() int {
	return int(int(class(self).GetInstanceCount()))
}

func (self Instance) SetInstanceCount(value int) {
	class(self).SetInstanceCount(int64(value))
}

func (self Instance) VisibleInstanceCount() int {
	return int(int(class(self).GetVisibleInstanceCount()))
}

func (self Instance) SetVisibleInstanceCount(value int) {
	class(self).SetVisibleInstanceCount(int64(value))
}

func (self Instance) Mesh() [1]gdclass.Mesh {
	return [1]gdclass.Mesh(class(self).GetMesh())
}

func (self Instance) SetMesh(value [1]gdclass.Mesh) {
	class(self).SetMesh(value)
}

func (self Instance) Buffer() []float32 {
	return []float32(slices.Collect(class(self).GetBuffer().Values()))
}

func (self Instance) SetBuffer(value []float32) {
	class(self).SetBuffer(Packed.New(value...))
}

func (self Instance) PhysicsInterpolationQuality() gdclass.MultiMeshPhysicsInterpolationQuality {
	return gdclass.MultiMeshPhysicsInterpolationQuality(class(self).GetPhysicsInterpolationQuality())
}

func (self Instance) SetPhysicsInterpolationQuality(value gdclass.MultiMeshPhysicsInterpolationQuality) {
	class(self).SetPhysicsInterpolationQuality(value)
}

//go:nosplit
func (self class) SetMesh(mesh [1]gdclass.Mesh) { //gd:MultiMesh.set_mesh
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMesh() [1]gdclass.Mesh { //gd:MultiMesh.get_mesh
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Mesh{gd.PointerWithOwnershipTransferredToGo[gdclass.Mesh](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseColors(enable bool) { //gd:MultiMesh.set_use_colors
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_use_colors, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingColors() bool { //gd:MultiMesh.is_using_colors
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_is_using_colors, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseCustomData(enable bool) { //gd:MultiMesh.set_use_custom_data
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_use_custom_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingCustomData() bool { //gd:MultiMesh.is_using_custom_data
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_is_using_custom_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTransformFormat(format gdclass.MultiMeshTransformFormat) { //gd:MultiMesh.set_transform_format
	var frame = callframe.New()
	callframe.Arg(frame, format)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_transform_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTransformFormat() gdclass.MultiMeshTransformFormat { //gd:MultiMesh.get_transform_format
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.MultiMeshTransformFormat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_transform_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInstanceCount(count int64) { //gd:MultiMesh.set_instance_count
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_instance_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInstanceCount() int64 { //gd:MultiMesh.get_instance_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_instance_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibleInstanceCount(count int64) { //gd:MultiMesh.set_visible_instance_count
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_visible_instance_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibleInstanceCount() int64 { //gd:MultiMesh.get_visible_instance_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_visible_instance_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPhysicsInterpolationQuality(quality gdclass.MultiMeshPhysicsInterpolationQuality) { //gd:MultiMesh.set_physics_interpolation_quality
	var frame = callframe.New()
	callframe.Arg(frame, quality)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_physics_interpolation_quality, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPhysicsInterpolationQuality() gdclass.MultiMeshPhysicsInterpolationQuality { //gd:MultiMesh.get_physics_interpolation_quality
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.MultiMeshPhysicsInterpolationQuality](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_physics_interpolation_quality, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [Transform3D] for a specific instance.
*/
//go:nosplit
func (self class) SetInstanceTransform(instance int64, transform Transform3D.BasisOrigin) { //gd:MultiMesh.set_instance_transform
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, gd.Transposed(transform))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_instance_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the [Transform2D] for a specific instance.
*/
//go:nosplit
func (self class) SetInstanceTransform2d(instance int64, transform Transform2D.OriginXY) { //gd:MultiMesh.set_instance_transform_2d
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, transform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_instance_transform_2d, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [Transform3D] of a specific instance.
*/
//go:nosplit
func (self class) GetInstanceTransform(instance int64) Transform3D.BasisOrigin { //gd:MultiMesh.get_instance_transform
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	var r_ret = callframe.Ret[Transform3D.BasisOrigin](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_instance_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = gd.Transposed(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the [Transform2D] of a specific instance.
*/
//go:nosplit
func (self class) GetInstanceTransform2d(instance int64) Transform2D.OriginXY { //gd:MultiMesh.get_instance_transform_2d
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	var r_ret = callframe.Ret[Transform2D.OriginXY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_instance_transform_2d, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) SetInstanceColor(instance int64, color Color.RGBA) { //gd:MultiMesh.set_instance_color
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_instance_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Gets a specific instance's color multiplier.
*/
//go:nosplit
func (self class) GetInstanceColor(instance int64) Color.RGBA { //gd:MultiMesh.get_instance_color
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_instance_color, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) SetInstanceCustomData(instance int64, custom_data Color.RGBA) { //gd:MultiMesh.set_instance_custom_data
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, custom_data)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_instance_custom_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the custom data that has been set for a specific instance.
*/
//go:nosplit
func (self class) GetInstanceCustomData(instance int64) Color.RGBA { //gd:MultiMesh.get_instance_custom_data
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_instance_custom_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
When using [i]physics interpolation[/i], this function allows you to prevent interpolation on an instance in the current physics tick.
This allows you to move instances instantaneously, and should usually be used when initially placing an instance such as a bullet to prevent graphical glitches.
*/
//go:nosplit
func (self class) ResetInstancePhysicsInterpolation(instance int64) { //gd:MultiMesh.reset_instance_physics_interpolation
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_reset_instance_physics_interpolation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetCustomAabb(aabb AABB.PositionSize) { //gd:MultiMesh.set_custom_aabb
	var frame = callframe.New()
	callframe.Arg(frame, aabb)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCustomAabb() AABB.PositionSize { //gd:MultiMesh.get_custom_aabb
	var frame = callframe.New()
	var r_ret = callframe.Ret[AABB.PositionSize](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the visibility axis-aligned bounding box in local space.
*/
//go:nosplit
func (self class) GetAabb() AABB.PositionSize { //gd:MultiMesh.get_aabb
	var frame = callframe.New()
	var r_ret = callframe.Ret[AABB.PositionSize](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_aabb, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetBuffer() Packed.Array[float32] { //gd:MultiMesh.get_buffer
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[float32](Array.Through(gd.PackedProxy[gd.PackedFloat32Array, float32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBuffer(buffer Packed.Array[float32]) { //gd:MultiMesh.set_buffer
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedFloat32Array, float32](buffer)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
An alternative to setting the [member buffer] property, which can be used with [i]physics interpolation[/i]. This method takes two arrays, and can set the data for the current and previous tick in one go. The renderer will automatically interpolate the data at each frame.
This is useful for situations where the order of instances may change from physics tick to tick, such as particle systems.
When the order of instances is coherent, the simpler alternative of setting [member buffer] can still be used with interpolation.
*/
//go:nosplit
func (self class) SetBufferInterpolated(buffer_curr Packed.Array[float32], buffer_prev Packed.Array[float32]) { //gd:MultiMesh.set_buffer_interpolated
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedFloat32Array, float32](buffer_curr)))
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedFloat32Array, float32](buffer_prev)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_buffer_interpolated, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsMultiMesh() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMultiMesh() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("MultiMesh", func(ptr gd.Object) any { return [1]gdclass.MultiMesh{*(*gdclass.MultiMesh)(unsafe.Pointer(&ptr))} })
}

type TransformFormat = gdclass.MultiMeshTransformFormat //gd:MultiMesh.TransformFormat

const (
	/*Use this when using 2D transforms.*/
	Transform2d TransformFormat = 0
	/*Use this when using 3D transforms.*/
	Transform3d TransformFormat = 1
)

type PhysicsInterpolationQuality = gdclass.MultiMeshPhysicsInterpolationQuality //gd:MultiMesh.PhysicsInterpolationQuality

const (
	/*Always interpolate using Basis lerping, which can produce warping artifacts in some situations.*/
	InterpQualityFast PhysicsInterpolationQuality = 0
	/*Attempt to interpolate using Basis slerping (spherical linear interpolation) where possible, otherwise fall back to lerping.*/
	InterpQualityHigh PhysicsInterpolationQuality = 1
)
