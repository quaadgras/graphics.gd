package AnimationNodeBlend3

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AnimationNodeSync"
import "grow.graphics/gd/gdclass/AnimationNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A resource to add to an [AnimationNodeBlendTree]. Blends two animations out of three linearly out of three based on the amount value.
This animation node has three inputs:
- The base animation to blend with
- A "-blend" animation to blend with when the blend amount is negative value
- A "+blend" animation to blend with when the blend amount is positive value
In general, the blend value should be in the [code][-1.0, 1.0][/code] range. Values outside of this range can blend amplified animations, however, [AnimationNodeAdd3] works better for this purpose.

*/
type Go [1]classdb.AnimationNodeBlend3
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AnimationNodeBlend3
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AnimationNodeBlend3"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsAnimationNodeBlend3() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationNodeBlend3() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsAnimationNodeSync() AnimationNodeSync.GD { return *((*AnimationNodeSync.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationNodeSync() AnimationNodeSync.Go { return *((*AnimationNodeSync.Go)(unsafe.Pointer(&self))) }
func (self class) AsAnimationNode() AnimationNode.GD { return *((*AnimationNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationNode() AnimationNode.Go { return *((*AnimationNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAnimationNodeSync(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAnimationNodeSync(), name)
	}
}
func init() {classdb.Register("AnimationNodeBlend3", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
