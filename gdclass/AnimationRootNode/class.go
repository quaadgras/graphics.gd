package AnimationRootNode

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AnimationNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[AnimationRootNode] is a base class for [AnimationNode]s that hold a complete animation. A complete animation refers to the output of an [AnimationNodeOutput] in an [AnimationNodeBlendTree] or the output of another [AnimationRootNode]. Used for [member AnimationTree.tree_root] or in other [AnimationRootNode]s.
Examples of built-in root nodes include [AnimationNodeBlendTree] (allows blending nodes between each other using various modes), [AnimationNodeStateMachine] (allows to configure blending and transitions between nodes using a state machine pattern), [AnimationNodeBlendSpace2D] (allows linear blending between [b]three[/b] [AnimationNode]s), [AnimationNodeBlendSpace1D] (allows linear blending only between [b]two[/b] [AnimationNode]s).

*/
type Go [1]classdb.AnimationRootNode
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AnimationRootNode
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AnimationRootNode"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsAnimationRootNode() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationRootNode() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsAnimationNode() AnimationNode.GD { return *((*AnimationNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationNode() AnimationNode.Go { return *((*AnimationNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAnimationNode(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAnimationNode(), name)
	}
}
func init() {classdb.Register("AnimationRootNode", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
