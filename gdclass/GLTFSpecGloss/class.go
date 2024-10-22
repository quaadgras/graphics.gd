package GLTFSpecGloss

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
KHR_materials_pbrSpecularGlossiness is an archived GLTF extension. This means that it is deprecated and not recommended for new files. However, it is still supported for loading old files.

*/
type Go [1]classdb.GLTFSpecGloss
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GLTFSpecGloss
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("GLTFSpecGloss"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) DiffuseImg() gdclass.Image {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Image(class(self).GetDiffuseImg(gc))
}

func (self Go) SetDiffuseImg(value gdclass.Image) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDiffuseImg(value)
}

func (self Go) DiffuseFactor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetDiffuseFactor())
}

func (self Go) SetDiffuseFactor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDiffuseFactor(value)
}

func (self Go) GlossFactor() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetGlossFactor()))
}

func (self Go) SetGlossFactor(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGlossFactor(gd.Float(value))
}

func (self Go) SpecularFactor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetSpecularFactor())
}

func (self Go) SetSpecularFactor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSpecularFactor(value)
}

func (self Go) SpecGlossImg() gdclass.Image {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Image(class(self).GetSpecGlossImg(gc))
}

func (self Go) SetSpecGlossImg(value gdclass.Image) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSpecGlossImg(value)
}

//go:nosplit
func (self class) GetDiffuseImg(ctx gd.Lifetime) gdclass.Image {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSpecGloss.Bind_get_diffuse_img, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDiffuseImg(diffuse_img gdclass.Image)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(diffuse_img[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSpecGloss.Bind_set_diffuse_img, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDiffuseFactor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSpecGloss.Bind_get_diffuse_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDiffuseFactor(diffuse_factor gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, diffuse_factor)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSpecGloss.Bind_set_diffuse_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlossFactor() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSpecGloss.Bind_get_gloss_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlossFactor(gloss_factor gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gloss_factor)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSpecGloss.Bind_set_gloss_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSpecularFactor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSpecGloss.Bind_get_specular_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSpecularFactor(specular_factor gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, specular_factor)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSpecGloss.Bind_set_specular_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSpecGlossImg(ctx gd.Lifetime) gdclass.Image {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSpecGloss.Bind_get_spec_gloss_img, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSpecGlossImg(spec_gloss_img gdclass.Image)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(spec_gloss_img[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSpecGloss.Bind_set_spec_gloss_img, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsGLTFSpecGloss() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGLTFSpecGloss() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("GLTFSpecGloss", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
