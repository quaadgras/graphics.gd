package VisualShaderNodeVaryingGetter

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/VisualShaderNodeVarying"
import "graphics.gd/objects/VisualShaderNode"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Outputs a value of a varying defined in the shader. You need to first create a varying that can be used in the given function, e.g. varying getter in Fragment shader requires a varying with mode set to [constant VisualShader.VARYING_MODE_VERTEX_TO_FRAG_LIGHT].
*/
type Instance [1]classdb.VisualShaderNodeVaryingGetter
type Any interface {
	gd.IsClass
	AsVisualShaderNodeVaryingGetter() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VisualShaderNodeVaryingGetter

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeVaryingGetter"))
	return Instance{*(*classdb.VisualShaderNodeVaryingGetter)(unsafe.Pointer(&object))}
}

func (self class) AsVisualShaderNodeVaryingGetter() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeVaryingGetter() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNodeVarying() VisualShaderNodeVarying.Advanced {
	return *((*VisualShaderNodeVarying.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeVarying() VisualShaderNodeVarying.Instance {
	return *((*VisualShaderNodeVarying.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsVisualShaderNodeVarying(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNodeVarying(), name)
	}
}
func init() {
	classdb.Register("VisualShaderNodeVaryingGetter", func(ptr gd.Object) any {
		return [1]classdb.VisualShaderNodeVaryingGetter{*(*classdb.VisualShaderNodeVaryingGetter)(unsafe.Pointer(&ptr))}
	})
}
