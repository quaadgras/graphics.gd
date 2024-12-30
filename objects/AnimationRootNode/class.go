package AnimationRootNode

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/AnimationNode"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
[AnimationRootNode] is a base class for [AnimationNode]s that hold a complete animation. A complete animation refers to the output of an [AnimationNodeOutput] in an [AnimationNodeBlendTree] or the output of another [AnimationRootNode]. Used for [member AnimationTree.tree_root] or in other [AnimationRootNode]s.
Examples of built-in root nodes include [AnimationNodeBlendTree] (allows blending nodes between each other using various modes), [AnimationNodeStateMachine] (allows to configure blending and transitions between nodes using a state machine pattern), [AnimationNodeBlendSpace2D] (allows linear blending between [b]three[/b] [AnimationNode]s), [AnimationNodeBlendSpace1D] (allows linear blending only between [b]two[/b] [AnimationNode]s).
*/
type Instance [1]classdb.AnimationRootNode
type Any interface {
	gd.IsClass
	AsAnimationRootNode() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AnimationRootNode

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationRootNode"))
	return Instance{classdb.AnimationRootNode(object)}
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAnimationNode(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAnimationNode(), name)
	}
}
func init() {
	classdb.Register("AnimationRootNode", func(ptr gd.Object) any { return classdb.AnimationRootNode(ptr) })
}
