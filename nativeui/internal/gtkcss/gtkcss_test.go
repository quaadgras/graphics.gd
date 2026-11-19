package gtkcss

import (
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

const sample = `
@define-color bg_color #c0c0c0;
@define-color border_dark #000000;
@define-color border_bright #ffffff;
@define-color shadowed shade(@bg_color, 0.5);
button {
  padding: 2px 4px;
  border: 1px solid;
  border-top-color: @border_bright;
  border-bottom-color: @border_dark;
  background-color: @bg_color;
  box-shadow: inset -1px -1px #808080, inset 1px 1px #dfdfdf; }
button:active, button:checked { box-shadow: inset 1px 1px #808080; }
button.suggested-action { background-color: alpha(@border_dark, 0.5); }
.dialog-action-area button { border: 3px solid red; }
check:not(:checked) { -gtk-icon-source: url("assets/unchecked.png"); }
check:checked { -gtk-icon-source: url("assets/checked.png"); }
label { color: mix(@border_dark, @border_bright, 0.25); }
`

func TestCompute(t *testing.T) {
	sheet, err := Parse(sample, "/theme")
	if err != nil {
		t.Fatal(err)
	}
	st := sheet.Compute(Path("window button"))
	if st.BorderWidth != [4]float64{1, 1, 1, 1} {
		t.Fatalf("border widths %v", st.BorderWidth)
	}
	if st.BorderColor[0] != (RGBA{1, 1, 1, 1}) || st.BorderColor[2] != (RGBA{0, 0, 0, 1}) {
		t.Fatalf("border colors %v", st.BorderColor)
	}
	if st.Padding != [4]float64{2, 4, 2, 4} {
		t.Fatalf("padding %v", st.Padding)
	}
	if len(st.Shadows) != 2 || !st.Shadows[0].Inset || st.Shadows[0].X != -1 {
		t.Fatalf("shadows %+v", st.Shadows)
	}
	if st := sheet.Compute(Path("window button:active")); len(st.Shadows) != 1 {
		t.Fatalf("active shadows %+v", st.Shadows)
	}
	if st := sheet.Compute(Path("window button.suggested-action")); st.BackgroundColor.A != 0.5 {
		t.Fatalf("suggested %+v", st.BackgroundColor)
	}
	if st := sheet.Compute(Path("dialog box.dialog-action-area button")); st.BorderWidth[0] != 3 || st.BorderColor[0] != (RGBA{1, 0, 0, 1}) {
		t.Fatalf("dialog button %+v", st)
	}
	if st := sheet.Compute(Path("checkbutton check")); st.IconSource.Path != "/theme/assets/unchecked.png" {
		t.Fatalf("unchecked icon %+v", st.IconSource)
	}
	if st := sheet.Compute(Path("checkbutton check:checked")); st.IconSource.Path != "/theme/assets/checked.png" {
		t.Fatalf("checked icon %+v", st.IconSource)
	}
	if st := sheet.Compute(Path("window label")); st.Color.R != 0.25 {
		t.Fatalf("mixed color %+v", st.Color)
	}
	if c, _ := sheet.Color("@shadowed", RGBA{}); c.R < 0.37 || c.R > 0.38 {
		t.Fatalf("shade %+v", c)
	}
}

// TestTheme loads a real theme named by NATIVEUI_GTK_THEME (a gtk-3.0
// directory), for example Chicago95.
func TestTheme(t *testing.T) {
	dir := os.Getenv("NATIVEUI_GTK_THEME")
	if dir == "" {
		t.Skip("NATIVEUI_GTK_THEME not set")
	}
	sheet, err := Load(filepath.Join(dir, "gtk.css"))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%d rules, %d colors", len(sheet.Rules), len(sheet.Colors))
	for _, p := range []string{"window.background button", "window.background button:active", "window.background entry",
		"window.background checkbutton check:checked", "window.background progressbar trough progress",
		"window.background notebook header.top tabs tab:checked", "window.background scale.horizontal contents trough slider"} {
		st := sheet.Compute(Path(p))
		t.Logf("%s:\n  bg %v border %v %v radius %v pad %v min %vx%v shadows %+v icon %+v bimg %+v",
			p, st.BackgroundColor, st.BorderWidth, st.BorderColor, st.Radius, st.Padding, st.MinWidth, st.MinHeight, st.Shadows, st.IconSource, st.BorderImage)
	}
}

// TestRender draws a few nodes of the theme named by NATIVEUI_GTK_THEME into
// NATIVEUI_RENDER_OUT (a directory) for inspection.
func TestRender(t *testing.T) {
	dir, out := os.Getenv("NATIVEUI_GTK_THEME"), os.Getenv("NATIVEUI_RENDER_OUT")
	if dir == "" || out == "" {
		t.Skip("NATIVEUI_GTK_THEME or NATIVEUI_RENDER_OUT not set")
	}
	sheet, err := Load(filepath.Join(dir, "gtk.css"))
	if err != nil {
		t.Fatal(err)
	}
	for name, node := range map[string]struct {
		path string
		w, h int
	}{
		"button":        {"window.background button", 75, 23},
		"button-active": {"window.background button:active", 75, 23},
		"entry":         {"window.background entry", 120, 22},
		"tab-checked":   {"window.background notebook header.top tabs tab:checked", 60, 20},
		"trough":        {"window.background progressbar trough", 120, 20},
		"switch":        {"window.background switch", 44, 20},
	} {
		img := Render(sheet.Compute(Path(node.path)), node.w, node.h)
		f, err := os.Create(filepath.Join(out, name+".png"))
		if err != nil {
			t.Fatal(err)
		}
		if err := png.Encode(f, img); err != nil {
			t.Fatal(err)
		}
		f.Close()
	}
}
