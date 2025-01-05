// Package GLTFMesh provides methods for working with GLTFMesh object instances.
package GLTFMesh

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
GLTFMesh handles 3D mesh data imported from GLTF files. It includes properties for blend channels, blend weights, instance materials, and the mesh itself.
*/
type Instance [1]gdclass.GLTFMesh
type Any interface {
	gd.IsClass
	AsGLTFMesh() Instance
}

/*
Gets additional arbitrary data in this [GLTFMesh] instance. This can be used to keep per-node state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the return value can be anything you set. If nothing was set, the return value is null.
*/
func (self Instance) GetAdditionalData(extension_name string) any {
	return any(class(self).GetAdditionalData(gd.NewStringName(extension_name)).Interface())
}

/*
Sets additional arbitrary data in this [GLTFMesh] instance. This can be used to keep per-node state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The first argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the second argument can be anything you want.
*/
func (self Instance) SetAdditionalData(extension_name string, additional_data any) {
	class(self).SetAdditionalData(gd.NewStringName(extension_name), gd.NewVariant(additional_data))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GLTFMesh

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFMesh"))
	return Instance{*(*gdclass.GLTFMesh)(unsafe.Pointer(&object))}
}

func (self Instance) OriginalName() string {
	return string(class(self).GetOriginalName().String())
}

func (self Instance) SetOriginalName(value string) {
	class(self).SetOriginalName(gd.NewString(value))
}

func (self Instance) Mesh() [1]gdclass.ImporterMesh {
	return [1]gdclass.ImporterMesh(class(self).GetMesh())
}

func (self Instance) SetMesh(value [1]gdclass.ImporterMesh) {
	class(self).SetMesh(value)
}

func (self Instance) BlendWeights() []float32 {
	return []float32(class(self).GetBlendWeights().AsSlice())
}

func (self Instance) SetBlendWeights(value []float32) {
	class(self).SetBlendWeights(gd.NewPackedFloat32Slice(value))
}

func (self Instance) InstanceMaterials() gd.Array {
	return gd.Array(class(self).GetInstanceMaterials())
}

func (self Instance) SetInstanceMaterials(value gd.Array) {
	class(self).SetInstanceMaterials(value)
}

//go:nosplit
func (self class) GetOriginalName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFMesh.Bind_get_original_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOriginalName(original_name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(original_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFMesh.Bind_set_original_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMesh() [1]gdclass.ImporterMesh {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFMesh.Bind_get_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.ImporterMesh{gd.PointerWithOwnershipTransferredToGo[gdclass.ImporterMesh](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMesh(mesh [1]gdclass.ImporterMesh) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFMesh.Bind_set_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBlendWeights() gd.PackedFloat32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFMesh.Bind_get_blend_weights, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedFloat32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBlendWeights(blend_weights gd.PackedFloat32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(blend_weights))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFMesh.Bind_set_blend_weights, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetInstanceMaterials() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFMesh.Bind_get_instance_materials, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInstanceMaterials(instance_materials gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(instance_materials))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFMesh.Bind_set_instance_materials, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Gets additional arbitrary data in this [GLTFMesh] instance. This can be used to keep per-node state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the return value can be anything you set. If nothing was set, the return value is null.
*/
//go:nosplit
func (self class) GetAdditionalData(extension_name gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension_name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFMesh.Bind_get_additional_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets additional arbitrary data in this [GLTFMesh] instance. This can be used to keep per-node state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The first argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the second argument can be anything you want.
*/
//go:nosplit
func (self class) SetAdditionalData(extension_name gd.StringName, additional_data gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension_name))
	callframe.Arg(frame, pointers.Get(additional_data))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFMesh.Bind_set_additional_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsGLTFMesh() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGLTFMesh() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("GLTFMesh", func(ptr gd.Object) any { return [1]gdclass.GLTFMesh{*(*gdclass.GLTFMesh)(unsafe.Pointer(&ptr))} })
}
