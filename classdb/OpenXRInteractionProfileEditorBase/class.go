// Package OpenXRInteractionProfileEditorBase provides methods for working with OpenXRInteractionProfileEditorBase object instances.
package OpenXRInteractionProfileEditorBase

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/BoxContainer"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Container"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/HBoxContainer"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
This is a base class for interaction profile editors used by the OpenXR action map editor. It can be used to create bespoke editors for specific interaction profiles.
*/
type Instance [1]gdclass.OpenXRInteractionProfileEditorBase

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsOpenXRInteractionProfileEditorBase() Instance
}

/*
Setup this editor for the provided [param action_map] and [param interaction_profile].
*/
func (self Instance) Setup(action_map [1]gdclass.OpenXRActionMap, interaction_profile [1]gdclass.OpenXRInteractionProfile) { //gd:OpenXRInteractionProfileEditorBase.setup
	Advanced(self).Setup(action_map, interaction_profile)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.OpenXRInteractionProfileEditorBase

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRInteractionProfileEditorBase"))
	casted := Instance{*(*gdclass.OpenXRInteractionProfileEditorBase)(unsafe.Pointer(&object))}
	return casted
}

/*
Setup this editor for the provided [param action_map] and [param interaction_profile].
*/
//go:nosplit
func (self class) Setup(action_map [1]gdclass.OpenXRActionMap, interaction_profile [1]gdclass.OpenXRInteractionProfile) { //gd:OpenXRInteractionProfileEditorBase.setup
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action_map[0])[0])
	callframe.Arg(frame, pointers.Get(interaction_profile[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInteractionProfileEditorBase.Bind_setup, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsOpenXRInteractionProfileEditorBase() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsOpenXRInteractionProfileEditorBase() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsHBoxContainer() HBoxContainer.Advanced {
	return *((*HBoxContainer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsHBoxContainer() HBoxContainer.Instance {
	return *((*HBoxContainer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsBoxContainer() BoxContainer.Advanced {
	return *((*BoxContainer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsBoxContainer() BoxContainer.Instance {
	return *((*BoxContainer.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(HBoxContainer.Advanced(self.AsHBoxContainer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(HBoxContainer.Instance(self.AsHBoxContainer()), name)
	}
}
func init() {
	gdclass.Register("OpenXRInteractionProfileEditorBase", func(ptr gd.Object) any {
		return [1]gdclass.OpenXRInteractionProfileEditorBase{*(*gdclass.OpenXRInteractionProfileEditorBase)(unsafe.Pointer(&ptr))}
	})
}
