package TextureButton

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/BaseButton"
import "grow.graphics/gd/objects/Control"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
[TextureButton] has the same functionality as [Button], except it uses sprites instead of Godot's [Theme] resource. It is faster to create, but it doesn't support localization like more complex [Control]s.
The "normal" state must contain a texture ([member texture_normal]); other textures are optional.
See also [BaseButton] which contains common properties and methods associated with this node.
*/
type Instance [1]classdb.TextureButton

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.TextureButton

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TextureButton"))
	return Instance{classdb.TextureButton(object)}
}

func (self Instance) TextureNormal() objects.Texture2D {
	return objects.Texture2D(class(self).GetTextureNormal())
}

func (self Instance) SetTextureNormal(value objects.Texture2D) {
	class(self).SetTextureNormal(value)
}

func (self Instance) TexturePressed() objects.Texture2D {
	return objects.Texture2D(class(self).GetTexturePressed())
}

func (self Instance) SetTexturePressed(value objects.Texture2D) {
	class(self).SetTexturePressed(value)
}

func (self Instance) TextureHover() objects.Texture2D {
	return objects.Texture2D(class(self).GetTextureHover())
}

func (self Instance) SetTextureHover(value objects.Texture2D) {
	class(self).SetTextureHover(value)
}

func (self Instance) TextureDisabled() objects.Texture2D {
	return objects.Texture2D(class(self).GetTextureDisabled())
}

func (self Instance) SetTextureDisabled(value objects.Texture2D) {
	class(self).SetTextureDisabled(value)
}

func (self Instance) TextureFocused() objects.Texture2D {
	return objects.Texture2D(class(self).GetTextureFocused())
}

func (self Instance) SetTextureFocused(value objects.Texture2D) {
	class(self).SetTextureFocused(value)
}

func (self Instance) TextureClickMask() objects.BitMap {
	return objects.BitMap(class(self).GetClickMask())
}

func (self Instance) SetTextureClickMask(value objects.BitMap) {
	class(self).SetClickMask(value)
}

func (self Instance) IgnoreTextureSize() bool {
	return bool(class(self).GetIgnoreTextureSize())
}

func (self Instance) SetIgnoreTextureSize(value bool) {
	class(self).SetIgnoreTextureSize(value)
}

func (self Instance) StretchMode() classdb.TextureButtonStretchMode {
	return classdb.TextureButtonStretchMode(class(self).GetStretchMode())
}

func (self Instance) SetStretchMode(value classdb.TextureButtonStretchMode) {
	class(self).SetStretchMode(value)
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

//go:nosplit
func (self class) SetTextureNormal(texture objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_texture_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetTexturePressed(texture objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_texture_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetTextureHover(texture objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_texture_hover, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetTextureDisabled(texture objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_texture_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetTextureFocused(texture objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_texture_focused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetClickMask(mask objects.BitMap) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mask[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_click_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetIgnoreTextureSize(ignore bool) {
	var frame = callframe.New()
	callframe.Arg(frame, ignore)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_ignore_texture_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetStretchMode(mode classdb.TextureButtonStretchMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetFlipH(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_flip_h, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsFlippedH() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_is_flipped_h, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlipV(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_flip_v, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsFlippedV() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_is_flipped_v, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTextureNormal() objects.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_texture_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTexturePressed() objects.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_texture_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTextureHover() objects.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_texture_hover, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTextureDisabled() objects.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_texture_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTextureFocused() objects.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_texture_focused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetClickMask() objects.BitMap {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_click_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.BitMap{classdb.BitMap(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetIgnoreTextureSize() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_ignore_texture_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetStretchMode() classdb.TextureButtonStretchMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextureButtonStretchMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTextureButton() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTextureButton() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsBaseButton() BaseButton.Advanced {
	return *((*BaseButton.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsBaseButton() BaseButton.Instance {
	return *((*BaseButton.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsControl() Control.Advanced { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsBaseButton(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsBaseButton(), name)
	}
}
func init() {
	classdb.Register("TextureButton", func(ptr gd.Object) any { return classdb.TextureButton(ptr) })
}

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
