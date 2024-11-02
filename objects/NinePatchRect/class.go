package NinePatchRect

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Control"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"
import "grow.graphics/gd/variant/Rect2"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Also known as 9-slice panels, [NinePatchRect] produces clean panels of any size based on a small texture. To do so, it splits the texture in a 3×3 grid. When you scale the node, it tiles the texture's edges horizontally or vertically, tiles the center on both axes, and leaves the corners unchanged.
*/
type Instance [1]classdb.NinePatchRect

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.NinePatchRect

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NinePatchRect"))
	return Instance{classdb.NinePatchRect(object)}
}

func (self Instance) Texture() objects.Texture2D {
	return objects.Texture2D(class(self).GetTexture())
}

func (self Instance) SetTexture(value objects.Texture2D) {
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

func (self Instance) AxisStretchHorizontal() classdb.NinePatchRectAxisStretchMode {
	return classdb.NinePatchRectAxisStretchMode(class(self).GetHAxisStretchMode())
}

func (self Instance) SetAxisStretchHorizontal(value classdb.NinePatchRectAxisStretchMode) {
	class(self).SetHAxisStretchMode(value)
}

func (self Instance) AxisStretchVertical() classdb.NinePatchRectAxisStretchMode {
	return classdb.NinePatchRectAxisStretchMode(class(self).GetVAxisStretchMode())
}

func (self Instance) SetAxisStretchVertical(value classdb.NinePatchRectAxisStretchMode) {
	class(self).SetVAxisStretchMode(value)
}

//go:nosplit
func (self class) SetTexture(texture objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NinePatchRect.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() objects.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NinePatchRect.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
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
func (self class) SetHAxisStretchMode(mode classdb.NinePatchRectAxisStretchMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NinePatchRect.Bind_set_h_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHAxisStretchMode() classdb.NinePatchRectAxisStretchMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NinePatchRectAxisStretchMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NinePatchRect.Bind_get_h_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVAxisStretchMode(mode classdb.NinePatchRectAxisStretchMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NinePatchRect.Bind_set_v_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVAxisStretchMode() classdb.NinePatchRectAxisStretchMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NinePatchRectAxisStretchMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NinePatchRect.Bind_get_v_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnTextureChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("texture_changed"), gd.NewCallable(cb), 0)
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
		return gd.VirtualByName(self.AsControl(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsControl(), name)
	}
}
func init() {
	classdb.Register("NinePatchRect", func(ptr gd.Object) any { return classdb.NinePatchRect(ptr) })
}

type AxisStretchMode = classdb.NinePatchRectAxisStretchMode

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
