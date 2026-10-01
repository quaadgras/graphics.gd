package main

import (
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/Sprite2D"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Vector2"
)

const (
	PadSpeed         = 150
	InitialBallSpeed = 80.0
)

// Pong is a minimal pong game rendered inside the SubViewport.
type Pong struct {
	Node2D.Extension[Pong] `gd:"Pong"`

	Ball        Sprite2D.Instance
	LeftPaddle  Sprite2D.Instance
	RightPaddle Sprite2D.Instance

	ballSpeed  Float.X
	screenSize Vector2.XY
	direction  Vector2.XY // Default ball direction.
	padSize    Vector2.XY
}

func NewPong() *Pong {
	return &Pong{
		ballSpeed:  InitialBallSpeed,
		screenSize: Vector2.New(640, 400),
		direction:  Vector2.Left,
		padSize:    Vector2.New(8, 32),
	}
}

func (p *Pong) Ready() {
	p.screenSize = p.AsCanvasItem().GetViewportRect().Size // Get actual size.
	p.padSize = p.LeftPaddle.Texture().GetSize()
}

func (p *Pong) Process(delta Float.X) {
	// Get ball position and pad rectangles.
	ballPos := p.Ball.AsNode2D().Position()
	leftRect := Rect2.PositionSize{Position: Vector2.Sub(p.LeftPaddle.AsNode2D().Position(), Vector2.MulX(p.padSize, 0.5)), Size: p.padSize}
	rightRect := Rect2.PositionSize{Position: Vector2.Sub(p.RightPaddle.AsNode2D().Position(), Vector2.MulX(p.padSize, 0.5)), Size: p.padSize}

	// Integrate new ball position.
	ballPos = Vector2.Add(ballPos, Vector2.MulX(p.direction, p.ballSpeed*delta))

	// Flip when touching roof or floor.
	if (ballPos.Y < 0 && p.direction.Y < 0) || (ballPos.Y > p.screenSize.Y && p.direction.Y > 0) {
		p.direction.Y = -p.direction.Y
	}

	// Flip, change direction, and increase speed when touching pads.
	if (Rect2.HasPoint(leftRect, ballPos) && p.direction.X < 0) || (Rect2.HasPoint(rightRect, ballPos) && p.direction.X > 0) {
		p.direction.X = -p.direction.X
		p.ballSpeed *= 1.1
		p.direction.Y = Float.Random()*2.0 - 1
		p.direction = Vector2.Normalized(p.direction)
	}

	// Check gameover.
	if ballPos.X < 0 || ballPos.X > p.screenSize.X {
		ballPos = Vector2.MulX(p.screenSize, 0.5)
		p.ballSpeed = InitialBallSpeed
		p.direction = Vector2.New(-1, 0)
	}

	p.Ball.AsNode2D().SetPosition(ballPos)

	// Move left pad.
	leftPos := p.LeftPaddle.AsNode2D().Position()
	if leftPos.Y > 0 && Input.IsActionPressed("left_move_up", false) {
		leftPos.Y += -PadSpeed * delta
	}
	if leftPos.Y < p.screenSize.Y && Input.IsActionPressed("left_move_down", false) {
		leftPos.Y += PadSpeed * delta
	}
	p.LeftPaddle.AsNode2D().SetPosition(leftPos)

	// Move right pad.
	rightPos := p.RightPaddle.AsNode2D().Position()
	if rightPos.Y > 0 && Input.IsActionPressed("right_move_up", false) {
		rightPos.Y += -PadSpeed * delta
	}
	if rightPos.Y < p.screenSize.Y && Input.IsActionPressed("right_move_down", false) {
		rightPos.Y += PadSpeed * delta
	}
	p.RightPaddle.AsNode2D().SetPosition(rightPos)
}
