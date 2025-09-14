/*
func _get_state():
    var state = {"zoom": zoom, "preferred_color": my_color}
    return state
*/

package main

func EditorPlugin_GetState() {
	GetState := func() map[any]any {
		return map[any]any{"zoom": nil, "preferred_color": nil}
	}
	_ = GetState
}
