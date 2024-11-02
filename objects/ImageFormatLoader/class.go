package ImageFormatLoader

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
The engine supports multiple image formats out of the box (PNG, SVG, JPEG, WebP to name a few), but you can choose to implement support for additional image formats by extending [ImageFormatLoaderExtension].
*/
type Instance [1]classdb.ImageFormatLoader

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ImageFormatLoader

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ImageFormatLoader"))
	return Instance{classdb.ImageFormatLoader(object)}
}

func (self class) AsImageFormatLoader() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsImageFormatLoader() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted      { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted   { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("ImageFormatLoader", func(ptr gd.Object) any { return classdb.ImageFormatLoader(ptr) })
}

type LoaderFlags = classdb.ImageFormatLoaderLoaderFlags

const (
	FlagNone          LoaderFlags = 0
	FlagForceLinear   LoaderFlags = 1
	FlagConvertColors LoaderFlags = 2
)
