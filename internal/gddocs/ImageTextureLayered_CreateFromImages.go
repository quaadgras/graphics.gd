/*
# Fill in an array of Images with different colors.
var images = []
const LAYERS = 6
for i in LAYERS:
    var image = Image.create_empty(128, 128, false, Image.FORMAT_RGB8)
    if i % 3 == 0:
        image.fill(Color.RED)
    elif i % 3 == 1:
        image.fill(Color.GREEN)
    else:
        image.fill(Color.BLUE)
    images.push_back(image)

# Create and save a 2D texture array. The array of images must have at least 1 Image.
var texture_2d_array = Texture2DArray.new()
texture_2d_array.create_from_images(images)
ResourceSaver.save(texture_2d_array, "res://texture_2d_array.res", ResourceSaver.FLAG_COMPRESS)

# Create and save a cubemap. The array of images must have exactly 6 Images.
# The cubemap's images are specified in this order: X+, X-, Y+, Y-, Z+, Z-
# (in Godot's coordinate system, so Y+ is "up" and Z- is "forward").
var cubemap = Cubemap.new()
cubemap.create_from_images(images)
ResourceSaver.save(cubemap, "res://cubemap.res", ResourceSaver.FLAG_COMPRESS)

# Create and save a cubemap array. The array of images must have a multiple of 6 Images.
# Each cubemap's images are specified in this order: X+, X-, Y+, Y-, Z+, Z-
# (in Godot's coordinate system, so Y+ is "up" and Z- is "forward").
var cubemap_array = CubemapArray.new()
cubemap_array.create_from_images(images)
ResourceSaver.save(cubemap_array, "res://cubemap_array.res", ResourceSaver.FLAG_COMPRESS)
*/

package main

import (
	"graphics.gd/classdb/Cubemap"
	"graphics.gd/classdb/Image"
	"graphics.gd/classdb/ImageTextureLayered"
	"graphics.gd/classdb/ResourceSaver"
	"graphics.gd/variant/Color"
)

func ImageTextureLayered_CreateFromImages() {
	// Fill in an array of Images with different colors.
	var images []Image.Instance
	const LAYERS = 6
	for i := range LAYERS {
		var image = Image.CreateEmpty(128, 128, false, Image.FormatRgb8)
		if i%3 == 0 {
			image.Fill(Color.W3C.Red)
		} else if i%3 == 1 {
			image.Fill(Color.W3C.Green)
		} else {
			image.Fill(Color.W3C.Blue)
		}
		images = append(images, image)
	}

	// Create and save a 2D texture array. The array of images must have at least 1 Image.
	var texture2DArray = ImageTextureLayered.New()
	texture2DArray.CreateFromImages(images)
	ResourceSaver.Save(texture2DArray.AsResource(), "res://texture_2d_array.res", ResourceSaver.FlagCompress)

	// Create and save a cubemap. The array of images must have exactly 6 Images.
	// The cubemap's images are specified in this order: X+, X-, Y+, Y-, Z+, Z-
	// (in Godot's coordinate system, so Y+ is "up" and Z- is "forward").
	var cubemap = Cubemap.New()
	cubemap.AsImageTextureLayered().CreateFromImages(images)
	ResourceSaver.Save(cubemap.AsResource(), "res://cubemap.res", ResourceSaver.FlagCompress)

	// Create and save a cubemap array. The array of images must have a multiple of 6 Images.
	// Each cubemap's images are specified in this order: X+, X-, Y+, Y-, Z+, Z-
	// (in Godot's coordinate system, so Y+ is "up" and Z- is "forward").
	var cubemapArray = Cubemap.New()
	cubemapArray.AsImageTextureLayered().CreateFromImages(images)
	ResourceSaver.Save(cubemapArray.AsResource(), "res://cubemap_array.res", ResourceSaver.FlagCompress)
}
