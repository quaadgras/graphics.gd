// Package SubViewport provides methods for working with SubViewport object instances.
package SubViewport

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Viewport"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector2i"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
[SubViewport] Isolates a rectangular region of a scene to be displayed independently. This can be used, for example, to display UI in 3D space.
[b]Note:[/b] [SubViewport] is a [Viewport] that isn't a [Window], i.e. it doesn't draw anything by itself. To display anything, [SubViewport] must have a non-zero size and be either put inside a [SubViewportContainer] or assigned to a [ViewportTexture].
*/
type Instance [1]gdclass.SubViewport
type Any interface {
	gd.IsClass
	AsSubViewport() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SubViewport

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SubViewport"))
	return Instance{*(*gdclass.SubViewport)(unsafe.Pointer(&object))}
}

func (self Instance) Size() Vector2i.XY {
	return Vector2i.XY(class(self).GetSize())
}

func (self Instance) SetSize(value Vector2i.XY) {
	class(self).SetSize(gd.Vector2i(value))
}

func (self Instance) Size2dOverride() Vector2i.XY {
	return Vector2i.XY(class(self).GetSize2dOverride())
}

func (self Instance) SetSize2dOverride(value Vector2i.XY) {
	class(self).SetSize2dOverride(gd.Vector2i(value))
}

func (self Instance) Size2dOverrideStretch() bool {
	return bool(class(self).IsSize2dOverrideStretchEnabled())
}

func (self Instance) SetSize2dOverrideStretch(value bool) {
	class(self).SetSize2dOverrideStretch(value)
}

func (self Instance) RenderTargetClearMode() gdclass.SubViewportClearMode {
	return gdclass.SubViewportClearMode(class(self).GetClearMode())
}

func (self Instance) SetRenderTargetClearMode(value gdclass.SubViewportClearMode) {
	class(self).SetClearMode(value)
}

func (self Instance) RenderTargetUpdateMode() gdclass.SubViewportUpdateMode {
	return gdclass.SubViewportUpdateMode(class(self).GetUpdateMode())
}

func (self Instance) SetRenderTargetUpdateMode(value gdclass.SubViewportUpdateMode) {
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
func (self class) SetUpdateMode(mode gdclass.SubViewportUpdateMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SubViewport.Bind_set_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetUpdateMode() gdclass.SubViewportUpdateMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.SubViewportUpdateMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SubViewport.Bind_get_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetClearMode(mode gdclass.SubViewportClearMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SubViewport.Bind_set_clear_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetClearMode() gdclass.SubViewportClearMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.SubViewportClearMode](frame)
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
		return gd.VirtualByName(Viewport.Advanced(self.AsViewport()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Viewport.Instance(self.AsViewport()), name)
	}
}
func init() {
	gdclass.Register("SubViewport", func(ptr gd.Object) any { return [1]gdclass.SubViewport{*(*gdclass.SubViewport)(unsafe.Pointer(&ptr))} })
}

type ClearMode = gdclass.SubViewportClearMode

const (
	/*Always clear the render target before drawing.*/
	ClearModeAlways ClearMode = 0
	/*Never clear the render target.*/
	ClearModeNever ClearMode = 1
	/*Clear the render target on the next frame, then switch to [constant CLEAR_MODE_NEVER].*/
	ClearModeOnce ClearMode = 2
)

type UpdateMode = gdclass.SubViewportUpdateMode

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
