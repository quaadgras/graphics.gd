// Package AnimationMixer provides methods for working with AnimationMixer object instances.
package AnimationMixer

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Quaternion"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Vector3"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

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
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var animation = [1]gdclass.Animation{pointers.New[gdclass.Animation]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(animation[0])
		var track = gd.UnsafeGet[int64](p_args, 1)

		var value = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))))
		defer pointers.End(gd.InternalVariant(value))
		var object_id = gd.UnsafeGet[int64](p_args, 3)

		var object_sub_idx = gd.UnsafeGet[int64](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, animation, int(track), value.Interface(), int(object_id), int(object_sub_idx))
		ptr, ok := pointers.End(gd.InternalVariant(variant.New(ret)))

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
func (self Instance) AddAnimationLibrary(name string, library [1]gdclass.AnimationLibrary) error { //gd:AnimationMixer.add_animation_library
	return error(gd.ToError(class(self).AddAnimationLibrary(String.Name(String.New(name)), library)))
}

/*
Removes the [AnimationLibrary] associated with the key [param name].
*/
func (self Instance) RemoveAnimationLibrary(name string) { //gd:AnimationMixer.remove_animation_library
	class(self).RemoveAnimationLibrary(String.Name(String.New(name)))
}

/*
Moves the [AnimationLibrary] associated with the key [param name] to the key [param newname].
*/
func (self Instance) RenameAnimationLibrary(name string, newname string) { //gd:AnimationMixer.rename_animation_library
	class(self).RenameAnimationLibrary(String.Name(String.New(name)), String.Name(String.New(newname)))
}

/*
Returns [code]true[/code] if the [AnimationMixer] stores an [AnimationLibrary] with key [param name].
*/
func (self Instance) HasAnimationLibrary(name string) bool { //gd:AnimationMixer.has_animation_library
	return bool(class(self).HasAnimationLibrary(String.Name(String.New(name))))
}

/*
Returns the first [AnimationLibrary] with key [param name] or [code]null[/code] if not found.
To get the [AnimationMixer]'s global animation library, use [code]get_animation_library("")[/code].
*/
func (self Instance) GetAnimationLibrary(name string) [1]gdclass.AnimationLibrary { //gd:AnimationMixer.get_animation_library
	return [1]gdclass.AnimationLibrary(class(self).GetAnimationLibrary(String.Name(String.New(name))))
}

/*
Returns the list of stored library keys.
*/
func (self Instance) GetAnimationLibraryList() []string { //gd:AnimationMixer.get_animation_library_list
	return []string(gd.ArrayAs[[]string](gd.InternalArray(class(self).GetAnimationLibraryList())))
}

/*
Returns [code]true[/code] if the [AnimationMixer] stores an [Animation] with key [param name].
*/
func (self Instance) HasAnimation(name string) bool { //gd:AnimationMixer.has_animation
	return bool(class(self).HasAnimation(String.Name(String.New(name))))
}

/*
Returns the [Animation] with the key [param name]. If the animation does not exist, [code]null[/code] is returned and an error is logged.
*/
func (self Instance) GetAnimation(name string) [1]gdclass.Animation { //gd:AnimationMixer.get_animation
	return [1]gdclass.Animation(class(self).GetAnimation(String.Name(String.New(name))))
}

/*
Returns the list of stored animation keys.
*/
func (self Instance) GetAnimationList() []string { //gd:AnimationMixer.get_animation_list
	return []string(class(self).GetAnimationList().Strings())
}

/*
Retrieve the motion delta of position with the [member root_motion_track] as a [Vector3] that can be used elsewhere.
If [member root_motion_track] is not a path to a track of type [constant Animation.TYPE_POSITION_3D], returns [code]Vector3(0, 0, 0)[/code].
See also [member root_motion_track] and [RootMotionView].
The most basic example is applying position to [CharacterBody3D]:
[codeblocks]
[gdscript]
var current_rotation

func _process(delta):

	if Input.is_action_just_pressed("animate"):
	    current_rotation = get_quaternion()
	    state_machine.travel("Animate")
	var velocity = current_rotation * animation_tree.get_root_motion_position() / delta
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
	var velocity = (animation_tree.get_root_motion_rotation_accumulator().inverse() * get_quaternion()) * animation_tree.get_root_motion_position() / delta
	set_velocity(velocity)
	move_and_slide()

[/gdscript]
[/codeblocks]
If [member root_motion_local] is [code]true[/code], return the pre-multiplied translation value with the inverted rotation.
In this case, the code can be written as follows:
[codeblocks]
[gdscript]
func _process(delta):

	if Input.is_action_just_pressed("animate"):
	    state_machine.travel("Animate")
	set_quaternion(get_quaternion() * animation_tree.get_root_motion_rotation())
	var velocity = get_quaternion() * animation_tree.get_root_motion_position() / delta
	set_velocity(velocity)
	move_and_slide()

[/gdscript]
[/codeblocks]
*/
func (self Instance) GetRootMotionPosition() Vector3.XYZ { //gd:AnimationMixer.get_root_motion_position
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
func (self Instance) GetRootMotionRotation() Quaternion.IJKX { //gd:AnimationMixer.get_root_motion_rotation
	return Quaternion.IJKX(class(self).GetRootMotionRotation())
}

/*
Retrieve the motion delta of scale with the [member root_motion_track] as a [Vector3] that can be used elsewhere.
If [member root_motion_track] is not a path to a track of type [constant Animation.TYPE_SCALE_3D], returns [code]Vector3(0, 0, 0)[/code].
See also [member root_motion_track] and [RootMotionView].
The most basic example is applying scale to [CharacterBody3D]:
[codeblocks]
[gdscript]
var current_scale = Vector3(1, 1, 1)
var scale_accum = Vector3(1, 1, 1)

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
func (self Instance) GetRootMotionScale() Vector3.XYZ { //gd:AnimationMixer.get_root_motion_scale
	return Vector3.XYZ(class(self).GetRootMotionScale())
}

/*
Retrieve the blended value of the position tracks with the [member root_motion_track] as a [Vector3] that can be used elsewhere.
This is useful in cases where you want to respect the initial key values of the animation.
For example, if an animation with only one key [code]Vector3(0, 0, 0)[/code] is played in the previous frame and then an animation with only one key [code]Vector3(1, 0, 1)[/code] is played in the next frame, the difference can be calculated as follows:
[codeblocks]
[gdscript]
var prev_root_motion_position_accumulator

func _process(delta):

	if Input.is_action_just_pressed("animate"):
	    state_machine.travel("Animate")
	var current_root_motion_position_accumulator = animation_tree.get_root_motion_position_accumulator()
	var difference = current_root_motion_position_accumulator - prev_root_motion_position_accumulator
	prev_root_motion_position_accumulator = current_root_motion_position_accumulator
	transform.origin += difference

[/gdscript]
[/codeblocks]
However, if the animation loops, an unintended discrete change may occur, so this is only useful for some simple use cases.
*/
func (self Instance) GetRootMotionPositionAccumulator() Vector3.XYZ { //gd:AnimationMixer.get_root_motion_position_accumulator
	return Vector3.XYZ(class(self).GetRootMotionPositionAccumulator())
}

/*
Retrieve the blended value of the rotation tracks with the [member root_motion_track] as a [Quaternion] that can be used elsewhere.
This is necessary to apply the root motion position correctly, taking rotation into account. See also [method get_root_motion_position].
Also, this is useful in cases where you want to respect the initial key values of the animation.
For example, if an animation with only one key [code]Quaternion(0, 0, 0, 1)[/code] is played in the previous frame and then an animation with only one key [code]Quaternion(0, 0.707, 0, 0.707)[/code] is played in the next frame, the difference can be calculated as follows:
[codeblocks]
[gdscript]
var prev_root_motion_rotation_accumulator

func _process(delta):

	if Input.is_action_just_pressed("animate"):
	    state_machine.travel("Animate")
	var current_root_motion_rotation_accumulator = animation_tree.get_root_motion_rotation_accumulator()
	var difference = prev_root_motion_rotation_accumulator.inverse() * current_root_motion_rotation_accumulator
	prev_root_motion_rotation_accumulator = current_root_motion_rotation_accumulator
	transform.basis *=  Basis(difference)

[/gdscript]
[/codeblocks]
However, if the animation loops, an unintended discrete change may occur, so this is only useful for some simple use cases.
*/
func (self Instance) GetRootMotionRotationAccumulator() Quaternion.IJKX { //gd:AnimationMixer.get_root_motion_rotation_accumulator
	return Quaternion.IJKX(class(self).GetRootMotionRotationAccumulator())
}

/*
Retrieve the blended value of the scale tracks with the [member root_motion_track] as a [Vector3] that can be used elsewhere.
For example, if an animation with only one key [code]Vector3(1, 1, 1)[/code] is played in the previous frame and then an animation with only one key [code]Vector3(2, 2, 2)[/code] is played in the next frame, the difference can be calculated as follows:
[codeblocks]
[gdscript]
var prev_root_motion_scale_accumulator

func _process(delta):

	if Input.is_action_just_pressed("animate"):
	    state_machine.travel("Animate")
	var current_root_motion_scale_accumulator = animation_tree.get_root_motion_scale_accumulator()
	var difference = current_root_motion_scale_accumulator - prev_root_motion_scale_accumulator
	prev_root_motion_scale_accumulator = current_root_motion_scale_accumulator
	transform.basis = transform.basis.scaled(difference)

[/gdscript]
[/codeblocks]
However, if the animation loops, an unintended discrete change may occur, so this is only useful for some simple use cases.
*/
func (self Instance) GetRootMotionScaleAccumulator() Vector3.XYZ { //gd:AnimationMixer.get_root_motion_scale_accumulator
	return Vector3.XYZ(class(self).GetRootMotionScaleAccumulator())
}

/*
[AnimationMixer] caches animated nodes. It may not notice if a node disappears; [method clear_caches] forces it to update the cache again.
*/
func (self Instance) ClearCaches() { //gd:AnimationMixer.clear_caches
	class(self).ClearCaches()
}

/*
Manually advance the animations by the specified time (in seconds).
*/
func (self Instance) Advance(delta Float.X) { //gd:AnimationMixer.advance
	class(self).Advance(float64(delta))
}

/*
If the animation track specified by [param name] has an option [constant Animation.UPDATE_CAPTURE], stores current values of the objects indicated by the track path as a cache. If there is already a captured cache, the old cache is discarded.
After this it will interpolate with current animation blending result during the playback process for the time specified by [param duration], working like a crossfade.
You can specify [param trans_type] as the curve for the interpolation. For better results, it may be appropriate to specify [constant Tween.TRANS_LINEAR] for cases where the first key of the track begins with a non-zero value or where the key value does not change, and [constant Tween.TRANS_QUAD] for cases where the key value changes linearly.
*/
func (self Instance) Capture(name string, duration Float.X) { //gd:AnimationMixer.capture
	class(self).Capture(String.Name(String.New(name)), float64(duration), 0, 0)
}

/*
Returns the key of [param animation] or an empty [StringName] if not found.
*/
func (self Instance) FindAnimation(animation [1]gdclass.Animation) string { //gd:AnimationMixer.find_animation
	return string(class(self).FindAnimation(animation).String())
}

/*
Returns the key for the [AnimationLibrary] that contains [param animation] or an empty [StringName] if not found.
*/
func (self Instance) FindAnimationLibrary(animation [1]gdclass.Animation) string { //gd:AnimationMixer.find_animation_library
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

func (self Instance) RootNode() string {
	return string(class(self).GetRootNode().String())
}

func (self Instance) SetRootNode(value string) {
	class(self).SetRootNode(Path.ToNode(String.New(value)))
}

func (self Instance) RootMotionTrack() string {
	return string(class(self).GetRootMotionTrack().String())
}

func (self Instance) SetRootMotionTrack(value string) {
	class(self).SetRootMotionTrack(Path.ToNode(String.New(value)))
}

func (self Instance) RootMotionLocal() bool {
	return bool(class(self).IsRootMotionLocal())
}

func (self Instance) SetRootMotionLocal(value bool) {
	class(self).SetRootMotionLocal(value)
}

func (self Instance) AudioMaxPolyphony() int {
	return int(int(class(self).GetAudioMaxPolyphony()))
}

func (self Instance) SetAudioMaxPolyphony(value int) {
	class(self).SetAudioMaxPolyphony(int64(value))
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
func (class) _post_process_key_value(impl func(ptr unsafe.Pointer, animation [1]gdclass.Animation, track int64, value variant.Any, object_id int64, object_sub_idx int64) variant.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var animation = [1]gdclass.Animation{pointers.New[gdclass.Animation]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(animation[0])
		var track = gd.UnsafeGet[int64](p_args, 1)

		var value = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))))
		defer pointers.End(gd.InternalVariant(value))
		var object_id = gd.UnsafeGet[int64](p_args, 3)

		var object_sub_idx = gd.UnsafeGet[int64](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, animation, track, value, object_id, object_sub_idx)
		ptr, ok := pointers.End(gd.InternalVariant(ret))

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
func (self class) AddAnimationLibrary(name String.Name, library [1]gdclass.AnimationLibrary) Error.Code { //gd:AnimationMixer.add_animation_library
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(library[0])[0])
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_add_animation_library, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Removes the [AnimationLibrary] associated with the key [param name].
*/
//go:nosplit
func (self class) RemoveAnimationLibrary(name String.Name) { //gd:AnimationMixer.remove_animation_library
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_remove_animation_library, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Moves the [AnimationLibrary] associated with the key [param name] to the key [param newname].
*/
//go:nosplit
func (self class) RenameAnimationLibrary(name String.Name, newname String.Name) { //gd:AnimationMixer.rename_animation_library
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(newname)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_rename_animation_library, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the [AnimationMixer] stores an [AnimationLibrary] with key [param name].
*/
//go:nosplit
func (self class) HasAnimationLibrary(name String.Name) bool { //gd:AnimationMixer.has_animation_library
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
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
func (self class) GetAnimationLibrary(name String.Name) [1]gdclass.AnimationLibrary { //gd:AnimationMixer.get_animation_library
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
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
func (self class) GetAnimationLibraryList() Array.Contains[String.Name] { //gd:AnimationMixer.get_animation_library_list
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_animation_library_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[String.Name]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the [AnimationMixer] stores an [Animation] with key [param name].
*/
//go:nosplit
func (self class) HasAnimation(name String.Name) bool { //gd:AnimationMixer.has_animation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
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
func (self class) GetAnimation(name String.Name) [1]gdclass.Animation { //gd:AnimationMixer.get_animation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
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
func (self class) GetAnimationList() Packed.Strings { //gd:AnimationMixer.get_animation_list
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_animation_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetActive(active bool) { //gd:AnimationMixer.set_active
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_set_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsActive() bool { //gd:AnimationMixer.is_active
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_is_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDeterministic(deterministic bool) { //gd:AnimationMixer.set_deterministic
	var frame = callframe.New()
	callframe.Arg(frame, deterministic)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_set_deterministic, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDeterministic() bool { //gd:AnimationMixer.is_deterministic
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_is_deterministic, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRootNode(path Path.ToNode) { //gd:AnimationMixer.set_root_node
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_set_root_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRootNode() Path.ToNode { //gd:AnimationMixer.get_root_node
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_root_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCallbackModeProcess(mode gdclass.AnimationMixerAnimationCallbackModeProcess) { //gd:AnimationMixer.set_callback_mode_process
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_set_callback_mode_process, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCallbackModeProcess() gdclass.AnimationMixerAnimationCallbackModeProcess { //gd:AnimationMixer.get_callback_mode_process
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AnimationMixerAnimationCallbackModeProcess](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_callback_mode_process, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCallbackModeMethod(mode gdclass.AnimationMixerAnimationCallbackModeMethod) { //gd:AnimationMixer.set_callback_mode_method
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_set_callback_mode_method, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCallbackModeMethod() gdclass.AnimationMixerAnimationCallbackModeMethod { //gd:AnimationMixer.get_callback_mode_method
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AnimationMixerAnimationCallbackModeMethod](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_callback_mode_method, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCallbackModeDiscrete(mode gdclass.AnimationMixerAnimationCallbackModeDiscrete) { //gd:AnimationMixer.set_callback_mode_discrete
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_set_callback_mode_discrete, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCallbackModeDiscrete() gdclass.AnimationMixerAnimationCallbackModeDiscrete { //gd:AnimationMixer.get_callback_mode_discrete
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AnimationMixerAnimationCallbackModeDiscrete](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_callback_mode_discrete, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAudioMaxPolyphony(max_polyphony int64) { //gd:AnimationMixer.set_audio_max_polyphony
	var frame = callframe.New()
	callframe.Arg(frame, max_polyphony)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_set_audio_max_polyphony, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAudioMaxPolyphony() int64 { //gd:AnimationMixer.get_audio_max_polyphony
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_audio_max_polyphony, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRootMotionTrack(path Path.ToNode) { //gd:AnimationMixer.set_root_motion_track
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_set_root_motion_track, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRootMotionTrack() Path.ToNode { //gd:AnimationMixer.get_root_motion_track
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_root_motion_track, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRootMotionLocal(enabled bool) { //gd:AnimationMixer.set_root_motion_local
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_set_root_motion_local, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsRootMotionLocal() bool { //gd:AnimationMixer.is_root_motion_local
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_is_root_motion_local, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
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
var current_rotation

func _process(delta):
    if Input.is_action_just_pressed("animate"):
        current_rotation = get_quaternion()
        state_machine.travel("Animate")
    var velocity = current_rotation * animation_tree.get_root_motion_position() / delta
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
    var velocity = (animation_tree.get_root_motion_rotation_accumulator().inverse() * get_quaternion()) * animation_tree.get_root_motion_position() / delta
    set_velocity(velocity)
    move_and_slide()
[/gdscript]
[/codeblocks]
If [member root_motion_local] is [code]true[/code], return the pre-multiplied translation value with the inverted rotation.
In this case, the code can be written as follows:
[codeblocks]
[gdscript]
func _process(delta):
    if Input.is_action_just_pressed("animate"):
        state_machine.travel("Animate")
    set_quaternion(get_quaternion() * animation_tree.get_root_motion_rotation())
    var velocity = get_quaternion() * animation_tree.get_root_motion_position() / delta
    set_velocity(velocity)
    move_and_slide()
[/gdscript]
[/codeblocks]
*/
//go:nosplit
func (self class) GetRootMotionPosition() Vector3.XYZ { //gd:AnimationMixer.get_root_motion_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
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
func (self class) GetRootMotionRotation() Quaternion.IJKX { //gd:AnimationMixer.get_root_motion_rotation
	var frame = callframe.New()
	var r_ret = callframe.Ret[Quaternion.IJKX](frame)
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
var current_scale = Vector3(1, 1, 1)
var scale_accum = Vector3(1, 1, 1)

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
func (self class) GetRootMotionScale() Vector3.XYZ { //gd:AnimationMixer.get_root_motion_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
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
var prev_root_motion_position_accumulator

func _process(delta):
    if Input.is_action_just_pressed("animate"):
        state_machine.travel("Animate")
    var current_root_motion_position_accumulator = animation_tree.get_root_motion_position_accumulator()
    var difference = current_root_motion_position_accumulator - prev_root_motion_position_accumulator
    prev_root_motion_position_accumulator = current_root_motion_position_accumulator
    transform.origin += difference
[/gdscript]
[/codeblocks]
However, if the animation loops, an unintended discrete change may occur, so this is only useful for some simple use cases.
*/
//go:nosplit
func (self class) GetRootMotionPositionAccumulator() Vector3.XYZ { //gd:AnimationMixer.get_root_motion_position_accumulator
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
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
var prev_root_motion_rotation_accumulator

func _process(delta):
    if Input.is_action_just_pressed("animate"):
        state_machine.travel("Animate")
    var current_root_motion_rotation_accumulator = animation_tree.get_root_motion_rotation_accumulator()
    var difference = prev_root_motion_rotation_accumulator.inverse() * current_root_motion_rotation_accumulator
    prev_root_motion_rotation_accumulator = current_root_motion_rotation_accumulator
    transform.basis *=  Basis(difference)
[/gdscript]
[/codeblocks]
However, if the animation loops, an unintended discrete change may occur, so this is only useful for some simple use cases.
*/
//go:nosplit
func (self class) GetRootMotionRotationAccumulator() Quaternion.IJKX { //gd:AnimationMixer.get_root_motion_rotation_accumulator
	var frame = callframe.New()
	var r_ret = callframe.Ret[Quaternion.IJKX](frame)
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
var prev_root_motion_scale_accumulator

func _process(delta):
    if Input.is_action_just_pressed("animate"):
        state_machine.travel("Animate")
    var current_root_motion_scale_accumulator = animation_tree.get_root_motion_scale_accumulator()
    var difference = current_root_motion_scale_accumulator - prev_root_motion_scale_accumulator
    prev_root_motion_scale_accumulator = current_root_motion_scale_accumulator
    transform.basis = transform.basis.scaled(difference)
[/gdscript]
[/codeblocks]
However, if the animation loops, an unintended discrete change may occur, so this is only useful for some simple use cases.
*/
//go:nosplit
func (self class) GetRootMotionScaleAccumulator() Vector3.XYZ { //gd:AnimationMixer.get_root_motion_scale_accumulator
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_get_root_motion_scale_accumulator, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
[AnimationMixer] caches animated nodes. It may not notice if a node disappears; [method clear_caches] forces it to update the cache again.
*/
//go:nosplit
func (self class) ClearCaches() { //gd:AnimationMixer.clear_caches
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_clear_caches, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Manually advance the animations by the specified time (in seconds).
*/
//go:nosplit
func (self class) Advance(delta float64) { //gd:AnimationMixer.advance
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
func (self class) Capture(name String.Name, duration float64, trans_type gdclass.TweenTransitionType, ease_type gdclass.TweenEaseType) { //gd:AnimationMixer.capture
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, duration)
	callframe.Arg(frame, trans_type)
	callframe.Arg(frame, ease_type)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_capture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetResetOnSaveEnabled(enabled bool) { //gd:AnimationMixer.set_reset_on_save_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_set_reset_on_save_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsResetOnSaveEnabled() bool { //gd:AnimationMixer.is_reset_on_save_enabled
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
func (self class) FindAnimation(animation [1]gdclass.Animation) String.Name { //gd:AnimationMixer.find_animation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(animation[0])[0])
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_find_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the key for the [AnimationLibrary] that contains [param animation] or an empty [StringName] if not found.
*/
//go:nosplit
func (self class) FindAnimationLibrary(animation [1]gdclass.Animation) String.Name { //gd:AnimationMixer.find_animation_library
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(animation[0])[0])
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationMixer.Bind_find_animation_library, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
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

type AnimationCallbackModeProcess = gdclass.AnimationMixerAnimationCallbackModeProcess //gd:AnimationMixer.AnimationCallbackModeProcess

const (
	/*Process animation during physics frames (see [constant Node.NOTIFICATION_INTERNAL_PHYSICS_PROCESS]). This is especially useful when animating physics bodies.*/
	AnimationCallbackModeProcessPhysics AnimationCallbackModeProcess = 0
	/*Process animation during process frames (see [constant Node.NOTIFICATION_INTERNAL_PROCESS]).*/
	AnimationCallbackModeProcessIdle AnimationCallbackModeProcess = 1
	/*Do not process animation. Use [method advance] to process the animation manually.*/
	AnimationCallbackModeProcessManual AnimationCallbackModeProcess = 2
)

type AnimationCallbackModeMethod = gdclass.AnimationMixerAnimationCallbackModeMethod //gd:AnimationMixer.AnimationCallbackModeMethod

const (
	/*Batch method calls during the animation process, then do the calls after events are processed. This avoids bugs involving deleting nodes or modifying the AnimationPlayer while playing.*/
	AnimationCallbackModeMethodDeferred AnimationCallbackModeMethod = 0
	/*Make method calls immediately when reached in the animation.*/
	AnimationCallbackModeMethodImmediate AnimationCallbackModeMethod = 1
)

type AnimationCallbackModeDiscrete = gdclass.AnimationMixerAnimationCallbackModeDiscrete //gd:AnimationMixer.AnimationCallbackModeDiscrete

const (
	/*An [constant Animation.UPDATE_DISCRETE] track value takes precedence when blending [constant Animation.UPDATE_CONTINUOUS] or [constant Animation.UPDATE_CAPTURE] track values and [constant Animation.UPDATE_DISCRETE] track values.*/
	AnimationCallbackModeDiscreteDominant AnimationCallbackModeDiscrete = 0
	/*An [constant Animation.UPDATE_CONTINUOUS] or [constant Animation.UPDATE_CAPTURE] track value takes precedence when blending the [constant Animation.UPDATE_CONTINUOUS] or [constant Animation.UPDATE_CAPTURE] track values and the [constant Animation.UPDATE_DISCRETE] track values. This is the default behavior for [AnimationPlayer].*/
	AnimationCallbackModeDiscreteRecessive AnimationCallbackModeDiscrete = 1
	/*Always treat the [constant Animation.UPDATE_DISCRETE] track value as [constant Animation.UPDATE_CONTINUOUS] with [constant Animation.INTERPOLATION_NEAREST]. This is the default behavior for [AnimationTree].
	  If a value track has un-interpolatable type key values, it is internally converted to use [constant ANIMATION_CALLBACK_MODE_DISCRETE_RECESSIVE] with [constant Animation.UPDATE_DISCRETE].
	  Un-interpolatable type list:
	  - [constant @GlobalScope.TYPE_NIL]
	  - [constant @GlobalScope.TYPE_NODE_PATH]
	  - [constant @GlobalScope.TYPE_RID]
	  - [constant @GlobalScope.TYPE_OBJECT]
	  - [constant @GlobalScope.TYPE_CALLABLE]
	  - [constant @GlobalScope.TYPE_SIGNAL]
	  - [constant @GlobalScope.TYPE_DICTIONARY]
	  - [constant @GlobalScope.TYPE_PACKED_BYTE_ARRAY]
	  [constant @GlobalScope.TYPE_BOOL] and [constant @GlobalScope.TYPE_INT] are treated as [constant @GlobalScope.TYPE_FLOAT] during blending and rounded when the result is retrieved.
	  It is same for arrays and vectors with them such as [constant @GlobalScope.TYPE_PACKED_INT32_ARRAY] or [constant @GlobalScope.TYPE_VECTOR2I], they are treated as [constant @GlobalScope.TYPE_PACKED_FLOAT32_ARRAY] or [constant @GlobalScope.TYPE_VECTOR2]. Also note that for arrays, the size is also interpolated.
	  [constant @GlobalScope.TYPE_STRING] and [constant @GlobalScope.TYPE_STRING_NAME] are interpolated between character codes and lengths, but note that there is a difference in algorithm between interpolation between keys and interpolation by blending.*/
	AnimationCallbackModeDiscreteForceContinuous AnimationCallbackModeDiscrete = 2
)
