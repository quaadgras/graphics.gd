// Package VisualShaderNodeTexture3D provides methods for working with VisualShaderNodeTexture3D object instances.
package VisualShaderNodeTexture3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/VisualShaderNodeSample3D"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Performs a lookup operation on the provided texture, with support for multiple texture sources to choose from.
*/
type Instance [1]gdclass.VisualShaderNodeTexture3D
type Any interface {
	gd.IsClass
	AsVisualShaderNodeTexture3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeTexture3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeTexture3D"))
	return Instance{*(*gdclass.VisualShaderNodeTexture3D)(unsafe.Pointer(&object))}
}

func (self Instance) Texture() [1]gdclass.Texture3D {
	return [1]gdclass.Texture3D(class(self).GetTexture())
}

func (self Instance) SetTexture(value [1]gdclass.Texture3D) {
	class(self).SetTexture(value)
}

//go:nosplit
func (self class) SetTexture(value [1]gdclass.Texture3D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(value[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTexture3D.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() [1]gdclass.Texture3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTexture3D.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Texture3D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture3D](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeTexture3D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeTexture3D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNodeSample3D() VisualShaderNodeSample3D.Advanced {
	return *((*VisualShaderNodeSample3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeSample3D() VisualShaderNodeSample3D.Instance {
	return *((*VisualShaderNodeSample3D.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(VisualShaderNodeSample3D.Advanced(self.AsVisualShaderNodeSample3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNodeSample3D.Instance(self.AsVisualShaderNodeSample3D()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeTexture3D", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeTexture3D{*(*gdclass.VisualShaderNodeTexture3D)(unsafe.Pointer(&ptr))}
	})
}
