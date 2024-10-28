package EditorSceneFormatImporterGLTF

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/EditorSceneFormatImporter"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

type Go [1]classdb.EditorSceneFormatImporterGLTF
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorSceneFormatImporterGLTF
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorSceneFormatImporterGLTF"))
	return Go{classdb.EditorSceneFormatImporterGLTF(object)}
}

func (self class) AsEditorSceneFormatImporterGLTF() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorSceneFormatImporterGLTF() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsEditorSceneFormatImporter() EditorSceneFormatImporter.GD { return *((*EditorSceneFormatImporter.GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorSceneFormatImporter() EditorSceneFormatImporter.Go { return *((*EditorSceneFormatImporter.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsEditorSceneFormatImporter(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsEditorSceneFormatImporter(), name)
	}
}
func init() {classdb.Register("EditorSceneFormatImporterGLTF", func(ptr gd.Object) any { return classdb.EditorSceneFormatImporterGLTF(ptr) })}
