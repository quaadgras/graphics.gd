// Package OmniLight3D provides methods for working with OmniLight3D object instances.
package OmniLight3D

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
import "graphics.gd/classdb/Light3D"
import "graphics.gd/classdb/VisualInstance3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"

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
An Omnidirectional light is a type of [Light3D] that emits light in all directions. The light is attenuated by distance and this attenuation can be configured by changing its energy, radius, and attenuation parameters.
[b]Note:[/b] When using the Mobile rendering method, only 8 omni lights can be displayed on each mesh resource. Attempting to display more than 8 omni lights on a single mesh resource will result in omni lights flickering in and out as the camera moves. When using the Compatibility rendering method, only 8 omni lights can be displayed on each mesh resource by default, but this can be increased by adjusting [member ProjectSettings.rendering/limits/opengl/max_lights_per_object].
[b]Note:[/b] When using the Mobile or Compatibility rendering methods, omni lights will only correctly affect meshes whose visibility AABB intersects with the light's AABB. If using a shader to deform the mesh in a way that makes it go outside its AABB, [member GeometryInstance3D.extra_cull_margin] must be increased on the mesh. Otherwise, the light may not be visible on the mesh.
*/
type Instance [1]gdclass.OmniLight3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsOmniLight3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.OmniLight3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OmniLight3D"))
	casted := Instance{*(*gdclass.OmniLight3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) OmniShadowMode() gdclass.OmniLight3DShadowMode {
	return gdclass.OmniLight3DShadowMode(class(self).GetShadowMode())
}

func (self Instance) SetOmniShadowMode(value gdclass.OmniLight3DShadowMode) {
	class(self).SetShadowMode(value)
}

//go:nosplit
func (self class) SetShadowMode(mode gdclass.OmniLight3DShadowMode) { //gd:OmniLight3D.set_shadow_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OmniLight3D.Bind_set_shadow_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetShadowMode() gdclass.OmniLight3DShadowMode { //gd:OmniLight3D.get_shadow_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.OmniLight3DShadowMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OmniLight3D.Bind_get_shadow_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsOmniLight3D() Advanced     { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsOmniLight3D() Instance  { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsLight3D() Light3D.Advanced { return *((*Light3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsLight3D() Light3D.Instance {
	return *((*Light3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualInstance3D() VisualInstance3D.Advanced {
	return *((*VisualInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualInstance3D() VisualInstance3D.Instance {
	return *((*VisualInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Light3D.Advanced(self.AsLight3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Light3D.Instance(self.AsLight3D()), name)
	}
}
func init() {
	gdclass.Register("OmniLight3D", func(ptr gd.Object) any { return [1]gdclass.OmniLight3D{*(*gdclass.OmniLight3D)(unsafe.Pointer(&ptr))} })
}

type ShadowMode = gdclass.OmniLight3DShadowMode //gd:OmniLight3D.ShadowMode

const (
	/*Shadows are rendered to a dual-paraboloid texture. Faster than [constant SHADOW_CUBE], but lower-quality.*/
	ShadowDualParaboloid ShadowMode = 0
	/*Shadows are rendered to a cubemap. Slower than [constant SHADOW_DUAL_PARABOLOID], but higher-quality.*/
	ShadowCube ShadowMode = 1
)
