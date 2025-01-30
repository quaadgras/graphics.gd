// Package RenderSceneDataExtension provides methods for working with RenderSceneDataExtension object instances.
package RenderSceneDataExtension

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/RenderSceneData"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Projection"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Vector3"

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
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
This class allows for a RenderSceneData implementation to be made in GDExtension.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=RenderSceneDataExtension)
*/
type Instance [1]gdclass.RenderSceneDataExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRenderSceneDataExtension() Instance
}
type Interface interface {
	//Implement this in GDExtension to return the camera [Transform3D].
	GetCamTransform() Transform3D.BasisOrigin
	//Implement this in GDExtension to return the camera [Projection].
	GetCamProjection() Projection.XYZW
	//Implement this in GDExtension to return the view count.
	GetViewCount() int
	//Implement this in GDExtension to return the eye offset for the given [param view].
	GetViewEyeOffset(view int) Vector3.XYZ
	//Implement this in GDExtension to return the view [Projection] for the given [param view].
	GetViewProjection(view int) Projection.XYZW
	//Implement this in GDExtension to return the [RID] of the uniform buffer containing the scene data as a UBO.
	GetUniformBuffer() RID.Any
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetCamTransform() (_ Transform3D.BasisOrigin)   { return }
func (self implementation) GetCamProjection() (_ Projection.XYZW)          { return }
func (self implementation) GetViewCount() (_ int)                          { return }
func (self implementation) GetViewEyeOffset(view int) (_ Vector3.XYZ)      { return }
func (self implementation) GetViewProjection(view int) (_ Projection.XYZW) { return }
func (self implementation) GetUniformBuffer() (_ RID.Any)                  { return }

/*
Implement this in GDExtension to return the camera [Transform3D].
*/
func (Instance) _get_cam_transform(impl func(ptr unsafe.Pointer) Transform3D.BasisOrigin) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, Transform3D.BasisOrigin(ret))
	}
}

/*
Implement this in GDExtension to return the camera [Projection].
*/
func (Instance) _get_cam_projection(impl func(ptr unsafe.Pointer) Projection.XYZW) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement this in GDExtension to return the view count.
*/
func (Instance) _get_view_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
Implement this in GDExtension to return the eye offset for the given [param view].
*/
func (Instance) _get_view_eye_offset(impl func(ptr unsafe.Pointer, view int) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var view = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(view))
		gd.UnsafeSet(p_back, Vector3.XYZ(ret))
	}
}

/*
Implement this in GDExtension to return the view [Projection] for the given [param view].
*/
func (Instance) _get_view_projection(impl func(ptr unsafe.Pointer, view int) Projection.XYZW) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var view = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(view))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement this in GDExtension to return the [RID] of the uniform buffer containing the scene data as a UBO.
*/
func (Instance) _get_uniform_buffer(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RenderSceneDataExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RenderSceneDataExtension"))
	casted := Instance{*(*gdclass.RenderSceneDataExtension)(unsafe.Pointer(&object))}
	return casted
}

/*
Implement this in GDExtension to return the camera [Transform3D].
*/
func (class) _get_cam_transform(impl func(ptr unsafe.Pointer) Transform3D.BasisOrigin) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement this in GDExtension to return the camera [Projection].
*/
func (class) _get_cam_projection(impl func(ptr unsafe.Pointer) Projection.XYZW) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement this in GDExtension to return the view count.
*/
func (class) _get_view_count(impl func(ptr unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement this in GDExtension to return the eye offset for the given [param view].
*/
func (class) _get_view_eye_offset(impl func(ptr unsafe.Pointer, view int64) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var view = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, view)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement this in GDExtension to return the view [Projection] for the given [param view].
*/
func (class) _get_view_projection(impl func(ptr unsafe.Pointer, view int64) Projection.XYZW) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var view = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, view)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement this in GDExtension to return the [RID] of the uniform buffer containing the scene data as a UBO.
*/
func (class) _get_uniform_buffer(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsRenderSceneDataExtension() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRenderSceneDataExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRenderSceneData() RenderSceneData.Advanced {
	return *((*RenderSceneData.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRenderSceneData() RenderSceneData.Instance {
	return *((*RenderSceneData.Instance)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_cam_transform":
		return reflect.ValueOf(self._get_cam_transform)
	case "_get_cam_projection":
		return reflect.ValueOf(self._get_cam_projection)
	case "_get_view_count":
		return reflect.ValueOf(self._get_view_count)
	case "_get_view_eye_offset":
		return reflect.ValueOf(self._get_view_eye_offset)
	case "_get_view_projection":
		return reflect.ValueOf(self._get_view_projection)
	case "_get_uniform_buffer":
		return reflect.ValueOf(self._get_uniform_buffer)
	default:
		return gd.VirtualByName(RenderSceneData.Advanced(self.AsRenderSceneData()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_cam_transform":
		return reflect.ValueOf(self._get_cam_transform)
	case "_get_cam_projection":
		return reflect.ValueOf(self._get_cam_projection)
	case "_get_view_count":
		return reflect.ValueOf(self._get_view_count)
	case "_get_view_eye_offset":
		return reflect.ValueOf(self._get_view_eye_offset)
	case "_get_view_projection":
		return reflect.ValueOf(self._get_view_projection)
	case "_get_uniform_buffer":
		return reflect.ValueOf(self._get_uniform_buffer)
	default:
		return gd.VirtualByName(RenderSceneData.Instance(self.AsRenderSceneData()), name)
	}
}
func init() {
	gdclass.Register("RenderSceneDataExtension", func(ptr gd.Object) any {
		return [1]gdclass.RenderSceneDataExtension{*(*gdclass.RenderSceneDataExtension)(unsafe.Pointer(&ptr))}
	})
}
