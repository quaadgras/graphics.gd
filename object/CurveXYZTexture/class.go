package CurveXYZTexture

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Texture2D"
import "grow.graphics/gd/object/Texture"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A 1D texture where the red, green, and blue color channels correspond to points on 3 [Curve] resources. Compared to using separate [CurveTexture]s, this further simplifies the task of saving curves as image files.
If you only need to store one curve within a single texture, use [CurveTexture] instead. See also [GradientTexture1D] and [GradientTexture2D].

*/
type Simple [1]classdb.CurveXYZTexture
func (self Simple) SetWidth(width int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetWidth(gd.Int(width))
}
func (self Simple) SetCurveX(curve [1]classdb.Curve) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCurveX(curve)
}
func (self Simple) GetCurveX() [1]classdb.Curve {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Curve(Expert(self).GetCurveX(gc))
}
func (self Simple) SetCurveY(curve [1]classdb.Curve) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCurveY(curve)
}
func (self Simple) GetCurveY() [1]classdb.Curve {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Curve(Expert(self).GetCurveY(gc))
}
func (self Simple) SetCurveZ(curve [1]classdb.Curve) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCurveZ(curve)
}
func (self Simple) GetCurveZ() [1]classdb.Curve {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Curve(Expert(self).GetCurveZ(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.CurveXYZTexture
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetWidth(width gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CurveXYZTexture.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCurveX(curve object.Curve)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(curve[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CurveXYZTexture.Bind_set_curve_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCurveX(ctx gd.Lifetime) object.Curve {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CurveXYZTexture.Bind_get_curve_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Curve
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCurveY(curve object.Curve)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(curve[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CurveXYZTexture.Bind_set_curve_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCurveY(ctx gd.Lifetime) object.Curve {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CurveXYZTexture.Bind_get_curve_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Curve
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCurveZ(curve object.Curve)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(curve[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CurveXYZTexture.Bind_set_curve_z, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCurveZ(ctx gd.Lifetime) object.Curve {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CurveXYZTexture.Bind_get_curve_z, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Curve
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsCurveXYZTexture() Expert { return self[0].AsCurveXYZTexture() }


//go:nosplit
func (self Simple) AsCurveXYZTexture() Simple { return self[0].AsCurveXYZTexture() }


//go:nosplit
func (self class) AsTexture2D() Texture2D.Expert { return self[0].AsTexture2D() }


//go:nosplit
func (self Simple) AsTexture2D() Texture2D.Simple { return self[0].AsTexture2D() }


//go:nosplit
func (self class) AsTexture() Texture.Expert { return self[0].AsTexture() }


//go:nosplit
func (self Simple) AsTexture() Texture.Simple { return self[0].AsTexture() }


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
func init() {classdb.Register("CurveXYZTexture", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
