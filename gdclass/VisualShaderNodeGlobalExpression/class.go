package VisualShaderNodeGlobalExpression

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualShaderNodeExpression"
import "grow.graphics/gd/gdclass/VisualShaderNodeGroupBase"
import "grow.graphics/gd/gdclass/VisualShaderNodeResizableBase"
import "grow.graphics/gd/gdclass/VisualShaderNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Custom Godot Shader Language expression, which is placed on top of the generated shader. You can place various function definitions inside to call later in [VisualShaderNodeExpression]s (which are injected in the main shader functions). You can also declare varyings, uniforms and global constants.

*/
type Go [1]classdb.VisualShaderNodeGlobalExpression
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeGlobalExpression
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeGlobalExpression"))
	return Go{classdb.VisualShaderNodeGlobalExpression(object)}
}

func (self class) AsVisualShaderNodeGlobalExpression() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeGlobalExpression() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNodeExpression() VisualShaderNodeExpression.GD { return *((*VisualShaderNodeExpression.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeExpression() VisualShaderNodeExpression.Go { return *((*VisualShaderNodeExpression.Go)(unsafe.Pointer(&self))) }
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
	default: return gd.VirtualByName(self.AsVisualShaderNodeExpression(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNodeExpression(), name)
	}
}
func init() {classdb.Register("VisualShaderNodeGlobalExpression", func(ptr gd.Object) any { return classdb.VisualShaderNodeGlobalExpression(ptr) })}
