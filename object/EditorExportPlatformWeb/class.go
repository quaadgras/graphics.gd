package EditorExportPlatformWeb

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/EditorExportPlatform"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The Web exporter customizes how a web build is handled. In the editor's "Export" window, it is created when adding a new "Web" preset.
[b]Note:[/b] Godot on Web is rendered inside a [code]<canvas>[/code] tag. Normally, the canvas cannot be positioned or resized manually, but otherwise acts as the main [Window] of the application.

*/
type Simple [1]classdb.EditorExportPlatformWeb
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorExportPlatformWeb
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsEditorExportPlatformWeb() Expert { return self[0].AsEditorExportPlatformWeb() }


//go:nosplit
func (self Simple) AsEditorExportPlatformWeb() Simple { return self[0].AsEditorExportPlatformWeb() }


//go:nosplit
func (self class) AsEditorExportPlatform() EditorExportPlatform.Expert { return self[0].AsEditorExportPlatform() }


//go:nosplit
func (self Simple) AsEditorExportPlatform() EditorExportPlatform.Simple { return self[0].AsEditorExportPlatform() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("EditorExportPlatformWeb", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
