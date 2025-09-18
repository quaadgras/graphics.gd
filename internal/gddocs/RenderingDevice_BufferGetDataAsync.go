/*
func _buffer_get_data_callback(array):
	value = array.decode_u32(0)

...

rd.buffer_get_data_async(buffer, _buffer_get_data_callback)
*/

package main

import "graphics.gd/variant/RID"

var buf_rid RID.Buffer

func RenderingDevice_BufferGetDataAsync() {
	rd.BufferGetDataAsync(buf_rid, func(array []byte) {
		// do something with buffer data
	})
}
