package SubViewport

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Viewport"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
[SubViewport] Isolates a rectangular region of a scene to be displayed independently. This can be used, for example, to display UI in 3D space.
[b]Note:[/b] [SubViewport] is a [Viewport] that isn't a [Window], i.e. it doesn't draw anything by itself. To display anything, [SubViewport] must have a non-zero size and be either put inside a [SubViewportContainer] or assigned to a [ViewportTexture].

*/
type Go [1]classdb.SubViewport
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SubViewport
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SubViewport"))
	return Go{classdb.SubViewport(object)}
}

func (self Go) Size() gd.Vector2i {
		return gd.Vector2i(class(self).GetSize())
}

func (self Go) SetSize(value gd.Vector2i) {
	class(self).SetSize(value)
}

func (self Go) Size2dOverride() gd.Vector2i {
		return gd.Vector2i(class(self).GetSize2dOverride())
}

func (self Go) SetSize2dOverride(value gd.Vector2i) {
	class(self).SetSize2dOverride(value)
}

func (self Go) Size2dOverrideStretch() bool {
		return bool(class(self).IsSize2dOverrideStretchEnabled())
}

func (self Go) SetSize2dOverrideStretch(value bool) {
	class(self).SetSize2dOverrideStretch(value)
}

func (self Go) RenderTargetClearMode() classdb.SubViewportClearMode {
		return classdb.SubViewportClearMode(class(self).GetClearMode())
}

func (self Go) SetRenderTargetClearMode(value classdb.SubViewportClearMode) {
	class(self).SetClearMode(value)
}

func (self Go) RenderTargetUpdateMode() classdb.SubViewportUpdateMode {
		return classdb.SubViewportUpdateMode(class(self).GetUpdateMode())
}

func (self Go) SetRenderTargetUpdateMode(value classdb.SubViewportUpdateMode) {
	class(self).SetUpdateMode(value)
}

//go:nosplit
func (self class) SetSize(size gd.Vector2i)  {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SubViewport.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSize() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SubViewport.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSize2dOverride(size gd.Vector2i)  {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SubViewport.Bind_set_size_2d_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSize2dOverride() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SubViewport.Bind_get_size_2d_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSize2dOverrideStretch(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SubViewport.Bind_set_size_2d_override_stretch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSize2dOverrideStretchEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SubViewport.Bind_is_size_2d_override_stretch_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUpdateMode(mode classdb.SubViewportUpdateMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SubViewport.Bind_set_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUpdateMode() classdb.SubViewportUpdateMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.SubViewportUpdateMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SubViewport.Bind_get_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetClearMode(mode classdb.SubViewportClearMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SubViewport.Bind_set_clear_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetClearMode() classdb.SubViewportClearMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.SubViewportClearMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SubViewport.Bind_get_clear_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSubViewport() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSubViewport() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsViewport() Viewport.GD { return *((*Viewport.GD)(unsafe.Pointer(&self))) }
func (self Go) AsViewport() Viewport.Go { return *((*Viewport.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsViewport(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsViewport(), name)
	}
}
func init() {classdb.Register("SubViewport", func(ptr gd.Object) any { return classdb.SubViewport(ptr) })}
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
