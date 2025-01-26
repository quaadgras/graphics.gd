// Package VisualShaderNodeTexture2DArray provides methods for working with VisualShaderNodeTexture2DArray object instances.
package VisualShaderNodeTexture2DArray

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
import "graphics.gd/classdb/VisualShaderNodeSample3D"
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
Translated to [code]uniform sampler2DArray[/code] in the shader language.
*/
type Instance [1]gdclass.VisualShaderNodeTexture2DArray

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeTexture2DArray() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeTexture2DArray

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeTexture2DArray"))
	casted := Instance{*(*gdclass.VisualShaderNodeTexture2DArray)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) TextureArray() [1]gdclass.Texture2DArray {
	return [1]gdclass.Texture2DArray(class(self).GetTextureArray())
}

func (self Instance) SetTextureArray(value [1]gdclass.Texture2DArray) {
	class(self).SetTextureArray(value)
}

//go:nosplit
func (self class) SetTextureArray(value [1]gdclass.Texture2DArray) { //gd:VisualShaderNodeTexture2DArray.set_texture_array
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(value[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTexture2DArray.Bind_set_texture_array, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureArray() [1]gdclass.Texture2DArray { //gd:VisualShaderNodeTexture2DArray.get_texture_array
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTexture2DArray.Bind_get_texture_array, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2DArray{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2DArray](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeTexture2DArray() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeTexture2DArray() Instance {
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
	gdclass.Register("VisualShaderNodeTexture2DArray", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeTexture2DArray{*(*gdclass.VisualShaderNodeTexture2DArray)(unsafe.Pointer(&ptr))}
	})
}
