/*
extends EditorPlugin

const SETTING_NAME = "addons/my_setting"
const SETTING_DEFAULT = 10.0

func _enter_tree():
	if not ProjectSettings.has_setting(SETTING_NAME):
		ProjectSettings.set_setting(SETTING_NAME, SETTING_DEFAULT)

	ProjectSettings.set_initial_value(SETTING_NAME, SETTING_DEFAULT)
*/

package main

import "graphics.gd/classdb/ProjectSettings"

func ProjectSettings_SetInitialValue() {
	const SETTING_NAME = "addons/my_setting"
	const SETTING_DEFAULT = 10.0

	if !ProjectSettings.HasSetting(SETTING_NAME) {
		ProjectSettings.SetSetting(SETTING_NAME, SETTING_DEFAULT)
	}
	ProjectSettings.SetInitialValue(SETTING_NAME, SETTING_DEFAULT)
}
