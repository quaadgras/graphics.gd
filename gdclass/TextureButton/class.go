package TextureButton

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/BaseButton"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[TextureButton] has the same functionality as [Button], except it uses sprites instead of Godot's [Theme] resource. It is faster to create, but it doesn't support localization like more complex [Control]s.
The "normal" state must contain a texture ([member texture_normal]); other textures are optional.
See also [BaseButton] which contains common properties and methods associated with this node.

*/
type Go [1]classdb.TextureButton
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.TextureButton
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("TextureButton"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) TextureNormal() gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Texture2D(class(self).GetTextureNormal(gc))
}

func (self Go) SetTextureNormal(value gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextureNormal(value)
}

func (self Go) TexturePressed() gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Texture2D(class(self).GetTexturePressed(gc))
}

func (self Go) SetTexturePressed(value gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTexturePressed(value)
}

func (self Go) TextureHover() gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Texture2D(class(self).GetTextureHover(gc))
}

func (self Go) SetTextureHover(value gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextureHover(value)
}

func (self Go) TextureDisabled() gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Texture2D(class(self).GetTextureDisabled(gc))
}

func (self Go) SetTextureDisabled(value gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextureDisabled(value)
}

func (self Go) TextureFocused() gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Texture2D(class(self).GetTextureFocused(gc))
}

func (self Go) SetTextureFocused(value gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextureFocused(value)
}

func (self Go) TextureClickMask() gdclass.BitMap {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.BitMap(class(self).GetClickMask(gc))
}

func (self Go) SetTextureClickMask(value gdclass.BitMap) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetClickMask(value)
}

func (self Go) IgnoreTextureSize() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetIgnoreTextureSize())
}

func (self Go) SetIgnoreTextureSize(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIgnoreTextureSize(value)
}

func (self Go) StretchMode() classdb.TextureButtonStretchMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextureButtonStretchMode(class(self).GetStretchMode())
}

func (self Go) SetStretchMode(value classdb.TextureButtonStretchMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStretchMode(value)
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

//go:nosplit
func (self class) SetTextureNormal(texture gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureButton.Bind_set_texture_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetTexturePressed(texture gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureButton.Bind_set_texture_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetTextureHover(texture gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureButton.Bind_set_texture_hover, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetTextureDisabled(texture gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureButton.Bind_set_texture_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetTextureFocused(texture gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureButton.Bind_set_texture_focused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetClickMask(mask gdclass.BitMap)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(mask[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureButton.Bind_set_click_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetIgnoreTextureSize(ignore bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ignore)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureButton.Bind_set_ignore_texture_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetStretchMode(mode classdb.TextureButtonStretchMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureButton.Bind_set_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFlipH(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureButton.Bind_set_flip_h, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFlippedH() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureButton.Bind_is_flipped_h, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFlipV(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureButton.Bind_set_flip_v, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFlippedV() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureButton.Bind_is_flipped_v, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetTextureNormal(ctx gd.Lifetime) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureButton.Bind_get_texture_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetTexturePressed(ctx gd.Lifetime) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureButton.Bind_get_texture_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetTextureHover(ctx gd.Lifetime) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureButton.Bind_get_texture_hover, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetTextureDisabled(ctx gd.Lifetime) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureButton.Bind_get_texture_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetTextureFocused(ctx gd.Lifetime) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureButton.Bind_get_texture_focused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetClickMask(ctx gd.Lifetime) gdclass.BitMap {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureButton.Bind_get_click_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.BitMap
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetIgnoreTextureSize() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureButton.Bind_get_ignore_texture_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetStretchMode() classdb.TextureButtonStretchMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextureButtonStretchMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureButton.Bind_get_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTextureButton() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTextureButton() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsBaseButton() BaseButton.GD { return *((*BaseButton.GD)(unsafe.Pointer(&self))) }
func (self Go) AsBaseButton() BaseButton.Go { return *((*BaseButton.Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsBaseButton(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsBaseButton(), name)
	}
}
func init() {classdb.Register("TextureButton", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type StretchMode = classdb.TextureButtonStretchMode

const (
/*Scale to fit the node's bounding rectangle.*/
	StretchScale StretchMode = 0
/*Tile inside the node's bounding rectangle.*/
	StretchTile StretchMode = 1
/*The texture keeps its original size and stays in the bounding rectangle's top-left corner.*/
	StretchKeep StretchMode = 2
/*The texture keeps its original size and stays centered in the node's bounding rectangle.*/
	StretchKeepCentered StretchMode = 3
/*Scale the texture to fit the node's bounding rectangle, but maintain the texture's aspect ratio.*/
	StretchKeepAspect StretchMode = 4
/*Scale the texture to fit the node's bounding rectangle, center it, and maintain its aspect ratio.*/
	StretchKeepAspectCentered StretchMode = 5
/*Scale the texture so that the shorter side fits the bounding rectangle. The other side clips to the node's limits.*/
	StretchKeepAspectCovered StretchMode = 6
)
