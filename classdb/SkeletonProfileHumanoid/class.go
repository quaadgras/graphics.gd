// Package SkeletonProfileHumanoid provides methods for working with SkeletonProfileHumanoid object instances.
package SkeletonProfileHumanoid

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/SkeletonProfile"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A [SkeletonProfile] as a preset that is optimized for the human form. This exists for standardization, so all parameters are read-only.
A humanoid skeleton profile contains 54 bones divided in 4 groups: [code]"Body"[/code], [code]"Face"[/code], [code]"LeftHand"[/code], and [code]"RightHand"[/code]. It is structured as follows:
[codeblock lang=text]
Root
└─ Hips

	├─ LeftUpperLeg
	│  └─ LeftLowerLeg
	│     └─ LeftFoot
	│        └─ LeftToes
	├─ RightUpperLeg
	│  └─ RightLowerLeg
	│     └─ RightFoot
	│        └─ RightToes
	└─ Spine
	    └─ Chest
	        └─ UpperChest
	            ├─ Neck
	            │   └─ Head
	            │       ├─ Jaw
	            │       ├─ LeftEye
	            │       └─ RightEye
	            ├─ LeftShoulder
	            │  └─ LeftUpperArm
	            │     └─ LeftLowerArm
	            │        └─ LeftHand
	            │           ├─ LeftThumbMetacarpal
	            │           │  └─ LeftThumbProximal
	            │           ├─ LeftIndexProximal
	            │           │  └─ LeftIndexIntermediate
	            │           │    └─ LeftIndexDistal
	            │           ├─ LeftMiddleProximal
	            │           │  └─ LeftMiddleIntermediate
	            │           │    └─ LeftMiddleDistal
	            │           ├─ LeftRingProximal
	            │           │  └─ LeftRingIntermediate
	            │           │    └─ LeftRingDistal
	            │           └─ LeftLittleProximal
	            │              └─ LeftLittleIntermediate
	            │                └─ LeftLittleDistal
	            └─ RightShoulder
	               └─ RightUpperArm
	                  └─ RightLowerArm
	                     └─ RightHand
	                        ├─ RightThumbMetacarpal
	                        │  └─ RightThumbProximal
	                        ├─ RightIndexProximal
	                        │  └─ RightIndexIntermediate
	                        │     └─ RightIndexDistal
	                        ├─ RightMiddleProximal
	                        │  └─ RightMiddleIntermediate
	                        │     └─ RightMiddleDistal
	                        ├─ RightRingProximal
	                        │  └─ RightRingIntermediate
	                        │     └─ RightRingDistal
	                        └─ RightLittleProximal
	                           └─ RightLittleIntermediate
	                             └─ RightLittleDistal

[/codeblock]
*/
type Instance [1]gdclass.SkeletonProfileHumanoid
type Any interface {
	gd.IsClass
	AsSkeletonProfileHumanoid() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SkeletonProfileHumanoid

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SkeletonProfileHumanoid"))
	return Instance{*(*gdclass.SkeletonProfileHumanoid)(unsafe.Pointer(&object))}
}

func (self class) AsSkeletonProfileHumanoid() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSkeletonProfileHumanoid() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsSkeletonProfile() SkeletonProfile.Advanced {
	return *((*SkeletonProfile.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSkeletonProfile() SkeletonProfile.Instance {
	return *((*SkeletonProfile.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(SkeletonProfile.Advanced(self.AsSkeletonProfile()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(SkeletonProfile.Instance(self.AsSkeletonProfile()), name)
	}
}
func init() {
	gdclass.Register("SkeletonProfileHumanoid", func(ptr gd.Object) any {
		return [1]gdclass.SkeletonProfileHumanoid{*(*gdclass.SkeletonProfileHumanoid)(unsafe.Pointer(&ptr))}
	})
}
