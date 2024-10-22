package RenderSceneBuffersExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/RenderSceneBuffers"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
type Go [1]classdb.RenderSceneBuffersExtension

/*
Implement this in GDExtension to handle the (re)sizing of a viewport.
*/
func (Go) _configure(impl func(ptr unsafe.Pointer, config gdclass.RenderSceneBuffersConfiguration) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var config gdclass.RenderSceneBuffersConfiguration
		config[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, config)
		gc.End()
	}
}

/*
Implement this in GDExtension to record a new FSR sharpness value.
*/
func (Go) _set_fsr_sharpness(impl func(ptr unsafe.Pointer, fsr_sharpness float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var fsr_sharpness = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, float64(fsr_sharpness))
		gc.End()
	}
}

/*
Implement this in GDExtension to change the texture mipmap bias.
*/
func (Go) _set_texture_mipmap_bias(impl func(ptr unsafe.Pointer, texture_mipmap_bias float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var texture_mipmap_bias = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, float64(texture_mipmap_bias))
		gc.End()
	}
}

/*
Implement this in GDExtension to react to the debanding flag changing.
*/
func (Go) _set_use_debanding(impl func(ptr unsafe.Pointer, use_debanding bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var use_debanding = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, use_debanding)
		gc.End()
	}
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RenderSceneBuffersExtension
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RenderSceneBuffersExtension"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Implement this in GDExtension to handle the (re)sizing of a viewport.
*/
func (class) _configure(impl func(ptr unsafe.Pointer, config gdclass.RenderSceneBuffersConfiguration) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var config gdclass.RenderSceneBuffersConfiguration
		config[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, config)
		ctx.End()
	}
}

/*
Implement this in GDExtension to record a new FSR sharpness value.
*/
func (class) _set_fsr_sharpness(impl func(ptr unsafe.Pointer, fsr_sharpness gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var fsr_sharpness = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, fsr_sharpness)
		ctx.End()
	}
}

/*
Implement this in GDExtension to change the texture mipmap bias.
*/
func (class) _set_texture_mipmap_bias(impl func(ptr unsafe.Pointer, texture_mipmap_bias gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var texture_mipmap_bias = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, texture_mipmap_bias)
		ctx.End()
	}
}

/*
Implement this in GDExtension to react to the debanding flag changing.
*/
func (class) _set_use_debanding(impl func(ptr unsafe.Pointer, use_debanding bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var use_debanding = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, use_debanding)
		ctx.End()
	}
}

func (self class) AsRenderSceneBuffersExtension() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRenderSceneBuffersExtension() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRenderSceneBuffers() RenderSceneBuffers.GD { return *((*RenderSceneBuffers.GD)(unsafe.Pointer(&self))) }
func (self Go) AsRenderSceneBuffers() RenderSceneBuffers.Go { return *((*RenderSceneBuffers.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_configure": return reflect.ValueOf(self._configure);
	case "_set_fsr_sharpness": return reflect.ValueOf(self._set_fsr_sharpness);
	case "_set_texture_mipmap_bias": return reflect.ValueOf(self._set_texture_mipmap_bias);
	case "_set_use_debanding": return reflect.ValueOf(self._set_use_debanding);
	default: return gd.VirtualByName(self.AsRenderSceneBuffers(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_configure": return reflect.ValueOf(self._configure);
	case "_set_fsr_sharpness": return reflect.ValueOf(self._set_fsr_sharpness);
	case "_set_texture_mipmap_bias": return reflect.ValueOf(self._set_texture_mipmap_bias);
	case "_set_use_debanding": return reflect.ValueOf(self._set_use_debanding);
	default: return gd.VirtualByName(self.AsRenderSceneBuffers(), name)
	}
}
func init() {classdb.Register("RenderSceneBuffersExtension", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
