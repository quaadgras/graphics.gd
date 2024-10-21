package MeshInstance3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/GeometryInstance3D"
import "grow.graphics/gd/object/VisualInstance3D"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
MeshInstance3D is a node that takes a [Mesh] resource and adds it to the current scenario by creating an instance of it. This is the class most often used render 3D geometry and can be used to instance a single [Mesh] in many places. This allows reusing geometry, which can save on resources. When a [Mesh] has to be instantiated more than thousands of times at close proximity, consider using a [MultiMesh] in a [MultiMeshInstance3D] instead.

*/
type Simple [1]classdb.MeshInstance3D
func (self Simple) SetMesh(mesh [1]classdb.Mesh) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMesh(mesh)
}
func (self Simple) GetMesh() [1]classdb.Mesh {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Mesh(Expert(self).GetMesh(gc))
}
func (self Simple) SetSkeletonPath(skeleton_path gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSkeletonPath(skeleton_path)
}
func (self Simple) GetSkeletonPath() gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetSkeletonPath(gc))
}
func (self Simple) SetSkin(skin [1]classdb.Skin) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSkin(skin)
}
func (self Simple) GetSkin() [1]classdb.Skin {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Skin(Expert(self).GetSkin(gc))
}
func (self Simple) GetSkinReference() [1]classdb.SkinReference {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.SkinReference(Expert(self).GetSkinReference(gc))
}
func (self Simple) GetSurfaceOverrideMaterialCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSurfaceOverrideMaterialCount()))
}
func (self Simple) SetSurfaceOverrideMaterial(surface int, material [1]classdb.Material) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSurfaceOverrideMaterial(gd.Int(surface), material)
}
func (self Simple) GetSurfaceOverrideMaterial(surface int) [1]classdb.Material {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Material(Expert(self).GetSurfaceOverrideMaterial(gc, gd.Int(surface)))
}
func (self Simple) GetActiveMaterial(surface int) [1]classdb.Material {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Material(Expert(self).GetActiveMaterial(gc, gd.Int(surface)))
}
func (self Simple) CreateTrimeshCollision() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CreateTrimeshCollision()
}
func (self Simple) CreateConvexCollision(clean bool, simplify bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CreateConvexCollision(clean, simplify)
}
func (self Simple) CreateMultipleConvexCollisions(settings [1]classdb.MeshConvexDecompositionSettings) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CreateMultipleConvexCollisions(settings)
}
func (self Simple) GetBlendShapeCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBlendShapeCount()))
}
func (self Simple) FindBlendShapeByName(name string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).FindBlendShapeByName(gc.StringName(name))))
}
func (self Simple) GetBlendShapeValue(blend_shape_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetBlendShapeValue(gd.Int(blend_shape_idx))))
}
func (self Simple) SetBlendShapeValue(blend_shape_idx int, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBlendShapeValue(gd.Int(blend_shape_idx), gd.Float(value))
}
func (self Simple) CreateDebugTangents() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CreateDebugTangents()
}
func (self Simple) BakeMeshFromCurrentBlendShapeMix(existing [1]classdb.ArrayMesh) [1]classdb.ArrayMesh {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.ArrayMesh(Expert(self).BakeMeshFromCurrentBlendShapeMix(gc, existing))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.MeshInstance3D
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshInstance3D.Bind_set_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMesh(ctx gd.Lifetime) object.Mesh {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshInstance3D.Bind_get_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Mesh
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkeletonPath(skeleton_path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(skeleton_path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshInstance3D.Bind_set_skeleton_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkeletonPath(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshInstance3D.Bind_get_skeleton_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkin(skin object.Skin)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(skin[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshInstance3D.Bind_set_skin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkin(ctx gd.Lifetime) object.Skin {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshInstance3D.Bind_get_skin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Skin
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the internal [SkinReference] containing the skeleton's [RID] attached to this RID. See also [method Resource.get_rid], [method SkinReference.get_skeleton], and [method RenderingServer.instance_attach_skeleton].
*/
//go:nosplit
func (self class) GetSkinReference(ctx gd.Lifetime) object.SkinReference {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshInstance3D.Bind_get_skin_reference, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.SkinReference
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the number of surface override materials. This is equivalent to [method Mesh.get_surface_count]. See also [method get_surface_override_material].
*/
//go:nosplit
func (self class) GetSurfaceOverrideMaterialCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshInstance3D.Bind_get_surface_override_material_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the override [param material] for the specified [param surface] of the [Mesh] resource. This material is associated with this [MeshInstance3D] rather than with [member mesh].
[b]Note:[/b] This assigns the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties, not the material within the [Mesh] resource. To set the material within the [Mesh] resource, use [method Mesh.surface_get_material] instead.
*/
//go:nosplit
func (self class) SetSurfaceOverrideMaterial(surface gd.Int, material object.Material)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, surface)
	callframe.Arg(frame, mmm.Get(material[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshInstance3D.Bind_set_surface_override_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the override [Material] for the specified [param surface] of the [Mesh] resource. See also [method get_surface_override_material_count].
[b]Note:[/b] This returns the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties, not the material within the [Mesh] resource. To get the material within the [Mesh] resource, use [method Mesh.surface_get_material] instead.
*/
//go:nosplit
func (self class) GetSurfaceOverrideMaterial(ctx gd.Lifetime, surface gd.Int) object.Material {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, surface)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshInstance3D.Bind_get_surface_override_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Material
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the [Material] that will be used by the [Mesh] when drawing. This can return the [member GeometryInstance3D.material_override], the surface override [Material] defined in this [MeshInstance3D], or the surface [Material] defined in the [member mesh]. For example, if [member GeometryInstance3D.material_override] is used, all surfaces will return the override material.
Returns [code]null[/code] if no material is active, including when [member mesh] is [code]null[/code].
*/
//go:nosplit
func (self class) GetActiveMaterial(ctx gd.Lifetime, surface gd.Int) object.Material {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, surface)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshInstance3D.Bind_get_active_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Material
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
This helper creates a [StaticBody3D] child node with a [ConcavePolygonShape3D] collision shape calculated from the mesh geometry. It's mainly used for testing.
*/
//go:nosplit
func (self class) CreateTrimeshCollision()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshInstance3D.Bind_create_trimesh_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
This helper creates a [StaticBody3D] child node with a [ConvexPolygonShape3D] collision shape calculated from the mesh geometry. It's mainly used for testing.
If [param clean] is [code]true[/code] (default), duplicate and interior vertices are removed automatically. You can set it to [code]false[/code] to make the process faster if not needed.
If [param simplify] is [code]true[/code], the geometry can be further simplified to reduce the number of vertices. Disabled by default.
*/
//go:nosplit
func (self class) CreateConvexCollision(clean bool, simplify bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, clean)
	callframe.Arg(frame, simplify)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshInstance3D.Bind_create_convex_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
This helper creates a [StaticBody3D] child node with multiple [ConvexPolygonShape3D] collision shapes calculated from the mesh geometry via convex decomposition. The convex decomposition operation can be controlled with parameters from the optional [param settings].
*/
//go:nosplit
func (self class) CreateMultipleConvexCollisions(settings object.MeshConvexDecompositionSettings)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(settings[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshInstance3D.Bind_create_multiple_convex_collisions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of blend shapes available. Produces an error if [member mesh] is [code]null[/code].
*/
//go:nosplit
func (self class) GetBlendShapeCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshInstance3D.Bind_get_blend_shape_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the index of the blend shape with the given [param name]. Returns [code]-1[/code] if no blend shape with this name exists, including when [member mesh] is [code]null[/code].
*/
//go:nosplit
func (self class) FindBlendShapeByName(name gd.StringName) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshInstance3D.Bind_find_blend_shape_by_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the value of the blend shape at the given [param blend_shape_idx]. Returns [code]0.0[/code] and produces an error if [member mesh] is [code]null[/code] or doesn't have a blend shape at that index.
*/
//go:nosplit
func (self class) GetBlendShapeValue(blend_shape_idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, blend_shape_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshInstance3D.Bind_get_blend_shape_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the value of the blend shape at [param blend_shape_idx] to [param value]. Produces an error if [member mesh] is [code]null[/code] or doesn't have a blend shape at that index.
*/
//go:nosplit
func (self class) SetBlendShapeValue(blend_shape_idx gd.Int, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, blend_shape_idx)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshInstance3D.Bind_set_blend_shape_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
This helper creates a [MeshInstance3D] child node with gizmos at every vertex calculated from the mesh geometry. It's mainly used for testing.
*/
//go:nosplit
func (self class) CreateDebugTangents()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshInstance3D.Bind_create_debug_tangents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Takes a snapshot from the current [ArrayMesh] with all blend shapes applied according to their current weights and bakes it to the provided [param existing] mesh. If no [param existing] mesh is provided a new [ArrayMesh] is created, baked and returned. Mesh surface materials are not copied.
[b]Performance:[/b] [Mesh] data needs to be received from the GPU, stalling the [RenderingServer] in the process.
*/
//go:nosplit
func (self class) BakeMeshFromCurrentBlendShapeMix(ctx gd.Lifetime, existing object.ArrayMesh) object.ArrayMesh {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(existing[0].AsPointer())[0])
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshInstance3D.Bind_bake_mesh_from_current_blend_shape_mix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.ArrayMesh
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsMeshInstance3D() Expert { return self[0].AsMeshInstance3D() }


//go:nosplit
func (self Simple) AsMeshInstance3D() Simple { return self[0].AsMeshInstance3D() }


//go:nosplit
func (self class) AsGeometryInstance3D() GeometryInstance3D.Expert { return self[0].AsGeometryInstance3D() }


//go:nosplit
func (self Simple) AsGeometryInstance3D() GeometryInstance3D.Simple { return self[0].AsGeometryInstance3D() }


//go:nosplit
func (self class) AsVisualInstance3D() VisualInstance3D.Expert { return self[0].AsVisualInstance3D() }


//go:nosplit
func (self Simple) AsVisualInstance3D() VisualInstance3D.Simple { return self[0].AsVisualInstance3D() }


//go:nosplit
func (self class) AsNode3D() Node3D.Expert { return self[0].AsNode3D() }


//go:nosplit
func (self Simple) AsNode3D() Node3D.Simple { return self[0].AsNode3D() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

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
func init() {classdb.Register("MeshInstance3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
