package Animation

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
type Go [1]classdb.Animation

/*
Adds a track to the Animation.
*/
func (self Go) AddTrack(atype classdb.AnimationTrackType) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).AddTrack(atype, gd.Int(-1))))
}

/*
Removes a track by specifying the track index.
*/
func (self Go) RemoveTrack(track_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveTrack(gd.Int(track_idx))
}

/*
Returns the amount of tracks in the animation.
*/
func (self Go) GetTrackCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetTrackCount()))
}

/*
Gets the type of a track.
*/
func (self Go) TrackGetType(track_idx int) classdb.AnimationTrackType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AnimationTrackType(class(self).TrackGetType(gd.Int(track_idx)))
}

/*
Gets the path of a track. For more information on the path format, see [method track_set_path].
*/
func (self Go) TrackGetPath(track_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).TrackGetPath(gc, gd.Int(track_idx)).String())
}

/*
Sets the path of a track. Paths must be valid scene-tree paths to a node and must be specified starting from the [member AnimationMixer.root_node] that will reproduce the animation. Tracks that control properties or bones must append their name after the path, separated by [code]":"[/code].
For example, [code]"character/skeleton:ankle"[/code] or [code]"character/mesh:transform/local"[/code].
*/
func (self Go) TrackSetPath(track_idx int, path string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).TrackSetPath(gd.Int(track_idx), gc.String(path).NodePath(gc))
}

/*
Returns the index of the specified track. If the track is not found, return -1.
*/
func (self Go) FindTrack(path string, atype classdb.AnimationTrackType) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).FindTrack(gc.String(path).NodePath(gc), atype)))
}

/*
Moves a track up.
*/
func (self Go) TrackMoveUp(track_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).TrackMoveUp(gd.Int(track_idx))
}

/*
Moves a track down.
*/
func (self Go) TrackMoveDown(track_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).TrackMoveDown(gd.Int(track_idx))
}

/*
Changes the index position of track [param track_idx] to the one defined in [param to_idx].
*/
func (self Go) TrackMoveTo(track_idx int, to_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).TrackMoveTo(gd.Int(track_idx), gd.Int(to_idx))
}

/*
Swaps the track [param track_idx]'s index position with the track [param with_idx].
*/
func (self Go) TrackSwap(track_idx int, with_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).TrackSwap(gd.Int(track_idx), gd.Int(with_idx))
}

/*
Sets the given track as imported or not.
*/
func (self Go) TrackSetImported(track_idx int, imported bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).TrackSetImported(gd.Int(track_idx), imported)
}

/*
Returns [code]true[/code] if the given track is imported. Else, return [code]false[/code].
*/
func (self Go) TrackIsImported(track_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).TrackIsImported(gd.Int(track_idx)))
}

/*
Enables/disables the given track. Tracks are enabled by default.
*/
func (self Go) TrackSetEnabled(track_idx int, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).TrackSetEnabled(gd.Int(track_idx), enabled)
}

/*
Returns [code]true[/code] if the track at index [param track_idx] is enabled.
*/
func (self Go) TrackIsEnabled(track_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).TrackIsEnabled(gd.Int(track_idx)))
}

/*
Inserts a key in a given 3D position track. Returns the key index.
*/
func (self Go) PositionTrackInsertKey(track_idx int, time float64, position gd.Vector3) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).PositionTrackInsertKey(gd.Int(track_idx), gd.Float(time), position)))
}

/*
Inserts a key in a given 3D rotation track. Returns the key index.
*/
func (self Go) RotationTrackInsertKey(track_idx int, time float64, rotation gd.Quaternion) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).RotationTrackInsertKey(gd.Int(track_idx), gd.Float(time), rotation)))
}

/*
Inserts a key in a given 3D scale track. Returns the key index.
*/
func (self Go) ScaleTrackInsertKey(track_idx int, time float64, scale gd.Vector3) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ScaleTrackInsertKey(gd.Int(track_idx), gd.Float(time), scale)))
}

/*
Inserts a key in a given blend shape track. Returns the key index.
*/
func (self Go) BlendShapeTrackInsertKey(track_idx int, time float64, amount float64) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).BlendShapeTrackInsertKey(gd.Int(track_idx), gd.Float(time), gd.Float(amount))))
}

/*
Returns the interpolated position value at the given time (in seconds). The [param track_idx] must be the index of a 3D position track.
*/
func (self Go) PositionTrackInterpolate(track_idx int, time_sec float64) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).PositionTrackInterpolate(gd.Int(track_idx), gd.Float(time_sec), false))
}

/*
Returns the interpolated rotation value at the given time (in seconds). The [param track_idx] must be the index of a 3D rotation track.
*/
func (self Go) RotationTrackInterpolate(track_idx int, time_sec float64) gd.Quaternion {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Quaternion(class(self).RotationTrackInterpolate(gd.Int(track_idx), gd.Float(time_sec), false))
}

/*
Returns the interpolated scale value at the given time (in seconds). The [param track_idx] must be the index of a 3D scale track.
*/
func (self Go) ScaleTrackInterpolate(track_idx int, time_sec float64) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).ScaleTrackInterpolate(gd.Int(track_idx), gd.Float(time_sec), false))
}

/*
Returns the interpolated blend shape value at the given time (in seconds). The [param track_idx] must be the index of a blend shape track.
*/
func (self Go) BlendShapeTrackInterpolate(track_idx int, time_sec float64) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).BlendShapeTrackInterpolate(gd.Int(track_idx), gd.Float(time_sec), false)))
}

/*
Inserts a generic key in a given track. Returns the key index.
*/
func (self Go) TrackInsertKey(track_idx int, time float64, key gd.Variant) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).TrackInsertKey(gd.Int(track_idx), gd.Float(time), key, gd.Float(1))))
}

/*
Removes a key by index in a given track.
*/
func (self Go) TrackRemoveKey(track_idx int, key_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).TrackRemoveKey(gd.Int(track_idx), gd.Int(key_idx))
}

/*
Removes a key at [param time] in a given track.
*/
func (self Go) TrackRemoveKeyAtTime(track_idx int, time float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).TrackRemoveKeyAtTime(gd.Int(track_idx), gd.Float(time))
}

/*
Sets the value of an existing key.
*/
func (self Go) TrackSetKeyValue(track_idx int, key int, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).TrackSetKeyValue(gd.Int(track_idx), gd.Int(key), value)
}

/*
Sets the transition curve (easing) for a specific key (see the built-in math function [method @GlobalScope.ease]).
*/
func (self Go) TrackSetKeyTransition(track_idx int, key_idx int, transition float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).TrackSetKeyTransition(gd.Int(track_idx), gd.Int(key_idx), gd.Float(transition))
}

/*
Sets the time of an existing key.
*/
func (self Go) TrackSetKeyTime(track_idx int, key_idx int, time float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).TrackSetKeyTime(gd.Int(track_idx), gd.Int(key_idx), gd.Float(time))
}

/*
Returns the transition curve (easing) for a specific key (see the built-in math function [method @GlobalScope.ease]).
*/
func (self Go) TrackGetKeyTransition(track_idx int, key_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).TrackGetKeyTransition(gd.Int(track_idx), gd.Int(key_idx))))
}

/*
Returns the number of keys in a given track.
*/
func (self Go) TrackGetKeyCount(track_idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).TrackGetKeyCount(gd.Int(track_idx))))
}

/*
Returns the value of a given key in a given track.
*/
func (self Go) TrackGetKeyValue(track_idx int, key_idx int) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).TrackGetKeyValue(gc, gd.Int(track_idx), gd.Int(key_idx)))
}

/*
Returns the time at which the key is located.
*/
func (self Go) TrackGetKeyTime(track_idx int, key_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).TrackGetKeyTime(gd.Int(track_idx), gd.Int(key_idx))))
}

/*
Finds the key index by time in a given track. Optionally, only find it if the approx/exact time is given.
If [param limit] is [code]true[/code], it does not return keys outside the animation range.
If [param backward] is [code]true[/code], the direction is reversed in methods that rely on one directional processing.
For example, in case [param find_mode] is [constant FIND_MODE_NEAREST], if there is no key in the current position just after seeked, the first key found is retrieved by searching before the position, but if [param backward] is [code]true[/code], the first key found is retrieved after the position.
*/
func (self Go) TrackFindKey(track_idx int, time float64) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).TrackFindKey(gd.Int(track_idx), gd.Float(time), 0, false, false)))
}

/*
Sets the interpolation type of a given track.
*/
func (self Go) TrackSetInterpolationType(track_idx int, interpolation classdb.AnimationInterpolationType) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).TrackSetInterpolationType(gd.Int(track_idx), interpolation)
}

/*
Returns the interpolation type of a given track.
*/
func (self Go) TrackGetInterpolationType(track_idx int) classdb.AnimationInterpolationType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AnimationInterpolationType(class(self).TrackGetInterpolationType(gd.Int(track_idx)))
}

/*
If [code]true[/code], the track at [param track_idx] wraps the interpolation loop.
*/
func (self Go) TrackSetInterpolationLoopWrap(track_idx int, interpolation bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).TrackSetInterpolationLoopWrap(gd.Int(track_idx), interpolation)
}

/*
Returns [code]true[/code] if the track at [param track_idx] wraps the interpolation loop. New tracks wrap the interpolation loop by default.
*/
func (self Go) TrackGetInterpolationLoopWrap(track_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).TrackGetInterpolationLoopWrap(gd.Int(track_idx)))
}

/*
Returns [code]true[/code] if the track is compressed, [code]false[/code] otherwise. See also [method compress].
*/
func (self Go) TrackIsCompressed(track_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).TrackIsCompressed(gd.Int(track_idx)))
}

/*
Sets the update mode (see [enum UpdateMode]) of a value track.
*/
func (self Go) ValueTrackSetUpdateMode(track_idx int, mode classdb.AnimationUpdateMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ValueTrackSetUpdateMode(gd.Int(track_idx), mode)
}

/*
Returns the update mode of a value track.
*/
func (self Go) ValueTrackGetUpdateMode(track_idx int) classdb.AnimationUpdateMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AnimationUpdateMode(class(self).ValueTrackGetUpdateMode(gd.Int(track_idx)))
}

/*
Returns the interpolated value at the given time (in seconds). The [param track_idx] must be the index of a value track.
A [param backward] mainly affects the direction of key retrieval of the track with [constant UPDATE_DISCRETE] converted by [constant AnimationMixer.ANIMATION_CALLBACK_MODE_DISCRETE_FORCE_CONTINUOUS] to match the result with [method track_find_key].
*/
func (self Go) ValueTrackInterpolate(track_idx int, time_sec float64) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).ValueTrackInterpolate(gc, gd.Int(track_idx), gd.Float(time_sec), false))
}

/*
Returns the method name of a method track.
*/
func (self Go) MethodTrackGetName(track_idx int, key_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).MethodTrackGetName(gc, gd.Int(track_idx), gd.Int(key_idx)).String())
}

/*
Returns the arguments values to be called on a method track for a given key in a given track.
*/
func (self Go) MethodTrackGetParams(track_idx int, key_idx int) gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(class(self).MethodTrackGetParams(gc, gd.Int(track_idx), gd.Int(key_idx)))
}

/*
Inserts a Bezier Track key at the given [param time] in seconds. The [param track_idx] must be the index of a Bezier Track.
[param in_handle] is the left-side weight of the added Bezier curve point, [param out_handle] is the right-side one, while [param value] is the actual value at this point.
*/
func (self Go) BezierTrackInsertKey(track_idx int, time float64, value float64) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).BezierTrackInsertKey(gd.Int(track_idx), gd.Float(time), gd.Float(value), gd.Vector2{0, 0}, gd.Vector2{0, 0})))
}

/*
Sets the value of the key identified by [param key_idx] to the given value. The [param track_idx] must be the index of a Bezier Track.
*/
func (self Go) BezierTrackSetKeyValue(track_idx int, key_idx int, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).BezierTrackSetKeyValue(gd.Int(track_idx), gd.Int(key_idx), gd.Float(value))
}

/*
Sets the in handle of the key identified by [param key_idx] to value [param in_handle]. The [param track_idx] must be the index of a Bezier Track.
*/
func (self Go) BezierTrackSetKeyInHandle(track_idx int, key_idx int, in_handle gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).BezierTrackSetKeyInHandle(gd.Int(track_idx), gd.Int(key_idx), in_handle, gd.Float(1.0))
}

/*
Sets the out handle of the key identified by [param key_idx] to value [param out_handle]. The [param track_idx] must be the index of a Bezier Track.
*/
func (self Go) BezierTrackSetKeyOutHandle(track_idx int, key_idx int, out_handle gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).BezierTrackSetKeyOutHandle(gd.Int(track_idx), gd.Int(key_idx), out_handle, gd.Float(1.0))
}

/*
Returns the value of the key identified by [param key_idx]. The [param track_idx] must be the index of a Bezier Track.
*/
func (self Go) BezierTrackGetKeyValue(track_idx int, key_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).BezierTrackGetKeyValue(gd.Int(track_idx), gd.Int(key_idx))))
}

/*
Returns the in handle of the key identified by [param key_idx]. The [param track_idx] must be the index of a Bezier Track.
*/
func (self Go) BezierTrackGetKeyInHandle(track_idx int, key_idx int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).BezierTrackGetKeyInHandle(gd.Int(track_idx), gd.Int(key_idx)))
}

/*
Returns the out handle of the key identified by [param key_idx]. The [param track_idx] must be the index of a Bezier Track.
*/
func (self Go) BezierTrackGetKeyOutHandle(track_idx int, key_idx int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).BezierTrackGetKeyOutHandle(gd.Int(track_idx), gd.Int(key_idx)))
}

/*
Returns the interpolated value at the given [param time] (in seconds). The [param track_idx] must be the index of a Bezier Track.
*/
func (self Go) BezierTrackInterpolate(track_idx int, time float64) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).BezierTrackInterpolate(gd.Int(track_idx), gd.Float(time))))
}

/*
Inserts an Audio Track key at the given [param time] in seconds. The [param track_idx] must be the index of an Audio Track.
[param stream] is the [AudioStream] resource to play. [param start_offset] is the number of seconds cut off at the beginning of the audio stream, while [param end_offset] is at the ending.
*/
func (self Go) AudioTrackInsertKey(track_idx int, time float64, stream gdclass.Resource) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).AudioTrackInsertKey(gd.Int(track_idx), gd.Float(time), stream, gd.Float(0), gd.Float(0))))
}

/*
Sets the stream of the key identified by [param key_idx] to value [param stream]. The [param track_idx] must be the index of an Audio Track.
*/
func (self Go) AudioTrackSetKeyStream(track_idx int, key_idx int, stream gdclass.Resource) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AudioTrackSetKeyStream(gd.Int(track_idx), gd.Int(key_idx), stream)
}

/*
Sets the start offset of the key identified by [param key_idx] to value [param offset]. The [param track_idx] must be the index of an Audio Track.
*/
func (self Go) AudioTrackSetKeyStartOffset(track_idx int, key_idx int, offset float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AudioTrackSetKeyStartOffset(gd.Int(track_idx), gd.Int(key_idx), gd.Float(offset))
}

/*
Sets the end offset of the key identified by [param key_idx] to value [param offset]. The [param track_idx] must be the index of an Audio Track.
*/
func (self Go) AudioTrackSetKeyEndOffset(track_idx int, key_idx int, offset float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AudioTrackSetKeyEndOffset(gd.Int(track_idx), gd.Int(key_idx), gd.Float(offset))
}

/*
Returns the audio stream of the key identified by [param key_idx]. The [param track_idx] must be the index of an Audio Track.
*/
func (self Go) AudioTrackGetKeyStream(track_idx int, key_idx int) gdclass.Resource {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Resource(class(self).AudioTrackGetKeyStream(gc, gd.Int(track_idx), gd.Int(key_idx)))
}

/*
Returns the start offset of the key identified by [param key_idx]. The [param track_idx] must be the index of an Audio Track.
Start offset is the number of seconds cut off at the beginning of the audio stream.
*/
func (self Go) AudioTrackGetKeyStartOffset(track_idx int, key_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).AudioTrackGetKeyStartOffset(gd.Int(track_idx), gd.Int(key_idx))))
}

/*
Returns the end offset of the key identified by [param key_idx]. The [param track_idx] must be the index of an Audio Track.
End offset is the number of seconds cut off at the ending of the audio stream.
*/
func (self Go) AudioTrackGetKeyEndOffset(track_idx int, key_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).AudioTrackGetKeyEndOffset(gd.Int(track_idx), gd.Int(key_idx))))
}

/*
Sets whether the track will be blended with other animations. If [code]true[/code], the audio playback volume changes depending on the blend value.
*/
func (self Go) AudioTrackSetUseBlend(track_idx int, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AudioTrackSetUseBlend(gd.Int(track_idx), enable)
}

/*
Returns [code]true[/code] if the track at [param track_idx] will be blended with other animations.
*/
func (self Go) AudioTrackIsUseBlend(track_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).AudioTrackIsUseBlend(gd.Int(track_idx)))
}

/*
Inserts a key with value [param animation] at the given [param time] (in seconds). The [param track_idx] must be the index of an Animation Track.
*/
func (self Go) AnimationTrackInsertKey(track_idx int, time float64, animation string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).AnimationTrackInsertKey(gd.Int(track_idx), gd.Float(time), gc.StringName(animation))))
}

/*
Sets the key identified by [param key_idx] to value [param animation]. The [param track_idx] must be the index of an Animation Track.
*/
func (self Go) AnimationTrackSetKeyAnimation(track_idx int, key_idx int, animation string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AnimationTrackSetKeyAnimation(gd.Int(track_idx), gd.Int(key_idx), gc.StringName(animation))
}

/*
Returns the animation name at the key identified by [param key_idx]. The [param track_idx] must be the index of an Animation Track.
*/
func (self Go) AnimationTrackGetKeyAnimation(track_idx int, key_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).AnimationTrackGetKeyAnimation(gc, gd.Int(track_idx), gd.Int(key_idx)).String())
}

/*
Clear the animation (clear all tracks and reset all).
*/
func (self Go) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Clear()
}

/*
Adds a new track to [param to_animation] that is a copy of the given track from this animation.
*/
func (self Go) CopyTrack(track_idx int, to_animation gdclass.Animation) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).CopyTrack(gd.Int(track_idx), to_animation)
}

/*
Compress the animation and all its tracks in-place. This will make [method track_is_compressed] return [code]true[/code] once called on this [Animation]. Compressed tracks require less memory to be played, and are designed to be used for complex 3D animations (such as cutscenes) imported from external 3D software. Compression is lossy, but the difference is usually not noticeable in real world conditions.
[b]Note:[/b] Compressed tracks have various limitations (such as not being editable from the editor), so only use compressed animations if you actually need them.
*/
func (self Go) Compress() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Compress(gd.Int(8192), gd.Int(120), gd.Float(4.0))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Animation
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Animation"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Length() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetLength()))
}

func (self Go) SetLength(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLength(gd.Float(value))
}

func (self Go) LoopMode() classdb.AnimationLoopMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AnimationLoopMode(class(self).GetLoopMode())
}

func (self Go) SetLoopMode(value classdb.AnimationLoopMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLoopMode(value)
}

func (self Go) Step() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetStep()))
}

func (self Go) SetStep(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStep(gd.Float(value))
}

func (self Go) CaptureIncluded() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsCaptureIncluded())
}

/*
Adds a track to the Animation.
*/
//go:nosplit
func (self class) AddTrack(atype classdb.AnimationTrackType, at_position gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, at_position)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_add_track, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes a track by specifying the track index.
*/
//go:nosplit
func (self class) RemoveTrack(track_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_remove_track, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the amount of tracks in the animation.
*/
//go:nosplit
func (self class) GetTrackCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_get_track_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets the type of a track.
*/
//go:nosplit
func (self class) TrackGetType(track_idx gd.Int) classdb.AnimationTrackType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Ret[classdb.AnimationTrackType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_get_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets the path of a track. For more information on the path format, see [method track_set_path].
*/
//go:nosplit
func (self class) TrackGetPath(ctx gd.Lifetime, track_idx gd.Int) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_get_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the path of a track. Paths must be valid scene-tree paths to a node and must be specified starting from the [member AnimationMixer.root_node] that will reproduce the animation. Tracks that control properties or bones must append their name after the path, separated by [code]":"[/code].
For example, [code]"character/skeleton:ankle"[/code] or [code]"character/mesh:transform/local"[/code].
*/
//go:nosplit
func (self class) TrackSetPath(track_idx gd.Int, path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_set_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the index of the specified track. If the track is not found, return -1.
*/
//go:nosplit
func (self class) FindTrack(path gd.NodePath, atype classdb.AnimationTrackType) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_find_track, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Moves a track up.
*/
//go:nosplit
func (self class) TrackMoveUp(track_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_move_up, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Moves a track down.
*/
//go:nosplit
func (self class) TrackMoveDown(track_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_move_down, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Changes the index position of track [param track_idx] to the one defined in [param to_idx].
*/
//go:nosplit
func (self class) TrackMoveTo(track_idx gd.Int, to_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, to_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_move_to, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Swaps the track [param track_idx]'s index position with the track [param with_idx].
*/
//go:nosplit
func (self class) TrackSwap(track_idx gd.Int, with_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, with_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_swap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the given track as imported or not.
*/
//go:nosplit
func (self class) TrackSetImported(track_idx gd.Int, imported bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, imported)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_set_imported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the given track is imported. Else, return [code]false[/code].
*/
//go:nosplit
func (self class) TrackIsImported(track_idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_is_imported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Enables/disables the given track. Tracks are enabled by default.
*/
//go:nosplit
func (self class) TrackSetEnabled(track_idx gd.Int, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the track at index [param track_idx] is enabled.
*/
//go:nosplit
func (self class) TrackIsEnabled(track_idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_is_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Inserts a key in a given 3D position track. Returns the key index.
*/
//go:nosplit
func (self class) PositionTrackInsertKey(track_idx gd.Int, time gd.Float, position gd.Vector3) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_position_track_insert_key, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Inserts a key in a given 3D rotation track. Returns the key index.
*/
//go:nosplit
func (self class) RotationTrackInsertKey(track_idx gd.Int, time gd.Float, rotation gd.Quaternion) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	callframe.Arg(frame, rotation)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_rotation_track_insert_key, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Inserts a key in a given 3D scale track. Returns the key index.
*/
//go:nosplit
func (self class) ScaleTrackInsertKey(track_idx gd.Int, time gd.Float, scale gd.Vector3) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	callframe.Arg(frame, scale)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_scale_track_insert_key, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Inserts a key in a given blend shape track. Returns the key index.
*/
//go:nosplit
func (self class) BlendShapeTrackInsertKey(track_idx gd.Int, time gd.Float, amount gd.Float) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	callframe.Arg(frame, amount)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_blend_shape_track_insert_key, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the interpolated position value at the given time (in seconds). The [param track_idx] must be the index of a 3D position track.
*/
//go:nosplit
func (self class) PositionTrackInterpolate(track_idx gd.Int, time_sec gd.Float, backward bool) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time_sec)
	callframe.Arg(frame, backward)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_position_track_interpolate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the interpolated rotation value at the given time (in seconds). The [param track_idx] must be the index of a 3D rotation track.
*/
//go:nosplit
func (self class) RotationTrackInterpolate(track_idx gd.Int, time_sec gd.Float, backward bool) gd.Quaternion {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time_sec)
	callframe.Arg(frame, backward)
	var r_ret = callframe.Ret[gd.Quaternion](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_rotation_track_interpolate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the interpolated scale value at the given time (in seconds). The [param track_idx] must be the index of a 3D scale track.
*/
//go:nosplit
func (self class) ScaleTrackInterpolate(track_idx gd.Int, time_sec gd.Float, backward bool) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time_sec)
	callframe.Arg(frame, backward)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_scale_track_interpolate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the interpolated blend shape value at the given time (in seconds). The [param track_idx] must be the index of a blend shape track.
*/
//go:nosplit
func (self class) BlendShapeTrackInterpolate(track_idx gd.Int, time_sec gd.Float, backward bool) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time_sec)
	callframe.Arg(frame, backward)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_blend_shape_track_interpolate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Inserts a generic key in a given track. Returns the key index.
*/
//go:nosplit
func (self class) TrackInsertKey(track_idx gd.Int, time gd.Float, key gd.Variant, transition gd.Float) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	callframe.Arg(frame, mmm.Get(key))
	callframe.Arg(frame, transition)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_insert_key, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes a key by index in a given track.
*/
//go:nosplit
func (self class) TrackRemoveKey(track_idx gd.Int, key_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_remove_key, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a key at [param time] in a given track.
*/
//go:nosplit
func (self class) TrackRemoveKeyAtTime(track_idx gd.Int, time gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_remove_key_at_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the value of an existing key.
*/
//go:nosplit
func (self class) TrackSetKeyValue(track_idx gd.Int, key gd.Int, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key)
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_set_key_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the transition curve (easing) for a specific key (see the built-in math function [method @GlobalScope.ease]).
*/
//go:nosplit
func (self class) TrackSetKeyTransition(track_idx gd.Int, key_idx gd.Int, transition gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	callframe.Arg(frame, transition)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_set_key_transition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the time of an existing key.
*/
//go:nosplit
func (self class) TrackSetKeyTime(track_idx gd.Int, key_idx gd.Int, time gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	callframe.Arg(frame, time)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_set_key_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the transition curve (easing) for a specific key (see the built-in math function [method @GlobalScope.ease]).
*/
//go:nosplit
func (self class) TrackGetKeyTransition(track_idx gd.Int, key_idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_get_key_transition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of keys in a given track.
*/
//go:nosplit
func (self class) TrackGetKeyCount(track_idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_get_key_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the value of a given key in a given track.
*/
//go:nosplit
func (self class) TrackGetKeyValue(ctx gd.Lifetime, track_idx gd.Int, key_idx gd.Int) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_get_key_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the time at which the key is located.
*/
//go:nosplit
func (self class) TrackGetKeyTime(track_idx gd.Int, key_idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_get_key_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) TrackFindKey(track_idx gd.Int, time gd.Float, find_mode classdb.AnimationFindMode, limit bool, backward bool) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	callframe.Arg(frame, find_mode)
	callframe.Arg(frame, limit)
	callframe.Arg(frame, backward)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_find_key, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the interpolation type of a given track.
*/
//go:nosplit
func (self class) TrackSetInterpolationType(track_idx gd.Int, interpolation classdb.AnimationInterpolationType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, interpolation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_set_interpolation_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the interpolation type of a given track.
*/
//go:nosplit
func (self class) TrackGetInterpolationType(track_idx gd.Int) classdb.AnimationInterpolationType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Ret[classdb.AnimationInterpolationType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_get_interpolation_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [code]true[/code], the track at [param track_idx] wraps the interpolation loop.
*/
//go:nosplit
func (self class) TrackSetInterpolationLoopWrap(track_idx gd.Int, interpolation bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, interpolation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_set_interpolation_loop_wrap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the track at [param track_idx] wraps the interpolation loop. New tracks wrap the interpolation loop by default.
*/
//go:nosplit
func (self class) TrackGetInterpolationLoopWrap(track_idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_get_interpolation_loop_wrap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the track is compressed, [code]false[/code] otherwise. See also [method compress].
*/
//go:nosplit
func (self class) TrackIsCompressed(track_idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_track_is_compressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the update mode (see [enum UpdateMode]) of a value track.
*/
//go:nosplit
func (self class) ValueTrackSetUpdateMode(track_idx gd.Int, mode classdb.AnimationUpdateMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_value_track_set_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the update mode of a value track.
*/
//go:nosplit
func (self class) ValueTrackGetUpdateMode(track_idx gd.Int) classdb.AnimationUpdateMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Ret[classdb.AnimationUpdateMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_value_track_get_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the interpolated value at the given time (in seconds). The [param track_idx] must be the index of a value track.
A [param backward] mainly affects the direction of key retrieval of the track with [constant UPDATE_DISCRETE] converted by [constant AnimationMixer.ANIMATION_CALLBACK_MODE_DISCRETE_FORCE_CONTINUOUS] to match the result with [method track_find_key].
*/
//go:nosplit
func (self class) ValueTrackInterpolate(ctx gd.Lifetime, track_idx gd.Int, time_sec gd.Float, backward bool) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time_sec)
	callframe.Arg(frame, backward)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_value_track_interpolate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the method name of a method track.
*/
//go:nosplit
func (self class) MethodTrackGetName(ctx gd.Lifetime, track_idx gd.Int, key_idx gd.Int) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_method_track_get_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the arguments values to be called on a method track for a given key in a given track.
*/
//go:nosplit
func (self class) MethodTrackGetParams(ctx gd.Lifetime, track_idx gd.Int, key_idx gd.Int) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_method_track_get_params, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Inserts a Bezier Track key at the given [param time] in seconds. The [param track_idx] must be the index of a Bezier Track.
[param in_handle] is the left-side weight of the added Bezier curve point, [param out_handle] is the right-side one, while [param value] is the actual value at this point.
*/
//go:nosplit
func (self class) BezierTrackInsertKey(track_idx gd.Int, time gd.Float, value gd.Float, in_handle gd.Vector2, out_handle gd.Vector2) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	callframe.Arg(frame, value)
	callframe.Arg(frame, in_handle)
	callframe.Arg(frame, out_handle)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_bezier_track_insert_key, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the value of the key identified by [param key_idx] to the given value. The [param track_idx] must be the index of a Bezier Track.
*/
//go:nosplit
func (self class) BezierTrackSetKeyValue(track_idx gd.Int, key_idx gd.Int, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_bezier_track_set_key_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the in handle of the key identified by [param key_idx] to value [param in_handle]. The [param track_idx] must be the index of a Bezier Track.
*/
//go:nosplit
func (self class) BezierTrackSetKeyInHandle(track_idx gd.Int, key_idx gd.Int, in_handle gd.Vector2, balanced_value_time_ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	callframe.Arg(frame, in_handle)
	callframe.Arg(frame, balanced_value_time_ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_bezier_track_set_key_in_handle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the out handle of the key identified by [param key_idx] to value [param out_handle]. The [param track_idx] must be the index of a Bezier Track.
*/
//go:nosplit
func (self class) BezierTrackSetKeyOutHandle(track_idx gd.Int, key_idx gd.Int, out_handle gd.Vector2, balanced_value_time_ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	callframe.Arg(frame, out_handle)
	callframe.Arg(frame, balanced_value_time_ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_bezier_track_set_key_out_handle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of the key identified by [param key_idx]. The [param track_idx] must be the index of a Bezier Track.
*/
//go:nosplit
func (self class) BezierTrackGetKeyValue(track_idx gd.Int, key_idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_bezier_track_get_key_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the in handle of the key identified by [param key_idx]. The [param track_idx] must be the index of a Bezier Track.
*/
//go:nosplit
func (self class) BezierTrackGetKeyInHandle(track_idx gd.Int, key_idx gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_bezier_track_get_key_in_handle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the out handle of the key identified by [param key_idx]. The [param track_idx] must be the index of a Bezier Track.
*/
//go:nosplit
func (self class) BezierTrackGetKeyOutHandle(track_idx gd.Int, key_idx gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_bezier_track_get_key_out_handle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the interpolated value at the given [param time] (in seconds). The [param track_idx] must be the index of a Bezier Track.
*/
//go:nosplit
func (self class) BezierTrackInterpolate(track_idx gd.Int, time gd.Float) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_bezier_track_interpolate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Inserts an Audio Track key at the given [param time] in seconds. The [param track_idx] must be the index of an Audio Track.
[param stream] is the [AudioStream] resource to play. [param start_offset] is the number of seconds cut off at the beginning of the audio stream, while [param end_offset] is at the ending.
*/
//go:nosplit
func (self class) AudioTrackInsertKey(track_idx gd.Int, time gd.Float, stream gdclass.Resource, start_offset gd.Float, end_offset gd.Float) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	callframe.Arg(frame, mmm.Get(stream[0].AsPointer())[0])
	callframe.Arg(frame, start_offset)
	callframe.Arg(frame, end_offset)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_audio_track_insert_key, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the stream of the key identified by [param key_idx] to value [param stream]. The [param track_idx] must be the index of an Audio Track.
*/
//go:nosplit
func (self class) AudioTrackSetKeyStream(track_idx gd.Int, key_idx gd.Int, stream gdclass.Resource)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	callframe.Arg(frame, mmm.Get(stream[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_audio_track_set_key_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the start offset of the key identified by [param key_idx] to value [param offset]. The [param track_idx] must be the index of an Audio Track.
*/
//go:nosplit
func (self class) AudioTrackSetKeyStartOffset(track_idx gd.Int, key_idx gd.Int, offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_audio_track_set_key_start_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the end offset of the key identified by [param key_idx] to value [param offset]. The [param track_idx] must be the index of an Audio Track.
*/
//go:nosplit
func (self class) AudioTrackSetKeyEndOffset(track_idx gd.Int, key_idx gd.Int, offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_audio_track_set_key_end_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the audio stream of the key identified by [param key_idx]. The [param track_idx] must be the index of an Audio Track.
*/
//go:nosplit
func (self class) AudioTrackGetKeyStream(ctx gd.Lifetime, track_idx gd.Int, key_idx gd.Int) gdclass.Resource {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_audio_track_get_key_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Resource
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the start offset of the key identified by [param key_idx]. The [param track_idx] must be the index of an Audio Track.
Start offset is the number of seconds cut off at the beginning of the audio stream.
*/
//go:nosplit
func (self class) AudioTrackGetKeyStartOffset(track_idx gd.Int, key_idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_audio_track_get_key_start_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the end offset of the key identified by [param key_idx]. The [param track_idx] must be the index of an Audio Track.
End offset is the number of seconds cut off at the ending of the audio stream.
*/
//go:nosplit
func (self class) AudioTrackGetKeyEndOffset(track_idx gd.Int, key_idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_audio_track_get_key_end_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the track will be blended with other animations. If [code]true[/code], the audio playback volume changes depending on the blend value.
*/
//go:nosplit
func (self class) AudioTrackSetUseBlend(track_idx gd.Int, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_audio_track_set_use_blend, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the track at [param track_idx] will be blended with other animations.
*/
//go:nosplit
func (self class) AudioTrackIsUseBlend(track_idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_audio_track_is_use_blend, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Inserts a key with value [param animation] at the given [param time] (in seconds). The [param track_idx] must be the index of an Animation Track.
*/
//go:nosplit
func (self class) AnimationTrackInsertKey(track_idx gd.Int, time gd.Float, animation gd.StringName) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, time)
	callframe.Arg(frame, mmm.Get(animation))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_animation_track_insert_key, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the key identified by [param key_idx] to value [param animation]. The [param track_idx] must be the index of an Animation Track.
*/
//go:nosplit
func (self class) AnimationTrackSetKeyAnimation(track_idx gd.Int, key_idx gd.Int, animation gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	callframe.Arg(frame, mmm.Get(animation))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_animation_track_set_key_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the animation name at the key identified by [param key_idx]. The [param track_idx] must be the index of an Animation Track.
*/
//go:nosplit
func (self class) AnimationTrackGetKeyAnimation(ctx gd.Lifetime, track_idx gd.Int, key_idx gd.Int) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, key_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_animation_track_get_key_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLength(time_sec gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, time_sec)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_set_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_get_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLoopMode(loop_mode classdb.AnimationLoopMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, loop_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_set_loop_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLoopMode() classdb.AnimationLoopMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationLoopMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_get_loop_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStep(size_sec gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size_sec)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_set_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStep() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_get_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Clear the animation (clear all tracks and reset all).
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a new track to [param to_animation] that is a copy of the given track from this animation.
*/
//go:nosplit
func (self class) CopyTrack(track_idx gd.Int, to_animation gdclass.Animation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, track_idx)
	callframe.Arg(frame, mmm.Get(to_animation[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_copy_track, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Compress the animation and all its tracks in-place. This will make [method track_is_compressed] return [code]true[/code] once called on this [Animation]. Compressed tracks require less memory to be played, and are designed to be used for complex 3D animations (such as cutscenes) imported from external 3D software. Compression is lossy, but the difference is usually not noticeable in real world conditions.
[b]Note:[/b] Compressed tracks have various limitations (such as not being editable from the editor), so only use compressed animations if you actually need them.
*/
//go:nosplit
func (self class) Compress(page_size gd.Int, fps gd.Int, split_tolerance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, page_size)
	callframe.Arg(frame, fps)
	callframe.Arg(frame, split_tolerance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_compress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCaptureIncluded() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Animation.Bind_is_capture_included, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAnimation() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimation() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("Animation", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type TrackType = classdb.AnimationTrackType

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
type InterpolationType = classdb.AnimationInterpolationType

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
type UpdateMode = classdb.AnimationUpdateMode

const (
/*Update between keyframes and hold the value.*/
	UpdateContinuous UpdateMode = 0
/*Update at the keyframes.*/
	UpdateDiscrete UpdateMode = 1
/*Same as [constant UPDATE_CONTINUOUS] but works as a flag to capture the value of the current object and perform interpolation in some methods. See also [method AnimationMixer.capture], [member AnimationPlayer.playback_auto_capture], and [method AnimationPlayer.play_with_capture].*/
	UpdateCapture UpdateMode = 2
)
type LoopMode = classdb.AnimationLoopMode

const (
/*At both ends of the animation, the animation will stop playing.*/
	LoopNone LoopMode = 0
/*At both ends of the animation, the animation will be repeated without changing the playback direction.*/
	LoopLinear LoopMode = 1
/*Repeats playback and reverse playback at both ends of the animation.*/
	LoopPingpong LoopMode = 2
)
type LoopedFlag = classdb.AnimationLoopedFlag

const (
/*This flag indicates that the animation proceeds without any looping.*/
	LoopedFlagNone LoopedFlag = 0
/*This flag indicates that the animation has reached the end of the animation and just after loop processed.*/
	LoopedFlagEnd LoopedFlag = 1
/*This flag indicates that the animation has reached the start of the animation and just after loop processed.*/
	LoopedFlagStart LoopedFlag = 2
)
type FindMode = classdb.AnimationFindMode

const (
/*Finds the nearest time key.*/
	FindModeNearest FindMode = 0
/*Finds only the key with approximating the time.*/
	FindModeApprox FindMode = 1
/*Finds only the key with matching the time.*/
	FindModeExact FindMode = 2
)
