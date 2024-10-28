package GLTFSpecGloss

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
KHR_materials_pbrSpecularGlossiness is an archived GLTF extension. This means that it is deprecated and not recommended for new files. However, it is still supported for loading old files.

*/
type Go [1]classdb.GLTFSpecGloss
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GLTFSpecGloss
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFSpecGloss"))
	return Go{classdb.GLTFSpecGloss(object)}
}

func (self Go) DiffuseImg() gdclass.Image {
		return gdclass.Image(class(self).GetDiffuseImg())
}

func (self Go) SetDiffuseImg(value gdclass.Image) {
	class(self).SetDiffuseImg(value)
}

func (self Go) DiffuseFactor() gd.Color {
		return gd.Color(class(self).GetDiffuseFactor())
}

func (self Go) SetDiffuseFactor(value gd.Color) {
	class(self).SetDiffuseFactor(value)
}

func (self Go) GlossFactor() float64 {
		return float64(float64(class(self).GetGlossFactor()))
}

func (self Go) SetGlossFactor(value float64) {
	class(self).SetGlossFactor(gd.Float(value))
}

func (self Go) SpecularFactor() gd.Color {
		return gd.Color(class(self).GetSpecularFactor())
}

func (self Go) SetSpecularFactor(value gd.Color) {
	class(self).SetSpecularFactor(value)
}

func (self Go) SpecGlossImg() gdclass.Image {
		return gdclass.Image(class(self).GetSpecGlossImg())
}

func (self Go) SetSpecGlossImg(value gdclass.Image) {
	class(self).SetSpecGlossImg(value)
}

//go:nosplit
func (self class) GetDiffuseImg() gdclass.Image {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_get_diffuse_img, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Image{classdb.Image(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDiffuseImg(diffuse_img gdclass.Image)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(diffuse_img[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_set_diffuse_img, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDiffuseFactor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_get_diffuse_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDiffuseFactor(diffuse_factor gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, diffuse_factor)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_set_diffuse_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlossFactor() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_get_gloss_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlossFactor(gloss_factor gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, gloss_factor)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_set_gloss_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSpecularFactor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_get_specular_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSpecularFactor(specular_factor gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, specular_factor)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_set_specular_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSpecGlossImg() gdclass.Image {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_get_spec_gloss_img, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Image{classdb.Image(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSpecGlossImg(spec_gloss_img gdclass.Image)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(spec_gloss_img[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_set_spec_gloss_img, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func init() {classdb.Register("GLTFSpecGloss", func(ptr gd.Object) any { return classdb.GLTFSpecGloss(ptr) })}
