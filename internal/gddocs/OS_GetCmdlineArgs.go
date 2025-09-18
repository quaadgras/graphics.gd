/*
[gdscript]
var arguments = {}
for argument in OS.get_cmdline_args():
	if argument.contains("="):
		var key_value = argument.split("=")
		arguments[key_value[0].trim_prefix("--")] = key_value[1]
	else:
		# Options without an argument will be present in the dictionary,
		# with the value set to an empty string.
		arguments[argument.trim_prefix("--")] = ""
[/gdscript]
[csharp]
var arguments = new Dictionary<string, string>();
foreach (var argument in OS.GetCmdlineArgs())
{
	if (argument.Contains('='))
	{
		string[] keyValue = argument.Split("=");
		arguments[keyValue[0].TrimPrefix("--")] = keyValue[1];
	}
	else
	{
		// Options without an argument will be present in the dictionary,
		// with the value set to an empty string.
		arguments[argument.TrimPrefix("--")] = "";
	}
}
[/csharp]
*/

package main

import (
	"strings"

	"graphics.gd/classdb/OS"
)

func OS_GetCmdlineArgs() {
	var arguments = map[string]string{}
	for _, argument := range OS.GetCmdlineArgs() {
		if strings.Contains(argument, "=") {
			_, keyValue, _ := strings.Cut(argument, "=")
			arguments[strings.TrimPrefix(keyValue, "--")] = keyValue
		} else {
			// Options without an argument will be present in the dictionary,
			// with the value set to an empty string.
			arguments[strings.TrimPrefix(argument, "--")] = ""
		}
	}
}
