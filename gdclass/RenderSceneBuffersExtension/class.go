package RenderSceneBuffersExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/RenderSceneBuffers"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
This class allows for a RenderSceneBuffer implementation to be made in GDExtension.

	// RenderSceneBuffersExtension methods that can be overridden by a [Class] that extends it.
	type RenderSceneBuffersExtension interface {
		//Implement this in GDExtension to handle the (re)sizing of a viewport.
		Configure(config gdclass.RenderSceneBuffersConfiguration)
		//Implement this in GDExtension to record a new FSR sharpness value.
		SetFsrSharpness(fsr_sharpness float64)
		//Implement this in GDExtension to change the texture mipmap bias.
		SetTextureMipmapBias(texture_mipmap_bias float64)
		//Implement this in GDExtension to react to the debanding flag changing.
		SetUseDebanding(use_debanding bool)
	}
*/
type Instance [1]classdb.RenderSceneBuffersExtension

/*
Implement this in GDExtension to handle the (re)sizing of a viewport.
*/
func (Instance) _configure(impl func(ptr unsafe.Pointer, config gdclass.RenderSceneBuffersConfiguration)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var config = gdclass.RenderSceneBuffersConfiguration{pointers.New[classdb.RenderSceneBuffersConfiguration]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(config[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, config)
	}
}

/*
Implement this in GDExtension to record a new FSR sharpness value.
*/
func (Instance) _set_fsr_sharpness(impl func(ptr unsafe.Pointer, fsr_sharpness float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var fsr_sharpness = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, float64(fsr_sharpness))
	}
}

/*
Implement this in GDExtension to change the texture mipmap bias.
*/
func (Instance) _set_texture_mipmap_bias(impl func(ptr unsafe.Pointer, texture_mipmap_bias float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var texture_mipmap_bias = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, float64(texture_mipmap_bias))
	}
}

/*
Implement this in GDExtension to react to the debanding flag changing.
*/
func (Instance) _set_use_debanding(impl func(ptr unsafe.Pointer, use_debanding bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var use_debanding = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, use_debanding)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RenderSceneBuffersExtension

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RenderSceneBuffersExtension"))
	return Instance{classdb.RenderSceneBuffersExtension(object)}
}

/*
Implement this in GDExtension to handle the (re)sizing of a viewport.
*/
func (class) _configure(impl func(ptr unsafe.Pointer, config gdclass.RenderSceneBuffersConfiguration)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var config = gdclass.RenderSceneBuffersConfiguration{pointers.New[classdb.RenderSceneBuffersConfiguration]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(config[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, config)
	}
}

/*
Implement this in GDExtension to record a new FSR sharpness value.
*/
func (class) _set_fsr_sharpness(impl func(ptr unsafe.Pointer, fsr_sharpness gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var fsr_sharpness = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, fsr_sharpness)
	}
}

/*
Implement this in GDExtension to change the texture mipmap bias.
*/
func (class) _set_texture_mipmap_bias(impl func(ptr unsafe.Pointer, texture_mipmap_bias gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var texture_mipmap_bias = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, texture_mipmap_bias)
	}
}

/*
Implement this in GDExtension to react to the debanding flag changing.
*/
func (class) _set_use_debanding(impl func(ptr unsafe.Pointer, use_debanding bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var use_debanding = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, use_debanding)
	}
}

func (self class) AsRenderSceneBuffersExtension() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRenderSceneBuffersExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRenderSceneBuffers() RenderSceneBuffers.Advanced {
	return *((*RenderSceneBuffers.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRenderSceneBuffers() RenderSceneBuffers.Instance {
	return *((*RenderSceneBuffers.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_configure":
		return reflect.ValueOf(self._configure)
	case "_set_fsr_sharpness":
		return reflect.ValueOf(self._set_fsr_sharpness)
	case "_set_texture_mipmap_bias":
		return reflect.ValueOf(self._set_texture_mipmap_bias)
	case "_set_use_debanding":
		return reflect.ValueOf(self._set_use_debanding)
	default:
		return gd.VirtualByName(self.AsRenderSceneBuffers(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_configure":
		return reflect.ValueOf(self._configure)
	case "_set_fsr_sharpness":
		return reflect.ValueOf(self._set_fsr_sharpness)
	case "_set_texture_mipmap_bias":
		return reflect.ValueOf(self._set_texture_mipmap_bias)
	case "_set_use_debanding":
		return reflect.ValueOf(self._set_use_debanding)
	default:
		return gd.VirtualByName(self.AsRenderSceneBuffers(), name)
	}
}
func init() {
	classdb.Register("RenderSceneBuffersExtension", func(ptr gd.Object) any { return classdb.RenderSceneBuffersExtension(ptr) })
}
