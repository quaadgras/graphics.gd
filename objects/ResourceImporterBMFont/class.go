package ResourceImporterBMFont

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/ResourceImporter"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
The BMFont format is a format created by the [url=https://www.angelcode.com/products/bmfont/]BMFont[/url] program. Many BMFont-compatible programs also exist, like [url=https://www.bmglyph.com/]BMGlyph[/url].
Compared to [ResourceImporterImageFont], [ResourceImporterBMFont] supports bitmap fonts with varying glyph widths/heights.
See also [ResourceImporterDynamicFont].
*/
type Instance [1]classdb.ResourceImporterBMFont
type Any interface {
	gd.IsClass
	AsResourceImporterBMFont() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ResourceImporterBMFont

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ResourceImporterBMFont"))
	return Instance{*(*classdb.ResourceImporterBMFont)(unsafe.Pointer(&object))}
}

func (self class) AsResourceImporterBMFont() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsResourceImporterBMFont() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResourceImporter() ResourceImporter.Advanced {
	return *((*ResourceImporter.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResourceImporter() ResourceImporter.Instance {
	return *((*ResourceImporter.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResourceImporter(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResourceImporter(), name)
	}
}
func init() {
	classdb.Register("ResourceImporterBMFont", func(ptr gd.Object) any {
		return [1]classdb.ResourceImporterBMFont{*(*classdb.ResourceImporterBMFont)(unsafe.Pointer(&ptr))}
	})
}
