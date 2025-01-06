// Package World3D provides methods for working with World3D object instances.
package World3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Class that has everything pertaining to a world: A physics space, a visual scenario, and a sound space. 3D nodes register their resources into the current 3D world.
*/
type Instance [1]gdclass.World3D
type Any interface {
	gd.IsClass
	AsWorld3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.World3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("World3D"))
	casted := Instance{*(*gdclass.World3D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Environment() [1]gdclass.Environment {
	return [1]gdclass.Environment(class(self).GetEnvironment())
}

func (self Instance) SetEnvironment(value [1]gdclass.Environment) {
	class(self).SetEnvironment(value)
}

func (self Instance) FallbackEnvironment() [1]gdclass.Environment {
	return [1]gdclass.Environment(class(self).GetFallbackEnvironment())
}

func (self Instance) SetFallbackEnvironment(value [1]gdclass.Environment) {
	class(self).SetFallbackEnvironment(value)
}

func (self Instance) CameraAttributes() [1]gdclass.CameraAttributes {
	return [1]gdclass.CameraAttributes(class(self).GetCameraAttributes())
}

func (self Instance) SetCameraAttributes(value [1]gdclass.CameraAttributes) {
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

func (self Instance) DirectSpaceState() [1]gdclass.PhysicsDirectSpaceState3D {
	return [1]gdclass.PhysicsDirectSpaceState3D(class(self).GetDirectSpaceState())
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
func (self class) SetEnvironment(env [1]gdclass.Environment) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(env[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_set_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnvironment() [1]gdclass.Environment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_get_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Environment{gd.PointerWithOwnershipTransferredToGo[gdclass.Environment](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFallbackEnvironment(env [1]gdclass.Environment) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(env[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_set_fallback_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFallbackEnvironment() [1]gdclass.Environment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_get_fallback_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Environment{gd.PointerWithOwnershipTransferredToGo[gdclass.Environment](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCameraAttributes(attributes [1]gdclass.CameraAttributes) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(attributes[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_set_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCameraAttributes() [1]gdclass.CameraAttributes {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_get_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.CameraAttributes{gd.PointerWithOwnershipTransferredToGo[gdclass.CameraAttributes](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetDirectSpaceState() [1]gdclass.PhysicsDirectSpaceState3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.World3D.Bind_get_direct_space_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.PhysicsDirectSpaceState3D{gd.PointerMustAssertInstanceID[gdclass.PhysicsDirectSpaceState3D](r_ret.Get())}
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("World3D", func(ptr gd.Object) any { return [1]gdclass.World3D{*(*gdclass.World3D)(unsafe.Pointer(&ptr))} })
}
