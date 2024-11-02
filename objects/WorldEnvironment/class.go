package WorldEnvironment

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
The [WorldEnvironment] node is used to configure the default [Environment] for the scene.
The parameters defined in the [WorldEnvironment] can be overridden by an [Environment] node set on the current [Camera3D]. Additionally, only one [WorldEnvironment] may be instantiated in a given scene at a time.
The [WorldEnvironment] allows the user to specify default lighting parameters (e.g. ambient lighting), various post-processing effects (e.g. SSAO, DOF, Tonemapping), and how to draw the background (e.g. solid color, skybox). Usually, these are added in order to improve the realism/color balance of the scene.
*/
type Instance [1]classdb.WorldEnvironment

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.WorldEnvironment

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("WorldEnvironment"))
	return Instance{classdb.WorldEnvironment(object)}
}

func (self Instance) Environment() objects.Environment {
	return objects.Environment(class(self).GetEnvironment())
}

func (self Instance) SetEnvironment(value objects.Environment) {
	class(self).SetEnvironment(value)
}

func (self Instance) CameraAttributes() objects.CameraAttributes {
	return objects.CameraAttributes(class(self).GetCameraAttributes())
}

func (self Instance) SetCameraAttributes(value objects.CameraAttributes) {
	class(self).SetCameraAttributes(value)
}

func (self Instance) Compositor() objects.Compositor {
	return objects.Compositor(class(self).GetCompositor())
}

func (self Instance) SetCompositor(value objects.Compositor) {
	class(self).SetCompositor(value)
}

//go:nosplit
func (self class) SetEnvironment(env objects.Environment) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(env[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldEnvironment.Bind_set_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnvironment() objects.Environment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldEnvironment.Bind_get_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Environment{classdb.Environment(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCameraAttributes(camera_attributes objects.CameraAttributes) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(camera_attributes[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldEnvironment.Bind_set_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCameraAttributes() objects.CameraAttributes {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldEnvironment.Bind_get_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.CameraAttributes{classdb.CameraAttributes(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCompositor(compositor objects.Compositor) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(compositor[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldEnvironment.Bind_set_compositor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCompositor() objects.Compositor {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldEnvironment.Bind_get_compositor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Compositor{classdb.Compositor(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsWorldEnvironment() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsWorldEnvironment() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced           { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance        { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode(), name)
	}
}
func init() {
	classdb.Register("WorldEnvironment", func(ptr gd.Object) any { return classdb.WorldEnvironment(ptr) })
}
