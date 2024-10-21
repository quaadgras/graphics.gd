package Sprite3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/SpriteBase3D"
import "grow.graphics/gd/object/GeometryInstance3D"
import "grow.graphics/gd/object/VisualInstance3D"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A node that displays a 2D texture in a 3D environment. The texture displayed can be a region from a larger atlas texture, or a frame from a sprite sheet animation. See also [SpriteBase3D] where properties such as the billboard mode are defined.

*/
type Simple [1]classdb.Sprite3D
func (self Simple) SetTexture(texture [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTexture(texture)
}
func (self Simple) GetTexture() [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetTexture(gc))
}
func (self Simple) SetRegionEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRegionEnabled(enabled)
}
func (self Simple) IsRegionEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsRegionEnabled())
}
func (self Simple) SetRegionRect(rect gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRegionRect(rect)
}
func (self Simple) GetRegionRect() gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(Expert(self).GetRegionRect())
}
func (self Simple) SetFrame(frame_ int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFrame(gd.Int(frame_))
}
func (self Simple) GetFrame() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetFrame()))
}
func (self Simple) SetFrameCoords(coords gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFrameCoords(coords)
}
func (self Simple) GetFrameCoords() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetFrameCoords())
}
func (self Simple) SetVframes(vframes int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVframes(gd.Int(vframes))
}
func (self Simple) GetVframes() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetVframes()))
}
func (self Simple) SetHframes(hframes int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHframes(gd.Int(hframes))
}
func (self Simple) GetHframes() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetHframes()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Sprite3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetTexture(texture object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sprite3D.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTexture(ctx gd.Lifetime) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sprite3D.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRegionEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sprite3D.Bind_set_region_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsRegionEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sprite3D.Bind_is_region_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sprite3D.Bind_set_region_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRegionRect() gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sprite3D.Bind_get_region_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFrame(frame_ gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sprite3D.Bind_set_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFrame() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sprite3D.Bind_get_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFrameCoords(coords gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sprite3D.Bind_set_frame_coords, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFrameCoords() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sprite3D.Bind_get_frame_coords, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVframes(vframes gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, vframes)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sprite3D.Bind_set_vframes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVframes() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sprite3D.Bind_get_vframes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHframes(hframes gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hframes)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sprite3D.Bind_set_hframes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHframes() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sprite3D.Bind_get_hframes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsSprite3D() Expert { return self[0].AsSprite3D() }


//go:nosplit
func (self Simple) AsSprite3D() Simple { return self[0].AsSprite3D() }


//go:nosplit
func (self class) AsSpriteBase3D() SpriteBase3D.Expert { return self[0].AsSpriteBase3D() }


//go:nosplit
func (self Simple) AsSpriteBase3D() SpriteBase3D.Simple { return self[0].AsSpriteBase3D() }


//go:nosplit
func (self class) AsGeometryInstance3D() GeometryInstance3D.Expert { return self[0].AsGeometryInstance3D() }


//go:nosplit
func (self Simple) AsGeometryInstance3D() GeometryInstance3D.Simple { return self[0].AsGeometryInstance3D() }


//go:nosplit
func (self class) AsVisualInstance3D() VisualInstance3D.Expert { return self[0].AsVisualInstance3D() }


//go:nosplit
func (self Simple) AsVisualInstance3D() VisualInstance3D.Simple { return self[0].AsVisualInstance3D() }


//go:nosplit
func (self class) AsNode3D() Node3D.Expert { return self[0].AsNode3D() }


//go:nosplit
func (self Simple) AsNode3D() Node3D.Simple { return self[0].AsNode3D() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("Sprite3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
