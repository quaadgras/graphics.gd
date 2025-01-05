// Package VisualShaderNodeComment provides methods for working with VisualShaderNodeComment object instances.
package VisualShaderNodeComment

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/VisualShaderNodeFrame"
import "graphics.gd/classdb/VisualShaderNodeResizableBase"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This node was replaced by [VisualShaderNodeFrame] and only exists to preserve compatibility. In the [VisualShader] editor it behaves exactly like [VisualShaderNodeFrame].
*/
type Instance [1]gdclass.VisualShaderNodeComment
type Any interface {
	gd.IsClass
	AsVisualShaderNodeComment() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeComment

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeComment"))
	return Instance{*(*gdclass.VisualShaderNodeComment)(unsafe.Pointer(&object))}
}

func (self Instance) Description() string {
	return string(class(self).GetDescription().String())
}

func (self Instance) SetDescription(value string) {
	class(self).SetDescription(gd.NewString(value))
}

//go:nosplit
func (self class) SetDescription(description gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(description))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeComment.Bind_set_description, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDescription() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeComment.Bind_get_description, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeComment() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualShaderNodeComment() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNodeFrame() VisualShaderNodeFrame.Advanced {
	return *((*VisualShaderNodeFrame.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeFrame() VisualShaderNodeFrame.Instance {
	return *((*VisualShaderNodeFrame.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNodeResizableBase() VisualShaderNodeResizableBase.Advanced {
	return *((*VisualShaderNodeResizableBase.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeResizableBase() VisualShaderNodeResizableBase.Instance {
	return *((*VisualShaderNodeResizableBase.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNode() VisualShaderNode.Advanced {
	return *((*VisualShaderNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNode() VisualShaderNode.Instance {
	return *((*VisualShaderNode.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(VisualShaderNodeFrame.Advanced(self.AsVisualShaderNodeFrame()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNodeFrame.Instance(self.AsVisualShaderNodeFrame()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeComment", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeComment{*(*gdclass.VisualShaderNodeComment)(unsafe.Pointer(&ptr))}
	})
}
