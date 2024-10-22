package PhysicsServer3DRenderingServerHandler

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

type Go [1]classdb.PhysicsServer3DRenderingServerHandler

/*
Called by the [PhysicsServer3D] to set the position for the [SoftBody3D] vertex at the index specified by [param vertex_id].
[b]Note:[/b] The [param vertex] parameter used to be of type [code]const void*[/code] prior to Godot 4.2.
*/
func (Go) _set_vertex(impl func(ptr unsafe.Pointer, vertex_id int, vertex gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var vertex_id = gd.UnsafeGet[gd.Int](p_args,0)
		var vertex = gd.UnsafeGet[gd.Vector3](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(vertex_id), vertex)
		gc.End()
	}
}

/*
Called by the [PhysicsServer3D] to set the normal for the [SoftBody3D] vertex at the index specified by [param vertex_id].
[b]Note:[/b] The [param normal] parameter used to be of type [code]const void*[/code] prior to Godot 4.2.
*/
func (Go) _set_normal(impl func(ptr unsafe.Pointer, vertex_id int, normal gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var vertex_id = gd.UnsafeGet[gd.Int](p_args,0)
		var normal = gd.UnsafeGet[gd.Vector3](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(vertex_id), normal)
		gc.End()
	}
}

/*
Called by the [PhysicsServer3D] to set the bounding box for the [SoftBody3D].
*/
func (Go) _set_aabb(impl func(ptr unsafe.Pointer, aabb gd.AABB) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var aabb = gd.UnsafeGet[gd.AABB](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, aabb)
		gc.End()
	}
}

/*
Sets the position for the [SoftBody3D] vertex at the index specified by [param vertex_id].
*/
func (self Go) SetVertex(vertex_id int, vertex gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVertex(gd.Int(vertex_id), vertex)
}

/*
Sets the normal for the [SoftBody3D] vertex at the index specified by [param vertex_id].
*/
func (self Go) SetNormal(vertex_id int, normal gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNormal(gd.Int(vertex_id), normal)
}

/*
Sets the bounding box for the [SoftBody3D].
*/
func (self Go) SetAabb(aabb gd.AABB) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAabb(aabb)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PhysicsServer3DRenderingServerHandler
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("PhysicsServer3DRenderingServerHandler"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Called by the [PhysicsServer3D] to set the position for the [SoftBody3D] vertex at the index specified by [param vertex_id].
[b]Note:[/b] The [param vertex] parameter used to be of type [code]const void*[/code] prior to Godot 4.2.
*/
func (class) _set_vertex(impl func(ptr unsafe.Pointer, vertex_id gd.Int, vertex gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var vertex_id = gd.UnsafeGet[gd.Int](p_args,0)
		var vertex = gd.UnsafeGet[gd.Vector3](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, vertex_id, vertex)
		ctx.End()
	}
}

/*
Called by the [PhysicsServer3D] to set the normal for the [SoftBody3D] vertex at the index specified by [param vertex_id].
[b]Note:[/b] The [param normal] parameter used to be of type [code]const void*[/code] prior to Godot 4.2.
*/
func (class) _set_normal(impl func(ptr unsafe.Pointer, vertex_id gd.Int, normal gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var vertex_id = gd.UnsafeGet[gd.Int](p_args,0)
		var normal = gd.UnsafeGet[gd.Vector3](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, vertex_id, normal)
		ctx.End()
	}
}

/*
Called by the [PhysicsServer3D] to set the bounding box for the [SoftBody3D].
*/
func (class) _set_aabb(impl func(ptr unsafe.Pointer, aabb gd.AABB) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var aabb = gd.UnsafeGet[gd.AABB](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, aabb)
		ctx.End()
	}
}

/*
Sets the position for the [SoftBody3D] vertex at the index specified by [param vertex_id].
*/
//go:nosplit
func (self class) SetVertex(vertex_id gd.Int, vertex gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, vertex_id)
	callframe.Arg(frame, vertex)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3DRenderingServerHandler.Bind_set_vertex, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the normal for the [SoftBody3D] vertex at the index specified by [param vertex_id].
*/
//go:nosplit
func (self class) SetNormal(vertex_id gd.Int, normal gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, vertex_id)
	callframe.Arg(frame, normal)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3DRenderingServerHandler.Bind_set_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the bounding box for the [SoftBody3D].
*/
//go:nosplit
func (self class) SetAabb(aabb gd.AABB)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, aabb)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3DRenderingServerHandler.Bind_set_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func init() {classdb.Register("PhysicsServer3DRenderingServerHandler", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
