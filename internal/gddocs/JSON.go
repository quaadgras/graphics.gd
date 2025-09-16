/*
var data_to_send = ["a", "b", "c"]
var json_string = JSON.stringify(data_to_send)
# Save data
# ...
# Retrieve data
var json = JSON.new()
var error = json.parse(json_string)
if error == OK:
	var data_received = json.data
	if typeof(data_received) == TYPE_ARRAY:
		print(data_received) # Prints the array.
	else:
		print("Unexpected data")
else:
	print("JSON Parse Error: ", json.get_error_message(), " in ", json_string, " at line ", json.get_error_line())
*/

package main

import (
	"graphics.gd/classdb/JSON"
)

func ExampleJSON() {
	var data_to_send = []string{"a", "b", "c"}
	var json_string = JSON.Stringify(data_to_send, "", false)
	// Save data
	// ...
	// Retrieve data
	var json = JSON.New()
	var error = json.Parse(json_string)
	if error == nil {
		var data_received = json.Data()
		if _, ok := data_received.([]any); ok {
			print(data_received) // Prints the array.
		} else {
			print("Unexpected data")
		}
	} else {
		print("JSON Parse Error: ", json.GetErrorMessage(), " in ", json_string, " at line ", json.GetErrorLine())
	}
}
