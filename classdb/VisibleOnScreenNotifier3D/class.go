// Package VisibleOnScreenNotifier3D provides methods for working with VisibleOnScreenNotifier3D object instances.
package VisibleOnScreenNotifier3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/VisualInstance3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/AABB"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
[VisibleOnScreenNotifier3D] represents a box-shaped region of 3D space. When any part of this region becomes visible on screen or in a [Camera3D]'s view, it will emit a [signal screen_entered] signal, and likewise it will emit a [signal screen_exited] signal when no part of it remains visible.
If you want a node to be enabled automatically when this region is visible on screen, use [VisibleOnScreenEnabler3D].
[b]Note:[/b] [VisibleOnScreenNotifier3D] uses an approximate heuristic that doesn't take walls and other occlusion into account, unless occlusion culling is used. It also won't function unless [member Node3D.visible] is set to [code]true[/code].
*/
type Instance [1]gdclass.VisibleOnScreenNotifier3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisibleOnScreenNotifier3D() Instance
}

/*
Returns [code]true[/code] if the bounding box is on the screen.
[b]Note:[/b] It takes one frame for the [VisibleOnScreenNotifier3D]'s visibility to be assessed once added to the scene tree, so this method will always return [code]false[/code] right after it is instantiated.
*/
func (self Instance) IsOnScreen() bool {
	return bool(class(self).IsOnScreen())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisibleOnScreenNotifier3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisibleOnScreenNotifier3D"))
	casted := Instance{*(*gdclass.VisibleOnScreenNotifier3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) SetAabb(value AABB.PositionSize) {
	class(self).SetAabb(gd.AABB(value))
}

//go:nosplit
func (self class) SetAabb(rect gd.AABB) {
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisibleOnScreenNotifier3D.Bind_set_aabb, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the bounding box is on the screen.
[b]Note:[/b] It takes one frame for the [VisibleOnScreenNotifier3D]'s visibility to be assessed once added to the scene tree, so this method will always return [code]false[/code] right after it is instantiated.
*/
//go:nosplit
func (self class) IsOnScreen() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisibleOnScreenNotifier3D.Bind_is_on_screen, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnScreenEntered(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("screen_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnScreenExited(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("screen_exited"), gd.NewCallable(cb), 0)
}

func (self class) AsVisibleOnScreenNotifier3D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisibleOnScreenNotifier3D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualInstance3D() VisualInstance3D.Advanced {
	return *((*VisualInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualInstance3D() VisualInstance3D.Instance {
	return *((*VisualInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualInstance3D.Advanced(self.AsVisualInstance3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualInstance3D.Instance(self.AsVisualInstance3D()), name)
	}
}
func init() {
	gdclass.Register("VisibleOnScreenNotifier3D", func(ptr gd.Object) any {
		return [1]gdclass.VisibleOnScreenNotifier3D{*(*gdclass.VisibleOnScreenNotifier3D)(unsafe.Pointer(&ptr))}
	})
}
