// Package AnimationRootNode provides methods for working with AnimationRootNode object instances.
package AnimationRootNode

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/classdb/AnimationNode"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any

/*
[AnimationRootNode] is a base class for [AnimationNode]s that hold a complete animation. A complete animation refers to the output of an [AnimationNodeOutput] in an [AnimationNodeBlendTree] or the output of another [AnimationRootNode]. Used for [member AnimationTree.tree_root] or in other [AnimationRootNode]s.
Examples of built-in root nodes include [AnimationNodeBlendTree] (allows blending nodes between each other using various modes), [AnimationNodeStateMachine] (allows to configure blending and transitions between nodes using a state machine pattern), [AnimationNodeBlendSpace2D] (allows linear blending between [b]three[/b] [AnimationNode]s), [AnimationNodeBlendSpace1D] (allows linear blending only between [b]two[/b] [AnimationNode]s).
*/
type Instance [1]gdclass.AnimationRootNode

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAnimationRootNode() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AnimationRootNode

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationRootNode"))
	casted := Instance{*(*gdclass.AnimationRootNode)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self class) AsAnimationRootNode() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimationRootNode() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsAnimationNode() AnimationNode.Advanced {
	return *((*AnimationNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationNode() AnimationNode.Instance {
	return *((*AnimationNode.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AnimationNode.Advanced(self.AsAnimationNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AnimationNode.Instance(self.AsAnimationNode()), name)
	}
}
func init() {
	gdclass.Register("AnimationRootNode", func(ptr gd.Object) any {
		return [1]gdclass.AnimationRootNode{*(*gdclass.AnimationRootNode)(unsafe.Pointer(&ptr))}
	})
}
