/*
func _texture_get_data_callback(array):
	value = array.decode_u32(0)

...

rd.texture_get_data_async(texture, 0, _texture_get_data_callback)
*/

package main

import (
	"graphics.gd/classdb/RenderingDevice"
	"graphics.gd/variant/RID"
)

var rd RenderingDevice.Instance
var tex_rid RID.Texture

func RenderingDevice_TextureGetDataAsync() {
	rd.TextureGetDataAsync(tex_rid, 0, func(array []byte) {
		// do something with texture data
	})
}
