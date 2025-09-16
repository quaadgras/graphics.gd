/*
[gdscript]
# autoload.gd
func _enter_tree():
	# Sets our custom multiplayer as the main one in SceneTree.
	get_tree().set_multiplayer(LogMultiplayer.new())
[/gdscript]
*/

package main

import (
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/variant/Object"
)

func ExampleSetCustomMultiplayer() {
	Object.To[SceneTree.Instance](Engine.GetMainLoop()).SetMultiplayer(new(LogMultiplayer).AsMultiplayerAPI())
}
