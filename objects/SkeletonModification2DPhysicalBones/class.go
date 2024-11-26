package SkeletonModification2DPhysicalBones

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/SkeletonModification2D"
import "grow.graphics/gd/objects/Resource"
import "grow.graphics/gd/variant/Path"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This modification takes the transforms of [PhysicalBone2D] nodes and applies them to [Bone2D] nodes. This allows the [Bone2D] nodes to react to physics thanks to the linked [PhysicalBone2D] nodes.
*/
type Instance [1]classdb.SkeletonModification2DPhysicalBones

/*
Sets the [PhysicalBone2D] node at [param joint_idx].
[b]Note:[/b] This is just the index used for this modification, not the bone index used in the [Skeleton2D].
*/
func (self Instance) SetPhysicalBoneNode(joint_idx int, physicalbone2d_node Path.String) {
	class(self).SetPhysicalBoneNode(gd.Int(joint_idx), gd.NewString(string(physicalbone2d_node)).NodePath())
}

/*
Returns the [PhysicalBone2D] node at [param joint_idx].
*/
func (self Instance) GetPhysicalBoneNode(joint_idx int) Path.String {
	return Path.String(class(self).GetPhysicalBoneNode(gd.Int(joint_idx)).String())
}

/*
Empties the list of [PhysicalBone2D] nodes and populates it with all [PhysicalBone2D] nodes that are children of the [Skeleton2D].
*/
func (self Instance) FetchPhysicalBones() {
	class(self).FetchPhysicalBones()
}

/*
Tell the [PhysicalBone2D] nodes to start simulating and interacting with the physics world.
Optionally, an array of bone names can be passed to this function, and that will cause only [PhysicalBone2D] nodes with those names to start simulating.
*/
func (self Instance) StartSimulation() {
	class(self).StartSimulation([1]gd.Array{}[0])
}

/*
Tell the [PhysicalBone2D] nodes to stop simulating and interacting with the physics world.
Optionally, an array of bone names can be passed to this function, and that will cause only [PhysicalBone2D] nodes with those names to stop simulating.
*/
func (self Instance) StopSimulation() {
	class(self).StopSimulation([1]gd.Array{}[0])
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.SkeletonModification2DPhysicalBones

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SkeletonModification2DPhysicalBones"))
	return Instance{classdb.SkeletonModification2DPhysicalBones(object)}
}

func (self Instance) PhysicalBoneChainLength() int {
	return int(int(class(self).GetPhysicalBoneChainLength()))
}

func (self Instance) SetPhysicalBoneChainLength(value int) {
	class(self).SetPhysicalBoneChainLength(gd.Int(value))
}

//go:nosplit
func (self class) SetPhysicalBoneChainLength(length gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DPhysicalBones.Bind_set_physical_bone_chain_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPhysicalBoneChainLength() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DPhysicalBones.Bind_get_physical_bone_chain_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [PhysicalBone2D] node at [param joint_idx].
[b]Note:[/b] This is just the index used for this modification, not the bone index used in the [Skeleton2D].
*/
//go:nosplit
func (self class) SetPhysicalBoneNode(joint_idx gd.Int, physicalbone2d_node gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, pointers.Get(physicalbone2d_node))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DPhysicalBones.Bind_set_physical_bone_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [PhysicalBone2D] node at [param joint_idx].
*/
//go:nosplit
func (self class) GetPhysicalBoneNode(joint_idx gd.Int) gd.NodePath {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DPhysicalBones.Bind_get_physical_bone_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

/*
Empties the list of [PhysicalBone2D] nodes and populates it with all [PhysicalBone2D] nodes that are children of the [Skeleton2D].
*/
//go:nosplit
func (self class) FetchPhysicalBones() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DPhysicalBones.Bind_fetch_physical_bones, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Tell the [PhysicalBone2D] nodes to start simulating and interacting with the physics world.
Optionally, an array of bone names can be passed to this function, and that will cause only [PhysicalBone2D] nodes with those names to start simulating.
*/
//go:nosplit
func (self class) StartSimulation(bones gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(bones))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DPhysicalBones.Bind_start_simulation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Tell the [PhysicalBone2D] nodes to stop simulating and interacting with the physics world.
Optionally, an array of bone names can be passed to this function, and that will cause only [PhysicalBone2D] nodes with those names to stop simulating.
*/
//go:nosplit
func (self class) StopSimulation(bones gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(bones))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DPhysicalBones.Bind_stop_simulation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsSkeletonModification2DPhysicalBones() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSkeletonModification2DPhysicalBones() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsSkeletonModification2D() SkeletonModification2D.Advanced {
	return *((*SkeletonModification2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSkeletonModification2D() SkeletonModification2D.Instance {
	return *((*SkeletonModification2D.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsSkeletonModification2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsSkeletonModification2D(), name)
	}
}
func init() {
	classdb.Register("SkeletonModification2DPhysicalBones", func(ptr gd.Object) any { return classdb.SkeletonModification2DPhysicalBones(ptr) })
}
