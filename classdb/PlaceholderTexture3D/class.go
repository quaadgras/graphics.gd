// Package PlaceholderTexture3D provides methods for working with PlaceholderTexture3D object instances.
package PlaceholderTexture3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Texture3D"
import "graphics.gd/classdb/Texture"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector3i"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This class is used when loading a project that uses a [Texture3D] subclass in 2 conditions:
- When running the project exported in dedicated server mode, only the texture's dimensions are kept (as they may be relied upon for gameplay purposes or positioning of other elements). This allows reducing the exported PCK's size significantly.
- When this subclass is missing due to using a different engine version or build (e.g. modules disabled).
[b]Note:[/b] This is not intended to be used as an actual texture for rendering. It is not guaranteed to work like one in shaders or materials (for example when calculating UV).
*/
type Instance [1]gdclass.PlaceholderTexture3D
type Any interface {
	gd.IsClass
	AsPlaceholderTexture3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PlaceholderTexture3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PlaceholderTexture3D"))
	return Instance{*(*gdclass.PlaceholderTexture3D)(unsafe.Pointer(&object))}
}

func (self Instance) Size() Vector3i.XYZ {
	return Vector3i.XYZ(class(self).GetSize())
}

func (self Instance) SetSize(value Vector3i.XYZ) {
	class(self).SetSize(gd.Vector3i(value))
}

//go:nosplit
func (self class) SetSize(size gd.Vector3i) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PlaceholderTexture3D.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSize() gd.Vector3i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PlaceholderTexture3D.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPlaceholderTexture3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPlaceholderTexture3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTexture3D() Texture3D.Advanced {
	return *((*Texture3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTexture3D() Texture3D.Instance {
	return *((*Texture3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsTexture() Texture.Advanced { return *((*Texture.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTexture() Texture.Instance {
	return *((*Texture.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(Texture3D.Advanced(self.AsTexture3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Texture3D.Instance(self.AsTexture3D()), name)
	}
}
func init() {
	gdclass.Register("PlaceholderTexture3D", func(ptr gd.Object) any {
		return [1]gdclass.PlaceholderTexture3D{*(*gdclass.PlaceholderTexture3D)(unsafe.Pointer(&ptr))}
	})
}
