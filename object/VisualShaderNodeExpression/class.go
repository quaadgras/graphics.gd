package VisualShaderNodeExpression

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
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
Custom Godot Shading Language expression, with a custom number of input and output ports.
The provided code is directly injected into the graph's matching shader function ([code]vertex[/code], [code]fragment[/code], or [code]light[/code]), so it cannot be used to declare functions, varyings, uniforms, or global constants. See [VisualShaderNodeGlobalExpression] for such global definitions.

*/
type Simple [1]classdb.VisualShaderNodeExpression
func (self Simple) SetExpression(expression string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetExpression(gc.String(expression))
}
func (self Simple) GetExpression() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetExpression(gc).String())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeExpression
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetExpression(expression gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(expression))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeExpression.Bind_set_expression, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetExpression(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeExpression.Bind_get_expression, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShaderNodeExpression() Expert { return self[0].AsVisualShaderNodeExpression() }


//go:nosplit
func (self Simple) AsVisualShaderNodeExpression() Simple { return self[0].AsVisualShaderNodeExpression() }


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

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("VisualShaderNodeExpression", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
