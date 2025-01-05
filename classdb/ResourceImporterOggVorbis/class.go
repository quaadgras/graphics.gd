// Package ResourceImporterOggVorbis provides methods for working with ResourceImporterOggVorbis object instances.
package ResourceImporterOggVorbis

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

/*
Ogg Vorbis is a lossy audio format, with better audio quality compared to [ResourceImporterMP3] at a given bitrate.
In most cases, it's recommended to use Ogg Vorbis over MP3. However, if you're using an MP3 sound source with no higher quality source available, then it's recommended to use the MP3 file directly to avoid double lossy compression.
Ogg Vorbis requires more CPU to decode than [ResourceImporterWAV]. If you need to play a lot of simultaneous sounds, it's recommended to use WAV for those sounds instead, especially if targeting low-end devices.
*/
type Instance [1]gdclass.ResourceImporterOggVorbis
type Any interface {
	gd.IsClass
	AsResourceImporterOggVorbis() Instance
}

/*
This method loads audio data from a PackedByteArray buffer into an AudioStreamOggVorbis object.
*/
func LoadFromBuffer(buffer []byte) [1]gdclass.AudioStreamOggVorbis {
	self := Instance{}
	return [1]gdclass.AudioStreamOggVorbis(class(self).LoadFromBuffer(gd.NewPackedByteSlice(buffer)))
}

/*
This method loads audio data from a file into an AudioStreamOggVorbis object. The file path is provided as a string.
*/
func LoadFromFile(path string) [1]gdclass.AudioStreamOggVorbis {
	self := Instance{}
	return [1]gdclass.AudioStreamOggVorbis(class(self).LoadFromFile(gd.NewString(path)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ResourceImporterOggVorbis

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ResourceImporterOggVorbis"))
	return Instance{*(*gdclass.ResourceImporterOggVorbis)(unsafe.Pointer(&object))}
}

/*
This method loads audio data from a PackedByteArray buffer into an AudioStreamOggVorbis object.
*/
//go:nosplit
func (self class) LoadFromBuffer(buffer gd.PackedByteArray) [1]gdclass.AudioStreamOggVorbis {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(buffer))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceImporterOggVorbis.Bind_load_from_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.AudioStreamOggVorbis{gd.PointerWithOwnershipTransferredToGo[gdclass.AudioStreamOggVorbis](r_ret.Get())}
	frame.Free()
	return ret
}

/*
This method loads audio data from a file into an AudioStreamOggVorbis object. The file path is provided as a string.
*/
//go:nosplit
func (self class) LoadFromFile(path gd.String) [1]gdclass.AudioStreamOggVorbis {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceImporterOggVorbis.Bind_load_from_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.AudioStreamOggVorbis{gd.PointerWithOwnershipTransferredToGo[gdclass.AudioStreamOggVorbis](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsResourceImporterOggVorbis() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResourceImporterOggVorbis() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
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
	gdclass.Register("ResourceImporterOggVorbis", func(ptr gd.Object) any {
		return [1]gdclass.ResourceImporterOggVorbis{*(*gdclass.ResourceImporterOggVorbis)(unsafe.Pointer(&ptr))}
	})
}
