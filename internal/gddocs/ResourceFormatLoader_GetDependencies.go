/*
func _get_dependencies(path, add_types):
	return [
		"uid://fqgvuwrkuixh::Script::res://script.gd",
		"uid://fqgvuwrkuixh::::res://script.gd",
		"res://script.gd::Script",
		"res://script.gd",
	]
*/

package main

func ResourceFormatLoader_GetDependencies() {
	GetDependencies := func(path string, addTypes bool) []string {
		return []string{
			"uid://fqgvuwrkuixh::Script::res://script.gd",
			"uid://fqgvuwrkuixh::::res://script.gd",
			"res://script.gd::Script",
			"res://script.gd",
		}
	}
	_ = GetDependencies
}
