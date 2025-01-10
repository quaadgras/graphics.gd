// Package NinePatchRect provides methods for working with NinePatchRect object instances.
package NinePatchRect

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Rect2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
Also known as 9-slice panels, [NinePatchRect] produces clean panels of any size based on a small texture. To do so, it splits the texture in a 3×3 grid. When you scale the node, it tiles the texture's edges horizontally or vertically, tiles the center on both axes, and leaves the corners unchanged.
*/
type Instance [1]gdclass.NinePatchRect

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsNinePatchRect() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.NinePatchRect

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NinePatchRect"))
	casted := Instance{*(*gdclass.NinePatchRect)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Texture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture())
}

func (self Instance) SetTexture(value [1]gdclass.Texture2D) {
	class(self).SetTexture(value)
}

func (self Instance) DrawCenter() bool {
	return bool(class(self).IsDrawCenterEnabled())
}

func (self Instance) SetDrawCenter(value bool) {
	class(self).SetDrawCenter(value)
}

func (self Instance) RegionRect() Rect2.PositionSize {
	return Rect2.PositionSize(class(self).GetRegionRect())
}

func (self Instance) SetRegionRect(value Rect2.PositionSize) {
	class(self).SetRegionRect(gd.Rect2(value))
}

func (self Instance) PatchMarginLeft() int {
	return int(int(class(self).GetPatchMargin(0)))
}

func (self Instance) SetPatchMarginLeft(value int) {
	class(self).SetPatchMargin(0, gd.Int(value))
}

func (self Instance) PatchMarginTop() int {
	return int(int(class(self).GetPatchMargin(1)))
}

func (self Instance) SetPatchMarginTop(value int) {
	class(self).SetPatchMargin(1, gd.Int(value))
}

func (self Instance) PatchMarginRight() int {
	return int(int(class(self).GetPatchMargin(2)))
}

func (self Instance) SetPatchMarginRight(value int) {
	class(self).SetPatchMargin(2, gd.Int(value))
}

func (self Instance) PatchMarginBottom() int {
	return int(int(class(self).GetPatchMargin(3)))
}

func (self Instance) SetPatchMarginBottom(value int) {
	class(self).SetPatchMargin(3, gd.Int(value))
}

func (self Instance) AxisStretchHorizontal() gdclass.NinePatchRectAxisStretchMode {
	return gdclass.NinePatchRectAxisStretchMode(class(self).GetHAxisStretchMode())
}

func (self Instance) SetAxisStretchHorizontal(value gdclass.NinePatchRectAxisStretchMode) {
	class(self).SetHAxisStretchMode(value)
}

func (self Instance) AxisStretchVertical() gdclass.NinePatchRectAxisStretchMode {
	return gdclass.NinePatchRectAxisStretchMode(class(self).GetVAxisStretchMode())
}

func (self Instance) SetAxisStretchVertical(value gdclass.NinePatchRectAxisStretchMode) {
	class(self).SetVAxisStretchMode(value)
}

//go:nosplit
func (self class) SetTexture(texture [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NinePatchRect.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() [1]gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NinePatchRect.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the size of the margin on the specified [enum Side] to [param value] pixels.
*/
//go:nosplit
func (self class) SetPatchMargin(margin Side, value gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NinePatchRect.Bind_set_patch_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the size of the margin on the specified [enum Side].
*/
//go:nosplit
func (self class) GetPatchMargin(margin Side) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NinePatchRect.Bind_get_patch_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRegionRect(rect gd.Rect2) {
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NinePatchRect.Bind_set_region_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRegionRect() gd.Rect2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NinePatchRect.Bind_get_region_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrawCenter(draw_center bool) {
	var frame = callframe.New()
	callframe.Arg(frame, draw_center)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NinePatchRect.Bind_set_draw_center, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsDrawCenterEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NinePatchRect.Bind_is_draw_center_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHAxisStretchMode(mode gdclass.NinePatchRectAxisStretchMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NinePatchRect.Bind_set_h_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHAxisStretchMode() gdclass.NinePatchRectAxisStretchMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NinePatchRectAxisStretchMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NinePatchRect.Bind_get_h_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVAxisStretchMode(mode gdclass.NinePatchRectAxisStretchMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NinePatchRect.Bind_set_v_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVAxisStretchMode() gdclass.NinePatchRectAxisStretchMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NinePatchRectAxisStretchMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NinePatchRect.Bind_get_v_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnTextureChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("texture_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsNinePatchRect() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNinePatchRect() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.Advanced  { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Control.Advanced(self.AsControl()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Control.Instance(self.AsControl()), name)
	}
}
func init() {
	gdclass.Register("NinePatchRect", func(ptr gd.Object) any {
		return [1]gdclass.NinePatchRect{*(*gdclass.NinePatchRect)(unsafe.Pointer(&ptr))}
	})
}

type AxisStretchMode = gdclass.NinePatchRectAxisStretchMode

const (
	/*Stretches the center texture across the NinePatchRect. This may cause the texture to be distorted.*/
	AxisStretchModeStretch AxisStretchMode = 0
	/*Repeats the center texture across the NinePatchRect. This won't cause any visible distortion. The texture must be seamless for this to work without displaying artifacts between edges.*/
	AxisStretchModeTile AxisStretchMode = 1
	/*Repeats the center texture across the NinePatchRect, but will also stretch the texture to make sure each tile is visible in full. This may cause the texture to be distorted, but less than [constant AXIS_STRETCH_MODE_STRETCH]. The texture must be seamless for this to work without displaying artifacts between edges.*/
	AxisStretchModeTileFit AxisStretchMode = 2
)

type Side int

const (
	/*Left side, usually used for [Control] or [StyleBox]-derived classes.*/
	SideLeft Side = 0
	/*Top side, usually used for [Control] or [StyleBox]-derived classes.*/
	SideTop Side = 1
	/*Right side, usually used for [Control] or [StyleBox]-derived classes.*/
	SideRight Side = 2
	/*Bottom side, usually used for [Control] or [StyleBox]-derived classes.*/
	SideBottom Side = 3
)
