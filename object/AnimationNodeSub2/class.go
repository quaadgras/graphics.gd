package AnimationNodeSub2

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/AnimationNodeSync"
import "grow.graphics/gd/object/AnimationNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A resource to add to an [AnimationNodeBlendTree]. Blends two animations subtractively based on the amount value.
This animation node is usually used for pre-calculation to cancel out any extra poses from the animation for the "add" animation source in [AnimationNodeAdd2] or [AnimationNodeAdd3].
In general, the blend value should be in the [code][0.0, 1.0][/code] range, but values outside of this range can be used for amplified or inverted animations.
[b]Note:[/b] This calculation is different from using a negative value in [AnimationNodeAdd2], since the transformation matrices do not satisfy the commutative law. [AnimationNodeSub2] multiplies the transformation matrix of the inverted animation from the left side, while negative [AnimationNodeAdd2] multiplies it from the right side.

*/
type Simple [1]classdb.AnimationNodeSub2
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AnimationNodeSub2
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsAnimationNodeSub2() Expert { return self[0].AsAnimationNodeSub2() }


//go:nosplit
func (self Simple) AsAnimationNodeSub2() Simple { return self[0].AsAnimationNodeSub2() }


//go:nosplit
func (self class) AsAnimationNodeSync() AnimationNodeSync.Expert { return self[0].AsAnimationNodeSync() }


//go:nosplit
func (self Simple) AsAnimationNodeSync() AnimationNodeSync.Simple { return self[0].AsAnimationNodeSync() }


//go:nosplit
func (self class) AsAnimationNode() AnimationNode.Expert { return self[0].AsAnimationNode() }


//go:nosplit
func (self Simple) AsAnimationNode() AnimationNode.Simple { return self[0].AsAnimationNode() }


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
func init() {classdb.Register("AnimationNodeSub2", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
