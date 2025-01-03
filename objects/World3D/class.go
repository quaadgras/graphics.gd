package World3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Class that has everything pertaining to a world: A physics space, a visual scenario, and a sound space. 3D nodes register their resources into the current 3D world.
*/
type Instance [1]classdb.World3D
type Any interface {
	gd.IsClass
	AsWorld3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.World3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("World3D"))
	return Instance{*(*classdb.World3D)(unsafe.Pointer(&object))}
}

func (self Instance) Environment() objects.Environment {
	return objects.Environment(class(self).GetEnvironment())
}

func (self Instance) SetEnvironment(value objects.Environment) {
	class(self).SetEnvironment(value)
}

func (self Instance) FallbackEnvironment() objects.Environment {
	return objects.Environment(class(self).GetFallbackEnvironment())
}

func (self Instance) SetFallbackEnvironment(value objects.Environment) {
	class(self).SetFallbackEnvironment(value)
}

func (self Instance) CameraAttributes() objects.CameraAttributes {
	return objects.CameraAttributes(class(self).GetCameraAttributes())
}

func (self Instance) SetCameraAttributes(value objects.CameraAttributes) {
	class(self).SetCameraAttributes(value)
}

func (self Instance) Space() Resource.ID {
	return Resource.ID(class(self).GetSpace())
}

func (self Instance) NavigationMap() Resource.ID {
	return Resource.ID(class(self).GetNavigationMap())
}

func (self Instance) Scenario() Resource.ID {
	return Resource.ID(class(self).GetScenario())
}

func (self Instance) DirectSpaceState() objects.PhysicsDirectSpaceState3D {
	return objects.PhysicsDirectSpaceState3D(class(self).GetDirectSpaceState())
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
func (self class) SetEnvironment(env objects.Environment) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(env[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_set_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnvironment() objects.Environment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_get_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Environment{gd.PointerWithOwnershipTransferredToGo[classdb.Environment](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFallbackEnvironment(env objects.Environment) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(env[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_set_fallback_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFallbackEnvironment() objects.Environment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_get_fallback_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Environment{gd.PointerWithOwnershipTransferredToGo[classdb.Environment](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCameraAttributes(attributes objects.CameraAttributes) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(attributes[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_set_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCameraAttributes() objects.CameraAttributes {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_get_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.CameraAttributes{gd.PointerWithOwnershipTransferredToGo[classdb.CameraAttributes](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetDirectSpaceState() objects.PhysicsDirectSpaceState3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_get_direct_space_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.PhysicsDirectSpaceState3D{gd.PointerMustAssertInstanceID[classdb.PhysicsDirectSpaceState3D](r_ret.Get())}
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
func init() {
	classdb.Register("World3D", func(ptr gd.Object) any { return [1]classdb.World3D{*(*classdb.World3D)(unsafe.Pointer(&ptr))} })
}
