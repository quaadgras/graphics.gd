package nativeui

import (
	"sync"

	"graphics.gd/classdb"
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/classdb/StyleBox"
	"graphics.gd/classdb/Texture2D"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/RID"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Vector2"
)

// anchor says where a layer goes in the stylebox's rectangle.
type anchor int

const (
	fillRect  anchor = iota // 9-sliced over the whole rectangle (minus the other layers' reserved width)
	leftEdge                // unstretched, at the left edge, centered vertically
	rightEdge               // 9-sliced, full height, fixed width at the right edge
)

// layer is one image of a layeredBox.
type layer struct {
	tex    Texture2D.Instance
	anchor anchor
	width  Float.X    // rightEdge: the width it takes; fillRect: width left out on the right
	margin [4]Float.X // 9-slice margins: left, top, right, bottom
}

// layeredBox is a StyleBox drawn from several images, for widgets that a
// single 9-slice can't draw: GTK 2's entry-shaped combo boxes (a field plus a
// button) and check buttons whose indicator sits left of the label.
type layeredBox struct {
	StyleBox.Extension[layeredBox] `gd:"NativeUILayeredBox"`

	layers []layer
}

var registerLayeredBox sync.Once

func newLayeredBox(layers ...layer) StyleBox.Instance {
	registerLayeredBox.Do(func() { classdb.Register[layeredBox]() })
	box := &layeredBox{layers: layers}
	return box.AsStyleBox()
}

func (b *layeredBox) Draw(item RID.CanvasItem, rect Rect2.PositionSize) {
	for _, l := range b.layers {
		size := l.tex.GetSize()
		texture := RID.CanvasTexture(l.tex.AsResource().GetRid())
		switch l.anchor {
		case leftEdge:
			at := Rect2.PositionSize{
				Position: Vector2.XY{X: rect.Position.X, Y: rect.Position.Y + (rect.Size.Y-size.Y)/2},
				Size:     size,
			}
			RenderingServer.CanvasItemAddTextureRect(item, at, texture, false, false)
		case rightEdge:
			at := Rect2.PositionSize{
				Position: Vector2.XY{X: rect.Position.X + rect.Size.X - l.width, Y: rect.Position.Y},
				Size:     Vector2.XY{X: l.width, Y: rect.Size.Y},
			}
			b.ninePatch(item, at, texture, size, l.margin)
		default:
			at := rect
			at.Size.X -= l.width
			b.ninePatch(item, at, texture, size, l.margin)
		}
	}
}

func (b *layeredBox) ninePatch(item RID.CanvasItem, rect Rect2.PositionSize, texture RID.CanvasTexture, size Vector2.XY, m [4]Float.X) {
	RenderingServer.CanvasItemAddNinePatch(item, rect, Rect2.PositionSize{Size: size}, texture,
		Vector2.XY{X: m[0], Y: m[1]}, Vector2.XY{X: m[2], Y: m[3]},
		RenderingServer.NinePatchStretch, RenderingServer.NinePatchStretch)
}
