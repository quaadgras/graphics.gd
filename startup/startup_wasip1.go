package startup

import (
	"fmt"
	"unsafe"

	gd "graphics.gd/internal"
)

//go:wasmimport gdextension get_godot_version_major
func get_godot_version_major() uint32

//go:wasmimport gdextension get_godot_version_minor
func get_godot_version_minor() uint32

//go:wasmimport gdextension get_godot_version_patch
func get_godot_version_patch() uint32

//go:wasmimport gdextension string_name_new_with_utf8_chars_and_len
func string_name_new_with_utf8_chars_and_len(chars unsafe.Pointer, len uint32) uint64

// If we are starting up under wasmip1 then that means we are peforming a hot-reload.
// we need to link up with our host and initialise the Global state, we will then
// receive all of the active extension classes so that they can be re-created.
func init() {
	gd.Global.GetGodotVersion = func() gd.Version {
		return gd.Version{
			Major: get_godot_version_major(),
			Minor: get_godot_version_minor(),
			Patch: get_godot_version_patch(),
		}
	}
	fmt.Println("Godot Version: ", gd.Global.GetGodotVersion().Major)
}
