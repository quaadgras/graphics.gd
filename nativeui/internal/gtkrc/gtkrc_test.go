package gtkrc

import (
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

const sample = `
gtk_color_scheme = "fg_color:#000000\nbg_color:#c0c0c0\nselected_bg_color:#000080"
style "default" {
	GtkButton::default_border = { 1, 1, 1, 1 }
	GtkHScale::slider_length = 11
	fg[NORMAL] = @fg_color
	bg[NORMAL] = @bg_color
	bg[SELECTED] = @selected_bg_color
	bg[PRELIGHT] = shade (1.1, @bg_color)
	base[NORMAL] = "#ffffff"
	xthickness = 2
	ythickness = 2
}
style "button" = "default" {
	xthickness = 3
	engine "pixmap" {
		image {
			function = BOX
			detail = "button"
			state = ACTIVE
			file = "pressed.png"
			border = { 2, 2, 2, 2 }
			stretch = TRUE
		}
	}
}
style "menuitem" { bg[PRELIGHT] = { 0.0, 0.0, 0.5 } }
class "GtkWidget" style "default"
widget_class "*GtkButton" style "button"
widget_class "*<GtkMenuItem>*" style "menuitem"
`

func parse(t *testing.T, src string) *Theme {
	dir := t.TempDir()
	path := filepath.Join(dir, "gtkrc")
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	theme, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	return theme
}

func TestEngineOptionsMerge(t *testing.T) {
	th := parse(t, `style "base" { engine "murrine" { menustyle = 0 progressbarstyle = 0 } }
style "bar" { engine "murrine" { roundness = 3 } }
class "GtkWidget" style "base"
class "GtkProgressBar" style "bar"`)
	st := th.Resolve(W("GtkWindow.GtkProgressBar"))
	if st.EngineParams["menustyle"] != "0" || st.EngineParams["progressbarstyle"] != "0" || st.EngineParams["roundness"] != "3" {
		t.Fatalf("%v", st.EngineParams)
	}
}

func TestColorSchemeSemicolons(t *testing.T) {
	th := parse(t, `gtk_color_scheme = "fg_color:#000;selected_bg_color:#75507b"
style "s" { bg[SELECTED] = @selected_bg_color }
class "GtkWidget" style "s"`)
	if c := th.Resolve(W("GtkWindow")).Bg[Selected]; c.R < 0.45 || c.R > 0.47 {
		t.Fatalf("selected bg %+v", c)
	}
}

func TestResolve(t *testing.T) {
	th := parse(t, sample)
	label := th.Resolve(W("GtkWindow.GtkVBox.GtkLabel"))
	if label.Bg[Normal] != (RGBA{192.0 / 255, 192.0 / 255, 192.0 / 255, 1}) || label.XThickness != 2 {
		t.Fatalf("label %+v", label)
	}
	button := th.Resolve(W("GtkWindow.GtkVBox.GtkButton"))
	if button.XThickness != 3 || button.Engine != "pixmap" || button.Bg[Selected].B < 0.5 {
		t.Fatalf("button %+v", button)
	}
	if v, _ := button.Prop("GtkButton", "default-border"); v != "{1,1,1,1}" {
		t.Fatalf("prop %q", v)
	}
	if _, ok := button.Match(Call{Function: "BOX", Detail: "button", State: Active, Shadow: "IN"}); !ok {
		t.Fatal("pressed image not matched")
	}
	if _, ok := button.Match(Call{Function: "BOX", Detail: "button", State: Normal, Shadow: "OUT"}); ok {
		t.Fatal("normal state should fall back")
	}
	item := th.Resolve(W("GtkWindow.GtkMenu.GtkCheckMenuItem"))
	if item.Bg[Prelight] != (RGBA{0, 0, 0.5, 1}) {
		t.Fatalf("menu item %+v", item.Bg)
	}
}

// TestTheme draws a few calls of the theme named by NATIVEUI_GTK2_THEME (a
// gtkrc file) into NATIVEUI_RENDER_OUT.
func TestTheme(t *testing.T) {
	rc, out := os.Getenv("NATIVEUI_GTK2_THEME"), os.Getenv("NATIVEUI_RENDER_OUT")
	if rc == "" || out == "" {
		t.Skip("NATIVEUI_GTK2_THEME or NATIVEUI_RENDER_OUT not set")
	}
	th, err := Load(rc)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%d styles, %d bindings", len(th.Styles), len(th.bindings))
	for name, d := range map[string]struct {
		widget string
		call   Call
		w, h   int
	}{
		"gtk2-button":  {"GtkWindow.GtkButton", Call{Function: "BOX", Detail: "button", Shadow: "OUT"}, 75, 23},
		"gtk2-pressed": {"GtkWindow.GtkButton", Call{Function: "BOX", Detail: "button", Shadow: "IN", State: Active}, 75, 23},
		"gtk2-check":   {"GtkWindow.GtkCheckButton", Call{Function: "CHECK", Detail: "checkbutton", Shadow: "IN"}, 13, 13},
		"gtk2-radio":   {"GtkWindow.GtkRadioButton", Call{Function: "OPTION", Detail: "radiobutton", Shadow: "IN"}, 13, 13},
		"gtk2-entry":   {"GtkWindow.GtkEntry", Call{Function: "SHADOW", Detail: "entry", Shadow: "IN"}, 120, 22},
		"gtk2-tab":     {"GtkWindow.GtkNotebook", Call{Function: "EXTENSION", Detail: "tab", GapSide: "BOTTOM"}, 60, 22},
		"gtk2-trough":  {"GtkWindow.GtkProgressBar", Call{Function: "BOX", Detail: "trough", Shadow: "IN"}, 120, 16},
		"gtk2-slider":  {"GtkWindow.GtkHScale", Call{Function: "SLIDER", Detail: "hscale", Shadow: "OUT", Orientation: "HORIZONTAL"}, 11, 21},
	} {
		img := th.Resolve(W(d.widget)).Draw(d.call, d.w, d.h)
		f, _ := os.Create(filepath.Join(out, name+".png"))
		png.Encode(f, img)
		f.Close()
	}
}
