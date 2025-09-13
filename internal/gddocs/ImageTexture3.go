/*
var texture = load("res://icon.svg")
var image = texture.get_image()
*/

package main

import (
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/Sprite2D"
	"graphics.gd/classdb/Texture2D"
)

func ExampleLoadImage(sprite Sprite2D.Instance) {
	var texture = Resource.Load[Texture2D.Instance]("res://icon.svg")
	var image = texture.GetImage()
	_ = image
}
