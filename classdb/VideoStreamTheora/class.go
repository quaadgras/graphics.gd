// Package VideoStreamTheora provides methods for working with VideoStreamTheora object instances.
package VideoStreamTheora

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/classdb/VideoStream"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode

/*
[VideoStream] resource handling the [url=https://www.theora.org/]Ogg Theora[/url] video format with [code].ogv[/code] extension. The Theora codec is decoded on the CPU.
[b]Note:[/b] While Ogg Theora videos can also have an [code].ogg[/code] extension, you will have to rename the extension to [code].ogv[/code] to use those videos within Godot.
*/
type Instance [1]gdclass.VideoStreamTheora

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVideoStreamTheora() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VideoStreamTheora

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VideoStreamTheora"))
	casted := Instance{*(*gdclass.VideoStreamTheora)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self class) AsVideoStreamTheora() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVideoStreamTheora() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsVideoStream() VideoStream.Advanced {
	return *((*VideoStream.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVideoStream() VideoStream.Instance {
	return *((*VideoStream.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(VideoStream.Advanced(self.AsVideoStream()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VideoStream.Instance(self.AsVideoStream()), name)
	}
}
func init() {
	gdclass.Register("VideoStreamTheora", func(ptr gd.Object) any {
		return [1]gdclass.VideoStreamTheora{*(*gdclass.VideoStreamTheora)(unsafe.Pointer(&ptr))}
	})
}
