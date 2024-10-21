package EditorSceneFormatImporterUFBX

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/EditorSceneFormatImporter"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
EditorSceneFormatImporterUFBX is designed to load FBX files and supports both binary and ASCII FBX files from version 3000 onward. This class supports various 3D object types like meshes, skins, blend shapes, materials, and rigging information. The class aims for feature parity with the official FBX SDK and supports FBX 7.4 specifications.

*/
type Simple [1]classdb.EditorSceneFormatImporterUFBX
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorSceneFormatImporterUFBX
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsEditorSceneFormatImporterUFBX() Expert { return self[0].AsEditorSceneFormatImporterUFBX() }


//go:nosplit
func (self Simple) AsEditorSceneFormatImporterUFBX() Simple { return self[0].AsEditorSceneFormatImporterUFBX() }


//go:nosplit
func (self class) AsEditorSceneFormatImporter() EditorSceneFormatImporter.Expert { return self[0].AsEditorSceneFormatImporter() }


//go:nosplit
func (self Simple) AsEditorSceneFormatImporter() EditorSceneFormatImporter.Simple { return self[0].AsEditorSceneFormatImporter() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("EditorSceneFormatImporterUFBX", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
