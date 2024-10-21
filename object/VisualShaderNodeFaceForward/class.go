package VisualShaderNodeFaceForward

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/VisualShaderNodeVectorBase"
import "grow.graphics/gd/object/VisualShaderNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Translates to [code]faceforward(N, I, Nref)[/code] in the shader language. The function has three vector parameters: [code]N[/code], the vector to orient, [code]I[/code], the incident vector, and [code]Nref[/code], the reference vector. If the dot product of [code]I[/code] and [code]Nref[/code] is smaller than zero the return value is [code]N[/code]. Otherwise, [code]-N[/code] is returned.

*/
type Simple [1]classdb.VisualShaderNodeFaceForward
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeFaceForward
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsVisualShaderNodeFaceForward() Expert { return self[0].AsVisualShaderNodeFaceForward() }


//go:nosplit
func (self Simple) AsVisualShaderNodeFaceForward() Simple { return self[0].AsVisualShaderNodeFaceForward() }


//go:nosplit
func (self class) AsVisualShaderNodeVectorBase() VisualShaderNodeVectorBase.Expert { return self[0].AsVisualShaderNodeVectorBase() }


//go:nosplit
func (self Simple) AsVisualShaderNodeVectorBase() VisualShaderNodeVectorBase.Simple { return self[0].AsVisualShaderNodeVectorBase() }


//go:nosplit
func (self class) AsVisualShaderNode() VisualShaderNode.Expert { return self[0].AsVisualShaderNode() }


//go:nosplit
func (self Simple) AsVisualShaderNode() VisualShaderNode.Simple { return self[0].AsVisualShaderNode() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("VisualShaderNodeFaceForward", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
