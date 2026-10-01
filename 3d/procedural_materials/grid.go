package main

// This creates the ImageTexture and assigns it to an existing material at runtime.
// By doing this at runtime, this avoids saving the raw image data in the scene
// file, which would make it much larger.

import (
	"graphics.gd/classdb/BaseMaterial3D"
	"graphics.gd/classdb/Image"
	"graphics.gd/classdb/ImageTexture"
	"graphics.gd/classdb/MeshInstance3D"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
)

const (
	TextureWidth  = 512
	TextureHeight = 512
	GridSize      = 32
	GridThickness = 4
)

type GridMaterial struct {
	MeshInstance3D.Extension[GridMaterial]
}

func (g *GridMaterial) Ready() {
	image := Image.Create(TextureWidth, TextureHeight, false, Image.FormatRgb8)
	// Use 1-dimensional loop as it's faster than a nested loop.
	for i := range TextureWidth * TextureHeight {
		x := i % TextureHeight
		y := i / TextureHeight
		var color Color.RGBA

		// Draw a grid with more contrasted points where X and Y lines meet.
		// Center the grid's lines so that lines are visible on all the texture's edges.
		if (x+GridThickness/2)%GridSize < GridThickness && (y+GridThickness/2)%GridSize < GridThickness {
			color.G = 0.8
		} else if (x+GridThickness/2)%GridSize < GridThickness || (y+GridThickness/2)%GridSize < GridThickness {
			color.G = 0.25
		}

		// Add some random noise for detail.
		color = Color.Add(color, Color.MulX(Color.RGBA{R: Float.Random(), G: Float.Random(), B: Float.Random()}, 0.1))

		image.SetPixel(x, y, color)
	}

	image.GenerateMipmaps()
	imageTexture := ImageTexture.CreateFromImage(image)
	material, _ := Object.As[BaseMaterial3D.Instance](g.AsMeshInstance3D().GetSurfaceOverrideMaterial(0))
	material.SetAlbedoTexture(imageTexture.AsTexture2D())

	image.MoreArgs().BumpMapToNormalMap(5.0)
	image.GenerateMipmaps()
	imageTextureNormal := ImageTexture.CreateFromImage(image)
	material.SetNormalTexture(imageTextureNormal.AsTexture2D())
}
