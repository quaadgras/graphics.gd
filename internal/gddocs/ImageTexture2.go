/*
var texture = load("res://icon.svg")
$Sprite2D.texture = texture
*/

package main

import (
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/Sprite2D"
	"graphics.gd/classdb/Texture2D"
)

func ExampleLoadImageTexture(sprite Sprite2D.Instance) {
	var texture = Resource.Load[Texture2D.Instance]("res://icon.svg")
	sprite.SetTexture(texture.AsTexture2D())
}
