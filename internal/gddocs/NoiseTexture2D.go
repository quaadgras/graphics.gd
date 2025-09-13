/*
var texture = NoiseTexture2D.new()
texture.noise = FastNoiseLite.new()
await texture.changed
var image = texture.get_image()
var data = image.get_data()
*/

package main

import (
	"graphics.gd/classdb/FastNoiseLite"
	"graphics.gd/classdb/NoiseTexture2D"
)

func ExampleNoiseTexture2D() {
	var texture = NoiseTexture2D.New()
	texture.SetNoise(FastNoiseLite.New().AsNoise())
	texture.AsResource().OnChanged(func() {
		var image = texture.AsTexture2D().GetImage()
		var data = image.GetData()
		_ = data
	})
}
