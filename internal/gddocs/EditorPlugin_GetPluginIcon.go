/*
[gdscript]
func _get_plugin_icon():
	# You can use a custom icon:
	return preload("res://addons/my_plugin/my_plugin_icon.svg")
	# Or use a built-in icon:
	return EditorInterface.get_editor_theme().get_icon("Node", "EditorIcons")
[/gdscript]
[csharp]
public override Texture2D _GetPluginIcon()
{
	// You can use a custom icon:
	return ResourceLoader.Load<Texture2D>("res://addons/my_plugin/my_plugin_icon.svg");
	// Or use a built-in icon:
	return EditorInterface.Singleton.GetEditorTheme().GetIcon("Node", "EditorIcons");
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/EditorInterface"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/Texture2D"
)

func EditorPlugin_GetPluginIcon() {
	GetPluginIcon := func() Texture2D.Instance {
		// You can use a custom icon:
		if true {
			return Resource.Load[Texture2D.Instance]("res://addons/my_plugin/my_plugin_icon.svg")
		}
		// Or use a built-in icon:
		return EditorInterface.GetEditorTheme().GetIcon("Node", "EditorIcons")
	}
	_ = GetPluginIcon
}
