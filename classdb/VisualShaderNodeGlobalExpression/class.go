// Package VisualShaderNodeGlobalExpression provides methods for working with VisualShaderNodeGlobalExpression object instances.
package VisualShaderNodeGlobalExpression

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
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/classdb/VisualShaderNodeExpression"
import "graphics.gd/classdb/VisualShaderNodeGroupBase"
import "graphics.gd/classdb/VisualShaderNodeResizableBase"
import "graphics.gd/classdb/VisualShaderNode"
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
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable

/*
Custom Godot Shader Language expression, which is placed on top of the generated shader. You can place various function definitions inside to call later in [VisualShaderNodeExpression]s (which are injected in the main shader functions). You can also declare varyings, uniforms and global constants.
*/
type Instance [1]gdclass.VisualShaderNodeGlobalExpression

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeGlobalExpression() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeGlobalExpression

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeGlobalExpression"))
	casted := Instance{*(*gdclass.VisualShaderNodeGlobalExpression)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self class) AsVisualShaderNodeGlobalExpression() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeGlobalExpression() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNodeExpression() VisualShaderNodeExpression.Advanced {
	return *((*VisualShaderNodeExpression.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeExpression() VisualShaderNodeExpression.Instance {
	return *((*VisualShaderNodeExpression.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNodeGroupBase() VisualShaderNodeGroupBase.Advanced {
	return *((*VisualShaderNodeGroupBase.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeGroupBase() VisualShaderNodeGroupBase.Instance {
	return *((*VisualShaderNodeGroupBase.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(VisualShaderNodeExpression.Advanced(self.AsVisualShaderNodeExpression()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNodeExpression.Instance(self.AsVisualShaderNodeExpression()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeGlobalExpression", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeGlobalExpression{*(*gdclass.VisualShaderNodeGlobalExpression)(unsafe.Pointer(&ptr))}
	})
}
