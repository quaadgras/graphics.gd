// Package InputEventShortcut provides methods for working with InputEventShortcut object instances.
package InputEventShortcut

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/classdb/InputEvent"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
InputEventShortcut is a special event that can be received in [method Node._input], [method Node._shortcut_input], and [method Node._unhandled_input]. It is typically sent by the editor's Command Palette to trigger actions, but can also be sent manually using [method Viewport.push_input].
*/
type Instance [1]gdclass.InputEventShortcut

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsInputEventShortcut() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.InputEventShortcut

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InputEventShortcut"))
	casted := Instance{*(*gdclass.InputEventShortcut)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Shortcut() [1]gdclass.Shortcut {
	return [1]gdclass.Shortcut(class(self).GetShortcut())
}

func (self Instance) SetShortcut(value [1]gdclass.Shortcut) {
	class(self).SetShortcut(value)
}

//go:nosplit
func (self class) SetShortcut(shortcut [1]gdclass.Shortcut) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shortcut[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventShortcut.Bind_set_shortcut, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetShortcut() [1]gdclass.Shortcut {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventShortcut.Bind_get_shortcut, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Shortcut{gd.PointerWithOwnershipTransferredToGo[gdclass.Shortcut](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsInputEventShortcut() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsInputEventShortcut() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsInputEvent() InputEvent.Advanced {
	return *((*InputEvent.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsInputEvent() InputEvent.Instance {
	return *((*InputEvent.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(InputEvent.Advanced(self.AsInputEvent()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(InputEvent.Instance(self.AsInputEvent()), name)
	}
}
func init() {
	gdclass.Register("InputEventShortcut", func(ptr gd.Object) any {
		return [1]gdclass.InputEventShortcut{*(*gdclass.InputEventShortcut)(unsafe.Pointer(&ptr))}
	})
}
