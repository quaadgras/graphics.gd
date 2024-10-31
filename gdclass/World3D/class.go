package World3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
Class that has everything pertaining to a world: A physics space, a visual scenario, and a sound space. 3D nodes register their resources into the current 3D world.
*/
type Instance [1]classdb.World3D

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.World3D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("World3D"))
	return Instance{classdb.World3D(object)}
}

func (self Instance) Environment() gdclass.Environment {
	return gdclass.Environment(class(self).GetEnvironment())
}

func (self Instance) SetEnvironment(value gdclass.Environment) {
	class(self).SetEnvironment(value)
}

func (self Instance) FallbackEnvironment() gdclass.Environment {
	return gdclass.Environment(class(self).GetFallbackEnvironment())
}

func (self Instance) SetFallbackEnvironment(value gdclass.Environment) {
	class(self).SetFallbackEnvironment(value)
}

func (self Instance) CameraAttributes() gdclass.CameraAttributes {
	return gdclass.CameraAttributes(class(self).GetCameraAttributes())
}

func (self Instance) SetCameraAttributes(value gdclass.CameraAttributes) {
	class(self).SetCameraAttributes(value)
}

func (self Instance) Space() gd.RID {
	return gd.RID(class(self).GetSpace())
}

func (self Instance) NavigationMap() gd.RID {
	return gd.RID(class(self).GetNavigationMap())
}

func (self Instance) Scenario() gd.RID {
	return gd.RID(class(self).GetScenario())
}

func (self Instance) DirectSpaceState() gdclass.PhysicsDirectSpaceState3D {
	return gdclass.PhysicsDirectSpaceState3D(class(self).GetDirectSpaceState())
}

//go:nosplit
func (self class) GetSpace() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_get_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetNavigationMap() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_get_navigation_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetScenario() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_get_scenario, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnvironment(env gdclass.Environment) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(env[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_set_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnvironment() gdclass.Environment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_get_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Environment{classdb.Environment(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFallbackEnvironment(env gdclass.Environment) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(env[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_set_fallback_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFallbackEnvironment() gdclass.Environment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_get_fallback_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Environment{classdb.Environment(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCameraAttributes(attributes gdclass.CameraAttributes) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(attributes[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_set_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCameraAttributes() gdclass.CameraAttributes {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_get_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.CameraAttributes{classdb.CameraAttributes(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetDirectSpaceState() gdclass.PhysicsDirectSpaceState3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_get_direct_space_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.PhysicsDirectSpaceState3D{classdb.PhysicsDirectSpaceState3D(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsWorld3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsWorld3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() { classdb.Register("World3D", func(ptr gd.Object) any { return classdb.World3D(ptr) }) }
