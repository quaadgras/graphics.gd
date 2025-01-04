// Package WorldEnvironment provides methods for working with WorldEnvironment object instances.
package WorldEnvironment

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Node"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
The [WorldEnvironment] node is used to configure the default [Environment] for the scene.
The parameters defined in the [WorldEnvironment] can be overridden by an [Environment] node set on the current [Camera3D]. Additionally, only one [WorldEnvironment] may be instantiated in a given scene at a time.
The [WorldEnvironment] allows the user to specify default lighting parameters (e.g. ambient lighting), various post-processing effects (e.g. SSAO, DOF, Tonemapping), and how to draw the background (e.g. solid color, skybox). Usually, these are added in order to improve the realism/color balance of the scene.
*/
type Instance [1]gdclass.WorldEnvironment
type Any interface {
	gd.IsClass
	AsWorldEnvironment() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.WorldEnvironment

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("WorldEnvironment"))
	return Instance{*(*gdclass.WorldEnvironment)(unsafe.Pointer(&object))}
}

func (self Instance) Environment() [1]gdclass.Environment {
	return [1]gdclass.Environment(class(self).GetEnvironment())
}

func (self Instance) SetEnvironment(value [1]gdclass.Environment) {
	class(self).SetEnvironment(value)
}

func (self Instance) CameraAttributes() [1]gdclass.CameraAttributes {
	return [1]gdclass.CameraAttributes(class(self).GetCameraAttributes())
}

func (self Instance) SetCameraAttributes(value [1]gdclass.CameraAttributes) {
	class(self).SetCameraAttributes(value)
}

func (self Instance) Compositor() [1]gdclass.Compositor {
	return [1]gdclass.Compositor(class(self).GetCompositor())
}

func (self Instance) SetCompositor(value [1]gdclass.Compositor) {
	class(self).SetCompositor(value)
}

//go:nosplit
func (self class) SetEnvironment(env [1]gdclass.Environment) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(env[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldEnvironment.Bind_set_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnvironment() [1]gdclass.Environment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldEnvironment.Bind_get_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Environment{gd.PointerWithOwnershipTransferredToGo[gdclass.Environment](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCameraAttributes(camera_attributes [1]gdclass.CameraAttributes) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(camera_attributes[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldEnvironment.Bind_set_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCameraAttributes() [1]gdclass.CameraAttributes {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldEnvironment.Bind_get_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.CameraAttributes{gd.PointerWithOwnershipTransferredToGo[gdclass.CameraAttributes](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCompositor(compositor [1]gdclass.Compositor) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(compositor[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldEnvironment.Bind_set_compositor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCompositor() [1]gdclass.Compositor {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldEnvironment.Bind_get_compositor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Compositor{gd.PointerWithOwnershipTransferredToGo[gdclass.Compositor](r_ret.Get())}
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
	gdclass.Register("WorldEnvironment", func(ptr gd.Object) any {
		return [1]gdclass.WorldEnvironment{*(*gdclass.WorldEnvironment)(unsafe.Pointer(&ptr))}
	})
}
