package AnimationNodeBlend3

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
var _ = pointers.Cycle

/*
A resource to add to an [AnimationNodeBlendTree]. Blends two animations out of three linearly out of three based on the amount value.
This animation node has three inputs:
- The base animation to blend with
- A "-blend" animation to blend with when the blend amount is negative value
- A "+blend" animation to blend with when the blend amount is positive value
In general, the blend value should be in the [code][-1.0, 1.0][/code] range. Values outside of this range can blend amplified animations, however, [AnimationNodeAdd3] works better for this purpose.
*/
type Instance [1]classdb.AnimationNodeBlend3
type Any interface {
	gd.IsClass
	AsAnimationNodeBlend3() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AnimationNodeBlend3

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationNodeBlend3"))
	return Instance{*(*classdb.AnimationNodeBlend3)(unsafe.Pointer(&object))}
}

func (self class) AsAnimationNodeBlend3() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimationNodeBlend3() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	classdb.Register("AnimationNodeBlend3", func(ptr gd.Object) any {
		return [1]classdb.AnimationNodeBlend3{*(*classdb.AnimationNodeBlend3)(unsafe.Pointer(&ptr))}
	})
}
