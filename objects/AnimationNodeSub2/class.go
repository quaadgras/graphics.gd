package AnimationNodeSub2

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/AnimationNodeSync"
import "graphics.gd/objects/AnimationNode"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A resource to add to an [AnimationNodeBlendTree]. Blends two animations subtractively based on the amount value.
This animation node is usually used for pre-calculation to cancel out any extra poses from the animation for the "add" animation source in [AnimationNodeAdd2] or [AnimationNodeAdd3].
In general, the blend value should be in the [code][0.0, 1.0][/code] range, but values outside of this range can be used for amplified or inverted animations.
[b]Note:[/b] This calculation is different from using a negative value in [AnimationNodeAdd2], since the transformation matrices do not satisfy the commutative law. [AnimationNodeSub2] multiplies the transformation matrix of the inverted animation from the left side, while negative [AnimationNodeAdd2] multiplies it from the right side.
*/
type Instance [1]classdb.AnimationNodeSub2
type Any interface {
	gd.IsClass
	AsAnimationNodeSub2() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AnimationNodeSub2

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationNodeSub2"))
	return Instance{classdb.AnimationNodeSub2(object)}
}

func (self class) AsAnimationNodeSub2() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimationNodeSub2() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsAnimationNodeSync() AnimationNodeSync.Advanced {
	return *((*AnimationNodeSync.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationNodeSync() AnimationNodeSync.Instance {
	return *((*AnimationNodeSync.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(self.AsAnimationNodeSync(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAnimationNodeSync(), name)
	}
}
func init() {
	classdb.Register("AnimationNodeSub2", func(ptr gd.Object) any { return [1]classdb.AnimationNodeSub2{classdb.AnimationNodeSub2(ptr)} })
}
