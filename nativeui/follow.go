package nativeui

import (
	"sync"

	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/variant/Object"
)

type follower struct {
	target    Control.Instance
	customize func(*Options)
}

var following struct {
	sync.Mutex
	hooked    bool
	followers []follower
}

// Follow themes target (and so its children) to match the system now, and
// rebuilds the theme whenever the system theme changes (dark mode, accent
// color, contrast...). customize, if given, adjusts the detected options
// before each rebuild, for example to force a look or a scale.
func Follow(target Control.Instance, customize ...func(*Options)) {
	f := follower{target: target}
	if len(customize) > 0 {
		f.customize = customize[0]
	}
	f.apply()
	following.Lock()
	defer following.Unlock()
	following.followers = append(following.followers, f)
	if !following.hooked {
		following.hooked = true
		DisplayServer.SetSystemThemeChangeCallback(systemThemeChanged)
	}
}

func (f follower) apply() {
	opts := Detect()
	if f.customize != nil {
		f.customize(&opts)
	}
	f.target.SetTheme(New(opts))
}

func systemThemeChanged() {
	following.Lock()
	alive := following.followers[:0]
	for _, f := range following.followers {
		if Object.InstanceIsValid(f.target) {
			alive = append(alive, f)
		}
	}
	following.followers = alive
	list := append([]follower(nil), alive...)
	following.Unlock()
	for _, f := range list {
		f.apply()
	}
}
