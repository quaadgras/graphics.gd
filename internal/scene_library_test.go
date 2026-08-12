//go:build !generate

package gd_test

import (
	"testing"

	"graphics.gd/classdb/Area2D"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/Texture2D"
	"graphics.gd/internal/gdreference"
)

// Scene-struct binding and resource library tests, see
// https://github.com/quaadgras/graphics.gd/discussions/283

type SceneMain struct {
	Node2D.Instance

	Bullets Node2D.Instance
	Player  struct {
		Area2D.Instance

		Sprite Node2D.Instance `gd:"Sprite"`
	}
	Missing   Node2D.Instance
	WrongType Area2D.Instance `gd:"Bullets"`
}

// testLibrary is deliberately declared at package level: Resource.Library
// must be safe to call before startup, deferring all engine access until a
// field is first used.
var testLibrary = Resource.Library[struct {
	Icon  Texture2D.Instance   `gd:"res://icon.svg"`
	Scene PackedScene.Instance `gd:"res://main.tscn"`
}]()

func makeSceneMain(t *testing.T) PackedScene.Is[SceneMain] {
	t.Helper()
	root := Node2D.New()
	root.AsNode().SetName("Main")
	bullets := Node2D.New()
	bullets.AsNode().SetName("Bullets")
	root.AsNode().AddChild(bullets.AsNode())
	bullets.AsNode().SetOwner(root.AsNode())
	player := Area2D.New()
	player.AsNode().SetName("Player")
	root.AsNode().AddChild(player.AsNode())
	player.AsNode().SetOwner(root.AsNode())
	sprite := Node2D.New()
	sprite.AsNode().SetName("Sprite")
	player.AsNode().AddChild(sprite.AsNode())
	sprite.AsNode().SetOwner(root.AsNode())
	packed := PackedScene.New()
	if err := packed.Pack(root.AsNode()); err != nil {
		t.Fatal(err)
	}
	return PackedScene.Is[SceneMain](packed)
}

func TestSceneStructBinding(t *testing.T) {
	scene := makeSceneMain(t)
	main := scene.Instantiate()
	if name := main.AsNode().Name(); name != "Main" {
		t.Fatalf("expected root to be bound, got name %q", name)
	}
	if name := main.Bullets.AsNode().Name(); name != "Bullets" {
		t.Fatalf("expected Bullets to resolve, got name %q", name)
	}
	if name := main.Player.AsNode().Name(); name != "Player" {
		t.Fatalf("expected Player to resolve, got name %q", name)
	}
	if name := main.Player.Sprite.AsNode().Name(); name != "Sprite" {
		t.Fatalf("expected Player/Sprite to resolve, got name %q", name)
	}
	if !gdreference.BadObject(main.Missing.AsObject()[0]) {
		t.Fatal("expected Missing to read as null, no such node exists")
	}
	if !gdreference.BadObject(main.WrongType.AsObject()[0]) {
		t.Fatal("expected WrongType to read as null, Bullets is not an Area2D")
	}
	main.AsNode().QueueFree()
}

func TestResourceLibrary(t *testing.T) {
	if testLibrary.Icon.GetWidth() == 0 {
		t.Fatal("expected res://icon.svg to load with a non-zero width")
	}
	if !testLibrary.Scene.CanInstantiate() {
		t.Fatal("expected res://main.tscn to load as an instantiable scene")
	}
}
