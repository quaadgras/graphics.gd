// Package PrimitiveMesh provides methods for working with PrimitiveMesh object instances.
package PrimitiveMesh

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
import "graphics.gd/classdb/Mesh"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/AABB"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
Base class for all primitive meshes. Handles applying a [Material] to a primitive mesh. Examples include [BoxMesh], [CapsuleMesh], [CylinderMesh], [PlaneMesh], [PrismMesh], and [SphereMesh].

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=PrimitiveMesh)
*/
type Instance [1]gdclass.PrimitiveMesh

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPrimitiveMesh() Instance
}
type Interface interface {
	//Override this method to customize how this primitive mesh should be generated. Should return an [Array] where each element is another Array of values required for the mesh (see the [enum Mesh.ArrayType] constants).
	CreateMeshArray() []any
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) CreateMeshArray() (_ []any) { return }

/*
Override this method to customize how this primitive mesh should be generated. Should return an [Array] where each element is another Array of values required for the mesh (see the [enum Mesh.ArrayType] constants).
*/
func (Instance) _create_mesh_array(impl func(ptr unsafe.Pointer) []any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalArray(gd.EngineArrayFromSlice(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns mesh arrays used to constitute surface of [Mesh]. The result can be passed to [method ArrayMesh.add_surface_from_arrays] to create a new surface. For example:
[codeblocks]
[gdscript]
var c = CylinderMesh.new()
var arr_mesh = ArrayMesh.new()
arr_mesh.add_surface_from_arrays(Mesh.PRIMITIVE_TRIANGLES, c.get_mesh_arrays())
[/gdscript]
[csharp]
var c = new CylinderMesh();
var arrMesh = new ArrayMesh();
arrMesh.AddSurfaceFromArrays(Mesh.PrimitiveType.Triangles, c.GetMeshArrays());
[/csharp]
[/codeblocks]
*/
func (self Instance) GetMeshArrays() []any {
	return []any(gd.ArrayAs[[]any](gd.InternalArray(class(self).GetMeshArrays())))
}

/*
Request an update of this primitive mesh based on its properties.
*/
func (self Instance) RequestUpdate() {
	class(self).RequestUpdate()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PrimitiveMesh

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PrimitiveMesh"))
	casted := Instance{*(*gdclass.PrimitiveMesh)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Material() [1]gdclass.Material {
	return [1]gdclass.Material(class(self).GetMaterial())
}

func (self Instance) SetMaterial(value [1]gdclass.Material) {
	class(self).SetMaterial(value)
}

func (self Instance) CustomAabb() AABB.PositionSize {
	return AABB.PositionSize(class(self).GetCustomAabb())
}

func (self Instance) SetCustomAabb(value AABB.PositionSize) {
	class(self).SetCustomAabb(gd.AABB(value))
}

func (self Instance) FlipFaces() bool {
	return bool(class(self).GetFlipFaces())
}

func (self Instance) SetFlipFaces(value bool) {
	class(self).SetFlipFaces(value)
}

func (self Instance) AddUv2() bool {
	return bool(class(self).GetAddUv2())
}

func (self Instance) SetAddUv2(value bool) {
	class(self).SetAddUv2(value)
}

func (self Instance) Uv2Padding() Float.X {
	return Float.X(Float.X(class(self).GetUv2Padding()))
}

func (self Instance) SetUv2Padding(value Float.X) {
	class(self).SetUv2Padding(gd.Float(value))
}

/*
Override this method to customize how this primitive mesh should be generated. Should return an [Array] where each element is another Array of values required for the mesh (see the [enum Mesh.ArrayType] constants).
*/
func (class) _create_mesh_array(impl func(ptr unsafe.Pointer) Array.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

//go:nosplit
func (self class) SetMaterial(material [1]gdclass.Material) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrimitiveMesh.Bind_set_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaterial() [1]gdclass.Material {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrimitiveMesh.Bind_get_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Material{gd.PointerWithOwnershipTransferredToGo[gdclass.Material](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns mesh arrays used to constitute surface of [Mesh]. The result can be passed to [method ArrayMesh.add_surface_from_arrays] to create a new surface. For example:
[codeblocks]
[gdscript]
var c = CylinderMesh.new()
var arr_mesh = ArrayMesh.new()
arr_mesh.add_surface_from_arrays(Mesh.PRIMITIVE_TRIANGLES, c.get_mesh_arrays())
[/gdscript]
[csharp]
var c = new CylinderMesh();
var arrMesh = new ArrayMesh();
arrMesh.AddSurfaceFromArrays(Mesh.PrimitiveType.Triangles, c.GetMeshArrays());
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GetMeshArrays() Array.Any {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrimitiveMesh.Bind_get_mesh_arrays, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCustomAabb(aabb gd.AABB) {
	var frame = callframe.New()
	callframe.Arg(frame, aabb)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrimitiveMesh.Bind_set_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCustomAabb() gd.AABB {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrimitiveMesh.Bind_get_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlipFaces(flip_faces bool) {
	var frame = callframe.New()
	callframe.Arg(frame, flip_faces)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrimitiveMesh.Bind_set_flip_faces, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFlipFaces() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrimitiveMesh.Bind_get_flip_faces, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAddUv2(add_uv2 bool) {
	var frame = callframe.New()
	callframe.Arg(frame, add_uv2)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrimitiveMesh.Bind_set_add_uv2, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAddUv2() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrimitiveMesh.Bind_get_add_uv2, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUv2Padding(uv2_padding gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, uv2_padding)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrimitiveMesh.Bind_set_uv2_padding, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUv2Padding() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrimitiveMesh.Bind_get_uv2_padding, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Request an update of this primitive mesh based on its properties.
*/
//go:nosplit
func (self class) RequestUpdate() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrimitiveMesh.Bind_request_update, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsPrimitiveMesh() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPrimitiveMesh() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsMesh() Mesh.Advanced        { return *((*Mesh.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMesh() Mesh.Instance     { return *((*Mesh.Instance)(unsafe.Pointer(&self))) }
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
	case "_create_mesh_array":
		return reflect.ValueOf(self._create_mesh_array)
	default:
		return gd.VirtualByName(Mesh.Advanced(self.AsMesh()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_create_mesh_array":
		return reflect.ValueOf(self._create_mesh_array)
	default:
		return gd.VirtualByName(Mesh.Instance(self.AsMesh()), name)
	}
}
func init() {
	gdclass.Register("PrimitiveMesh", func(ptr gd.Object) any {
		return [1]gdclass.PrimitiveMesh{*(*gdclass.PrimitiveMesh)(unsafe.Pointer(&ptr))}
	})
}
