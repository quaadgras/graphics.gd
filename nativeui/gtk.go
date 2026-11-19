//go:build (linux && !android) || freebsd || openbsd || netbsd || dragonfly || nativeui_all

package nativeui

import (
	"image"
	"image/draw"
	"math"
	"os"
	"path/filepath"

	"graphics.gd/classdb/Image"
	"graphics.gd/classdb/ImageTexture"
	"graphics.gd/classdb/StyleBox"
	"graphics.gd/classdb/StyleBoxEmpty"
	"graphics.gd/classdb/StyleBoxFlat"
	"graphics.gd/classdb/StyleBoxTexture"
	"graphics.gd/classdb/SystemFont"
	"graphics.gd/classdb/Texture2D"
	"graphics.gd/classdb/Theme"
	"graphics.gd/nativeui/internal/gtkcss"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
)

func init() { gtkTheme, gtk2Theme = buildGTK, buildGTK2 }

// gtkSheet loads the GTK 3 stylesheet of a theme directory (one containing
// gtk-3.0/gtk.css), preferring gtk-dark.css in dark mode.
func gtkSheet(dir string, dark bool) (*gtkcss.Sheet, error) {
	if dark {
		if path := filepath.Join(dir, "gtk-3.0", "gtk-dark.css"); fileExists(path) {
			return gtkcss.Load(path)
		}
	}
	return gtkcss.Load(filepath.Join(dir, "gtk-3.0", "gtk.css"))
}

func fileExists(path string) bool { _, err := os.Stat(path); return err == nil }

type gtkBuilder struct {
	sheet  *gtkcss.Sheet
	theme  Theme.Instance
	k      float64
	family string // the theme's default font
	weight int
}

// root is prefixed to every node path: widgets live in a window.
// (inside a container, as in any real window: some themes style direct
// children of the window differently).
const gtkRoot = "window.background box.vertical "

func (g *gtkBuilder) style(path string) gtkcss.Style {
	return scaleStyle(g.sheet.Compute(gtkcss.Path(gtkRoot+path)), g.k)
}

func scaleStyle(st gtkcss.Style, k float64) gtkcss.Style {
	if k == 1 {
		return st
	}
	for i := range 4 {
		st.BorderWidth[i] = math.Round(st.BorderWidth[i] * k)
		st.Radius[i] *= k
		st.Padding[i] *= k
		st.Margin[i] *= k
	}
	st.MinWidth *= k
	st.MinHeight *= k
	for i := range st.Shadows {
		sh := &st.Shadows[i]
		sh.X, sh.Y, sh.Blur, sh.Spread = math.Round(sh.X*k), math.Round(sh.Y*k), sh.Blur*k, sh.Spread*k
	}
	st.OutlineWidth = math.Round(st.OutlineWidth * k)
	st.OutlineOffset = math.Round(st.OutlineOffset * k)
	st.FontSize *= k
	return st
}

func gdColor(c gtkcss.RGBA) Color.RGBA {
	return Color.RGBA{R: Float.X(c.R), G: Float.X(c.G), B: Float.X(c.B), A: Float.X(c.A)}
}

func (g *gtkBuilder) texture(img *image.NRGBA) Texture2D.Instance { return textureFrom(img) }

func textureFrom(img *image.NRGBA) Texture2D.Instance {
	b := img.Bounds()
	return ImageTexture.CreateFromImage(Image.CreateFromData(b.Dx(), b.Dy(), false, Image.FormatRgba8, img.Pix)).AsTexture2D()
}

// iconImage loads an -gtk-icon-source image, scaled to the theme's scale and
// tinted if it is a symbolic one.
func (g *gtkBuilder) iconImage(src gtkcss.Image, tint gtkcss.RGBA) *image.NRGBA {
	if src.Kind != "url" {
		return nil
	}
	img := gtkcss.LoadImage(src.Path)
	if img == nil {
		return nil
	}
	out := scaleNearest(img, g.k)
	if src.Recolor {
		for i := 0; i < len(out.Pix); i += 4 {
			out.Pix[i+0] = uint8(tint.R*255 + 0.5)
			out.Pix[i+1] = uint8(tint.G*255 + 0.5)
			out.Pix[i+2] = uint8(tint.B*255 + 0.5)
			out.Pix[i+3] = uint8(float64(out.Pix[i+3]) * tint.A)
		}
	}
	return out
}

func scaleNearest(img image.Image, k float64) *image.NRGBA {
	b := img.Bounds()
	w, h := int(math.Round(float64(b.Dx())*k)), int(math.Round(float64(b.Dy())*k))
	out := image.NewNRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			out.Set(x, y, img.At(b.Min.X+int(float64(x)/k), b.Min.Y+int(float64(y)/k)))
		}
	}
	return out
}

// box is the StyleBox for a node drawn at a typical size; Godot's 9-slice
// stretches it to the control's size.
func (g *gtkBuilder) box(st gtkcss.Style, w, h int) StyleBox.Instance {
	if !visible(st) {
		sb := StyleBoxEmpty.New().AsStyleBox()
		g.contentMargins(sb, st)
		return sb
	}
	w, h = int(float64(w)*g.k), int(float64(h)*g.k)
	img := gtkcss.Render(st, w, h)
	ext := st.Outside()
	sb := StyleBoxTexture.New()
	sb.SetTexture(g.texture(img))
	// The fixed (non-stretched) margins hold the borders, corners and inset shadows.
	var reach [4]float64
	for _, sh := range st.Shadows {
		if !sh.Inset {
			continue
		}
		r := sh.Blur + sh.Spread
		reach[0] = math.Max(reach[0], math.Max(sh.Y, 0)+r)
		reach[1] = math.Max(reach[1], math.Max(-sh.X, 0)+r)
		reach[2] = math.Max(reach[2], math.Max(-sh.Y, 0)+r)
		reach[3] = math.Max(reach[3], math.Max(sh.X, 0)+r)
	}
	if st.OutlineWidth > 0 {
		o := math.Max(0, -st.OutlineOffset)
		for i := range reach {
			reach[i] = math.Max(reach[i], o)
		}
	}
	corner := [4]float64{
		math.Max(st.Radius[0], st.Radius[1]), math.Max(st.Radius[1], st.Radius[2]),
		math.Max(st.Radius[2], st.Radius[3]), math.Max(st.Radius[3], st.Radius[0]),
	}
	var m [4]float64
	for i := range m {
		m[i] = math.Ceil(st.BorderWidth[i]+math.Max(reach[i], corner[i])) + float64(ext[i])
	}
	limit := func(a, b *float64, size int) { // leave at least a pixel to stretch
		for *a+*b >= float64(size) && (*a > 0 || *b > 0) {
			if *a >= *b {
				*a--
			} else {
				*b--
			}
		}
	}
	limit(&m[0], &m[2], img.Bounds().Dy())
	limit(&m[3], &m[1], img.Bounds().Dx())
	sb.SetTextureMarginTop(Float.X(m[0]))
	sb.SetTextureMarginRight(Float.X(m[1]))
	sb.SetTextureMarginBottom(Float.X(m[2]))
	sb.SetTextureMarginLeft(Float.X(m[3]))
	sb.SetExpandMarginTop(Float.X(ext[0]))
	sb.SetExpandMarginRight(Float.X(ext[1]))
	sb.SetExpandMarginBottom(Float.X(ext[2]))
	sb.SetExpandMarginLeft(Float.X(ext[3]))
	if patterned(st) {
		sb.SetAxisStretchHorizontal(StyleBoxTexture.AxisStretchModeTileFit)
		sb.SetAxisStretchVertical(StyleBoxTexture.AxisStretchModeTileFit)
	}
	g.contentMargins(sb.AsStyleBox(), st)
	return sb.AsStyleBox()
}

func visible(st gtkcss.Style) bool {
	if st.BackgroundColor.A > 0 || len(st.BackgroundImage) > 0 || len(st.Shadows) > 0 || st.OutlineWidth > 0 || st.BorderImage.Kind != "" {
		return true
	}
	for i := range st.BorderWidth {
		if st.BorderWidth[i] > 0 && st.BorderColor[i].A > 0 {
			return true
		}
	}
	return false
}

func patterned(st gtkcss.Style) bool {
	if st.OutlineWidth > 0 && st.OutlineStyle != "solid" {
		return true
	}
	for _, s := range st.BorderStyle {
		if s == "dotted" || s == "dashed" {
			return true
		}
	}
	return false
}

func (g *gtkBuilder) contentMargins(sb StyleBox.Instance, st gtkcss.Style) {
	sb.SetContentMarginTop(Float.X(st.BorderWidth[0] + st.Padding[0]))
	sb.SetContentMarginRight(Float.X(st.BorderWidth[1] + st.Padding[1]))
	sb.SetContentMarginBottom(Float.X(st.BorderWidth[2] + st.Padding[2]))
	sb.SetContentMarginLeft(Float.X(st.BorderWidth[3] + st.Padding[3]))
}

func noOutline(st gtkcss.Style) gtkcss.Style {
	st.OutlineWidth = 0
	return st
}

// outlineOnly keeps just a node's focus outline, for Godot's focus styleboxes.
func outlineOnly(st gtkcss.Style) gtkcss.Style {
	return gtkcss.Style{OutlineWidth: st.OutlineWidth, OutlineOffset: st.OutlineOffset, OutlineStyle: st.OutlineStyle,
		OutlineColor: st.OutlineColor, Opacity: 1}
}

// states maps Godot state names to GTK pseudo-classes.
var buttonStates = []struct{ godot, gtk string }{
	{"normal", ""}, {"hover", ":hover"}, {"pressed", ":active"}, {"hover_pressed", ":hover:active"}, {"disabled", ":disabled"},
}

var godotFontColors = map[string]string{
	"normal": "font_color", "hover": "font_hover_color", "pressed": "font_pressed_color",
	"hover_pressed": "font_hover_pressed_color", "disabled": "font_disabled_color",
}

func (g *gtkBuilder) buttonLike(themeType, path string, w, h int) {
	if lw := g.style(path + " label").FontWeight; lw != g.weight {
		g.theme.SetFont("font", themeType, g.font(g.family, lw).AsFont())
	}
	for _, s := range buttonStates {
		st := g.style(path + s.gtk)
		g.theme.SetStylebox(s.godot, themeType, g.box(st, w, h))
		g.theme.SetColor(godotFontColors[s.godot], themeType, gdColor(g.style(path+s.gtk+" label").Color))
	}
	g.theme.SetColor("font_focus_color", themeType, gdColor(g.style(path+":focus label").Color))
	focus := outlineOnly(g.style(path + ":focus"))
	if focus.OutlineWidth > 0 {
		g.theme.SetStylebox("focus", themeType, g.box(focus, w, h))
	} else {
		g.theme.SetStylebox("focus", themeType, StyleBoxEmpty.New().AsStyleBox())
	}
}

// iconNRGBA renders a node at its minimum size with its -gtk-icon-source
// centered on it (GTK draws check and radio indicators this way).
func (g *gtkBuilder) iconNRGBA(path string) *image.NRGBA {
	st := g.style(path)
	icon := g.iconImage(st.IconSource, st.Color)
	w := int(math.Max(1, st.MinWidth+st.Padding[1]+st.Padding[3]+st.BorderWidth[1]+st.BorderWidth[3]))
	h := int(math.Max(1, st.MinHeight+st.Padding[0]+st.Padding[2]+st.BorderWidth[0]+st.BorderWidth[2]))
	if !visible(st) && icon != nil {
		return icon
	}
	img := gtkcss.Render(st, w, h)
	if icon != nil {
		ext := st.Outside()
		ib := icon.Bounds()
		at := image.Pt(ext[3]+(w-ib.Dx())/2, ext[0]+(h-ib.Dy())/2)
		draw.Draw(img, ib.Add(at), icon, image.Point{}, draw.Over)
	}
	return img
}

func (g *gtkBuilder) font(family string, weight int) SystemFont.Instance {
	f := SystemFont.New()
	// Godot resolves "sans-serif", not fontconfig's "Sans" alias.
	f.SetFontNames(fontNames(family, "sans-serif"))
	f.SetFontWeight(weight)
	return f
}

func buildGTK(opts Options) (Theme.Instance, error) {
	sheet, err := gtkSheet(opts.GTKTheme, opts.Dark)
	if err != nil {
		return Theme.Nil, err
	}
	g := &gtkBuilder{sheet: sheet, theme: Theme.New(), k: float64(opts.Scale), weight: 400}
	t := g.theme
	empty := StyleBoxEmpty.New().AsStyleBox()

	win := scaleStyle(sheet.Compute(gtkcss.Path("window.background")), g.k)
	label := g.style("label")
	t.SetDefaultBaseScale(Float.X(g.k))
	// The desktop's font, unless the theme's CSS sets one.
	family, weight, size := opts.Font.Family, opts.Font.Weight, float64(opts.Font.Size)*g.k
	if label.FontFamily != "" {
		family = label.FontFamily
	}
	if weight == 0 {
		weight = label.FontWeight
	}
	if size <= 0 {
		size = label.FontSize
	}
	g.family, g.weight = family, weight
	t.SetDefaultFont(g.font(family, weight).AsFont())
	t.SetDefaultFontSize(int(math.Round(size)))
	bold := g.font(family, max(weight, 700)).AsFont()

	// Window background, labels and frames.
	bg := StyleBoxFlat.New()
	bg.SetBgColor(gdColor(win.BackgroundColor))
	t.SetStylebox("panel", "Panel", bg.AsStyleBox())
	t.SetStylebox("panel", "PanelContainer", bg.AsStyleBox())
	t.SetColor("font_color", "Label", gdColor(label.Color))
	t.SetTypeVariation(HeaderLabel, "Label")
	t.SetFont("font", HeaderLabel, bold)
	t.SetTypeVariation(DimLabel, "Label")
	t.SetColor("font_color", DimLabel, gdColor(g.style("label.dim-label").Color))
	t.SetTypeVariation(Card, "PanelContainer")
	frame := g.style("frame border")
	frame.Padding = [4]float64{8 * g.k, 8 * g.k, 8 * g.k, 8 * g.k}
	t.SetStylebox("panel", Card, g.box(frame, 64, 64))

	// Buttons.
	g.buttonLike("Button", "button.text-button", 75, 23)
	t.SetTypeVariation(AccentButton, "Button")
	g.buttonLike(AccentButton, "button.text-button.suggested-action", 75, 23)
	g.buttonLike("OptionButton", "combobox box.linked button.combo", 100, 23)
	arrow := g.style("combobox box.linked button.combo")
	t.SetIcon("arrow", "OptionButton", g.texture(triangle(int(7*g.k), arrow.Color, false)))
	t.SetConstant("arrow_margin", "OptionButton", int(4*g.k))

	// Check boxes, radio buttons and switches.
	for _, typ := range []string{"CheckBox", "CheckButton"} {
		for _, name := range []string{"normal", "hover", "pressed", "hover_pressed", "disabled", "focus"} {
			t.SetStylebox(name, typ, empty)
		}
		lab := g.style("checkbutton label")
		t.SetColor("font_color", typ, gdColor(lab.Color))
		t.SetColor("font_hover_color", typ, gdColor(g.style("checkbutton:hover label").Color))
		t.SetColor("font_pressed_color", typ, gdColor(g.style("checkbutton:checked label").Color))
		t.SetColor("font_hover_pressed_color", typ, gdColor(g.style("checkbutton:checked:hover label").Color))
		t.SetColor("font_focus_color", typ, gdColor(lab.Color))
		t.SetColor("font_disabled_color", typ, gdColor(g.style("checkbutton:disabled label").Color))
		t.SetConstant("h_separation", typ, int(lab.Padding[3]+lab.BorderWidth[3]+2*g.k))
	}
	for name, path := range map[string]string{
		"unchecked": "checkbutton check", "checked": "checkbutton check:checked",
		"unchecked_disabled": "checkbutton check:disabled", "checked_disabled": "checkbutton check:disabled:checked",
		"radio_unchecked": "radiobutton radio", "radio_checked": "radiobutton radio:checked",
		"radio_unchecked_disabled": "radiobutton radio:disabled", "radio_checked_disabled": "radiobutton radio:disabled:checked",
	} {
		t.SetIcon(name, "CheckBox", g.texture(g.iconNRGBA(path)))
	}
	for _, on := range []bool{false, true} {
		for _, disabled := range []bool{false, true} {
			name := map[bool]string{false: "unchecked", true: "checked"}[on]
			if disabled {
				name += "_disabled"
			}
			tex := g.texture(g.switchImage(on, disabled))
			t.SetIcon(name, "CheckButton", tex)
			t.SetIcon(name+"_mirrored", "CheckButton", tex)
		}
	}

	// Text fields.
	// GTK entries never draw a focus outline (only buttons and the like do).
	entry := noOutline(g.style("entry"))
	sel := g.style("entry selection")
	placeholder := entry.Color
	placeholder.A *= 0.5
	for _, typ := range []string{"LineEdit", "TextEdit"} {
		t.SetColor("font_color", typ, gdColor(entry.Color))
		t.SetColor("caret_color", typ, gdColor(entry.Color))
		t.SetColor("font_placeholder_color", typ, gdColor(placeholder))
		t.SetColor("font_selected_color", typ, gdColor(sel.Color))
		t.SetColor("selection_color", typ, gdColor(sel.BackgroundColor))
		t.SetColor("font_readonly_color", typ, gdColor(g.style("entry:disabled").Color))
		t.SetColor("background_color", typ, Color.RGBA{})
	}
	t.SetStylebox("normal", "LineEdit", g.box(entry, 120, 22))
	t.SetStylebox("focus", "LineEdit", g.box(noOutline(g.style("entry:focus")), 120, 22))
	t.SetStylebox("read_only", "LineEdit", g.box(noOutline(g.style("entry:disabled")), 120, 22))
	text := g.style("scrolledwindow.frame")
	if !visible(text) || text.BorderWidth == [4]float64{} {
		text = entry // themes without a framed scrolled window: draw it like an entry
	}
	text.BackgroundColor = g.style("scrolledwindow.frame textview.view text").BackgroundColor
	text.Padding = [4]float64{2 * g.k, 2 * g.k, 2 * g.k, 2 * g.k}
	t.SetStylebox("normal", "TextEdit", g.box(text, 120, 64))
	t.SetStylebox("focus", "TextEdit", g.box(text, 120, 64))
	t.SetStylebox("read_only", "TextEdit", g.box(text, 120, 64))

	// Spin boxes: Godot stacks the buttons vertically, like Windows did.
	spinButton := "button.text-button"
	for _, s := range []struct{ godot, gtk string }{{"", ""}, {"_hovered", ":hover"}, {"_pressed", ":active"}, {"_disabled", ":disabled"}} {
		st := g.style(spinButton + s.gtk)
		st.Padding = [4]float64{}
		sb := g.box(st, 16, 11)
		t.SetStylebox("up_background"+s.godot, "SpinBox", sb)
		t.SetStylebox("down_background"+s.godot, "SpinBox", sb)
	}
	up, down := g.texture(triangle(int(7*g.k), label.Color, true)), g.texture(triangle(int(7*g.k), label.Color, false))
	for _, s := range []string{"", "_hover", "_pressed", "_disabled"} {
		t.SetIcon("up"+s, "SpinBox", up)
		t.SetIcon("down"+s, "SpinBox", down)
	}
	t.SetStylebox("field_and_buttons_separator", "SpinBox", empty)
	t.SetStylebox("up_down_buttons_separator", "SpinBox", empty)
	t.SetConstant("buttons_width", "SpinBox", int(16*g.k))

	// Sliders.
	trough := g.style("scale.horizontal contents trough")
	troughH := trough.MinHeight + trough.BorderWidth[0] + trough.BorderWidth[2] + trough.Padding[0] + trough.Padding[2]
	tb := g.box(trough, 100, int(math.Ceil(troughH/g.k)))
	tb.SetContentMarginTop(Float.X(troughH / 2))
	tb.SetContentMarginBottom(Float.X(troughH / 2))
	t.SetStylebox("slider", "HSlider", tb)
	highlight := g.style("scale.horizontal contents trough highlight")
	t.SetStylebox("grabber_area", "HSlider", g.box(highlight, 32, int(math.Ceil(troughH/g.k))))
	t.SetStylebox("grabber_area_highlight", "HSlider", g.box(highlight, 32, int(math.Ceil(troughH/g.k))))
	slider := g.style("scale.horizontal contents trough slider")
	sw := slider.MinWidth + slider.BorderWidth[1] + slider.BorderWidth[3] + slider.Padding[1] + slider.Padding[3]
	// negative margins grow the slider past the trough, into the scale's padding.
	scale := g.style("scale.horizontal")
	sh := troughH + math.Min(-slider.Margin[0], scale.Padding[0]) + math.Min(-slider.Margin[2], scale.Padding[2])
	if sh < slider.MinHeight+slider.BorderWidth[0]+slider.BorderWidth[2] {
		sh = slider.MinHeight + slider.BorderWidth[0] + slider.BorderWidth[2] + slider.Padding[0] + slider.Padding[2]
	}
	t.SetIcon("grabber", "HSlider", g.texture(gtkcss.Render(slider, int(sw), int(sh))))
	t.SetIcon("grabber_highlight", "HSlider", g.texture(gtkcss.Render(g.style("scale.horizontal contents trough slider:hover"), int(sw), int(sh))))
	t.SetIcon("grabber_disabled", "HSlider", g.texture(gtkcss.Render(g.style("scale.horizontal:disabled contents trough slider:disabled"), int(sw), int(sh))))
	t.SetConstant("center_grabber", "HSlider", 0)

	// Progress bars: GTK draws the bar inside the trough's padding.
	pt := g.style("progressbar.horizontal trough")
	ph := pt.MinHeight + pt.BorderWidth[0] + pt.BorderWidth[2] + pt.Padding[0] + pt.Padding[2]
	ptBox := g.box(pt, 120, int(math.Ceil(ph/g.k)))
	ptBox.SetContentMarginTop(Float.X(ph / 2)) // Godot sizes the bar from its background's margins
	ptBox.SetContentMarginBottom(Float.X(ph / 2))
	t.SetStylebox("background", "ProgressBar", ptBox)
	fill := g.box(g.style("progressbar.horizontal trough progress"), 32, int(math.Ceil(ph/g.k)))
	fill.SetContentMarginTop(0)
	fill.SetContentMarginBottom(0)
	if tex, ok := Object.As[StyleBoxTexture.Instance](fill); ok {
		tex.SetExpandMarginTop(-Float.X(pt.BorderWidth[0] + pt.Padding[0]))
		tex.SetExpandMarginBottom(-Float.X(pt.BorderWidth[2] + pt.Padding[2]))
		tex.SetExpandMarginLeft(-Float.X(pt.BorderWidth[3] + pt.Padding[3]))
	}
	t.SetStylebox("fill", "ProgressBar", fill)
	t.SetColor("font_color", "ProgressBar", gdColor(label.Color))

	// Separators.
	sep := g.style("separator.horizontal")
	sepH := math.Max(1, sep.MinHeight+sep.BorderWidth[0]+sep.BorderWidth[2])
	sepBox := g.box(sep, 32, int(math.Ceil(sepH/g.k)))
	sepBox.SetContentMarginTop(Float.X(sepH / 2))
	sepBox.SetContentMarginBottom(Float.X(sepH / 2))
	t.SetStylebox("separator", "HSeparator", sepBox)
	t.SetConstant("separation", "HSeparator", int(8*g.k))

	// Tabs.
	tab := "notebook header.top tabs tab"
	for _, typ := range []string{"TabContainer", "TabBar"} {
		for _, s := range []struct{ godot, gtk string }{{"tab_unselected", ""}, {"tab_hovered", ":hover"}, {"tab_selected", ":checked"}, {"tab_disabled", ":disabled"}} {
			st := g.style(tab + s.gtk)
			sb := g.box(st, 60, 20)
			// negative margins let the selected tab overlap its neighbours
			if tex, ok := Object.As[StyleBoxTexture.Instance](sb); ok {
				tex.SetExpandMarginLeft(Float.X(-st.Margin[3]))
				tex.SetExpandMarginRight(Float.X(-st.Margin[1]))
				tex.SetExpandMarginBottom(Float.X(-st.Margin[2]))
			}
			t.SetStylebox(s.godot, typ, sb)
		}
		t.SetStylebox("tab_focus", typ, empty)
		// GTK draws the line under the tabs on the header.
		hdr := g.style("notebook header.top")
		hdr.Padding = [4]float64{}
		t.SetStylebox("tabbar_background", typ, g.box(hdr, 64, 24))
		t.SetColor("font_selected_color", typ, gdColor(g.style(tab+":checked label").Color))
		t.SetColor("font_unselected_color", typ, gdColor(g.style(tab+" label").Color))
		t.SetColor("font_hovered_color", typ, gdColor(g.style(tab+":hover label").Color))
		t.SetColor("font_disabled_color", typ, gdColor(g.style(tab+":disabled label").Color))
	}
	// The page area: GTK 3 draws it on the notebook's stack, which follows the header.
	nb := g.style("notebook +stack")
	if !visible(nb) {
		nb = g.style("notebook")
	}
	nb.Padding = [4]float64{8 * g.k, 8 * g.k, 8 * g.k, 8 * g.k}
	if nb.BackgroundColor.A == 0 {
		nb.BackgroundColor = win.BackgroundColor
	}
	header := g.style("notebook header.top")
	for _, typ := range []string{"TabContainer", "TabBar"} {
		t.SetConstant("side_margin", typ, int(header.Margin[3]))
	}
	t.SetStylebox("panel", "TabContainer", g.box(nb, 64, 64))

	// Menus: GTK draws the popup window's decoration around the menu.
	menu := g.style("window.popup menu")
	if !visible(menu) || menu.BorderWidth == [4]float64{} {
		deco := g.style("window.popup decoration")
		if deco.BorderWidth != [4]float64{} || len(deco.Shadows) > 0 {
			deco.BackgroundColor = menu.BackgroundColor
			deco.Padding = menu.Padding
			menu = deco
		}
	}
	if menu.BorderWidth == [4]float64{} {
		menu = g.style("button.text-button") // no frame at all: use a raised bevel
		menu.BackgroundColor = g.style("window.popup menu").BackgroundColor
	}
	menu.Padding = [4]float64{2 * g.k, 2 * g.k, 2 * g.k, 2 * g.k}
	t.SetStylebox("panel", "PopupMenu", g.box(menu, 64, 64))
	item := g.style("window.popup menu menuitem:hover")
	hover := g.box(item, 64, 20)
	t.SetStylebox("hover", "PopupMenu", hover)
	msep := g.style("window.popup menu separator")
	msepH := math.Max(1, msep.MinHeight+msep.BorderWidth[0]+msep.BorderWidth[2])
	msepBox := g.box(msep, 32, int(math.Ceil(msepH/g.k)))
	msepBox.SetContentMarginTop(Float.X(msepH / 2))
	msepBox.SetContentMarginBottom(Float.X(msepH / 2))
	t.SetStylebox("separator", "PopupMenu", msepBox)
	mi := g.style("window.popup menu menuitem label")
	t.SetColor("font_color", "PopupMenu", gdColor(mi.Color))
	t.SetColor("font_accelerator_color", "PopupMenu", gdColor(mi.Color))
	t.SetColor("font_hover_color", "PopupMenu", gdColor(g.style("window.popup menu menuitem:hover label").Color))
	t.SetColor("font_disabled_color", "PopupMenu", gdColor(g.style("window.popup menu menuitem:disabled label").Color))
	t.SetConstant("v_separation", "PopupMenu", int(6*g.k))
	t.SetConstant("item_start_padding", "PopupMenu", int(4*g.k))
	t.SetConstant("item_end_padding", "PopupMenu", int(8*g.k))
	for name, path := range map[string]string{
		"checked": "window.popup menu menuitem check:checked", "unchecked": "window.popup menu menuitem check",
		"radio_checked": "window.popup menu menuitem radio:checked", "radio_unchecked": "window.popup menu menuitem radio",
		"checked_disabled": "window.popup menu menuitem:disabled check:checked:disabled", "unchecked_disabled": "window.popup menu menuitem:disabled check:disabled",
		"radio_checked_disabled": "window.popup menu menuitem:disabled radio:checked:disabled", "radio_unchecked_disabled": "window.popup menu menuitem:disabled radio:disabled",
	} {
		t.SetIcon(name, "PopupMenu", g.texture(g.iconNRGBA(path)))
	}
	t.SetStylebox("embedded_border", "Window", empty)
	t.SetStylebox("embedded_unfocused_border", "Window", empty)
	return t, nil
}

// switchImage composes a GTK 3 switch: the trough with its slider on one
// side and the I/O mark GTK draws on the other.
func (g *gtkBuilder) switchImage(on, disabled bool) *image.NRGBA {
	state := ""
	if on {
		state += ":checked"
	}
	if disabled {
		state += ":disabled"
	}
	sw := g.style("switch" + state)
	sl := g.style("switch" + state + " slider")
	slW := math.Max(sl.MinWidth, 1) + sl.BorderWidth[1] + sl.BorderWidth[3] + sl.Padding[1] + sl.Padding[3]
	slH := math.Max(sl.MinHeight, 1) + sl.BorderWidth[0] + sl.BorderWidth[2] + sl.Padding[0] + sl.Padding[2]
	half := slW + sl.Margin[1] + sl.Margin[3]
	inner := [4]float64{sw.BorderWidth[0] + sw.Padding[0], sw.BorderWidth[1] + sw.Padding[1], sw.BorderWidth[2] + sw.Padding[2], sw.BorderWidth[3] + sw.Padding[3]}
	w := int(math.Round(2*half + inner[1] + inner[3]))
	h := int(math.Round(slH + sl.Margin[0] + sl.Margin[2] + inner[0] + inner[2]))
	img := gtkcss.Render(sw, w, h)
	ext := sw.Outside()
	x := float64(ext[3]) + inner[3] + sl.Margin[3]
	if on {
		x += half
	}
	y := float64(ext[0]) + inner[0] + sl.Margin[0]
	slider := gtkcss.Render(sl, int(slW), int(slH))
	se := sl.Outside()
	draw.Draw(img, slider.Bounds().Add(image.Pt(int(x)-se[3], int(y)-se[0])), slider, image.Point{}, draw.Over)
	// GTK 3 marks the empty half: "|" when on, "o" when off (unless the
	// theme hides them).
	mark := g.style("switch" + state + " image").Color
	if mark.A == 0 {
		return img
	}
	cx := float64(ext[3]) + inner[3] + half/2
	if !on {
		cx += half
	}
	cy := float64(h) / 2
	r := math.Max(2, math.Round(slH*0.22))
	cv := canvasFrom(img)
	col := Color.RGBA{R: Float.X(mark.R), G: Float.X(mark.G), B: Float.X(mark.B), A: Float.X(mark.A)}
	if on {
		cv.fill(roundRect(math.Round(cx)+0.5, cy, 0.5*g.k, r, 0), col, 0)
	} else {
		cv.fill(ring(circle(math.Round(cx), cy, r), 1*g.k), col, 0)
	}
	return cv.nrgba()
}

// triangle is a solid pixel-art arrow of the given width.
func triangle(w int, c gtkcss.RGBA, up bool) *image.NRGBA {
	if w%2 == 0 {
		w++
	}
	h := (w + 1) / 2
	img := image.NewNRGBA(image.Rect(0, 0, w, h))
	col := Color.RGBA{R: Float.X(c.R), G: Float.X(c.G), B: Float.X(c.B), A: Float.X(c.A)}
	for row := 0; row < h; row++ {
		inset := row
		if up {
			inset = h - 1 - row
		}
		for x := inset; x < w-inset; x++ {
			img.Pix[(row*w+x)*4+0] = uint8(col.R*255 + 0.5)
			img.Pix[(row*w+x)*4+1] = uint8(col.G*255 + 0.5)
			img.Pix[(row*w+x)*4+2] = uint8(col.B*255 + 0.5)
			img.Pix[(row*w+x)*4+3] = uint8(col.A*255 + 0.5)
		}
	}
	return img
}
