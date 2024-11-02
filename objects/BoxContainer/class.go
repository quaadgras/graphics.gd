package BoxContainer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Container"
import "grow.graphics/gd/objects/Control"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A container that arranges its child controls horizontally or vertically, rearranging them automatically when their minimum size changes.
*/
type Instance [1]classdb.BoxContainer

/*
Adds a [Control] node to the box as a spacer. If [param begin] is [code]true[/code], it will insert the [Control] node in front of all other children.
*/
func (self Instance) AddSpacer(begin bool) objects.Control {
	return objects.Control(class(self).AddSpacer(begin))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.BoxContainer

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("BoxContainer"))
	return Instance{classdb.BoxContainer(object)}
}

func (self Instance) Alignment() classdb.BoxContainerAlignmentMode {
	return classdb.BoxContainerAlignmentMode(class(self).GetAlignment())
}

func (self Instance) SetAlignment(value classdb.BoxContainerAlignmentMode) {
	class(self).SetAlignment(value)
}

func (self Instance) Vertical() bool {
	return bool(class(self).IsVertical())
}

func (self Instance) SetVertical(value bool) {
	class(self).SetVertical(value)
}

/*
Adds a [Control] node to the box as a spacer. If [param begin] is [code]true[/code], it will insert the [Control] node in front of all other children.
*/
//go:nosplit
func (self class) AddSpacer(begin bool) objects.Control {
	var frame = callframe.New()
	callframe.Arg(frame, begin)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoxContainer.Bind_add_spacer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Control{classdb.Control(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlignment(alignment classdb.BoxContainerAlignmentMode) {
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoxContainer.Bind_set_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlignment() classdb.BoxContainerAlignmentMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BoxContainerAlignmentMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoxContainer.Bind_get_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVertical(vertical bool) {
	var frame = callframe.New()
	callframe.Arg(frame, vertical)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoxContainer.Bind_set_vertical, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsVertical() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoxContainer.Bind_is_vertical, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsBoxContainer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsBoxContainer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsContainer() Container.Advanced {
	return *((*Container.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsContainer() Container.Instance {
	return *((*Container.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsControl() Control.Advanced { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(self.AsContainer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsContainer(), name)
	}
}
func init() {
	classdb.Register("BoxContainer", func(ptr gd.Object) any { return classdb.BoxContainer(ptr) })
}

type AlignmentMode = classdb.BoxContainerAlignmentMode

const (
	/*The child controls will be arranged at the beginning of the container, i.e. top if orientation is vertical, left if orientation is horizontal (right for RTL layout).*/
	AlignmentBegin AlignmentMode = 0
	/*The child controls will be centered in the container.*/
	AlignmentCenter AlignmentMode = 1
	/*The child controls will be arranged at the end of the container, i.e. bottom if orientation is vertical, right if orientation is horizontal (left for RTL layout).*/
	AlignmentEnd AlignmentMode = 2
)
