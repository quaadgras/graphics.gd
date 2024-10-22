package SkinReference

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
An internal object containing a mapping from a [Skin] used within the context of a particular [MeshInstance3D] to refer to the skeleton's [RID] in the RenderingServer.
See also [method MeshInstance3D.get_skin_reference] and [method RenderingServer.instance_attach_skeleton].
Note that despite the similar naming, the skeleton RID used in the [RenderingServer] does not have a direct one-to-one correspondence to a [Skeleton3D] node.
In particular, a [Skeleton3D] node with no [MeshInstance3D] children may be unknown to the [RenderingServer].
On the other hand, a [Skeleton3D] with multiple [MeshInstance3D] nodes which each have different [member MeshInstance3D.skin] objects may have multiple SkinReference instances (and hence, multiple skeleton [RID]s).

*/
type Go [1]classdb.SkinReference

/*
Returns the [RID] owned by this SkinReference, as returned by [method RenderingServer.skeleton_create].
*/
func (self Go) GetSkeleton() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetSkeleton())
}

/*
Returns the [Skin] connected to this SkinReference. In the case of [MeshInstance3D] with no [member MeshInstance3D.skin] assigned, this will reference an internal default [Skin] owned by that [MeshInstance3D].
Note that a single [Skin] may have more than one [SkinReference] in the case that it is shared by meshes across multiple [Skeleton3D] nodes.
*/
func (self Go) GetSkin() gdclass.Skin {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Skin(class(self).GetSkin(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SkinReference
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SkinReference"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Returns the [RID] owned by this SkinReference, as returned by [method RenderingServer.skeleton_create].
*/
//go:nosplit
func (self class) GetSkeleton() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkinReference.Bind_get_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Skin] connected to this SkinReference. In the case of [MeshInstance3D] with no [member MeshInstance3D.skin] assigned, this will reference an internal default [Skin] owned by that [MeshInstance3D].
Note that a single [Skin] may have more than one [SkinReference] in the case that it is shared by meshes across multiple [Skeleton3D] nodes.
*/
//go:nosplit
func (self class) GetSkin(ctx gd.Lifetime) gdclass.Skin {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkinReference.Bind_get_skin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Skin
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsSkinReference() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSkinReference() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("SkinReference", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
