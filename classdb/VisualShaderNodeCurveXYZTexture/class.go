// Package VisualShaderNodeCurveXYZTexture provides methods for working with VisualShaderNodeCurveXYZTexture object instances.
package VisualShaderNodeCurveXYZTexture

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

/*
Comes with a built-in editor for texture's curves.
*/
type Instance [1]gdclass.VisualShaderNodeCurveXYZTexture

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeCurveXYZTexture() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeCurveXYZTexture

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeCurveXYZTexture"))
	casted := Instance{*(*gdclass.VisualShaderNodeCurveXYZTexture)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Texture() [1]gdclass.CurveXYZTexture {
	return [1]gdclass.CurveXYZTexture(class(self).GetTexture())
}

func (self Instance) SetTexture(value [1]gdclass.CurveXYZTexture) {
	class(self).SetTexture(value)
}

//go:nosplit
func (self class) SetTexture(texture [1]gdclass.CurveXYZTexture) { //gd:VisualShaderNodeCurveXYZTexture.set_texture
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCurveXYZTexture.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() [1]gdclass.CurveXYZTexture { //gd:VisualShaderNodeCurveXYZTexture.get_texture
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCurveXYZTexture.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.CurveXYZTexture{gd.PointerWithOwnershipTransferredToGo[gdclass.CurveXYZTexture](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeCurveXYZTexture() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeCurveXYZTexture() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(VisualShaderNodeResizableBase.Advanced(self.AsVisualShaderNodeResizableBase()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNodeResizableBase.Instance(self.AsVisualShaderNodeResizableBase()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeCurveXYZTexture", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeCurveXYZTexture{*(*gdclass.VisualShaderNodeCurveXYZTexture)(unsafe.Pointer(&ptr))}
	})
}
