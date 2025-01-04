// Package ScriptLanguage provides methods for working with ScriptLanguage object instances.
package ScriptLanguage

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type Instance [1]gdclass.ScriptLanguage
type Any interface {
	gd.IsClass
	AsScriptLanguage() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ScriptLanguage

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ScriptLanguage"))
	return Instance{*(*gdclass.ScriptLanguage)(unsafe.Pointer(&object))}
}

func (self class) AsScriptLanguage() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsScriptLanguage() Instance { return *((*Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	gdclass.Register("ScriptLanguage", func(ptr gd.Object) any {
		return [1]gdclass.ScriptLanguage{*(*gdclass.ScriptLanguage)(unsafe.Pointer(&ptr))}
	})
}

type ScriptNameCasing = gdclass.ScriptLanguageScriptNameCasing

const (
	ScriptNameCasingAuto       ScriptNameCasing = 0
	ScriptNameCasingPascalCase ScriptNameCasing = 1
	ScriptNameCasingSnakeCase  ScriptNameCasing = 2
	ScriptNameCasingKebabCase  ScriptNameCasing = 3
)
