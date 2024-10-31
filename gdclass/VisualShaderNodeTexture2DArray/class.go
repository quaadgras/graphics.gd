package VisualShaderNodeTexture2DArray

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualShaderNodeSample3D"
import "grow.graphics/gd/gdclass/VisualShaderNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
Translated to [code]uniform sampler2DArray[/code] in the shader language.
*/
type Instance [1]classdb.VisualShaderNodeTexture2DArray

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VisualShaderNodeTexture2DArray

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeTexture2DArray"))
	return Instance{classdb.VisualShaderNodeTexture2DArray(object)}
}

func (self Instance) TextureArray() gdclass.Texture2DArray {
	return gdclass.Texture2DArray(class(self).GetTextureArray())
}

func (self Instance) SetTextureArray(value gdclass.Texture2DArray) {
	class(self).SetTextureArray(value)
}

//go:nosplit
func (self class) SetTextureArray(value gdclass.Texture2DArray) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(value[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTexture2DArray.Bind_set_texture_array, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureArray() gdclass.Texture2DArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTexture2DArray.Bind_get_texture_array, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2DArray{classdb.Texture2DArray(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNodeSample3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNodeSample3D(), name)
	}
}
func init() {
	classdb.Register("VisualShaderNodeTexture2DArray", func(ptr gd.Object) any { return classdb.VisualShaderNodeTexture2DArray(ptr) })
}
