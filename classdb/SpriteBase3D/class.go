// Package SpriteBase3D provides methods for working with SpriteBase3D object instances.
package SpriteBase3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/classdb/GeometryInstance3D"
import "graphics.gd/classdb/VisualInstance3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Rect2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
A node that displays 2D texture information in a 3D environment. See also [Sprite3D] where many other properties are defined.
*/
type Instance [1]gdclass.SpriteBase3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSpriteBase3D() Instance
}

/*
Returns the rectangle representing this sprite.
*/
func (self Instance) GetItemRect() Rect2.PositionSize { //gd:SpriteBase3D.get_item_rect
	return Rect2.PositionSize(class(self).GetItemRect())
}

/*
Returns a [TriangleMesh] with the sprite's vertices following its current configuration (such as its [member axis] and [member pixel_size]).
*/
func (self Instance) GenerateTriangleMesh() [1]gdclass.TriangleMesh { //gd:SpriteBase3D.generate_triangle_mesh
	return [1]gdclass.TriangleMesh(class(self).GenerateTriangleMesh())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SpriteBase3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SpriteBase3D"))
	casted := Instance{*(*gdclass.SpriteBase3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Centered() bool {
	return bool(class(self).IsCentered())
}

func (self Instance) SetCentered(value bool) {
	class(self).SetCentered(value)
}

func (self Instance) Offset() Vector2.XY {
	return Vector2.XY(class(self).GetOffset())
}

func (self Instance) SetOffset(value Vector2.XY) {
	class(self).SetOffset(gd.Vector2(value))
}

func (self Instance) FlipH() bool {
	return bool(class(self).IsFlippedH())
}

func (self Instance) SetFlipH(value bool) {
	class(self).SetFlipH(value)
}

func (self Instance) FlipV() bool {
	return bool(class(self).IsFlippedV())
}

func (self Instance) SetFlipV(value bool) {
	class(self).SetFlipV(value)
}

func (self Instance) Modulate() Color.RGBA {
	return Color.RGBA(class(self).GetModulate())
}

func (self Instance) SetModulate(value Color.RGBA) {
	class(self).SetModulate(gd.Color(value))
}

func (self Instance) PixelSize() Float.X {
	return Float.X(Float.X(class(self).GetPixelSize()))
}

func (self Instance) SetPixelSize(value Float.X) {
	class(self).SetPixelSize(gd.Float(value))
}

func (self Instance) Axis() gd.Vector3Axis {
	return gd.Vector3Axis(class(self).GetAxis())
}

func (self Instance) SetAxis(value gd.Vector3Axis) {
	class(self).SetAxis(value)
}

func (self Instance) Billboard() gdclass.BaseMaterial3DBillboardMode {
	return gdclass.BaseMaterial3DBillboardMode(class(self).GetBillboardMode())
}

func (self Instance) SetBillboard(value gdclass.BaseMaterial3DBillboardMode) {
	class(self).SetBillboardMode(value)
}

func (self Instance) Transparent() bool {
	return bool(class(self).GetDrawFlag(0))
}

func (self Instance) SetTransparent(value bool) {
	class(self).SetDrawFlag(0, value)
}

func (self Instance) Shaded() bool {
	return bool(class(self).GetDrawFlag(1))
}

func (self Instance) SetShaded(value bool) {
	class(self).SetDrawFlag(1, value)
}

func (self Instance) DoubleSided() bool {
	return bool(class(self).GetDrawFlag(2))
}

func (self Instance) SetDoubleSided(value bool) {
	class(self).SetDrawFlag(2, value)
}

func (self Instance) NoDepthTest() bool {
	return bool(class(self).GetDrawFlag(3))
}

func (self Instance) SetNoDepthTest(value bool) {
	class(self).SetDrawFlag(3, value)
}

func (self Instance) FixedSize() bool {
	return bool(class(self).GetDrawFlag(4))
}

func (self Instance) SetFixedSize(value bool) {
	class(self).SetDrawFlag(4, value)
}

func (self Instance) AlphaCut() gdclass.SpriteBase3DAlphaCutMode {
	return gdclass.SpriteBase3DAlphaCutMode(class(self).GetAlphaCutMode())
}

func (self Instance) SetAlphaCut(value gdclass.SpriteBase3DAlphaCutMode) {
	class(self).SetAlphaCutMode(value)
}

func (self Instance) AlphaScissorThreshold() Float.X {
	return Float.X(Float.X(class(self).GetAlphaScissorThreshold()))
}

func (self Instance) SetAlphaScissorThreshold(value Float.X) {
	class(self).SetAlphaScissorThreshold(gd.Float(value))
}

func (self Instance) AlphaHashScale() Float.X {
	return Float.X(Float.X(class(self).GetAlphaHashScale()))
}

func (self Instance) SetAlphaHashScale(value Float.X) {
	class(self).SetAlphaHashScale(gd.Float(value))
}

func (self Instance) AlphaAntialiasingMode() gdclass.BaseMaterial3DAlphaAntiAliasing {
	return gdclass.BaseMaterial3DAlphaAntiAliasing(class(self).GetAlphaAntialiasing())
}

func (self Instance) SetAlphaAntialiasingMode(value gdclass.BaseMaterial3DAlphaAntiAliasing) {
	class(self).SetAlphaAntialiasing(value)
}

func (self Instance) AlphaAntialiasingEdge() Float.X {
	return Float.X(Float.X(class(self).GetAlphaAntialiasingEdge()))
}

func (self Instance) SetAlphaAntialiasingEdge(value Float.X) {
	class(self).SetAlphaAntialiasingEdge(gd.Float(value))
}

func (self Instance) TextureFilter() gdclass.BaseMaterial3DTextureFilter {
	return gdclass.BaseMaterial3DTextureFilter(class(self).GetTextureFilter())
}

func (self Instance) SetTextureFilter(value gdclass.BaseMaterial3DTextureFilter) {
	class(self).SetTextureFilter(value)
}

func (self Instance) RenderPriority() int {
	return int(int(class(self).GetRenderPriority()))
}

func (self Instance) SetRenderPriority(value int) {
	class(self).SetRenderPriority(gd.Int(value))
}

//go:nosplit
func (self class) SetCentered(centered bool) { //gd:SpriteBase3D.set_centered
	var frame = callframe.New()
	callframe.Arg(frame, centered)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_set_centered, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCentered() bool { //gd:SpriteBase3D.is_centered
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_is_centered, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOffset(offset gd.Vector2) { //gd:SpriteBase3D.set_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOffset() gd.Vector2 { //gd:SpriteBase3D.get_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlipH(flip_h bool) { //gd:SpriteBase3D.set_flip_h
	var frame = callframe.New()
	callframe.Arg(frame, flip_h)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_set_flip_h, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFlippedH() bool { //gd:SpriteBase3D.is_flipped_h
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_is_flipped_h, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlipV(flip_v bool) { //gd:SpriteBase3D.set_flip_v
	var frame = callframe.New()
	callframe.Arg(frame, flip_v)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_set_flip_v, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFlippedV() bool { //gd:SpriteBase3D.is_flipped_v
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_is_flipped_v, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetModulate(modulate gd.Color) { //gd:SpriteBase3D.set_modulate
	var frame = callframe.New()
	callframe.Arg(frame, modulate)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_set_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetModulate() gd.Color { //gd:SpriteBase3D.get_modulate
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_get_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRenderPriority(priority gd.Int) { //gd:SpriteBase3D.set_render_priority
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_set_render_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRenderPriority() gd.Int { //gd:SpriteBase3D.get_render_priority
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_get_render_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPixelSize(pixel_size gd.Float) { //gd:SpriteBase3D.set_pixel_size
	var frame = callframe.New()
	callframe.Arg(frame, pixel_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_set_pixel_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPixelSize() gd.Float { //gd:SpriteBase3D.get_pixel_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_get_pixel_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAxis(axis gd.Vector3Axis) { //gd:SpriteBase3D.set_axis
	var frame = callframe.New()
	callframe.Arg(frame, axis)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_set_axis, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAxis() gd.Vector3Axis { //gd:SpriteBase3D.get_axis
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3Axis](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_get_axis, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [code]true[/code], the specified flag will be enabled. See [enum SpriteBase3D.DrawFlags] for a list of flags.
*/
//go:nosplit
func (self class) SetDrawFlag(flag gdclass.SpriteBase3DDrawFlags, enabled bool) { //gd:SpriteBase3D.set_draw_flag
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_set_draw_flag, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the value of the specified flag.
*/
//go:nosplit
func (self class) GetDrawFlag(flag gdclass.SpriteBase3DDrawFlags) bool { //gd:SpriteBase3D.get_draw_flag
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_get_draw_flag, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlphaCutMode(mode gdclass.SpriteBase3DAlphaCutMode) { //gd:SpriteBase3D.set_alpha_cut_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_set_alpha_cut_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlphaCutMode() gdclass.SpriteBase3DAlphaCutMode { //gd:SpriteBase3D.get_alpha_cut_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.SpriteBase3DAlphaCutMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_get_alpha_cut_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlphaScissorThreshold(threshold gd.Float) { //gd:SpriteBase3D.set_alpha_scissor_threshold
	var frame = callframe.New()
	callframe.Arg(frame, threshold)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_set_alpha_scissor_threshold, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlphaScissorThreshold() gd.Float { //gd:SpriteBase3D.get_alpha_scissor_threshold
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_get_alpha_scissor_threshold, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlphaHashScale(threshold gd.Float) { //gd:SpriteBase3D.set_alpha_hash_scale
	var frame = callframe.New()
	callframe.Arg(frame, threshold)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_set_alpha_hash_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlphaHashScale() gd.Float { //gd:SpriteBase3D.get_alpha_hash_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_get_alpha_hash_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlphaAntialiasing(alpha_aa gdclass.BaseMaterial3DAlphaAntiAliasing) { //gd:SpriteBase3D.set_alpha_antialiasing
	var frame = callframe.New()
	callframe.Arg(frame, alpha_aa)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_set_alpha_antialiasing, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlphaAntialiasing() gdclass.BaseMaterial3DAlphaAntiAliasing { //gd:SpriteBase3D.get_alpha_antialiasing
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DAlphaAntiAliasing](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_get_alpha_antialiasing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlphaAntialiasingEdge(edge gd.Float) { //gd:SpriteBase3D.set_alpha_antialiasing_edge
	var frame = callframe.New()
	callframe.Arg(frame, edge)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_set_alpha_antialiasing_edge, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlphaAntialiasingEdge() gd.Float { //gd:SpriteBase3D.get_alpha_antialiasing_edge
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_get_alpha_antialiasing_edge, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBillboardMode(mode gdclass.BaseMaterial3DBillboardMode) { //gd:SpriteBase3D.set_billboard_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_set_billboard_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBillboardMode() gdclass.BaseMaterial3DBillboardMode { //gd:SpriteBase3D.get_billboard_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DBillboardMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_get_billboard_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureFilter(mode gdclass.BaseMaterial3DTextureFilter) { //gd:SpriteBase3D.set_texture_filter
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_set_texture_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureFilter() gdclass.BaseMaterial3DTextureFilter { //gd:SpriteBase3D.get_texture_filter
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DTextureFilter](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_get_texture_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the rectangle representing this sprite.
*/
//go:nosplit
func (self class) GetItemRect() gd.Rect2 { //gd:SpriteBase3D.get_item_rect
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_get_item_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a [TriangleMesh] with the sprite's vertices following its current configuration (such as its [member axis] and [member pixel_size]).
*/
//go:nosplit
func (self class) GenerateTriangleMesh() [1]gdclass.TriangleMesh { //gd:SpriteBase3D.generate_triangle_mesh
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpriteBase3D.Bind_generate_triangle_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TriangleMesh{gd.PointerWithOwnershipTransferredToGo[gdclass.TriangleMesh](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsSpriteBase3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSpriteBase3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsGeometryInstance3D() GeometryInstance3D.Advanced {
	return *((*GeometryInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGeometryInstance3D() GeometryInstance3D.Instance {
	return *((*GeometryInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualInstance3D() VisualInstance3D.Advanced {
	return *((*VisualInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualInstance3D() VisualInstance3D.Instance {
	return *((*VisualInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(GeometryInstance3D.Advanced(self.AsGeometryInstance3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(GeometryInstance3D.Instance(self.AsGeometryInstance3D()), name)
	}
}
func init() {
	gdclass.Register("SpriteBase3D", func(ptr gd.Object) any {
		return [1]gdclass.SpriteBase3D{*(*gdclass.SpriteBase3D)(unsafe.Pointer(&ptr))}
	})
}

type DrawFlags = gdclass.SpriteBase3DDrawFlags //gd:SpriteBase3D.DrawFlags

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

type AlphaCutMode = gdclass.SpriteBase3DAlphaCutMode //gd:SpriteBase3D.AlphaCutMode

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
