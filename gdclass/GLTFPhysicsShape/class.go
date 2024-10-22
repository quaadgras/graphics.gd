package GLTFPhysicsShape

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Represents a physics shape as defined by the [code]OMI_physics_shape[/code] or [code]OMI_collider[/code] GLTF extensions. This class is an intermediary between the GLTF data and Godot's nodes, and it's abstracted in a way that allows adding support for different GLTF physics extensions in the future.

*/
type Go [1]classdb.GLTFPhysicsShape

/*
Creates a new GLTFPhysicsShape instance from the given Godot [CollisionShape3D] node.
*/
func (self Go) FromNode(shape_node gdclass.CollisionShape3D) gdclass.GLTFPhysicsShape {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.GLTFPhysicsShape(class(self).FromNode(gc, shape_node))
}

/*
Converts this GLTFPhysicsShape instance into a Godot [CollisionShape3D] node.
*/
func (self Go) ToNode() gdclass.CollisionShape3D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.CollisionShape3D(class(self).ToNode(gc, false))
}

/*
Creates a new GLTFPhysicsShape instance from the given Godot [Shape3D] resource.
*/
func (self Go) FromResource(shape_resource gdclass.Shape3D) gdclass.GLTFPhysicsShape {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.GLTFPhysicsShape(class(self).FromResource(gc, shape_resource))
}

/*
Converts this GLTFPhysicsShape instance into a Godot [Shape3D] resource.
*/
func (self Go) ToResource() gdclass.Shape3D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Shape3D(class(self).ToResource(gc, false))
}

/*
Creates a new GLTFPhysicsShape instance by parsing the given [Dictionary].
*/
func (self Go) FromDictionary(dictionary gd.Dictionary) gdclass.GLTFPhysicsShape {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.GLTFPhysicsShape(class(self).FromDictionary(gc, dictionary))
}

/*
Serializes this GLTFPhysicsShape instance into a [Dictionary] in the format defined by [code]OMI_physics_shape[/code].
*/
func (self Go) ToDictionary() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(class(self).ToDictionary(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GLTFPhysicsShape
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("GLTFPhysicsShape"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) ShapeType() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetShapeType(gc).String())
}

func (self Go) SetShapeType(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShapeType(gc.String(value))
}

func (self Go) Size() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetSize())
}

func (self Go) SetSize(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSize(value)
}

func (self Go) Radius() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetRadius()))
}

func (self Go) SetRadius(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRadius(gd.Float(value))
}

func (self Go) Height() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetHeight()))
}

func (self Go) SetHeight(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHeight(gd.Float(value))
}

func (self Go) IsTrigger() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetIsTrigger())
}

func (self Go) SetIsTrigger(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIsTrigger(value)
}

func (self Go) MeshIndex() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMeshIndex()))
}

func (self Go) SetMeshIndex(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMeshIndex(gd.Int(value))
}

func (self Go) ImporterMesh() gdclass.ImporterMesh {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.ImporterMesh(class(self).GetImporterMesh(gc))
}

func (self Go) SetImporterMesh(value gdclass.ImporterMesh) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetImporterMesh(value)
}

/*
Creates a new GLTFPhysicsShape instance from the given Godot [CollisionShape3D] node.
*/
//go:nosplit
func (self class) FromNode(ctx gd.Lifetime, shape_node gdclass.CollisionShape3D) gdclass.GLTFPhysicsShape {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(shape_node[0].AsPointer())[0])
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.GLTFPhysicsShape.Bind_from_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.GLTFPhysicsShape
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Converts this GLTFPhysicsShape instance into a Godot [CollisionShape3D] node.
*/
//go:nosplit
func (self class) ToNode(ctx gd.Lifetime, cache_shapes bool) gdclass.CollisionShape3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_shapes)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsShape.Bind_to_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.CollisionShape3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates a new GLTFPhysicsShape instance from the given Godot [Shape3D] resource.
*/
//go:nosplit
func (self class) FromResource(ctx gd.Lifetime, shape_resource gdclass.Shape3D) gdclass.GLTFPhysicsShape {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(shape_resource[0].AsPointer())[0])
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.GLTFPhysicsShape.Bind_from_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.GLTFPhysicsShape
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Converts this GLTFPhysicsShape instance into a Godot [Shape3D] resource.
*/
//go:nosplit
func (self class) ToResource(ctx gd.Lifetime, cache_shapes bool) gdclass.Shape3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_shapes)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsShape.Bind_to_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Shape3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates a new GLTFPhysicsShape instance by parsing the given [Dictionary].
*/
//go:nosplit
func (self class) FromDictionary(ctx gd.Lifetime, dictionary gd.Dictionary) gdclass.GLTFPhysicsShape {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(dictionary))
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.GLTFPhysicsShape.Bind_from_dictionary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.GLTFPhysicsShape
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
func (self class) GetImporterMesh(ctx gd.Lifetime) gdclass.ImporterMesh {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsShape.Bind_get_importer_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.ImporterMesh
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetImporterMesh(importer_mesh gdclass.ImporterMesh)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(importer_mesh[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsShape.Bind_set_importer_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsGLTFPhysicsShape() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGLTFPhysicsShape() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("GLTFPhysicsShape", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
