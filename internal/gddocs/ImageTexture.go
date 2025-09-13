/*
var image = Image.load_from_file("res://icon.svg")
var texture = ImageTexture.create_from_image(image)
$Sprite2D.texture = texture
*/

package main

import (
	"graphics.gd/classdb/Image"
	"graphics.gd/classdb/ImageTexture"
	"graphics.gd/classdb/Sprite2D"
)

func ExampleImageTexture(sprite Sprite2D.Instance) {
	var image = Image.LoadFromFile("res://icon.svg")
	var texture = ImageTexture.CreateFromImage(image)
	sprite.SetTexture(texture.AsTexture2D())
}
