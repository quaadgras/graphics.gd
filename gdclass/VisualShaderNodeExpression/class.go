package VisualShaderNodeExpression

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualShaderNodeGroupBase"
import "grow.graphics/gd/gdclass/VisualShaderNodeResizableBase"
import "grow.graphics/gd/gdclass/VisualShaderNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Custom Godot Shading Language expression, with a custom number of input and output ports.
The provided code is directly injected into the graph's matching shader function ([code]vertex[/code], [code]fragment[/code], or [code]light[/code]), so it cannot be used to declare functions, varyings, uniforms, or global constants. See [VisualShaderNodeGlobalExpression] for such global definitions.

*/
type Go [1]classdb.VisualShaderNodeExpression
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeExpression
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VisualShaderNodeExpression"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Expression() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetExpression(gc).String())
}

func (self Go) SetExpression(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetExpression(gc.String(value))
}

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
func (self class) AsVisualShaderNodeExpression() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeExpression() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNodeGroupBase() VisualShaderNodeGroupBase.GD { return *((*VisualShaderNodeGroupBase.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeGroupBase() VisualShaderNodeGroupBase.Go { return *((*VisualShaderNodeGroupBase.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNodeResizableBase() VisualShaderNodeResizableBase.GD { return *((*VisualShaderNodeResizableBase.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeResizableBase() VisualShaderNodeResizableBase.Go { return *((*VisualShaderNodeResizableBase.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNode() VisualShaderNode.GD { return *((*VisualShaderNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNode() VisualShaderNode.Go { return *((*VisualShaderNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNodeGroupBase(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNodeGroupBase(), name)
	}
}
func init() {classdb.Register("VisualShaderNodeExpression", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
