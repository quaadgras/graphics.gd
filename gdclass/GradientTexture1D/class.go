package GradientTexture1D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Texture2D"
import "grow.graphics/gd/gdclass/Texture"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A 1D texture that obtains colors from a [Gradient] to fill the texture data. The texture is filled by sampling the gradient for each pixel. Therefore, the texture does not necessarily represent an exact copy of the gradient, as it may miss some colors if there are not enough pixels. See also [GradientTexture2D], [CurveTexture] and [CurveXYZTexture].

*/
type Go [1]classdb.GradientTexture1D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GradientTexture1D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("GradientTexture1D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Gradient() gdclass.Gradient {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Gradient(class(self).GetGradient(gc))
}

func (self Go) SetGradient(value gdclass.Gradient) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGradient(value)
}

func (self Go) UseHdr() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsUsingHdr())
}

func (self Go) SetUseHdr(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseHdr(value)
}

//go:nosplit
func (self class) SetGradient(gradient gdclass.Gradient)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(gradient[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GradientTexture1D.Bind_set_gradient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGradient(ctx gd.Lifetime) gdclass.Gradient {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GradientTexture1D.Bind_get_gradient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Gradient
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWidth(width gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GradientTexture1D.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetUseHdr(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GradientTexture1D.Bind_set_use_hdr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingHdr() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GradientTexture1D.Bind_is_using_hdr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsGradientTexture1D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGradientTexture1D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture2D() Texture2D.GD { return *((*Texture2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture2D() Texture2D.Go { return *((*Texture2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture() Texture.GD { return *((*Texture.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture() Texture.Go { return *((*Texture.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTexture2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTexture2D(), name)
	}
}
func init() {classdb.Register("GradientTexture1D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
