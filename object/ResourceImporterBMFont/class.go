package ResourceImporterBMFont

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/ResourceImporter"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The BMFont format is a format created by the [url=https://www.angelcode.com/products/bmfont/]BMFont[/url] program. Many BMFont-compatible programs also exist, like [url=https://www.bmglyph.com/]BMGlyph[/url].
Compared to [ResourceImporterImageFont], [ResourceImporterBMFont] supports bitmap fonts with varying glyph widths/heights.
See also [ResourceImporterDynamicFont].

*/
type Simple [1]classdb.ResourceImporterBMFont
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ResourceImporterBMFont
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsResourceImporterBMFont() Expert { return self[0].AsResourceImporterBMFont() }


//go:nosplit
func (self Simple) AsResourceImporterBMFont() Simple { return self[0].AsResourceImporterBMFont() }


//go:nosplit
func (self class) AsResourceImporter() ResourceImporter.Expert { return self[0].AsResourceImporter() }


//go:nosplit
func (self Simple) AsResourceImporter() ResourceImporter.Simple { return self[0].AsResourceImporter() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("ResourceImporterBMFont", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
