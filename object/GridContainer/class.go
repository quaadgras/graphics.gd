package GridContainer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Container"
import "grow.graphics/gd/object/Control"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[GridContainer] arranges its child controls in a grid layout. The number of columns is specified by the [member columns] property, whereas the number of rows depends on how many are needed for the child controls. The number of rows and columns is preserved for every size of the container.
[b]Note:[/b] [GridContainer] only works with child nodes inheriting from [Control]. It won't rearrange child nodes inheriting from [Node2D].

*/
type Simple [1]classdb.GridContainer
func (self Simple) SetColumns(columns int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetColumns(gd.Int(columns))
}
func (self Simple) GetColumns() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetColumns()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.GridContainer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetColumns(columns gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, columns)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridContainer.Bind_set_columns, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColumns() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GridContainer.Bind_get_columns, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsGridContainer() Expert { return self[0].AsGridContainer() }


//go:nosplit
func (self Simple) AsGridContainer() Simple { return self[0].AsGridContainer() }


//go:nosplit
func (self class) AsContainer() Container.Expert { return self[0].AsContainer() }


//go:nosplit
func (self Simple) AsContainer() Container.Simple { return self[0].AsContainer() }


//go:nosplit
func (self class) AsControl() Control.Expert { return self[0].AsControl() }


//go:nosplit
func (self Simple) AsControl() Control.Simple { return self[0].AsControl() }


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
func init() {classdb.Register("GridContainer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
