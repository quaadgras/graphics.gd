/*
[gdscript]
# Stores the node's non-internal groups only (as an array of StringNames).
var non_internal_groups = []
for group in get_groups():
	if not str(group).begins_with("_"):
		non_internal_groups.push_back(group)
[/gdscript]
[csharp]
// Stores the node's non-internal groups only (as a List of StringNames).
List<string> nonInternalGroups = new List<string>();
foreach (string group in GetGroups())
{
	if (!group.BeginsWith("_"))
		nonInternalGroups.Add(group);
}
[/csharp]
*/

package main

import "strings"

func Node_GetGroups() {
	// Stores the node's non-internal groups only (as a slice of strings).
	var nonInternalGroups []string
	for _, group := range node.GetGroups() {
		if !strings.HasPrefix(group, "_") {
			nonInternalGroups = append(nonInternalGroups, group)
		}
	}
}
