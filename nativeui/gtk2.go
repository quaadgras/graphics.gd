//go:build (linux && !android) || freebsd || openbsd || netbsd || dragonfly || nativeui_all

package nativeui

import (
	"image"
	"math"
	"path/filepath"
	"strings"

	"graphics.gd/classdb/StyleBox"
	"graphics.gd/classdb/StyleBoxEmpty"
	"graphics.gd/classdb/StyleBoxFlat"
	"graphics.gd/classdb/StyleBoxTexture"
	"graphics.gd/classdb/SystemFont"
	"graphics.gd/classdb/Theme"
	"graphics.gd/nativeui/internal/desktop"
	"graphics.gd/nativeui/internal/gtkrc"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
)

type gtk2Builder struct {
	rc    *gtkrc.Theme
	theme Theme.Instance
	k     float64
}

func rc2(c gtkrc.RGBA) Color.RGBA {
	return Color.RGBA{R: Float.X(c.R), G: Float.X(c.G), B: Float.X(c.B), A: Float.X(c.A)}
}

// widgets live in a window's box, as in a real GTK 2 window.
func (g *gtk2Builder) style(classPath string) *gtkrc.Style {
	return g.rc.Resolve(gtkrc.W("GtkWindow.GtkVBox." + classPath))
}

func (g *gtk2Builder) scaled(img *image.NRGBA) *image.NRGBA {
	if g.k == 1 {
		return img
	}
	return scaleNearest(img, g.k)
}

// box draws calls into a w×h image and makes a 9-slice StyleBox from it, with
// fixed margins of the style's thickness (or the pixmap's border).
func (g *gtk2Builder) box(st *gtkrc.Style, w, h int, content [4]float64, calls ...gtkrc.Call) StyleBox.Instance {
	img := image.NewNRGBA(image.Rect(0, 0, w, h))
	margin := [4]int{st.YThickness, st.XThickness, st.YThickness, st.XThickness} // top, right, bottom, left
	for _, c := range calls {
		st.DrawInto(img, img.Bounds(), c)
		if pm, ok := st.Match(c); ok && pm.File != "" {
			b := pm.Border // left, right, top, bottom
			margin = [4]int{max(margin[0], b[2]), max(margin[1], b[1]), max(margin[2], b[3]), max(margin[3], b[0])}
		}
	}
	sb := StyleBoxTexture.New()
	sb.SetTexture(textureFrom(g.scaled(img)))
	k := g.k
	sb.SetTextureMarginTop(Float.X(float64(margin[0]) * k))
	sb.SetTextureMarginRight(Float.X(float64(margin[1]) * k))
	sb.SetTextureMarginBottom(Float.X(float64(margin[2]) * k))
	sb.SetTextureMarginLeft(Float.X(float64(margin[3]) * k))
	sb.AsStyleBox().SetContentMarginTop(Float.X(content[0] * k))
	sb.AsStyleBox().SetContentMarginRight(Float.X(content[1] * k))
	sb.AsStyleBox().SetContentMarginBottom(Float.X(content[2] * k))
	sb.AsStyleBox().SetContentMarginLeft(Float.X(content[3] * k))
	return sb.AsStyleBox()
}

func (g *gtk2Builder) icon(st *gtkrc.Style, w, h int, c gtkrc.Call) image.Image {
	return g.scaled(st.Draw(c, w, h))
}

func pad(x, y float64) [4]float64 { return [4]float64{y, x, y, x} }

func (g *gtk2Builder) font(family string, weight int) SystemFont.Instance {
	f := SystemFont.New()
	f.SetFontNames(fontNames(family, "sans-serif"))
	f.SetFontWeight(weight)
	return f
}

func buildGTK2(opts options) (Theme.Instance, error) {
	rc, err := gtkrc.Load(filepath.Join(opts.GTKTheme, "gtk-2.0", "gtkrc"))
	if err != nil {
		return Theme.Nil, err
	}
	g := &gtk2Builder{rc: rc, theme: Theme.New(), k: float64(opts.Scale)}
	t := g.theme
	empty := StyleBoxEmpty.New().AsStyleBox()
	win := g.rc.Resolve(gtkrc.W("GtkWindow"))
	label := g.style("GtkLabel")

	// Font: the desktop's, or the theme's font_name.
	family, size, weight := opts.Font.Family, float64(opts.Font.Size), opts.Font.Weight
	if label.Font != "" && family == "" {
		f := desktop.ParseFont(label.Font)
		family, size, weight = f.Family, f.Size*96/72, f.Weight
	}
	if family == "" {
		family, size, weight = "Sans", 10*96.0/72, 400
	}
	t.SetDefaultBaseScale(Float.X(g.k))
	t.SetDefaultFont(g.font(family, weight).AsFont())
	t.SetDefaultFontSize(int(math.Round(size * g.k)))
	bold := g.font(family, 700).AsFont()

	bg := StyleBoxFlat.New()
	bg.SetBgColor(rc2(win.Bg[gtkrc.Normal]))
	t.SetStylebox("panel", "Panel", bg.AsStyleBox())
	t.SetStylebox("panel", "PanelContainer", bg.AsStyleBox())
	t.SetColor("font_color", "Label", rc2(label.Fg[gtkrc.Normal]))
	t.SetTypeVariation(variationHeaderLabel, "Label")
	t.SetFont("font", variationHeaderLabel, bold)
	t.SetTypeVariation(variationDimLabel, "Label")
	t.SetColor("font_color", variationDimLabel, rc2(label.Fg[gtkrc.Insensitive]))
	frame := g.style("GtkFrame")
	t.SetTypeVariation(variationCard, "PanelContainer")
	t.SetStylebox("panel", variationCard, g.box(frame, 64, 64, pad(8, 8), gtkrc.Call{Function: "SHADOW", Shadow: "ETCHED_IN", Detail: "frame"}))

	// Buttons: GTK 2 pads the label by the thickness, inner border, focus padding and focus line.
	button := g.style("GtkHBox.GtkButton")
	bpad := pad(float64(button.XThickness+4), float64(button.YThickness+2))
	states := []struct {
		godot  string
		state  int
		shadow string
	}{{"normal", gtkrc.Normal, "OUT"}, {"hover", gtkrc.Prelight, "OUT"}, {"pressed", gtkrc.Active, "IN"}, {"hover_pressed", gtkrc.Active, "IN"}, {"disabled", gtkrc.Insensitive, "OUT"}}
	buttonLike := func(typ string, st *gtkrc.Style, w int) {
		for _, s := range states {
			t.SetStylebox(s.godot, typ, g.box(st, w, 23, bpad, gtkrc.Call{Function: "BOX", Detail: "button", State: s.state, Shadow: s.shadow}))
			t.SetColor(godotFontColors[s.godot], typ, rc2(st.Fg[s.state]))
		}
		t.SetColor("font_focus_color", typ, rc2(st.Fg[gtkrc.Normal]))
		focus := image.NewNRGBA(image.Rect(0, 0, w, 23))
		in := st.XThickness + 1
		st.DrawInto(focus, image.Rect(in, in, w-in, 23-in), gtkrc.Call{Function: "FOCUS", Detail: "button"})
		fb := StyleBoxTexture.New()
		fb.SetTexture(textureFrom(g.scaled(focus)))
		fb.SetTextureMarginAll(Float.X(float64(in+1) * g.k))
		fb.SetAxisStretchHorizontal(StyleBoxTexture.AxisStretchModeTileFit)
		fb.SetAxisStretchVertical(StyleBoxTexture.AxisStretchModeTileFit)
		t.SetStylebox("focus", typ, fb.AsStyleBox())
	}
	buttonLike("Button", button, 75)
	t.SetTypeVariation(variationAccentButton, "Button")
	combo := g.style("GtkHBox.GtkComboBox.GtkToggleButton")
	buttonLike("OptionButton", combo, 100)
	arrow := asNRGBA(g.icon(combo, 9, 9, gtkrc.Call{Function: "ARROW", Detail: "arrow", ArrowDirection: "DOWN", Shadow: "OUT", Widget: "GtkComboBox"}))
	t.SetIcon("arrow", "OptionButton", textureFrom(g.scaled(arrow)))
	t.SetConstant("arrow_margin", "OptionButton", int(4*g.k))
	if v, _ := g.style("GtkHBox.GtkComboBox").Prop("GtkComboBox", "appears-as-list"); v == "1" || v == "TRUE" {
		g.listCombo(combo, arrow.Bounds().Dx())
	}

	// Check boxes and radio buttons (GTK 2 has no switch: CheckButtons are check boxes).
	check, radio := g.style("GtkCheckButton"), g.style("GtkRadioButton")
	isize := check.PropInt("GtkCheckButton", "indicator-size", 13)
	defer g.leftIndicators(check, isize)
	for _, typ := range []string{"CheckBox", "CheckButton"} {
		for _, name := range []string{"normal", "hover", "pressed", "hover_pressed", "disabled", "focus"} {
			t.SetStylebox(name, typ, empty)
		}
		for _, s := range states {
			t.SetColor(godotFontColors[s.godot], typ, rc2(check.Fg[s.state]))
		}
		t.SetColor("font_focus_color", typ, rc2(check.Fg[gtkrc.Normal]))
		t.SetConstant("h_separation", typ, int(float64(check.PropInt("GtkCheckButton", "indicator-spacing", 2)*2+2)*g.k))
	}
	for _, v := range []struct {
		name   string
		st     *gtkrc.Style
		fn     string
		detail string
		state  int
		shadow string
	}{
		{"unchecked", check, "CHECK", "checkbutton", gtkrc.Normal, "OUT"},
		{"checked", check, "CHECK", "checkbutton", gtkrc.Normal, "IN"},
		{"unchecked_disabled", check, "CHECK", "checkbutton", gtkrc.Insensitive, "OUT"},
		{"checked_disabled", check, "CHECK", "checkbutton", gtkrc.Insensitive, "IN"},
		{"radio_unchecked", radio, "OPTION", "radiobutton", gtkrc.Normal, "OUT"},
		{"radio_checked", radio, "OPTION", "radiobutton", gtkrc.Normal, "IN"},
		{"radio_unchecked_disabled", radio, "OPTION", "radiobutton", gtkrc.Insensitive, "OUT"},
		{"radio_checked_disabled", radio, "OPTION", "radiobutton", gtkrc.Insensitive, "IN"},
	} {
		tex := textureFrom(asNRGBA(g.icon(v.st, isize, isize, gtkrc.Call{Function: v.fn, Detail: v.detail, State: v.state, Shadow: v.shadow})))
		t.SetIcon(v.name, "CheckBox", tex)
		if !strings.HasPrefix(v.name, "radio") {
			t.SetIcon(v.name, "CheckButton", tex)
			t.SetIcon(v.name+"_mirrored", "CheckButton", tex)
		}
	}

	// Text fields.
	entry := g.style("GtkEntry")
	epad := pad(float64(entry.XThickness+2), float64(entry.YThickness+1))
	entryBox := func(state int, focused bool) StyleBox.Instance {
		return g.box(entry, 120, 22, epad,
			gtkrc.Call{Function: "FLAT_BOX", Detail: "entry_bg", State: state},
			gtkrc.Call{Function: "SHADOW", Detail: "entry", Shadow: "IN", State: state, Focused: focused})
	}
	t.SetStylebox("normal", "LineEdit", entryBox(gtkrc.Normal, false))
	t.SetStylebox("focus", "LineEdit", entryBox(gtkrc.Normal, true))
	t.SetStylebox("read_only", "LineEdit", entryBox(gtkrc.Insensitive, false))
	view := g.style("GtkScrolledWindow.GtkTextView")
	sw := g.style("GtkScrolledWindow")
	textBox := g.box(sw, 120, 64, pad(float64(sw.XThickness+2), float64(sw.YThickness+2)),
		gtkrc.Call{Function: "FLAT_BOX", Detail: "base", State: gtkrc.Normal},
		gtkrc.Call{Function: "SHADOW", Detail: "scrolled_window", Shadow: "IN"})
	_ = view
	for _, name := range []string{"normal", "focus", "read_only"} {
		t.SetStylebox(name, "TextEdit", textBox)
	}
	for _, typ := range []string{"LineEdit", "TextEdit"} {
		t.SetColor("font_color", typ, rc2(entry.Text[gtkrc.Normal]))
		t.SetColor("caret_color", typ, rc2(entry.Text[gtkrc.Normal]))
		t.SetColor("font_placeholder_color", typ, rc2(entry.Text[gtkrc.Insensitive]))
		t.SetColor("font_selected_color", typ, rc2(entry.Text[gtkrc.Selected]))
		t.SetColor("selection_color", typ, rc2(entry.Base[gtkrc.Selected]))
		t.SetColor("font_readonly_color", typ, rc2(entry.Text[gtkrc.Insensitive]))
		t.SetColor("background_color", typ, Color.RGBA{})
	}

	// Spin boxes: GTK 2 stacks the buttons too.
	spin := g.style("GtkSpinButton")
	for _, s := range []struct {
		godot  string
		state  int
		shadow string
	}{{"", gtkrc.Normal, "OUT"}, {"_hovered", gtkrc.Prelight, "OUT"}, {"_pressed", gtkrc.Active, "IN"}, {"_disabled", gtkrc.Insensitive, "OUT"}} {
		t.SetStylebox("up_background"+s.godot, "SpinBox", g.box(spin, 16, 11, [4]float64{}, gtkrc.Call{Function: "BOX", Detail: "spinbutton_up", State: s.state, Shadow: s.shadow}))
		t.SetStylebox("down_background"+s.godot, "SpinBox", g.box(spin, 16, 11, [4]float64{}, gtkrc.Call{Function: "BOX", Detail: "spinbutton_down", State: s.state, Shadow: s.shadow}))
	}
	for _, s := range []struct {
		suffix string
		state  int
	}{{"", gtkrc.Normal}, {"_hover", gtkrc.Prelight}, {"_pressed", gtkrc.Active}, {"_disabled", gtkrc.Insensitive}} {
		t.SetIcon("up"+s.suffix, "SpinBox", textureFrom(asNRGBA(g.icon(spin, 7, 4, gtkrc.Call{Function: "ARROW", Detail: "spinbutton", ArrowDirection: "UP", State: s.state, Shadow: "OUT"}))))
		t.SetIcon("down"+s.suffix, "SpinBox", textureFrom(asNRGBA(g.icon(spin, 7, 4, gtkrc.Call{Function: "ARROW", Detail: "spinbutton", ArrowDirection: "DOWN", State: s.state, Shadow: "OUT"}))))
	}
	for _, name := range []string{"up_icon_modulate", "up_hover_icon_modulate", "up_pressed_icon_modulate", "down_icon_modulate", "down_hover_icon_modulate", "down_pressed_icon_modulate"} {
		t.SetColor(name, "SpinBox", Color.RGBA{R: 1, G: 1, B: 1, A: 1})
	}
	t.SetStylebox("field_and_buttons_separator", "SpinBox", empty)
	t.SetStylebox("up_down_buttons_separator", "SpinBox", empty)
	t.SetConstant("buttons_width", "SpinBox", int(16*g.k))

	// Sliders: the trough is as tall as the slider is wide.
	scale := g.style("GtkHScale")
	sliderLen := scale.PropInt("GtkHScale", "slider-length", scale.PropInt("GtkScale", "slider-length", 31))
	sliderWidth := scale.PropInt("GtkHScale", "slider-width", scale.PropInt("GtkScale", "slider-width", scale.PropInt("GtkRange", "slider-width", 14)))
	trough := g.box(scale, 100, sliderWidth, [4]float64{float64(sliderWidth) / 2, 0, float64(sliderWidth) / 2, 0},
		gtkrc.Call{Function: "BOX", Detail: "trough", Shadow: "IN", State: gtkrc.Active, Orientation: "HORIZONTAL", Widget: "GtkHScale"})
	t.SetStylebox("slider", "HSlider", trough)
	t.SetStylebox("grabber_area", "HSlider", empty)
	t.SetStylebox("grabber_area_highlight", "HSlider", empty)
	side, _ := scale.Prop("GtkRange", "trough-side-details")
	if v, ok := scale.Prop("GtkScale", "trough-side-details"); ok {
		side = v
	}
	if side == "1" || strings.EqualFold(side, "TRUE") {
		// the part before the slider is drawn as GTK's "trough-lower"
		lower := g.box(scale, 32, sliderWidth, [4]float64{float64(sliderWidth) / 2, 0, float64(sliderWidth) / 2, 0},
			gtkrc.Call{Function: "BOX", Detail: "trough-lower", Shadow: "IN", State: gtkrc.Active, Orientation: "HORIZONTAL"})
		t.SetStylebox("grabber_area", "HSlider", lower)
		t.SetStylebox("grabber_area_highlight", "HSlider", lower)
	}
	for name, state := range map[string]int{"grabber": gtkrc.Normal, "grabber_highlight": gtkrc.Prelight, "grabber_disabled": gtkrc.Insensitive} {
		t.SetIcon(name, "HSlider", textureFrom(asNRGBA(g.icon(scale, sliderLen, sliderWidth, gtkrc.Call{Function: "SLIDER", Detail: "hscale", State: state, Shadow: "OUT", Orientation: "HORIZONTAL"}))))
	}
	t.SetConstant("center_grabber", "HSlider", 0)

	// Progress bars.
	progress := g.style("GtkProgressBar")
	ph := progress.PropInt("GtkProgressBar", "min-horizontal-bar-height", 16) + 2*progress.YThickness
	ptBox := g.box(progress, 120, ph, [4]float64{float64(ph) / 2, 0, float64(ph) / 2, 0},
		gtkrc.Call{Function: "BOX", Detail: "trough", Shadow: "IN", State: gtkrc.Normal, Widget: "GtkProgressBar"})
	t.SetStylebox("background", "ProgressBar", ptBox)
	// drawn several stripe periods wide and tiled, not stretched, so that
	// engines' diagonal stripes keep their shape
	barH := ph - 2*progress.YThickness
	fill := g.box(progress, 4*2*barH+2*progress.XThickness+2, barH, [4]float64{},
		gtkrc.Call{Function: "BOX", Detail: "bar", Shadow: "OUT", State: gtkrc.Prelight, Widget: "GtkProgressBar"})
	if tex, ok := Object.As[StyleBoxTexture.Instance](fill); ok {
		tex.SetAxisStretchHorizontal(StyleBoxTexture.AxisStretchModeTile)
		// keep rounded ends out of the repeated middle
		end := Float.X(float64(barH)/2+float64(progress.XThickness)+1) * Float.X(g.k)
		tex.SetTextureMarginLeft(end)
		tex.SetTextureMarginRight(end)
	}
	if tex, ok := Object.As[StyleBoxTexture.Instance](fill); ok {
		tex.SetExpandMarginTop(-Float.X(float64(progress.YThickness) * g.k))
		tex.SetExpandMarginBottom(-Float.X(float64(progress.YThickness) * g.k))
		tex.SetExpandMarginLeft(-Float.X(float64(progress.XThickness) * g.k))
	}
	t.SetStylebox("fill", "ProgressBar", fill)
	t.SetColor("font_color", "ProgressBar", rc2(progress.Fg[gtkrc.Normal]))

	// Separators.
	sep := g.style("GtkHSeparator")
	t.SetStylebox("separator", "HSeparator", g.box(sep, 32, 4, [4]float64{2, 0, 2, 0}, gtkrc.Call{Function: "HLINE", Detail: "hseparator"}))
	t.SetConstant("separation", "HSeparator", int(8*g.k))

	// Tabs.
	nb := g.style("GtkNotebook")
	overlap := nb.PropInt("GtkNotebook", "tab-overlap", 2)
	for _, typ := range []string{"TabContainer", "TabBar"} {
		for name, state := range map[string]int{"tab_selected": gtkrc.Normal, "tab_unselected": gtkrc.Active, "tab_hovered": gtkrc.Active, "tab_disabled": gtkrc.Insensitive} {
			h := 22
			if state != gtkrc.Normal {
				h = 20 // GTK 2 draws the other tabs a little lower
			}
			content := pad(float64(nb.XThickness+4), float64(nb.YThickness+1))
			if state != gtkrc.Normal {
				content[0] += 2
			}
			sb := g.box(nb, 60, h, content,
				gtkrc.Call{Function: "EXTENSION", Detail: "tab", State: state, GapSide: "BOTTOM", Shadow: "OUT"})
			if tex, ok := Object.As[StyleBoxTexture.Instance](sb); ok && state == gtkrc.Normal {
				tex.SetExpandMarginLeft(Float.X(float64(overlap) / 2 * g.k))
				tex.SetExpandMarginRight(Float.X(float64(overlap) / 2 * g.k))
				tex.SetExpandMarginBottom(Float.X(float64(nb.YThickness) * g.k))
			}
			t.SetStylebox(name, typ, sb)
		}
		t.SetStylebox("tab_focus", typ, empty)
		t.SetStylebox("tabbar_background", typ, empty)
		t.SetColor("font_selected_color", typ, rc2(nb.Fg[gtkrc.Normal]))
		t.SetColor("font_unselected_color", typ, rc2(nb.Fg[gtkrc.Active]))
		t.SetColor("font_hovered_color", typ, rc2(nb.Fg[gtkrc.Prelight]))
		t.SetColor("font_disabled_color", typ, rc2(nb.Fg[gtkrc.Insensitive]))
		t.SetConstant("side_margin", typ, int(2*g.k))
	}
	t.SetStylebox("panel", "TabContainer", g.box(nb, 64, 64, pad(8, 8),
		gtkrc.Call{Function: "BOX_GAP", Detail: "notebook", Shadow: "OUT", GapSide: "TOP", State: gtkrc.Normal}))

	// Menus.
	menu := g.rc.Resolve(gtkrc.W("GtkWindow.GtkMenu"))
	item := g.rc.Resolve(gtkrc.W("GtkWindow.GtkMenu.GtkMenuItem"))
	t.SetStylebox("panel", "PopupMenu", g.box(menu, 64, 64, pad(float64(menu.XThickness+1), float64(menu.YThickness+1)),
		gtkrc.Call{Function: "BOX", Detail: "menu", Shadow: "OUT", State: gtkrc.Normal}))
	t.SetStylebox("hover", "PopupMenu", g.box(item, 64, 20, pad(4, 2),
		gtkrc.Call{Function: "BOX", Detail: "menuitem", Shadow: "OUT", State: gtkrc.Prelight}))
	t.SetStylebox("separator", "PopupMenu", g.box(item, 32, 4, [4]float64{2, 0, 2, 0}, gtkrc.Call{Function: "HLINE", Detail: "menuitem"}))
	t.SetColor("font_color", "PopupMenu", rc2(item.Fg[gtkrc.Normal]))
	t.SetColor("font_accelerator_color", "PopupMenu", rc2(item.Fg[gtkrc.Normal]))
	t.SetColor("font_hover_color", "PopupMenu", rc2(item.Fg[gtkrc.Prelight]))
	t.SetColor("font_disabled_color", "PopupMenu", rc2(item.Fg[gtkrc.Insensitive]))
	checkItem := g.rc.Resolve(gtkrc.W("GtkWindow.GtkMenu.GtkCheckMenuItem"))
	csize := checkItem.PropInt("GtkCheckMenuItem", "indicator-size", 13)
	t.SetIcon("checked", "PopupMenu", textureFrom(asNRGBA(g.icon(checkItem, csize, csize, gtkrc.Call{Function: "CHECK", Detail: "check", Shadow: "IN"}))))
	t.SetIcon("unchecked", "PopupMenu", textureFrom(image.NewNRGBA(image.Rect(0, 0, csize, csize))))
	t.SetConstant("v_separation", "PopupMenu", int(6*g.k))
	t.SetConstant("item_start_padding", "PopupMenu", int(4*g.k))
	t.SetConstant("item_end_padding", "PopupMenu", int(8*g.k))
	t.SetStylebox("embedded_border", "Window", empty)
	t.SetStylebox("embedded_unfocused_border", "Window", empty)
	return t, nil
}

func asNRGBA(img image.Image) *image.NRGBA {
	if n, ok := img.(*image.NRGBA); ok {
		return n
	}
	b := img.Bounds()
	out := image.NewNRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	for y := 0; y < b.Dy(); y++ {
		for x := 0; x < b.Dx(); x++ {
			out.Set(x, y, img.At(b.Min.X+x, b.Min.Y+y))
		}
	}
	return out
}

// listCombo draws OptionButtons as GTK 2 combo boxes in "appears-as-list"
// mode: a sunken field with the arrow button beside it.
func (g *gtk2Builder) listCombo(button *gtkrc.Style, arrowW int) {
	t := g.theme
	entry := g.style("GtkEntry")
	const h = 22
	btnW := arrowW + 2*button.XThickness + 4
	field := image.NewNRGBA(image.Rect(0, 0, 64, h))
	entry.DrawInto(field, field.Bounds(), gtkrc.Call{Function: "FLAT_BOX", Detail: "entry_bg", State: gtkrc.Normal})
	entry.DrawInto(field, field.Bounds(), gtkrc.Call{Function: "SHADOW", Detail: "entry", Shadow: "IN"})
	fieldTex := textureFrom(g.scaled(field))
	k := Float.X(g.k)
	em := [4]Float.X{Float.X(entry.XThickness) * k, Float.X(entry.YThickness) * k, Float.X(entry.XThickness) * k, Float.X(entry.YThickness) * k}
	bm := [4]Float.X{Float.X(button.XThickness) * k, Float.X(button.YThickness) * k, Float.X(button.XThickness) * k, Float.X(button.YThickness) * k}
	for _, s := range []struct {
		godot  string
		state  int
		shadow string
	}{{"normal", gtkrc.Normal, "OUT"}, {"hover", gtkrc.Prelight, "OUT"}, {"pressed", gtkrc.Active, "IN"}, {"hover_pressed", gtkrc.Active, "IN"}, {"disabled", gtkrc.Insensitive, "OUT"}} {
		btn := button.Draw(gtkrc.Call{Function: "BOX", Detail: "button", State: s.state, Shadow: s.shadow}, btnW, h)
		sb := newLayeredBox(
			layer{tex: fieldTex, anchor: fillRect, width: Float.X(btnW) * k, margin: em},
			layer{tex: textureFrom(g.scaled(btn)), anchor: rightEdge, width: Float.X(btnW) * k, margin: bm},
		)
		sb.SetContentMarginLeft(Float.X(entry.XThickness+2) * k)
		sb.SetContentMarginRight(Float.X(btnW+2) * k)
		sb.SetContentMarginTop(Float.X(entry.YThickness+1) * k)
		sb.SetContentMarginBottom(Float.X(entry.YThickness+1) * k)
		t.SetStylebox(s.godot, "OptionButton", sb)
		t.SetColor(godotFontColors[s.godot], "OptionButton", rc2(entry.Text[gtkrc.Normal]))
	}
	t.SetColor("font_disabled_color", "OptionButton", rc2(entry.Text[gtkrc.Insensitive]))
	t.SetColor("font_focus_color", "OptionButton", rc2(entry.Text[gtkrc.Normal]))
	t.SetStylebox("focus", "OptionButton", StyleBoxEmpty.New().AsStyleBox())
	t.SetConstant("arrow_margin", "OptionButton", int(float64(btnW-arrowW)/2*g.k))
}

// leftIndicators draws CheckButtons as GTK 2 check buttons (GTK 2 has no
// switch): the indicator left of the label, drawn by the state's stylebox
// since Godot puts CheckButton icons on the right. Godot uses the same
// "disabled" stylebox whether the button is on or off, so disabled ones
// always show as off.
func (g *gtk2Builder) leftIndicators(check *gtkrc.Style, size int) {
	t := g.theme
	k := Float.X(g.k)
	spacing := check.PropInt("GtkCheckButton", "indicator-spacing", 2)
	for _, s := range []struct {
		godot  string
		state  int
		shadow string
	}{{"normal", gtkrc.Normal, "OUT"}, {"hover", gtkrc.Prelight, "OUT"}, {"pressed", gtkrc.Normal, "IN"}, {"hover_pressed", gtkrc.Prelight, "IN"}, {"disabled", gtkrc.Insensitive, "OUT"}} {
		ind := check.Draw(gtkrc.Call{Function: "CHECK", Detail: "checkbutton", State: s.state, Shadow: s.shadow}, size, size)
		sb := newLayeredBox(layer{tex: textureFrom(g.scaled(ind)), anchor: leftEdge})
		sb.SetContentMarginLeft(Float.X(size+2*spacing+2) * k)
		sb.SetContentMarginTop(Float.X(spacing) * k)
		sb.SetContentMarginBottom(Float.X(spacing) * k)
		t.SetStylebox(s.godot, "CheckButton", sb)
	}
	none := textureFrom(image.NewNRGBA(image.Rect(0, 0, 1, 1)))
	for _, name := range []string{"checked", "unchecked", "checked_disabled", "unchecked_disabled"} {
		t.SetIcon(name, "CheckButton", none)
		t.SetIcon(name+"_mirrored", "CheckButton", none)
	}
	t.SetConstant("h_separation", "CheckButton", 0)
}
