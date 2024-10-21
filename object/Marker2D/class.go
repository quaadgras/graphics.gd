package Marker2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Node2D"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Generic 2D position hint for editing. It's just like a plain [Node2D], but it displays as a cross in the 2D editor at all times. You can set the cross' visual size by using the gizmo in the 2D editor while the node is selected.

*/
type Simple [1]classdb.Marker2D
func (self Simple) SetGizmoExtents(extents float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGizmoExtents(gd.Float(extents))
}
func (self Simple) GetGizmoExtents() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetGizmoExtents()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Marker2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetGizmoExtents(extents gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, extents)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Marker2D.Bind_set_gizmo_extents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGizmoExtents() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Marker2D.Bind_get_gizmo_extents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsMarker2D() Expert { return self[0].AsMarker2D() }


//go:nosplit
func (self Simple) AsMarker2D() Simple { return self[0].AsMarker2D() }


//go:nosplit
func (self class) AsNode2D() Node2D.Expert { return self[0].AsNode2D() }


//go:nosplit
func (self Simple) AsNode2D() Node2D.Simple { return self[0].AsNode2D() }


//go:nosplit
func (self class) AsCanvasItem() CanvasItem.Expert { return self[0].AsCanvasItem() }


//go:nosplit
func (self Simple) AsCanvasItem() CanvasItem.Simple { return self[0].AsCanvasItem() }


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
func init() {classdb.Register("Marker2D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
