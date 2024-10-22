package SpriteBase3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/GeometryInstance3D"
import "grow.graphics/gd/gdclass/VisualInstance3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A node that displays 2D texture information in a 3D environment. See also [Sprite3D] where many other properties are defined.

*/
type Go [1]classdb.SpriteBase3D

/*
Returns the rectangle representing this sprite.
*/
func (self Go) GetItemRect() gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(class(self).GetItemRect())
}

/*
Returns a [TriangleMesh] with the sprite's vertices following its current configuration (such as its [member axis] and [member pixel_size]).
*/
func (self Go) GenerateTriangleMesh() gdclass.TriangleMesh {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.TriangleMesh(class(self).GenerateTriangleMesh(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SpriteBase3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SpriteBase3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Centered() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsCentered())
}

func (self Go) SetCentered(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCentered(value)
}

func (self Go) Offset() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetOffset())
}

func (self Go) SetOffset(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOffset(value)
}

func (self Go) FlipH() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsFlippedH())
}

func (self Go) SetFlipH(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFlipH(value)
}

func (self Go) FlipV() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsFlippedV())
}

func (self Go) SetFlipV(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFlipV(value)
}

func (self Go) Modulate() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetModulate())
}

func (self Go) SetModulate(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetModulate(value)
}

func (self Go) PixelSize() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetPixelSize()))
}

func (self Go) SetPixelSize(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPixelSize(gd.Float(value))
}

func (self Go) Axis() gd.Vector3Axis {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3Axis(class(self).GetAxis())
}

func (self Go) SetAxis(value gd.Vector3Axis) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAxis(value)
}

func (self Go) Billboard() classdb.BaseMaterial3DBillboardMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.BaseMaterial3DBillboardMode(class(self).GetBillboardMode())
}

func (self Go) SetBillboard(value classdb.BaseMaterial3DBillboardMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBillboardMode(value)
}

func (self Go) Transparent() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetDrawFlag(0))
}

func (self Go) SetTransparent(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawFlag(0, value)
}

func (self Go) Shaded() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetDrawFlag(1))
}

func (self Go) SetShaded(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawFlag(1, value)
}

func (self Go) DoubleSided() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetDrawFlag(2))
}

func (self Go) SetDoubleSided(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawFlag(2, value)
}

func (self Go) NoDepthTest() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetDrawFlag(3))
}

func (self Go) SetNoDepthTest(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawFlag(3, value)
}

func (self Go) FixedSize() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetDrawFlag(4))
}

func (self Go) SetFixedSize(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawFlag(4, value)
}

func (self Go) AlphaCut() classdb.SpriteBase3DAlphaCutMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.SpriteBase3DAlphaCutMode(class(self).GetAlphaCutMode())
}

func (self Go) SetAlphaCut(value classdb.SpriteBase3DAlphaCutMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAlphaCutMode(value)
}

func (self Go) AlphaScissorThreshold() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAlphaScissorThreshold()))
}

func (self Go) SetAlphaScissorThreshold(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAlphaScissorThreshold(gd.Float(value))
}

func (self Go) AlphaHashScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAlphaHashScale()))
}

func (self Go) SetAlphaHashScale(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAlphaHashScale(gd.Float(value))
}

func (self Go) AlphaAntialiasingMode() classdb.BaseMaterial3DAlphaAntiAliasing {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.BaseMaterial3DAlphaAntiAliasing(class(self).GetAlphaAntialiasing())
}

func (self Go) SetAlphaAntialiasingMode(value classdb.BaseMaterial3DAlphaAntiAliasing) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAlphaAntialiasing(value)
}

func (self Go) AlphaAntialiasingEdge() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAlphaAntialiasingEdge()))
}

func (self Go) SetAlphaAntialiasingEdge(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAlphaAntialiasingEdge(gd.Float(value))
}

func (self Go) TextureFilter() classdb.BaseMaterial3DTextureFilter {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.BaseMaterial3DTextureFilter(class(self).GetTextureFilter())
}

func (self Go) SetTextureFilter(value classdb.BaseMaterial3DTextureFilter) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextureFilter(value)
}

func (self Go) RenderPriority() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetRenderPriority()))
}

func (self Go) SetRenderPriority(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRenderPriority(gd.Int(value))
}

//go:nosplit
func (self class) SetCentered(centered bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, centered)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_set_centered, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCentered() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_is_centered, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOffset(offset gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOffset() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFlipH(flip_h bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flip_h)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_set_flip_h, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFlippedH() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_is_flipped_h, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFlipV(flip_v bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flip_v)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_set_flip_v, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFlippedV() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_is_flipped_v, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetModulate(modulate gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_set_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetModulate() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_get_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRenderPriority(priority gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_set_render_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRenderPriority() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_get_render_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPixelSize(pixel_size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pixel_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_set_pixel_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPixelSize() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_get_pixel_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAxis(axis gd.Vector3Axis)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, axis)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_set_axis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAxis() gd.Vector3Axis {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3Axis](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_get_axis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [code]true[/code], the specified flag will be enabled. See [enum SpriteBase3D.DrawFlags] for a list of flags.
*/
//go:nosplit
func (self class) SetDrawFlag(flag classdb.SpriteBase3DDrawFlags, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_set_draw_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of the specified flag.
*/
//go:nosplit
func (self class) GetDrawFlag(flag classdb.SpriteBase3DDrawFlags) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_get_draw_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlphaCutMode(mode classdb.SpriteBase3DAlphaCutMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_set_alpha_cut_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlphaCutMode() classdb.SpriteBase3DAlphaCutMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.SpriteBase3DAlphaCutMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_get_alpha_cut_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlphaScissorThreshold(threshold gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, threshold)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_set_alpha_scissor_threshold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlphaScissorThreshold() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_get_alpha_scissor_threshold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlphaHashScale(threshold gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, threshold)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_set_alpha_hash_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlphaHashScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_get_alpha_hash_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlphaAntialiasing(alpha_aa classdb.BaseMaterial3DAlphaAntiAliasing)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, alpha_aa)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_set_alpha_antialiasing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlphaAntialiasing() classdb.BaseMaterial3DAlphaAntiAliasing {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DAlphaAntiAliasing](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_get_alpha_antialiasing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlphaAntialiasingEdge(edge gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, edge)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_set_alpha_antialiasing_edge, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlphaAntialiasingEdge() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_get_alpha_antialiasing_edge, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBillboardMode(mode classdb.BaseMaterial3DBillboardMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_set_billboard_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBillboardMode() classdb.BaseMaterial3DBillboardMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DBillboardMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_get_billboard_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextureFilter(mode classdb.BaseMaterial3DTextureFilter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_set_texture_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextureFilter() classdb.BaseMaterial3DTextureFilter {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DTextureFilter](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_get_texture_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the rectangle representing this sprite.
*/
//go:nosplit
func (self class) GetItemRect() gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_get_item_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a [TriangleMesh] with the sprite's vertices following its current configuration (such as its [member axis] and [member pixel_size]).
*/
//go:nosplit
func (self class) GenerateTriangleMesh(ctx gd.Lifetime) gdclass.TriangleMesh {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpriteBase3D.Bind_generate_triangle_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.TriangleMesh
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsSpriteBase3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSpriteBase3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsGeometryInstance3D() GeometryInstance3D.GD { return *((*GeometryInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsGeometryInstance3D() GeometryInstance3D.Go { return *((*GeometryInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualInstance3D() VisualInstance3D.GD { return *((*VisualInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualInstance3D() VisualInstance3D.Go { return *((*VisualInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsGeometryInstance3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsGeometryInstance3D(), name)
	}
}
func init() {classdb.Register("SpriteBase3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type DrawFlags = classdb.SpriteBase3DDrawFlags

const (
/*If set, the texture's transparency and the opacity are used to make those parts of the sprite invisible.*/
	FlagTransparent DrawFlags = 0
/*If set, lights in the environment affect the sprite.*/
	FlagShaded DrawFlags = 1
/*If set, texture can be seen from the back as well. If not, the texture is invisible when looking at it from behind.*/
	FlagDoubleSided DrawFlags = 2
/*Disables the depth test, so this object is drawn on top of all others. However, objects drawn after it in the draw order may cover it.*/
	FlagDisableDepthTest DrawFlags = 3
/*Label is scaled by depth so that it always appears the same size on screen.*/
	FlagFixedSize DrawFlags = 4
/*Represents the size of the [enum DrawFlags] enum.*/
	FlagMax DrawFlags = 5
)
type AlphaCutMode = classdb.SpriteBase3DAlphaCutMode

const (
/*This mode performs standard alpha blending. It can display translucent areas, but transparency sorting issues may be visible when multiple transparent materials are overlapping.*/
	AlphaCutDisabled AlphaCutMode = 0
/*This mode only allows fully transparent or fully opaque pixels. Harsh edges will be visible unless some form of screen-space antialiasing is enabled (see [member ProjectSettings.rendering/anti_aliasing/quality/screen_space_aa]). On the bright side, this mode doesn't suffer from transparency sorting issues when multiple transparent materials are overlapping. This mode is also known as [i]alpha testing[/i] or [i]1-bit transparency[/i].*/
	AlphaCutDiscard AlphaCutMode = 1
/*This mode draws fully opaque pixels in the depth prepass. This is slower than [constant ALPHA_CUT_DISABLED] or [constant ALPHA_CUT_DISCARD], but it allows displaying translucent areas and smooth edges while using proper sorting.*/
	AlphaCutOpaquePrepass AlphaCutMode = 2
/*This mode draws cuts off all values below a spatially-deterministic threshold, the rest will remain opaque.*/
	AlphaCutHash AlphaCutMode = 3
)
