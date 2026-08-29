package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/AnimatedSprite2D"
	"graphics.gd/classdb/Area2D"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMouseMotion"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/PhysicsServer2D"
	"graphics.gd/classdb/World2D"
	"graphics.gd/startup"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/RID"
	"graphics.gd/variant/Transform2D"
	"graphics.gd/variant/Vector2"

	"grow.graphics/eg/2d/bullet_shower/graphics"
)

const (
	BulletCount = 500
	SpeedMin    = 20
	SpeedMax    = 80
)


type Bullet struct {
	Position Vector2.XY
	Speed    Float.X
	Body     RID.Body2D
}

type Bullets struct {
	Node2D.Extension[Bullets] `gd:"BulletShower"`

	bullets [BulletCount]Bullet
	shape   RID.Shape2D
}

func (b *Bullets) Ready() {
	Node2D := b.AsNode2D()
	b.shape = PhysicsServer2D.CircleShapeCreate()
	PhysicsServer2D.ShapeSetData(b.shape, 8)
	for i := range b.bullets {
		var bullet = Bullet{
			Speed: Float.RandomBetween(SpeedMin, SpeedMax),
			Body:  PhysicsServer2D.BodyCreate(),
			Position: Vector2.New(
				Float.RandomBetween(0, Node2D.AsCanvasItem().GetViewportRect().Size.X)+Node2D.AsCanvasItem().GetViewportRect().Size.X,
				Float.RandomBetween(0, Node2D.AsCanvasItem().GetViewportRect().Size.Y),
			),
		}
		PhysicsServer2D.BodySetSpace(bullet.Body, RID.Space2D(World2D.Instance(Node2D.AsCanvasItem().GetWorld2d()).Space()))
		PhysicsServer2D.BodyAddShape(bullet.Body, b.shape, false)
		PhysicsServer2D.BodySetCollisionMask(bullet.Body, 0)
		var transform = Transform2D.New()
		transform.Origin = bullet.Position
		PhysicsServer2D.BodySetState(bullet.Body, PhysicsServer2D.BodyStateTransform, transform)
		b.bullets[i] = bullet
	}
}

func (b *Bullets) Process(delta Float.X) {
	b.AsCanvasItem().QueueRedraw()
}

func (b *Bullets) PhysicsProcess(delta Float.X) {
	var transform = Transform2D.New()
	var offset = b.AsCanvasItem().GetViewportRect().Size.X
	for i := range b.bullets {
		var bullet = &b.bullets[i]
		bullet.Position.X -= bullet.Speed * delta
		if bullet.Position.X < -16 {
			bullet.Position.X = offset // Move the bullet back to the right when it left the screen.
		}
		transform.Origin = bullet.Position
		PhysicsServer2D.BodySetState(bullet.Body, PhysicsServer2D.BodyStateTransform, transform)
	}
}

func (b *Bullets) Draw() {
	var offset = Vector2.MulX(Vector2.Neg(graphics.PNG.Bullet.GetSize()), 0.5)
	for _, bullet := range b.bullets {
		b.AsCanvasItem().DrawTexture(graphics.PNG.Bullet, Vector2.Add(bullet.Position, offset))
	}
}

func (b *Bullets) ExitTree() {
	for _, bullet := range b.bullets {
		PhysicsServer2D.FreeRid(RID.Any(bullet.Body))
	}
	PhysicsServer2D.FreeRid(RID.Any(b.shape))
	clear(b.bullets[:])
}

type Player struct {
	Area2D.Extension[Player] `gd:"BulletShowerPlayer"`

	Sprite AnimatedSprite2D.Instance `gd:"AnimatedSprite2D"`

	touching int
}

func (p *Player) Ready() {
	Input.SetMouseMode(Input.MouseModeHidden)
}

func (p *Player) Input(event InputEvent.Instance) {
	if event, ok := Object.As[InputEventMouseMotion.Instance](event); ok {
		p.AsNode2D().SetPosition(event.AsInputEventMouse().Position())
	}
}

func (p *Player) OnBodyShapeEntered(RID.Body2D, Node2D.Instance, int, int) {
	p.touching++
	if p.touching > 1 {
		p.Sprite.SetFrame(1)
	}
}

func (p *Player) OnBodyShapeExited(RID.Body2D, Node2D.Instance, int, int) {
	p.touching--
	if p.touching == 0 {
		p.Sprite.SetFrame(0)
	}
}

func main() {
	classdb.Register[Player]()
	classdb.Register[Bullets]()
	startup.Scene()
}
