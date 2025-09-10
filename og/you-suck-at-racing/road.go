// WARNING! You will not understand the maths in this file, it makes heavy use of 2d geometry with a catch.
// The catch is, the roads rotate around the car. Yes. You read that right. (Don't judge). Youngwiseone knows!

package main

import (
	"fmt"
	"math"

	"graphics.gd/classdb/CanvasItem"
	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/Texture2D"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector2"
)

// Doubly Linked List of road tiles.
type RoadTile struct {
	Pos      Vector2.XY
	Angle    Angle.Radians
	ID       int
	Texture  Texture2D.Instance
	previous *RoadTile
	next     *RoadTile
}

type Road struct {
	Node2D.Extension[Road]
	car_pos       Vector2.XY
	head          *RoadTile
	resize_offset Vector2.XY
}

const (
	Left  = false
	Right = true
)

func RandomRoadTile(base *RoadTile) *RoadTile {
	size := Resource.Load[Texture2D.Instance]("res://roadtile-1.png")
	sizeVec := size.GetSize()
	angle := base.Angle

	// Adjust angle for slight curves
	if Float.RandomBetween(0, 1) < 0.5 {
		angle += math.Pi / 80
	} else {
		angle -= math.Pi / 80
	}

	angle = max(min(angle, math.Pi/2), -math.Pi/2) // Limit to [-90°, 90°]

	// Place new tile above the base tile (decreasing Y for Y+ down)
	return &RoadTile{
		Pos:     Vector2.Add(base.Pos, Vector2.Rotated(Vector2.New(0, -sizeVec.Y+5), angle)),
		ID:      (base.ID + 1) % 6,
		Texture: Resource.Load[Texture2D.Instance](fmt.Sprintf("res://roadtile-%d.png", (base.ID+1)%6)),
		Angle:   angle,
	}
}

func (road *Road) Ready() {
	road.AsCanvasItem().SetZIndex(-1)
	// Initialize head as a dummy node
	road.head = &RoadTile{
		Angle: math.Pi / 2, // Start with a vertical road
	}
	game_size_int := DisplayServer.WindowGetSize(0)
	game_size := Vector2.New(game_size_int.X, game_size_int.Y)
	size := Resource.Load[Texture2D.Instance]("res://roadtile-1.png")
	sizeVec := size.GetSize()
	current := road.head
	// Create initial straight road tiles
	for y := Float.X(0); y <= Float.X(game_size.Y); y += sizeVec.Y {
		newTile := &RoadTile{
			Pos:      Vector2.New(Float.X(game_size.X/2), y),
			Texture:  Resource.Load[Texture2D.Instance]("res://roadtile-0.png"),
			ID:       0,
			Angle:    0, // Straight up initially
			previous: current,
		}
		current.next = newTile
		current = newTile
		AddToScene(Vector2.Zero, Vector2.Add(current.Pos, Vector2.Rotated(Vector2.New(-sizeVec.X-game_size.X/2+Float.RandomBetween(0, game_size.X/2), 0), current.Angle)), Resource.Load[Texture2D.Instance]("res://tree.png"))
		AddToScene(Vector2.Zero, Vector2.Add(current.Pos, Vector2.Rotated(Vector2.New(+sizeVec.X+game_size.X/2-Float.RandomBetween(0, game_size.X/2), 0), current.Angle)), Resource.Load[Texture2D.Instance]("res://tree.png"))

	}

}

func collisionLineToCircle(from, upto Vector2.XY, center Vector2.XY, radius Float.X) bool {
	var (
		line      = Vector2.Sub(upto, from)
		to_center = Vector2.Sub(center, from)
		line_len  = Vector2.Length(line)
	)
	if line_len == 0 {
		return false
	}
	var (
		line_dir       = Vector2.DivX(line, line_len)
		proj           = Vector2.Dot(to_center, line_dir)
		closest        = Vector2.Add(from, Vector2.MulX(line_dir, proj))
		dist_to_center = Vector2.Length(Vector2.Sub(center, closest))
	)
	return dist_to_center <= radius && proj >= 0 && proj <= line_len
}

func (r *Road) Travel(dt Float.X, pos Vector2.XY, speed Float.X, angle Angle.Radians, resize_offset Vector2.XY) bool {
	r.AsCanvasItem().QueueRedraw()
	r.resize_offset = resize_offset

	if GameOver {
		return false
	}

	speed = Angle.Cos(Angle.Radians(angle)) * speed * dt
	r.car_pos = pos

	game_size_int := DisplayServer.WindowGetSize(0)
	game_size := Vector2.New(game_size_int.X, game_size_int.Y)
	size := Resource.Load[Texture2D.Instance]("res://roadtile-1.png")
	if size == Texture2D.Nil {
		return false
	}
	sizeVec := size.GetSize()
	// Move all tiles downward (increasing Y)
	current := r.head.next
	for current != nil {
		current.Pos.Y += speed // Move down to simulate car moving up
		if current.Pos.Y+resize_offset.Y > Float.X(game_size.Y)+sizeVec.Y/2*10 {
			// Remove tile if it's off-screen
			// Unlink from list
			if current.previous != nil {
				current.previous.next = current.next
			}
			if current.next != nil {
				current.next.previous = current.previous
			}
		}
		current = current.next
	}
	// Find first tile (since we add at the head)
	current = r.head.next
	if current == nil {
		// If list is empty, add a new tile at Y=0
		newTile := &RoadTile{
			Pos:      Vector2.New(Float.X(game_size.X/2), 0),
			Texture:  Resource.Load[Texture2D.Instance]("res://roadtile-0.png"),
			ID:       0,
			Angle:    0,
			previous: r.head,
		}
		r.head.next = newTile
		current = newTile
	}
	// Add new tiles at the top if the first tile's top is too far down
	for current != nil && current.Pos.Y+resize_offset.Y-sizeVec.Y/2 > -sizeVec.Y*50 {
		newTile := RandomRoadTile(current)
		if newTile == nil {
			break
		}
		newTile.next = current
		newTile.previous = r.head
		current.previous = newTile
		r.head.next = newTile
		current = newTile

		AddToScene(resize_offset, Vector2.Add(current.Pos, Vector2.Rotated(Vector2.New(-sizeVec.X-game_size.X/2+Float.RandomBetween(0, game_size.X/2), 0), current.Angle)), Resource.Load[Texture2D.Instance]("res://tree.png"))
		AddToScene(resize_offset, Vector2.Add(current.Pos, Vector2.Rotated(Vector2.New(+sizeVec.X+game_size.X/2-Float.RandomBetween(0, game_size.X/2), 0), current.Angle)), Resource.Load[Texture2D.Instance]("res://tree.png"))
	}
	// Check collision
	current = r.head.next
	for current != nil {
		worldPos := Vector2.Add(Vector2.Add(current.Pos, Vector2.New(-r.car_pos.X, 0)), resize_offset)
		var (
			roadLeftStart  = Vector2.Add(worldPos, Vector2.Rotated(Vector2.New(sizeVec.X/2-10, 0), current.Angle))
			roadLeftEnd    = Vector2.Add(worldPos, Vector2.Rotated(Vector2.New(sizeVec.X/2-10, -sizeVec.Y), current.Angle))
			roadRightStart = Vector2.Add(worldPos, Vector2.Rotated(Vector2.New(-sizeVec.X/2+10, 0), current.Angle))
			roadRightEnd   = Vector2.Add(worldPos, Vector2.Rotated(Vector2.New(-sizeVec.X/2+10, -sizeVec.Y), current.Angle))
		)
		if collisionLineToCircle(roadLeftStart, roadLeftEnd, Vector2.New(game_size.X/2, game_size.Y/2), 32) ||
			collisionLineToCircle(roadRightStart, roadRightEnd, Vector2.New(game_size.X/2, game_size.Y/2), 32) {
			return true
		}
		current = current.next
	}
	return false
}

func (road *Road) Draw() {
	if road.head == nil || road.head.next == nil {
		return
	}
	road.head.next.drawTo(road.AsCanvasItem(), Vector2.Add(Vector2.New(-road.car_pos.X, 0), road.resize_offset))
}

func (tile *RoadTile) drawTo(canvas CanvasItem.Instance, offset Vector2.XY) {
	if tile.Texture == Texture2D.Nil {
		return
	}
	CanvasItem.Expanded(canvas).DrawSetTransform(Vector2.Add(tile.Pos, offset), tile.Angle, Vector2.One)
	canvas.DrawTexture(tile.Texture, Vector2.New(-tile.Texture.GetWidth()/2, -tile.Texture.GetHeight()/2))

	if Debug {
		size := tile.Texture.GetSize()
		canvas.DrawLine(
			Vector2.Add(Vector2.Zero, Vector2.New(size.X/2-10, 0)),
			Vector2.Add(Vector2.Zero, Vector2.New(size.X/2-10, -size.Y)),
			Color.W3C.Red,
		)
		canvas.DrawLine(
			Vector2.Add(Vector2.Zero, Vector2.New(-size.X/2+10, 0)),
			Vector2.Add(Vector2.Zero, Vector2.New(-size.X/2+10, -size.Y)),
			Color.W3C.Red,
		)
	}

	if tile.next != nil {
		tile.next.drawTo(canvas, offset)
	}
}
