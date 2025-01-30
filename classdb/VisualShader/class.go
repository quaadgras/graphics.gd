// Package VisualShader provides methods for working with VisualShader object instances.
package VisualShader

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
import "graphics.gd/classdb/Shader"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Vector2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

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
func (self Instance) SetMode(mode gdclass.ShaderMode) { //gd:VisualShader.set_mode
	class(self).SetMode(mode)
}

/*
Adds the specified [param node] to the shader.
*/
func (self Instance) AddNode(atype gdclass.VisualShaderType, node [1]gdclass.VisualShaderNode, position Vector2.XY, id int) { //gd:VisualShader.add_node
	class(self).AddNode(atype, node, Vector2.XY(position), int64(id))
}

/*
Returns the shader node instance with specified [param type] and [param id].
*/
func (self Instance) GetNode(atype gdclass.VisualShaderType, id int) [1]gdclass.VisualShaderNode { //gd:VisualShader.get_node
	return [1]gdclass.VisualShaderNode(class(self).GetNode(atype, int64(id)))
}

/*
Sets the position of the specified node.
*/
func (self Instance) SetNodePosition(atype gdclass.VisualShaderType, id int, position Vector2.XY) { //gd:VisualShader.set_node_position
	class(self).SetNodePosition(atype, int64(id), Vector2.XY(position))
}

/*
Returns the position of the specified node within the shader graph.
*/
func (self Instance) GetNodePosition(atype gdclass.VisualShaderType, id int) Vector2.XY { //gd:VisualShader.get_node_position
	return Vector2.XY(class(self).GetNodePosition(atype, int64(id)))
}

/*
Returns the list of all nodes in the shader with the specified type.
*/
func (self Instance) GetNodeList(atype gdclass.VisualShaderType) []int32 { //gd:VisualShader.get_node_list
	return []int32(slices.Collect(class(self).GetNodeList(atype).Values()))
}

/*
Returns next valid node ID that can be added to the shader graph.
*/
func (self Instance) GetValidNodeId(atype gdclass.VisualShaderType) int { //gd:VisualShader.get_valid_node_id
	return int(int(class(self).GetValidNodeId(atype)))
}

/*
Removes the specified node from the shader.
*/
func (self Instance) RemoveNode(atype gdclass.VisualShaderType, id int) { //gd:VisualShader.remove_node
	class(self).RemoveNode(atype, int64(id))
}

/*
Replaces the specified node with a node of new class type.
*/
func (self Instance) ReplaceNode(atype gdclass.VisualShaderType, id int, new_class string) { //gd:VisualShader.replace_node
	class(self).ReplaceNode(atype, int64(id), String.Name(String.New(new_class)))
}

/*
Returns [code]true[/code] if the specified node and port connection exist.
*/
func (self Instance) IsNodeConnection(atype gdclass.VisualShaderType, from_node int, from_port int, to_node int, to_port int) bool { //gd:VisualShader.is_node_connection
	return bool(class(self).IsNodeConnection(atype, int64(from_node), int64(from_port), int64(to_node), int64(to_port)))
}

/*
Returns [code]true[/code] if the specified nodes and ports can be connected together.
*/
func (self Instance) CanConnectNodes(atype gdclass.VisualShaderType, from_node int, from_port int, to_node int, to_port int) bool { //gd:VisualShader.can_connect_nodes
	return bool(class(self).CanConnectNodes(atype, int64(from_node), int64(from_port), int64(to_node), int64(to_port)))
}

/*
Connects the specified nodes and ports.
*/
func (self Instance) ConnectNodes(atype gdclass.VisualShaderType, from_node int, from_port int, to_node int, to_port int) error { //gd:VisualShader.connect_nodes
	return error(gd.ToError(class(self).ConnectNodes(atype, int64(from_node), int64(from_port), int64(to_node), int64(to_port))))
}

/*
Connects the specified nodes and ports.
*/
func (self Instance) DisconnectNodes(atype gdclass.VisualShaderType, from_node int, from_port int, to_node int, to_port int) { //gd:VisualShader.disconnect_nodes
	class(self).DisconnectNodes(atype, int64(from_node), int64(from_port), int64(to_node), int64(to_port))
}

/*
Connects the specified nodes and ports, even if they can't be connected. Such connection is invalid and will not function properly.
*/
func (self Instance) ConnectNodesForced(atype gdclass.VisualShaderType, from_node int, from_port int, to_node int, to_port int) { //gd:VisualShader.connect_nodes_forced
	class(self).ConnectNodesForced(atype, int64(from_node), int64(from_port), int64(to_node), int64(to_port))
}

/*
Returns the list of connected nodes with the specified type.
*/
func (self Instance) GetNodeConnections(atype gdclass.VisualShaderType) []map[string]interface{} { //gd:VisualShader.get_node_connections
	return []map[string]interface{}(gd.ArrayAs[[]map[string]interface{}](gd.InternalArray(class(self).GetNodeConnections(atype))))
}

/*
Attaches the given node to the given frame.
*/
func (self Instance) AttachNodeToFrame(atype gdclass.VisualShaderType, id int, frame_ int) { //gd:VisualShader.attach_node_to_frame
	class(self).AttachNodeToFrame(atype, int64(id), int64(frame_))
}

/*
Detaches the given node from the frame it is attached to.
*/
func (self Instance) DetachNodeFromFrame(atype gdclass.VisualShaderType, id int) { //gd:VisualShader.detach_node_from_frame
	class(self).DetachNodeFromFrame(atype, int64(id))
}

/*
Adds a new varying value node to the shader.
*/
func (self Instance) AddVarying(name string, mode gdclass.VisualShaderVaryingMode, atype gdclass.VisualShaderVaryingType) { //gd:VisualShader.add_varying
	class(self).AddVarying(String.New(name), mode, atype)
}

/*
Removes a varying value node with the given [param name]. Prints an error if a node with this name is not found.
*/
func (self Instance) RemoveVarying(name string) { //gd:VisualShader.remove_varying
	class(self).RemoveVarying(String.New(name))
}

/*
Returns [code]true[/code] if the shader has a varying with the given [param name].
*/
func (self Instance) HasVarying(name string) bool { //gd:VisualShader.has_varying
	return bool(class(self).HasVarying(String.New(name)))
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
	class(self).SetGraphOffset(Vector2.XY(value))
}

/*
Sets the mode of this shader.
*/
//go:nosplit
func (self class) SetMode(mode gdclass.ShaderMode) { //gd:VisualShader.set_mode
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
func (self class) AddNode(atype gdclass.VisualShaderType, node [1]gdclass.VisualShaderNode, position Vector2.XY, id int64) { //gd:VisualShader.add_node
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
func (self class) GetNode(atype gdclass.VisualShaderType, id int64) [1]gdclass.VisualShaderNode { //gd:VisualShader.get_node
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
func (self class) SetNodePosition(atype gdclass.VisualShaderType, id int64, position Vector2.XY) { //gd:VisualShader.set_node_position
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
func (self class) GetNodePosition(atype gdclass.VisualShaderType, id int64) Vector2.XY { //gd:VisualShader.get_node_position
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_get_node_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the list of all nodes in the shader with the specified type.
*/
//go:nosplit
func (self class) GetNodeList(atype gdclass.VisualShaderType) Packed.Array[int32] { //gd:VisualShader.get_node_list
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_get_node_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns next valid node ID that can be added to the shader graph.
*/
//go:nosplit
func (self class) GetValidNodeId(atype gdclass.VisualShaderType) int64 { //gd:VisualShader.get_valid_node_id
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_get_valid_node_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the specified node from the shader.
*/
//go:nosplit
func (self class) RemoveNode(atype gdclass.VisualShaderType, id int64) { //gd:VisualShader.remove_node
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
func (self class) ReplaceNode(atype gdclass.VisualShaderType, id int64, new_class String.Name) { //gd:VisualShader.replace_node
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(new_class)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_replace_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the specified node and port connection exist.
*/
//go:nosplit
func (self class) IsNodeConnection(atype gdclass.VisualShaderType, from_node int64, from_port int64, to_node int64, to_port int64) bool { //gd:VisualShader.is_node_connection
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
func (self class) CanConnectNodes(atype gdclass.VisualShaderType, from_node int64, from_port int64, to_node int64, to_port int64) bool { //gd:VisualShader.can_connect_nodes
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
func (self class) ConnectNodes(atype gdclass.VisualShaderType, from_node int64, from_port int64, to_node int64, to_port int64) Error.Code { //gd:VisualShader.connect_nodes
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, from_node)
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, to_node)
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_connect_nodes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Connects the specified nodes and ports.
*/
//go:nosplit
func (self class) DisconnectNodes(atype gdclass.VisualShaderType, from_node int64, from_port int64, to_node int64, to_port int64) { //gd:VisualShader.disconnect_nodes
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
func (self class) ConnectNodesForced(atype gdclass.VisualShaderType, from_node int64, from_port int64, to_node int64, to_port int64) { //gd:VisualShader.connect_nodes_forced
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
func (self class) GetNodeConnections(atype gdclass.VisualShaderType) Array.Contains[Dictionary.Any] { //gd:VisualShader.get_node_connections
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_get_node_connections, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Dictionary.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGraphOffset(offset Vector2.XY) { //gd:VisualShader.set_graph_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_set_graph_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGraphOffset() Vector2.XY { //gd:VisualShader.get_graph_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_get_graph_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Attaches the given node to the given frame.
*/
//go:nosplit
func (self class) AttachNodeToFrame(atype gdclass.VisualShaderType, id int64, frame_ int64) { //gd:VisualShader.attach_node_to_frame
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
func (self class) DetachNodeFromFrame(atype gdclass.VisualShaderType, id int64) { //gd:VisualShader.detach_node_from_frame
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
func (self class) AddVarying(name String.Readable, mode gdclass.VisualShaderVaryingMode, atype gdclass.VisualShaderVaryingType) { //gd:VisualShader.add_varying
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
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
func (self class) RemoveVarying(name String.Readable) { //gd:VisualShader.remove_varying
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_remove_varying, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the shader has a varying with the given [param name].
*/
//go:nosplit
func (self class) HasVarying(name String.Readable) bool { //gd:VisualShader.has_varying
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
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

type Type = gdclass.VisualShaderType //gd:VisualShader.Type

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

type VaryingMode = gdclass.VisualShaderVaryingMode //gd:VisualShader.VaryingMode

const (
	/*Varying is passed from [code]Vertex[/code] function to [code]Fragment[/code] and [code]Light[/code] functions.*/
	VaryingModeVertexToFragLight VaryingMode = 0
	/*Varying is passed from [code]Fragment[/code] function to [code]Light[/code] function.*/
	VaryingModeFragToLight VaryingMode = 1
	/*Represents the size of the [enum VaryingMode] enum.*/
	VaryingModeMax VaryingMode = 2
)

type VaryingType = gdclass.VisualShaderVaryingType //gd:VisualShader.VaryingType

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
