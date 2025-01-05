// Package VoxelGI provides methods for working with VoxelGI object instances.
package VoxelGI

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/VisualInstance3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector3"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
[VoxelGI]s are used to provide high-quality real-time indirect light and reflections to scenes. They precompute the effect of objects that emit light and the effect of static geometry to simulate the behavior of complex light in real-time. [VoxelGI]s need to be baked before having a visible effect. However, once baked, dynamic objects will receive light from them. Furthermore, lights can be fully dynamic or baked.
[b]Note:[/b] [VoxelGI] is only supported in the Forward+ rendering method, not Mobile or Compatibility.
[b]Procedural generation:[/b] [VoxelGI] can be baked in an exported project, which makes it suitable for procedurally generated or user-built levels as long as all the geometry is generated in advance. For games where geometry is generated at any time during gameplay, SDFGI is more suitable (see [member Environment.sdfgi_enabled]).
[b]Performance:[/b] [VoxelGI] is relatively demanding on the GPU and is not suited to low-end hardware such as integrated graphics (consider [LightmapGI] instead). To improve performance, adjust [member ProjectSettings.rendering/global_illumination/voxel_gi/quality] and enable [member ProjectSettings.rendering/global_illumination/gi/use_half_resolution] in the Project Settings. To provide a fallback for low-end hardware, consider adding an option to disable [VoxelGI] in your project's options menus. A [VoxelGI] node can be disabled by hiding it.
[b]Note:[/b] Meshes should have sufficiently thick walls to avoid light leaks (avoid one-sided walls). For interior levels, enclose your level geometry in a sufficiently large box and bridge the loops to close the mesh. To further prevent light leaks, you can also strategically place temporary [MeshInstance3D] nodes with their [member GeometryInstance3D.gi_mode] set to [constant GeometryInstance3D.GI_MODE_STATIC]. These temporary nodes can then be hidden after baking the [VoxelGI] node.
*/
type Instance [1]gdclass.VoxelGI
type Any interface {
	gd.IsClass
	AsVoxelGI() Instance
}

/*
Bakes the effect from all [GeometryInstance3D]s marked with [constant GeometryInstance3D.GI_MODE_STATIC] and [Light3D]s marked with either [constant Light3D.BAKE_STATIC] or [constant Light3D.BAKE_DYNAMIC]. If [param create_visual_debug] is [code]true[/code], after baking the light, this will generate a [MultiMesh] that has a cube representing each solid cell with each cube colored to the cell's albedo color. This can be used to visualize the [VoxelGI]'s data and debug any issues that may be occurring.
[b]Note:[/b] [method bake] works from the editor and in exported projects. This makes it suitable for procedurally generated or user-built levels. Baking a [VoxelGI] node generally takes from 5 to 20 seconds in most scenes. Reducing [member subdiv] can speed up baking.
[b]Note:[/b] [GeometryInstance3D]s and [Light3D]s must be fully ready before [method bake] is called. If you are procedurally creating those and some meshes or lights are missing from your baked [VoxelGI], use [code]call_deferred("bake")[/code] instead of calling [method bake] directly.
*/
func (self Instance) Bake() {
	class(self).Bake([1][1]gdclass.Node{}[0], false)
}

/*
Calls [method bake] with [code]create_visual_debug[/code] enabled.
*/
func (self Instance) DebugBake() {
	class(self).DebugBake()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VoxelGI

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VoxelGI"))
	return Instance{*(*gdclass.VoxelGI)(unsafe.Pointer(&object))}
}

func (self Instance) Subdiv() gdclass.VoxelGISubdiv {
	return gdclass.VoxelGISubdiv(class(self).GetSubdiv())
}

func (self Instance) SetSubdiv(value gdclass.VoxelGISubdiv) {
	class(self).SetSubdiv(value)
}

func (self Instance) Size() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetSize())
}

func (self Instance) SetSize(value Vector3.XYZ) {
	class(self).SetSize(gd.Vector3(value))
}

func (self Instance) CameraAttributes() [1]gdclass.CameraAttributes {
	return [1]gdclass.CameraAttributes(class(self).GetCameraAttributes())
}

func (self Instance) SetCameraAttributes(value [1]gdclass.CameraAttributes) {
	class(self).SetCameraAttributes(value)
}

func (self Instance) Data() [1]gdclass.VoxelGIData {
	return [1]gdclass.VoxelGIData(class(self).GetProbeData())
}

func (self Instance) SetData(value [1]gdclass.VoxelGIData) {
	class(self).SetProbeData(value)
}

//go:nosplit
func (self class) SetProbeData(data [1]gdclass.VoxelGIData) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(data[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGI.Bind_set_probe_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetProbeData() [1]gdclass.VoxelGIData {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGI.Bind_get_probe_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.VoxelGIData{gd.PointerWithOwnershipTransferredToGo[gdclass.VoxelGIData](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSubdiv(subdiv gdclass.VoxelGISubdiv) {
	var frame = callframe.New()
	callframe.Arg(frame, subdiv)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGI.Bind_set_subdiv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubdiv() gdclass.VoxelGISubdiv {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VoxelGISubdiv](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGI.Bind_get_subdiv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSize(size gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGI.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSize() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGI.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCameraAttributes(camera_attributes [1]gdclass.CameraAttributes) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(camera_attributes[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGI.Bind_set_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCameraAttributes() [1]gdclass.CameraAttributes {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGI.Bind_get_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.CameraAttributes{gd.PointerWithOwnershipTransferredToGo[gdclass.CameraAttributes](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Bakes the effect from all [GeometryInstance3D]s marked with [constant GeometryInstance3D.GI_MODE_STATIC] and [Light3D]s marked with either [constant Light3D.BAKE_STATIC] or [constant Light3D.BAKE_DYNAMIC]. If [param create_visual_debug] is [code]true[/code], after baking the light, this will generate a [MultiMesh] that has a cube representing each solid cell with each cube colored to the cell's albedo color. This can be used to visualize the [VoxelGI]'s data and debug any issues that may be occurring.
[b]Note:[/b] [method bake] works from the editor and in exported projects. This makes it suitable for procedurally generated or user-built levels. Baking a [VoxelGI] node generally takes from 5 to 20 seconds in most scenes. Reducing [member subdiv] can speed up baking.
[b]Note:[/b] [GeometryInstance3D]s and [Light3D]s must be fully ready before [method bake] is called. If you are procedurally creating those and some meshes or lights are missing from your baked [VoxelGI], use [code]call_deferred("bake")[/code] instead of calling [method bake] directly.
*/
//go:nosplit
func (self class) Bake(from_node [1]gdclass.Node, create_visual_debug bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(from_node[0])[0])
	callframe.Arg(frame, create_visual_debug)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGI.Bind_bake, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Calls [method bake] with [code]create_visual_debug[/code] enabled.
*/
//go:nosplit
func (self class) DebugBake() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGI.Bind_debug_bake, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsVoxelGI() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVoxelGI() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(VisualInstance3D.Advanced(self.AsVisualInstance3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualInstance3D.Instance(self.AsVisualInstance3D()), name)
	}
}
func init() {
	gdclass.Register("VoxelGI", func(ptr gd.Object) any { return [1]gdclass.VoxelGI{*(*gdclass.VoxelGI)(unsafe.Pointer(&ptr))} })
}

type Subdiv = gdclass.VoxelGISubdiv

const (
	/*Use 64 subdivisions. This is the lowest quality setting, but the fastest. Use it if you can, but especially use it on lower-end hardware.*/
	Subdiv64 Subdiv = 0
	/*Use 128 subdivisions. This is the default quality setting.*/
	Subdiv128 Subdiv = 1
	/*Use 256 subdivisions.*/
	Subdiv256 Subdiv = 2
	/*Use 512 subdivisions. This is the highest quality setting, but the slowest. On lower-end hardware, this could cause the GPU to stall.*/
	Subdiv512 Subdiv = 3
	/*Represents the size of the [enum Subdiv] enum.*/
	SubdivMax Subdiv = 4
)
