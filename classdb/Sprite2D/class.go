// Package Sprite2D provides methods for working with Sprite2D object instances.
package Sprite2D

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/Rect2"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Vector2i"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
A node that displays a 2D texture. The texture displayed can be a region from a larger atlas texture, or a frame from a sprite sheet animation.
*/
type Instance [1]gdclass.Sprite2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSprite2D() Instance
}

/*
Returns [code]true[/code], if the pixel at the given position is opaque and [code]false[/code] in other case. The position is in local coordinates.
[b]Note:[/b] It also returns [code]false[/code], if the sprite's texture is [code]null[/code] or if the given position is invalid.
*/
func (self Instance) IsPixelOpaque(pos Vector2.XY) bool { //gd:Sprite2D.is_pixel_opaque
	return bool(class(self).IsPixelOpaque(Vector2.XY(pos)))
}

/*
Returns a [Rect2] representing the Sprite2D's boundary in local coordinates.
[b]Example:[/b] Detect if the Sprite2D was clicked:
[codeblocks]
[gdscript]
func _input(event):

	if event is InputEventMouseButton and event.pressed and event.button_index == MOUSE_BUTTON_LEFT:
	    if get_rect().has_point(to_local(event.position)):
	        print("A click!")

[/gdscript]
[csharp]
public override void _Input(InputEvent @event)

	{
	    if (@event is InputEventMouseButton inputEventMouse)
	    {
	        if (inputEventMouse.Pressed && inputEventMouse.ButtonIndex == MouseButton.Left)
	        {
	            if (GetRect().HasPoint(ToLocal(inputEventMouse.Position)))
	            {
	                GD.Print("A click!");
	            }
	        }
	    }
	}

[/csharp]
[/codeblocks]
*/
func (self Instance) GetRect() Rect2.PositionSize { //gd:Sprite2D.get_rect
	return Rect2.PositionSize(class(self).GetRect())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Sprite2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Sprite2D"))
	casted := Instance{*(*gdclass.Sprite2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Texture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture())
}

func (self Instance) SetTexture(value [1]gdclass.Texture2D) {
	class(self).SetTexture(value)
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
	class(self).SetOffset(Vector2.XY(value))
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

func (self Instance) Hframes() int {
	return int(int(class(self).GetHframes()))
}

func (self Instance) SetHframes(value int) {
	class(self).SetHframes(int64(value))
}

func (self Instance) Vframes() int {
	return int(int(class(self).GetVframes()))
}

func (self Instance) SetVframes(value int) {
	class(self).SetVframes(int64(value))
}

func (self Instance) Frame() int {
	return int(int(class(self).GetFrame()))
}

func (self Instance) SetFrame(value int) {
	class(self).SetFrame(int64(value))
}

func (self Instance) FrameCoords() Vector2i.XY {
	return Vector2i.XY(class(self).GetFrameCoords())
}

func (self Instance) SetFrameCoords(value Vector2i.XY) {
	class(self).SetFrameCoords(Vector2i.XY(value))
}

func (self Instance) RegionEnabled() bool {
	return bool(class(self).IsRegionEnabled())
}

func (self Instance) SetRegionEnabled(value bool) {
	class(self).SetRegionEnabled(value)
}

func (self Instance) RegionRect() Rect2.PositionSize {
	return Rect2.PositionSize(class(self).GetRegionRect())
}

func (self Instance) SetRegionRect(value Rect2.PositionSize) {
	class(self).SetRegionRect(Rect2.PositionSize(value))
}

func (self Instance) RegionFilterClipEnabled() bool {
	return bool(class(self).IsRegionFilterClipEnabled())
}

func (self Instance) SetRegionFilterClipEnabled(value bool) {
	class(self).SetRegionFilterClipEnabled(value)
}

//go:nosplit
func (self class) SetTexture(texture [1]gdclass.Texture2D) { //gd:Sprite2D.set_texture
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() [1]gdclass.Texture2D { //gd:Sprite2D.get_texture
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCentered(centered bool) { //gd:Sprite2D.set_centered
	var frame = callframe.New()
	callframe.Arg(frame, centered)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_set_centered, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCentered() bool { //gd:Sprite2D.is_centered
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_is_centered, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOffset(offset Vector2.XY) { //gd:Sprite2D.set_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOffset() Vector2.XY { //gd:Sprite2D.get_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlipH(flip_h bool) { //gd:Sprite2D.set_flip_h
	var frame = callframe.New()
	callframe.Arg(frame, flip_h)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_set_flip_h, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFlippedH() bool { //gd:Sprite2D.is_flipped_h
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_is_flipped_h, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlipV(flip_v bool) { //gd:Sprite2D.set_flip_v
	var frame = callframe.New()
	callframe.Arg(frame, flip_v)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_set_flip_v, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFlippedV() bool { //gd:Sprite2D.is_flipped_v
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_is_flipped_v, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRegionEnabled(enabled bool) { //gd:Sprite2D.set_region_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_set_region_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsRegionEnabled() bool { //gd:Sprite2D.is_region_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_is_region_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code], if the pixel at the given position is opaque and [code]false[/code] in other case. The position is in local coordinates.
[b]Note:[/b] It also returns [code]false[/code], if the sprite's texture is [code]null[/code] or if the given position is invalid.
*/
//go:nosplit
func (self class) IsPixelOpaque(pos Vector2.XY) bool { //gd:Sprite2D.is_pixel_opaque
	var frame = callframe.New()
	callframe.Arg(frame, pos)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_is_pixel_opaque, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRegionRect(rect Rect2.PositionSize) { //gd:Sprite2D.set_region_rect
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_set_region_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRegionRect() Rect2.PositionSize { //gd:Sprite2D.get_region_rect
	var frame = callframe.New()
	var r_ret = callframe.Ret[Rect2.PositionSize](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_get_region_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRegionFilterClipEnabled(enabled bool) { //gd:Sprite2D.set_region_filter_clip_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_set_region_filter_clip_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsRegionFilterClipEnabled() bool { //gd:Sprite2D.is_region_filter_clip_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_is_region_filter_clip_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrame(frame_ int64) { //gd:Sprite2D.set_frame
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_set_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrame() int64 { //gd:Sprite2D.get_frame
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_get_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrameCoords(coords Vector2i.XY) { //gd:Sprite2D.set_frame_coords
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_set_frame_coords, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrameCoords() Vector2i.XY { //gd:Sprite2D.get_frame_coords
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2i.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_get_frame_coords, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVframes(vframes int64) { //gd:Sprite2D.set_vframes
	var frame = callframe.New()
	callframe.Arg(frame, vframes)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_set_vframes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVframes() int64 { //gd:Sprite2D.get_vframes
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_get_vframes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHframes(hframes int64) { //gd:Sprite2D.set_hframes
	var frame = callframe.New()
	callframe.Arg(frame, hframes)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_set_hframes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHframes() int64 { //gd:Sprite2D.get_hframes
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_get_hframes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a [Rect2] representing the Sprite2D's boundary in local coordinates.
[b]Example:[/b] Detect if the Sprite2D was clicked:
[codeblocks]
[gdscript]
func _input(event):
    if event is InputEventMouseButton and event.pressed and event.button_index == MOUSE_BUTTON_LEFT:
        if get_rect().has_point(to_local(event.position)):
            print("A click!")
[/gdscript]
[csharp]
public override void _Input(InputEvent @event)
{
    if (@event is InputEventMouseButton inputEventMouse)
    {
        if (inputEventMouse.Pressed && inputEventMouse.ButtonIndex == MouseButton.Left)
        {
            if (GetRect().HasPoint(ToLocal(inputEventMouse.Position)))
            {
                GD.Print("A click!");
            }
        }
    }
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GetRect() Rect2.PositionSize { //gd:Sprite2D.get_rect
	var frame = callframe.New()
	var r_ret = callframe.Ret[Rect2.PositionSize](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sprite2D.Bind_get_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnFrameChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("frame_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTextureChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("texture_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsSprite2D() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSprite2D() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.Advanced    { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Node2D.Advanced(self.AsNode2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Instance(self.AsNode2D()), name)
	}
}
func init() {
	gdclass.Register("Sprite2D", func(ptr gd.Object) any { return [1]gdclass.Sprite2D{*(*gdclass.Sprite2D)(unsafe.Pointer(&ptr))} })
}
