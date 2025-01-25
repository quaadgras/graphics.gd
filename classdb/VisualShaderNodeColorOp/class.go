// Package VisualShaderNodeColorOp provides methods for working with VisualShaderNodeColorOp object instances.
package VisualShaderNodeColorOp

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

/*
Applies [member operator] to two color inputs.
*/
type Instance [1]gdclass.VisualShaderNodeColorOp

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeColorOp() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeColorOp

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeColorOp"))
	casted := Instance{*(*gdclass.VisualShaderNodeColorOp)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Operator() gdclass.VisualShaderNodeColorOpOperator {
	return gdclass.VisualShaderNodeColorOpOperator(class(self).GetOperator())
}

func (self Instance) SetOperator(value gdclass.VisualShaderNodeColorOpOperator) {
	class(self).SetOperator(value)
}

//go:nosplit
func (self class) SetOperator(op gdclass.VisualShaderNodeColorOpOperator) {
	var frame = callframe.New()
	callframe.Arg(frame, op)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeColorOp.Bind_set_operator, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOperator() gdclass.VisualShaderNodeColorOpOperator {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeColorOpOperator](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeColorOp.Bind_get_operator, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeColorOp() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualShaderNodeColorOp() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(VisualShaderNode.Advanced(self.AsVisualShaderNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNode.Instance(self.AsVisualShaderNode()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeColorOp", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeColorOp{*(*gdclass.VisualShaderNodeColorOp)(unsafe.Pointer(&ptr))}
	})
}

type Operator = gdclass.VisualShaderNodeColorOpOperator

const (
	/*Produce a screen effect with the following formula:
	  [codeblock]
	  result = vec3(1.0) - (vec3(1.0) - a) * (vec3(1.0) - b);
	  [/codeblock]*/
	OpScreen Operator = 0
	/*Produce a difference effect with the following formula:
	  [codeblock]
	  result = abs(a - b);
	  [/codeblock]*/
	OpDifference Operator = 1
	/*Produce a darken effect with the following formula:
	  [codeblock]
	  result = min(a, b);
	  [/codeblock]*/
	OpDarken Operator = 2
	/*Produce a lighten effect with the following formula:
	  [codeblock]
	  result = max(a, b);
	  [/codeblock]*/
	OpLighten Operator = 3
	/*Produce an overlay effect with the following formula:
	  [codeblock]
	  for (int i = 0; i < 3; i++) {
	      float base = a[i];
	      float blend = b[i];
	      if (base < 0.5) {
	          result[i] = 2.0 * base * blend;
	      } else {
	          result[i] = 1.0 - 2.0 * (1.0 - blend) * (1.0 - base);
	      }
	  }
	  [/codeblock]*/
	OpOverlay Operator = 4
	/*Produce a dodge effect with the following formula:
	  [codeblock]
	  result = a / (vec3(1.0) - b);
	  [/codeblock]*/
	OpDodge Operator = 5
	/*Produce a burn effect with the following formula:
	  [codeblock]
	  result = vec3(1.0) - (vec3(1.0) - a) / b;
	  [/codeblock]*/
	OpBurn Operator = 6
	/*Produce a soft light effect with the following formula:
	  [codeblock]
	  for (int i = 0; i < 3; i++) {
	      float base = a[i];
	      float blend = b[i];
	      if (base < 0.5) {
	          result[i] = base * (blend + 0.5);
	      } else {
	          result[i] = 1.0 - (1.0 - base) * (1.0 - (blend - 0.5));
	      }
	  }
	  [/codeblock]*/
	OpSoftLight Operator = 7
	/*Produce a hard light effect with the following formula:
	  [codeblock]
	  for (int i = 0; i < 3; i++) {
	      float base = a[i];
	      float blend = b[i];
	      if (base < 0.5) {
	          result[i] = base * (2.0 * blend);
	      } else {
	          result[i] = 1.0 - (1.0 - base) * (1.0 - 2.0 * (blend - 0.5));
	      }
	  }
	  [/codeblock]*/
	OpHardLight Operator = 8
	/*Represents the size of the [enum Operator] enum.*/
	OpMax Operator = 9
)
