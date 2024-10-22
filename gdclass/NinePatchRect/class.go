package NinePatchRect

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Also known as 9-slice panels, [NinePatchRect] produces clean panels of any size based on a small texture. To do so, it splits the texture in a 3×3 grid. When you scale the node, it tiles the texture's edges horizontally or vertically, tiles the center on both axes, and leaves the corners unchanged.

*/
type Go [1]classdb.NinePatchRect
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.NinePatchRect
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("NinePatchRect"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Texture() gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Texture2D(class(self).GetTexture(gc))
}

func (self Go) SetTexture(value gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTexture(value)
}

func (self Go) DrawCenter() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDrawCenterEnabled())
}

func (self Go) SetDrawCenter(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawCenter(value)
}

func (self Go) RegionRect() gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Rect2(class(self).GetRegionRect())
}

func (self Go) SetRegionRect(value gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRegionRect(value)
}

func (self Go) PatchMarginLeft() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetPatchMargin(0)))
}

func (self Go) SetPatchMarginLeft(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPatchMargin(0, gd.Int(value))
}

func (self Go) PatchMarginTop() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetPatchMargin(1)))
}

func (self Go) SetPatchMarginTop(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPatchMargin(1, gd.Int(value))
}

func (self Go) PatchMarginRight() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetPatchMargin(2)))
}

func (self Go) SetPatchMarginRight(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPatchMargin(2, gd.Int(value))
}

func (self Go) PatchMarginBottom() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetPatchMargin(3)))
}

func (self Go) SetPatchMarginBottom(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPatchMargin(3, gd.Int(value))
}

func (self Go) AxisStretchHorizontal() classdb.NinePatchRectAxisStretchMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.NinePatchRectAxisStretchMode(class(self).GetHAxisStretchMode())
}

func (self Go) SetAxisStretchHorizontal(value classdb.NinePatchRectAxisStretchMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHAxisStretchMode(value)
}

func (self Go) AxisStretchVertical() classdb.NinePatchRectAxisStretchMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.NinePatchRectAxisStretchMode(class(self).GetVAxisStretchMode())
}

func (self Go) SetAxisStretchVertical(value classdb.NinePatchRectAxisStretchMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVAxisStretchMode(value)
}

//go:nosplit
func (self class) SetTexture(texture gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NinePatchRect.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTexture(ctx gd.Lifetime) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NinePatchRect.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the size of the margin on the specified [enum Side] to [param value] pixels.
*/
//go:nosplit
func (self class) SetPatchMargin(margin gd.Side, value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NinePatchRect.Bind_set_patch_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the size of the margin on the specified [enum Side].
*/
//go:nosplit
func (self class) GetPatchMargin(margin gd.Side) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NinePatchRect.Bind_get_patch_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRegionRect(rect gd.Rect2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NinePatchRect.Bind_set_region_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRegionRect() gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NinePatchRect.Bind_get_region_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDrawCenter(draw_center bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, draw_center)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NinePatchRect.Bind_set_draw_center, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDrawCenterEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NinePatchRect.Bind_is_draw_center_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHAxisStretchMode(mode classdb.NinePatchRectAxisStretchMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NinePatchRect.Bind_set_h_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHAxisStretchMode() classdb.NinePatchRectAxisStretchMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NinePatchRectAxisStretchMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NinePatchRect.Bind_get_h_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVAxisStretchMode(mode classdb.NinePatchRectAxisStretchMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NinePatchRect.Bind_set_v_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVAxisStretchMode() classdb.NinePatchRectAxisStretchMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NinePatchRectAxisStretchMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NinePatchRect.Bind_get_v_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnTextureChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("texture_changed"), gc.Callable(cb), 0)
}


func (self class) AsNinePatchRect() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsNinePatchRect() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsControl(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsControl(), name)
	}
}
func init() {classdb.Register("NinePatchRect", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type AxisStretchMode = classdb.NinePatchRectAxisStretchMode

const (
/*Stretches the center texture across the NinePatchRect. This may cause the texture to be distorted.*/
	AxisStretchModeStretch AxisStretchMode = 0
/*Repeats the center texture across the NinePatchRect. This won't cause any visible distortion. The texture must be seamless for this to work without displaying artifacts between edges.*/
	AxisStretchModeTile AxisStretchMode = 1
/*Repeats the center texture across the NinePatchRect, but will also stretch the texture to make sure each tile is visible in full. This may cause the texture to be distorted, but less than [constant AXIS_STRETCH_MODE_STRETCH]. The texture must be seamless for this to work without displaying artifacts between edges.*/
	AxisStretchModeTileFit AxisStretchMode = 2
)
