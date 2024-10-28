package GLTFPhysicsBody

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Represents a physics body as an intermediary between the [code]OMI_physics_body[/code] GLTF data and Godot's nodes, and it's abstracted in a way that allows adding support for different GLTF physics extensions in the future.

*/
type Go [1]classdb.GLTFPhysicsBody

/*
Creates a new GLTFPhysicsBody instance from the given Godot [CollisionObject3D] node.
*/
func (self Go) FromNode(body_node gdclass.CollisionObject3D) gdclass.GLTFPhysicsBody {
	return gdclass.GLTFPhysicsBody(class(self).FromNode(body_node))
}

/*
Converts this GLTFPhysicsBody instance into a Godot [CollisionObject3D] node.
*/
func (self Go) ToNode() gdclass.CollisionObject3D {
	return gdclass.CollisionObject3D(class(self).ToNode())
}

/*
Creates a new GLTFPhysicsBody instance by parsing the given [Dictionary] in the [code]OMI_physics_body[/code] GLTF extension format.
*/
func (self Go) FromDictionary(dictionary gd.Dictionary) gdclass.GLTFPhysicsBody {
	return gdclass.GLTFPhysicsBody(class(self).FromDictionary(dictionary))
}

/*
Serializes this GLTFPhysicsBody instance into a [Dictionary]. It will be in the format expected by the [code]OMI_physics_body[/code] GLTF extension.
*/
func (self Go) ToDictionary() gd.Dictionary {
	return gd.Dictionary(class(self).ToDictionary())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GLTFPhysicsBody
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFPhysicsBody"))
	return Go{classdb.GLTFPhysicsBody(object)}
}

func (self Go) BodyType() string {
		return string(class(self).GetBodyType().String())
}

func (self Go) SetBodyType(value string) {
	class(self).SetBodyType(gd.NewString(value))
}

func (self Go) Mass() float64 {
		return float64(float64(class(self).GetMass()))
}

func (self Go) SetMass(value float64) {
	class(self).SetMass(gd.Float(value))
}

func (self Go) LinearVelocity() gd.Vector3 {
		return gd.Vector3(class(self).GetLinearVelocity())
}

func (self Go) SetLinearVelocity(value gd.Vector3) {
	class(self).SetLinearVelocity(value)
}

func (self Go) AngularVelocity() gd.Vector3 {
		return gd.Vector3(class(self).GetAngularVelocity())
}

func (self Go) SetAngularVelocity(value gd.Vector3) {
	class(self).SetAngularVelocity(value)
}

func (self Go) CenterOfMass() gd.Vector3 {
		return gd.Vector3(class(self).GetCenterOfMass())
}

func (self Go) SetCenterOfMass(value gd.Vector3) {
	class(self).SetCenterOfMass(value)
}

func (self Go) InertiaDiagonal() gd.Vector3 {
		return gd.Vector3(class(self).GetInertiaDiagonal())
}

func (self Go) SetInertiaDiagonal(value gd.Vector3) {
	class(self).SetInertiaDiagonal(value)
}

func (self Go) InertiaOrientation() gd.Quaternion {
		return gd.Quaternion(class(self).GetInertiaOrientation())
}

func (self Go) SetInertiaOrientation(value gd.Quaternion) {
	class(self).SetInertiaOrientation(value)
}

func (self Go) InertiaTensor() gd.Basis {
		return gd.Basis(class(self).GetInertiaTensor())
}

func (self Go) SetInertiaTensor(value gd.Basis) {
	class(self).SetInertiaTensor(value)
}

/*
Creates a new GLTFPhysicsBody instance from the given Godot [CollisionObject3D] node.
*/
//go:nosplit
func (self class) FromNode(body_node gdclass.CollisionObject3D) gdclass.GLTFPhysicsBody {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(body_node[0])[0])
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_from_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.GLTFPhysicsBody{classdb.GLTFPhysicsBody(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Converts this GLTFPhysicsBody instance into a Godot [CollisionObject3D] node.
*/
//go:nosplit
func (self class) ToNode() gdclass.CollisionObject3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_to_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.CollisionObject3D{classdb.CollisionObject3D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Creates a new GLTFPhysicsBody instance by parsing the given [Dictionary] in the [code]OMI_physics_body[/code] GLTF extension format.
*/
//go:nosplit
func (self class) FromDictionary(dictionary gd.Dictionary) gdclass.GLTFPhysicsBody {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(dictionary))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_from_dictionary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.GLTFPhysicsBody{classdb.GLTFPhysicsBody(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Serializes this GLTFPhysicsBody instance into a [Dictionary]. It will be in the format expected by the [code]OMI_physics_body[/code] GLTF extension.
*/
//go:nosplit
func (self class) ToDictionary() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_to_dictionary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetBodyType() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_get_body_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBodyType(body_type gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(body_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_set_body_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMass() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_get_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMass(mass gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, mass)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_set_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLinearVelocity() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_get_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLinearVelocity(linear_velocity gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, linear_velocity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_set_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAngularVelocity() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_get_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAngularVelocity(angular_velocity gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, angular_velocity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_set_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCenterOfMass() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_get_center_of_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCenterOfMass(center_of_mass gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, center_of_mass)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_set_center_of_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInertiaDiagonal() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_get_inertia_diagonal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInertiaDiagonal(inertia_diagonal gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, inertia_diagonal)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_set_inertia_diagonal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInertiaOrientation() gd.Quaternion {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Quaternion](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_get_inertia_orientation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInertiaOrientation(inertia_orientation gd.Quaternion)  {
	var frame = callframe.New()
	callframe.Arg(frame, inertia_orientation)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_set_inertia_orientation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInertiaTensor() gd.Basis {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Basis](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_get_inertia_tensor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInertiaTensor(inertia_tensor gd.Basis)  {
	var frame = callframe.New()
	callframe.Arg(frame, inertia_tensor)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_set_inertia_tensor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsGLTFPhysicsBody() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGLTFPhysicsBody() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("GLTFPhysicsBody", func(ptr gd.Object) any { return classdb.GLTFPhysicsBody(ptr) })}
