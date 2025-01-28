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
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Packed"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Transform2D"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/AABB"

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
	class(self).SetInstanceTransform(gd.Int(instance), gd.Transform3D(transform))
}

/*
Sets the [Transform2D] for a specific instance.
*/
func (self Instance) SetInstanceTransform2d(instance int, transform Transform2D.OriginXY) { //gd:MultiMesh.set_instance_transform_2d
	class(self).SetInstanceTransform2d(gd.Int(instance), gd.Transform2D(transform))
}

/*
Returns the [Transform3D] of a specific instance.
*/
func (self Instance) GetInstanceTransform(instance int) Transform3D.BasisOrigin { //gd:MultiMesh.get_instance_transform
	return Transform3D.BasisOrigin(class(self).GetInstanceTransform(gd.Int(instance)))
}

/*
Returns the [Transform2D] of a specific instance.
*/
func (self Instance) GetInstanceTransform2d(instance int) Transform2D.OriginXY { //gd:MultiMesh.get_instance_transform_2d
	return Transform2D.OriginXY(class(self).GetInstanceTransform2d(gd.Int(instance)))
}

/*
Sets the color of a specific instance by [i]multiplying[/i] the mesh's existing vertex colors. This allows for different color tinting per instance.
[b]Note:[/b] Each component is stored in 32 bits in the Forward+ and Mobile rendering methods, but is packed into 16 bits in the Compatibility rendering method.
For the color to take effect, ensure that [member use_colors] is [code]true[/code] on the [MultiMesh] and [member BaseMaterial3D.vertex_color_use_as_albedo] is [code]true[/code] on the material. If you intend to set an absolute color instead of tinting, make sure the material's albedo color is set to pure white ([code]Color(1, 1, 1)[/code]).
*/
func (self Instance) SetInstanceColor(instance int, color Color.RGBA) { //gd:MultiMesh.set_instance_color
	class(self).SetInstanceColor(gd.Int(instance), gd.Color(color))
}

/*
Gets a specific instance's color multiplier.
*/
func (self Instance) GetInstanceColor(instance int) Color.RGBA { //gd:MultiMesh.get_instance_color
	return Color.RGBA(class(self).GetInstanceColor(gd.Int(instance)))
}

/*
Sets custom data for a specific instance. [param custom_data] is a [Color] type only to contain 4 floating-point numbers.
[b]Note:[/b] Each number is stored in 32 bits in the Forward+ and Mobile rendering methods, but is packed into 16 bits in the Compatibility rendering method.
For the custom data to be used, ensure that [member use_custom_data] is [code]true[/code].
This custom instance data has to be manually accessed in your custom shader using [code]INSTANCE_CUSTOM[/code].
*/
func (self Instance) SetInstanceCustomData(instance int, custom_data Color.RGBA) { //gd:MultiMesh.set_instance_custom_data
	class(self).SetInstanceCustomData(gd.Int(instance), gd.Color(custom_data))
}

/*
Returns the custom data that has been set for a specific instance.
*/
func (self Instance) GetInstanceCustomData(instance int) Color.RGBA { //gd:MultiMesh.get_instance_custom_data
	return Color.RGBA(class(self).GetInstanceCustomData(gd.Int(instance)))
}

/*
Returns the visibility axis-aligned bounding box in local space.
*/
func (self Instance) GetAabb() AABB.PositionSize { //gd:MultiMesh.get_aabb
	return AABB.PositionSize(class(self).GetAabb())
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
	class(self).SetCustomAabb(gd.AABB(value))
}

func (self Instance) InstanceCount() int {
	return int(int(class(self).GetInstanceCount()))
}

func (self Instance) SetInstanceCount(value int) {
	class(self).SetInstanceCount(gd.Int(value))
}

func (self Instance) VisibleInstanceCount() int {
	return int(int(class(self).GetVisibleInstanceCount()))
}

func (self Instance) SetVisibleInstanceCount(value int) {
	class(self).SetVisibleInstanceCount(gd.Int(value))
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
func (self class) SetInstanceCount(count gd.Int) { //gd:MultiMesh.set_instance_count
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_instance_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInstanceCount() gd.Int { //gd:MultiMesh.get_instance_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_instance_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibleInstanceCount(count gd.Int) { //gd:MultiMesh.set_visible_instance_count
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_visible_instance_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibleInstanceCount() gd.Int { //gd:MultiMesh.get_visible_instance_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_visible_instance_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [Transform3D] for a specific instance.
*/
//go:nosplit
func (self class) SetInstanceTransform(instance gd.Int, transform gd.Transform3D) { //gd:MultiMesh.set_instance_transform
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, transform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_instance_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the [Transform2D] for a specific instance.
*/
//go:nosplit
func (self class) SetInstanceTransform2d(instance gd.Int, transform gd.Transform2D) { //gd:MultiMesh.set_instance_transform_2d
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
func (self class) GetInstanceTransform(instance gd.Int) gd.Transform3D { //gd:MultiMesh.get_instance_transform
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_instance_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Transform2D] of a specific instance.
*/
//go:nosplit
func (self class) GetInstanceTransform2d(instance gd.Int) gd.Transform2D { //gd:MultiMesh.get_instance_transform_2d
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	var r_ret = callframe.Ret[gd.Transform2D](frame)
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
func (self class) SetInstanceColor(instance gd.Int, color gd.Color) { //gd:MultiMesh.set_instance_color
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
func (self class) GetInstanceColor(instance gd.Int) gd.Color { //gd:MultiMesh.get_instance_color
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	var r_ret = callframe.Ret[gd.Color](frame)
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
func (self class) SetInstanceCustomData(instance gd.Int, custom_data gd.Color) { //gd:MultiMesh.set_instance_custom_data
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
func (self class) GetInstanceCustomData(instance gd.Int) gd.Color { //gd:MultiMesh.get_instance_custom_data
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_instance_custom_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCustomAabb(aabb gd.AABB) { //gd:MultiMesh.set_custom_aabb
	var frame = callframe.New()
	callframe.Arg(frame, aabb)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCustomAabb() gd.AABB { //gd:MultiMesh.get_custom_aabb
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the visibility axis-aligned bounding box in local space.
*/
//go:nosplit
func (self class) GetAabb() gd.AABB { //gd:MultiMesh.get_aabb
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
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
