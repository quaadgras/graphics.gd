package DirectionalLight3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Light3D"
import "grow.graphics/gd/object/VisualInstance3D"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A directional light is a type of [Light3D] node that models an infinite number of parallel rays covering the entire scene. It is used for lights with strong intensity that are located far away from the scene to model sunlight or moonlight. The worldspace location of the DirectionalLight3D transform (origin) is ignored. Only the basis is used to determine light direction.

*/
type Simple [1]classdb.DirectionalLight3D
func (self Simple) SetShadowMode(mode classdb.DirectionalLight3DShadowMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShadowMode(mode)
}
func (self Simple) GetShadowMode() classdb.DirectionalLight3DShadowMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.DirectionalLight3DShadowMode(Expert(self).GetShadowMode())
}
func (self Simple) SetBlendSplits(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBlendSplits(enabled)
}
func (self Simple) IsBlendSplitsEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsBlendSplitsEnabled())
}
func (self Simple) SetSkyMode(mode classdb.DirectionalLight3DSkyMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSkyMode(mode)
}
func (self Simple) GetSkyMode() classdb.DirectionalLight3DSkyMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.DirectionalLight3DSkyMode(Expert(self).GetSkyMode())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.DirectionalLight3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetShadowMode(mode classdb.DirectionalLight3DShadowMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirectionalLight3D.Bind_set_shadow_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShadowMode() classdb.DirectionalLight3DShadowMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.DirectionalLight3DShadowMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirectionalLight3D.Bind_get_shadow_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBlendSplits(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirectionalLight3D.Bind_set_blend_splits, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsBlendSplitsEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirectionalLight3D.Bind_is_blend_splits_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkyMode(mode classdb.DirectionalLight3DSkyMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirectionalLight3D.Bind_set_sky_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkyMode() classdb.DirectionalLight3DSkyMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.DirectionalLight3DSkyMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirectionalLight3D.Bind_get_sky_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsDirectionalLight3D() Expert { return self[0].AsDirectionalLight3D() }


//go:nosplit
func (self Simple) AsDirectionalLight3D() Simple { return self[0].AsDirectionalLight3D() }


//go:nosplit
func (self class) AsLight3D() Light3D.Expert { return self[0].AsLight3D() }


//go:nosplit
func (self Simple) AsLight3D() Light3D.Simple { return self[0].AsLight3D() }


//go:nosplit
func (self class) AsVisualInstance3D() VisualInstance3D.Expert { return self[0].AsVisualInstance3D() }


//go:nosplit
func (self Simple) AsVisualInstance3D() VisualInstance3D.Simple { return self[0].AsVisualInstance3D() }


//go:nosplit
func (self class) AsNode3D() Node3D.Expert { return self[0].AsNode3D() }


//go:nosplit
func (self Simple) AsNode3D() Node3D.Simple { return self[0].AsNode3D() }


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
func init() {classdb.Register("DirectionalLight3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type ShadowMode = classdb.DirectionalLight3DShadowMode

const (
/*Renders the entire scene's shadow map from an orthogonal point of view. This is the fastest directional shadow mode. May result in blurrier shadows on close objects.*/
	ShadowOrthogonal ShadowMode = 0
/*Splits the view frustum in 2 areas, each with its own shadow map. This shadow mode is a compromise between [constant SHADOW_ORTHOGONAL] and [constant SHADOW_PARALLEL_4_SPLITS] in terms of performance.*/
	ShadowParallel2Splits ShadowMode = 1
/*Splits the view frustum in 4 areas, each with its own shadow map. This is the slowest directional shadow mode.*/
	ShadowParallel4Splits ShadowMode = 2
)
type SkyMode = classdb.DirectionalLight3DSkyMode

const (
/*Makes the light visible in both scene lighting and sky rendering.*/
	SkyModeLightAndSky SkyMode = 0
/*Makes the light visible in scene lighting only (including direct lighting and global illumination). When using this mode, the light will not be visible from sky shaders.*/
	SkyModeLightOnly SkyMode = 1
/*Makes the light visible to sky shaders only. When using this mode the light will not cast light into the scene (either through direct lighting or through global illumination), but can be accessed through sky shaders. This can be useful, for example, when you want to control sky effects without illuminating the scene (during a night cycle, for example).*/
	SkyModeSkyOnly SkyMode = 2
)
