// Package ResourceImporterMP3 provides methods for working with ResourceImporterMP3 object instances.
package ResourceImporterMP3

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/ResourceImporter"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
MP3 is a lossy audio format, with worse audio quality compared to [ResourceImporterOggVorbis] at a given bitrate.
In most cases, it's recommended to use Ogg Vorbis over MP3. However, if you're using an MP3 sound source with no higher quality source available, then it's recommended to use the MP3 file directly to avoid double lossy compression.
MP3 requires more CPU to decode than [ResourceImporterWAV]. If you need to play a lot of simultaneous sounds, it's recommended to use WAV for those sounds instead, especially if targeting low-end devices.
*/
type Instance [1]gdclass.ResourceImporterMP3

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsResourceImporterMP3() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ResourceImporterMP3

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ResourceImporterMP3"))
	casted := Instance{*(*gdclass.ResourceImporterMP3)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self class) AsResourceImporterMP3() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsResourceImporterMP3() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResourceImporter() ResourceImporter.Advanced {
	return *((*ResourceImporter.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResourceImporter() ResourceImporter.Instance {
	return *((*ResourceImporter.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(ResourceImporter.Advanced(self.AsResourceImporter()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(ResourceImporter.Instance(self.AsResourceImporter()), name)
	}
}
func init() {
	gdclass.Register("ResourceImporterMP3", func(ptr gd.Object) any {
		return [1]gdclass.ResourceImporterMP3{*(*gdclass.ResourceImporterMP3)(unsafe.Pointer(&ptr))}
	})
}
