package PrimitiveMesh

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Mesh"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Base class for all primitive meshes. Handles applying a [Material] to a primitive mesh. Examples include [BoxMesh], [CapsuleMesh], [CylinderMesh], [PlaneMesh], [PrismMesh], and [SphereMesh].
	// PrimitiveMesh methods that can be overridden by a [Class] that extends it.
	type PrimitiveMesh interface {
		//Override this method to customize how this primitive mesh should be generated. Should return an [Array] where each element is another Array of values required for the mesh (see the [enum Mesh.ArrayType] constants).
		CreateMeshArray() gd.Array
	}

*/
type Go [1]classdb.PrimitiveMesh

/*
Override this method to customize how this primitive mesh should be generated. Should return an [Array] where each element is another Array of values required for the mesh (see the [enum Mesh.ArrayType] constants).
*/
func (Go) _create_mesh_array(impl func(ptr unsafe.Pointer) gd.Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
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
func (self Go) GetMeshArrays() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(class(self).GetMeshArrays(gc))
}

/*
Request an update of this primitive mesh based on its properties.
*/
func (self Go) RequestUpdate() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RequestUpdate()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PrimitiveMesh
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("PrimitiveMesh"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Material() gdclass.Material {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Material(class(self).GetMaterial(gc))
}

func (self Go) SetMaterial(value gdclass.Material) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaterial(value)
}

func (self Go) CustomAabb() gd.AABB {
	gc := gd.GarbageCollector(); _ = gc
		return gd.AABB(class(self).GetCustomAabb())
}

func (self Go) SetCustomAabb(value gd.AABB) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCustomAabb(value)
}

func (self Go) FlipFaces() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetFlipFaces())
}

func (self Go) SetFlipFaces(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFlipFaces(value)
}

func (self Go) AddUv2() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetAddUv2())
}

func (self Go) SetAddUv2(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAddUv2(value)
}

func (self Go) Uv2Padding() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetUv2Padding()))
}

func (self Go) SetUv2Padding(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUv2Padding(gd.Float(value))
}

/*
Override this method to customize how this primitive mesh should be generated. Should return an [Array] where each element is another Array of values required for the mesh (see the [enum Mesh.ArrayType] constants).
*/
func (class) _create_mesh_array(impl func(ptr unsafe.Pointer) gd.Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

//go:nosplit
func (self class) SetMaterial(material gdclass.Material)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(material[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrimitiveMesh.Bind_set_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaterial(ctx gd.Lifetime) gdclass.Material {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrimitiveMesh.Bind_get_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Material
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
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
func (self class) GetMeshArrays(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrimitiveMesh.Bind_get_mesh_arrays, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCustomAabb(aabb gd.AABB)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, aabb)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrimitiveMesh.Bind_set_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCustomAabb() gd.AABB {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrimitiveMesh.Bind_get_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFlipFaces(flip_faces bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flip_faces)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrimitiveMesh.Bind_set_flip_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFlipFaces() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrimitiveMesh.Bind_get_flip_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAddUv2(add_uv2 bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, add_uv2)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrimitiveMesh.Bind_set_add_uv2, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAddUv2() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrimitiveMesh.Bind_get_add_uv2, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUv2Padding(uv2_padding gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, uv2_padding)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrimitiveMesh.Bind_set_uv2_padding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUv2Padding() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrimitiveMesh.Bind_get_uv2_padding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Request an update of this primitive mesh based on its properties.
*/
//go:nosplit
func (self class) RequestUpdate()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrimitiveMesh.Bind_request_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsPrimitiveMesh() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPrimitiveMesh() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsMesh() Mesh.GD { return *((*Mesh.GD)(unsafe.Pointer(&self))) }
func (self Go) AsMesh() Mesh.Go { return *((*Mesh.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_create_mesh_array": return reflect.ValueOf(self._create_mesh_array);
	default: return gd.VirtualByName(self.AsMesh(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_create_mesh_array": return reflect.ValueOf(self._create_mesh_array);
	default: return gd.VirtualByName(self.AsMesh(), name)
	}
}
func init() {classdb.Register("PrimitiveMesh", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
