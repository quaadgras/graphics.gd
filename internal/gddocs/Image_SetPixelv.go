/*
[gdscript]
var img_width = 10
var img_height = 5
var img = Image.create(img_width, img_height, false, Image.FORMAT_RGBA8)

img.set_pixelv(Vector2i(1, 2), Color.RED) # Sets the color at (1, 2) to red.
[/gdscript]
[csharp]
int imgWidth = 10;
int imgHeight = 5;
var img = Image.Create(imgWidth, imgHeight, false, Image.Format.Rgba8);

img.SetPixelv(new Vector2I(1, 2), Colors.Red); // Sets the color at (1, 2) to red.
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Image"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2i"
)

func Image_SetPixelv() {
	var imgWidth = 10
	var imgHeight = 5
	var img = Image.Create(imgWidth, imgHeight, false, Image.FormatRgba8)
	img.SetPixelv(Vector2i.New(1, 2), Color.W3C.Red) // Sets the color at (1, 2) to red.
}
