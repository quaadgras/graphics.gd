package ResourceImporterBMFont

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/ResourceImporter"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
The BMFont format is a format created by the [url=https://www.angelcode.com/products/bmfont/]BMFont[/url] program. Many BMFont-compatible programs also exist, like [url=https://www.bmglyph.com/]BMGlyph[/url].
Compared to [ResourceImporterImageFont], [ResourceImporterBMFont] supports bitmap fonts with varying glyph widths/heights.
See also [ResourceImporterDynamicFont].

*/
type Go [1]classdb.ResourceImporterBMFont
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ResourceImporterBMFont
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ResourceImporterBMFont"))
	return Go{classdb.ResourceImporterBMFont(object)}
}

func (self class) AsResourceImporterBMFont() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsResourceImporterBMFont() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResourceImporter() ResourceImporter.GD { return *((*ResourceImporter.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResourceImporter() ResourceImporter.Go { return *((*ResourceImporter.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResourceImporter(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResourceImporter(), name)
	}
}
func init() {classdb.Register("ResourceImporterBMFont", func(ptr gd.Object) any { return classdb.ResourceImporterBMFont(ptr) })}
