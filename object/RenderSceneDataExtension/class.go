package RenderSceneDataExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/RenderSceneData"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class allows for a RenderSceneData implementation to be made in GDExtension.
	// RenderSceneDataExtension methods that can be overridden by a [Class] that extends it.
	type RenderSceneDataExtension interface {
		//Implement this in GDExtension to return the camera [Transform3D].
		GetCamTransform() gd.Transform3D
		//Implement this in GDExtension to return the camera [Projection].
		GetCamProjection() gd.Projection
		//Implement this in GDExtension to return the view count.
		GetViewCount() gd.Int
		//Implement this in GDExtension to return the eye offset for the given [param view].
		GetViewEyeOffset(view gd.Int) gd.Vector3
		//Implement this in GDExtension to return the view [Projection] for the given [param view].
		GetViewProjection(view gd.Int) gd.Projection
		//Implement this in GDExtension to return the [RID] of the uniform buffer containing the scene data as a UBO.
		GetUniformBuffer() gd.RID
	}

*/
type Simple [1]classdb.RenderSceneDataExtension
func (Simple) _get_cam_transform(impl func(ptr unsafe.Pointer) gd.Transform3D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _get_cam_projection(impl func(ptr unsafe.Pointer) gd.Projection, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _get_view_count(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _get_view_eye_offset(impl func(ptr unsafe.Pointer, view int) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var view = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(view))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _get_view_projection(impl func(ptr unsafe.Pointer, view int) gd.Projection, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var view = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(view))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _get_uniform_buffer(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
type class [1]classdb.RenderSceneDataExtension
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Implement this in GDExtension to return the camera [Transform3D].
*/
func (class) _get_cam_transform(impl func(ptr unsafe.Pointer) gd.Transform3D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Implement this in GDExtension to return the camera [Projection].
*/
func (class) _get_cam_projection(impl func(ptr unsafe.Pointer) gd.Projection, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Implement this in GDExtension to return the view count.
*/
func (class) _get_view_count(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Implement this in GDExtension to return the eye offset for the given [param view].
*/
func (class) _get_view_eye_offset(impl func(ptr unsafe.Pointer, view gd.Int) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var view = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, view)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Implement this in GDExtension to return the view [Projection] for the given [param view].
*/
func (class) _get_view_projection(impl func(ptr unsafe.Pointer, view gd.Int) gd.Projection, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var view = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, view)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Implement this in GDExtension to return the [RID] of the uniform buffer containing the scene data as a UBO.
*/
func (class) _get_uniform_buffer(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
func (self class) AsRenderSceneDataExtension() Expert { return self[0].AsRenderSceneDataExtension() }


//go:nosplit
func (self Simple) AsRenderSceneDataExtension() Simple { return self[0].AsRenderSceneDataExtension() }


//go:nosplit
func (self class) AsRenderSceneData() RenderSceneData.Expert { return self[0].AsRenderSceneData() }


//go:nosplit
func (self Simple) AsRenderSceneData() RenderSceneData.Simple { return self[0].AsRenderSceneData() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_cam_transform": return reflect.ValueOf(self._get_cam_transform);
	case "_get_cam_projection": return reflect.ValueOf(self._get_cam_projection);
	case "_get_view_count": return reflect.ValueOf(self._get_view_count);
	case "_get_view_eye_offset": return reflect.ValueOf(self._get_view_eye_offset);
	case "_get_view_projection": return reflect.ValueOf(self._get_view_projection);
	case "_get_uniform_buffer": return reflect.ValueOf(self._get_uniform_buffer);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_get_cam_transform": return reflect.ValueOf(self._get_cam_transform);
	case "_get_cam_projection": return reflect.ValueOf(self._get_cam_projection);
	case "_get_view_count": return reflect.ValueOf(self._get_view_count);
	case "_get_view_eye_offset": return reflect.ValueOf(self._get_view_eye_offset);
	case "_get_view_projection": return reflect.ValueOf(self._get_view_projection);
	case "_get_uniform_buffer": return reflect.ValueOf(self._get_uniform_buffer);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("RenderSceneDataExtension", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
