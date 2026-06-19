/*
func decode_data(string, allow_objects = false):
	return JSON.to_native(JSON.parse_string(string), allow_objects)
*/

package main

import "graphics.gd/classdb/JSON"

func decodeData(jsonString string, allowObjects bool) any {
	return JSON.ToNative(JSON.ParseString(jsonString), allowObjects)
}
