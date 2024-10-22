package VisibleOnScreenEnabler3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisibleOnScreenNotifier3D"
import "grow.graphics/gd/gdclass/VisualInstance3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[VisibleOnScreenEnabler3D] contains a box-shaped region of 3D space and a target node. The target node will be automatically enabled (via its [member Node.process_mode] property) when any part of this region becomes visible on the screen, and automatically disabled otherwise. This can for example be used to activate enemies only when the player approaches them.
See [VisibleOnScreenNotifier3D] if you only want to be notified when the region is visible on screen.
[b]Note:[/b] [VisibleOnScreenEnabler3D] uses an approximate heuristic that doesn't take walls and other occlusion into account, unless occlusion culling is used. It also won't function unless [member Node3D.visible] is set to [code]true[/code].

*/
type Go [1]classdb.VisibleOnScreenEnabler3D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisibleOnScreenEnabler3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VisibleOnScreenEnabler3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) EnableMode() classdb.VisibleOnScreenEnabler3DEnableMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.VisibleOnScreenEnabler3DEnableMode(class(self).GetEnableMode())
}

func (self Go) SetEnableMode(value classdb.VisibleOnScreenEnabler3DEnableMode) {
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
func (self class) SetEnableMode(mode classdb.VisibleOnScreenEnabler3DEnableMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisibleOnScreenEnabler3D.Bind_set_enable_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnableMode() classdb.VisibleOnScreenEnabler3DEnableMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisibleOnScreenEnabler3DEnableMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisibleOnScreenEnabler3D.Bind_get_enable_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisibleOnScreenEnabler3D.Bind_set_enable_node_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnableNodePath(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisibleOnScreenEnabler3D.Bind_get_enable_node_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsVisibleOnScreenEnabler3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisibleOnScreenEnabler3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisibleOnScreenNotifier3D() VisibleOnScreenNotifier3D.GD { return *((*VisibleOnScreenNotifier3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisibleOnScreenNotifier3D() VisibleOnScreenNotifier3D.Go { return *((*VisibleOnScreenNotifier3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualInstance3D() VisualInstance3D.GD { return *((*VisualInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualInstance3D() VisualInstance3D.Go { return *((*VisualInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisibleOnScreenNotifier3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisibleOnScreenNotifier3D(), name)
	}
}
func init() {classdb.Register("VisibleOnScreenEnabler3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type EnableMode = classdb.VisibleOnScreenEnabler3DEnableMode

const (
/*Corresponds to [constant Node.PROCESS_MODE_INHERIT].*/
	EnableModeInherit EnableMode = 0
/*Corresponds to [constant Node.PROCESS_MODE_ALWAYS].*/
	EnableModeAlways EnableMode = 1
/*Corresponds to [constant Node.PROCESS_MODE_WHEN_PAUSED].*/
	EnableModeWhenPaused EnableMode = 2
)
