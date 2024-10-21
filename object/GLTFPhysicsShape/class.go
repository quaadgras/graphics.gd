package GLTFPhysicsShape

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
Represents a physics shape as defined by the [code]OMI_physics_shape[/code] or [code]OMI_collider[/code] GLTF extensions. This class is an intermediary between the GLTF data and Godot's nodes, and it's abstracted in a way that allows adding support for different GLTF physics extensions in the future.

*/
type Simple [1]classdb.GLTFPhysicsShape
func (self Simple) FromNode(shape_node [1]classdb.CollisionShape3D) [1]classdb.GLTFPhysicsShape {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.GLTFPhysicsShape(Expert(self).FromNode(gc, shape_node))
}
func (self Simple) ToNode(cache_shapes bool) [1]classdb.CollisionShape3D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.CollisionShape3D(Expert(self).ToNode(gc, cache_shapes))
}
func (self Simple) FromResource(shape_resource [1]classdb.Shape3D) [1]classdb.GLTFPhysicsShape {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.GLTFPhysicsShape(Expert(self).FromResource(gc, shape_resource))
}
func (self Simple) ToResource(cache_shapes bool) [1]classdb.Shape3D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Shape3D(Expert(self).ToResource(gc, cache_shapes))
}
func (self Simple) FromDictionary(dictionary gd.Dictionary) [1]classdb.GLTFPhysicsShape {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.GLTFPhysicsShape(Expert(self).FromDictionary(gc, dictionary))
}
func (self Simple) ToDictionary() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).ToDictionary(gc))
}
func (self Simple) GetShapeType() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetShapeType(gc).String())
}
func (self Simple) SetShapeType(shape_type string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShapeType(gc.String(shape_type))
}
func (self Simple) GetSize() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetSize())
}
func (self Simple) SetSize(size gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSize(size)
}
func (self Simple) GetRadius() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetRadius()))
}
func (self Simple) SetRadius(radius float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRadius(gd.Float(radius))
}
func (self Simple) GetHeight() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetHeight()))
}
func (self Simple) SetHeight(height float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHeight(gd.Float(height))
}
func (self Simple) GetIsTrigger() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetIsTrigger())
}
func (self Simple) SetIsTrigger(is_trigger bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetIsTrigger(is_trigger)
}
func (self Simple) GetMeshIndex() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMeshIndex()))
}
func (self Simple) SetMeshIndex(mesh_index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMeshIndex(gd.Int(mesh_index))
}
func (self Simple) GetImporterMesh() [1]classdb.ImporterMesh {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.ImporterMesh(Expert(self).GetImporterMesh(gc))
}
func (self Simple) SetImporterMesh(importer_mesh [1]classdb.ImporterMesh) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetImporterMesh(importer_mesh)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.GLTFPhysicsShape
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Creates a new GLTFPhysicsShape instance from the given Godot [CollisionShape3D] node.
*/
//go:nosplit
func (self class) FromNode(ctx gd.Lifetime, shape_node object.CollisionShape3D) object.GLTFPhysicsShape {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(shape_node[0].AsPointer())[0])
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.GLTFPhysicsShape.Bind_from_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.GLTFPhysicsShape
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Converts this GLTFPhysicsShape instance into a Godot [CollisionShape3D] node.
*/
//go:nosplit
func (self class) ToNode(ctx gd.Lifetime, cache_shapes bool) object.CollisionShape3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_shapes)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsShape.Bind_to_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.CollisionShape3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates a new GLTFPhysicsShape instance from the given Godot [Shape3D] resource.
*/
//go:nosplit
func (self class) FromResource(ctx gd.Lifetime, shape_resource object.Shape3D) object.GLTFPhysicsShape {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(shape_resource[0].AsPointer())[0])
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.GLTFPhysicsShape.Bind_from_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.GLTFPhysicsShape
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Converts this GLTFPhysicsShape instance into a Godot [Shape3D] resource.
*/
//go:nosplit
func (self class) ToResource(ctx gd.Lifetime, cache_shapes bool) object.Shape3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_shapes)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsShape.Bind_to_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Shape3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates a new GLTFPhysicsShape instance by parsing the given [Dictionary].
*/
//go:nosplit
func (self class) FromDictionary(ctx gd.Lifetime, dictionary gd.Dictionary) object.GLTFPhysicsShape {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(dictionary))
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.GLTFPhysicsShape.Bind_from_dictionary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.GLTFPhysicsShape
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Serializes this GLTFPhysicsShape instance into a [Dictionary] in the format defined by [code]OMI_physics_shape[/code].
*/
//go:nosplit
func (self class) ToDictionary(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsShape.Bind_to_dictionary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetShapeType(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsShape.Bind_get_shape_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShapeType(shape_type gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(shape_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsShape.Bind_set_shape_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSize() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsShape.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSize(size gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsShape.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRadius() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsShape.Bind_get_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRadius(radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsShape.Bind_set_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHeight() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsShape.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHeight(height gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsShape.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetIsTrigger() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsShape.Bind_get_is_trigger, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIsTrigger(is_trigger bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, is_trigger)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsShape.Bind_set_is_trigger, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMeshIndex() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsShape.Bind_get_mesh_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMeshIndex(mesh_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mesh_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsShape.Bind_set_mesh_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetImporterMesh(ctx gd.Lifetime) object.ImporterMesh {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsShape.Bind_get_importer_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.ImporterMesh
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetImporterMesh(importer_mesh object.ImporterMesh)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(importer_mesh[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsShape.Bind_set_importer_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsGLTFPhysicsShape() Expert { return self[0].AsGLTFPhysicsShape() }


//go:nosplit
func (self Simple) AsGLTFPhysicsShape() Simple { return self[0].AsGLTFPhysicsShape() }


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
func init() {classdb.Register("GLTFPhysicsShape", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
