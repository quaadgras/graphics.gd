package VisibleOnScreenNotifier3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/VisualInstance3D"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[VisibleOnScreenNotifier3D] represents a box-shaped region of 3D space. When any part of this region becomes visible on screen or in a [Camera3D]'s view, it will emit a [signal screen_entered] signal, and likewise it will emit a [signal screen_exited] signal when no part of it remains visible.
If you want a node to be enabled automatically when this region is visible on screen, use [VisibleOnScreenEnabler3D].
[b]Note:[/b] [VisibleOnScreenNotifier3D] uses an approximate heuristic that doesn't take walls and other occlusion into account, unless occlusion culling is used. It also won't function unless [member Node3D.visible] is set to [code]true[/code].

*/
type Simple [1]classdb.VisibleOnScreenNotifier3D
func (self Simple) SetAabb(rect gd.AABB) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAabb(rect)
}
func (self Simple) IsOnScreen() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsOnScreen())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisibleOnScreenNotifier3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetAabb(rect gd.AABB)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisibleOnScreenNotifier3D.Bind_set_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the bounding box is on the screen.
[b]Note:[/b] It takes one frame for the [VisibleOnScreenNotifier3D]'s visibility to be assessed once added to the scene tree, so this method will always return [code]false[/code] right after it is instantiated.
*/
//go:nosplit
func (self class) IsOnScreen() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisibleOnScreenNotifier3D.Bind_is_on_screen, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisibleOnScreenNotifier3D() Expert { return self[0].AsVisibleOnScreenNotifier3D() }


//go:nosplit
func (self Simple) AsVisibleOnScreenNotifier3D() Simple { return self[0].AsVisibleOnScreenNotifier3D() }


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
func init() {classdb.Register("VisibleOnScreenNotifier3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
