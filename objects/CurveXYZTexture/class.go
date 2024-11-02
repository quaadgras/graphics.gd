package CurveXYZTexture

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Texture2D"
import "grow.graphics/gd/objects/Texture"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A 1D texture where the red, green, and blue color channels correspond to points on 3 [Curve] resources. Compared to using separate [CurveTexture]s, this further simplifies the task of saving curves as image files.
If you only need to store one curve within a single texture, use [CurveTexture] instead. See also [GradientTexture1D] and [GradientTexture2D].
*/
type Instance [1]classdb.CurveXYZTexture

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.CurveXYZTexture

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CurveXYZTexture"))
	return Instance{classdb.CurveXYZTexture(object)}
}

func (self Instance) CurveX() objects.Curve {
	return objects.Curve(class(self).GetCurveX())
}

func (self Instance) SetCurveX(value objects.Curve) {
	class(self).SetCurveX(value)
}

func (self Instance) CurveY() objects.Curve {
	return objects.Curve(class(self).GetCurveY())
}

func (self Instance) SetCurveY(value objects.Curve) {
	class(self).SetCurveY(value)
}

func (self Instance) CurveZ() objects.Curve {
	return objects.Curve(class(self).GetCurveZ())
}

func (self Instance) SetCurveZ(value objects.Curve) {
	class(self).SetCurveZ(value)
}

//go:nosplit
func (self class) SetWidth(width gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CurveXYZTexture.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetCurveX(curve objects.Curve) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CurveXYZTexture.Bind_set_curve_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCurveX() objects.Curve {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CurveXYZTexture.Bind_get_curve_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Curve{classdb.Curve(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCurveY(curve objects.Curve) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CurveXYZTexture.Bind_set_curve_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCurveY() objects.Curve {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CurveXYZTexture.Bind_get_curve_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Curve{classdb.Curve(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCurveZ(curve objects.Curve) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CurveXYZTexture.Bind_set_curve_z, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCurveZ() objects.Curve {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CurveXYZTexture.Bind_get_curve_z, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Curve{classdb.Curve(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsCurveXYZTexture() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCurveXYZTexture() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTexture2D() Texture2D.Advanced {
	return *((*Texture2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTexture2D() Texture2D.Instance {
	return *((*Texture2D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsTexture() Texture.Advanced { return *((*Texture.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTexture() Texture.Instance {
	return *((*Texture.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(self.AsTexture2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsTexture2D(), name)
	}
}
func init() {
	classdb.Register("CurveXYZTexture", func(ptr gd.Object) any { return classdb.CurveXYZTexture(ptr) })
}
