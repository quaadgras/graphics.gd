package SubViewport

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Viewport"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[SubViewport] Isolates a rectangular region of a scene to be displayed independently. This can be used, for example, to display UI in 3D space.
[b]Note:[/b] [SubViewport] is a [Viewport] that isn't a [Window], i.e. it doesn't draw anything by itself. To display anything, [SubViewport] must have a non-zero size and be either put inside a [SubViewportContainer] or assigned to a [ViewportTexture].

*/
type Simple [1]classdb.SubViewport
func (self Simple) SetSize(size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSize(size)
}
func (self Simple) GetSize() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetSize())
}
func (self Simple) SetSize2dOverride(size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSize2dOverride(size)
}
func (self Simple) GetSize2dOverride() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetSize2dOverride())
}
func (self Simple) SetSize2dOverrideStretch(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSize2dOverrideStretch(enable)
}
func (self Simple) IsSize2dOverrideStretchEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSize2dOverrideStretchEnabled())
}
func (self Simple) SetUpdateMode(mode classdb.SubViewportUpdateMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUpdateMode(mode)
}
func (self Simple) GetUpdateMode() classdb.SubViewportUpdateMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.SubViewportUpdateMode(Expert(self).GetUpdateMode())
}
func (self Simple) SetClearMode(mode classdb.SubViewportClearMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetClearMode(mode)
}
func (self Simple) GetClearMode() classdb.SubViewportClearMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.SubViewportClearMode(Expert(self).GetClearMode())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.SubViewport
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetSize(size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SubViewport.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSize() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SubViewport.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSize2dOverride(size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SubViewport.Bind_set_size_2d_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSize2dOverride() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SubViewport.Bind_get_size_2d_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSize2dOverrideStretch(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SubViewport.Bind_set_size_2d_override_stretch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSize2dOverrideStretchEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SubViewport.Bind_is_size_2d_override_stretch_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUpdateMode(mode classdb.SubViewportUpdateMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SubViewport.Bind_set_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUpdateMode() classdb.SubViewportUpdateMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.SubViewportUpdateMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SubViewport.Bind_get_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetClearMode(mode classdb.SubViewportClearMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SubViewport.Bind_set_clear_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetClearMode() classdb.SubViewportClearMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.SubViewportClearMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SubViewport.Bind_get_clear_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsSubViewport() Expert { return self[0].AsSubViewport() }


//go:nosplit
func (self Simple) AsSubViewport() Simple { return self[0].AsSubViewport() }


//go:nosplit
func (self class) AsViewport() Viewport.Expert { return self[0].AsViewport() }


//go:nosplit
func (self Simple) AsViewport() Viewport.Simple { return self[0].AsViewport() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("SubViewport", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type ClearMode = classdb.SubViewportClearMode

const (
/*Always clear the render target before drawing.*/
	ClearModeAlways ClearMode = 0
/*Never clear the render target.*/
	ClearModeNever ClearMode = 1
/*Clear the render target on the next frame, then switch to [constant CLEAR_MODE_NEVER].*/
	ClearModeOnce ClearMode = 2
)
type UpdateMode = classdb.SubViewportUpdateMode

const (
/*Do not update the render target.*/
	UpdateDisabled UpdateMode = 0
/*Update the render target once, then switch to [constant UPDATE_DISABLED].*/
	UpdateOnce UpdateMode = 1
/*Update the render target only when it is visible. This is the default value.*/
	UpdateWhenVisible UpdateMode = 2
/*Update the render target only when its parent is visible.*/
	UpdateWhenParentVisible UpdateMode = 3
/*Always update the render target.*/
	UpdateAlways UpdateMode = 4
)
