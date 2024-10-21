package GLTFPhysicsBody

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Represents a physics body as an intermediary between the [code]OMI_physics_body[/code] GLTF data and Godot's nodes, and it's abstracted in a way that allows adding support for different GLTF physics extensions in the future.

*/
type Simple [1]classdb.GLTFPhysicsBody
func (self Simple) FromNode(body_node [1]classdb.CollisionObject3D) [1]classdb.GLTFPhysicsBody {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.GLTFPhysicsBody(Expert(self).FromNode(gc, body_node))
}
func (self Simple) ToNode() [1]classdb.CollisionObject3D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.CollisionObject3D(Expert(self).ToNode(gc))
}
func (self Simple) FromDictionary(dictionary gd.Dictionary) [1]classdb.GLTFPhysicsBody {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.GLTFPhysicsBody(Expert(self).FromDictionary(gc, dictionary))
}
func (self Simple) ToDictionary() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).ToDictionary(gc))
}
func (self Simple) GetBodyType() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetBodyType(gc).String())
}
func (self Simple) SetBodyType(body_type string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBodyType(gc.String(body_type))
}
func (self Simple) GetMass() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMass()))
}
func (self Simple) SetMass(mass float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMass(gd.Float(mass))
}
func (self Simple) GetLinearVelocity() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetLinearVelocity())
}
func (self Simple) SetLinearVelocity(linear_velocity gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLinearVelocity(linear_velocity)
}
func (self Simple) GetAngularVelocity() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetAngularVelocity())
}
func (self Simple) SetAngularVelocity(angular_velocity gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAngularVelocity(angular_velocity)
}
func (self Simple) GetCenterOfMass() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetCenterOfMass())
}
func (self Simple) SetCenterOfMass(center_of_mass gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCenterOfMass(center_of_mass)
}
func (self Simple) GetInertiaDiagonal() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetInertiaDiagonal())
}
func (self Simple) SetInertiaDiagonal(inertia_diagonal gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInertiaDiagonal(inertia_diagonal)
}
func (self Simple) GetInertiaOrientation() gd.Quaternion {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Quaternion(Expert(self).GetInertiaOrientation())
}
func (self Simple) SetInertiaOrientation(inertia_orientation gd.Quaternion) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInertiaOrientation(inertia_orientation)
}
func (self Simple) GetInertiaTensor() gd.Basis {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Basis(Expert(self).GetInertiaTensor())
}
func (self Simple) SetInertiaTensor(inertia_tensor gd.Basis) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInertiaTensor(inertia_tensor)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.GLTFPhysicsBody
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Creates a new GLTFPhysicsBody instance from the given Godot [CollisionObject3D] node.
*/
//go:nosplit
func (self class) FromNode(ctx gd.Lifetime, body_node object.CollisionObject3D) object.GLTFPhysicsBody {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(body_node[0].AsPointer())[0])
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.GLTFPhysicsBody.Bind_from_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.GLTFPhysicsBody
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Converts this GLTFPhysicsBody instance into a Godot [CollisionObject3D] node.
*/
//go:nosplit
func (self class) ToNode(ctx gd.Lifetime) object.CollisionObject3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsBody.Bind_to_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.CollisionObject3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates a new GLTFPhysicsBody instance by parsing the given [Dictionary] in the [code]OMI_physics_body[/code] GLTF extension format.
*/
//go:nosplit
func (self class) FromDictionary(ctx gd.Lifetime, dictionary gd.Dictionary) object.GLTFPhysicsBody {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(dictionary))
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.GLTFPhysicsBody.Bind_from_dictionary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.GLTFPhysicsBody
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Serializes this GLTFPhysicsBody instance into a [Dictionary]. It will be in the format expected by the [code]OMI_physics_body[/code] GLTF extension.
*/
//go:nosplit
func (self class) ToDictionary(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsBody.Bind_to_dictionary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetBodyType(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsBody.Bind_get_body_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBodyType(body_type gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(body_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsBody.Bind_set_body_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMass() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsBody.Bind_get_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMass(mass gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mass)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsBody.Bind_set_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLinearVelocity() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsBody.Bind_get_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLinearVelocity(linear_velocity gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, linear_velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsBody.Bind_set_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAngularVelocity() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsBody.Bind_get_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAngularVelocity(angular_velocity gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, angular_velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsBody.Bind_set_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCenterOfMass() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsBody.Bind_get_center_of_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCenterOfMass(center_of_mass gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, center_of_mass)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsBody.Bind_set_center_of_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInertiaDiagonal() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsBody.Bind_get_inertia_diagonal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInertiaDiagonal(inertia_diagonal gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, inertia_diagonal)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsBody.Bind_set_inertia_diagonal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInertiaOrientation() gd.Quaternion {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Quaternion](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsBody.Bind_get_inertia_orientation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInertiaOrientation(inertia_orientation gd.Quaternion)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, inertia_orientation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsBody.Bind_set_inertia_orientation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInertiaTensor() gd.Basis {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Basis](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsBody.Bind_get_inertia_tensor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInertiaTensor(inertia_tensor gd.Basis)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, inertia_tensor)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFPhysicsBody.Bind_set_inertia_tensor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsGLTFPhysicsBody() Expert { return self[0].AsGLTFPhysicsBody() }


//go:nosplit
func (self Simple) AsGLTFPhysicsBody() Simple { return self[0].AsGLTFPhysicsBody() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("GLTFPhysicsBody", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
