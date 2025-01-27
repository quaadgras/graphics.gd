// Package FramebufferCacheRD provides methods for working with FramebufferCacheRD object instances.
package FramebufferCacheRD

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

/*
Framebuffer cache manager for Rendering Device based renderers. Provides a way to create a framebuffer and reuse it in subsequent calls for as long as the used textures exists. Framebuffers will automatically be cleaned up when dependent objects are freed.
*/
type Instance [1]gdclass.FramebufferCacheRD

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsFramebufferCacheRD() Instance
}

/*
Creates, or obtains a cached, framebuffer. [param textures] lists textures accessed. [param passes] defines the subpasses and texture allocation, if left empty a single pass is created and textures are allocated depending on their usage flags. [param views] defines the number of views used when rendering.
*/
func GetCacheMultipass(textures [][]RID.Texture, passes [][1]gdclass.RDFramebufferPass, views int) RID.Framebuffer { //gd:FramebufferCacheRD.get_cache_multipass
	self := Instance{}
	return RID.Framebuffer(class(self).GetCacheMultipass(gd.ArrayFromSlice[Array.Contains[gd.RID]](textures), gd.ArrayFromSlice[Array.Contains[[1]gdclass.RDFramebufferPass]](passes), gd.Int(views)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.FramebufferCacheRD

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("FramebufferCacheRD"))
	casted := Instance{*(*gdclass.FramebufferCacheRD)(unsafe.Pointer(&object))}
	return casted
}

/*
Creates, or obtains a cached, framebuffer. [param textures] lists textures accessed. [param passes] defines the subpasses and texture allocation, if left empty a single pass is created and textures are allocated depending on their usage flags. [param views] defines the number of views used when rendering.
*/
//go:nosplit
func (self class) GetCacheMultipass(textures Array.Contains[gd.RID], passes Array.Contains[[1]gdclass.RDFramebufferPass], views gd.Int) gd.RID { //gd:FramebufferCacheRD.get_cache_multipass
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(textures)))
	callframe.Arg(frame, pointers.Get(gd.InternalArray(passes)))
	callframe.Arg(frame, views)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FramebufferCacheRD.Bind_get_cache_multipass, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsFramebufferCacheRD() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsFramebufferCacheRD() Instance { return *((*Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Instance(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("FramebufferCacheRD", func(ptr gd.Object) any {
		return [1]gdclass.FramebufferCacheRD{*(*gdclass.FramebufferCacheRD)(unsafe.Pointer(&ptr))}
	})
}
