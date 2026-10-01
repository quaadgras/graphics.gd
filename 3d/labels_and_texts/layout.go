package main

import (
	"fmt"
	"strings"

	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Label3D"
	"graphics.gd/classdb/LineEdit"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/TextServer"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
)

const (
	// The margin to apply between the name and health percentage (in pixels).
	HealthMargin = 25

	// The health bar width (in number of characters).
	// Higher can be more precise, at the cost of lower performance
	// (since more characters may need to be rendered at once).
	BarWidth = 100
)

// Label3DLayout simulates layout by adjusting Label3Ds' `offset` properties.
// This can be done in the inspector, or via code with the help of `Font.get_string_size()`.
// For proper billboard behavior, Label3D's `offset` property
// must be adjusted instead of adjusting the `position` property.
type Label3DLayout struct {
	Node3D.Extension[Label3DLayout]

	Name                Label3D.Instance
	Name2               Label.Instance
	Health              Label3D.Instance
	HealthBarForeground Label3D.Instance
	HealthBarBackground Label3D.Instance
	LineEdit            LineEdit.Instance

	health  int
	counter Float.X
}

func (l *Label3DLayout) Ready() {
	l.LineEdit.SetText(l.Name.Text())
}

func (l *Label3DLayout) Process(delta Float.X) {
	// Animate the health percentage.
	l.counter += delta
	l.SetHealth(int(Float.Round(50 + Angle.Sin(Angle.Radians(l.counter*0.5))*50)))
}

func (l *Label3DLayout) OnLineEditTextChanged(newText string) {
	l.Name.SetText(newText)

	// Adjust name's font size to fit within the allowed width.
	l.Name.SetFontSize(32)
	for l.Name.Font().MoreArgs().GetStringSize(l.Name.Text(), l.Name.HorizontalAlignment(), -1, l.Name.FontSize(), TextServer.JustificationKashida|TextServer.JustificationWordBound, TextServer.DirectionAuto, TextServer.OrientationHorizontal).X > l.Name.Width() {
		l.Name.SetFontSize(l.Name.FontSize() - 1)
	}
}

func (l *Label3DLayout) SetHealth(health int) {
	l.health = health

	l.Health.SetText(fmt.Sprintf("%d%%", health))
	if health <= 30 {
		// Low health alert.
		l.Health.SetModulate(Color.RGBA{1, 0.2, 0.1, 1})
		l.Health.SetOutlineModulate(Color.RGBA{0.2, 0.1, 0.0, 1})
		l.HealthBarForeground.SetModulate(Color.RGBA{1, 0.2, 0.1, 1})
		l.HealthBarForeground.SetOutlineModulate(Color.RGBA{0.2, 0.1, 0.0, 1})
		l.HealthBarBackground.SetOutlineModulate(Color.RGBA{0.2, 0.1, 0.0, 1})
		l.HealthBarBackground.SetModulate(Color.RGBA{0.2, 0.1, 0.0, 1})
	} else {
		l.Health.SetModulate(Color.RGBA{0.8, 1, 0.4, 1})
		l.Health.SetOutlineModulate(Color.RGBA{0.15, 0.2, 0.15, 1})
		l.HealthBarForeground.SetModulate(Color.RGBA{0.8, 1, 0.4, 1})
		l.HealthBarForeground.SetOutlineModulate(Color.RGBA{0.15, 0.2, 0.15, 1})
		l.HealthBarBackground.SetOutlineModulate(Color.RGBA{0.15, 0.2, 0.15, 1})
		l.HealthBarBackground.SetModulate(Color.RGBA{0.15, 0.2, 0.15, 1})
	}

	// Construct a health bar with `|` symbols brought very close to each other using
	// a custom FontVariation on the HealthBarForeground and HealthBarBackground nodes.
	l.HealthBarForeground.SetText(strings.Repeat("|", int(Float.Round(Float.X(health)/100.0*BarWidth))))
	l.HealthBarBackground.SetText(strings.Repeat("|", BarWidth))
}
