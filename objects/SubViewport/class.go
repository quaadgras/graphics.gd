package SubViewport

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Viewport"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
[SubViewport] Isolates a rectangular region of a scene to be displayed independently. This can be used, for example, to display UI in 3D space.
[b]Note:[/b] [SubViewport] is a [Viewport] that isn't a [Window], i.e. it doesn't draw anything by itself. To display anything, [SubViewport] must have a non-zero size and be either put inside a [SubViewportContainer] or assigned to a [ViewportTexture].
*/
type Instance [1]classdb.SubViewport

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.SubViewport

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SubViewport"))
	return Instance{classdb.SubViewport(object)}
}

func (self Instance) Size() gd.Vector2i {
	return gd.Vector2i(class(self).GetSize())
}

func (self Instance) SetSize(value gd.Vector2i) {
	class(self).SetSize(value)
}

func (self Instance) Size2dOverride() gd.Vector2i {
	return gd.Vector2i(class(self).GetSize2dOverride())
}

func (self Instance) SetSize2dOverride(value gd.Vector2i) {
	class(self).SetSize2dOverride(value)
}

func (self Instance) Size2dOverrideStretch() bool {
	return bool(class(self).IsSize2dOverrideStretchEnabled())
}

func (self Instance) SetSize2dOverrideStretch(value bool) {
	class(self).SetSize2dOverrideStretch(value)
}

func (self Instance) RenderTargetClearMode() classdb.SubViewportClearMode {
	return classdb.SubViewportClearMode(class(self).GetClearMode())
}

func (self Instance) SetRenderTargetClearMode(value classdb.SubViewportClearMode) {
	class(self).SetClearMode(value)
}

func (self Instance) RenderTargetUpdateMode() classdb.SubViewportUpdateMode {
	return classdb.SubViewportUpdateMode(class(self).GetUpdateMode())
}

func (self Instance) SetRenderTargetUpdateMode(value classdb.SubViewportUpdateMode) {
	class(self).SetUpdateMode(value)
}

//go:nosplit
func (self class) SetSize(size gd.Vector2i) {
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
func (self class) SetSize2dOverride(size gd.Vector2i) {
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
func (self class) SetSize2dOverrideStretch(enable bool) {
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
func (self class) SetUpdateMode(mode classdb.SubViewportUpdateMode) {
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
func (self class) SetClearMode(mode classdb.SubViewportClearMode) {
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
func (self class) AsSubViewport() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSubViewport() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsViewport() Viewport.Advanced {
	return *((*Viewport.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsViewport() Viewport.Instance {
	return *((*Viewport.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsViewport(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsViewport(), name)
	}
}
func init() {
	classdb.Register("SubViewport", func(ptr gd.Object) any { return classdb.SubViewport(ptr) })
}

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
