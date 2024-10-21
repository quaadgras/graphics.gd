package RenderDataExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/RenderData"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class allows for a RenderData implementation to be made in GDExtension.
	// RenderDataExtension methods that can be overridden by a [Class] that extends it.
	type RenderDataExtension interface {
		//Implement this in GDExtension to return the implementation's [RenderSceneBuffers] object.
		GetRenderSceneBuffers() object.RenderSceneBuffers
		//Implement this in GDExtension to return the implementation's [RenderSceneDataExtension] object.
		GetRenderSceneData() object.RenderSceneData
		//Implement this in GDExtension to return the [RID] of the implementation's environment object.
		GetEnvironment() gd.RID
		//Implement this in GDExtension to return the [RID] for the implementation's camera attributes object.
		GetCameraAttributes() gd.RID
	}

*/
type Simple [1]classdb.RenderDataExtension
func (Simple) _get_render_scene_buffers(impl func(ptr unsafe.Pointer) [1]classdb.RenderSceneBuffers, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _get_render_scene_data(impl func(ptr unsafe.Pointer) [1]classdb.RenderSceneData, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _get_environment(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _get_camera_attributes(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RenderDataExtension
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Implement this in GDExtension to return the implementation's [RenderSceneBuffers] object.
*/
func (class) _get_render_scene_buffers(impl func(ptr unsafe.Pointer) object.RenderSceneBuffers, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		ctx.End()
	}
}

/*
Implement this in GDExtension to return the implementation's [RenderSceneDataExtension] object.
*/
func (class) _get_render_scene_data(impl func(ptr unsafe.Pointer) object.RenderSceneData, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		ctx.End()
	}
}

/*
Implement this in GDExtension to return the [RID] of the implementation's environment object.
*/
func (class) _get_environment(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Implement this in GDExtension to return the [RID] for the implementation's camera attributes object.
*/
func (class) _get_camera_attributes(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}


//go:nosplit
func (self class) AsRenderDataExtension() Expert { return self[0].AsRenderDataExtension() }


//go:nosplit
func (self Simple) AsRenderDataExtension() Simple { return self[0].AsRenderDataExtension() }


//go:nosplit
func (self class) AsRenderData() RenderData.Expert { return self[0].AsRenderData() }


//go:nosplit
func (self Simple) AsRenderData() RenderData.Simple { return self[0].AsRenderData() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_render_scene_buffers": return reflect.ValueOf(self._get_render_scene_buffers);
	case "_get_render_scene_data": return reflect.ValueOf(self._get_render_scene_data);
	case "_get_environment": return reflect.ValueOf(self._get_environment);
	case "_get_camera_attributes": return reflect.ValueOf(self._get_camera_attributes);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_get_render_scene_buffers": return reflect.ValueOf(self._get_render_scene_buffers);
	case "_get_render_scene_data": return reflect.ValueOf(self._get_render_scene_data);
	case "_get_environment": return reflect.ValueOf(self._get_environment);
	case "_get_camera_attributes": return reflect.ValueOf(self._get_camera_attributes);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("RenderDataExtension", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
