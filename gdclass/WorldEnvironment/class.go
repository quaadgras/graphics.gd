package WorldEnvironment

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
The [WorldEnvironment] node is used to configure the default [Environment] for the scene.
The parameters defined in the [WorldEnvironment] can be overridden by an [Environment] node set on the current [Camera3D]. Additionally, only one [WorldEnvironment] may be instantiated in a given scene at a time.
The [WorldEnvironment] allows the user to specify default lighting parameters (e.g. ambient lighting), various post-processing effects (e.g. SSAO, DOF, Tonemapping), and how to draw the background (e.g. solid color, skybox). Usually, these are added in order to improve the realism/color balance of the scene.

*/
type Go [1]classdb.WorldEnvironment
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.WorldEnvironment
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("WorldEnvironment"))
	return Go{classdb.WorldEnvironment(object)}
}

func (self Go) Environment() gdclass.Environment {
		return gdclass.Environment(class(self).GetEnvironment())
}

func (self Go) SetEnvironment(value gdclass.Environment) {
	class(self).SetEnvironment(value)
}

func (self Go) CameraAttributes() gdclass.CameraAttributes {
		return gdclass.CameraAttributes(class(self).GetCameraAttributes())
}

func (self Go) SetCameraAttributes(value gdclass.CameraAttributes) {
	class(self).SetCameraAttributes(value)
}

func (self Go) Compositor() gdclass.Compositor {
		return gdclass.Compositor(class(self).GetCompositor())
}

func (self Go) SetCompositor(value gdclass.Compositor) {
	class(self).SetCompositor(value)
}

//go:nosplit
func (self class) SetEnvironment(env gdclass.Environment)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(env[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldEnvironment.Bind_set_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnvironment() gdclass.Environment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldEnvironment.Bind_get_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Environment{classdb.Environment(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCameraAttributes(camera_attributes gdclass.CameraAttributes)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(camera_attributes[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldEnvironment.Bind_set_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCameraAttributes() gdclass.CameraAttributes {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldEnvironment.Bind_get_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.CameraAttributes{classdb.CameraAttributes(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCompositor(compositor gdclass.Compositor)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(compositor[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldEnvironment.Bind_set_compositor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCompositor() gdclass.Compositor {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldEnvironment.Bind_get_compositor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Compositor{classdb.Compositor(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsWorldEnvironment() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsWorldEnvironment() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}
func init() {classdb.Register("WorldEnvironment", func(ptr gd.Object) any { return classdb.WorldEnvironment(ptr) })}
