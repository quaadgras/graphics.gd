package VisibleOnScreenEnabler2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/VisibleOnScreenNotifier2D"
import "grow.graphics/gd/objects/Node2D"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
[VisibleOnScreenEnabler2D] contains a rectangular region of 2D space and a target node. The target node will be automatically enabled (via its [member Node.process_mode] property) when any part of this region becomes visible on the screen, and automatically disabled otherwise. This can for example be used to activate enemies only when the player approaches them.
See [VisibleOnScreenNotifier2D] if you only want to be notified when the region is visible on screen.
[b]Note:[/b] [VisibleOnScreenEnabler2D] uses the render culling code to determine whether it's visible on screen, so it won't function unless [member CanvasItem.visible] is set to [code]true[/code].
*/
type Instance [1]classdb.VisibleOnScreenEnabler2D

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VisibleOnScreenEnabler2D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisibleOnScreenEnabler2D"))
	return Instance{classdb.VisibleOnScreenEnabler2D(object)}
}

func (self Instance) EnableMode() classdb.VisibleOnScreenEnabler2DEnableMode {
	return classdb.VisibleOnScreenEnabler2DEnableMode(class(self).GetEnableMode())
}

func (self Instance) SetEnableMode(value classdb.VisibleOnScreenEnabler2DEnableMode) {
	class(self).SetEnableMode(value)
}

func (self Instance) EnableNodePath() string {
	return string(class(self).GetEnableNodePath().String())
}

func (self Instance) SetEnableNodePath(value string) {
	class(self).SetEnableNodePath(gd.NewString(value).NodePath())
}

//go:nosplit
func (self class) SetEnableMode(mode classdb.VisibleOnScreenEnabler2DEnableMode) {
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
func (self class) SetEnableNodePath(path gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisibleOnScreenEnabler2D.Bind_set_enable_node_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableNodePath() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisibleOnScreenEnabler2D.Bind_get_enable_node_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsVisibleOnScreenEnabler2D() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisibleOnScreenEnabler2D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisibleOnScreenNotifier2D() VisibleOnScreenNotifier2D.Advanced {
	return *((*VisibleOnScreenNotifier2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisibleOnScreenNotifier2D() VisibleOnScreenNotifier2D.Instance {
	return *((*VisibleOnScreenNotifier2D.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(self.AsVisibleOnScreenNotifier2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisibleOnScreenNotifier2D(), name)
	}
}
func init() {
	classdb.Register("VisibleOnScreenEnabler2D", func(ptr gd.Object) any { return classdb.VisibleOnScreenEnabler2D(ptr) })
}

type EnableMode = classdb.VisibleOnScreenEnabler2DEnableMode

const (
	/*Corresponds to [constant Node.PROCESS_MODE_INHERIT].*/
	EnableModeInherit EnableMode = 0
	/*Corresponds to [constant Node.PROCESS_MODE_ALWAYS].*/
	EnableModeAlways EnableMode = 1
	/*Corresponds to [constant Node.PROCESS_MODE_WHEN_PAUSED].*/
	EnableModeWhenPaused EnableMode = 2
)
