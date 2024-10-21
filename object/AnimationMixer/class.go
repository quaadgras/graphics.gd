package AnimationMixer

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
Base class for [AnimationPlayer] and [AnimationTree] to manage animation lists. It also has general properties and methods for playback and blending.
After instantiating the playback information data within the extended class, the blending is processed by the [AnimationMixer].
	// AnimationMixer methods that can be overridden by a [Class] that extends it.
	type AnimationMixer interface {
		//A virtual function for processing after getting a key during playback.
		PostProcessKeyValue(animation object.Animation, track gd.Int, value gd.Variant, object_id gd.Int, object_sub_idx gd.Int) gd.Variant
	}

*/
type Simple [1]classdb.AnimationMixer
func (Simple) _post_process_key_value(impl func(ptr unsafe.Pointer, animation [1]classdb.Animation, track int, value gd.Variant, object_id int, object_sub_idx int) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var animation [1]classdb.Animation
		animation[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var track = gd.UnsafeGet[gd.Int](p_args,1)
		var value = mmm.Let[gd.Variant](gc.Lifetime, gc.API, gd.UnsafeGet[[3]uintptr](p_args,2))
		var object_id = gd.UnsafeGet[gd.Int](p_args,3)
		var object_sub_idx = gd.UnsafeGet[gd.Int](p_args,4)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, animation, int(track), value, int(object_id), int(object_sub_idx))
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (self Simple) AddAnimationLibrary(name string, library [1]classdb.AnimationLibrary) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).AddAnimationLibrary(gc.StringName(name), library))
}
func (self Simple) RemoveAnimationLibrary(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveAnimationLibrary(gc.StringName(name))
}
func (self Simple) RenameAnimationLibrary(name string, newname string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RenameAnimationLibrary(gc.StringName(name), gc.StringName(newname))
}
func (self Simple) HasAnimationLibrary(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasAnimationLibrary(gc.StringName(name)))
}
func (self Simple) GetAnimationLibrary(name string) [1]classdb.AnimationLibrary {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.AnimationLibrary(Expert(self).GetAnimationLibrary(gc, gc.StringName(name)))
}
func (self Simple) GetAnimationLibraryList() gd.ArrayOf[gd.StringName] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.StringName](Expert(self).GetAnimationLibraryList(gc))
}
func (self Simple) HasAnimation(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasAnimation(gc.StringName(name)))
}
func (self Simple) GetAnimation(name string) [1]classdb.Animation {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Animation(Expert(self).GetAnimation(gc, gc.StringName(name)))
}
func (self Simple) GetAnimationList() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetAnimationList(gc))
}
func (self Simple) SetActive(active bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetActive(active)
}
func (self Simple) IsActive() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsActive())
}
func (self Simple) SetDeterministic(deterministic bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDeterministic(deterministic)
}
func (self Simple) IsDeterministic() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDeterministic())
}
func (self Simple) SetRootNode(path gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRootNode(path)
}
func (self Simple) GetRootNode() gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetRootNode(gc))
}
func (self Simple) SetCallbackModeProcess(mode classdb.AnimationMixerAnimationCallbackModeProcess) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCallbackModeProcess(mode)
}
func (self Simple) GetCallbackModeProcess() classdb.AnimationMixerAnimationCallbackModeProcess {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AnimationMixerAnimationCallbackModeProcess(Expert(self).GetCallbackModeProcess())
}
func (self Simple) SetCallbackModeMethod(mode classdb.AnimationMixerAnimationCallbackModeMethod) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCallbackModeMethod(mode)
}
func (self Simple) GetCallbackModeMethod() classdb.AnimationMixerAnimationCallbackModeMethod {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AnimationMixerAnimationCallbackModeMethod(Expert(self).GetCallbackModeMethod())
}
func (self Simple) SetCallbackModeDiscrete(mode classdb.AnimationMixerAnimationCallbackModeDiscrete) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCallbackModeDiscrete(mode)
}
func (self Simple) GetCallbackModeDiscrete() classdb.AnimationMixerAnimationCallbackModeDiscrete {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AnimationMixerAnimationCallbackModeDiscrete(Expert(self).GetCallbackModeDiscrete())
}
func (self Simple) SetAudioMaxPolyphony(max_polyphony int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAudioMaxPolyphony(gd.Int(max_polyphony))
}
func (self Simple) GetAudioMaxPolyphony() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetAudioMaxPolyphony()))
}
func (self Simple) SetRootMotionTrack(path gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRootMotionTrack(path)
}
func (self Simple) GetRootMotionTrack() gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetRootMotionTrack(gc))
}
func (self Simple) GetRootMotionPosition() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetRootMotionPosition())
}
func (self Simple) GetRootMotionRotation() gd.Quaternion {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Quaternion(Expert(self).GetRootMotionRotation())
}
func (self Simple) GetRootMotionScale() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetRootMotionScale())
}
func (self Simple) GetRootMotionPositionAccumulator() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetRootMotionPositionAccumulator())
}
func (self Simple) GetRootMotionRotationAccumulator() gd.Quaternion {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Quaternion(Expert(self).GetRootMotionRotationAccumulator())
}
func (self Simple) GetRootMotionScaleAccumulator() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetRootMotionScaleAccumulator())
}
func (self Simple) ClearCaches() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearCaches()
}
func (self Simple) Advance(delta float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Advance(gd.Float(delta))
}
func (self Simple) Capture(name string, duration float64, trans_type classdb.TweenTransitionType, ease_type classdb.TweenEaseType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Capture(gc.StringName(name), gd.Float(duration), trans_type, ease_type)
}
func (self Simple) SetResetOnSaveEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetResetOnSaveEnabled(enabled)
}
func (self Simple) IsResetOnSaveEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsResetOnSaveEnabled())
}
func (self Simple) FindAnimation(animation [1]classdb.Animation) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).FindAnimation(gc, animation).String())
}
func (self Simple) FindAnimationLibrary(animation [1]classdb.Animation) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).FindAnimationLibrary(gc, animation).String())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AnimationMixer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
A virtual function for processing after getting a key during playback.
*/
func (class) _post_process_key_value(impl func(ptr unsafe.Pointer, animation object.Animation, track gd.Int, value gd.Variant, object_id gd.Int, object_sub_idx gd.Int) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var animation object.Animation
		animation[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var track = gd.UnsafeGet[gd.Int](p_args,1)
		var value = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, gd.UnsafeGet[[3]uintptr](p_args,2))
		var object_id = gd.UnsafeGet[gd.Int](p_args,3)
		var object_sub_idx = gd.UnsafeGet[gd.Int](p_args,4)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, animation, track, value, object_id, object_sub_idx)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Adds [param library] to the animation player, under the key [param name].
AnimationMixer has a global library by default with an empty string as key. For adding an animation to the global library:
[codeblocks]
[gdscript]
var global_library = mixer.get_animation_library("")
global_library.add_animation("animation_name", animation_resource)
[/gdscript]
[/codeblocks]
*/
//go:nosplit
func (self class) AddAnimationLibrary(name gd.StringName, library object.AnimationLibrary) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(library[0].AsPointer())[0])
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_add_animation_library, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the [AnimationLibrary] associated with the key [param name].
*/
//go:nosplit
func (self class) RemoveAnimationLibrary(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_remove_animation_library, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Moves the [AnimationLibrary] associated with the key [param name] to the key [param newname].
*/
//go:nosplit
func (self class) RenameAnimationLibrary(name gd.StringName, newname gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(newname))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_rename_animation_library, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the [AnimationMixer] stores an [AnimationLibrary] with key [param name].
*/
//go:nosplit
func (self class) HasAnimationLibrary(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_has_animation_library, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the first [AnimationLibrary] with key [param name] or [code]null[/code] if not found.
To get the [AnimationMixer]'s global animation library, use [code]get_animation_library("")[/code].
*/
//go:nosplit
func (self class) GetAnimationLibrary(ctx gd.Lifetime, name gd.StringName) object.AnimationLibrary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_get_animation_library, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.AnimationLibrary
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the list of stored library keys.
*/
//go:nosplit
func (self class) GetAnimationLibraryList(ctx gd.Lifetime) gd.ArrayOf[gd.StringName] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_get_animation_library_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.StringName](ret)
}
/*
Returns [code]true[/code] if the [AnimationMixer] stores an [Animation] with key [param name].
*/
//go:nosplit
func (self class) HasAnimation(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_has_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Animation] with the key [param name]. If the animation does not exist, [code]null[/code] is returned and an error is logged.
*/
//go:nosplit
func (self class) GetAnimation(ctx gd.Lifetime, name gd.StringName) object.Animation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_get_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Animation
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the list of stored animation keys.
*/
//go:nosplit
func (self class) GetAnimationList(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_get_animation_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetActive(active bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_set_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsActive() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_is_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDeterministic(deterministic bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, deterministic)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_set_deterministic, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDeterministic() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_is_deterministic, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRootNode(path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_set_root_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRootNode(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_get_root_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCallbackModeProcess(mode classdb.AnimationMixerAnimationCallbackModeProcess)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_set_callback_mode_process, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCallbackModeProcess() classdb.AnimationMixerAnimationCallbackModeProcess {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationMixerAnimationCallbackModeProcess](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_get_callback_mode_process, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCallbackModeMethod(mode classdb.AnimationMixerAnimationCallbackModeMethod)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_set_callback_mode_method, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCallbackModeMethod() classdb.AnimationMixerAnimationCallbackModeMethod {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationMixerAnimationCallbackModeMethod](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_get_callback_mode_method, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCallbackModeDiscrete(mode classdb.AnimationMixerAnimationCallbackModeDiscrete)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_set_callback_mode_discrete, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCallbackModeDiscrete() classdb.AnimationMixerAnimationCallbackModeDiscrete {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationMixerAnimationCallbackModeDiscrete](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_get_callback_mode_discrete, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAudioMaxPolyphony(max_polyphony gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_polyphony)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_set_audio_max_polyphony, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAudioMaxPolyphony() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_get_audio_max_polyphony, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRootMotionTrack(path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_set_root_motion_track, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRootMotionTrack(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_get_root_motion_track, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Retrieve the motion delta of position with the [member root_motion_track] as a [Vector3] that can be used elsewhere.
If [member root_motion_track] is not a path to a track of type [constant Animation.TYPE_POSITION_3D], returns [code]Vector3(0, 0, 0)[/code].
See also [member root_motion_track] and [RootMotionView].
The most basic example is applying position to [CharacterBody3D]:
[codeblocks]
[gdscript]
var current_rotation: Quaternion

func _process(delta):
    if Input.is_action_just_pressed("animate"):
        current_rotation = get_quaternion()
        state_machine.travel("Animate")
    var velocity: Vector3 = current_rotation * animation_tree.get_root_motion_position() / delta
    set_velocity(velocity)
    move_and_slide()
[/gdscript]
[/codeblocks]
By using this in combination with [method get_root_motion_rotation_accumulator], you can apply the root motion position more correctly to account for the rotation of the node.
[codeblocks]
[gdscript]
func _process(delta):
    if Input.is_action_just_pressed("animate"):
        state_machine.travel("Animate")
    set_quaternion(get_quaternion() * animation_tree.get_root_motion_rotation())
    var velocity: Vector3 = (animation_tree.get_root_motion_rotation_accumulator().inverse() * get_quaternion()) * animation_tree.get_root_motion_position() / delta
    set_velocity(velocity)
    move_and_slide()
[/gdscript]
[/codeblocks]
*/
//go:nosplit
func (self class) GetRootMotionPosition() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_get_root_motion_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Retrieve the motion delta of rotation with the [member root_motion_track] as a [Quaternion] that can be used elsewhere.
If [member root_motion_track] is not a path to a track of type [constant Animation.TYPE_ROTATION_3D], returns [code]Quaternion(0, 0, 0, 1)[/code].
See also [member root_motion_track] and [RootMotionView].
The most basic example is applying rotation to [CharacterBody3D]:
[codeblocks]
[gdscript]
func _process(delta):
    if Input.is_action_just_pressed("animate"):
        state_machine.travel("Animate")
    set_quaternion(get_quaternion() * animation_tree.get_root_motion_rotation())
[/gdscript]
[/codeblocks]
*/
//go:nosplit
func (self class) GetRootMotionRotation() gd.Quaternion {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Quaternion](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_get_root_motion_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Retrieve the motion delta of scale with the [member root_motion_track] as a [Vector3] that can be used elsewhere.
If [member root_motion_track] is not a path to a track of type [constant Animation.TYPE_SCALE_3D], returns [code]Vector3(0, 0, 0)[/code].
See also [member root_motion_track] and [RootMotionView].
The most basic example is applying scale to [CharacterBody3D]:
[codeblocks]
[gdscript]
var current_scale: Vector3 = Vector3(1, 1, 1)
var scale_accum: Vector3 = Vector3(1, 1, 1)

func _process(delta):
    if Input.is_action_just_pressed("animate"):
        current_scale = get_scale()
        scale_accum = Vector3(1, 1, 1)
        state_machine.travel("Animate")
    scale_accum += animation_tree.get_root_motion_scale()
    set_scale(current_scale * scale_accum)
[/gdscript]
[/codeblocks]
*/
//go:nosplit
func (self class) GetRootMotionScale() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_get_root_motion_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Retrieve the blended value of the position tracks with the [member root_motion_track] as a [Vector3] that can be used elsewhere.
This is useful in cases where you want to respect the initial key values of the animation.
For example, if an animation with only one key [code]Vector3(0, 0, 0)[/code] is played in the previous frame and then an animation with only one key [code]Vector3(1, 0, 1)[/code] is played in the next frame, the difference can be calculated as follows:
[codeblocks]
[gdscript]
var prev_root_motion_position_accumulator: Vector3

func _process(delta):
    if Input.is_action_just_pressed("animate"):
        state_machine.travel("Animate")
    var current_root_motion_position_accumulator: Vector3 = animation_tree.get_root_motion_position_accumulator()
    var difference: Vector3 = current_root_motion_position_accumulator - prev_root_motion_position_accumulator
    prev_root_motion_position_accumulator = current_root_motion_position_accumulator
    transform.origin += difference
[/gdscript]
[/codeblocks]
However, if the animation loops, an unintended discrete change may occur, so this is only useful for some simple use cases.
*/
//go:nosplit
func (self class) GetRootMotionPositionAccumulator() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_get_root_motion_position_accumulator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Retrieve the blended value of the rotation tracks with the [member root_motion_track] as a [Quaternion] that can be used elsewhere.
This is necessary to apply the root motion position correctly, taking rotation into account. See also [method get_root_motion_position].
Also, this is useful in cases where you want to respect the initial key values of the animation.
For example, if an animation with only one key [code]Quaternion(0, 0, 0, 1)[/code] is played in the previous frame and then an animation with only one key [code]Quaternion(0, 0.707, 0, 0.707)[/code] is played in the next frame, the difference can be calculated as follows:
[codeblocks]
[gdscript]
var prev_root_motion_rotation_accumulator: Quaternion

func _process(delta):
    if Input.is_action_just_pressed("animate"):
        state_machine.travel("Animate")
    var current_root_motion_rotation_accumulator: Quaternion = animation_tree.get_root_motion_rotation_accumulator()
    var difference: Quaternion = prev_root_motion_rotation_accumulator.inverse() * current_root_motion_rotation_accumulator
    prev_root_motion_rotation_accumulator = current_root_motion_rotation_accumulator
    transform.basis *=  Basis(difference)
[/gdscript]
[/codeblocks]
However, if the animation loops, an unintended discrete change may occur, so this is only useful for some simple use cases.
*/
//go:nosplit
func (self class) GetRootMotionRotationAccumulator() gd.Quaternion {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Quaternion](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_get_root_motion_rotation_accumulator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Retrieve the blended value of the scale tracks with the [member root_motion_track] as a [Vector3] that can be used elsewhere.
For example, if an animation with only one key [code]Vector3(1, 1, 1)[/code] is played in the previous frame and then an animation with only one key [code]Vector3(2, 2, 2)[/code] is played in the next frame, the difference can be calculated as follows:
[codeblocks]
[gdscript]
var prev_root_motion_scale_accumulator: Vector3

func _process(delta):
    if Input.is_action_just_pressed("animate"):
        state_machine.travel("Animate")
    var current_root_motion_scale_accumulator: Vector3 = animation_tree.get_root_motion_scale_accumulator()
    var difference: Vector3 = current_root_motion_scale_accumulator - prev_root_motion_scale_accumulator
    prev_root_motion_scale_accumulator = current_root_motion_scale_accumulator
    transform.basis = transform.basis.scaled(difference)
[/gdscript]
[/codeblocks]
However, if the animation loops, an unintended discrete change may occur, so this is only useful for some simple use cases.
*/
//go:nosplit
func (self class) GetRootMotionScaleAccumulator() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_get_root_motion_scale_accumulator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
[AnimationMixer] caches animated nodes. It may not notice if a node disappears; [method clear_caches] forces it to update the cache again.
*/
//go:nosplit
func (self class) ClearCaches()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_clear_caches, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Manually advance the animations by the specified time (in seconds).
*/
//go:nosplit
func (self class) Advance(delta gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, delta)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_advance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If the animation track specified by [param name] has an option [constant Animation.UPDATE_CAPTURE], stores current values of the objects indicated by the track path as a cache. If there is already a captured cache, the old cache is discarded.
After this it will interpolate with current animation blending result during the playback process for the time specified by [param duration], working like a crossfade.
You can specify [param trans_type] as the curve for the interpolation. For better results, it may be appropriate to specify [constant Tween.TRANS_LINEAR] for cases where the first key of the track begins with a non-zero value or where the key value does not change, and [constant Tween.TRANS_QUAD] for cases where the key value changes linearly.
*/
//go:nosplit
func (self class) Capture(name gd.StringName, duration gd.Float, trans_type classdb.TweenTransitionType, ease_type classdb.TweenEaseType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, duration)
	callframe.Arg(frame, trans_type)
	callframe.Arg(frame, ease_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_capture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetResetOnSaveEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_set_reset_on_save_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsResetOnSaveEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_is_reset_on_save_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the key of [param animation] or an empty [StringName] if not found.
*/
//go:nosplit
func (self class) FindAnimation(ctx gd.Lifetime, animation object.Animation) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(animation[0].AsPointer())[0])
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_find_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the key for the [AnimationLibrary] that contains [param animation] or an empty [StringName] if not found.
*/
//go:nosplit
func (self class) FindAnimationLibrary(ctx gd.Lifetime, animation object.Animation) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(animation[0].AsPointer())[0])
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationMixer.Bind_find_animation_library, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsAnimationMixer() Expert { return self[0].AsAnimationMixer() }


//go:nosplit
func (self Simple) AsAnimationMixer() Simple { return self[0].AsAnimationMixer() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_post_process_key_value": return reflect.ValueOf(self._post_process_key_value);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_post_process_key_value": return reflect.ValueOf(self._post_process_key_value);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("AnimationMixer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type AnimationCallbackModeProcess = classdb.AnimationMixerAnimationCallbackModeProcess

const (
/*Process animation during physics frames (see [constant Node.NOTIFICATION_INTERNAL_PHYSICS_PROCESS]). This is especially useful when animating physics bodies.*/
	AnimationCallbackModeProcessPhysics AnimationCallbackModeProcess = 0
/*Process animation during process frames (see [constant Node.NOTIFICATION_INTERNAL_PROCESS]).*/
	AnimationCallbackModeProcessIdle AnimationCallbackModeProcess = 1
/*Do not process animation. Use [method advance] to process the animation manually.*/
	AnimationCallbackModeProcessManual AnimationCallbackModeProcess = 2
)
type AnimationCallbackModeMethod = classdb.AnimationMixerAnimationCallbackModeMethod

const (
/*Batch method calls during the animation process, then do the calls after events are processed. This avoids bugs involving deleting nodes or modifying the AnimationPlayer while playing.*/
	AnimationCallbackModeMethodDeferred AnimationCallbackModeMethod = 0
/*Make method calls immediately when reached in the animation.*/
	AnimationCallbackModeMethodImmediate AnimationCallbackModeMethod = 1
)
type AnimationCallbackModeDiscrete = classdb.AnimationMixerAnimationCallbackModeDiscrete

const (
/*An [constant Animation.UPDATE_DISCRETE] track value takes precedence when blending [constant Animation.UPDATE_CONTINUOUS] or [constant Animation.UPDATE_CAPTURE] track values and [constant Animation.UPDATE_DISCRETE] track values.*/
	AnimationCallbackModeDiscreteDominant AnimationCallbackModeDiscrete = 0
/*An [constant Animation.UPDATE_CONTINUOUS] or [constant Animation.UPDATE_CAPTURE] track value takes precedence when blending the [constant Animation.UPDATE_CONTINUOUS] or [constant Animation.UPDATE_CAPTURE] track values and the [constant Animation.UPDATE_DISCRETE] track values. This is the default behavior for [AnimationPlayer].*/
	AnimationCallbackModeDiscreteRecessive AnimationCallbackModeDiscrete = 1
/*Always treat the [constant Animation.UPDATE_DISCRETE] track value as [constant Animation.UPDATE_CONTINUOUS] with [constant Animation.INTERPOLATION_NEAREST]. This is the default behavior for [AnimationTree].
If a value track has non-numeric type key values, it is internally converted to use [constant ANIMATION_CALLBACK_MODE_DISCRETE_RECESSIVE] with [constant Animation.UPDATE_DISCRETE].*/
	AnimationCallbackModeDiscreteForceContinuous AnimationCallbackModeDiscrete = 2
)
