package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/classdb/GUI"
	"graphics.gd/classdb/Material"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/Sprite2D"
	"graphics.gd/classdb/Texture2D"
	"graphics.gd/startup"
	"graphics.gd/variant/Vector2"
)

var sprite = Sprite2D.New()

func main() {
	classdb.Register[MyFirstShader]()
	classdb.Register[MyFirstShader2D]()
	classdb.Register[MyFirstShader2D_Texture]()
	classdb.Register[MyFirstShader2D_UV]()
	startup.LoadingScene()

	size := DisplayServer.WindowGetSize(0)

	sprite.SetTexture(Resource.Load[Texture2D.Instance]("res://icon.png"))
	sprite.AsNode2D().SetPosition(Vector2.New(size.X/2, size.Y/2))
	sprite.AsNode2D().SetScale(Vector2.New(5, 5))

	var materials = []func() Material.Any{
		func() Material.Any { return new(MyFirstShader2D_Texture) },
		func() Material.Any { return new(MyFirstShader) },
		func() Material.Any { return new(MyFirstShader2D) },
		func() Material.Any { return new(MyFirstShader2D_UV) },
	}
	var index int
	sprite.AsCanvasItem().SetMaterial(materials[index]().AsMaterial())

	button := Button.New()
	button.SetText("Switch Shader")
	button.AsControl().SetAnchorsAndOffsetsPreset(Control.PresetCenterBottom)
	button.SetAlignment(GUI.HorizontalAlignmentCenter)
	buttonSize := button.AsControl().Size()
	button.AsControl().SetPosition(Vector2.New(-buttonSize.X/2, -100))

	button.AsBaseButton().OnPressed(func() {
		index++
		index %= len(materials)
		sprite.AsCanvasItem().SetMaterial(materials[index]().AsMaterial())
	})
	SceneTree.Add(sprite)
	SceneTree.Add(button)
	startup.Scene()
}
