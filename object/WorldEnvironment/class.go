package WorldEnvironment

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The [WorldEnvironment] node is used to configure the default [Environment] for the scene.
The parameters defined in the [WorldEnvironment] can be overridden by an [Environment] node set on the current [Camera3D]. Additionally, only one [WorldEnvironment] may be instantiated in a given scene at a time.
The [WorldEnvironment] allows the user to specify default lighting parameters (e.g. ambient lighting), various post-processing effects (e.g. SSAO, DOF, Tonemapping), and how to draw the background (e.g. solid color, skybox). Usually, these are added in order to improve the realism/color balance of the scene.

*/
type Simple [1]classdb.WorldEnvironment
func (self Simple) SetEnvironment(env [1]classdb.Environment) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnvironment(env)
}
func (self Simple) GetEnvironment() [1]classdb.Environment {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Environment(Expert(self).GetEnvironment(gc))
}
func (self Simple) SetCameraAttributes(camera_attributes [1]classdb.CameraAttributes) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCameraAttributes(camera_attributes)
}
func (self Simple) GetCameraAttributes() [1]classdb.CameraAttributes {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.CameraAttributes(Expert(self).GetCameraAttributes(gc))
}
func (self Simple) SetCompositor(compositor [1]classdb.Compositor) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCompositor(compositor)
}
func (self Simple) GetCompositor() [1]classdb.Compositor {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Compositor(Expert(self).GetCompositor(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.WorldEnvironment
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetEnvironment(env object.Environment)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(env[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WorldEnvironment.Bind_set_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnvironment(ctx gd.Lifetime) object.Environment {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WorldEnvironment.Bind_get_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Environment
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCameraAttributes(camera_attributes object.CameraAttributes)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(camera_attributes[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WorldEnvironment.Bind_set_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCameraAttributes(ctx gd.Lifetime) object.CameraAttributes {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WorldEnvironment.Bind_get_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.CameraAttributes
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCompositor(compositor object.Compositor)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(compositor[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WorldEnvironment.Bind_set_compositor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCompositor(ctx gd.Lifetime) object.Compositor {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WorldEnvironment.Bind_get_compositor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Compositor
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsWorldEnvironment() Expert { return self[0].AsWorldEnvironment() }


//go:nosplit
func (self Simple) AsWorldEnvironment() Simple { return self[0].AsWorldEnvironment() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("WorldEnvironment", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
