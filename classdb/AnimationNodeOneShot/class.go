// Package AnimationNodeOneShot provides methods for working with AnimationNodeOneShot object instances.
package AnimationNodeOneShot

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/AnimationNodeSync"
import "graphics.gd/classdb/AnimationNode"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A resource to add to an [AnimationNodeBlendTree]. This animation node will execute a sub-animation and return once it finishes. Blend times for fading in and out can be customized, as well as filters.
After setting the request and changing the animation playback, the one-shot node automatically clears the request on the next process frame by setting its [code]request[/code] value to [constant ONE_SHOT_REQUEST_NONE].
[codeblocks]
[gdscript]
# Play child animation connected to "shot" port.
animation_tree.set("parameters/OneShot/request", AnimationNodeOneShot.ONE_SHOT_REQUEST_FIRE)
# Alternative syntax (same result as above).
animation_tree["parameters/OneShot/request"] = AnimationNodeOneShot.ONE_SHOT_REQUEST_FIRE

# Abort child animation connected to "shot" port.
animation_tree.set("parameters/OneShot/request", AnimationNodeOneShot.ONE_SHOT_REQUEST_ABORT)
# Alternative syntax (same result as above).
animation_tree["parameters/OneShot/request"] = AnimationNodeOneShot.ONE_SHOT_REQUEST_ABORT

# Abort child animation with fading out connected to "shot" port.
animation_tree.set("parameters/OneShot/request", AnimationNodeOneShot.ONE_SHOT_REQUEST_FADE_OUT)
# Alternative syntax (same result as above).
animation_tree["parameters/OneShot/request"] = AnimationNodeOneShot.ONE_SHOT_REQUEST_FADE_OUT

# Get current state (read-only).
animation_tree.get("parameters/OneShot/active")
# Alternative syntax (same result as above).
animation_tree["parameters/OneShot/active"]

# Get current internal state (read-only).
animation_tree.get("parameters/OneShot/internal_active")
# Alternative syntax (same result as above).
animation_tree["parameters/OneShot/internal_active"]
[/gdscript]
[csharp]
// Play child animation connected to "shot" port.
animationTree.Set("parameters/OneShot/request", (int)AnimationNodeOneShot.OneShotRequest.Fire);

// Abort child animation connected to "shot" port.
animationTree.Set("parameters/OneShot/request", (int)AnimationNodeOneShot.OneShotRequest.Abort);

// Abort child animation with fading out connected to "shot" port.
animationTree.Set("parameters/OneShot/request", (int)AnimationNodeOneShot.OneShotRequest.FadeOut);

// Get current state (read-only).
animationTree.Get("parameters/OneShot/active");

// Get current internal state (read-only).
animationTree.Get("parameters/OneShot/internal_active");
[/csharp]
[/codeblocks]
*/
type Instance [1]gdclass.AnimationNodeOneShot
type Any interface {
	gd.IsClass
	AsAnimationNodeOneShot() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AnimationNodeOneShot

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationNodeOneShot"))
	return Instance{*(*gdclass.AnimationNodeOneShot)(unsafe.Pointer(&object))}
}

func (self Instance) MixMode() gdclass.AnimationNodeOneShotMixMode {
	return gdclass.AnimationNodeOneShotMixMode(class(self).GetMixMode())
}

func (self Instance) SetMixMode(value gdclass.AnimationNodeOneShotMixMode) {
	class(self).SetMixMode(value)
}

func (self Instance) FadeinTime() Float.X {
	return Float.X(Float.X(class(self).GetFadeinTime()))
}

func (self Instance) SetFadeinTime(value Float.X) {
	class(self).SetFadeinTime(gd.Float(value))
}

func (self Instance) FadeinCurve() [1]gdclass.Curve {
	return [1]gdclass.Curve(class(self).GetFadeinCurve())
}

func (self Instance) SetFadeinCurve(value [1]gdclass.Curve) {
	class(self).SetFadeinCurve(value)
}

func (self Instance) FadeoutTime() Float.X {
	return Float.X(Float.X(class(self).GetFadeoutTime()))
}

func (self Instance) SetFadeoutTime(value Float.X) {
	class(self).SetFadeoutTime(gd.Float(value))
}

func (self Instance) FadeoutCurve() [1]gdclass.Curve {
	return [1]gdclass.Curve(class(self).GetFadeoutCurve())
}

func (self Instance) SetFadeoutCurve(value [1]gdclass.Curve) {
	class(self).SetFadeoutCurve(value)
}

func (self Instance) BreakLoopAtEnd() bool {
	return bool(class(self).IsLoopBrokenAtEnd())
}

func (self Instance) SetBreakLoopAtEnd(value bool) {
	class(self).SetBreakLoopAtEnd(value)
}

func (self Instance) Autorestart() bool {
	return bool(class(self).HasAutorestart())
}

func (self Instance) SetAutorestart(value bool) {
	class(self).SetAutorestart(value)
}

func (self Instance) AutorestartDelay() Float.X {
	return Float.X(Float.X(class(self).GetAutorestartDelay()))
}

func (self Instance) SetAutorestartDelay(value Float.X) {
	class(self).SetAutorestartDelay(gd.Float(value))
}

func (self Instance) AutorestartRandomDelay() Float.X {
	return Float.X(Float.X(class(self).GetAutorestartRandomDelay()))
}

func (self Instance) SetAutorestartRandomDelay(value Float.X) {
	class(self).SetAutorestartRandomDelay(gd.Float(value))
}

//go:nosplit
func (self class) SetFadeinTime(time gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, time)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_set_fadein_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFadeinTime() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_get_fadein_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFadeinCurve(curve [1]gdclass.Curve) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_set_fadein_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFadeinCurve() [1]gdclass.Curve {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_get_fadein_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Curve{gd.PointerWithOwnershipTransferredToGo[gdclass.Curve](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFadeoutTime(time gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, time)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_set_fadeout_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFadeoutTime() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_get_fadeout_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFadeoutCurve(curve [1]gdclass.Curve) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_set_fadeout_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFadeoutCurve() [1]gdclass.Curve {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_get_fadeout_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Curve{gd.PointerWithOwnershipTransferredToGo[gdclass.Curve](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBreakLoopAtEnd(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_set_break_loop_at_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsLoopBrokenAtEnd() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_is_loop_broken_at_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutorestart(active bool) {
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_set_autorestart, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) HasAutorestart() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_has_autorestart, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutorestartDelay(time gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, time)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_set_autorestart_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutorestartDelay() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_get_autorestart_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutorestartRandomDelay(time gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, time)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_set_autorestart_random_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutorestartRandomDelay() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_get_autorestart_random_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMixMode(mode gdclass.AnimationNodeOneShotMixMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_set_mix_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMixMode() gdclass.AnimationNodeOneShotMixMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AnimationNodeOneShotMixMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_get_mix_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAnimationNodeOneShot() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimationNodeOneShot() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsAnimationNodeSync() AnimationNodeSync.Advanced {
	return *((*AnimationNodeSync.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationNodeSync() AnimationNodeSync.Instance {
	return *((*AnimationNodeSync.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAnimationNode() AnimationNode.Advanced {
	return *((*AnimationNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationNode() AnimationNode.Instance {
	return *((*AnimationNode.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(AnimationNodeSync.Advanced(self.AsAnimationNodeSync()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AnimationNodeSync.Instance(self.AsAnimationNodeSync()), name)
	}
}
func init() {
	gdclass.Register("AnimationNodeOneShot", func(ptr gd.Object) any {
		return [1]gdclass.AnimationNodeOneShot{*(*gdclass.AnimationNodeOneShot)(unsafe.Pointer(&ptr))}
	})
}

type OneShotRequest = gdclass.AnimationNodeOneShotOneShotRequest

const (
	/*The default state of the request. Nothing is done.*/
	OneShotRequestNone OneShotRequest = 0
	/*The request to play the animation connected to "shot" port.*/
	OneShotRequestFire OneShotRequest = 1
	/*The request to stop the animation connected to "shot" port.*/
	OneShotRequestAbort OneShotRequest = 2
	/*The request to fade out the animation connected to "shot" port.*/
	OneShotRequestFadeOut OneShotRequest = 3
)

type MixMode = gdclass.AnimationNodeOneShotMixMode

const (
	/*Blends two animations. See also [AnimationNodeBlend2].*/
	MixModeBlend MixMode = 0
	/*Blends two animations additively. See also [AnimationNodeAdd2].*/
	MixModeAdd MixMode = 1
)
