// Package VisualShaderNodeVectorDecompose provides methods for working with VisualShaderNodeVectorDecompose object instances.
package VisualShaderNodeVectorDecompose

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/VisualShaderNodeVectorBase"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
Takes a [code]vec2[/code], [code]vec3[/code] or [code]vec4[/code] and decomposes it into scalar values that can be used as separate outputs.
*/
type Instance [1]gdclass.VisualShaderNodeVectorDecompose

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeVectorDecompose() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeVectorDecompose

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeVectorDecompose"))
	casted := Instance{*(*gdclass.VisualShaderNodeVectorDecompose)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self class) AsVisualShaderNodeVectorDecompose() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeVectorDecompose() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNodeVectorBase() VisualShaderNodeVectorBase.Advanced {
	return *((*VisualShaderNodeVectorBase.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeVectorBase() VisualShaderNodeVectorBase.Instance {
	return *((*VisualShaderNodeVectorBase.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(VisualShaderNodeVectorBase.Advanced(self.AsVisualShaderNodeVectorBase()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNodeVectorBase.Instance(self.AsVisualShaderNodeVectorBase()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeVectorDecompose", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeVectorDecompose{*(*gdclass.VisualShaderNodeVectorDecompose)(unsafe.Pointer(&ptr))}
	})
}
