// Package nativeui builds Godot themes that make Control nodes look like the
// native widgets of the platform the program is running on, including light
// and dark mode, the user's accent colour, display scale and font.
//
//	nativeui.Use() // every Control, now and whenever the system theme changes
//
// Use makes the native look Godot's default theme. New returns the same
// theme as a Theme of your own, as the system is when it's called.
//
// Everything is drawn by Godot, so the controls behave the same everywhere;
// they only look at home. nativeui is theme-only: it styles Godot's own
// Button, CheckBox, LineEdit, TabContainer... and adds no widgets of its
// own. Theme type variations mark the roles a theme can't infer, for use
// with Control.SetThemeTypeVariation:
//
//   - "AccentButton", a Button for the default action (Save, OK)
//   - "HeaderLabel", a Label heading a group of controls
//   - "DimLabel", a Label for secondary text
//   - "Card", a PanelContainer grouping related controls
//
// # Looks
//
// The look, and its generation, come from the running system:
//
//   - Windows: Windows 11 (WinUI/Fluent), Windows 10, or the user's
//     contrast theme; the accent shades, text size and transparency
//     setting from the registry, the scale from the DPI.
//   - macOS: macOS 11-15, or macOS 26 and later (Liquid Glass tab bars,
//     unless Reduce Transparency is on).
//   - iOS: iOS 15-18 or iOS 26 and later.
//   - Android: Material 3 with the wallpaper's dynamic colour, or One UI on
//     Samsung devices; the scale from the DPI (dp).
//   - Linux and the BSDs: the desktop's own theme. GTK 3 themes are read
//     from their CSS, GTK 2 themes from their gtkrc (with Go ports of the
//     common theme engines: Clearlooks, Murrine, Industrial, Mist, ThinIce,
//     Crux, Redmond95), KDE and LXQt get Breeze or Fusion in the desktop's
//     colour scheme, COSMIC is read from its theme files; otherwise
//     Adwaita. The desktop's interface font is used.
//   - Web: the browser's own form controls (Chromium, Firefox or Safari),
//     from the user agent, and the page's dark mode. Browsers don't let
//     WebAssembly use installed fonts, so text is in Godot's built-in font.
//
// Sizes are the platform's: each look's padding, radii and indicator sizes
// were measured from the real controls (WinUI, AppKit, UIKit, Material
// Components, browsers, GTK and Qt) and are multiplied by the display's
// scale.
//
// Each platform only compiles its own looks.
package nativeui

import (
	"image"

	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Image"
	"graphics.gd/classdb/ImageTexture"
	"graphics.gd/classdb/StyleBoxTexture"
	"graphics.gd/nativeui/internal/gtkcss"
	"math"
	"strings"

	"graphics.gd/classdb/Font"
	"graphics.gd/classdb/StyleBox"
	"graphics.gd/classdb/StyleBoxEmpty"
	"graphics.gd/classdb/StyleBoxFlat"
	"graphics.gd/classdb/StyleBoxLine"
	"graphics.gd/classdb/SystemFont"
	"graphics.gd/classdb/Texture2D"
	"graphics.gd/classdb/Theme"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector2"
)

// Look selects which desktop's widgets to imitate.
type lookID int

const (
	lookAuto    lookID = iota // pick from the operating system
	lookMacOS                 // macOS Sonoma/Sequoia (AppKit)
	lookWindows               // Windows 11 (Fluent / WinUI 3)
	lookAdwaita               // GNOME (libadwaita)
	lookAndroid               // Android (Material 3)
	lookIOS                   // iOS / iPadOS (UIKit)
	lookGTK                   // a GTK 3 theme, read from options.GTKTheme (Linux and the BSDs)
	lookGTK2                  // a GTK 2 theme (gtk-2.0/gtkrc), read from options.GTKTheme
	lookBreeze                // KDE Plasma (Breeze), coloured by options.Palette
	lookFusion                // Qt's Fusion (LXQt and Qt outside KDE), coloured by options.Palette
	lookCOSMIC                // System76's COSMIC desktop, from its theme files
	lookWeb                   // the browser's default form controls (options.Browser picks the engine)
	lookOneUI                 // Samsung's One UI (Galaxy phones and tablets)
)

func (l lookID) String() string {
	switch l {
	case lookMacOS:
		return "macOS"
	case lookWindows:
		return "Windows 11"
	case lookAdwaita:
		return "Adwaita"
	case lookAndroid:
		return "Android (Material 3)"
	case lookIOS:
		return "iOS"
	case lookGTK:
		return "GTK"
	case lookGTK2:
		return "GTK 2"
	case lookBreeze:
		return "Breeze"
	case lookFusion:
		return "Fusion"
	case lookCOSMIC:
		return "COSMIC"
	case lookWeb:
		return "Web"
	case lookOneUI:
		return "One UI"
	}
	return "Auto"
}

// Theme type variations that New adds, for use with
// Control.SetThemeTypeVariation.
const (
	variationAccentButton = "AccentButton" // the default/suggested action (blue button)
	variationHeaderLabel  = "HeaderLabel"  // bold group heading
	variationDimLabel     = "DimLabel"     // secondary text
	variationCard         = "Card"         // PanelContainer drawn as a grouped box / card
)

type builder struct {
	theme Theme.Instance
	s     spec
	scale float64
}

func (b *builder) px(v float64) float64 { return math.Round(v * b.scale) }

// gtkTheme builds the GTK look, where it is compiled in.
var gtkTheme, gtk2Theme func(options) (Theme.Instance, error)

// New returns a theme with the native look of the system, as it is now:
// for a Theme of your own, where Use would make it Godot's default theme.
// It doesn't follow later changes to the system's theme.
func New() Theme.Instance { return newTheme(detect()) }

// newTheme builds a theme for the given options.
func newTheme(opts options) Theme.Instance {
	if opts.Look == lookAuto {
		opts.Look = platformLook(opts.Version)
	}
	if opts.Scale <= 0 {
		opts.Scale = 1
	}
	if (opts.Look == lookGTK || opts.Look == lookGTK2) && gtkTheme != nil && opts.GTKTheme != "" {
		build := gtkTheme
		if opts.Look == lookGTK2 {
			build = gtk2Theme
		}
		if theme, err := build(opts); err == nil {
			return theme
		} else {
			Engine.RaiseWarning("nativeui: " + err.Error())
		}
		opts.Look = platformLook(opts.Version)
	}
	b := &builder{theme: Theme.New(), s: lookSpec(opts), scale: float64(opts.Scale)}
	if opts.Font.Family != "" {
		b.s.fonts = append([]string{opts.Font.Family}, b.s.fonts...)
	}
	if opts.Font.Size > 0 {
		b.s.fontSize = int(math.Round(float64(opts.Font.Size)))
	}
	if opts.textScale > 0 {
		b.s.fontSize = int(math.Round(float64(b.s.fontSize) * float64(opts.textScale)))
	}
	b.build()
	return b.theme
}

// fontNames maps fontconfig's generic aliases to the names Godot resolves.
func fontNames(families ...string) []string {
	var names []string
	for _, f := range families {
		switch strings.ToLower(f) {
		case "sans", "sans serif", "sans-serif":
			f = "sans-serif"
		case "serif":
			f = "serif"
		case "monospace", "mono":
			f = "monospace"
		}
		names = append(names, f)
	}
	return names
}

func (b *builder) font(weight int) Font.Instance {
	f := SystemFont.New()
	f.SetFontNames(fontNames(b.s.fonts...))
	f.SetFontWeight(weight)
	return f.AsFont()
}

func (b *builder) style(x box) StyleBox.Instance {
	if x.gradN > 0 || x.borderBottom.A > 0 {
		return b.gradientBox(x)
	}
	if x.glass {
		return newGlassBox(x, b.scale)
	}
	var sb StyleBox.Instance
	switch {
	case x.empty || (x.bg.A == 0 && x.border.A == 0 && x.shadow.A == 0):
		sb = StyleBoxEmpty.New().AsStyleBox()
	case x.line:
		l := StyleBoxLine.New()
		l.SetColor(x.border)
		l.SetThickness(int(math.Max(1, b.px(float64(x.lineThick)))))
		sb = l.AsStyleBox()
	default:
		f := StyleBoxFlat.New()
		f.SetBgColor(x.bg)
		f.SetDrawCenter(!x.noCenter && x.bg.A > 0)
		f.SetBorderColor(x.border)
		f.SetBorderWidthLeft(int(b.px(float64(x.bw[0]))))
		f.SetBorderWidthTop(int(b.px(float64(x.bw[1]))))
		f.SetBorderWidthRight(int(b.px(float64(x.bw[2]))))
		f.SetBorderWidthBottom(int(b.px(float64(x.bw[3]))))
		f.SetCornerRadiusAll(int(b.px(float64(x.radius))))
		if x.squareBottom {
			f.SetCornerRadiusBottomLeft(0)
			f.SetCornerRadiusBottomRight(0)
		}
		f.SetCornerDetail(12)
		f.SetAntiAliasing(true)
		f.SetAntiAliasingSize(0.75)
		if x.shadow.A > 0 {
			f.SetShadowColor(x.shadow)
			f.SetShadowSize(int(b.px(float64(x.shadowSize))))
			f.SetShadowOffset(Vector2.XY{X: 0, Y: Float.X(b.px(x.shadowY*2) / 2)})
		}
		if x.expand > 0 {
			f.SetExpandMarginAll(Float.X(b.px(x.expand)))
		}
		if x.expandTop > 0 {
			f.SetExpandMarginTop(Float.X(b.px(x.expand + x.expandTop)))
		}
		sb = f.AsStyleBox()
	}
	if x.setPad {
		sb.SetContentMarginLeft(Float.X(b.px(x.pad[0])))
		sb.SetContentMarginTop(Float.X(b.px(x.pad[1])))
		sb.SetContentMarginRight(Float.X(b.px(x.pad[2])))
		sb.SetContentMarginBottom(Float.X(b.px(x.pad[3])))
	}
	return sb
}

func (b *builder) styles(themeType string, boxes map[string]box) {
	for name, x := range boxes {
		b.theme.SetStylebox(name, themeType, b.style(x))
	}
}

func (b *builder) colors(themeType string, c Color.RGBA, names ...string) {
	for _, name := range names {
		b.theme.SetColor(name, themeType, c)
	}
}

func (b *builder) icons(themeType string, tex Texture2D.Instance, names ...string) {
	for _, name := range names {
		b.theme.SetIcon(name, themeType, tex)
	}
}

func (b *builder) constant(themeType, name string, v float64) {
	b.theme.SetConstant(name, themeType, int(b.px(v)))
}

var buttonFontColors = []string{"font_color", "font_hover_color", "font_focus_color", "font_hover_pressed_color", "font_pressed_color"}

func (b *builder) build() {
	s, t := b.s, b.theme
	t.SetDefaultBaseScale(Float.X(b.scale))
	t.SetDefaultFont(b.font(400))
	t.SetDefaultFontSize(int(b.px(float64(s.fontSize))))
	bold := b.font(s.boldWeight)
	focus := b.style(s.focus)
	if s.noFocusRing {
		focus = StyleBoxEmpty.New().AsStyleBox()
	}
	empty := StyleBoxEmpty.New().AsStyleBox()

	// Window background.
	b.styles("Panel", map[string]box{"panel": {bg: s.window}})
	b.styles("PanelContainer", map[string]box{"panel": {bg: s.window}})
	t.SetTypeVariation(variationCard, "PanelContainer")
	b.styles(variationCard, map[string]box{"panel": s.card})

	// Labels.
	b.colors("Label", s.text, "font_color")
	t.SetTypeVariation(variationHeaderLabel, "Label")
	t.SetFont("font", variationHeaderLabel, bold)
	if s.headerColor.A > 0 {
		b.colors(variationHeaderLabel, s.headerColor, "font_color")
	}
	if s.headerFontSize > 0 {
		t.SetFontSize("font_size", variationHeaderLabel, int(b.px(float64(s.headerFontSize))))
	}
	if s.entryFontSize > 0 {
		for _, typ := range []string{"LineEdit", "TextEdit"} {
			t.SetFontSize("font_size", typ, int(b.px(float64(s.entryFontSize))))
		}
	}
	t.SetTypeVariation(variationDimLabel, "Label")
	b.colors(variationDimLabel, s.dim, "font_color")
	if len(s.labelFonts) > 0 {
		// Text outside controls is in the document's font (browsers'
		// serif), check box and radio labels included.
		label := func(weight int) Font.Instance {
			f := SystemFont.New()
			f.SetFontNames(fontNames(s.labelFonts...))
			f.SetFontWeight(weight)
			return f.AsFont()
		}
		size := int(b.px(float64(s.labelSize)))
		for _, typ := range []string{"Label", "CheckBox", "CheckButton"} {
			t.SetFont("font", typ, label(400))
			t.SetFontSize("font_size", typ, size)
		}
		t.SetFont("font", variationHeaderLabel, label(700))
	}

	// Push buttons.
	b.styles("Button", map[string]box{
		"normal": s.button, "hover": s.buttonHover, "pressed": s.buttonPressed,
		"hover_pressed": s.buttonPressed, "disabled": s.buttonDisabled,
	})
	t.SetStylebox("focus", "Button", focus)
	b.colors("Button", s.buttonText, buttonFontColors...)
	b.colors("Button", s.disabled, "font_disabled_color")
	if s.buttonBold {
		t.SetFont("font", "Button", bold)
	}
	t.SetTypeVariation(variationAccentButton, "Button")
	b.styles(variationAccentButton, map[string]box{
		"normal": s.accentButton, "hover": s.accentHover, "pressed": s.accentPressed, "hover_pressed": s.accentPressed,
	})
	accentText := s.onAccent
	if s.accentText.A > 0 {
		accentText = s.accentText
	}
	b.colors(variationAccentButton, accentText, buttonFontColors...)

	// Pop-up buttons.
	b.styles("OptionButton", map[string]box{
		"normal": s.button, "hover": s.buttonHover, "pressed": s.buttonPressed,
		"hover_pressed": s.buttonPressed, "disabled": s.buttonDisabled,
	})
	t.SetStylebox("focus", "OptionButton", focus)
	b.colors("OptionButton", s.buttonText, buttonFontColors...)
	b.colors("OptionButton", s.disabled, "font_disabled_color")
	b.icons("OptionButton", b.arrowIcon(), "arrow")
	b.constant("OptionButton", "arrow_margin", 6)
	if s.buttonBold {
		t.SetFont("font", "OptionButton", bold)
	}

	// Check boxes, radio buttons and switches.
	for _, typ := range []string{"CheckBox", "CheckButton"} {
		row := empty
		if s.checkRow > 0 {
			p := math.Max(0, (s.checkRow-s.check.size)/2)
			row = b.style(box{empty: true}.padded(0, p, 0, p))
		}
		for _, name := range []string{"normal", "hover", "pressed", "hover_pressed", "disabled"} {
			t.SetStylebox(name, typ, row)
		}
		t.SetStylebox("focus", typ, empty)
		b.colors(typ, s.text, buttonFontColors...)
		b.colors(typ, s.disabled, "font_disabled_color")
		b.constant(typ, "h_separation", 8)
	}
	b.icons("CheckBox", b.checkIcon(false, false, false), "unchecked")
	b.icons("CheckBox", b.checkIcon(true, false, false), "checked")
	b.icons("CheckBox", b.checkIcon(false, false, true), "unchecked_disabled")
	b.icons("CheckBox", b.checkIcon(true, false, true), "checked_disabled")
	b.icons("CheckBox", b.checkIcon(false, true, false), "radio_unchecked")
	b.icons("CheckBox", b.checkIcon(true, true, false), "radio_checked")
	b.icons("CheckBox", b.checkIcon(false, true, true), "radio_unchecked_disabled")
	b.icons("CheckBox", b.checkIcon(true, true, true), "radio_checked_disabled")
	if s.switchAsCheck { // no switch control: switches are check boxes
		b.icons("CheckButton", b.checkIcon(false, false, false), "unchecked", "unchecked_mirrored")
		b.icons("CheckButton", b.checkIcon(true, false, false), "checked", "checked_mirrored")
		b.icons("CheckButton", b.checkIcon(false, false, true), "unchecked_disabled", "unchecked_disabled_mirrored")
		b.icons("CheckButton", b.checkIcon(true, false, true), "checked_disabled", "checked_disabled_mirrored")
	} else {
		b.icons("CheckButton", b.switchIcon(false, false), "unchecked", "unchecked_mirrored")
		b.icons("CheckButton", b.switchIcon(true, false), "checked", "checked_mirrored")
		b.icons("CheckButton", b.switchIcon(false, true), "unchecked_disabled", "unchecked_disabled_mirrored")
		b.icons("CheckButton", b.switchIcon(true, true), "checked_disabled", "checked_disabled_mirrored")
	}

	// Text fields.
	selection := s.accent
	selection.A = 0.3
	for _, typ := range []string{"LineEdit", "TextEdit"} {
		b.styles(typ, map[string]box{"normal": s.entry, "read_only": s.entry})
		t.SetStylebox("focus", typ, b.style(s.entryFocus))
		b.colors(typ, s.text, "font_color", "caret_color", "font_selected_color")
		b.colors(typ, s.dim, "font_placeholder_color")
		b.colors(typ, s.disabled, "font_readonly_color", "font_uneditable_color")
		b.colors(typ, selection, "selection_color")
		b.colors(typ, Color.RGBA{}, "background_color")
	}
	if s.entryFocus == s.entry && !s.noFocusRing {
		// Looks without a distinct focused field draw the focus ring around it.
		t.SetStylebox("focus", "LineEdit", b.ringed(s.entry))
		t.SetStylebox("focus", "TextEdit", b.ringed(s.entry))
	}

	// Spin boxes.
	updown := b.updownIcon()
	b.icons("SpinBox", updown, "updown")
	for _, name := range []string{"up", "up_hover", "up_pressed", "up_disabled"} {
		b.theme.SetIcon(name, "SpinBox", b.chevronIcon(true))
	}
	for _, name := range []string{"down", "down_hover", "down_pressed", "down_disabled"} {
		b.theme.SetIcon(name, "SpinBox", b.chevronIcon(false))
	}
	for _, name := range []string{"up_background", "up_background_hovered", "up_background_pressed", "up_background_disabled",
		"down_background", "down_background_hovered", "down_background_pressed", "down_background_disabled",
		"field_and_buttons_separator", "up_down_buttons_separator"} {
		t.SetStylebox(name, "SpinBox", empty)
	}
	b.colors("SpinBox", s.dim, "up_icon_modulate", "up_hover_icon_modulate", "up_pressed_icon_modulate",
		"down_icon_modulate", "down_hover_icon_modulate", "down_pressed_icon_modulate")

	// Sliders and progress bars.
	b.styles("HSlider", map[string]box{"slider": s.sliderTrack, "grabber_area": s.sliderFill, "grabber_area_highlight": s.sliderFill})
	b.icons("HSlider", b.thumbIcon(false), "grabber", "grabber_highlight")
	b.icons("HSlider", b.thumbIcon(true), "grabber_disabled")
	b.constant("HSlider", "center_grabber", 0)
	b.styles("ProgressBar", map[string]box{"background": s.progressTrack, "fill": s.progressFill})
	b.colors("ProgressBar", s.text, "font_color")

	// Separators.
	sep := StyleBoxLine.New()
	sep.SetColor(s.separator)
	sep.SetThickness(int(math.Max(1, b.px(1))))
	t.SetStylebox("separator", "HSeparator", sep.AsStyleBox())
	b.constant("HSeparator", "separation", 12)

	// Tabs.
	for _, typ := range []string{"TabContainer", "TabBar"} {
		b.styles(typ, map[string]box{
			"tab_selected": s.tabSelected, "tab_unselected": s.tabUnselected, "tab_hovered": s.tabHovered,
			"tab_disabled": s.tabUnselected, "tabbar_background": {},
		})
		t.SetStylebox("tab_focus", typ, empty)
		b.colors(typ, s.tabSelectedText, "font_selected_color")
		b.colors(typ, s.text, "font_hovered_color")
		b.colors(typ, s.dim, "font_unselected_color")
		b.colors(typ, s.disabled, "font_disabled_color")
		b.constant(typ, "side_margin", 0)
		if s.tabFontSize > 0 {
			t.SetFontSize("font_size", typ, int(b.px(float64(s.tabFontSize))))
		}
	}
	panel := s.card
	if s.tabPanel != nil {
		panel = *s.tabPanel
	}
	b.styles("TabContainer", map[string]box{"panel": panel})
	tabSep := s.tabSeparation
	if tabSep == 0 {
		tabSep = 2
	}
	b.constant("TabContainer", "tab_separation", tabSep)
	b.constant("TabBar", "h_separation", tabSep)

	// Menus.
	b.styles("PopupMenu", map[string]box{"panel": s.popup, "hover": s.popupHover})
	sepMenu := StyleBoxLine.New()
	sepMenu.SetColor(s.separator)
	sepMenu.SetGrowBegin(-Float.X(b.px(6)))
	sepMenu.SetGrowEnd(-Float.X(b.px(6)))
	t.SetStylebox("separator", "PopupMenu", sepMenu.AsStyleBox())
	b.colors("PopupMenu", s.text, "font_color", "font_accelerator_color")
	b.colors("PopupMenu", s.popupHoverText, "font_hover_color")
	b.colors("PopupMenu", s.disabled, "font_disabled_color")
	b.constant("PopupMenu", "v_separation", 6)
	b.constant("PopupMenu", "item_start_padding", 10)
	b.constant("PopupMenu", "item_end_padding", 10)
	b.icons("PopupMenu", b.menuCheckIcon(), "checked", "radio_checked")
	b.icons("PopupMenu", b.blankIcon(), "unchecked", "radio_unchecked")
	t.SetStylebox("embedded_border", "Window", empty)
	t.SetStylebox("embedded_unfocused_border", "Window", empty)
}

// gradientBox draws a box with a vertical gradient fill through the CSS box
// renderer into a 9-slice texture (StyleBoxFlat has no gradients).
func (b *builder) gradientBox(x box) StyleBox.Instance {
	k := b.scale
	toCSS := func(c Color.RGBA) gtkcss.RGBA {
		return gtkcss.RGBA{R: float64(c.R), G: float64(c.G), B: float64(c.B), A: float64(c.A)}
	}
	var st gtkcss.Style
	img := gtkcss.Image{Kind: "linear-gradient", Angle: 180}
	for i, c := range x.grad[:x.gradN] {
		img.Stops = append(img.Stops, gtkcss.Stop{Color: toCSS(c), Pos: float64(i) / float64(max(1, x.gradN-1))})
	}
	if x.gradN > 0 {
		st.BackgroundImage = []gtkcss.Image{img}
	} else {
		st.BackgroundColor = toCSS(x.bg)
	}
	for i, side := range [4]int{1, 2, 3, 0} { // gtkcss: top, right, bottom, left; box: left, top, right, bottom
		st.BorderWidth[i] = math.Round(float64(x.bw[side]) * k)
		st.BorderColor[i] = toCSS(x.border)
		st.BorderStyle[i] = "solid"
	}
	if x.borderBottom.A > 0 {
		st.BorderColor[2] = toCSS(x.borderBottom)
	}
	for i := range st.Radius {
		st.Radius[i] = float64(x.radius) * k
	}
	if x.squareBottom {
		st.Radius[2], st.Radius[3] = 0, 0
	}
	w, h := int(math.Round(48*k)), int(math.Round(28*k))
	tex := StyleBoxTexture.New()
	tex.SetTexture(textureOf(gtkcss.Render(st, w, h)))
	m := math.Ceil(float64(x.radius)*k) + math.Ceil(float64(max(x.bw[0], x.bw[1], x.bw[2], x.bw[3]))*k) + 1
	tex.SetTextureMarginAll(Float.X(m))
	sb := tex.AsStyleBox()
	if x.setPad {
		sb.SetContentMarginLeft(Float.X(b.px(x.pad[0])))
		sb.SetContentMarginTop(Float.X(b.px(x.pad[1])))
		sb.SetContentMarginRight(Float.X(b.px(x.pad[2])))
		sb.SetContentMarginBottom(Float.X(b.px(x.pad[3])))
	}
	return sb
}

func textureOf(img *image.NRGBA) Texture2D.Instance {
	b := img.Bounds()
	return ImageTexture.CreateFromImage(Image.CreateFromData(b.Dx(), b.Dy(), false, Image.FormatRgba8, img.Pix)).AsTexture2D()
}

// ringed is a text field drawn with its focus ring around it: the ring is the
// border, drawn in the margin the box grows by.
func (b *builder) ringed(entry box) StyleBox.Instance {
	ring := b.s.focus
	w := ring.bw[0]
	entry.border = ring.border
	entry.bw = borderAll(w)
	entry.expand = float64(w)
	entry.radius += w
	return b.style(entry)
}
