// Package MultiMeshInstance3D provides methods for working with MultiMeshInstance3D object instances.
package MultiMeshInstance3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/GeometryInstance3D"
import "graphics.gd/classdb/VisualInstance3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
[MultiMeshInstance3D] is a specialized node to instance [GeometryInstance3D]s based on a [MultiMesh] resource.
This is useful to optimize the rendering of a high number of instances of a given mesh (for example trees in a forest or grass strands).
*/
type Instance [1]gdclass.MultiMeshInstance3D
type Any interface {
	gd.IsClass
	AsMultiMeshInstance3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.MultiMeshInstance3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MultiMeshInstance3D"))
	return Instance{*(*gdclass.MultiMeshInstance3D)(unsafe.Pointer(&object))}
}

func (self Instance) Multimesh() [1]gdclass.MultiMesh {
	return [1]gdclass.MultiMesh(class(self).GetMultimesh())
}

func (self Instance) SetMultimesh(value [1]gdclass.MultiMesh) {
	class(self).SetMultimesh(value)
}

//go:nosplit
func (self class) SetMultimesh(multimesh [1]gdclass.MultiMesh) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(multimesh[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMeshInstance3D.Bind_set_multimesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMultimesh() [1]gdclass.MultiMesh {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMeshInstance3D.Bind_get_multimesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.MultiMesh{gd.PointerWithOwnershipTransferredToGo[gdclass.MultiMesh](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsMultiMeshInstance3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMultiMeshInstance3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsGeometryInstance3D() GeometryInstance3D.Advanced {
	return *((*GeometryInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGeometryInstance3D() GeometryInstance3D.Instance {
	return *((*GeometryInstance3D.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsGeometryInstance3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsGeometryInstance3D(), name)
	}
}
func init() {
	gdclass.Register("MultiMeshInstance3D", func(ptr gd.Object) any {
		return [1]gdclass.MultiMeshInstance3D{*(*gdclass.MultiMeshInstance3D)(unsafe.Pointer(&ptr))}
	})
}
