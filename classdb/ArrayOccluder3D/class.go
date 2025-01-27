// Package ArrayOccluder3D provides methods for working with ArrayOccluder3D object instances.
package ArrayOccluder3D

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
import "graphics.gd/classdb/Occluder3D"
import "graphics.gd/classdb/Resource"
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

/*
[ArrayOccluder3D] stores an arbitrary 3D polygon shape that can be used by the engine's occlusion culling system. This is analogous to [ArrayMesh], but for occluders.
See [OccluderInstance3D]'s documentation for instructions on setting up occlusion culling.
*/
type Instance [1]gdclass.ArrayOccluder3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsArrayOccluder3D() Instance
}

/*
Sets [member indices] and [member vertices], while updating the final occluder only once after both values are set.
*/
func (self Instance) SetArrays(vertices []Vector3.XYZ, indices []int32) { //gd:ArrayOccluder3D.set_arrays
	class(self).SetArrays(gd.NewPackedVector3Slice(*(*[]gd.Vector3)(unsafe.Pointer(&vertices))), gd.NewPackedInt32Slice(indices))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ArrayOccluder3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ArrayOccluder3D"))
	casted := Instance{*(*gdclass.ArrayOccluder3D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) SetVertices(value []Vector3.XYZ) {
	class(self).SetVertices(gd.NewPackedVector3Slice(*(*[]gd.Vector3)(unsafe.Pointer(&value))))
}

func (self Instance) SetIndices(value []int32) {
	class(self).SetIndices(gd.NewPackedInt32Slice(value))
}

/*
Sets [member indices] and [member vertices], while updating the final occluder only once after both values are set.
*/
//go:nosplit
func (self class) SetArrays(vertices gd.PackedVector3Array, indices gd.PackedInt32Array) { //gd:ArrayOccluder3D.set_arrays
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(vertices))
	callframe.Arg(frame, pointers.Get(indices))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayOccluder3D.Bind_set_arrays, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetVertices(vertices gd.PackedVector3Array) { //gd:ArrayOccluder3D.set_vertices
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(vertices))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayOccluder3D.Bind_set_vertices, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetIndices(indices gd.PackedInt32Array) { //gd:ArrayOccluder3D.set_indices
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(indices))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ArrayOccluder3D.Bind_set_indices, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsArrayOccluder3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsArrayOccluder3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsOccluder3D() Occluder3D.Advanced {
	return *((*Occluder3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsOccluder3D() Occluder3D.Instance {
	return *((*Occluder3D.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(Occluder3D.Advanced(self.AsOccluder3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Occluder3D.Instance(self.AsOccluder3D()), name)
	}
}
func init() {
	gdclass.Register("ArrayOccluder3D", func(ptr gd.Object) any {
		return [1]gdclass.ArrayOccluder3D{*(*gdclass.ArrayOccluder3D)(unsafe.Pointer(&ptr))}
	})
}
