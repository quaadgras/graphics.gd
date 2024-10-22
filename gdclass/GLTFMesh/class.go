package GLTFMesh

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
GLTFMesh handles 3D mesh data imported from GLTF files. It includes properties for blend channels, blend weights, instance materials, and the mesh itself.

*/
type Go [1]classdb.GLTFMesh

/*
Gets additional arbitrary data in this [GLTFMesh] instance. This can be used to keep per-node state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the return value can be anything you set. If nothing was set, the return value is null.
*/
func (self Go) GetAdditionalData(extension_name string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).GetAdditionalData(gc, gc.StringName(extension_name)))
}

/*
Sets additional arbitrary data in this [GLTFMesh] instance. This can be used to keep per-node state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The first argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the second argument can be anything you want.
*/
func (self Go) SetAdditionalData(extension_name string, additional_data gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAdditionalData(gc.StringName(extension_name), additional_data)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GLTFMesh
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("GLTFMesh"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) OriginalName() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetOriginalName(gc).String())
}

func (self Go) SetOriginalName(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOriginalName(gc.String(value))
}

func (self Go) Mesh() gdclass.ImporterMesh {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.ImporterMesh(class(self).GetMesh(gc))
}

func (self Go) SetMesh(value gdclass.ImporterMesh) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMesh(value)
}

func (self Go) BlendWeights() []float32 {
	gc := gd.GarbageCollector(); _ = gc
		return []float32(class(self).GetBlendWeights(gc).AsSlice())
}

func (self Go) SetBlendWeights(value []float32) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBlendWeights(gc.PackedFloat32Slice(value))
}

func (self Go) InstanceMaterials() gd.ArrayOf[gdclass.Material] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gdclass.Material](class(self).GetInstanceMaterials(gc))
}

func (self Go) SetInstanceMaterials(value gd.ArrayOf[gdclass.Material]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetInstanceMaterials(value)
}

//go:nosplit
func (self class) GetOriginalName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFMesh.Bind_get_original_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOriginalName(original_name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(original_name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFMesh.Bind_set_original_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMesh(ctx gd.Lifetime) gdclass.ImporterMesh {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFMesh.Bind_get_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.ImporterMesh
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMesh(mesh gdclass.ImporterMesh)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(mesh[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFMesh.Bind_set_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBlendWeights(ctx gd.Lifetime) gd.PackedFloat32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFMesh.Bind_get_blend_weights, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedFloat32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBlendWeights(blend_weights gd.PackedFloat32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(blend_weights))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFMesh.Bind_set_blend_weights, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInstanceMaterials(ctx gd.Lifetime) gd.ArrayOf[gdclass.Material] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFMesh.Bind_get_instance_materials, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.Material](ret)
}
//go:nosplit
func (self class) SetInstanceMaterials(instance_materials gd.ArrayOf[gdclass.Material])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(instance_materials))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFMesh.Bind_set_instance_materials, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets additional arbitrary data in this [GLTFMesh] instance. This can be used to keep per-node state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the return value can be anything you set. If nothing was set, the return value is null.
*/
//go:nosplit
func (self class) GetAdditionalData(ctx gd.Lifetime, extension_name gd.StringName) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(extension_name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFMesh.Bind_get_additional_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets additional arbitrary data in this [GLTFMesh] instance. This can be used to keep per-node state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The first argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the second argument can be anything you want.
*/
//go:nosplit
func (self class) SetAdditionalData(extension_name gd.StringName, additional_data gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(extension_name))
	callframe.Arg(frame, mmm.Get(additional_data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFMesh.Bind_set_additional_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsGLTFMesh() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGLTFMesh() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("GLTFMesh", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
