package GLTFPhysicsShape

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"
import "grow.graphics/gd/variant/Dictionary"
import "grow.graphics/gd/variant/Vector3"
import "grow.graphics/gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Represents a physics shape as defined by the [code]OMI_physics_shape[/code] or [code]OMI_collider[/code] GLTF extensions. This class is an intermediary between the GLTF data and Godot's nodes, and it's abstracted in a way that allows adding support for different GLTF physics extensions in the future.
*/
type Instance [1]classdb.GLTFPhysicsShape

/*
Creates a new GLTFPhysicsShape instance from the given Godot [CollisionShape3D] node.
*/
func (self Instance) FromNode(shape_node objects.CollisionShape3D) objects.GLTFPhysicsShape {
	return objects.GLTFPhysicsShape(class(self).FromNode(shape_node))
}

/*
Converts this GLTFPhysicsShape instance into a Godot [CollisionShape3D] node.
*/
func (self Instance) ToNode() objects.CollisionShape3D {
	return objects.CollisionShape3D(class(self).ToNode(false))
}

/*
Creates a new GLTFPhysicsShape instance from the given Godot [Shape3D] resource.
*/
func (self Instance) FromResource(shape_resource objects.Shape3D) objects.GLTFPhysicsShape {
	return objects.GLTFPhysicsShape(class(self).FromResource(shape_resource))
}

/*
Converts this GLTFPhysicsShape instance into a Godot [Shape3D] resource.
*/
func (self Instance) ToResource() objects.Shape3D {
	return objects.Shape3D(class(self).ToResource(false))
}

/*
Creates a new GLTFPhysicsShape instance by parsing the given [Dictionary].
*/
func (self Instance) FromDictionary(dictionary Dictionary.Any) objects.GLTFPhysicsShape {
	return objects.GLTFPhysicsShape(class(self).FromDictionary(dictionary))
}

/*
Serializes this GLTFPhysicsShape instance into a [Dictionary] in the format defined by [code]OMI_physics_shape[/code].
*/
func (self Instance) ToDictionary() Dictionary.Any {
	return Dictionary.Any(class(self).ToDictionary())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.GLTFPhysicsShape

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFPhysicsShape"))
	return Instance{classdb.GLTFPhysicsShape(object)}
}

func (self Instance) ShapeType() string {
	return string(class(self).GetShapeType().String())
}

func (self Instance) SetShapeType(value string) {
	class(self).SetShapeType(gd.NewString(value))
}

func (self Instance) Size() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetSize())
}

func (self Instance) SetSize(value Vector3.XYZ) {
	class(self).SetSize(gd.Vector3(value))
}

func (self Instance) Radius() Float.X {
	return Float.X(Float.X(class(self).GetRadius()))
}

func (self Instance) SetRadius(value Float.X) {
	class(self).SetRadius(gd.Float(value))
}

func (self Instance) Height() Float.X {
	return Float.X(Float.X(class(self).GetHeight()))
}

func (self Instance) SetHeight(value Float.X) {
	class(self).SetHeight(gd.Float(value))
}

func (self Instance) IsTrigger() bool {
	return bool(class(self).GetIsTrigger())
}

func (self Instance) SetIsTrigger(value bool) {
	class(self).SetIsTrigger(value)
}

func (self Instance) MeshIndex() int {
	return int(int(class(self).GetMeshIndex()))
}

func (self Instance) SetMeshIndex(value int) {
	class(self).SetMeshIndex(gd.Int(value))
}

func (self Instance) ImporterMesh() objects.ImporterMesh {
	return objects.ImporterMesh(class(self).GetImporterMesh())
}

func (self Instance) SetImporterMesh(value objects.ImporterMesh) {
	class(self).SetImporterMesh(value)
}

/*
Creates a new GLTFPhysicsShape instance from the given Godot [CollisionShape3D] node.
*/
//go:nosplit
func (self class) FromNode(shape_node objects.CollisionShape3D) objects.GLTFPhysicsShape {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shape_node[0])[0])
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_from_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.GLTFPhysicsShape{classdb.GLTFPhysicsShape(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Converts this GLTFPhysicsShape instance into a Godot [CollisionShape3D] node.
*/
//go:nosplit
func (self class) ToNode(cache_shapes bool) objects.CollisionShape3D {
	var frame = callframe.New()
	callframe.Arg(frame, cache_shapes)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_to_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.CollisionShape3D{classdb.CollisionShape3D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Creates a new GLTFPhysicsShape instance from the given Godot [Shape3D] resource.
*/
//go:nosplit
func (self class) FromResource(shape_resource objects.Shape3D) objects.GLTFPhysicsShape {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shape_resource[0])[0])
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_from_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.GLTFPhysicsShape{classdb.GLTFPhysicsShape(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Converts this GLTFPhysicsShape instance into a Godot [Shape3D] resource.
*/
//go:nosplit
func (self class) ToResource(cache_shapes bool) objects.Shape3D {
	var frame = callframe.New()
	callframe.Arg(frame, cache_shapes)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_to_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Shape3D{classdb.Shape3D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Creates a new GLTFPhysicsShape instance by parsing the given [Dictionary].
*/
//go:nosplit
func (self class) FromDictionary(dictionary gd.Dictionary) objects.GLTFPhysicsShape {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(dictionary))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_from_dictionary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.GLTFPhysicsShape{classdb.GLTFPhysicsShape(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Serializes this GLTFPhysicsShape instance into a [Dictionary] in the format defined by [code]OMI_physics_shape[/code].
*/
//go:nosplit
func (self class) ToDictionary() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_to_dictionary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetShapeType() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_get_shape_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShapeType(shape_type gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shape_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_set_shape_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSize() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSize(size gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_get_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRadius(radius gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_set_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHeight() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHeight(height gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetIsTrigger() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_get_is_trigger, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIsTrigger(is_trigger bool) {
	var frame = callframe.New()
	callframe.Arg(frame, is_trigger)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_set_is_trigger, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMeshIndex() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_get_mesh_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMeshIndex(mesh_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mesh_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_set_mesh_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetImporterMesh() objects.ImporterMesh {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_get_importer_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.ImporterMesh{classdb.ImporterMesh(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetImporterMesh(importer_mesh objects.ImporterMesh) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(importer_mesh[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_set_importer_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsGLTFPhysicsShape() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGLTFPhysicsShape() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	classdb.Register("GLTFPhysicsShape", func(ptr gd.Object) any { return classdb.GLTFPhysicsShape(ptr) })
}
