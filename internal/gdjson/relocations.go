package gdjson

import "strings"

var Relocations = map[string]string{
	"TreeItem.get_tree":                              "Tree.get",
	"BaseButton.get_button_group":                    "ButtonGroup.get",
	"BaseButton.set_button_group":                    "ButtonGroup.set",
	"EditorNode3DGizmo.get_plugin":                   "EditorNode3DGizmoPlugin.get",
	"SkeletonModificationStack2D.add_modification":   "SkeletonModification2D.add",
	"SkeletonModificationStack2D.set_modification":   "SkeletonModification2D.set",
	"SkeletonModificationStack2D.get_modification":   "SkeletonModification2D.get",
	"SkeletonModificationStack2D.get_skeleton":       "Skeleton2D.get",
	"TextEdit.get_syntax_highlighter":                "SyntaxHighlighter.get",
	"TextEdit.set_syntax_highlighter":                "SyntaxHighlighter.set",
	"PackedScene.get_state":                          "SceneState.get",
	"Node.get_viewport":                              "Viewport.get",
	"Node.get_window":                                "Window.get",
	"Node.get_last_exclusive_window":                 "Window.get_last_exclusive",
	"Node.get_tree":                                  "SceneTree.get",
	"Tween.bind_node":                                "Node.bind_tween",
	"CanvasItem.draw_style_box":                      "StyleBox.draw",
	"Viewport.get_embedded_subwindows":               "Window.get_embedded_in_view",
	"OpenXRExtensionWrapperExtension.get_openxr_api": "OpenXRAPIExtension.get_from_wrapper_extension",
	"GLTFNode.get_scene_node_path":                   "GLTFState.get_scene_node_path",
	"GLTFBufferView.load_buffer_view_data":           "GLTFState.load_buffer_view_data",
	"Shape3D.get_debug_mesh":                         "ArrayMesh.get_debug",
	"Resource.get_local_scene":                       "Node.resource_local_scene",
	"StyleBox.draw":                                  "",
	"Tween.tween_property":                           "PropertyTweener.Make",
	"Tween.tween_method":                             "MethodTweener.Make",
}

var RelocationsReverse = map[string]map[string]string{}

func init() {
	for k, v := range Relocations {
		class, method, _ := strings.Cut(v, ".")
		if RelocationsReverse[class] == nil {
			RelocationsReverse[class] = map[string]string{}
		}
		RelocationsReverse[class][method] = k
	}
}
