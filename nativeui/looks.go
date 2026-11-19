package nativeui

import (
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
)

// box describes a StyleBoxFlat in unscaled pixels.
type box struct {
	grad         [4]Color.RGBA // a vertical gradient fill, gradN evenly spaced stops top to bottom (instead of bg)
	gradN        int
	bg, border   Color.RGBA
	bw           [4]int // left, top, right, bottom
	radius       int
	shadow       Color.RGBA
	shadowSize   int
	shadowY      float64
	pad          [4]float64 // left, top, right, bottom content margins
	expand       float64
	expandTop    float64 // extra expand margin above (macOS tab views' boxes reach up behind the tabs)
	noCenter     bool
	line         bool // draw as a 1px horizontal StyleBoxLine using border
	lineThick    int
	empty        bool
	setPad       bool
	squareBottom bool       // only round the top corners (Material filled text fields)
	glass        bool       // Liquid Glass (glass.go), bg is its tint
	borderBottom Color.RGBA // a different colour for the bottom border (WinUI's darker bottom edge)
}

func (b box) padded(l, t, r, bot float64) box {
	b.pad = [4]float64{l, t, r, bot}
	b.setPad = true
	return b
}

func (b box) with(bg Color.RGBA) box { b.bg = bg; return b }

// gradient gives a box a vertical gradient fill through the given stops.
func (b box) gradient(stops ...Color.RGBA) box {
	b.gradN = copy(b.grad[:], stops)
	return b
}

func borderAll(w int) [4]int { return [4]int{w, w, w, w} }

type checkSpec struct {
	size, radius, borderW, markW float64
	bg, border                   Color.RGBA
	radioDot                     float64    // radius of the inner dot, as a fraction of size
	radioRing                    bool       // checked radios are an accent ring + dot (Material), not a filled disc
	radioGap                     Color.RGBA // with radioRing: the fill between ring and dot (zero: none)
	// optional checked colours (zero: accent fill, accent border, onAccent mark)
	onFill, onBorder, mark Color.RGBA
}

type switchSpec struct {
	w, h              float64
	off, offBorder    Color.RGBA
	on                Color.RGBA // zero: accent
	offBorderW        float64
	knobOff, knobOn   Color.RGBA
	knobOffR, knobOnR float64
	knobW             float64 // a capsule knob this wide (macOS 26+), zero for a round one
	knobShadow        bool
}

type thumbSpec struct {
	size    float64 // height (and width, unless w is set)
	w       float64 // width of a capsule handle (macOS 26+), zero for size
	fill    Color.RGBA
	border  Color.RGBA
	shadow  bool
	dot     float64    // radius of an inner accent dot (Windows), 0 for none
	radius  float64    // corner radius of a square handle (Fusion); 0 for a round one
	top     Color.RGBA // with radius: gradient from top to fill
	gapW    float64    // a bar handle (Material 3, 2024): the gap either side cut out of the track
	borderW float64    // ring width (zero: 1)
	gap     Color.RGBA // the gap's colour (the background)
}

// spec is everything that differs between looks.
type spec struct {
	fonts      []string
	fontSize   int
	boldWeight int
	buttonBold bool

	window, text, dim, disabled, accent, onAccent, separator Color.RGBA

	button, buttonHover, buttonPressed, buttonDisabled box
	accentButton, accentHover, accentPressed           box
	entry, entryFocus                                  box
	focus                                              box
	popup, popupHover                                  box
	popupHoverText                                     Color.RGBA
	card                                               box
	tabSelected, tabUnselected, tabHovered             box
	progressTrack, progressFill                        box
	sliderTrack, sliderFill                            box

	check  checkSpec
	toggle switchSpec
	thumb  thumbSpec

	arrowBadge  bool // macOS pop-up buttons draw their arrows in an accent badge
	arrowUpDown bool // macOS 26+ pop-up buttons: plain up/down chevrons

	// Optional overrides, zero means s.text (or s.dim for arrowColor).
	buttonText, tabSelectedText, arrowColor Color.RGBA
	accentText                              Color.RGBA // default/accent button text (zero: onAccent)
	noFocusRing                             bool       // touch looks have no keyboard focus ring
	tabSeparation                           float64    // between tabs (zero: 2)
	checkRow                                float64    // minimum height of check box and switch rows (zero: the indicator's)
	tabPanel                                *box       // the tab container's page (nil: card)
	labelFonts                              []string   // text outside controls, when it differs (browsers' document font)
	labelSize                               int
	switchAsCheck                           bool       // the platform has no switch: CheckButton draws a check box
	tabFontSize                             int        // tab labels, when smaller than the text (iOS segmented controls)
	headerColor                             Color.RGBA // HeaderLabel (zero: text)
	headerFontSize                          int        // HeaderLabel (zero: the font size)
	entryFontSize                           int        // text fields, when bigger than the text (One UI)
}

func rgba(r, g, b, a Float.X) Color.RGBA { return Color.RGBA{R: r / 255, G: g / 255, B: b / 255, A: a} }
func black(a Float.X) Color.RGBA         { return Color.RGBA{A: a} }
func white(a Float.X) Color.RGBA         { return Color.RGBA{R: 1, G: 1, B: 1, A: a} }
func hex(s string) Color.RGBA            { return Color.Hex(s) }

// blend mixes a towards b by t, opaque.
func blend(a, b Color.RGBA, t float32) Color.RGBA {
	return Color.RGBA{R: a.R + (b.R-a.R)*t, G: a.G + (b.G-a.G)*t, B: a.B + (b.B-a.B)*t, A: 1}
}

func withAlpha(c Color.RGBA, a float32) Color.RGBA { c.A = a; return c }

func shade(c Color.RGBA, f Float.X) Color.RGBA {
	return Color.RGBA{R: c.R * f, G: c.G * f, B: c.B * f, A: c.A}
}

// looks holds the looks compiled into this build: each platform only gets
// its own (see the build constraints on the look_*.go files), the
// nativeui_all build tag includes all of them.
var looks = map[Look]func(Options) spec{}

func register(look Look, fn func(Options) spec) { looks[look] = fn }

// Available reports the looks compiled into this build.
func Available() []Look {
	var list []Look
	for look := MacOS; look <= OneUI; look++ {
		if looks[look] != nil {
			list = append(list, look)
		}
	}
	return list
}

func lookSpec(opts Options) spec {
	fn := looks[opts.Look]
	if fn == nil {
		fn = looks[platformLook(opts.Version)]
	}
	if fn == nil {
		for _, l := range Available() {
			fn = looks[l]
			break
		}
	}
	s := fn(opts)
	if opts.Accent.A > 0 && !(opts.HighContrast && opts.Palette != nil) { // contrast themes bring their own highlight
		s.recolor(opts.Accent)
	}
	if s.buttonText.A == 0 {
		s.buttonText = s.text
	}
	if s.tabSelectedText.A == 0 {
		s.tabSelectedText = s.text
	}
	if s.arrowColor.A == 0 {
		s.arrowColor = s.dim
	}
	return s
}

// recolor replaces the look's default accent with the system one.
func (s *spec) recolor(accent Color.RGBA) {
	old := s.accent
	fix := func(c *Color.RGBA) {
		if c.R == old.R && c.G == old.G && c.B == old.B {
			*c = Color.RGBA{R: accent.R, G: accent.G, B: accent.B, A: c.A}
		}
	}
	for _, b := range []*box{&s.accentButton, &s.accentHover, &s.accentPressed, &s.entryFocus, &s.focus,
		&s.popupHover, &s.tabSelected, &s.progressFill, &s.sliderFill} {
		fix(&b.bg)
		fix(&b.border)
		fix(&b.borderBottom)
	}
	if s.accentHover.bg == accent { // only where the look fills its accent buttons with the accent
		s.accentHover.bg = shade(accent, 1.08)
		s.accentPressed.bg = shade(accent, 0.85)
	}
	s.accent = accent
}
