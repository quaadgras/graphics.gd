package main

import (
	"fmt"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/FontFile"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Resource"
	"graphics.gd/startup"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
)

var (
	traditionalFont = Resource.Load[FontFile.Instance]("res://montserrat_semibold.ttf")
	msdfFont        = Resource.Load[FontFile.Instance]("res://montserrat_semibold_msdf.ttf")
)

// SDFFontDemo toggles the animated label between traditional and
// multi-channel signed distance field font rendering.
type SDFFontDemo struct {
	Control.Extension[SDFFontDemo] `gd:"SDFFontDemo"`

	FontLabel Label.Instance `gd:"%FontLabel"`
	FontMode  Label.Instance `gd:"%FontMode"`
	Value     Label.Instance `gd:"%Value"`
}

// usingMSDF reports whether the label's current font uses multi-channel
// signed distance field rendering.
func (s *SDFFontDemo) usingMSDF() bool {
	font, ok := Object.As[FontFile.Instance](s.FontLabel.AsControl().GetThemeFont("font"))
	return ok && font.MultichannelSignedDistanceField()
}

func (s *SDFFontDemo) Input(event InputEvent.Instance) {
	if event.IsActionPressed("toggle_msdf_font") {
		if s.usingMSDF() {
			s.FontLabel.AsControl().AddThemeFontOverride("font", traditionalFont.AsFont())
		} else {
			s.FontLabel.AsControl().AddThemeFontOverride("font", msdfFont.AsFont())
		}
		s.UpdateLabel()
	}
}

func (s *SDFFontDemo) UpdateLabel() {
	mode := "Traditional"
	if s.usingMSDF() {
		mode = "MSDF"
	}
	s.FontMode.SetText(fmt.Sprintf("Font rendering: %s", mode))
}

func (s *SDFFontDemo) OnOutlineSizeValueChanged(value Float.X) {
	s.FontLabel.AsControl().AddThemeConstantOverride("outline_size", int(value))
	s.Value.SetText(fmt.Sprint(value))
}

func main() {
	classdb.Register[SDFFontDemo]()
	startup.Scene()
}
