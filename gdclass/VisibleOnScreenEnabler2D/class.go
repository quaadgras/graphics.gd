package VisibleOnScreenEnabler2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisibleOnScreenNotifier2D"
import "grow.graphics/gd/gdclass/Node2D"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
[VisibleOnScreenEnabler2D] contains a rectangular region of 2D space and a target node. The target node will be automatically enabled (via its [member Node.process_mode] property) when any part of this region becomes visible on the screen, and automatically disabled otherwise. This can for example be used to activate enemies only when the player approaches them.
See [VisibleOnScreenNotifier2D] if you only want to be notified when the region is visible on screen.
[b]Note:[/b] [VisibleOnScreenEnabler2D] uses the render culling code to determine whether it's visible on screen, so it won't function unless [member CanvasItem.visible] is set to [code]true[/code].

*/
type Go [1]classdb.VisibleOnScreenEnabler2D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisibleOnScreenEnabler2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisibleOnScreenEnabler2D"))
	return Go{classdb.VisibleOnScreenEnabler2D(object)}
}

func (self Go) EnableMode() classdb.VisibleOnScreenEnabler2DEnableMode {
		return classdb.VisibleOnScreenEnabler2DEnableMode(class(self).GetEnableMode())
}

func (self Go) SetEnableMode(value classdb.VisibleOnScreenEnabler2DEnableMode) {
	class(self).SetEnableMode(value)
}

func (self Go) EnableNodePath() string {
		return string(class(self).GetEnableNodePath().String())
}

func (self Go) SetEnableNodePath(value string) {
	class(self).SetEnableNodePath(gd.NewString(value).NodePath())
}

//go:nosplit
func (self class) SetEnableMode(mode classdb.VisibleOnScreenEnabler2DEnableMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisibleOnScreenEnabler2D.Bind_set_enable_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnableMode() classdb.VisibleOnScreenEnabler2DEnableMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisibleOnScreenEnabler2DEnableMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisibleOnScreenEnabler2D.Bind_get_enable_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnableNodePath(path gd.NodePath)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisibleOnScreenEnabler2D.Bind_set_enable_node_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnableNodePath() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisibleOnScreenEnabler2D.Bind_get_enable_node_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsVisibleOnScreenEnabler2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisibleOnScreenEnabler2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisibleOnScreenNotifier2D() VisibleOnScreenNotifier2D.GD { return *((*VisibleOnScreenNotifier2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisibleOnScreenNotifier2D() VisibleOnScreenNotifier2D.Go { return *((*VisibleOnScreenNotifier2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.GD { return *((*Node2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode2D() Node2D.Go { return *((*Node2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisibleOnScreenNotifier2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisibleOnScreenNotifier2D(), name)
	}
}
func init() {classdb.Register("VisibleOnScreenEnabler2D", func(ptr gd.Object) any { return classdb.VisibleOnScreenEnabler2D(ptr) })}
type EnableMode = classdb.VisibleOnScreenEnabler2DEnableMode

const (
/*Corresponds to [constant Node.PROCESS_MODE_INHERIT].*/
	EnableModeInherit EnableMode = 0
/*Corresponds to [constant Node.PROCESS_MODE_ALWAYS].*/
	EnableModeAlways EnableMode = 1
/*Corresponds to [constant Node.PROCESS_MODE_WHEN_PAUSED].*/
	EnableModeWhenPaused EnableMode = 2
)
