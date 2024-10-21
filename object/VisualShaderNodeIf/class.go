package VisualShaderNodeIf

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/VisualShaderNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This visual shader node has six input ports:
- Port [b]1[/b] and [b]2[/b] provide the two floating-point numbers [code]a[/code] and [code]b[/code] that will be compared.
- Port [b]3[/b] is the tolerance, which allows similar floating-point numbers to be considered equal.
- Ports [b]4[/b], [b]5[/b], and [b]6[/b] are the possible outputs, returned if [code]a == b[/code], [code]a > b[/code], or [code]a < b[/code] respectively.

*/
type Simple [1]classdb.VisualShaderNodeIf
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeIf
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsVisualShaderNodeIf() Expert { return self[0].AsVisualShaderNodeIf() }


//go:nosplit
func (self Simple) AsVisualShaderNodeIf() Simple { return self[0].AsVisualShaderNodeIf() }


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
func init() {classdb.Register("VisualShaderNodeIf", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
