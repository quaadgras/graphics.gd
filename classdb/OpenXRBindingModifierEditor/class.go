// Package OpenXRBindingModifierEditor provides methods for working with OpenXRBindingModifierEditor object instances.
package OpenXRBindingModifierEditor

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Container"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/PanelContainer"
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
This is the default binding modifier editor used in the OpenXR action map.
*/
type Instance [1]gdclass.OpenXRBindingModifierEditor

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsOpenXRBindingModifierEditor() Instance
}

/*
Returns the [OpenXRBindingModifier] currently being edited.
*/
func (self Instance) GetBindingModifier() [1]gdclass.OpenXRBindingModifier { //gd:OpenXRBindingModifierEditor.get_binding_modifier
	return [1]gdclass.OpenXRBindingModifier(Advanced(self).GetBindingModifier())
}

/*
Setup this editor for the provided [param action_map] and [param binding_modifier].
*/
func (self Instance) Setup(action_map [1]gdclass.OpenXRActionMap, binding_modifier [1]gdclass.OpenXRBindingModifier) { //gd:OpenXRBindingModifierEditor.setup
	Advanced(self).Setup(action_map, binding_modifier)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.OpenXRBindingModifierEditor

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRBindingModifierEditor"))
	casted := Instance{*(*gdclass.OpenXRBindingModifierEditor)(unsafe.Pointer(&object))}
	return casted
}

/*
Returns the [OpenXRBindingModifier] currently being edited.
*/
//go:nosplit
func (self class) GetBindingModifier() [1]gdclass.OpenXRBindingModifier { //gd:OpenXRBindingModifierEditor.get_binding_modifier
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRBindingModifierEditor.Bind_get_binding_modifier, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.OpenXRBindingModifier{gd.PointerWithOwnershipTransferredToGo[gdclass.OpenXRBindingModifier](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Setup this editor for the provided [param action_map] and [param binding_modifier].
*/
//go:nosplit
func (self class) Setup(action_map [1]gdclass.OpenXRActionMap, binding_modifier [1]gdclass.OpenXRBindingModifier) { //gd:OpenXRBindingModifierEditor.setup
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action_map[0])[0])
	callframe.Arg(frame, pointers.Get(binding_modifier[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRBindingModifierEditor.Bind_setup, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnBindingModifierRemoved(cb func(binding_modifier_editor Object.Instance)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("binding_modifier_removed"), gd.NewCallable(cb), 0)
}

func (self class) AsOpenXRBindingModifierEditor() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsOpenXRBindingModifierEditor() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsPanelContainer() PanelContainer.Advanced {
	return *((*PanelContainer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPanelContainer() PanelContainer.Instance {
	return *((*PanelContainer.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(PanelContainer.Advanced(self.AsPanelContainer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(PanelContainer.Instance(self.AsPanelContainer()), name)
	}
}
func init() {
	gdclass.Register("OpenXRBindingModifierEditor", func(ptr gd.Object) any {
		return [1]gdclass.OpenXRBindingModifierEditor{*(*gdclass.OpenXRBindingModifierEditor)(unsafe.Pointer(&ptr))}
	})
}
