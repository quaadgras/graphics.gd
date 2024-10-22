package VideoStreamTheora

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VideoStream"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[VideoStream] resource handling the [url=https://www.theora.org/]Ogg Theora[/url] video format with [code].ogv[/code] extension. The Theora codec is decoded on the CPU.
[b]Note:[/b] While Ogg Theora videos can also have an [code].ogg[/code] extension, you will have to rename the extension to [code].ogv[/code] to use those videos within Godot.

*/
type Go [1]classdb.VideoStreamTheora
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VideoStreamTheora
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VideoStreamTheora"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsVideoStreamTheora() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVideoStreamTheora() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVideoStream() VideoStream.GD { return *((*VideoStream.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVideoStream() VideoStream.Go { return *((*VideoStream.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVideoStream(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVideoStream(), name)
	}
}
func init() {classdb.Register("VideoStreamTheora", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
