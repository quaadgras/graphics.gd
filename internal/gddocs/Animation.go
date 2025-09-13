/*
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
*/

package main

import "graphics.gd/classdb/Animation"

func ExampleAnimation() {
	// This creates an animation that makes the node "Enemy" move to the right by
	// 100 pixels in 2.0 seconds.
	var animation = Animation.New()
	var track_index = animation.AddTrack(Animation.TypeValue)
	animation.TrackSetPath(track_index, "Enemy:position:x")
	animation.TrackInsertKey(track_index, 0.0, 0)
	animation.TrackInsertKey(track_index, 2.0, 100)
	animation.SetLength(2.0)
}
