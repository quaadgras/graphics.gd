// Package desktop reads the appearance settings of Linux and BSD desktops
// (which desktop is running, its GTK theme, font and color scheme) from the
// places each desktop keeps them. It is pure Go so that it can be tested
// against fake home directories.
package desktop

import (
	"bufio"
	"context"
	"encoding/xml"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Desktop identifies a desktop environment.
type Desktop string

const (
	Unknown  Desktop = ""
	GNOME    Desktop = "GNOME"
	KDE      Desktop = "KDE"
	XFCE     Desktop = "XFCE"
	LXDE     Desktop = "LXDE"
	LXQt     Desktop = "LXQt"
	MATE     Desktop = "MATE"
	Cinnamon Desktop = "Cinnamon"
	Budgie   Desktop = "Budgie"
	Pantheon Desktop = "Pantheon"
	COSMIC   Desktop = "COSMIC"
	Unity    Desktop = "Unity"
	Deepin   Desktop = "Deepin"
)

// Settings are the appearance settings found.
type Settings struct {
	Desktop   Desktop
	GTKTheme  string  // theme name, as the user picked it ("Chicago95", "Adwaita-dark"...)
	ThemeDir  string  // directory of the theme, if it has GTK 3 files (empty for built-in themes)
	ThemeDir2 string  // directory of the theme, if it has GTK 2 files (gtk-2.0/gtkrc)
	Dark      bool    // the theme is a dark variant, or the desktop prefers dark
	Font      Font    // the interface font
	IconTheme string  // freedesktop icon theme name
	DPI       float64 // font DPI (96 by default)
	Sources   []string
	// KDE is the kdeglobals colour scheme, "Window.BackgroundNormal" ->
	// {r, g, b} (0-255), on Plasma.
	KDE map[string][3]uint8
	// QtStyle is the Qt widget style LXQt is set to use ("Fusion", "Breeze"...).
	QtStyle string
}

// Font is a parsed Pango font description, e.g. "DejaVu Sans Bold 10".
type Font struct {
	Family string
	Weight int     // CSS weight: 400 regular, 700 bold
	Size   float64 // points
}

// Env is the environment settings are read from; the zero value is the real one.
type Env struct {
	Getenv func(string) string
	// Run runs a command and returns its standard output (gsettings).
	Run func(name string, args ...string) (string, error)
}

func (e Env) getenv(k string) string {
	if e.Getenv != nil {
		return e.Getenv(k)
	}
	return os.Getenv(k)
}

func (e Env) run(name string, args ...string) (string, error) {
	if e.Run != nil {
		return e.Run(name, args...)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	out, err := exec.CommandContext(ctx, name, args...).Output()
	return string(out), err
}

func (e Env) home() string {
	if h := e.getenv("HOME"); h != "" {
		return h
	}
	h, _ := os.UserHomeDir()
	return h
}

// ConfigHome is $XDG_CONFIG_HOME, or ~/.config.
func (e Env) ConfigHome() string { return e.configHome() }

func (e Env) configHome() string {
	if d := e.getenv("XDG_CONFIG_HOME"); d != "" {
		return d
	}
	return filepath.Join(e.home(), ".config")
}

func (e Env) configDirs() []string {
	dirs := []string{e.configHome()}
	if d := e.getenv("XDG_CONFIG_DIRS"); d != "" {
		dirs = append(dirs, filepath.SplitList(d)...)
	} else {
		dirs = append(dirs, "/etc/xdg")
	}
	return dirs
}

func (e Env) dataDirs() []string {
	home := e.getenv("XDG_DATA_HOME")
	if home == "" {
		home = filepath.Join(e.home(), ".local", "share")
	}
	dirs := []string{home}
	if d := e.getenv("XDG_DATA_DIRS"); d != "" {
		dirs = append(dirs, filepath.SplitList(d)...)
	} else {
		dirs = append(dirs, "/usr/local/share", "/usr/share")
	}
	return dirs
}

// DetectDesktop works out which desktop environment is running.
func (e Env) DetectDesktop() Desktop {
	names := e.getenv("XDG_CURRENT_DESKTOP") + ":" + e.getenv("XDG_SESSION_DESKTOP") + ":" + e.getenv("DESKTOP_SESSION")
	for _, name := range strings.FieldsFunc(strings.ToLower(names), func(r rune) bool { return r == ':' || r == ';' }) {
		switch {
		case name == "cosmic":
			return COSMIC
		case name == "kde" || name == "plasma" || strings.HasPrefix(name, "plasma"):
			return KDE
		case name == "xfce" || strings.HasPrefix(name, "xfce"):
			return XFCE
		case name == "lxqt":
			return LXQt
		case name == "lxde" || name == "lubuntu":
			return LXDE
		case name == "mate":
			return MATE
		case name == "x-cinnamon" || name == "cinnamon":
			return Cinnamon
		case name == "budgie" || strings.HasPrefix(name, "budgie"):
			return Budgie
		case name == "pantheon":
			return Pantheon
		case name == "unity":
			return Unity
		case name == "deepin" || name == "dde":
			return Deepin
		case name == "gnome" || strings.HasPrefix(name, "gnome") || name == "ubuntu":
			return GNOME
		}
	}
	switch {
	case e.getenv("KDE_FULL_SESSION") != "":
		return KDE
	case e.getenv("GNOME_DESKTOP_SESSION_ID") != "":
		return GNOME
	case e.getenv("MATE_DESKTOP_SESSION_ID") != "":
		return MATE
	}
	return Unknown
}

// Read reads the current desktop's appearance settings.
func (e Env) Read() Settings {
	s := Settings{Desktop: e.DetectDesktop(), DPI: 96, Font: Font{Family: "Sans", Weight: 400, Size: 10}}
	// lowest priority first: later sources override earlier ones.
	e.gtkSettingsIni(&s)
	switch s.Desktop {
	case XFCE:
		e.xfconf(&s)
	case LXDE:
		e.lxsession(&s)
	case GNOME, Budgie, Pantheon, Unity, COSMIC:
		e.gsettings(&s, "org.gnome.desktop.interface")
	case Cinnamon:
		e.gsettings(&s, "org.cinnamon.desktop.interface")
	case MATE:
		e.gsettings(&s, "org.mate.interface")
	case KDE:
		e.kdeglobals(&s)
	case LXQt:
		e.lxqt(&s)
	}
	if theme := e.getenv("GTK_THEME"); theme != "" {
		s.GTKTheme = theme
		s.Sources = append(s.Sources, "GTK_THEME")
	}
	name, variant, _ := strings.Cut(s.GTKTheme, ":")
	if variant == "dark" || strings.HasSuffix(strings.ToLower(name), "-dark") {
		s.Dark = true
	}
	s.ThemeDir = e.FindTheme(name, "gtk-3.0/gtk.css")
	s.ThemeDir2 = e.FindTheme(name, "gtk-2.0/gtkrc")
	return s
}

// FindTheme returns the directory of the named theme if it has the given
// file ("gtk-3.0/gtk.css", "gtk-2.0/gtkrc").
func (e Env) FindTheme(name, file string) string {
	if name == "" || strings.ContainsAny(name, "/\\") {
		return ""
	}
	dirs := []string{}
	for _, d := range e.dataDirs() {
		dirs = append(dirs, filepath.Join(d, "themes"))
		if d == e.dataDirs()[0] {
			dirs = append(dirs, filepath.Join(e.home(), ".themes"))
		}
	}
	for _, d := range dirs {
		dir := filepath.Join(d, name)
		if _, err := os.Stat(filepath.Join(dir, filepath.FromSlash(file))); err == nil {
			return dir
		}
	}
	return ""
}

func (e Env) gtkSettingsIni(s *Settings) {
	dirs := e.configDirs()
	for i := len(dirs) - 1; i >= 0; i-- { // system first, user last
		path := filepath.Join(dirs[i], "gtk-3.0", "settings.ini")
		values := readIni(path, "Settings")
		if values == nil {
			continue
		}
		s.Sources = append(s.Sources, path)
		if v := values["gtk-theme-name"]; v != "" {
			s.GTKTheme = v
		}
		if v := values["gtk-font-name"]; v != "" {
			s.Font = ParseFont(v)
		}
		if v := values["gtk-icon-theme-name"]; v != "" {
			s.IconTheme = v
		}
		if v := values["gtk-application-prefer-dark-theme"]; v == "1" || v == "true" {
			s.Dark = true
		}
	}
}

// readIni reads one section of an ini file, nil if the file doesn't exist.
func readIni(path, section string) map[string]string {
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()
	values := map[string]string{}
	in := section == ""
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || line[0] == '#' || line[0] == ';' {
			continue
		}
		if line[0] == '[' {
			in = strings.TrimSuffix(line[1:], "]") == section
			continue
		}
		if !in {
			continue
		}
		k, v, ok := strings.Cut(line, "=")
		if ok {
			values[strings.TrimSpace(k)] = strings.Trim(strings.TrimSpace(v), `"`)
		}
	}
	return values
}

// xfconf reads XFCE's xsettings channel (what xfsettsd serves over XSETTINGS).
func (e Env) xfconf(s *Settings) {
	dirs := e.configDirs()
	for i := len(dirs) - 1; i >= 0; i-- {
		path := filepath.Join(dirs[i], "xfce4", "xfconf", "xfce-perchannel-xml", "xsettings.xml")
		props, err := readXfconf(path)
		if err != nil {
			continue
		}
		s.Sources = append(s.Sources, path)
		if v := props["Net/ThemeName"]; v != "" {
			s.GTKTheme = v
		}
		if v := props["Net/IconThemeName"]; v != "" {
			s.IconTheme = v
		}
		if v := props["Gtk/FontName"]; v != "" {
			s.Font = ParseFont(v)
		}
		if v := props["Xft/DPI"]; v != "" {
			if dpi, err := strconv.ParseFloat(v, 64); err == nil && dpi > 0 {
				s.DPI = dpi
			}
		}
	}
}

type xfconfProperty struct {
	Name     string           `xml:"name,attr"`
	Value    string           `xml:"value,attr"`
	Children []xfconfProperty `xml:"property"`
}

func readXfconf(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var channel struct {
		Properties []xfconfProperty `xml:"property"`
	}
	if err := xml.Unmarshal(data, &channel); err != nil {
		return nil, err
	}
	props := map[string]string{}
	var walk func(prefix string, list []xfconfProperty)
	walk = func(prefix string, list []xfconfProperty) {
		for _, p := range list {
			if len(p.Children) > 0 {
				walk(prefix+p.Name+"/", p.Children)
			} else {
				props[prefix+p.Name] = p.Value
			}
		}
	}
	walk("", channel.Properties)
	return props, nil
}

// lxsession reads LXDE's desktop.conf (what lxsession serves over XSETTINGS).
func (e Env) lxsession(s *Settings) {
	profile := e.getenv("DESKTOP_SESSION")
	if profile == "" || strings.ContainsAny(profile, "/\\") {
		profile = "LXDE"
	}
	dirs := e.configDirs()
	for i := len(dirs) - 1; i >= 0; i-- {
		path := filepath.Join(dirs[i], "lxsession", profile, "desktop.conf")
		values := readIni(path, "GTK")
		if values == nil {
			continue
		}
		s.Sources = append(s.Sources, path)
		if v := values["sNet/ThemeName"]; v != "" {
			s.GTKTheme = v
		}
		if v := values["sNet/IconThemeName"]; v != "" {
			s.IconTheme = v
		}
		if v := values["sGtk/FontName"]; v != "" {
			s.Font = ParseFont(v)
		}
	}
}

// gsettings asks GSettings (dconf) for a desktop's interface settings.
func (e Env) gsettings(s *Settings, schema string) {
	get := func(key string) string {
		out, err := e.run("gsettings", "get", schema, key)
		if err != nil {
			return ""
		}
		return strings.Trim(strings.TrimSpace(out), "'")
	}
	found := false
	if v := get("gtk-theme"); v != "" {
		s.GTKTheme, found = v, true
	}
	if v := get("font-name"); v != "" {
		s.Font, found = ParseFont(v), true
	}
	if v := get("icon-theme"); v != "" {
		s.IconTheme = v
	}
	if schema == "org.gnome.desktop.interface" && get("color-scheme") == "prefer-dark" {
		s.Dark = true
	}
	if found {
		s.Sources = append(s.Sources, "gsettings "+schema)
	}
}

// kdeglobals reads Plasma's font; its GTK theme is in gtk-3.0/settings.ini.
func (e Env) kdeglobals(s *Settings) {
	path := filepath.Join(e.configHome(), "kdeglobals")
	values := readIni(path, "General")
	if values == nil {
		return
	}
	s.Sources = append(s.Sources, path)
	if v := values["font"]; v != "" {
		s.Font = parseQtFont(v)
	}
	s.KDE = map[string][3]uint8{}
	for _, group := range []string{"Window", "Button", "View", "Selection"} {
		for k, v := range readIni(path, "Colors:"+group) {
			parts := strings.Split(v, ",")
			if len(parts) < 3 {
				continue
			}
			var c [3]uint8
			ok := true
			for i := range 3 {
				n, err := strconv.Atoi(strings.TrimSpace(parts[i]))
				if err != nil || n < 0 || n > 255 {
					ok = false
				}
				c[i] = uint8(n)
			}
			if ok {
				s.KDE[group+"."+k] = c
			}
		}
	}
	if c, ok := s.KDE["Window.BackgroundNormal"]; ok && int(c[0])+int(c[1])+int(c[2]) < 3*128 {
		s.Dark = true
	}
}

// lxqt reads LXQt's font.
func (e Env) lxqt(s *Settings) {
	path := filepath.Join(e.configHome(), "lxqt", "lxqt.conf")
	values := readIni(path, "Qt")
	if values == nil {
		return
	}
	s.Sources = append(s.Sources, path)
	if v := values["font"]; v != "" {
		s.Font = parseQtFont(strings.Trim(v, `"`))
	}
	s.QtStyle = values["style"]
	// LXQt's palette, in the same keys as KDE's colour scheme
	keys := map[string]string{
		"window_color": "Window.BackgroundNormal", "window_text_color": "Window.ForegroundNormal",
		"button_color": "Button.BackgroundNormal", "button_text_color": "Button.ForegroundNormal",
		"base_color": "View.BackgroundNormal", "text_color": "View.ForegroundNormal",
		"highlight_color": "Selection.BackgroundNormal", "highlighted_text_color": "Selection.ForegroundNormal",
	}
	for k, v := range readIni(path, "Palette") {
		if key, ok := keys[k]; ok {
			if c, ok := parseHex(v); ok {
				if s.KDE == nil {
					s.KDE = map[string][3]uint8{}
				}
				s.KDE[key] = c
			}
		}
	}
}

func parseHex(v string) ([3]uint8, bool) {
	v = strings.TrimPrefix(strings.TrimSpace(v), "#")
	if len(v) == 8 { // #aarrggbb
		v = v[2:]
	}
	if len(v) != 6 {
		return [3]uint8{}, false
	}
	n, err := strconv.ParseUint(v, 16, 32)
	if err != nil {
		return [3]uint8{}, false
	}
	return [3]uint8{uint8(n >> 16), uint8(n >> 8), uint8(n)}, true
}

// parseQtFont parses QFont::toString(): "Noto Sans,10,-1,5,50,0,0,0,0,0".
func parseQtFont(v string) Font {
	f := Font{Family: "Sans", Weight: 400, Size: 10}
	parts := strings.Split(v, ",")
	if len(parts) > 0 && parts[0] != "" {
		f.Family = parts[0]
	}
	if len(parts) > 1 {
		if size, err := strconv.ParseFloat(parts[1], 64); err == nil && size > 0 {
			f.Size = size
		}
	}
	if len(parts) > 4 {
		if w, err := strconv.Atoi(parts[4]); err == nil {
			switch {
			case w <= 0:
			case w > 100: // Qt 6 uses CSS weights
				f.Weight = w
			case w >= 75: // Qt 5 weights: 50 normal, 63 demibold, 75 bold
				f.Weight = 700
			case w >= 63:
				f.Weight = 600
			}
		}
	}
	return f
}

var pangoWeights = map[string]int{
	"thin": 100, "ultra-light": 200, "extra-light": 200, "ultralight": 200, "light": 300, "semi-light": 350, "book": 380,
	"regular": 400, "normal": 400, "medium": 500, "semi-bold": 600, "semibold": 600, "demi-bold": 600,
	"bold": 700, "ultra-bold": 800, "extra-bold": 800, "heavy": 900, "black": 900, "ultra-heavy": 1000,
}

var pangoStyles = map[string]bool{"italic": true, "oblique": true, "condensed": true, "semi-condensed": true,
	"expanded": true, "small-caps": true, "roman": true}

// ParseFont parses a Pango font description: "[FAMILY-LIST] [STYLE-OPTIONS] [SIZE]".
func ParseFont(desc string) Font {
	f := Font{Family: "Sans", Weight: 400, Size: 10}
	words := strings.Fields(strings.ReplaceAll(desc, ",", " "))
	if n := len(words); n > 0 {
		last := strings.TrimSuffix(words[n-1], "px")
		if size, err := strconv.ParseFloat(last, 64); err == nil {
			f.Size = size
			if strings.HasSuffix(words[n-1], "px") {
				f.Size = size * 72 / 96
			}
			words = words[:n-1]
		}
	}
	for len(words) > 0 {
		w := strings.ToLower(words[len(words)-1])
		if weight, ok := pangoWeights[w]; ok {
			f.Weight = weight
		} else if !pangoStyles[w] {
			break
		}
		words = words[:len(words)-1]
	}
	if len(words) > 0 {
		f.Family = strings.Join(words, " ")
	}
	return f
}

// Pixels is the font size in pixels at the settings' DPI.
func (s Settings) Pixels() float64 { return s.Font.Size * s.DPI / 72 }
