// Package GLTFPhysicsShape provides methods for working with GLTFPhysicsShape object instances.
package GLTFPhysicsShape

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Vector3"

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
Represents a physics shape as defined by the [code]OMI_physics_shape[/code] or [code]OMI_collider[/code] GLTF extensions. This class is an intermediary between the GLTF data and Godot's nodes, and it's abstracted in a way that allows adding support for different GLTF physics extensions in the future.
*/
type Instance [1]gdclass.GLTFPhysicsShape

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGLTFPhysicsShape() Instance
}

/*
Creates a new GLTFPhysicsShape instance from the given Godot [CollisionShape3D] node.
*/
func FromNode(shape_node [1]gdclass.CollisionShape3D) [1]gdclass.GLTFPhysicsShape { //gd:GLTFPhysicsShape.from_node
	self := Instance{}
	return [1]gdclass.GLTFPhysicsShape(class(self).FromNode(shape_node))
}

/*
Converts this GLTFPhysicsShape instance into a Godot [CollisionShape3D] node.
*/
func (self Instance) ToNode() [1]gdclass.CollisionShape3D { //gd:GLTFPhysicsShape.to_node
	return [1]gdclass.CollisionShape3D(class(self).ToNode(false))
}

/*
Creates a new GLTFPhysicsShape instance from the given Godot [Shape3D] resource.
*/
func FromResource(shape_resource [1]gdclass.Shape3D) [1]gdclass.GLTFPhysicsShape { //gd:GLTFPhysicsShape.from_resource
	self := Instance{}
	return [1]gdclass.GLTFPhysicsShape(class(self).FromResource(shape_resource))
}

/*
Converts this GLTFPhysicsShape instance into a Godot [Shape3D] resource.
*/
func (self Instance) ToResource() [1]gdclass.Shape3D { //gd:GLTFPhysicsShape.to_resource
	return [1]gdclass.Shape3D(class(self).ToResource(false))
}

/*
Creates a new GLTFPhysicsShape instance by parsing the given [Dictionary].
*/
func FromDictionary(dictionary Structure) [1]gdclass.GLTFPhysicsShape { //gd:GLTFPhysicsShape.from_dictionary
	self := Instance{}
	return [1]gdclass.GLTFPhysicsShape(class(self).FromDictionary(gd.DictionaryFromMap(dictionary)))
}

/*
Serializes this GLTFPhysicsShape instance into a [Dictionary] in the format defined by [code]OMI_physics_shape[/code].
*/
func (self Instance) ToDictionary() Structure { //gd:GLTFPhysicsShape.to_dictionary
	return Structure(gd.DictionaryAs[Structure](class(self).ToDictionary()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GLTFPhysicsShape

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFPhysicsShape"))
	casted := Instance{*(*gdclass.GLTFPhysicsShape)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) ShapeType() string {
	return string(class(self).GetShapeType().String())
}

func (self Instance) SetShapeType(value string) {
	class(self).SetShapeType(String.New(value))
}

func (self Instance) Size() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetSize())
}

func (self Instance) SetSize(value Vector3.XYZ) {
	class(self).SetSize(Vector3.XYZ(value))
}

func (self Instance) Radius() Float.X {
	return Float.X(Float.X(class(self).GetRadius()))
}

func (self Instance) SetRadius(value Float.X) {
	class(self).SetRadius(float64(value))
}

func (self Instance) Height() Float.X {
	return Float.X(Float.X(class(self).GetHeight()))
}

func (self Instance) SetHeight(value Float.X) {
	class(self).SetHeight(float64(value))
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
	class(self).SetMeshIndex(int64(value))
}

func (self Instance) ImporterMesh() [1]gdclass.ImporterMesh {
	return [1]gdclass.ImporterMesh(class(self).GetImporterMesh())
}

func (self Instance) SetImporterMesh(value [1]gdclass.ImporterMesh) {
	class(self).SetImporterMesh(value)
}

/*
Creates a new GLTFPhysicsShape instance from the given Godot [CollisionShape3D] node.
*/
//go:nosplit
func (self class) FromNode(shape_node [1]gdclass.CollisionShape3D) [1]gdclass.GLTFPhysicsShape { //gd:GLTFPhysicsShape.from_node
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shape_node[0])[0])
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_from_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.GLTFPhysicsShape{gd.PointerWithOwnershipTransferredToGo[gdclass.GLTFPhysicsShape](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Converts this GLTFPhysicsShape instance into a Godot [CollisionShape3D] node.
*/
//go:nosplit
func (self class) ToNode(cache_shapes bool) [1]gdclass.CollisionShape3D { //gd:GLTFPhysicsShape.to_node
	var frame = callframe.New()
	callframe.Arg(frame, cache_shapes)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_to_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.CollisionShape3D{gd.PointerWithOwnershipTransferredToGo[gdclass.CollisionShape3D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates a new GLTFPhysicsShape instance from the given Godot [Shape3D] resource.
*/
//go:nosplit
func (self class) FromResource(shape_resource [1]gdclass.Shape3D) [1]gdclass.GLTFPhysicsShape { //gd:GLTFPhysicsShape.from_resource
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shape_resource[0])[0])
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_from_resource, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.GLTFPhysicsShape{gd.PointerWithOwnershipTransferredToGo[gdclass.GLTFPhysicsShape](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Converts this GLTFPhysicsShape instance into a Godot [Shape3D] resource.
*/
//go:nosplit
func (self class) ToResource(cache_shapes bool) [1]gdclass.Shape3D { //gd:GLTFPhysicsShape.to_resource
	var frame = callframe.New()
	callframe.Arg(frame, cache_shapes)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_to_resource, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Shape3D{gd.PointerWithOwnershipTransferredToGo[gdclass.Shape3D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates a new GLTFPhysicsShape instance by parsing the given [Dictionary].
*/
//go:nosplit
func (self class) FromDictionary(dictionary Dictionary.Any) [1]gdclass.GLTFPhysicsShape { //gd:GLTFPhysicsShape.from_dictionary
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalDictionary(dictionary)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_from_dictionary, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.GLTFPhysicsShape{gd.PointerWithOwnershipTransferredToGo[gdclass.GLTFPhysicsShape](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Serializes this GLTFPhysicsShape instance into a [Dictionary] in the format defined by [code]OMI_physics_shape[/code].
*/
//go:nosplit
func (self class) ToDictionary() Dictionary.Any { //gd:GLTFPhysicsShape.to_dictionary
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_to_dictionary, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetShapeType() String.Readable { //gd:GLTFPhysicsShape.get_shape_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_get_shape_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShapeType(shape_type String.Readable) { //gd:GLTFPhysicsShape.set_shape_type
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(shape_type)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_set_shape_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSize() Vector3.XYZ { //gd:GLTFPhysicsShape.get_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSize(size Vector3.XYZ) { //gd:GLTFPhysicsShape.set_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRadius() float64 { //gd:GLTFPhysicsShape.get_radius
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_get_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRadius(radius float64) { //gd:GLTFPhysicsShape.set_radius
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_set_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHeight() float64 { //gd:GLTFPhysicsShape.get_height
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHeight(height float64) { //gd:GLTFPhysicsShape.set_height
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetIsTrigger() bool { //gd:GLTFPhysicsShape.get_is_trigger
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_get_is_trigger, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIsTrigger(is_trigger bool) { //gd:GLTFPhysicsShape.set_is_trigger
	var frame = callframe.New()
	callframe.Arg(frame, is_trigger)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_set_is_trigger, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMeshIndex() int64 { //gd:GLTFPhysicsShape.get_mesh_index
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_get_mesh_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMeshIndex(mesh_index int64) { //gd:GLTFPhysicsShape.set_mesh_index
	var frame = callframe.New()
	callframe.Arg(frame, mesh_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_set_mesh_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetImporterMesh() [1]gdclass.ImporterMesh { //gd:GLTFPhysicsShape.get_importer_mesh
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_get_importer_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.ImporterMesh{gd.PointerWithOwnershipTransferredToGo[gdclass.ImporterMesh](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetImporterMesh(importer_mesh [1]gdclass.ImporterMesh) { //gd:GLTFPhysicsShape.set_importer_mesh
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(importer_mesh[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsShape.Bind_set_importer_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gdclass.Register("GLTFPhysicsShape", func(ptr gd.Object) any {
		return [1]gdclass.GLTFPhysicsShape{*(*gdclass.GLTFPhysicsShape)(unsafe.Pointer(&ptr))}
	})
}

type Structure map[interface{}]interface{}
