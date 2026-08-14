package startup

import (
	"os"

	"graphics.gd/classdb/EditorInterface"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/ProjectSettings"
	"graphics.gd/classdb/SceneTree"
	gd "graphics.gd/internal"
	"graphics.gd/internal/packagegen"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Object"
)

// Whenever a graphics.gd project is opened inside the editor, type-safe Go
// packages are automatically generated for the contents of the graphics
// directory: each directory containing resources gets a graphics.gen.go
// file with a Resource.Library declaration per file extension, a struct
// mirroring the tree of each scene and constants for the project's input
// actions, global groups and named layers, see
// https://github.com/quaadgras/graphics.gd/discussions/283
//
// The packages are kept up to date while the editor is open: whenever the
// editor reports filesystem changes, regeneration re-runs once the changes
// settle. Generation can be disabled by turning off the
// [GraphicsPackagesSetting] project setting (under "Graphics Gd" in the
// project settings dialog).

// GraphicsPackagesSetting is the boolean project setting that controls
// whether Go packages are generated for the graphics directory when the
// project is opened inside the editor. Defaults to enabled.
const GraphicsPackagesSetting = "graphics_gd/generate/graphics_packages"

func init() {
	gd.EditorStartupFunctions = append(gd.EditorStartupFunctions, func() {
		Callable.Defer(Callable.New(func() {
			if !Engine.IsEditorHint() {
				return
			}
			// Register the setting so it shows up in the project settings
			// dialog: registering the initial value keeps it out of
			// project.godot until the user actually turns it off.
			if !ProjectSettings.HasSetting(GraphicsPackagesSetting) {
				ProjectSettings.SetSetting(GraphicsPackagesSetting, true)
			}
			ProjectSettings.SetInitialValue(GraphicsPackagesSetting, true)
			ProjectSettings.SetAsBasic(GraphicsPackagesSetting, true)
			makeGraphicsPackages()
			// Regenerate when the filesystem settles after a change: each
			// change during the debounce window restarts it, so a burst of
			// imports triggers a single regeneration.
			var pending int
			EditorInterface.GetResourceFilesystem().OnFilesystemChanged(func() {
				pending++
				generation := pending
				tree, ok := Object.As[SceneTree.Instance](Engine.GetMainLoop())
				if !ok {
					return
				}
				tree.CreateTimer(0.5).OnTimeout(func() {
					if pending == generation {
						makeGraphicsPackages()
					}
				})
			})
		}))
	})
}

func makeGraphicsPackages() {
	if enabled, ok := ProjectSettings.GetSetting(GraphicsPackagesSetting, true).(bool); ok && !enabled {
		return
	}
	if err := packagegen.All(ProjectSettings.GlobalizePath("res://")); err != nil {
		Engine.RaiseWarning("graphics package generation: " + err.Error())
		os.Stderr.WriteString("graphics package generation: " + err.Error() + "\n")
	}
}
