package XRFaceModifier3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Node3D"
import "graphics.gd/objects/Node"
import "graphics.gd/variant/Path"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This node applies weights from a [XRFaceTracker] to a mesh with supporting face blend shapes.
The [url=https://docs.vrcft.io/docs/tutorial-avatars/tutorial-avatars-extras/unified-blendshapes]Unified Expressions[/url] blend shapes are supported, as well as ARKit and SRanipal blend shapes.
The node attempts to identify blend shapes based on name matching. Blend shapes should match the names listed in the [url=https://docs.vrcft.io/docs/tutorial-avatars/tutorial-avatars-extras/compatibility/overview]Unified Expressions Compatibility[/url] chart.
*/
type Instance [1]classdb.XRFaceModifier3D
type Any interface {
	gd.IsClass
	AsXRFaceModifier3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.XRFaceModifier3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XRFaceModifier3D"))
	return Instance{*(*classdb.XRFaceModifier3D)(unsafe.Pointer(&object))}
}

func (self Instance) FaceTracker() string {
	return string(class(self).GetFaceTracker().String())
}

func (self Instance) SetFaceTracker(value string) {
	class(self).SetFaceTracker(gd.NewStringName(value))
}

func (self Instance) Target() Path.String {
	return Path.String(class(self).GetTarget().String())
}

func (self Instance) SetTarget(value Path.String) {
	class(self).SetTarget(gd.NewString(string(value)).NodePath())
}

//go:nosplit
func (self class) SetFaceTracker(tracker_name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(tracker_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRFaceModifier3D.Bind_set_face_tracker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFaceTracker() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRFaceModifier3D.Bind_get_face_tracker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTarget(target gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(target))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRFaceModifier3D.Bind_set_target, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTarget() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRFaceModifier3D.Bind_get_target, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsXRFaceModifier3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsXRFaceModifier3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced       { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance    { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced           { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance        { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {
	classdb.Register("XRFaceModifier3D", func(ptr gd.Object) any {
		return [1]classdb.XRFaceModifier3D{*(*classdb.XRFaceModifier3D)(unsafe.Pointer(&ptr))}
	})
}
