/*
class MyExportPlugin extends EditorExportPlugin:
	func _get_name() -> String:
		return "MyExportPlugin"

	func _supports_platform(platform) -> bool:
		if platform is EditorExportPlatformPC:
			# Run on all desktop platforms including Windows, MacOS and Linux.
			return true
		return false

	func _get_export_options_overrides(platform) -> Dictionary:
		# Override "Embed PCK" to always be enabled.
		return {
			"binary_format/embed_pck": true,
		}
*/

package main

import (
	"graphics.gd/classdb/EditorExportPlatform"
	"graphics.gd/classdb/EditorExportPlatformPC"
	"graphics.gd/classdb/EditorExportPlugin"
	"graphics.gd/variant/Object"
)

type MyExportPlugin struct {
	EditorExportPlugin.Extension[MyExportPlugin]
}

func (m *MyExportPlugin) GetName() string { return "MyExportPlugin" }

func (m *MyExportPlugin) SupportsPlatform(platform EditorExportPlatform.Instance) bool {
	return Object.Is[EditorExportPlatformPC.Instance](platform)
}

func (m *MyExportPlugin) GetExportOptionsOverrides(platform EditorExportPlatform.Instance) map[string]any {
	return map[string]any{
		"binary_format/embed_pck": true,
	}
}
