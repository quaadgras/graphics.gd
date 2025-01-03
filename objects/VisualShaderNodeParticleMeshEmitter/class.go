package VisualShaderNodeParticleMeshEmitter

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/VisualShaderNodeParticleEmitter"
import "graphics.gd/objects/VisualShaderNode"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
[VisualShaderNodeParticleEmitter] that makes the particles emitted in a shape of the assigned [member mesh]. It will emit from the mesh's surfaces, either all or only the specified one.
*/
type Instance [1]classdb.VisualShaderNodeParticleMeshEmitter
type Any interface {
	gd.IsClass
	AsVisualShaderNodeParticleMeshEmitter() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VisualShaderNodeParticleMeshEmitter

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeParticleMeshEmitter"))
	return Instance{*(*classdb.VisualShaderNodeParticleMeshEmitter)(unsafe.Pointer(&object))}
}

func (self Instance) Mesh() objects.Mesh {
	return objects.Mesh(class(self).GetMesh())
}

func (self Instance) SetMesh(value objects.Mesh) {
	class(self).SetMesh(value)
}

func (self Instance) UseAllSurfaces() bool {
	return bool(class(self).IsUseAllSurfaces())
}

func (self Instance) SetUseAllSurfaces(value bool) {
	class(self).SetUseAllSurfaces(value)
}

func (self Instance) SurfaceIndex() int {
	return int(int(class(self).GetSurfaceIndex()))
}

func (self Instance) SetSurfaceIndex(value int) {
	class(self).SetSurfaceIndex(gd.Int(value))
}

//go:nosplit
func (self class) SetMesh(mesh objects.Mesh) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleMeshEmitter.Bind_set_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMesh() objects.Mesh {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleMeshEmitter.Bind_get_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Mesh{gd.PointerWithOwnershipTransferredToGo[classdb.Mesh](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseAllSurfaces(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleMeshEmitter.Bind_set_use_all_surfaces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsUseAllSurfaces() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleMeshEmitter.Bind_is_use_all_surfaces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSurfaceIndex(surface_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, surface_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleMeshEmitter.Bind_set_surface_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSurfaceIndex() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleMeshEmitter.Bind_get_surface_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeParticleMeshEmitter() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeParticleMeshEmitter() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNodeParticleEmitter() VisualShaderNodeParticleEmitter.Advanced {
	return *((*VisualShaderNodeParticleEmitter.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeParticleEmitter() VisualShaderNodeParticleEmitter.Instance {
	return *((*VisualShaderNodeParticleEmitter.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNode() VisualShaderNode.Advanced {
	return *((*VisualShaderNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNode() VisualShaderNode.Instance {
	return *((*VisualShaderNode.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNodeParticleEmitter(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNodeParticleEmitter(), name)
	}
}
func init() {
	classdb.Register("VisualShaderNodeParticleMeshEmitter", func(ptr gd.Object) any {
		return [1]classdb.VisualShaderNodeParticleMeshEmitter{*(*classdb.VisualShaderNodeParticleMeshEmitter)(unsafe.Pointer(&ptr))}
	})
}
