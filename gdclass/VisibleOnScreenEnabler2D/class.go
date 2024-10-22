package VisibleOnScreenEnabler2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
var _ mmm.Lifetime

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


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VisibleOnScreenEnabler2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) EnableMode() classdb.VisibleOnScreenEnabler2DEnableMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.VisibleOnScreenEnabler2DEnableMode(class(self).GetEnableMode())
}

func (self Go) SetEnableMode(value classdb.VisibleOnScreenEnabler2DEnableMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnableMode(value)
}

func (self Go) EnableNodePath() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetEnableNodePath(gc).String())
}

func (self Go) SetEnableNodePath(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnableNodePath(gc.String(value).NodePath(gc))
}

//go:nosplit
func (self class) SetEnableMode(mode classdb.VisibleOnScreenEnabler2DEnableMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisibleOnScreenEnabler2D.Bind_set_enable_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnableMode() classdb.VisibleOnScreenEnabler2DEnableMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisibleOnScreenEnabler2DEnableMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisibleOnScreenEnabler2D.Bind_get_enable_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnableNodePath(path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisibleOnScreenEnabler2D.Bind_set_enable_node_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnableNodePath(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisibleOnScreenEnabler2D.Bind_get_enable_node_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
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
func init() {classdb.Register("VisibleOnScreenEnabler2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type EnableMode = classdb.VisibleOnScreenEnabler2DEnableMode

const (
/*Corresponds to [constant Node.PROCESS_MODE_INHERIT].*/
	EnableModeInherit EnableMode = 0
/*Corresponds to [constant Node.PROCESS_MODE_ALWAYS].*/
	EnableModeAlways EnableMode = 1
/*Corresponds to [constant Node.PROCESS_MODE_WHEN_PAUSED].*/
	EnableModeWhenPaused EnableMode = 2
)
