// Package Animation provides methods for working with Animation object instances.
package Animation

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
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
import "graphics.gd/variant/Vector2"
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
This resource holds data that can be used to animate anything in the engine. Animations are divided into tracks and each track must be linked to a node. The state of that node can be changed through time, by adding timed keys (events) to the track.
[codeblocks]
[gdscript]
# This creates an animation that makes the node "Enemy" move to the right by
# 100 pixels in 2.0 seconds.
var animation = Animation.new()
var track_index = animation.add_track(Animation.TYPE_VALUE)
animation.track_set_path(track_index, "Enemy:position:x")
animation.track_insert_key(track_index, 0.0, 0)
animation.track_insert_key(track_index, 2.0, 100)
animation.length = 2.0
[/gdscript]
[csharp]
// This creates an animation that makes the node "Enemy" move to the right by
// 100 pixels in 2.0 seconds.
var animation = new Animation();
int trackIndex = animation.AddTrack(Animation.TrackType.Value);
animation.TrackSetPath(trackIndex, "Enemy:position:x");
animation.TrackInsertKey(trackIndex, 0.0f, 0);
animation.TrackInsertKey(trackIndex, 2.0f, 100);
animation.Length = 2.0f;
[/csharp]
[/codeblocks]
Animations are just data containers, and must be added to nodes such as an [AnimationPlayer] to be played back. Animation tracks have different types, each with its own set of dedicated methods. Check [enum TrackType] to see available types.
[b]Note:[/b] For 3D position/rotation/scale, using the dedicated [constant TYPE_POSITION_3D], [constant TYPE_ROTATION_3D] and [constant TYPE_SCALE_3D] track types instead of [constant TYPE_VALUE] is recommended for performance reasons.
*/
type Instance [1]gdclass.Animation

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAnimation() Instance
}

/*
Adds a track to the Animation.
*/
func (self Instance) AddTrack(atype gdclass.AnimationTrackType) int { //gd:Animation.add_track
	return int(int(class(self).AddTrack(atype, int64(-1))))
}

/*
Removes a track by specifying the track index.
*/
func (self Instance) RemoveTrack(track_idx int) { //gd:Animation.remove_track
	class(self).RemoveTrack(int64(track_idx))
}

/*
Returns the amount of tracks in the animation.
*/
func (self Instance) GetTrackCount() int { //gd:Animation.get_track_count
	return int(int(class(self).GetTrackCount()))
}

/*
Gets the type of a track.
*/
func (self Instance) TrackGetType(track_idx int) gdclass.AnimationTrackType { //gd:Animation.track_get_type
	return gdclass.AnimationTrackType(class(self).TrackGetType(int64(track_idx)))
}

/*
Gets the path of a track. For more information on the path format, see [method track_set_path].
*/
func (self Instance) TrackGetPath(track_idx int) string { //gd:Animation.track_get_path
	return string(class(self).TrackGetPath(int64(track_idx)).String())
}

/*
Sets the path of a track. Paths must be valid scene-tree paths to a node and must be specified starting from the [member AnimationMixer.root_node] that will reproduce the animation. Tracks that control properties or bones must append their name after the path, separated by [code]":"[/code].
For example, [code]"character/skeleton:ankle"[/code] or [code]"character/mesh:transform/local"[/code].
*/
func (self Instance) TrackSetPath(track_idx int, path string) { //gd:Animation.track_set_path
	class(self).TrackSetPath(int64(track_idx), Path.ToNode(String.New(path)))
}

/*
Returns the index of the specified track. If the track is not found, return -1.
*/
func (self Instance) FindTrack(path string, atype gdclass.AnimationTrackType) int { //gd:Animation.find_track
	return int(int(class(self).FindTrack(Path.ToNode(String.New(path)), atype)))
}

/*
Moves a track up.
*/
func (self Instance) TrackMoveUp(track_idx int) { //gd:Animation.track_move_up
	class(self).TrackMoveUp(int64(track_idx))
}

/*
Moves a track down.
*/
func (self Instance) TrackMoveDown(track_idx int) { //gd:Animation.track_move_down
	class(self).TrackMoveDown(int64(track_idx))
}

/*
Changes the index position of track [param track_idx] to the one defined in [param to_idx].
*/
func (self Instance) TrackMoveTo(track_idx int, to_idx int) { //gd:Animation.track_move_to
	class(self).TrackMoveTo(int64(track_idx), int64(to_idx))
}

/*
Swaps the track [param track_idx]'s index position with the track [param with_idx].
*/
func (self Instance) TrackSwap(track_idx int, with_idx int) { //gd:Animation.track_swap
	class(self).TrackSwap(int64(track_idx), int64(with_idx))
}

/*
Sets the given track as imported or not.
*/
func (self Instance) TrackSetImported(track_idx int, imported bool) { //gd:Animation.track_set_imported
	class(self).TrackSetImported(int64(track_idx), imported)
}

/*
Returns [code]true[/code] if the given track is imported. Else, return [code]false[/code].
*/
func (self Instance) TrackIsImported(track_idx int) bool { //gd:Animation.track_is_imported
	return bool(class(self).TrackIsImported(int64(track_idx)))
}

/*
Enables/disables the given track. Tracks are enabled by default.
*/
func (self Instance) TrackSetEnabled(track_idx int, enabled bool) { //gd:Animation.track_set_enabled
	class(self).TrackSetEnabled(int64(track_idx), enabled)
}

/*
Returns [code]true[/code] if the track at index [param track_idx] is enabled.
*/
func (self Instance) TrackIsEnabled(track_idx int) bool { //gd:Animation.track_is_enabled
	return bool(class(self).TrackIsEnabled(int64(track_idx)))
}

/*
Inserts a key in a given 3D position track. Returns the key index.
*/
func (self Instance) PositionTrackInsertKey(track_idx int, time Float.X, position Vector3.XYZ) int { //gd:Animation.position_track_insert_key
	return int(int(class(self).PositionTrackInsertKey(int64(track_idx), float64(time), Vector3.XYZ(position))))
}

/*
Inserts a key in a given 3D rotation track. Returns the key index.
*/
func (self Instance) RotationTrackInsertKey(track_idx int, time Float.X, rotation Quaternion.IJKX) int { //gd:Animation.rotation_track_insert_key
	return int(int(class(self).RotationTrackInsertKey(int64(track_idx), float64(time), rotation)))
}

/*
Inserts a key in a given 3D scale track. Returns the key index.
*/
func (self Instance) ScaleTrackInsertKey(track_idx int, time Float.X, scale Vector3.XYZ) int { //gd:Animation.scale_track_insert_key
	return int(int(class(self).ScaleTrackInsertKey(int64(track_idx), float64(time), Vector3.XYZ(scale))))
}

/*
Inserts a key in a given blend shape track. Returns the key index.
*/
func (self Instance) BlendShapeTrackInsertKey(track_idx int, time Float.X, amount Float.X) int { //gd:Animation.blend_shape_track_insert_key
	return int(int(class(self).BlendShapeTrackInsertKey(int64(track_idx), float64(time), float64(amount))))
}

/*
Returns the interpolated position value at the given time (in seconds). The [param track_idx] must be the index of a 3D position track.
*/
func (self Instance) PositionTrackInterpolate(track_idx int, time_sec Float.X) Vector3.XYZ { //gd:Animation.position_track_interpolate
	return Vector3.XYZ(class(self).PositionTrackInterpolate(int64(track_idx), float64(time_sec), false))
}

/*
Returns the interpolated rotation value at the given time (in seconds). The [param track_idx] must be the index of a 3D rotation track.
*/
func (self Instance) RotationTrackInterpolate(track_idx int, time_sec Float.X) Quaternion.IJKX { //gd:Animation.rotation_track_interpolate
	return Quaternion.IJKX(class(self).RotationTrackInterpolate(int64(track_idx), float64(time_sec), false))
}

/*
Returns the interpolated scale value at the given time (in seconds). The [param track_idx] must be the index of a 3D scale track.
*/
func (self Instance) ScaleTrackInterpolate(track_idx int, time_sec Float.X) Vector3.XYZ { //gd:Animation.scale_track_interpolate
	return Vector3.XYZ(class(self).ScaleTrackInterpolate(int64(track_idx), float64(time_sec), false))
}

/*
Returns the interpolated blend shape value at the given time (in seconds). The [param track_idx] must be the index of a blend shape track.
*/
func (self Instance) BlendShapeTrackInterpolate(track_idx int, time_sec Float.X) Float.X { //gd:Animation.blend_shape_track_interpolate
	return Float.X(Float.X(class(self).BlendShapeTrackInterpolate(int64(track_idx), float64(time_sec), false)))
}

/*
Inserts a generic key in a given track. Returns the key index.
*/
func (self Instance) TrackInsertKey(track_idx int, time Float.X, key any) int { //gd:Animation.track_insert_key
	return int(int(class(self).TrackInsertKey(int64(track_idx), float64(time), variant.New(key), float64(1))))
}

/*
Removes a key by index in a given track.
*/
func (self Instance) TrackRemoveKey(track_idx int, key_idx int) { //gd:Animation.track_remove_key
	class(self).TrackRemoveKey(int64(track_idx), int64(key_idx))
}

/*
Removes a key at [param time] in a given track.
*/
func (self Instance) TrackRemoveKeyAtTime(track_idx int, time Float.X) { //gd:Animation.track_remove_key_at_time
	class(self).TrackRemoveKeyAtTime(int64(track_idx), float64(time))
}

/*
Sets the value of an existing key.
*/
func (self Instance) TrackSetKeyValue(track_idx int, key int, value any) { //gd:Animation.track_set_key_value
	class(self).TrackSetKeyValue(int64(track_idx), int64(key), variant.New(value))
}

/*
Sets the transition curve (easing) for a specific key (see the built-in math function [method @GlobalScope.ease]).
*/
func (self Instance) TrackSetKeyTransition(track_idx int, key_idx int, transition Float.X) { //gd:Animation.track_set_key_transition
	class(self).TrackSetKeyTransition(int64(track_idx), int64(key_idx), float64(transition))
}

/*
Sets the time of an existing key.
*/
func (self Instance) TrackSetKeyTime(track_idx int, key_idx int, time Float.X) { //gd:Animation.track_set_key_time
	class(self).TrackSetKeyTime(int64(track_idx), int64(key_idx), float64(time))
}

/*
Returns the transition curve (easing) for a specific key (see the built-in math function [method @GlobalScope.ease]).
*/
func (self Instance) TrackGetKeyTransition(track_idx int, key_idx int) Float.X { //gd:Animation.track_get_key_transition
	return Float.X(Float.X(class(self).TrackGetKeyTransition(int64(track_idx), int64(key_idx))))
}

/*
Returns the number of keys in a given track.
*/
func (self Instance) TrackGetKeyCount(track_idx int) int { //gd:Animation.track_get_key_count
	return int(int(class(self).TrackGetKeyCount(int64(track_idx))))
}

/*
Returns the value of a given key in a given track.
*/
func (self Instance) TrackGetKeyValue(track_idx int, key_idx int) any { //gd:Animation.track_get_key_value
	return any(class(self).TrackGetKeyValue(int64(track_idx), int64(key_idx)).Interface())
}

/*
Returns the time at which the key is located.
*/
func (self Instance) TrackGetKeyTime(track_idx int, key_idx int) Float.X { //gd:Animation.track_get_key_time
	return Float.X(Float.X(class(self).TrackGetKeyTime(int64(track_idx), int64(key_idx))))
}

/*
Finds the key index by time in a given track. Optionally, only find it if the approx/exact time is given.
If [param limit] is [code]true[/code], it does not return keys outside the animation range.
If [param backward] is [code]true[/code], the direction is reversed in methods that rely on one directional processing.
For example, in case [param find_mode] is [constant FIND_MODE_NEAREST], if there is no key in the current position just after seeked, the first key found is retrieved by searching before the position, but if [param backward] is [code]true[/code], the first key found is retrieved after the position.
*/
func (self Instance) TrackFindKey(track_idx int, time Float.X) int { //gd:Animation.track_find_key
	return int(int(class(self).TrackFindKey(int64(track_idx), float64(time), 0, false, false)))
}

/*
Sets the interpolation type of a given track.
*/
func (self Instance) TrackSetInterpolationType(track_idx int, interpolation gdclass.AnimationInterpolationType) { //gd:Animation.track_set_interpolation_type
	class(self).TrackSetInterpolationType(int64(track_idx), interpolation)
}

/*
Returns the interpolation type of a given track.
*/
func (self Instance) TrackGetInterpolationType(track_idx int) gdclass.AnimationInterpolationType { //gd:Animation.track_get_interpolation_type
	return gdclass.AnimationInterpolationType(class(self).TrackGetInterpolationType(int64(track_idx)))
}

/*
If [code]true[/code], the track at [param track_idx] wraps the interpolation loop.
*/
func (self Instance) TrackSetInterpolationLoopWrap(track_idx int, interpolation bool) { //gd:Animation.track_set_interpolation_loop_wrap
	class(self).TrackSetInterpolationLoopWrap(int64(track_idx), interpolation)
}

/*
Returns [code]true[/code] if the track at [param track_idx] wraps the interpolation loop. New tracks wrap the interpolation loop by default.
*/
func (self Instance) TrackGetInterpolationLoopWrap(track_idx int) bool { //gd:Animation.track_get_interpolation_loop_wrap
	return bool(class(self).TrackGetInterpolationLoopWrap(int64(track_idx)))
}

/*
Returns [code]true[/code] if the track is compressed, [code]false[/code] otherwise. See also [method compress].
*/
func (self Instance) TrackIsCompressed(track_idx int) bool { //gd:Animation.track_is_compressed
	return bool(class(self).TrackIsCompressed(int64(track_idx)))
}

/*
Sets the update mode (see [enum UpdateMode]) of a value track.
*/
func (self Instance) ValueTrackSetUpdateMode(track_idx int, mode gdclass.AnimationUpdateMode) { //gd:Animation.value_track_set_update_mode
	class(self).ValueTrackSetUpdateMode(int64(track_idx), mode)
}

/*
Returns the update mode of a value track.
*/
func (self Instance) ValueTrackGetUpdateMode(track_idx int) gdclass.AnimationUpdateMode { //gd:Animation.value_track_get_update_mode
	return gdclass.AnimationUpdateMode(class(self).ValueTrackGetUpdateMode(int64(track_idx)))
}

/*
Returns the interpolated value at the given time (in seconds). The [param track_idx] must be the index of a value track.
A [param backward] mainly affects the direction of key retrieval of the track with [constant UPDATE_DISCRETE] converted by [constant AnimationMixer.ANIMATION_CALLBACK_MODE_DISCRETE_FORCE_CONTINUOUS] to match the result with [method track_find_key].
*/
func (self Instance) ValueTrackInterpolate(track_idx int, time_sec Float.X) any { //gd:Animation.value_track_interpolate
	return any(class(self).ValueTrackInterpolate(int64(track_idx), float64(time_sec), false).Interface())
}

/*
Returns the method name of a method track.
*/
func (self Instance) MethodTrackGetName(track_idx int, key_idx int) string { //gd:Animation.method_track_get_name
	return string(class(self).MethodTrackGetName(int64(track_idx), int64(key_idx)).String())
}

/*
Returns the arguments values to be called on a method track for a given key in a given track.
*/
func (self Instance) MethodTrackGetParams(track_idx int, key_idx int) []any { //gd:Animation.method_track_get_params
	return []any(gd.ArrayAs[[]any](gd.InternalArray(class(self).MethodTrackGetParams(int64(track_idx), int64(key_idx)))))
}

/*
Inserts a Bezier Track key at the given [param time] in seconds. The [param track_idx] must be the index of a Bezier Track.
[param in_handle] is the left-side weight of the added Bezier curve point, [param out_handle] is the right-side one, while [param value] is the actual value at this point.
*/
func (self Instance) BezierTrackInsertKey(track_idx int, time Float.X, value Float.X) int { //gd:Animation.bezier_track_insert_key
	return int(int(class(self).BezierTrackInsertKey(int64(track_idx), float64(time), float64(value), Vector2.XY(gd.Vector2{0, 0}), Vector2.XY(gd.Vector2{0, 0}))))
}

/*
Sets the value of the key identified by [param key_idx] to the given value. The [param track_idx] must be the index of a Bezier Track.
*/
func (self Instance) BezierTrackSetKeyValue(track_idx int, key_idx int, value Float.X) { //gd:Animation.bezier_track_set_key_value
	class(self).BezierTrackSetKeyValue(int64(track_idx), int64(key_idx), float64(value))
}

/*
Sets the in handle of the key identified by [param key_idx] to value [param in_handle]. The [param track_idx] must be the index of a Bezier Track.
*/
func (self Instance) BezierTrackSetKeyInHandle(track_idx int, key_idx int, in_handle Vector2.XY) { //gd:Animation.bezier_track_set_key_in_handle
	class(self).BezierTrackSetKeyInHandle(int64(track_idx), int64(key_idx), Vector2.XY(in_handle), float64(1.0))
}

/*
Sets the out handle of the key identified by [param key_idx] to value [param out_handle]. The [param track_idx] must be the index of a Bezier Track.
*/
func (self Instance) BezierTrackSetKeyOutHandle(track_idx int, key_idx int, out_handle Vector2.XY) { //gd:Animation.bezier_track_set_key_out_handle
	class(self).BezierTrackSetKeyOutHandle(int64(track_idx), int64(key_idx), Vector2.XY(out_handle), float64(1.0))
}

/*
Returns the value of the key identified by [param key_idx]. The [param track_idx] must be the index of a Bezier Track.
*/
func (self Instance) BezierTrackGetKeyValue(track_idx int, key_idx int) Float.X { //gd:Animation.bezier_track_get_key_value
	return Float.X(Float.X(class(self).BezierTrackGetKeyValue(int64(track_idx), int64(key_idx))))
}

/*
Returns the in handle of the key identified by [param key_idx]. The [param track_idx] must be the index of a Bezier Track.
*/
func (self Instance) BezierTrackGetKeyInHandle(track_idx int, key_idx int) Vector2.XY { //gd:Animation.bezier_track_get_key_in_handle
	return Vector2.XY(class(self).BezierTrackGetKeyInHandle(int64(track_idx), int64(key_idx)))
}

/*
Returns the out handle of the key identified by [param key_idx]. The [param track_idx] must be the index of a Bezier Track.
*/
func (self Instance) BezierTrackGetKeyOutHandle(track_idx int, key_idx int) Vector2.XY { //gd:Animation.bezier_track_get_key_out_handle
	return Vector2.XY(class(self).BezierTrackGetKeyOutHandle(int64(track_idx), int64(key_idx)))
}

/*
Returns the interpolated value at the given [param time] (in seconds). The [param track_idx] must be the index of a Bezier Track.
*/
func (self Instance) BezierTrackInterpolate(track_idx int, time Float.X) Float.X { //gd:Animation.bezier_track_interpolate
	return Float.X(Float.X(class(self).BezierTrackInterpolate(int64(track_idx), float64(time))))
}

/*
Inserts an Audio Track key at the given [param time] in seconds. The [param track_idx] must be the index of an Audio Track.
[param stream] is the [AudioStream] resource to play. [param start_offset] is the number of seconds cut off at the beginning of the audio stream, while [param end_offset] is at the ending.
*/
func (self Instance) AudioTrackInsertKey(track_idx int, time Float.X, stream [1]gdclass.Resource) int { //gd:Animation.audio_track_insert_key
	return int(int(class(self).AudioTrackInsertKey(int64(track_idx), float64(time), stream, float64(0), float64(0))))
}

/*
Sets the stream of the key identified by [param key_idx] to value [param stream]. The [param track_idx] must be the index of an Audio Track.
*/
func (self Instance) AudioTrackSetKeyStream(track_idx int, key_idx int, stream [1]gdclass.Resource) { //gd:Animation.audio_track_set_key_stream
	class(self).AudioTrackSetKeyStream(int64(track_idx), int64(key_idx), stream)
}

/*
Sets the start offset of the key identified by [param key_idx] to value [param offset]. The [param track_idx] must be the index of an Audio Track.
*/
func (self Instance) AudioTrackSetKeyStartOffset(track_idx int, key_idx int, offset Float.X) { //gd:Animation.audio_track_set_key_start_offset
	class(self).AudioTrackSetKeyStartOffset(int64(track_idx), int64(key_idx), float64(offset))
}

/*
Sets the end offset of the key identified by [param key_idx] to value [param offset]. The [param track_idx] must be the index of an Audio Track.
*/
func (self Instance) AudioTrackSetKeyEndOffset(track_idx int, key_idx int, offset Float.X) { //gd:Animation.audio_track_set_key_end_offset
	class(self).AudioTrackSetKeyEndOffset(int64(track_idx), int64(key_idx), float64(offset))
}

/*
Returns the audio stream of the key identified by [param key_idx]. The [param track_idx] must be the index of an Audio Track.
*/
func (self Instance) AudioTrackGetKeyStream(track_idx int, key_idx int) [1]gdclass.Resource { //gd:Animation.audio_track_get_key_stream
	return [1]gdclass.Resource(class(self).AudioTrackGetKeyStream(int64(track_idx), int64(key_idx)))
}

/*
Returns the start offset of the key identified by [param key_idx]. The [param track_idx] must be the index of an Audio Track.
Start offset is the number of seconds cut off at the beginning of the audio stream.
*/
func (self Instance) AudioTrackGetKeyStartOffset(track_idx int, key_idx int) Float.X { //gd:Animation.audio_track_get_key_start_offset
	return Float.X(Float.X(class(self).AudioTrackGetKeyStartOffset(int64(track_idx), int64(key_idx))))
}

/*
Returns the end offset of the key identified by [param key_idx]. The [param track_idx] must be the index of an Audio Track.
End offset is the number of seconds cut off at the ending of the audio stream.
*/
func (self Instance) AudioTrackGetKeyEndOffset(track_idx int, key_idx int) Float.X { //gd:Animation.audio_track_get_key_end_offset
	return Float.X(Float.X(class(self).AudioTrackGetKeyEndOffset(int64(track_idx), int64(key_idx))))
}

/*
Sets whether the track will be blended with other animations. If [code]true[/code], the audio playback volume changes depending on the blend value.
*/
func (self Instance) AudioTrackSetUseBlend(track_idx int, enable bool) { //gd:Animation.audio_track_set_use_blend
	class(self).AudioTrackSetUseBlend(int64(track_idx), enable)
}

/*
Returns [code]true[/code] if the track at [param track_idx] will be blended with other animations.
*/
func (self Instance) AudioTrackIsUseBlend(track_idx int) bool { //gd:Animation.audio_track_is_use_blend
	return bool(class(self).AudioTrackIsUseBlend(int64(track_idx)))
}

/*
Inserts a key with value [param animation] at the given [param time] (in seconds). The [param track_idx] must be the index of an Animation Track.
*/
func (self Instance) AnimationTrackInsertKey(track_idx int, time Float.X, animation string) int { //gd:Animation.animation_track_insert_key
	return int(int(class(self).AnimationTrackInsertKey(int64(track_idx), float64(time), String.Name(String.New(animation)))))
}

/*
Sets the key identified by [param key_idx] to value [param animation]. The [param track_idx] must be the index of an Animation Track.
*/
func (self Instance) AnimationTrackSetKeyAnimation(track_idx int, key_idx int, animation string) { //gd:Animation.animation_track_set_key_animation
	class(self).AnimationTrackSetKeyAnimation(int64(track_idx), int64(key_idx), String.Name(String.New(animation)))
}

/*
Returns the animation name at the key identified by [param key_idx]. The [param track_idx] must be the index of an Animation Track.
*/
func (self Instance) AnimationTrackGetKeyAnimation(track_idx int, key_idx int) string { //gd:Animation.animation_track_get_key_animation
	return string(class(self).AnimationTrackGetKeyAnimation(int64(track_idx), int64(key_idx)).String())
}

/*
Clear the animation (clear all tracks and reset all).
*/
func (self Instance) Clear() { //gd:Animation.clear
	class(self).Clear()
}

/*
Adds a new track to [param to_animation] that is a copy of the given track from this animation.
*/
func (self Instance) CopyTrack(track_idx int, to_animation [1]gdclass.Animation) { //gd:Animation.copy_track
	class(self).CopyTrack(int64(track_idx), to_animation)
}

/*
Compress the animation and all its tracks in-place. This will make [method track_is_compressed] return [code]true[/code] once called on this [Animation]. Compressed tracks require less memory to be played, and are designed to be used for complex 3D animations (such as cutscenes) imported from external 3D software. Compression is lossy, but the difference is usually not noticeable in real world conditions.
[b]Note:[/b] Compressed tracks have various limitations (such as not being editable from the editor), so only use compressed animations if you actually need them.
*/
func (self Instance) Compress() { //gd:Animation.compress
	class(self).Compress(int64(8192), int64(120), float64(4.0))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Animation

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Animation"))
	casted := Instance{*(*gdclass.Animation)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Length() Float.X {
	return Float.X(Float.X(class(self).GetLength()))
}

func (self Instance) SetLength(value Float.X) {
	class(self).SetLength(float64(value))
}

func (self Instance) LoopMode() gdclass.AnimationLoopMode {
	return gdclass.AnimationLoopMode(class(self).GetLoopMode())
}

func (self Instance) SetLoopMode(value gdclass.AnimationLoopMode) {
	class(self).SetLoopMode(value)
}

func (self Instance) Step() Float.X {
	return Float.X(Float.X(class(self).GetStep()))
}

func (self Instance) SetStep(value Float.X) {
	class(self).SetStep(float64(value))
}

func (self Instance) CaptureIncluded() bool {
	return bool(class(self).IsCaptureIncluded())
}

/*
Adds a track to the Animation.
*/
//go:nosplit
func (self class) AddTrack(atype gdclass.AnimationTrackType, at_position int64) int64 { //gd:Animation.add_track
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, at_position)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_add_track, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes a track by specifying the track index.
*/
//go:nosplit
func (self class) RemoveTrack(track_idx int64) { //gd:Animation.remove_track
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_remove_track, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the amount of tracks in the animation.
*/
//go:nosplit
func (self class) GetTrackCount() int64 { //gd:Animation.get_track_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_get_track_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets the type of a track.
*/
//go:nosplit
func (self class) TrackGetType(track_idx int64) gdclass.AnimationTrackType { //gd:Animation.track_get_type
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Ret[gdclass.AnimationTrackType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_get_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets the path of a track. For more information on the path format, see [method track_set_path].
*/
//go:nosplit
func (self class) TrackGetPath(track_idx int64) Path.ToNode { //gd:Animation.track_get_path
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_get_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Sets the path of a track. Paths must be valid scene-tree paths to a node and must be specified starting from the [member AnimationMixer.root_node] that will reproduce the animation. Tracks that control properties or bones must append their name after the path, separated by [code]":"[/code].
For example, [code]"character/skeleton:ankle"[/code] or [code]"character/mesh:transform/local"[/code].
*/
//go:nosplit
func (self class) TrackSetPath(track_idx int64, path Path.ToNode) { //gd:Animation.track_set_path
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_set_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the index of the specified track. If the track is not found, return -1.
*/
//go:nosplit
func (self class) FindTrack(path Path.ToNode, atype gdclass.AnimationTrackType) int64 { //gd:Animation.find_track
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_find_track, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Moves a track up.
*/
//go:nosplit
func (self class) TrackMoveUp(track_idx int64) { //gd:Animation.track_move_up
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_move_up, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Moves a track down.
*/
//go:nosplit
func (self class) TrackMoveDown(track_idx int64) { //gd:Animation.track_move_down
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_move_down, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Changes the index position of track [param track_idx] to the one defined in [param to_idx].
*/
//go:nosplit
func (self class) TrackMoveTo(track_idx int64, to_idx int64) { //gd:Animation.track_move_to
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, to_idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_move_to, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Swaps the track [param track_idx]'s index position with the track [param with_idx].
*/
//go:nosplit
func (self class) TrackSwap(track_idx int64, with_idx int64) { //gd:Animation.track_swap
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, with_idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_swap, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the given track as imported or not.
*/
//go:nosplit
func (self class) TrackSetImported(track_idx int64, imported bool) { //gd:Animation.track_set_imported
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, imported)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_set_imported, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the given track is imported. Else, return [code]false[/code].
*/
//go:nosplit
func (self class) TrackIsImported(track_idx int64) bool { //gd:Animation.track_is_imported
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_is_imported, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Enables/disables the given track. Tracks are enabled by default.
*/
//go:nosplit
func (self class) TrackSetEnabled(track_idx int64, enabled bool) { //gd:Animation.track_set_enabled
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_set_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the track at index [param track_idx] is enabled.
*/
//go:nosplit
func (self class) TrackIsEnabled(track_idx int64) bool { //gd:Animation.track_is_enabled
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_is_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Inserts a key in a given 3D position track. Returns the key index.
*/
//go:nosplit
func (self class) PositionTrackInsertKey(track_idx int64, time float64, position Vector3.XYZ) int64 { //gd:Animation.position_track_insert_key
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_position_track_insert_key, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Inserts a key in a given 3D rotation track. Returns the key index.
*/
//go:nosplit
func (self class) RotationTrackInsertKey(track_idx int64, time float64, rotation Quaternion.IJKX) int64 { //gd:Animation.rotation_track_insert_key
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	callframe.Arg(frame, rotation)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_rotation_track_insert_key, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Inserts a key in a given 3D scale track. Returns the key index.
*/
//go:nosplit
func (self class) ScaleTrackInsertKey(track_idx int64, time float64, scale Vector3.XYZ) int64 { //gd:Animation.scale_track_insert_key
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	callframe.Arg(frame, scale)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_scale_track_insert_key, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Inserts a key in a given blend shape track. Returns the key index.
*/
//go:nosplit
func (self class) BlendShapeTrackInsertKey(track_idx int64, time float64, amount float64) int64 { //gd:Animation.blend_shape_track_insert_key
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	callframe.Arg(frame, amount)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_blend_shape_track_insert_key, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the interpolated position value at the given time (in seconds). The [param track_idx] must be the index of a 3D position track.
*/
//go:nosplit
func (self class) PositionTrackInterpolate(track_idx int64, time_sec float64, backward bool) Vector3.XYZ { //gd:Animation.position_track_interpolate
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time_sec)
	callframe.Arg(frame, backward)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_position_track_interpolate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the interpolated rotation value at the given time (in seconds). The [param track_idx] must be the index of a 3D rotation track.
*/
//go:nosplit
func (self class) RotationTrackInterpolate(track_idx int64, time_sec float64, backward bool) Quaternion.IJKX { //gd:Animation.rotation_track_interpolate
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time_sec)
	callframe.Arg(frame, backward)
	var r_ret = callframe.Ret[Quaternion.IJKX](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_rotation_track_interpolate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the interpolated scale value at the given time (in seconds). The [param track_idx] must be the index of a 3D scale track.
*/
//go:nosplit
func (self class) ScaleTrackInterpolate(track_idx int64, time_sec float64, backward bool) Vector3.XYZ { //gd:Animation.scale_track_interpolate
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time_sec)
	callframe.Arg(frame, backward)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_scale_track_interpolate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the interpolated blend shape value at the given time (in seconds). The [param track_idx] must be the index of a blend shape track.
*/
//go:nosplit
func (self class) BlendShapeTrackInterpolate(track_idx int64, time_sec float64, backward bool) float64 { //gd:Animation.blend_shape_track_interpolate
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time_sec)
	callframe.Arg(frame, backward)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_blend_shape_track_interpolate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Inserts a generic key in a given track. Returns the key index.
*/
//go:nosplit
func (self class) TrackInsertKey(track_idx int64, time float64, key variant.Any, transition float64) int64 { //gd:Animation.track_insert_key
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(key)))
	callframe.Arg(frame, transition)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_insert_key, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes a key by index in a given track.
*/
//go:nosplit
func (self class) TrackRemoveKey(track_idx int64, key_idx int64) { //gd:Animation.track_remove_key
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_remove_key, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a key at [param time] in a given track.
*/
//go:nosplit
func (self class) TrackRemoveKeyAtTime(track_idx int64, time float64) { //gd:Animation.track_remove_key_at_time
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_remove_key_at_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the value of an existing key.
*/
//go:nosplit
func (self class) TrackSetKeyValue(track_idx int64, key int64, value variant.Any) { //gd:Animation.track_set_key_value
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key)
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(value)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_set_key_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the transition curve (easing) for a specific key (see the built-in math function [method @GlobalScope.ease]).
*/
//go:nosplit
func (self class) TrackSetKeyTransition(track_idx int64, key_idx int64, transition float64) { //gd:Animation.track_set_key_transition
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	callframe.Arg(frame, transition)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_set_key_transition, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the time of an existing key.
*/
//go:nosplit
func (self class) TrackSetKeyTime(track_idx int64, key_idx int64, time float64) { //gd:Animation.track_set_key_time
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	callframe.Arg(frame, time)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_set_key_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the transition curve (easing) for a specific key (see the built-in math function [method @GlobalScope.ease]).
*/
//go:nosplit
func (self class) TrackGetKeyTransition(track_idx int64, key_idx int64) float64 { //gd:Animation.track_get_key_transition
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_get_key_transition, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of keys in a given track.
*/
//go:nosplit
func (self class) TrackGetKeyCount(track_idx int64) int64 { //gd:Animation.track_get_key_count
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_get_key_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the value of a given key in a given track.
*/
//go:nosplit
func (self class) TrackGetKeyValue(track_idx int64, key_idx int64) variant.Any { //gd:Animation.track_get_key_value
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_get_key_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the time at which the key is located.
*/
//go:nosplit
func (self class) TrackGetKeyTime(track_idx int64, key_idx int64) float64 { //gd:Animation.track_get_key_time
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_get_key_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Finds the key index by time in a given track. Optionally, only find it if the approx/exact time is given.
If [param limit] is [code]true[/code], it does not return keys outside the animation range.
If [param backward] is [code]true[/code], the direction is reversed in methods that rely on one directional processing.
For example, in case [param find_mode] is [constant FIND_MODE_NEAREST], if there is no key in the current position just after seeked, the first key found is retrieved by searching before the position, but if [param backward] is [code]true[/code], the first key found is retrieved after the position.
*/
//go:nosplit
func (self class) TrackFindKey(track_idx int64, time float64, find_mode gdclass.AnimationFindMode, limit bool, backward bool) int64 { //gd:Animation.track_find_key
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	callframe.Arg(frame, find_mode)
	callframe.Arg(frame, limit)
	callframe.Arg(frame, backward)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_find_key, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the interpolation type of a given track.
*/
//go:nosplit
func (self class) TrackSetInterpolationType(track_idx int64, interpolation gdclass.AnimationInterpolationType) { //gd:Animation.track_set_interpolation_type
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, interpolation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_set_interpolation_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the interpolation type of a given track.
*/
//go:nosplit
func (self class) TrackGetInterpolationType(track_idx int64) gdclass.AnimationInterpolationType { //gd:Animation.track_get_interpolation_type
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Ret[gdclass.AnimationInterpolationType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_get_interpolation_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [code]true[/code], the track at [param track_idx] wraps the interpolation loop.
*/
//go:nosplit
func (self class) TrackSetInterpolationLoopWrap(track_idx int64, interpolation bool) { //gd:Animation.track_set_interpolation_loop_wrap
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, interpolation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_set_interpolation_loop_wrap, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the track at [param track_idx] wraps the interpolation loop. New tracks wrap the interpolation loop by default.
*/
//go:nosplit
func (self class) TrackGetInterpolationLoopWrap(track_idx int64) bool { //gd:Animation.track_get_interpolation_loop_wrap
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_get_interpolation_loop_wrap, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the track is compressed, [code]false[/code] otherwise. See also [method compress].
*/
//go:nosplit
func (self class) TrackIsCompressed(track_idx int64) bool { //gd:Animation.track_is_compressed
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_track_is_compressed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the update mode (see [enum UpdateMode]) of a value track.
*/
//go:nosplit
func (self class) ValueTrackSetUpdateMode(track_idx int64, mode gdclass.AnimationUpdateMode) { //gd:Animation.value_track_set_update_mode
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_value_track_set_update_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the update mode of a value track.
*/
//go:nosplit
func (self class) ValueTrackGetUpdateMode(track_idx int64) gdclass.AnimationUpdateMode { //gd:Animation.value_track_get_update_mode
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Ret[gdclass.AnimationUpdateMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_value_track_get_update_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the interpolated value at the given time (in seconds). The [param track_idx] must be the index of a value track.
A [param backward] mainly affects the direction of key retrieval of the track with [constant UPDATE_DISCRETE] converted by [constant AnimationMixer.ANIMATION_CALLBACK_MODE_DISCRETE_FORCE_CONTINUOUS] to match the result with [method track_find_key].
*/
//go:nosplit
func (self class) ValueTrackInterpolate(track_idx int64, time_sec float64, backward bool) variant.Any { //gd:Animation.value_track_interpolate
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time_sec)
	callframe.Arg(frame, backward)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_value_track_interpolate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the method name of a method track.
*/
//go:nosplit
func (self class) MethodTrackGetName(track_idx int64, key_idx int64) String.Name { //gd:Animation.method_track_get_name
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_method_track_get_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the arguments values to be called on a method track for a given key in a given track.
*/
//go:nosplit
func (self class) MethodTrackGetParams(track_idx int64, key_idx int64) Array.Any { //gd:Animation.method_track_get_params
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_method_track_get_params, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Inserts a Bezier Track key at the given [param time] in seconds. The [param track_idx] must be the index of a Bezier Track.
[param in_handle] is the left-side weight of the added Bezier curve point, [param out_handle] is the right-side one, while [param value] is the actual value at this point.
*/
//go:nosplit
func (self class) BezierTrackInsertKey(track_idx int64, time float64, value float64, in_handle Vector2.XY, out_handle Vector2.XY) int64 { //gd:Animation.bezier_track_insert_key
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	callframe.Arg(frame, value)
	callframe.Arg(frame, in_handle)
	callframe.Arg(frame, out_handle)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_bezier_track_insert_key, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the value of the key identified by [param key_idx] to the given value. The [param track_idx] must be the index of a Bezier Track.
*/
//go:nosplit
func (self class) BezierTrackSetKeyValue(track_idx int64, key_idx int64, value float64) { //gd:Animation.bezier_track_set_key_value
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_bezier_track_set_key_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the in handle of the key identified by [param key_idx] to value [param in_handle]. The [param track_idx] must be the index of a Bezier Track.
*/
//go:nosplit
func (self class) BezierTrackSetKeyInHandle(track_idx int64, key_idx int64, in_handle Vector2.XY, balanced_value_time_ratio float64) { //gd:Animation.bezier_track_set_key_in_handle
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	callframe.Arg(frame, in_handle)
	callframe.Arg(frame, balanced_value_time_ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_bezier_track_set_key_in_handle, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the out handle of the key identified by [param key_idx] to value [param out_handle]. The [param track_idx] must be the index of a Bezier Track.
*/
//go:nosplit
func (self class) BezierTrackSetKeyOutHandle(track_idx int64, key_idx int64, out_handle Vector2.XY, balanced_value_time_ratio float64) { //gd:Animation.bezier_track_set_key_out_handle
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	callframe.Arg(frame, out_handle)
	callframe.Arg(frame, balanced_value_time_ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_bezier_track_set_key_out_handle, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the value of the key identified by [param key_idx]. The [param track_idx] must be the index of a Bezier Track.
*/
//go:nosplit
func (self class) BezierTrackGetKeyValue(track_idx int64, key_idx int64) float64 { //gd:Animation.bezier_track_get_key_value
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_bezier_track_get_key_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the in handle of the key identified by [param key_idx]. The [param track_idx] must be the index of a Bezier Track.
*/
//go:nosplit
func (self class) BezierTrackGetKeyInHandle(track_idx int64, key_idx int64) Vector2.XY { //gd:Animation.bezier_track_get_key_in_handle
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_bezier_track_get_key_in_handle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the out handle of the key identified by [param key_idx]. The [param track_idx] must be the index of a Bezier Track.
*/
//go:nosplit
func (self class) BezierTrackGetKeyOutHandle(track_idx int64, key_idx int64) Vector2.XY { //gd:Animation.bezier_track_get_key_out_handle
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_bezier_track_get_key_out_handle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the interpolated value at the given [param time] (in seconds). The [param track_idx] must be the index of a Bezier Track.
*/
//go:nosplit
func (self class) BezierTrackInterpolate(track_idx int64, time float64) float64 { //gd:Animation.bezier_track_interpolate
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_bezier_track_interpolate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Inserts an Audio Track key at the given [param time] in seconds. The [param track_idx] must be the index of an Audio Track.
[param stream] is the [AudioStream] resource to play. [param start_offset] is the number of seconds cut off at the beginning of the audio stream, while [param end_offset] is at the ending.
*/
//go:nosplit
func (self class) AudioTrackInsertKey(track_idx int64, time float64, stream [1]gdclass.Resource, start_offset float64, end_offset float64) int64 { //gd:Animation.audio_track_insert_key
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	callframe.Arg(frame, pointers.Get(stream[0])[0])
	callframe.Arg(frame, start_offset)
	callframe.Arg(frame, end_offset)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_audio_track_insert_key, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the stream of the key identified by [param key_idx] to value [param stream]. The [param track_idx] must be the index of an Audio Track.
*/
//go:nosplit
func (self class) AudioTrackSetKeyStream(track_idx int64, key_idx int64, stream [1]gdclass.Resource) { //gd:Animation.audio_track_set_key_stream
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	callframe.Arg(frame, pointers.Get(stream[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_audio_track_set_key_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the start offset of the key identified by [param key_idx] to value [param offset]. The [param track_idx] must be the index of an Audio Track.
*/
//go:nosplit
func (self class) AudioTrackSetKeyStartOffset(track_idx int64, key_idx int64, offset float64) { //gd:Animation.audio_track_set_key_start_offset
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_audio_track_set_key_start_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the end offset of the key identified by [param key_idx] to value [param offset]. The [param track_idx] must be the index of an Audio Track.
*/
//go:nosplit
func (self class) AudioTrackSetKeyEndOffset(track_idx int64, key_idx int64, offset float64) { //gd:Animation.audio_track_set_key_end_offset
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_audio_track_set_key_end_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the audio stream of the key identified by [param key_idx]. The [param track_idx] must be the index of an Audio Track.
*/
//go:nosplit
func (self class) AudioTrackGetKeyStream(track_idx int64, key_idx int64) [1]gdclass.Resource { //gd:Animation.audio_track_get_key_stream
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_audio_track_get_key_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Resource{gd.PointerWithOwnershipTransferredToGo[gdclass.Resource](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the start offset of the key identified by [param key_idx]. The [param track_idx] must be the index of an Audio Track.
Start offset is the number of seconds cut off at the beginning of the audio stream.
*/
//go:nosplit
func (self class) AudioTrackGetKeyStartOffset(track_idx int64, key_idx int64) float64 { //gd:Animation.audio_track_get_key_start_offset
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_audio_track_get_key_start_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the end offset of the key identified by [param key_idx]. The [param track_idx] must be the index of an Audio Track.
End offset is the number of seconds cut off at the ending of the audio stream.
*/
//go:nosplit
func (self class) AudioTrackGetKeyEndOffset(track_idx int64, key_idx int64) float64 { //gd:Animation.audio_track_get_key_end_offset
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_audio_track_get_key_end_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether the track will be blended with other animations. If [code]true[/code], the audio playback volume changes depending on the blend value.
*/
//go:nosplit
func (self class) AudioTrackSetUseBlend(track_idx int64, enable bool) { //gd:Animation.audio_track_set_use_blend
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_audio_track_set_use_blend, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the track at [param track_idx] will be blended with other animations.
*/
//go:nosplit
func (self class) AudioTrackIsUseBlend(track_idx int64) bool { //gd:Animation.audio_track_is_use_blend
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_audio_track_is_use_blend, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Inserts a key with value [param animation] at the given [param time] (in seconds). The [param track_idx] must be the index of an Animation Track.
*/
//go:nosplit
func (self class) AnimationTrackInsertKey(track_idx int64, time float64, animation String.Name) int64 { //gd:Animation.animation_track_insert_key
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(animation)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_animation_track_insert_key, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the key identified by [param key_idx] to value [param animation]. The [param track_idx] must be the index of an Animation Track.
*/
//go:nosplit
func (self class) AnimationTrackSetKeyAnimation(track_idx int64, key_idx int64, animation String.Name) { //gd:Animation.animation_track_set_key_animation
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(animation)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_animation_track_set_key_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the animation name at the key identified by [param key_idx]. The [param track_idx] must be the index of an Animation Track.
*/
//go:nosplit
func (self class) AnimationTrackGetKeyAnimation(track_idx int64, key_idx int64) String.Name { //gd:Animation.animation_track_get_key_animation
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_animation_track_get_key_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLength(time_sec float64) { //gd:Animation.set_length
	var frame = callframe.New()
	callframe.Arg(frame, time_sec)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_set_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLength() float64 { //gd:Animation.get_length
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_get_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLoopMode(loop_mode gdclass.AnimationLoopMode) { //gd:Animation.set_loop_mode
	var frame = callframe.New()
	callframe.Arg(frame, loop_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_set_loop_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLoopMode() gdclass.AnimationLoopMode { //gd:Animation.get_loop_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AnimationLoopMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_get_loop_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStep(size_sec float64) { //gd:Animation.set_step
	var frame = callframe.New()
	callframe.Arg(frame, size_sec)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_set_step, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStep() float64 { //gd:Animation.get_step
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_get_step, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Clear the animation (clear all tracks and reset all).
*/
//go:nosplit
func (self class) Clear() { //gd:Animation.clear
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a new track to [param to_animation] that is a copy of the given track from this animation.
*/
//go:nosplit
func (self class) CopyTrack(track_idx int64, to_animation [1]gdclass.Animation) { //gd:Animation.copy_track
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, pointers.Get(to_animation[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_copy_track, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Compress the animation and all its tracks in-place. This will make [method track_is_compressed] return [code]true[/code] once called on this [Animation]. Compressed tracks require less memory to be played, and are designed to be used for complex 3D animations (such as cutscenes) imported from external 3D software. Compression is lossy, but the difference is usually not noticeable in real world conditions.
[b]Note:[/b] Compressed tracks have various limitations (such as not being editable from the editor), so only use compressed animations if you actually need them.
*/
//go:nosplit
func (self class) Compress(page_size int64, fps int64, split_tolerance float64) { //gd:Animation.compress
	var frame = callframe.New()
	callframe.Arg(frame, page_size)
	callframe.Arg(frame, fps)
	callframe.Arg(frame, split_tolerance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_compress, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCaptureIncluded() bool { //gd:Animation.is_capture_included
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Animation.Bind_is_capture_included, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAnimation() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimation() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("Animation", func(ptr gd.Object) any { return [1]gdclass.Animation{*(*gdclass.Animation)(unsafe.Pointer(&ptr))} })
}

type TrackType = gdclass.AnimationTrackType //gd:Animation.TrackType

const (
	/*Value tracks set values in node properties, but only those which can be interpolated. For 3D position/rotation/scale, using the dedicated [constant TYPE_POSITION_3D], [constant TYPE_ROTATION_3D] and [constant TYPE_SCALE_3D] track types instead of [constant TYPE_VALUE] is recommended for performance reasons.*/
	TypeValue TrackType = 0
	/*3D position track (values are stored in [Vector3]s).*/
	TypePosition3d TrackType = 1
	/*3D rotation track (values are stored in [Quaternion]s).*/
	TypeRotation3d TrackType = 2
	/*3D scale track (values are stored in [Vector3]s).*/
	TypeScale3d TrackType = 3
	/*Blend shape track.*/
	TypeBlendShape TrackType = 4
	/*Method tracks call functions with given arguments per key.*/
	TypeMethod TrackType = 5
	/*Bezier tracks are used to interpolate a value using custom curves. They can also be used to animate sub-properties of vectors and colors (e.g. alpha value of a [Color]).*/
	TypeBezier TrackType = 6
	/*Audio tracks are used to play an audio stream with either type of [AudioStreamPlayer]. The stream can be trimmed and previewed in the animation.*/
	TypeAudio TrackType = 7
	/*Animation tracks play animations in other [AnimationPlayer] nodes.*/
	TypeAnimation TrackType = 8
)

type InterpolationType = gdclass.AnimationInterpolationType //gd:Animation.InterpolationType

const (
	/*No interpolation (nearest value).*/
	InterpolationNearest InterpolationType = 0
	/*Linear interpolation.*/
	InterpolationLinear InterpolationType = 1
	/*Cubic interpolation. This looks smoother than linear interpolation, but is more expensive to interpolate. Stick to [constant INTERPOLATION_LINEAR] for complex 3D animations imported from external software, even if it requires using a higher animation framerate in return.*/
	InterpolationCubic InterpolationType = 2
	/*Linear interpolation with shortest path rotation.
	  [b]Note:[/b] The result value is always normalized and may not match the key value.*/
	InterpolationLinearAngle InterpolationType = 3
	/*Cubic interpolation with shortest path rotation.
	  [b]Note:[/b] The result value is always normalized and may not match the key value.*/
	InterpolationCubicAngle InterpolationType = 4
)

type UpdateMode = gdclass.AnimationUpdateMode //gd:Animation.UpdateMode

const (
	/*Update between keyframes and hold the value.*/
	UpdateContinuous UpdateMode = 0
	/*Update at the keyframes.*/
	UpdateDiscrete UpdateMode = 1
	/*Same as [constant UPDATE_CONTINUOUS] but works as a flag to capture the value of the current object and perform interpolation in some methods. See also [method AnimationMixer.capture], [member AnimationPlayer.playback_auto_capture], and [method AnimationPlayer.play_with_capture].*/
	UpdateCapture UpdateMode = 2
)

type LoopMode = gdclass.AnimationLoopMode //gd:Animation.LoopMode

const (
	/*At both ends of the animation, the animation will stop playing.*/
	LoopNone LoopMode = 0
	/*At both ends of the animation, the animation will be repeated without changing the playback direction.*/
	LoopLinear LoopMode = 1
	/*Repeats playback and reverse playback at both ends of the animation.*/
	LoopPingpong LoopMode = 2
)

type LoopedFlag = gdclass.AnimationLoopedFlag //gd:Animation.LoopedFlag

const (
	/*This flag indicates that the animation proceeds without any looping.*/
	LoopedFlagNone LoopedFlag = 0
	/*This flag indicates that the animation has reached the end of the animation and just after loop processed.*/
	LoopedFlagEnd LoopedFlag = 1
	/*This flag indicates that the animation has reached the start of the animation and just after loop processed.*/
	LoopedFlagStart LoopedFlag = 2
)

type FindMode = gdclass.AnimationFindMode //gd:Animation.FindMode

const (
	/*Finds the nearest time key.*/
	FindModeNearest FindMode = 0
	/*Finds only the key with approximating the time.*/
	FindModeApprox FindMode = 1
	/*Finds only the key with matching the time.*/
	FindModeExact FindMode = 2
)
