// Package GLTFNode provides methods for working with GLTFNode object instances.
package GLTFNode

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Quaternion"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Represents a GLTF node. GLTF nodes may have names, transforms, children (other GLTF nodes), and more specialized properties (represented by their own classes).
GLTF nodes generally exist inside of [GLTFState] which represents all data of a GLTF file. Most of GLTFNode's properties are indices of other data in the GLTF file. You can extend a GLTF node with additional properties by using [method get_additional_data] and [method set_additional_data].
*/
type Instance [1]gdclass.GLTFNode

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGLTFNode() Instance
}

/*
Gets additional arbitrary data in this [GLTFNode] instance. This can be used to keep per-node state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the return value can be anything you set. If nothing was set, the return value is null.
*/
func (self Instance) GetAdditionalData(extension_name string) any {
	return any(class(self).GetAdditionalData(gd.NewStringName(extension_name)).Interface())
}

/*
Sets additional arbitrary data in this [GLTFNode] instance. This can be used to keep per-node state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The first argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the second argument can be anything you want.
*/
func (self Instance) SetAdditionalData(extension_name string, additional_data any) {
	class(self).SetAdditionalData(gd.NewStringName(extension_name), gd.NewVariant(additional_data))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GLTFNode

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFNode"))
	casted := Instance{*(*gdclass.GLTFNode)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) OriginalName() string {
	return string(class(self).GetOriginalName().String())
}

func (self Instance) SetOriginalName(value string) {
	class(self).SetOriginalName(gd.NewString(value))
}

func (self Instance) Parent() int {
	return int(int(class(self).GetParent()))
}

func (self Instance) SetParent(value int) {
	class(self).SetParent(gd.Int(value))
}

func (self Instance) Height() int {
	return int(int(class(self).GetHeight()))
}

func (self Instance) SetHeight(value int) {
	class(self).SetHeight(gd.Int(value))
}

func (self Instance) Xform() Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetXform())
}

func (self Instance) SetXform(value Transform3D.BasisOrigin) {
	class(self).SetXform(gd.Transform3D(value))
}

func (self Instance) Mesh() int {
	return int(int(class(self).GetMesh()))
}

func (self Instance) SetMesh(value int) {
	class(self).SetMesh(gd.Int(value))
}

func (self Instance) Camera() int {
	return int(int(class(self).GetCamera()))
}

func (self Instance) SetCamera(value int) {
	class(self).SetCamera(gd.Int(value))
}

func (self Instance) Skin() int {
	return int(int(class(self).GetSkin()))
}

func (self Instance) SetSkin(value int) {
	class(self).SetSkin(gd.Int(value))
}

func (self Instance) Skeleton() int {
	return int(int(class(self).GetSkeleton()))
}

func (self Instance) SetSkeleton(value int) {
	class(self).SetSkeleton(gd.Int(value))
}

func (self Instance) Position() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetPosition())
}

func (self Instance) SetPosition(value Vector3.XYZ) {
	class(self).SetPosition(gd.Vector3(value))
}

func (self Instance) Rotation() Quaternion.IJKX {
	return Quaternion.IJKX(class(self).GetRotation())
}

func (self Instance) SetRotation(value Quaternion.IJKX) {
	class(self).SetRotation(gd.Quaternion(value))
}

func (self Instance) Scale() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetScale())
}

func (self Instance) SetScale(value Vector3.XYZ) {
	class(self).SetScale(gd.Vector3(value))
}

func (self Instance) Children() []int32 {
	return []int32(class(self).GetChildren().AsSlice())
}

func (self Instance) SetChildren(value []int32) {
	class(self).SetChildren(gd.NewPackedInt32Slice(value))
}

func (self Instance) Light() int {
	return int(int(class(self).GetLight()))
}

func (self Instance) SetLight(value int) {
	class(self).SetLight(gd.Int(value))
}

//go:nosplit
func (self class) GetOriginalName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_get_original_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOriginalName(original_name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(original_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_set_original_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetParent() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_get_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetParent(parent gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, parent)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_set_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHeight() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHeight(height gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetXform() gd.Transform3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_get_xform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetXform(xform gd.Transform3D) {
	var frame = callframe.New()
	callframe.Arg(frame, xform)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_set_xform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMesh() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_get_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMesh(mesh gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mesh)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_set_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCamera() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_get_camera, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCamera(camera gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, camera)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_set_camera, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkin() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_get_skin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkin(skin gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, skin)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_set_skin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkeleton() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_get_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkeleton(skeleton gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, skeleton)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_set_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPosition() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPosition(position gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_set_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRotation() gd.Quaternion {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Quaternion](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_get_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRotation(rotation gd.Quaternion) {
	var frame = callframe.New()
	callframe.Arg(frame, rotation)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_set_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetScale() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_get_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScale(scale gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_set_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetChildren() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_get_children, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetChildren(children gd.PackedInt32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(children))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_set_children, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLight() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_get_light, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLight(light gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, light)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_set_light, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Gets additional arbitrary data in this [GLTFNode] instance. This can be used to keep per-node state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the return value can be anything you set. If nothing was set, the return value is null.
*/
//go:nosplit
func (self class) GetAdditionalData(extension_name gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension_name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_get_additional_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets additional arbitrary data in this [GLTFNode] instance. This can be used to keep per-node state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The first argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the second argument can be anything you want.
*/
//go:nosplit
func (self class) SetAdditionalData(extension_name gd.StringName, additional_data gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension_name))
	callframe.Arg(frame, pointers.Get(additional_data))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFNode.Bind_set_additional_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsGLTFNode() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGLTFNode() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("GLTFNode", func(ptr gd.Object) any { return [1]gdclass.GLTFNode{*(*gdclass.GLTFNode)(unsafe.Pointer(&ptr))} })
}
