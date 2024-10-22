package VisualShaderNodeParticleMeshEmitter

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualShaderNodeParticleEmitter"
import "grow.graphics/gd/gdclass/VisualShaderNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[VisualShaderNodeParticleEmitter] that makes the particles emitted in a shape of the assigned [member mesh]. It will emit from the mesh's surfaces, either all or only the specified one.

*/
type Go [1]classdb.VisualShaderNodeParticleMeshEmitter
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeParticleMeshEmitter
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VisualShaderNodeParticleMeshEmitter"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Mesh() gdclass.Mesh {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Mesh(class(self).GetMesh(gc))
}

func (self Go) SetMesh(value gdclass.Mesh) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMesh(value)
}

func (self Go) UseAllSurfaces() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsUseAllSurfaces())
}

func (self Go) SetUseAllSurfaces(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseAllSurfaces(value)
}

func (self Go) SurfaceIndex() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSurfaceIndex()))
}

func (self Go) SetSurfaceIndex(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSurfaceIndex(gd.Int(value))
}

//go:nosplit
func (self class) SetMesh(mesh gdclass.Mesh)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(mesh[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeParticleMeshEmitter.Bind_set_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMesh(ctx gd.Lifetime) gdclass.Mesh {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeParticleMeshEmitter.Bind_get_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Mesh
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseAllSurfaces(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeParticleMeshEmitter.Bind_set_use_all_surfaces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUseAllSurfaces() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeParticleMeshEmitter.Bind_is_use_all_surfaces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSurfaceIndex(surface_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, surface_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeParticleMeshEmitter.Bind_set_surface_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSurfaceIndex() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeParticleMeshEmitter.Bind_get_surface_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeParticleMeshEmitter() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeParticleMeshEmitter() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNodeParticleEmitter() VisualShaderNodeParticleEmitter.GD { return *((*VisualShaderNodeParticleEmitter.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeParticleEmitter() VisualShaderNodeParticleEmitter.Go { return *((*VisualShaderNodeParticleEmitter.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNode() VisualShaderNode.GD { return *((*VisualShaderNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNode() VisualShaderNode.Go { return *((*VisualShaderNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNodeParticleEmitter(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNodeParticleEmitter(), name)
	}
}
func init() {classdb.Register("VisualShaderNodeParticleMeshEmitter", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
