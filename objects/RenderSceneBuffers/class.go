package RenderSceneBuffers

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
Abstract scene buffers object, created for each viewport for which 3D rendering is done. It manages any additional buffers used during rendering and will discard buffers when the viewport is resized.
[b]Note:[/b] This is an internal rendering server object, do not instantiate this from script.
*/
type Instance [1]classdb.RenderSceneBuffers

/*
This method is called by the rendering server when the associated viewports configuration is changed. It will discard the old buffers and recreate the internal buffers used.
*/
func (self Instance) Configure(config objects.RenderSceneBuffersConfiguration) {
	class(self).Configure(config)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RenderSceneBuffers

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RenderSceneBuffers"))
	return Instance{classdb.RenderSceneBuffers(object)}
}

/*
This method is called by the rendering server when the associated viewports configuration is changed. It will discard the old buffers and recreate the internal buffers used.
*/
//go:nosplit
func (self class) Configure(config objects.RenderSceneBuffersConfiguration) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(config[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffers.Bind_configure, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsRenderSceneBuffers() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRenderSceneBuffers() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted       { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

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
	classdb.Register("RenderSceneBuffers", func(ptr gd.Object) any { return classdb.RenderSceneBuffers(ptr) })
}
