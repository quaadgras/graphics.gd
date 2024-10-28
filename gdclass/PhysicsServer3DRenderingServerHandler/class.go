package PhysicsServer3DRenderingServerHandler

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

type Go [1]classdb.PhysicsServer3DRenderingServerHandler

/*
Called by the [PhysicsServer3D] to set the position for the [SoftBody3D] vertex at the index specified by [param vertex_id].
[b]Note:[/b] The [param vertex] parameter used to be of type [code]const void*[/code] prior to Godot 4.2.
*/
func (Go) _set_vertex(impl func(ptr unsafe.Pointer, vertex_id int, vertex gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var vertex_id = gd.UnsafeGet[gd.Int](p_args,0)
		var vertex = gd.UnsafeGet[gd.Vector3](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(vertex_id), vertex)
	}
}

/*
Called by the [PhysicsServer3D] to set the normal for the [SoftBody3D] vertex at the index specified by [param vertex_id].
[b]Note:[/b] The [param normal] parameter used to be of type [code]const void*[/code] prior to Godot 4.2.
*/
func (Go) _set_normal(impl func(ptr unsafe.Pointer, vertex_id int, normal gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var vertex_id = gd.UnsafeGet[gd.Int](p_args,0)
		var normal = gd.UnsafeGet[gd.Vector3](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(vertex_id), normal)
	}
}

/*
Called by the [PhysicsServer3D] to set the bounding box for the [SoftBody3D].
*/
func (Go) _set_aabb(impl func(ptr unsafe.Pointer, aabb gd.AABB) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var aabb = gd.UnsafeGet[gd.AABB](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, aabb)
	}
}

/*
Sets the position for the [SoftBody3D] vertex at the index specified by [param vertex_id].
*/
func (self Go) SetVertex(vertex_id int, vertex gd.Vector3) {
	class(self).SetVertex(gd.Int(vertex_id), vertex)
}

/*
Sets the normal for the [SoftBody3D] vertex at the index specified by [param vertex_id].
*/
func (self Go) SetNormal(vertex_id int, normal gd.Vector3) {
	class(self).SetNormal(gd.Int(vertex_id), normal)
}

/*
Sets the bounding box for the [SoftBody3D].
*/
func (self Go) SetAabb(aabb gd.AABB) {
	class(self).SetAabb(aabb)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PhysicsServer3DRenderingServerHandler
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsServer3DRenderingServerHandler"))
	return Go{classdb.PhysicsServer3DRenderingServerHandler(object)}
}

/*
Called by the [PhysicsServer3D] to set the position for the [SoftBody3D] vertex at the index specified by [param vertex_id].
[b]Note:[/b] The [param vertex] parameter used to be of type [code]const void*[/code] prior to Godot 4.2.
*/
func (class) _set_vertex(impl func(ptr unsafe.Pointer, vertex_id gd.Int, vertex gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var vertex_id = gd.UnsafeGet[gd.Int](p_args,0)
		var vertex = gd.UnsafeGet[gd.Vector3](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, vertex_id, vertex)
	}
}

/*
Called by the [PhysicsServer3D] to set the normal for the [SoftBody3D] vertex at the index specified by [param vertex_id].
[b]Note:[/b] The [param normal] parameter used to be of type [code]const void*[/code] prior to Godot 4.2.
*/
func (class) _set_normal(impl func(ptr unsafe.Pointer, vertex_id gd.Int, normal gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var vertex_id = gd.UnsafeGet[gd.Int](p_args,0)
		var normal = gd.UnsafeGet[gd.Vector3](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, vertex_id, normal)
	}
}

/*
Called by the [PhysicsServer3D] to set the bounding box for the [SoftBody3D].
*/
func (class) _set_aabb(impl func(ptr unsafe.Pointer, aabb gd.AABB) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var aabb = gd.UnsafeGet[gd.AABB](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, aabb)
	}
}

/*
Sets the position for the [SoftBody3D] vertex at the index specified by [param vertex_id].
*/
//go:nosplit
func (self class) SetVertex(vertex_id gd.Int, vertex gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, vertex_id)
	callframe.Arg(frame, vertex)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsServer3DRenderingServerHandler.Bind_set_vertex, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the normal for the [SoftBody3D] vertex at the index specified by [param vertex_id].
*/
//go:nosplit
func (self class) SetNormal(vertex_id gd.Int, normal gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, vertex_id)
	callframe.Arg(frame, normal)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsServer3DRenderingServerHandler.Bind_set_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the bounding box for the [SoftBody3D].
*/
//go:nosplit
func (self class) SetAabb(aabb gd.AABB)  {
	var frame = callframe.New()
	callframe.Arg(frame, aabb)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsServer3DRenderingServerHandler.Bind_set_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsPhysicsServer3DRenderingServerHandler() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPhysicsServer3DRenderingServerHandler() Go { return *((*Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_set_vertex": return reflect.ValueOf(self._set_vertex);
	case "_set_normal": return reflect.ValueOf(self._set_normal);
	case "_set_aabb": return reflect.ValueOf(self._set_aabb);
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_set_vertex": return reflect.ValueOf(self._set_vertex);
	case "_set_normal": return reflect.ValueOf(self._set_normal);
	case "_set_aabb": return reflect.ValueOf(self._set_aabb);
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("PhysicsServer3DRenderingServerHandler", func(ptr gd.Object) any { return classdb.PhysicsServer3DRenderingServerHandler(ptr) })}
