package SkeletonProfileHumanoid

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/SkeletonProfile"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

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
type Instance [1]classdb.SkeletonProfileHumanoid

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.SkeletonProfileHumanoid

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SkeletonProfileHumanoid"))
	return Instance{classdb.SkeletonProfileHumanoid(object)}
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsSkeletonProfile(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsSkeletonProfile(), name)
	}
}
func init() {
	classdb.Register("SkeletonProfileHumanoid", func(ptr gd.Object) any { return classdb.SkeletonProfileHumanoid(ptr) })
}
