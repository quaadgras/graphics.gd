// Package AnimationMixer provides methods for working with AnimationMixer object instances.
package AnimationMixer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/NodePath"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Quaternion"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Base class for [AnimationPlayer] and [AnimationTree] to manage animation lists. It also has general properties and methods for playback and blending.
After instantiating the playback information data within the extended class, the blending is processed by the [AnimationMixer].

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=AnimationMixer)
*/
type Instance [1]gdclass.AnimationMixer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAnimationMixer() Instance
}
type Interface interface {
	//A virtual function for processing after getting a key during playback.
	PostProcessKeyValue(animation [1]gdclass.Animation, track int, value any, object_id int, object_sub_idx int) any
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) PostProcessKeyValue(animation [1]gdclass.Animation, track int, value any, object_id int, object_sub_idx int) (_ any) {
	return
}

/*
A virtual function for processing after getting a key during playback.
*/
func (Instance) _post_process_key_value(impl func(ptr unsafe.Pointer, animation [1]gdclass.Animation, track int, value any, object_id int, object_sub_idx int) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var animation = [1]gdclass.Animation{pointers.New[gdclass.Animation]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 0))})}
		defer pointers.End(animation[0])
		var track = gd.UnsafeGet[gd.Int](p_args, 1)
		var value = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))
		defer pointers.End(value)
		var object_id = gd.UnsafeGet[gd.Int](p_args, 3)
		var object_sub_idx = gd.UnsafeGet[gd.Int](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, animation, int(track), value.Interface(), int(object_id), int(object_sub_idx))
		ptr, ok := pointers.End(gd.NewVariant(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
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
func (self Instance) AddAnimationLibrary(name string, library [1]gdclass.AnimationLibrary) error {
	return error(class(self).AddAnimationLibrary(gd.NewStringName(name), library))
}

/*
Removes the [AnimationLibrary] associated with the key [param name].
*/
func (self Instance) RemoveAnimationLibrary(name string) {
	class(self).RemoveAnimationLibrary(gd.NewStringName(name))
}

/*
Moves the [AnimationLibrary] associated with the key [param name] to the key [param newname].
*/
func (self Instance) RenameAnimationLibrary(name string, newname string) {
	class(self).RenameAnimationLibrary(gd.NewStringName(name), gd.NewStringName(newname))
}

/*
Returns [code]true[/code] if the [AnimationMixer] stores an [AnimationLibrary] with key [param name].
*/
func (self Instance) HasAnimationLibrary(name string) bool {
	return bool(class(self).HasAnimationLibrary(gd.NewStringName(name)))
}

/*
Returns the first [AnimationLibrary] with key [param name] or [code]null[/code] if not found.
To get the [AnimationMixer]'s global animation library, use [code]get_animation_library("")[/code].
*/
func (self Instance) GetAnimationLibrary(name string) [1]gdclass.AnimationLibrary {
	return [1]gdclass.AnimationLibrary(class(self).GetAnimationLibrary(gd.NewStringName(name)))
}

/*
Returns the list of stored library keys.
*/
func (self Instance) GetAnimationLibraryList() gd.Array {
	return gd.Array(class(self).GetAnimationLibraryList())
}

/*
Returns [code]true[/code] if the [AnimationMixer] stores an [Animation] with key [param name].
*/
func (self Instance) HasAnimation(name string) bool {
	return bool(class(self).HasAnimation(gd.NewStringName(name)))
}

/*
Returns the [Animation] with the key [param name]. If the animation does not exist, [code]null[/code] is returned and an error is logged.
*/
func (self Instance) GetAnimation(name string) [1]gdclass.Animation {
	return [1]gdclass.Animation(class(self).GetAnimation(gd.NewStringName(name)))
}

/*
Returns the list of stored animation keys.
*/
func (self Instance) GetAnimationList() []string {
	return []string(class(self).GetAnimationList().Strings())
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
func (self Instance) GetRootMotionPosition() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetRootMotionPosition())
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
func (self Instance) GetRootMotionRotation() Quaternion.IJKX {
	return Quaternion.IJKX(class(self).GetRootMotionRotation())
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
func (self Instance) GetRootMotionScale() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetRootMotionScale())
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
func (self Instance) GetRootMotionPositionAccumulator() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetRootMotionPositionAccumulator())
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
func (self Instance) GetRootMotionRotationAccumulator() Quaternion.IJKX {
	return Quaternion.IJKX(class(self).GetRootMotionRotationAccumulator())
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
func (self Instance) GetRootMotionScaleAccumulator() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetRootMotionScaleAccumulator())
}

/*
[AnimationMixer] caches animated nodes. It may not notice if a node disappears; [method clear_caches] forces it to update the cache again.
*/
func (self Instance) ClearCaches() {
	class(self).ClearCaches()
}

/*
Manually advance the animations by the specified time (in seconds).
*/
func (self Instance) Advance(delta Float.X) {
	class(self).Advance(gd.Float(delta))
}

/*
If the animation track specified by [param name] has an option [constant Animation.UPDATE_CAPTURE], stores current values of the objects indicated by the track path as a cache. If there is already a captured cache, the old cache is discarded.
After this it will interpolate with current animation blending result during the playback process for the time specified by [param duration], working like a crossfade.
You can specify [param trans_type] as the curve for the interpolation. For better results, it may be appropriate to specify [constant Tween.TRANS_LINEAR] for cases where the first key of the track begins with a non-zero value or where the key value does not change, and [constant Tween.TRANS_QUAD] for cases where the key value changes linearly.
*/
func (self Instance) Capture(name string, duration Float.X) {
	class(self).Capture(gd.NewStringName(name), gd.Float(duration), 0, 0)
}

/*
Returns the key of [param animation] or an empty [StringName] if not found.
*/
func (self Instance) FindAnimation(animation [1]gdclass.Animation) string {
	return string(class(self).FindAnimation(animation).String())
}

/*
Returns the key for the [AnimationLibrary] that contains [param animation] or an empty [StringName] if not found.
*/
func (self Instance) FindAnimationLibrary(animation [1]gdclass.Animation) string {
	return string(class(self).FindAnimationLibrary(animation).String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AnimationMixer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationMixer"))
	casted := Instance{*(*gdclass.AnimationMixer)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Active() bool {
	return bool(class(self).IsActive())
}

func (self Instance) SetActive(value bool) {
	class(self).SetActive(value)
}

func (self Instance) Deterministic() bool {
	return bool(class(self).IsDeterministic())
}

func (self Instance) SetDeterministic(value bool) {
	class(self).SetDeterministic(value)
}

func (self Instance) ResetOnSave() bool {
	return bool(class(self).IsResetOnSaveEnabled())
}

func (self Instance) SetResetOnSave(value bool) {
	class(self).SetResetOnSaveEnabled(value)
}

func (self Instance) RootNode() NodePath.String {
	return NodePath.String(class(self).GetRootNode().String())
}

func (self Instance) SetRootNode(value NodePath.String) {
	class(self).SetRootNode(gd.NewString(string(value)).NodePath())
}

func (self Instance) RootMotionTrack() NodePath.String {
	return NodePath.String(class(self).GetRootMotionTrack().String())
}

func (self Instance) SetRootMotionTrack(value NodePath.String) {
	class(self).SetRootMotionTrack(gd.NewString(string(value)).NodePath())
}

func (self Instance) AudioMaxPolyphony() int {
	return int(int(class(self).GetAudioMaxPolyphony()))
}

func (self Instance) SetAudioMaxPolyphony(value int) {
	class(self).SetAudioMaxPolyphony(gd.Int(value))
}

func (self Instance) CallbackModeProcess() gdclass.AnimationMixerAnimationCallbackModeProcess {
	return gdclass.AnimationMixerAnimationCallbackModeProcess(class(self).GetCallbackModeProcess())
}

func (self Instance) SetCallbackModeProcess(value gdclass.AnimationMixerAnimationCallbackModeProcess) {
	class(self).SetCallbackModeProcess(value)
}

func (self Instance) CallbackModeMethod() gdclass.AnimationMixerAnimationCallbackModeMethod {
	return gdclass.AnimationMixerAnimationCallbackModeMethod(class(self).GetCallbackModeMethod())
}

func (self Instance) SetCallbackModeMethod(value gdclass.AnimationMixerAnimationCallbackModeMethod) {
	class(self).SetCallbackModeMethod(value)
}

func (self Instance) CallbackModeDiscrete() gdclass.AnimationMixerAnimationCallbackModeDiscrete {
	return gdclass.AnimationMixerAnimationCallbackModeDiscrete(class(self).GetCallbackModeDiscrete())
}

func (self Instance) SetCallbackModeDiscrete(value gdclass.AnimationMixerAnimationCallbackModeDiscrete) {
	class(self).SetCallbackModeDiscrete(value)
}

/*
A virtual function for processing after getting a key during playback.
*/
func (class) _post_process_key_value(impl func(ptr unsafe.Pointer, animation [1]gdclass.Animation, track gd.Int, value gd.Variant, object_id gd.Int, object_sub_idx gd.Int) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var animation = [1]gdclass.Animation{pointers.New[gdclass.Animation]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(animation[0])
		var track = gd.UnsafeGet[gd.Int](p_args, 1)
		var value = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))
		var object_id = gd.UnsafeGet[gd.Int](p_args, 3)
		var object_sub_idx = gd.UnsafeGet[gd.Int](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, animation, track, value, object_id, object_sub_idx)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
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
func (self class) AddAnimationLibrary(name gd.StringName, library [1]gdclass.AnimationLibrary) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(library[0])[0])
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_add_animation_library, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the [AnimationLibrary] associated with the key [param name].
*/
//go:nosplit
func (self class) RemoveAnimationLibrary(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_remove_animation_library, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Moves the [AnimationLibrary] associated with the key [param name] to the key [param newname].
*/
//go:nosplit
func (self class) RenameAnimationLibrary(name gd.StringName, newname gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(newname))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_rename_animation_library, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the [AnimationMixer] stores an [AnimationLibrary] with key [param name].
*/
//go:nosplit
func (self class) HasAnimationLibrary(name gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_has_animation_library, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the first [AnimationLibrary] with key [param name] or [code]null[/code] if not found.
To get the [AnimationMixer]'s global animation library, use [code]get_animation_library("")[/code].
*/
//go:nosplit
func (self class) GetAnimationLibrary(name gd.StringName) [1]gdclass.AnimationLibrary {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_animation_library, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.AnimationLibrary{gd.PointerWithOwnershipTransferredToGo[gdclass.AnimationLibrary](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the list of stored library keys.
*/
//go:nosplit
func (self class) GetAnimationLibraryList() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_animation_library_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the [AnimationMixer] stores an [Animation] with key [param name].
*/
//go:nosplit
func (self class) HasAnimation(name gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_has_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Animation] with the key [param name]. If the animation does not exist, [code]null[/code] is returned and an error is logged.
*/
//go:nosplit
func (self class) GetAnimation(name gd.StringName) [1]gdclass.Animation {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Animation{gd.PointerWithOwnershipTransferredToGo[gdclass.Animation](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the list of stored animation keys.
*/
//go:nosplit
func (self class) GetAnimationList() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_animation_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetActive(active bool) {
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_set_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsActive() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_is_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDeterministic(deterministic bool) {
	var frame = callframe.New()
	callframe.Arg(frame, deterministic)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_set_deterministic, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDeterministic() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_is_deterministic, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRootNode(path gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_set_root_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRootNode() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_root_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCallbackModeProcess(mode gdclass.AnimationMixerAnimationCallbackModeProcess) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_set_callback_mode_process, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCallbackModeProcess() gdclass.AnimationMixerAnimationCallbackModeProcess {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AnimationMixerAnimationCallbackModeProcess](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_callback_mode_process, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCallbackModeMethod(mode gdclass.AnimationMixerAnimationCallbackModeMethod) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_set_callback_mode_method, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCallbackModeMethod() gdclass.AnimationMixerAnimationCallbackModeMethod {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AnimationMixerAnimationCallbackModeMethod](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_callback_mode_method, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCallbackModeDiscrete(mode gdclass.AnimationMixerAnimationCallbackModeDiscrete) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_set_callback_mode_discrete, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCallbackModeDiscrete() gdclass.AnimationMixerAnimationCallbackModeDiscrete {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AnimationMixerAnimationCallbackModeDiscrete](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_callback_mode_discrete, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAudioMaxPolyphony(max_polyphony gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, max_polyphony)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_set_audio_max_polyphony, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAudioMaxPolyphony() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_audio_max_polyphony, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRootMotionTrack(path gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_set_root_motion_track, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRootMotionTrack() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_root_motion_track, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_root_motion_position, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Quaternion](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_root_motion_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_root_motion_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_root_motion_position_accumulator, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Quaternion](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_root_motion_rotation_accumulator, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_root_motion_scale_accumulator, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
[AnimationMixer] caches animated nodes. It may not notice if a node disappears; [method clear_caches] forces it to update the cache again.
*/
//go:nosplit
func (self class) ClearCaches() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_clear_caches, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Manually advance the animations by the specified time (in seconds).
*/
//go:nosplit
func (self class) Advance(delta gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, delta)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_advance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
If the animation track specified by [param name] has an option [constant Animation.UPDATE_CAPTURE], stores current values of the objects indicated by the track path as a cache. If there is already a captured cache, the old cache is discarded.
After this it will interpolate with current animation blending result during the playback process for the time specified by [param duration], working like a crossfade.
You can specify [param trans_type] as the curve for the interpolation. For better results, it may be appropriate to specify [constant Tween.TRANS_LINEAR] for cases where the first key of the track begins with a non-zero value or where the key value does not change, and [constant Tween.TRANS_QUAD] for cases where the key value changes linearly.
*/
//go:nosplit
func (self class) Capture(name gd.StringName, duration gd.Float, trans_type gdclass.TweenTransitionType, ease_type gdclass.TweenEaseType) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, duration)
	callframe.Arg(frame, trans_type)
	callframe.Arg(frame, ease_type)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_capture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetResetOnSaveEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_set_reset_on_save_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsResetOnSaveEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_is_reset_on_save_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the key of [param animation] or an empty [StringName] if not found.
*/
//go:nosplit
func (self class) FindAnimation(animation [1]gdclass.Animation) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(animation[0])[0])
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_find_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the key for the [AnimationLibrary] that contains [param animation] or an empty [StringName] if not found.
*/
//go:nosplit
func (self class) FindAnimationLibrary(animation [1]gdclass.Animation) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(animation[0])[0])
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_find_animation_library, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
func (self Instance) OnAnimationListChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("animation_list_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAnimationLibrariesUpdated(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("animation_libraries_updated"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAnimationFinished(cb func(anim_name string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("animation_finished"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAnimationStarted(cb func(anim_name string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("animation_started"), gd.NewCallable(cb), 0)
}

func (self Instance) OnCachesCleared(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("caches_cleared"), gd.NewCallable(cb), 0)
}

func (self Instance) OnMixerApplied(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("mixer_applied"), gd.NewCallable(cb), 0)
}

func (self Instance) OnMixerUpdated(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("mixer_updated"), gd.NewCallable(cb), 0)
}

func (self class) AsAnimationMixer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimationMixer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced         { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance      { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_post_process_key_value":
		return reflect.ValueOf(self._post_process_key_value)
	default:
		return gd.VirtualByName(Node.Advanced(self.AsNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_post_process_key_value":
		return reflect.ValueOf(self._post_process_key_value)
	default:
		return gd.VirtualByName(Node.Instance(self.AsNode()), name)
	}
}
func init() {
	gdclass.Register("AnimationMixer", func(ptr gd.Object) any {
		return [1]gdclass.AnimationMixer{*(*gdclass.AnimationMixer)(unsafe.Pointer(&ptr))}
	})
}

type AnimationCallbackModeProcess = gdclass.AnimationMixerAnimationCallbackModeProcess

const (
	/*Process animation during physics frames (see [constant Node.NOTIFICATION_INTERNAL_PHYSICS_PROCESS]). This is especially useful when animating physics bodies.*/
	AnimationCallbackModeProcessPhysics AnimationCallbackModeProcess = 0
	/*Process animation during process frames (see [constant Node.NOTIFICATION_INTERNAL_PROCESS]).*/
	AnimationCallbackModeProcessIdle AnimationCallbackModeProcess = 1
	/*Do not process animation. Use [method advance] to process the animation manually.*/
	AnimationCallbackModeProcessManual AnimationCallbackModeProcess = 2
)

type AnimationCallbackModeMethod = gdclass.AnimationMixerAnimationCallbackModeMethod

const (
	/*Batch method calls during the animation process, then do the calls after events are processed. This avoids bugs involving deleting nodes or modifying the AnimationPlayer while playing.*/
	AnimationCallbackModeMethodDeferred AnimationCallbackModeMethod = 0
	/*Make method calls immediately when reached in the animation.*/
	AnimationCallbackModeMethodImmediate AnimationCallbackModeMethod = 1
)

type AnimationCallbackModeDiscrete = gdclass.AnimationMixerAnimationCallbackModeDiscrete

const (
	/*An [constant Animation.UPDATE_DISCRETE] track value takes precedence when blending [constant Animation.UPDATE_CONTINUOUS] or [constant Animation.UPDATE_CAPTURE] track values and [constant Animation.UPDATE_DISCRETE] track values.*/
	AnimationCallbackModeDiscreteDominant AnimationCallbackModeDiscrete = 0
	/*An [constant Animation.UPDATE_CONTINUOUS] or [constant Animation.UPDATE_CAPTURE] track value takes precedence when blending the [constant Animation.UPDATE_CONTINUOUS] or [constant Animation.UPDATE_CAPTURE] track values and the [constant Animation.UPDATE_DISCRETE] track values. This is the default behavior for [AnimationPlayer].*/
	AnimationCallbackModeDiscreteRecessive AnimationCallbackModeDiscrete = 1
	/*Always treat the [constant Animation.UPDATE_DISCRETE] track value as [constant Animation.UPDATE_CONTINUOUS] with [constant Animation.INTERPOLATION_NEAREST]. This is the default behavior for [AnimationTree].
	  If a value track has non-numeric type key values, it is internally converted to use [constant ANIMATION_CALLBACK_MODE_DISCRETE_RECESSIVE] with [constant Animation.UPDATE_DISCRETE].*/
	AnimationCallbackModeDiscreteForceContinuous AnimationCallbackModeDiscrete = 2
)

type Error = gd.Error

const (
	/*Methods that return [enum Error] return [constant OK] when no error occurred.
	  Since [constant OK] has value 0, and all other error constants are positive integers, it can also be used in boolean checks.
	  [b]Example:[/b]
	  [codeblock]
	  var error = method_that_returns_error()
	  if error != OK:
	      printerr("Failure!")

	  # Or, alternatively:
	  if error:
	      printerr("Still failing!")
	  [/codeblock]
	  [b]Note:[/b] Many functions do not return an error code, but will print error messages to standard output.*/
	Ok Error = 0
	/*Generic error.*/
	Failed Error = 1
	/*Unavailable error.*/
	ErrUnavailable Error = 2
	/*Unconfigured error.*/
	ErrUnconfigured Error = 3
	/*Unauthorized error.*/
	ErrUnauthorized Error = 4
	/*Parameter range error.*/
	ErrParameterRangeError Error = 5
	/*Out of memory (OOM) error.*/
	ErrOutOfMemory Error = 6
	/*File: Not found error.*/
	ErrFileNotFound Error = 7
	/*File: Bad drive error.*/
	ErrFileBadDrive Error = 8
	/*File: Bad path error.*/
	ErrFileBadPath Error = 9
	/*File: No permission error.*/
	ErrFileNoPermission Error = 10
	/*File: Already in use error.*/
	ErrFileAlreadyInUse Error = 11
	/*File: Can't open error.*/
	ErrFileCantOpen Error = 12
	/*File: Can't write error.*/
	ErrFileCantWrite Error = 13
	/*File: Can't read error.*/
	ErrFileCantRead Error = 14
	/*File: Unrecognized error.*/
	ErrFileUnrecognized Error = 15
	/*File: Corrupt error.*/
	ErrFileCorrupt Error = 16
	/*File: Missing dependencies error.*/
	ErrFileMissingDependencies Error = 17
	/*File: End of file (EOF) error.*/
	ErrFileEof Error = 18
	/*Can't open error.*/
	ErrCantOpen Error = 19
	/*Can't create error.*/
	ErrCantCreate Error = 20
	/*Query failed error.*/
	ErrQueryFailed Error = 21
	/*Already in use error.*/
	ErrAlreadyInUse Error = 22
	/*Locked error.*/
	ErrLocked Error = 23
	/*Timeout error.*/
	ErrTimeout Error = 24
	/*Can't connect error.*/
	ErrCantConnect Error = 25
	/*Can't resolve error.*/
	ErrCantResolve Error = 26
	/*Connection error.*/
	ErrConnectionError Error = 27
	/*Can't acquire resource error.*/
	ErrCantAcquireResource Error = 28
	/*Can't fork process error.*/
	ErrCantFork Error = 29
	/*Invalid data error.*/
	ErrInvalidData Error = 30
	/*Invalid parameter error.*/
	ErrInvalidParameter Error = 31
	/*Already exists error.*/
	ErrAlreadyExists Error = 32
	/*Does not exist error.*/
	ErrDoesNotExist Error = 33
	/*Database: Read error.*/
	ErrDatabaseCantRead Error = 34
	/*Database: Write error.*/
	ErrDatabaseCantWrite Error = 35
	/*Compilation failed error.*/
	ErrCompilationFailed Error = 36
	/*Method not found error.*/
	ErrMethodNotFound Error = 37
	/*Linking failed error.*/
	ErrLinkFailed Error = 38
	/*Script failed error.*/
	ErrScriptFailed Error = 39
	/*Cycling link (import cycle) error.*/
	ErrCyclicLink Error = 40
	/*Invalid declaration error.*/
	ErrInvalidDeclaration Error = 41
	/*Duplicate symbol error.*/
	ErrDuplicateSymbol Error = 42
	/*Parse error.*/
	ErrParseError Error = 43
	/*Busy error.*/
	ErrBusy Error = 44
	/*Skip error.*/
	ErrSkip Error = 45
	/*Help error. Used internally when passing [code]--version[/code] or [code]--help[/code] as executable options.*/
	ErrHelp Error = 46
	/*Bug error, caused by an implementation issue in the method.
	  [b]Note:[/b] If a built-in method returns this code, please open an issue on [url=https://github.com/godotengine/godot/issues]the GitHub Issue Tracker[/url].*/
	ErrBug Error = 47
	/*Printer on fire error (This is an easter egg, no built-in methods return this error code).*/
	ErrPrinterOnFire Error = 48
)
