package ScriptLanguage

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

type Go [1]classdb.ScriptLanguage
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ScriptLanguage
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ScriptLanguage"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsScriptLanguage() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsScriptLanguage() Go { return *((*Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("ScriptLanguage", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type ScriptNameCasing = classdb.ScriptLanguageScriptNameCasing

const (
	ScriptNameCasingAuto ScriptNameCasing = 0
	ScriptNameCasingPascalCase ScriptNameCasing = 1
	ScriptNameCasingSnakeCase ScriptNameCasing = 2
	ScriptNameCasingKebabCase ScriptNameCasing = 3
)
