// Package SkeletonModification2DPhysicalBones provides methods for working with SkeletonModification2DPhysicalBones object instances.
package SkeletonModification2DPhysicalBones

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
import "graphics.gd/classdb/SkeletonModification2D"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
This modification takes the transforms of [PhysicalBone2D] nodes and applies them to [Bone2D] nodes. This allows the [Bone2D] nodes to react to physics thanks to the linked [PhysicalBone2D] nodes.
*/
type Instance [1]gdclass.SkeletonModification2DPhysicalBones
type Expanded [1]gdclass.SkeletonModification2DPhysicalBones

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSkeletonModification2DPhysicalBones() Instance
}

/*
Sets the [PhysicalBone2D] node at [param joint_idx].
[b]Note:[/b] This is just the index used for this modification, not the bone index used in the [Skeleton2D].
*/
func (self Instance) SetPhysicalBoneNode(joint_idx int, physicalbone2d_node string) { //gd:SkeletonModification2DPhysicalBones.set_physical_bone_node
	Advanced(self).SetPhysicalBoneNode(int64(joint_idx), Path.ToNode(String.New(physicalbone2d_node)))
}

/*
Returns the [PhysicalBone2D] node at [param joint_idx].
*/
func (self Instance) GetPhysicalBoneNode(joint_idx int) string { //gd:SkeletonModification2DPhysicalBones.get_physical_bone_node
	return string(Advanced(self).GetPhysicalBoneNode(int64(joint_idx)).String())
}

/*
Empties the list of [PhysicalBone2D] nodes and populates it with all [PhysicalBone2D] nodes that are children of the [Skeleton2D].
*/
func (self Instance) FetchPhysicalBones() { //gd:SkeletonModification2DPhysicalBones.fetch_physical_bones
	Advanced(self).FetchPhysicalBones()
}

/*
Tell the [PhysicalBone2D] nodes to start simulating and interacting with the physics world.
Optionally, an array of bone names can be passed to this function, and that will cause only [PhysicalBone2D] nodes with those names to start simulating.
*/
func (self Instance) StartSimulation() { //gd:SkeletonModification2DPhysicalBones.start_simulation
	Advanced(self).StartSimulation(gd.ArrayFromSlice[Array.Contains[String.Name]]([1][]string{}[0]))
}

/*
Tell the [PhysicalBone2D] nodes to start simulating and interacting with the physics world.
Optionally, an array of bone names can be passed to this function, and that will cause only [PhysicalBone2D] nodes with those names to start simulating.
*/
func (self Expanded) StartSimulation(bones []string) { //gd:SkeletonModification2DPhysicalBones.start_simulation
	Advanced(self).StartSimulation(gd.ArrayFromSlice[Array.Contains[String.Name]](bones))
}

/*
Tell the [PhysicalBone2D] nodes to stop simulating and interacting with the physics world.
Optionally, an array of bone names can be passed to this function, and that will cause only [PhysicalBone2D] nodes with those names to stop simulating.
*/
func (self Instance) StopSimulation() { //gd:SkeletonModification2DPhysicalBones.stop_simulation
	Advanced(self).StopSimulation(gd.ArrayFromSlice[Array.Contains[String.Name]]([1][]string{}[0]))
}

/*
Tell the [PhysicalBone2D] nodes to stop simulating and interacting with the physics world.
Optionally, an array of bone names can be passed to this function, and that will cause only [PhysicalBone2D] nodes with those names to stop simulating.
*/
func (self Expanded) StopSimulation(bones []string) { //gd:SkeletonModification2DPhysicalBones.stop_simulation
	Advanced(self).StopSimulation(gd.ArrayFromSlice[Array.Contains[String.Name]](bones))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SkeletonModification2DPhysicalBones

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SkeletonModification2DPhysicalBones"))
	casted := Instance{*(*gdclass.SkeletonModification2DPhysicalBones)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) PhysicalBoneChainLength() int {
	return int(int(class(self).GetPhysicalBoneChainLength()))
}

func (self Instance) SetPhysicalBoneChainLength(value int) {
	class(self).SetPhysicalBoneChainLength(int64(value))
}

//go:nosplit
func (self class) SetPhysicalBoneChainLength(length int64) { //gd:SkeletonModification2DPhysicalBones.set_physical_bone_chain_length
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DPhysicalBones.Bind_set_physical_bone_chain_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPhysicalBoneChainLength() int64 { //gd:SkeletonModification2DPhysicalBones.get_physical_bone_chain_length
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DPhysicalBones.Bind_get_physical_bone_chain_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [PhysicalBone2D] node at [param joint_idx].
[b]Note:[/b] This is just the index used for this modification, not the bone index used in the [Skeleton2D].
*/
//go:nosplit
func (self class) SetPhysicalBoneNode(joint_idx int64, physicalbone2d_node Path.ToNode) { //gd:SkeletonModification2DPhysicalBones.set_physical_bone_node
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(physicalbone2d_node)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DPhysicalBones.Bind_set_physical_bone_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [PhysicalBone2D] node at [param joint_idx].
*/
//go:nosplit
func (self class) GetPhysicalBoneNode(joint_idx int64) Path.ToNode { //gd:SkeletonModification2DPhysicalBones.get_physical_bone_node
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DPhysicalBones.Bind_get_physical_bone_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Empties the list of [PhysicalBone2D] nodes and populates it with all [PhysicalBone2D] nodes that are children of the [Skeleton2D].
*/
//go:nosplit
func (self class) FetchPhysicalBones() { //gd:SkeletonModification2DPhysicalBones.fetch_physical_bones
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DPhysicalBones.Bind_fetch_physical_bones, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Tell the [PhysicalBone2D] nodes to start simulating and interacting with the physics world.
Optionally, an array of bone names can be passed to this function, and that will cause only [PhysicalBone2D] nodes with those names to start simulating.
*/
//go:nosplit
func (self class) StartSimulation(bones Array.Contains[String.Name]) { //gd:SkeletonModification2DPhysicalBones.start_simulation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(bones)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DPhysicalBones.Bind_start_simulation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Tell the [PhysicalBone2D] nodes to stop simulating and interacting with the physics world.
Optionally, an array of bone names can be passed to this function, and that will cause only [PhysicalBone2D] nodes with those names to stop simulating.
*/
//go:nosplit
func (self class) StopSimulation(bones Array.Contains[String.Name]) { //gd:SkeletonModification2DPhysicalBones.stop_simulation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(bones)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DPhysicalBones.Bind_stop_simulation, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(SkeletonModification2D.Advanced(self.AsSkeletonModification2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(SkeletonModification2D.Instance(self.AsSkeletonModification2D()), name)
	}
}
func init() {
	gdclass.Register("SkeletonModification2DPhysicalBones", func(ptr gd.Object) any {
		return [1]gdclass.SkeletonModification2DPhysicalBones{*(*gdclass.SkeletonModification2DPhysicalBones)(unsafe.Pointer(&ptr))}
	})
}
