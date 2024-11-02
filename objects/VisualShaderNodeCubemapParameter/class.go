package VisualShaderNodeCubemapParameter

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/VisualShaderNodeTextureParameter"
import "grow.graphics/gd/objects/VisualShaderNodeParameter"
import "grow.graphics/gd/objects/VisualShaderNode"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Translated to [code]uniform samplerCube[/code] in the shader language. The output value can be used as port for [VisualShaderNodeCubemap].
*/
type Instance [1]classdb.VisualShaderNodeCubemapParameter

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VisualShaderNodeCubemapParameter

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeCubemapParameter"))
	return Instance{classdb.VisualShaderNodeCubemapParameter(object)}
}

func (self class) AsVisualShaderNodeCubemapParameter() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeCubemapParameter() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNodeTextureParameter() VisualShaderNodeTextureParameter.Advanced {
	return *((*VisualShaderNodeTextureParameter.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeTextureParameter() VisualShaderNodeTextureParameter.Instance {
	return *((*VisualShaderNodeTextureParameter.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNodeParameter() VisualShaderNodeParameter.Advanced {
	return *((*VisualShaderNodeParameter.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeParameter() VisualShaderNodeParameter.Instance {
	return *((*VisualShaderNodeParameter.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsVisualShaderNodeTextureParameter(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNodeTextureParameter(), name)
	}
}
func init() {
	classdb.Register("VisualShaderNodeCubemapParameter", func(ptr gd.Object) any { return classdb.VisualShaderNodeCubemapParameter(ptr) })
}
