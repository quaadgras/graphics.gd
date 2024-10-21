package GLTFSpecGloss

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
KHR_materials_pbrSpecularGlossiness is an archived GLTF extension. This means that it is deprecated and not recommended for new files. However, it is still supported for loading old files.

*/
type Simple [1]classdb.GLTFSpecGloss
func (self Simple) GetDiffuseImg() [1]classdb.Image {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Image(Expert(self).GetDiffuseImg(gc))
}
func (self Simple) SetDiffuseImg(diffuse_img [1]classdb.Image) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDiffuseImg(diffuse_img)
}
func (self Simple) GetDiffuseFactor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetDiffuseFactor())
}
func (self Simple) SetDiffuseFactor(diffuse_factor gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDiffuseFactor(diffuse_factor)
}
func (self Simple) GetGlossFactor() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetGlossFactor()))
}
func (self Simple) SetGlossFactor(gloss_factor float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlossFactor(gd.Float(gloss_factor))
}
func (self Simple) GetSpecularFactor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetSpecularFactor())
}
func (self Simple) SetSpecularFactor(specular_factor gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSpecularFactor(specular_factor)
}
func (self Simple) GetSpecGlossImg() [1]classdb.Image {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Image(Expert(self).GetSpecGlossImg(gc))
}
func (self Simple) SetSpecGlossImg(spec_gloss_img [1]classdb.Image) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSpecGlossImg(spec_gloss_img)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.GLTFSpecGloss
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) GetDiffuseImg(ctx gd.Lifetime) object.Image {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSpecGloss.Bind_get_diffuse_img, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDiffuseImg(diffuse_img object.Image)  {
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
func (self class) GetSpecGlossImg(ctx gd.Lifetime) object.Image {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSpecGloss.Bind_get_spec_gloss_img, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSpecGlossImg(spec_gloss_img object.Image)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(spec_gloss_img[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFSpecGloss.Bind_set_spec_gloss_img, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsGLTFSpecGloss() Expert { return self[0].AsGLTFSpecGloss() }


//go:nosplit
func (self Simple) AsGLTFSpecGloss() Simple { return self[0].AsGLTFSpecGloss() }


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

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("GLTFSpecGloss", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
