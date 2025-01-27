// Package XRHandModifier3D provides methods for working with XRHandModifier3D object instances.
package XRHandModifier3D

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
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/classdb/SkeletonModifier3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"

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

/*
This node uses hand tracking data from an [XRHandTracker] to pose the skeleton of a hand mesh.
Positioning of hands is performed by creating an [XRNode3D] ancestor of the hand mesh driven by the same [XRHandTracker].
The hand tracking position-data is scaled by [member Skeleton3D.motion_scale] when applied to the skeleton, which can be used to adjust the tracked hand to match the scale of the hand model.
*/
type Instance [1]gdclass.XRHandModifier3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsXRHandModifier3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.XRHandModifier3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XRHandModifier3D"))
	casted := Instance{*(*gdclass.XRHandModifier3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) HandTracker() string {
	return string(class(self).GetHandTracker().String())
}

func (self Instance) SetHandTracker(value string) {
	class(self).SetHandTracker(String.Name(String.New(value)))
}

func (self Instance) BoneUpdate() gdclass.XRHandModifier3DBoneUpdate {
	return gdclass.XRHandModifier3DBoneUpdate(class(self).GetBoneUpdate())
}

func (self Instance) SetBoneUpdate(value gdclass.XRHandModifier3DBoneUpdate) {
	class(self).SetBoneUpdate(value)
}

//go:nosplit
func (self class) SetHandTracker(tracker_name String.Name) { //gd:XRHandModifier3D.set_hand_tracker
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(tracker_name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandModifier3D.Bind_set_hand_tracker, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHandTracker() String.Name { //gd:XRHandModifier3D.get_hand_tracker
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandModifier3D.Bind_get_hand_tracker, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBoneUpdate(bone_update gdclass.XRHandModifier3DBoneUpdate) { //gd:XRHandModifier3D.set_bone_update
	var frame = callframe.New()
	callframe.Arg(frame, bone_update)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandModifier3D.Bind_set_bone_update, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBoneUpdate() gdclass.XRHandModifier3DBoneUpdate { //gd:XRHandModifier3D.get_bone_update
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.XRHandModifier3DBoneUpdate](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandModifier3D.Bind_get_bone_update, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsXRHandModifier3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsXRHandModifier3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsSkeletonModifier3D() SkeletonModifier3D.Advanced {
	return *((*SkeletonModifier3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSkeletonModifier3D() SkeletonModifier3D.Instance {
	return *((*SkeletonModifier3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(SkeletonModifier3D.Advanced(self.AsSkeletonModifier3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(SkeletonModifier3D.Instance(self.AsSkeletonModifier3D()), name)
	}
}
func init() {
	gdclass.Register("XRHandModifier3D", func(ptr gd.Object) any {
		return [1]gdclass.XRHandModifier3D{*(*gdclass.XRHandModifier3D)(unsafe.Pointer(&ptr))}
	})
}

type BoneUpdate = gdclass.XRHandModifier3DBoneUpdate //gd:XRHandModifier3D.BoneUpdate

const (
	/*The skeleton's bones are fully updated (both position and rotation) to match the tracked bones.*/
	BoneUpdateFull BoneUpdate = 0
	/*The skeleton's bones are only rotated to align with the tracked bones, preserving bone length.*/
	BoneUpdateRotationOnly BoneUpdate = 1
	/*Represents the size of the [enum BoneUpdate] enum.*/
	BoneUpdateMax BoneUpdate = 2
)
