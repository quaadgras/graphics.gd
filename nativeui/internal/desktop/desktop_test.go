package desktop

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func write(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

// fakeEnv is an environment rooted at a temporary home with no system dirs.
func fakeEnv(t *testing.T, vars map[string]string, gsettings map[string]string) (Env, string) {
	home := t.TempDir()
	system := filepath.Join(home, "system")
	env := map[string]string{
		"HOME":            home,
		"XDG_CONFIG_DIRS": filepath.Join(system, "xdg"),
		"XDG_DATA_DIRS":   filepath.Join(system, "share"),
	}
	for k, v := range vars {
		env[k] = v
	}
	return Env{
		Getenv: func(k string) string { return env[k] },
		Run: func(name string, args ...string) (string, error) {
			if v, ok := gsettings[args[len(args)-1]]; ok {
				return "'" + v + "'\n", nil
			}
			return "", errors.New("no such key")
		},
	}, home
}

func theme(t *testing.T, dir string) {
	write(t, filepath.Join(dir, "gtk-3.0", "gtk.css"), "window { background-color: #c0c0c0; }")
}

func TestXFCE(t *testing.T) {
	env, home := fakeEnv(t, map[string]string{"XDG_CURRENT_DESKTOP": "XFCE"}, nil)
	write(t, filepath.Join(home, ".config/xfce4/xfconf/xfce-perchannel-xml/xsettings.xml"), `<?xml version="1.0" encoding="UTF-8"?>
<channel name="xsettings" version="1.0">
  <property name="Net" type="empty">
    <property name="ThemeName" type="string" value="Chicago95"/>
    <property name="IconThemeName" type="string" value="Chicago95"/>
  </property>
  <property name="Xft" type="empty">
    <property name="DPI" type="int" value="120"/>
  </property>
  <property name="Gtk" type="empty">
    <property name="FontName" type="string" value="Helvetica Bold 9"/>
  </property>
</channel>`)
	theme(t, filepath.Join(home, ".themes", "Chicago95"))
	s := env.Read()
	if s.Desktop != XFCE || s.GTKTheme != "Chicago95" || s.IconTheme != "Chicago95" {
		t.Fatalf("%+v", s)
	}
	if s.ThemeDir != filepath.Join(home, ".themes", "Chicago95") {
		t.Fatalf("theme dir %q", s.ThemeDir)
	}
	if s.Font != (Font{"Helvetica", 700, 9}) || s.DPI != 120 || s.Pixels() != 15 {
		t.Fatalf("font %+v dpi %v", s.Font, s.DPI)
	}
}

func TestLXDE(t *testing.T) {
	env, home := fakeEnv(t, map[string]string{"XDG_CURRENT_DESKTOP": "LXDE", "DESKTOP_SESSION": "LXDE"}, nil)
	write(t, filepath.Join(home, "system/xdg/lxsession/LXDE/desktop.conf"), "[GTK]\nsNet/ThemeName=Clearlooks\nsGtk/FontName=Sans 10\n")
	write(t, filepath.Join(home, ".config/lxsession/LXDE/desktop.conf"), "[Session]\nwindow_manager=openbox-lxde\n[GTK]\nsNet/ThemeName=Redmond\nsGtk/FontName=Tahoma 8\n")
	theme(t, filepath.Join(home, "system/share/themes/Redmond"))
	s := env.Read()
	if s.Desktop != LXDE || s.GTKTheme != "Redmond" || s.Font != (Font{"Tahoma", 400, 8}) {
		t.Fatalf("%+v", s)
	}
	if s.ThemeDir != filepath.Join(home, "system/share/themes/Redmond") {
		t.Fatalf("theme dir %q", s.ThemeDir)
	}
}

func TestGNOME(t *testing.T) {
	env, _ := fakeEnv(t, map[string]string{"XDG_CURRENT_DESKTOP": "ubuntu:GNOME"}, map[string]string{
		"gtk-theme":    "Adwaita",
		"font-name":    "Cantarell 11",
		"color-scheme": "prefer-dark",
	})
	s := env.Read()
	if s.Desktop != GNOME || s.GTKTheme != "Adwaita" || !s.Dark || s.ThemeDir != "" || s.Font.Family != "Cantarell" {
		t.Fatalf("%+v", s)
	}
}

func TestGTKThemeOverride(t *testing.T) {
	env, home := fakeEnv(t, map[string]string{"XDG_CURRENT_DESKTOP": "sway", "GTK_THEME": "adw-gtk3:dark"}, nil)
	write(t, filepath.Join(home, ".config/gtk-3.0/settings.ini"), "[Settings]\ngtk-theme-name=Arc\ngtk-font-name=Noto Sans Medium 10\n")
	theme(t, filepath.Join(home, ".local/share/themes/adw-gtk3"))
	s := env.Read()
	if s.Desktop != Unknown || s.GTKTheme != "adw-gtk3:dark" || !s.Dark || s.Font != (Font{"Noto Sans", 500, 10}) {
		t.Fatalf("%+v", s)
	}
	if s.ThemeDir == "" {
		t.Fatal("theme not found")
	}
}

func TestKDE(t *testing.T) {
	env, home := fakeEnv(t, map[string]string{"XDG_CURRENT_DESKTOP": "KDE"}, nil)
	write(t, filepath.Join(home, ".config/kdeglobals"), "[General]\nfont=Noto Sans,10,-1,5,50,0,0,0,0,0\n\n[Colors:Window]\nBackgroundNormal=32,35,38\nForegroundNormal=252,252,252\n\n[Colors:Selection]\nBackgroundNormal=61,174,233\n")
	s := env.Read()
	if s.Desktop != KDE || s.Font != (Font{"Noto Sans", 400, 10}) {
		t.Fatalf("%+v", s)
	}
	if s.KDE["Window.BackgroundNormal"] != [3]uint8{32, 35, 38} || s.KDE["Selection.BackgroundNormal"] != [3]uint8{61, 174, 233} || !s.Dark {
		t.Fatalf("colours %v dark %v", s.KDE, s.Dark)
	}
}

func TestLXQt(t *testing.T) {
	env, home := fakeEnv(t, map[string]string{"XDG_CURRENT_DESKTOP": "LXQt"}, nil)
	write(t, filepath.Join(home, ".config/lxqt/lxqt.conf"), "[Qt]\nstyle=Fusion\nfont=\"Noto Sans,11,-1,5,400,0,0,0,0,0,0,0,0,0,0,1\"\n\n[Palette]\nwindow_color=#31363b\nhighlight_color=#ff3daee9\n")
	s := env.Read()
	if s.Desktop != LXQt || s.QtStyle != "Fusion" || s.Font.Size != 11 {
		t.Fatalf("%+v", s)
	}
	if s.KDE["Window.BackgroundNormal"] != [3]uint8{0x31, 0x36, 0x3b} || s.KDE["Selection.BackgroundNormal"] != [3]uint8{0x3d, 0xae, 0xe9} {
		t.Fatalf("palette %v", s.KDE)
	}
}

func TestParseFont(t *testing.T) {
	for desc, want := range map[string]Font{
		"DejaVu Sans 10":            {"DejaVu Sans", 400, 10},
		"Ubuntu Bold Italic 11":     {"Ubuntu", 700, 11},
		"Sans":                      {"Sans", 400, 10},
		"Noto Sans, Cantarell 12px": {"Noto Sans Cantarell", 400, 9},
		"MS Sans Serif 8":           {"MS Sans Serif", 400, 8},
	} {
		if got := ParseFont(desc); got != want {
			t.Errorf("%q: got %+v want %+v", desc, got, want)
		}
	}
}
