package VisualShaderNodeGlobalExpression

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/VisualShaderNodeExpression"
import "grow.graphics/gd/object/VisualShaderNodeGroupBase"
import "grow.graphics/gd/object/VisualShaderNodeResizableBase"
import "grow.graphics/gd/object/VisualShaderNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Custom Godot Shader Language expression, which is placed on top of the generated shader. You can place various function definitions inside to call later in [VisualShaderNodeExpression]s (which are injected in the main shader functions). You can also declare varyings, uniforms and global constants.

*/
type Simple [1]classdb.VisualShaderNodeGlobalExpression
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeGlobalExpression
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsVisualShaderNodeGlobalExpression() Expert { return self[0].AsVisualShaderNodeGlobalExpression() }


//go:nosplit
func (self Simple) AsVisualShaderNodeGlobalExpression() Simple { return self[0].AsVisualShaderNodeGlobalExpression() }


//go:nosplit
func (self class) AsVisualShaderNodeExpression() VisualShaderNodeExpression.Expert { return self[0].AsVisualShaderNodeExpression() }


//go:nosplit
func (self Simple) AsVisualShaderNodeExpression() VisualShaderNodeExpression.Simple { return self[0].AsVisualShaderNodeExpression() }


//go:nosplit
func (self class) AsVisualShaderNodeGroupBase() VisualShaderNodeGroupBase.Expert { return self[0].AsVisualShaderNodeGroupBase() }


//go:nosplit
func (self Simple) AsVisualShaderNodeGroupBase() VisualShaderNodeGroupBase.Simple { return self[0].AsVisualShaderNodeGroupBase() }


//go:nosplit
func (self class) AsVisualShaderNodeResizableBase() VisualShaderNodeResizableBase.Expert { return self[0].AsVisualShaderNodeResizableBase() }


//go:nosplit
func (self Simple) AsVisualShaderNodeResizableBase() VisualShaderNodeResizableBase.Simple { return self[0].AsVisualShaderNodeResizableBase() }


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
func init() {classdb.Register("VisualShaderNodeGlobalExpression", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
