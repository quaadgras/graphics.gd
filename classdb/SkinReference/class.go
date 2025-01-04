package SkinReference

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
An internal object containing a mapping from a [Skin] used within the context of a particular [MeshInstance3D] to refer to the skeleton's [RID] in the RenderingServer.
See also [method MeshInstance3D.get_skin_reference] and [method RenderingServer.instance_attach_skeleton].
Note that despite the similar naming, the skeleton RID used in the [RenderingServer] does not have a direct one-to-one correspondence to a [Skeleton3D] node.
In particular, a [Skeleton3D] node with no [MeshInstance3D] children may be unknown to the [RenderingServer].
On the other hand, a [Skeleton3D] with multiple [MeshInstance3D] nodes which each have different [member MeshInstance3D.skin] objects may have multiple SkinReference instances (and hence, multiple skeleton [RID]s).
*/
type Instance [1]gdclass.SkinReference
type Any interface {
	gd.IsClass
	AsSkinReference() Instance
}

/*
Returns the [RID] owned by this SkinReference, as returned by [method RenderingServer.skeleton_create].
*/
func (self Instance) GetSkeleton() Resource.ID {
	return Resource.ID(class(self).GetSkeleton())
}

/*
Returns the [Skin] connected to this SkinReference. In the case of [MeshInstance3D] with no [member MeshInstance3D.skin] assigned, this will reference an internal default [Skin] owned by that [MeshInstance3D].
Note that a single [Skin] may have more than one [SkinReference] in the case that it is shared by meshes across multiple [Skeleton3D] nodes.
*/
func (self Instance) GetSkin() [1]gdclass.Skin {
	return [1]gdclass.Skin(class(self).GetSkin())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SkinReference

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SkinReference"))
	return Instance{*(*gdclass.SkinReference)(unsafe.Pointer(&object))}
}

/*
Returns the [RID] owned by this SkinReference, as returned by [method RenderingServer.skeleton_create].
*/
//go:nosplit
func (self class) GetSkeleton() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkinReference.Bind_get_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Skin] connected to this SkinReference. In the case of [MeshInstance3D] with no [member MeshInstance3D.skin] assigned, this will reference an internal default [Skin] owned by that [MeshInstance3D].
Note that a single [Skin] may have more than one [SkinReference] in the case that it is shared by meshes across multiple [Skeleton3D] nodes.
*/
//go:nosplit
func (self class) GetSkin() [1]gdclass.Skin {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkinReference.Bind_get_skin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Skin{gd.PointerWithOwnershipTransferredToGo[gdclass.Skin](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsSkinReference() Advanced      { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSkinReference() Instance   { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	gdclass.Register("SkinReference", func(ptr gd.Object) any {
		return [1]gdclass.SkinReference{*(*gdclass.SkinReference)(unsafe.Pointer(&ptr))}
	})
}
