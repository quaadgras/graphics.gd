// Package VisualShader provides methods for working with VisualShader object instances.
package VisualShader

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Shader"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This class provides a graph-like visual editor for creating a [Shader]. Although [VisualShader]s do not require coding, they share the same logic with script shaders. They use [VisualShaderNode]s that can be connected to each other to control the flow of the shader. The visual shader graph is converted to a script shader behind the scenes.
*/
type Instance [1]gdclass.VisualShader

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShader() Instance
}

/*
Sets the mode of this shader.
*/
func (self Instance) SetMode(mode gdclass.ShaderMode) {
	class(self).SetMode(mode)
}

/*
Adds the specified [param node] to the shader.
*/
func (self Instance) AddNode(atype gdclass.VisualShaderType, node [1]gdclass.VisualShaderNode, position Vector2.XY, id int) {
	class(self).AddNode(atype, node, gd.Vector2(position), gd.Int(id))
}

/*
Returns the shader node instance with specified [param type] and [param id].
*/
func (self Instance) GetNode(atype gdclass.VisualShaderType, id int) [1]gdclass.VisualShaderNode {
	return [1]gdclass.VisualShaderNode(class(self).GetNode(atype, gd.Int(id)))
}

/*
Sets the position of the specified node.
*/
func (self Instance) SetNodePosition(atype gdclass.VisualShaderType, id int, position Vector2.XY) {
	class(self).SetNodePosition(atype, gd.Int(id), gd.Vector2(position))
}

/*
Returns the position of the specified node within the shader graph.
*/
func (self Instance) GetNodePosition(atype gdclass.VisualShaderType, id int) Vector2.XY {
	return Vector2.XY(class(self).GetNodePosition(atype, gd.Int(id)))
}

/*
Returns the list of all nodes in the shader with the specified type.
*/
func (self Instance) GetNodeList(atype gdclass.VisualShaderType) []int32 {
	return []int32(class(self).GetNodeList(atype).AsSlice())
}

/*
Returns next valid node ID that can be added to the shader graph.
*/
func (self Instance) GetValidNodeId(atype gdclass.VisualShaderType) int {
	return int(int(class(self).GetValidNodeId(atype)))
}

/*
Removes the specified node from the shader.
*/
func (self Instance) RemoveNode(atype gdclass.VisualShaderType, id int) {
	class(self).RemoveNode(atype, gd.Int(id))
}

/*
Replaces the specified node with a node of new class type.
*/
func (self Instance) ReplaceNode(atype gdclass.VisualShaderType, id int, new_class string) {
	class(self).ReplaceNode(atype, gd.Int(id), gd.NewStringName(new_class))
}

/*
Returns [code]true[/code] if the specified node and port connection exist.
*/
func (self Instance) IsNodeConnection(atype gdclass.VisualShaderType, from_node int, from_port int, to_node int, to_port int) bool {
	return bool(class(self).IsNodeConnection(atype, gd.Int(from_node), gd.Int(from_port), gd.Int(to_node), gd.Int(to_port)))
}

/*
Returns [code]true[/code] if the specified nodes and ports can be connected together.
*/
func (self Instance) CanConnectNodes(atype gdclass.VisualShaderType, from_node int, from_port int, to_node int, to_port int) bool {
	return bool(class(self).CanConnectNodes(atype, gd.Int(from_node), gd.Int(from_port), gd.Int(to_node), gd.Int(to_port)))
}

/*
Connects the specified nodes and ports.
*/
func (self Instance) ConnectNodes(atype gdclass.VisualShaderType, from_node int, from_port int, to_node int, to_port int) error {
	return error(gd.ToError(class(self).ConnectNodes(atype, gd.Int(from_node), gd.Int(from_port), gd.Int(to_node), gd.Int(to_port))))
}

/*
Connects the specified nodes and ports.
*/
func (self Instance) DisconnectNodes(atype gdclass.VisualShaderType, from_node int, from_port int, to_node int, to_port int) {
	class(self).DisconnectNodes(atype, gd.Int(from_node), gd.Int(from_port), gd.Int(to_node), gd.Int(to_port))
}

/*
Connects the specified nodes and ports, even if they can't be connected. Such connection is invalid and will not function properly.
*/
func (self Instance) ConnectNodesForced(atype gdclass.VisualShaderType, from_node int, from_port int, to_node int, to_port int) {
	class(self).ConnectNodesForced(atype, gd.Int(from_node), gd.Int(from_port), gd.Int(to_node), gd.Int(to_port))
}

/*
Returns the list of connected nodes with the specified type.
*/
func (self Instance) GetNodeConnections(atype gdclass.VisualShaderType) gd.Array {
	return gd.Array(class(self).GetNodeConnections(atype))
}

/*
Attaches the given node to the given frame.
*/
func (self Instance) AttachNodeToFrame(atype gdclass.VisualShaderType, id int, frame_ int) {
	class(self).AttachNodeToFrame(atype, gd.Int(id), gd.Int(frame_))
}

/*
Detaches the given node from the frame it is attached to.
*/
func (self Instance) DetachNodeFromFrame(atype gdclass.VisualShaderType, id int) {
	class(self).DetachNodeFromFrame(atype, gd.Int(id))
}

/*
Adds a new varying value node to the shader.
*/
func (self Instance) AddVarying(name string, mode gdclass.VisualShaderVaryingMode, atype gdclass.VisualShaderVaryingType) {
	class(self).AddVarying(gd.NewString(name), mode, atype)
}

/*
Removes a varying value node with the given [param name]. Prints an error if a node with this name is not found.
*/
func (self Instance) RemoveVarying(name string) {
	class(self).RemoveVarying(gd.NewString(name))
}

/*
Returns [code]true[/code] if the shader has a varying with the given [param name].
*/
func (self Instance) HasVarying(name string) bool {
	return bool(class(self).HasVarying(gd.NewString(name)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShader

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShader"))
	casted := Instance{*(*gdclass.VisualShader)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) GraphOffset() Vector2.XY {
	return Vector2.XY(class(self).GetGraphOffset())
}

func (self Instance) SetGraphOffset(value Vector2.XY) {
	class(self).SetGraphOffset(gd.Vector2(value))
}

/*
Sets the mode of this shader.
*/
//go:nosplit
func (self class) SetMode(mode gdclass.ShaderMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_set_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds the specified [param node] to the shader.
*/
//go:nosplit
func (self class) AddNode(atype gdclass.VisualShaderType, node [1]gdclass.VisualShaderNode, position gd.Vector2, id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, pointers.Get(node[0])[0])
	callframe.Arg(frame, position)
	callframe.Arg(frame, id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_add_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the shader node instance with specified [param type] and [param id].
*/
//go:nosplit
func (self class) GetNode(atype gdclass.VisualShaderType, id gd.Int) [1]gdclass.VisualShaderNode {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_get_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.VisualShaderNode{gd.PointerWithOwnershipTransferredToGo[gdclass.VisualShaderNode](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the position of the specified node.
*/
//go:nosplit
func (self class) SetNodePosition(atype gdclass.VisualShaderType, id gd.Int, position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_set_node_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the position of the specified node within the shader graph.
*/
//go:nosplit
func (self class) GetNodePosition(atype gdclass.VisualShaderType, id gd.Int) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_get_node_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the list of all nodes in the shader with the specified type.
*/
//go:nosplit
func (self class) GetNodeList(atype gdclass.VisualShaderType) gd.PackedInt32Array {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_get_node_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns next valid node ID that can be added to the shader graph.
*/
//go:nosplit
func (self class) GetValidNodeId(atype gdclass.VisualShaderType) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_get_valid_node_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the specified node from the shader.
*/
//go:nosplit
func (self class) RemoveNode(atype gdclass.VisualShaderType, id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_remove_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Replaces the specified node with a node of new class type.
*/
//go:nosplit
func (self class) ReplaceNode(atype gdclass.VisualShaderType, id gd.Int, new_class gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(new_class))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_replace_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the specified node and port connection exist.
*/
//go:nosplit
func (self class) IsNodeConnection(atype gdclass.VisualShaderType, from_node gd.Int, from_port gd.Int, to_node gd.Int, to_port gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, from_node)
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, to_node)
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_is_node_connection, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the specified nodes and ports can be connected together.
*/
//go:nosplit
func (self class) CanConnectNodes(atype gdclass.VisualShaderType, from_node gd.Int, from_port gd.Int, to_node gd.Int, to_port gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, from_node)
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, to_node)
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_can_connect_nodes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Connects the specified nodes and ports.
*/
//go:nosplit
func (self class) ConnectNodes(atype gdclass.VisualShaderType, from_node gd.Int, from_port gd.Int, to_node gd.Int, to_port gd.Int) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, from_node)
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, to_node)
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_connect_nodes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Connects the specified nodes and ports.
*/
//go:nosplit
func (self class) DisconnectNodes(atype gdclass.VisualShaderType, from_node gd.Int, from_port gd.Int, to_node gd.Int, to_port gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, from_node)
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, to_node)
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_disconnect_nodes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Connects the specified nodes and ports, even if they can't be connected. Such connection is invalid and will not function properly.
*/
//go:nosplit
func (self class) ConnectNodesForced(atype gdclass.VisualShaderType, from_node gd.Int, from_port gd.Int, to_node gd.Int, to_port gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, from_node)
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, to_node)
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_connect_nodes_forced, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the list of connected nodes with the specified type.
*/
//go:nosplit
func (self class) GetNodeConnections(atype gdclass.VisualShaderType) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_get_node_connections, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGraphOffset(offset gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_set_graph_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGraphOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_get_graph_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Attaches the given node to the given frame.
*/
//go:nosplit
func (self class) AttachNodeToFrame(atype gdclass.VisualShaderType, id gd.Int, frame_ gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	callframe.Arg(frame, frame_)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_attach_node_to_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Detaches the given node from the frame it is attached to.
*/
//go:nosplit
func (self class) DetachNodeFromFrame(atype gdclass.VisualShaderType, id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_detach_node_from_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a new varying value node to the shader.
*/
//go:nosplit
func (self class) AddVarying(name gd.String, mode gdclass.VisualShaderVaryingMode, atype gdclass.VisualShaderVaryingType) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, mode)
	callframe.Arg(frame, atype)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_add_varying, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a varying value node with the given [param name]. Prints an error if a node with this name is not found.
*/
//go:nosplit
func (self class) RemoveVarying(name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_remove_varying, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the shader has a varying with the given [param name].
*/
//go:nosplit
func (self class) HasVarying(name gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_has_varying, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShader() Advanced     { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualShader() Instance  { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsShader() Shader.Advanced    { return *((*Shader.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsShader() Shader.Instance { return *((*Shader.Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Shader.Advanced(self.AsShader()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Shader.Instance(self.AsShader()), name)
	}
}
func init() {
	gdclass.Register("VisualShader", func(ptr gd.Object) any {
		return [1]gdclass.VisualShader{*(*gdclass.VisualShader)(unsafe.Pointer(&ptr))}
	})
}

type Type = gdclass.VisualShaderType

const (
	/*A vertex shader, operating on vertices.*/
	TypeVertex Type = 0
	/*A fragment shader, operating on fragments (pixels).*/
	TypeFragment Type = 1
	/*A shader for light calculations.*/
	TypeLight Type = 2
	/*A function for the "start" stage of particle shader.*/
	TypeStart Type = 3
	/*A function for the "process" stage of particle shader.*/
	TypeProcess Type = 4
	/*A function for the "collide" stage (particle collision handler) of particle shader.*/
	TypeCollide Type = 5
	/*A function for the "start" stage of particle shader, with customized output.*/
	TypeStartCustom Type = 6
	/*A function for the "process" stage of particle shader, with customized output.*/
	TypeProcessCustom Type = 7
	/*A shader for 3D environment's sky.*/
	TypeSky Type = 8
	/*A compute shader that runs for each froxel of the volumetric fog map.*/
	TypeFog Type = 9
	/*Represents the size of the [enum Type] enum.*/
	TypeMax Type = 10
)

type VaryingMode = gdclass.VisualShaderVaryingMode

const (
	/*Varying is passed from [code]Vertex[/code] function to [code]Fragment[/code] and [code]Light[/code] functions.*/
	VaryingModeVertexToFragLight VaryingMode = 0
	/*Varying is passed from [code]Fragment[/code] function to [code]Light[/code] function.*/
	VaryingModeFragToLight VaryingMode = 1
	/*Represents the size of the [enum VaryingMode] enum.*/
	VaryingModeMax VaryingMode = 2
)

type VaryingType = gdclass.VisualShaderVaryingType

const (
	/*Varying is of type [float].*/
	VaryingTypeFloat VaryingType = 0
	/*Varying is of type [int].*/
	VaryingTypeInt VaryingType = 1
	/*Varying is of type unsigned [int].*/
	VaryingTypeUint VaryingType = 2
	/*Varying is of type [Vector2].*/
	VaryingTypeVector2d VaryingType = 3
	/*Varying is of type [Vector3].*/
	VaryingTypeVector3d VaryingType = 4
	/*Varying is of type [Vector4].*/
	VaryingTypeVector4d VaryingType = 5
	/*Varying is of type [bool].*/
	VaryingTypeBoolean VaryingType = 6
	/*Varying is of type [Transform3D].*/
	VaryingTypeTransform VaryingType = 7
	/*Represents the size of the [enum VaryingType] enum.*/
	VaryingTypeMax VaryingType = 8
)

type Error = gd.Error

const (
	/*Methods that return [enum Error] return [constant OK] when no error occurred.
	  Since [constant OK] has value 0, and all other error constants are positive integers, it can also be used in boolean checks.
	  [b]Example:[/b]
	  [codeblock]
	  var error = method_that_returns_error()
	  if error != OK:
	      printerr("Failure!")

	  # Or, alternatively:
	  if error:
	      printerr("Still failing!")
	  [/codeblock]
	  [b]Note:[/b] Many functions do not return an error code, but will print error messages to standard output.*/
	Ok Error = 0
	/*Generic error.*/
	Failed Error = 1
	/*Unavailable error.*/
	ErrUnavailable Error = 2
	/*Unconfigured error.*/
	ErrUnconfigured Error = 3
	/*Unauthorized error.*/
	ErrUnauthorized Error = 4
	/*Parameter range error.*/
	ErrParameterRangeError Error = 5
	/*Out of memory (OOM) error.*/
	ErrOutOfMemory Error = 6
	/*File: Not found error.*/
	ErrFileNotFound Error = 7
	/*File: Bad drive error.*/
	ErrFileBadDrive Error = 8
	/*File: Bad path error.*/
	ErrFileBadPath Error = 9
	/*File: No permission error.*/
	ErrFileNoPermission Error = 10
	/*File: Already in use error.*/
	ErrFileAlreadyInUse Error = 11
	/*File: Can't open error.*/
	ErrFileCantOpen Error = 12
	/*File: Can't write error.*/
	ErrFileCantWrite Error = 13
	/*File: Can't read error.*/
	ErrFileCantRead Error = 14
	/*File: Unrecognized error.*/
	ErrFileUnrecognized Error = 15
	/*File: Corrupt error.*/
	ErrFileCorrupt Error = 16
	/*File: Missing dependencies error.*/
	ErrFileMissingDependencies Error = 17
	/*File: End of file (EOF) error.*/
	ErrFileEof Error = 18
	/*Can't open error.*/
	ErrCantOpen Error = 19
	/*Can't create error.*/
	ErrCantCreate Error = 20
	/*Query failed error.*/
	ErrQueryFailed Error = 21
	/*Already in use error.*/
	ErrAlreadyInUse Error = 22
	/*Locked error.*/
	ErrLocked Error = 23
	/*Timeout error.*/
	ErrTimeout Error = 24
	/*Can't connect error.*/
	ErrCantConnect Error = 25
	/*Can't resolve error.*/
	ErrCantResolve Error = 26
	/*Connection error.*/
	ErrConnectionError Error = 27
	/*Can't acquire resource error.*/
	ErrCantAcquireResource Error = 28
	/*Can't fork process error.*/
	ErrCantFork Error = 29
	/*Invalid data error.*/
	ErrInvalidData Error = 30
	/*Invalid parameter error.*/
	ErrInvalidParameter Error = 31
	/*Already exists error.*/
	ErrAlreadyExists Error = 32
	/*Does not exist error.*/
	ErrDoesNotExist Error = 33
	/*Database: Read error.*/
	ErrDatabaseCantRead Error = 34
	/*Database: Write error.*/
	ErrDatabaseCantWrite Error = 35
	/*Compilation failed error.*/
	ErrCompilationFailed Error = 36
	/*Method not found error.*/
	ErrMethodNotFound Error = 37
	/*Linking failed error.*/
	ErrLinkFailed Error = 38
	/*Script failed error.*/
	ErrScriptFailed Error = 39
	/*Cycling link (import cycle) error.*/
	ErrCyclicLink Error = 40
	/*Invalid declaration error.*/
	ErrInvalidDeclaration Error = 41
	/*Duplicate symbol error.*/
	ErrDuplicateSymbol Error = 42
	/*Parse error.*/
	ErrParseError Error = 43
	/*Busy error.*/
	ErrBusy Error = 44
	/*Skip error.*/
	ErrSkip Error = 45
	/*Help error. Used internally when passing [code]--version[/code] or [code]--help[/code] as executable options.*/
	ErrHelp Error = 46
	/*Bug error, caused by an implementation issue in the method.
	  [b]Note:[/b] If a built-in method returns this code, please open an issue on [url=https://github.com/godotengine/godot/issues]the GitHub Issue Tracker[/url].*/
	ErrBug Error = 47
	/*Printer on fire error (This is an easter egg, no built-in methods return this error code).*/
	ErrPrinterOnFire Error = 48
)
