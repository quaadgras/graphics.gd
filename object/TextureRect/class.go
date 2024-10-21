package TextureRect

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Control"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A control that displays a texture, for example an icon inside a GUI. The texture's placement can be controlled with the [member stretch_mode] property. It can scale, tile, or stay centered inside its bounding rectangle.

*/
type Simple [1]classdb.TextureRect
func (self Simple) SetTexture(texture [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTexture(texture)
}
func (self Simple) GetTexture() [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetTexture(gc))
}
func (self Simple) SetExpandMode(expand_mode classdb.TextureRectExpandMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetExpandMode(expand_mode)
}
func (self Simple) GetExpandMode() classdb.TextureRectExpandMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextureRectExpandMode(Expert(self).GetExpandMode())
}
func (self Simple) SetFlipH(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFlipH(enable)
}
func (self Simple) IsFlippedH() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsFlippedH())
}
func (self Simple) SetFlipV(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFlipV(enable)
}
func (self Simple) IsFlippedV() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsFlippedV())
}
func (self Simple) SetStretchMode(stretch_mode classdb.TextureRectStretchMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStretchMode(stretch_mode)
}
func (self Simple) GetStretchMode() classdb.TextureRectStretchMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextureRectStretchMode(Expert(self).GetStretchMode())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.TextureRect
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetTexture(texture object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureRect.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTexture(ctx gd.Lifetime) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureRect.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetExpandMode(expand_mode classdb.TextureRectExpandMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, expand_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureRect.Bind_set_expand_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetExpandMode() classdb.TextureRectExpandMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextureRectExpandMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureRect.Bind_get_expand_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFlipH(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureRect.Bind_set_flip_h, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFlippedH() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureRect.Bind_is_flipped_h, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureRect.Bind_set_flip_v, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFlippedV() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureRect.Bind_is_flipped_v, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStretchMode(stretch_mode classdb.TextureRectStretchMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stretch_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureRect.Bind_set_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStretchMode() classdb.TextureRectStretchMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextureRectStretchMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureRect.Bind_get_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsTextureRect() Expert { return self[0].AsTextureRect() }


//go:nosplit
func (self Simple) AsTextureRect() Simple { return self[0].AsTextureRect() }


//go:nosplit
func (self class) AsControl() Control.Expert { return self[0].AsControl() }


//go:nosplit
func (self Simple) AsControl() Control.Simple { return self[0].AsControl() }


//go:nosplit
func (self class) AsCanvasItem() CanvasItem.Expert { return self[0].AsCanvasItem() }


//go:nosplit
func (self Simple) AsCanvasItem() CanvasItem.Simple { return self[0].AsCanvasItem() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("TextureRect", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type ExpandMode = classdb.TextureRectExpandMode

const (
/*The minimum size will be equal to texture size, i.e. [TextureRect] can't be smaller than the texture.*/
	ExpandKeepSize ExpandMode = 0
/*The size of the texture won't be considered for minimum size calculation, so the [TextureRect] can be shrunk down past the texture size.*/
	ExpandIgnoreSize ExpandMode = 1
/*The height of the texture will be ignored. Minimum width will be equal to the current height. Useful for horizontal layouts, e.g. inside [HBoxContainer].*/
	ExpandFitWidth ExpandMode = 2
/*Same as [constant EXPAND_FIT_WIDTH], but keeps texture's aspect ratio.*/
	ExpandFitWidthProportional ExpandMode = 3
/*The width of the texture will be ignored. Minimum height will be equal to the current width. Useful for vertical layouts, e.g. inside [VBoxContainer].*/
	ExpandFitHeight ExpandMode = 4
/*Same as [constant EXPAND_FIT_HEIGHT], but keeps texture's aspect ratio.*/
	ExpandFitHeightProportional ExpandMode = 5
)
type StretchMode = classdb.TextureRectStretchMode

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
/*Scale the texture to fit the node's bounding rectangle, center it and maintain its aspect ratio.*/
	StretchKeepAspectCentered StretchMode = 5
/*Scale the texture so that the shorter side fits the bounding rectangle. The other side clips to the node's limits.*/
	StretchKeepAspectCovered StretchMode = 6
)
