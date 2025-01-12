// Package StyleBoxLine provides methods for working with StyleBoxLine object instances.
package StyleBoxLine

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/StyleBox"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A [StyleBox] that displays a single line of a given color and thickness. The line can be either horizontal or vertical. Useful for separators.
*/
type Instance [1]gdclass.StyleBoxLine

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsStyleBoxLine() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.StyleBoxLine

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StyleBoxLine"))
	casted := Instance{*(*gdclass.StyleBoxLine)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Color() Color.RGBA {
	return Color.RGBA(class(self).GetColor())
}

func (self Instance) SetColor(value Color.RGBA) {
	class(self).SetColor(gd.Color(value))
}

func (self Instance) GrowBegin() Float.X {
	return Float.X(Float.X(class(self).GetGrowBegin()))
}

func (self Instance) SetGrowBegin(value Float.X) {
	class(self).SetGrowBegin(gd.Float(value))
}

func (self Instance) GrowEnd() Float.X {
	return Float.X(Float.X(class(self).GetGrowEnd()))
}

func (self Instance) SetGrowEnd(value Float.X) {
	class(self).SetGrowEnd(gd.Float(value))
}

func (self Instance) Thickness() int {
	return int(int(class(self).GetThickness()))
}

func (self Instance) SetThickness(value int) {
	class(self).SetThickness(gd.Int(value))
}

func (self Instance) Vertical() bool {
	return bool(class(self).IsVertical())
}

func (self Instance) SetVertical(value bool) {
	class(self).SetVertical(value)
}

//go:nosplit
func (self class) SetColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxLine.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxLine.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetThickness(thickness gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, thickness)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxLine.Bind_set_thickness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetThickness() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxLine.Bind_get_thickness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGrowBegin(offset gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxLine.Bind_set_grow_begin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGrowBegin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxLine.Bind_get_grow_begin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGrowEnd(offset gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxLine.Bind_set_grow_end, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGrowEnd() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxLine.Bind_get_grow_end, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVertical(vertical bool) {
	var frame = callframe.New()
	callframe.Arg(frame, vertical)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxLine.Bind_set_vertical, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsVertical() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxLine.Bind_is_vertical, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsStyleBoxLine() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsStyleBoxLine() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsStyleBox() StyleBox.Advanced {
	return *((*StyleBox.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsStyleBox() StyleBox.Instance {
	return *((*StyleBox.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(StyleBox.Advanced(self.AsStyleBox()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(StyleBox.Instance(self.AsStyleBox()), name)
	}
}
func init() {
	gdclass.Register("StyleBoxLine", func(ptr gd.Object) any {
		return [1]gdclass.StyleBoxLine{*(*gdclass.StyleBoxLine)(unsafe.Pointer(&ptr))}
	})
}
