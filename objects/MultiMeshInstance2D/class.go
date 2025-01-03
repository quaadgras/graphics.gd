package MultiMeshInstance2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Node2D"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
[MultiMeshInstance2D] is a specialized node to instance a [MultiMesh] resource in 2D.
Usage is the same as [MultiMeshInstance3D].
*/
type Instance [1]classdb.MultiMeshInstance2D
type Any interface {
	gd.IsClass
	AsMultiMeshInstance2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.MultiMeshInstance2D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MultiMeshInstance2D"))
	return Instance{*(*classdb.MultiMeshInstance2D)(unsafe.Pointer(&object))}
}

func (self Instance) Multimesh() objects.MultiMesh {
	return objects.MultiMesh(class(self).GetMultimesh())
}

func (self Instance) SetMultimesh(value objects.MultiMesh) {
	class(self).SetMultimesh(value)
}

func (self Instance) Texture() objects.Texture2D {
	return objects.Texture2D(class(self).GetTexture())
}

func (self Instance) SetTexture(value objects.Texture2D) {
	class(self).SetTexture(value)
}

//go:nosplit
func (self class) SetMultimesh(multimesh objects.MultiMesh) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(multimesh[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMeshInstance2D.Bind_set_multimesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMultimesh() objects.MultiMesh {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMeshInstance2D.Bind_get_multimesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.MultiMesh{gd.PointerWithOwnershipTransferredToGo[classdb.MultiMesh](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTexture(texture objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMeshInstance2D.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() objects.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMeshInstance2D.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{gd.PointerWithOwnershipTransferredToGo[classdb.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}
func (self Instance) OnTextureChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("texture_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsMultiMeshInstance2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMultiMeshInstance2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.Advanced          { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance       { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() {
	classdb.Register("MultiMeshInstance2D", func(ptr gd.Object) any {
		return [1]classdb.MultiMeshInstance2D{*(*classdb.MultiMeshInstance2D)(unsafe.Pointer(&ptr))}
	})
}
