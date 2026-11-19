package cosmic

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	v, err := Parse(`(
    base: (red: 0.1, green: 0.2, blue: 0.3, alpha: 1.0),
    component: ( focus: ( red: 1.0, green: 0.68, blue: 0.0, alpha: 1.0, ), ),
    radius_s: (8.0, 8.0, 8.0, 8.0),
    name: "cosmic-dark", // a comment
    opt: Some(4), none: None, flag: true,
)`)
	if err != nil {
		t.Fatal(err)
	}
	if c, ok := Color(v, "component", "focus"); !ok || c.G != 0.68 {
		t.Fatalf("focus %v %v", c, ok)
	}
	if c, _ := Color(v, "base"); c != (RGBA{0.1, 0.2, 0.3, 1}) {
		t.Fatalf("base %v", c)
	}
	if Get(v, "name") != "cosmic-dark" || Get(v, "opt") != 4.0 || Get(v, "flag") != true {
		t.Fatalf("%v", v)
	}
	if l, ok := Get(v, "radius_s").([]Value); !ok || len(l) != 4 {
		t.Fatalf("radius %v", Get(v, "radius_s"))
	}
}

func TestLoad(t *testing.T) {
	cfg := t.TempDir()
	write := func(rel, s string) {
		p := filepath.Join(cfg, "cosmic", rel)
		os.MkdirAll(filepath.Dir(p), 0o755)
		os.WriteFile(p, []byte(s), 0o644)
	}
	write("com.system76.CosmicTheme.Mode/v1/is_dark", "true")
	write("com.system76.CosmicTheme.Dark/v1/accent", "(base: (red: 1.0, green: 0.5, blue: 0.0, alpha: 1.0))")
	write("com.system76.CosmicTheme.Dark/v1/corner_radii", "(radius_s: (8.0, 8.0, 8.0, 8.0))")
	th, err := Load(cfg)
	if err != nil {
		t.Fatal(err)
	}
	if c, ok := th.Color("accent", "base"); !th.Dark || !ok || c.G != 0.5 || th.Radius("radius_s", 0) != 8 {
		t.Fatalf("%+v", th)
	}
}

// TestReal loads this machine's theme when there is one.
func TestReal(t *testing.T) {
	home, _ := os.UserHomeDir()
	th, err := Load(filepath.Join(home, ".config"))
	if err != nil {
		t.Skip("no COSMIC theme here")
	}
	accent, _ := th.Color("accent", "base")
	bg, _ := th.Color("background", "base")
	t.Logf("dark %v, %d files, accent %v, background %v, radius_s %v", th.Dark, len(th.Files), accent, bg, th.Radius("radius_s", -1))
}
