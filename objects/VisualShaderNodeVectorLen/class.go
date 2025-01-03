package VisualShaderNodeVectorLen

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/VisualShaderNodeVectorBase"
import "graphics.gd/objects/VisualShaderNode"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Translated to [code]length(p0)[/code] in the shader language.
*/
type Instance [1]classdb.VisualShaderNodeVectorLen
type Any interface {
	gd.IsClass
	AsVisualShaderNodeVectorLen() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VisualShaderNodeVectorLen

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeVectorLen"))
	return Instance{*(*classdb.VisualShaderNodeVectorLen)(unsafe.Pointer(&object))}
}

func (self class) AsVisualShaderNodeVectorLen() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeVectorLen() Instance {
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNodeVectorBase(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNodeVectorBase(), name)
	}
}
func init() {
	classdb.Register("VisualShaderNodeVectorLen", func(ptr gd.Object) any {
		return [1]classdb.VisualShaderNodeVectorLen{*(*classdb.VisualShaderNodeVectorLen)(unsafe.Pointer(&ptr))}
	})
}
