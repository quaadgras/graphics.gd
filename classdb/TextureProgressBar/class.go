// Package TextureProgressBar provides methods for working with TextureProgressBar object instances.
package TextureProgressBar

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Range"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
TextureProgressBar works like [ProgressBar], but uses up to 3 textures instead of Godot's [Theme] resource. It can be used to create horizontal, vertical and radial progress bars.
*/
type Instance [1]gdclass.TextureProgressBar
type Any interface {
	gd.IsClass
	AsTextureProgressBar() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TextureProgressBar

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TextureProgressBar"))
	casted := Instance{*(*gdclass.TextureProgressBar)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) FillMode() int {
	return int(int(class(self).GetFillMode()))
}

func (self Instance) SetFillMode(value int) {
	class(self).SetFillMode(gd.Int(value))
}

func (self Instance) RadialInitialAngle() Float.X {
	return Float.X(Float.X(class(self).GetRadialInitialAngle()))
}

func (self Instance) SetRadialInitialAngle(value Float.X) {
	class(self).SetRadialInitialAngle(gd.Float(value))
}

func (self Instance) RadialFillDegrees() Float.X {
	return Float.X(Float.X(class(self).GetFillDegrees()))
}

func (self Instance) SetRadialFillDegrees(value Float.X) {
	class(self).SetFillDegrees(gd.Float(value))
}

func (self Instance) RadialCenterOffset() Vector2.XY {
	return Vector2.XY(class(self).GetRadialCenterOffset())
}

func (self Instance) SetRadialCenterOffset(value Vector2.XY) {
	class(self).SetRadialCenterOffset(gd.Vector2(value))
}

func (self Instance) NinePatchStretch() bool {
	return bool(class(self).GetNinePatchStretch())
}

func (self Instance) SetNinePatchStretch(value bool) {
	class(self).SetNinePatchStretch(value)
}

func (self Instance) StretchMarginLeft() int {
	return int(int(class(self).GetStretchMargin(0)))
}

func (self Instance) SetStretchMarginLeft(value int) {
	class(self).SetStretchMargin(0, gd.Int(value))
}

func (self Instance) StretchMarginTop() int {
	return int(int(class(self).GetStretchMargin(1)))
}

func (self Instance) SetStretchMarginTop(value int) {
	class(self).SetStretchMargin(1, gd.Int(value))
}

func (self Instance) StretchMarginRight() int {
	return int(int(class(self).GetStretchMargin(2)))
}

func (self Instance) SetStretchMarginRight(value int) {
	class(self).SetStretchMargin(2, gd.Int(value))
}

func (self Instance) StretchMarginBottom() int {
	return int(int(class(self).GetStretchMargin(3)))
}

func (self Instance) SetStretchMarginBottom(value int) {
	class(self).SetStretchMargin(3, gd.Int(value))
}

func (self Instance) TextureUnder() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetUnderTexture())
}

func (self Instance) SetTextureUnder(value [1]gdclass.Texture2D) {
	class(self).SetUnderTexture(value)
}

func (self Instance) TextureOver() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetOverTexture())
}

func (self Instance) SetTextureOver(value [1]gdclass.Texture2D) {
	class(self).SetOverTexture(value)
}

func (self Instance) TextureProgress() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetProgressTexture())
}

func (self Instance) SetTextureProgress(value [1]gdclass.Texture2D) {
	class(self).SetProgressTexture(value)
}

func (self Instance) TextureProgressOffset() Vector2.XY {
	return Vector2.XY(class(self).GetTextureProgressOffset())
}

func (self Instance) SetTextureProgressOffset(value Vector2.XY) {
	class(self).SetTextureProgressOffset(gd.Vector2(value))
}

func (self Instance) TintUnder() Color.RGBA {
	return Color.RGBA(class(self).GetTintUnder())
}

func (self Instance) SetTintUnder(value Color.RGBA) {
	class(self).SetTintUnder(gd.Color(value))
}

func (self Instance) TintOver() Color.RGBA {
	return Color.RGBA(class(self).GetTintOver())
}

func (self Instance) SetTintOver(value Color.RGBA) {
	class(self).SetTintOver(gd.Color(value))
}

func (self Instance) TintProgress() Color.RGBA {
	return Color.RGBA(class(self).GetTintProgress())
}

func (self Instance) SetTintProgress(value Color.RGBA) {
	class(self).SetTintProgress(gd.Color(value))
}

//go:nosplit
func (self class) SetUnderTexture(tex [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(tex[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_under_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetUnderTexture() [1]gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_under_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProgressTexture(tex [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(tex[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_progress_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetProgressTexture() [1]gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_progress_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOverTexture(tex [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(tex[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_over_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOverTexture() [1]gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_over_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFillMode(mode gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_fill_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFillMode() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_fill_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTintUnder(tint gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, tint)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_tint_under, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTintUnder() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_tint_under, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTintProgress(tint gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, tint)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_tint_progress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTintProgress() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_tint_progress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTintOver(tint gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, tint)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_tint_over, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTintOver() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_tint_over, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureProgressOffset(offset gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_texture_progress_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureProgressOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_texture_progress_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRadialInitialAngle(mode gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_radial_initial_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRadialInitialAngle() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_radial_initial_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRadialCenterOffset(mode gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_radial_center_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRadialCenterOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_radial_center_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFillDegrees(mode gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_fill_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFillDegrees() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_fill_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the stretch margin with the specified index. See [member stretch_margin_bottom] and related properties.
*/
//go:nosplit
func (self class) SetStretchMargin(margin Side, value gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_stretch_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the stretch margin with the specified index. See [member stretch_margin_bottom] and related properties.
*/
//go:nosplit
func (self class) GetStretchMargin(margin Side) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_stretch_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNinePatchStretch(stretch bool) {
	var frame = callframe.New()
	callframe.Arg(frame, stretch)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_nine_patch_stretch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetNinePatchStretch() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_nine_patch_stretch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTextureProgressBar() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTextureProgressBar() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRange() Range.Advanced           { return *((*Range.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRange() Range.Instance        { return *((*Range.Instance)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.Advanced       { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Range.Advanced(self.AsRange()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Range.Instance(self.AsRange()), name)
	}
}
func init() {
	gdclass.Register("TextureProgressBar", func(ptr gd.Object) any {
		return [1]gdclass.TextureProgressBar{*(*gdclass.TextureProgressBar)(unsafe.Pointer(&ptr))}
	})
}

type FillMode = gdclass.TextureProgressBarFillMode

const (
	/*The [member texture_progress] fills from left to right.*/
	FillLeftToRight FillMode = 0
	/*The [member texture_progress] fills from right to left.*/
	FillRightToLeft FillMode = 1
	/*The [member texture_progress] fills from top to bottom.*/
	FillTopToBottom FillMode = 2
	/*The [member texture_progress] fills from bottom to top.*/
	FillBottomToTop FillMode = 3
	/*Turns the node into a radial bar. The [member texture_progress] fills clockwise. See [member radial_center_offset], [member radial_initial_angle] and [member radial_fill_degrees] to control the way the bar fills up.*/
	FillClockwise FillMode = 4
	/*Turns the node into a radial bar. The [member texture_progress] fills counterclockwise. See [member radial_center_offset], [member radial_initial_angle] and [member radial_fill_degrees] to control the way the bar fills up.*/
	FillCounterClockwise FillMode = 5
	/*The [member texture_progress] fills from the center, expanding both towards the left and the right.*/
	FillBilinearLeftAndRight FillMode = 6
	/*The [member texture_progress] fills from the center, expanding both towards the top and the bottom.*/
	FillBilinearTopAndBottom FillMode = 7
	/*Turns the node into a radial bar. The [member texture_progress] fills radially from the center, expanding both clockwise and counterclockwise. See [member radial_center_offset], [member radial_initial_angle] and [member radial_fill_degrees] to control the way the bar fills up.*/
	FillClockwiseAndCounterClockwise FillMode = 8
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
