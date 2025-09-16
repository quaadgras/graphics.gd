---
title: Declarative Children
slug: guide/classdb/children
---

If a Go structs extends a [Node](https://pkg.go.dev/graphics.gd/classdb/Node)-derived class, any exported
[Node](https://pkg.go.dev/graphics.gd/classdb/Node)-derived fields will be populated automatically.
This means `graphics.gd` will look for any children that match the same name as the field (or the name within the `gd`
struct tag) and either populate the field with the first child found or create a new instance of the field's type if
no child is found.

This is useful for ensuring that any nodes that your struct depends on to be located in the scene tree are present.
The `gd` struct tag can either be a name, a path, or a uniquely-named node, ie `%MyNode`.

```go
package main

import (
	"fmt"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Label"
	"graphics.gd/startup"
)

type MyGUI struct {
	Control.Extension[MyNode]

	Label struct {
		Label.Instance `gd:"Label"` // will point to the "Label" child

		ClickMe Button.Instance `gd:"ClickMe"` // will point to a "ClickMe" child
	}
}

func (m *MyGUI) Ready() {
	m.Label.SetText("My GUI")
	m.Label.ClickMe.AsBaseButton().OnPressed(func() {
		fmt.Println("Button clicked!")
	})
}

func main() {
	classdb.Register[MyGUI]()
	startup.LoadingScene()
	SceneTree.Add(new(MyGUI))
	startup.Scene()
}
```
