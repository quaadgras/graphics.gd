package ScriptLanguage

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

type Simple [1]classdb.ScriptLanguage
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ScriptLanguage
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsScriptLanguage() Expert { return self[0].AsScriptLanguage() }


//go:nosplit
func (self Simple) AsScriptLanguage() Simple { return self[0].AsScriptLanguage() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("ScriptLanguage", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type ScriptNameCasing = classdb.ScriptLanguageScriptNameCasing

const (
	ScriptNameCasingAuto ScriptNameCasing = 0
	ScriptNameCasingPascalCase ScriptNameCasing = 1
	ScriptNameCasingSnakeCase ScriptNameCasing = 2
	ScriptNameCasingKebabCase ScriptNameCasing = 3
)
