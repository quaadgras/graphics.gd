/*
func encode_data(value, full_objects = false):
	return JSON.stringify(JSON.from_native(value, full_objects))
*/

package main

import "graphics.gd/classdb/JSON"

func encodeData(value any, fullObjects bool) string {
	return JSON.Stringify(JSON.FromNative(value, fullObjects), "", false)
}
