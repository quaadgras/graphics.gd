package VisualShader

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Shader"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class provides a graph-like visual editor for creating a [Shader]. Although [VisualShader]s do not require coding, they share the same logic with script shaders. They use [VisualShaderNode]s that can be connected to each other to control the flow of the shader. The visual shader graph is converted to a script shader behind the scenes.

*/
type Simple [1]classdb.VisualShader
func (self Simple) SetMode(mode classdb.ShaderMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMode(mode)
}
func (self Simple) AddNode(atype classdb.VisualShaderType, node [1]classdb.VisualShaderNode, position gd.Vector2, id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddNode(atype, node, position, gd.Int(id))
}
func (self Simple) GetNode(atype classdb.VisualShaderType, id int) [1]classdb.VisualShaderNode {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.VisualShaderNode(Expert(self).GetNode(gc, atype, gd.Int(id)))
}
func (self Simple) SetNodePosition(atype classdb.VisualShaderType, id int, position gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNodePosition(atype, gd.Int(id), position)
}
func (self Simple) GetNodePosition(atype classdb.VisualShaderType, id int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetNodePosition(atype, gd.Int(id)))
}
func (self Simple) GetNodeList(atype classdb.VisualShaderType) gd.PackedInt32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt32Array(Expert(self).GetNodeList(gc, atype))
}
func (self Simple) GetValidNodeId(atype classdb.VisualShaderType) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetValidNodeId(atype)))
}
func (self Simple) RemoveNode(atype classdb.VisualShaderType, id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveNode(atype, gd.Int(id))
}
func (self Simple) ReplaceNode(atype classdb.VisualShaderType, id int, new_class string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ReplaceNode(atype, gd.Int(id), gc.StringName(new_class))
}
func (self Simple) IsNodeConnection(atype classdb.VisualShaderType, from_node int, from_port int, to_node int, to_port int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsNodeConnection(atype, gd.Int(from_node), gd.Int(from_port), gd.Int(to_node), gd.Int(to_port)))
}
func (self Simple) CanConnectNodes(atype classdb.VisualShaderType, from_node int, from_port int, to_node int, to_port int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).CanConnectNodes(atype, gd.Int(from_node), gd.Int(from_port), gd.Int(to_node), gd.Int(to_port)))
}
func (self Simple) ConnectNodes(atype classdb.VisualShaderType, from_node int, from_port int, to_node int, to_port int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).ConnectNodes(atype, gd.Int(from_node), gd.Int(from_port), gd.Int(to_node), gd.Int(to_port)))
}
func (self Simple) DisconnectNodes(atype classdb.VisualShaderType, from_node int, from_port int, to_node int, to_port int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DisconnectNodes(atype, gd.Int(from_node), gd.Int(from_port), gd.Int(to_node), gd.Int(to_port))
}
func (self Simple) ConnectNodesForced(atype classdb.VisualShaderType, from_node int, from_port int, to_node int, to_port int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ConnectNodesForced(atype, gd.Int(from_node), gd.Int(from_port), gd.Int(to_node), gd.Int(to_port))
}
func (self Simple) GetNodeConnections(atype classdb.VisualShaderType) gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Dictionary](Expert(self).GetNodeConnections(gc, atype))
}
func (self Simple) SetGraphOffset(offset gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGraphOffset(offset)
}
func (self Simple) GetGraphOffset() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetGraphOffset())
}
func (self Simple) AttachNodeToFrame(atype classdb.VisualShaderType, id int, frame_ int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AttachNodeToFrame(atype, gd.Int(id), gd.Int(frame_))
}
func (self Simple) DetachNodeFromFrame(atype classdb.VisualShaderType, id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DetachNodeFromFrame(atype, gd.Int(id))
}
func (self Simple) AddVarying(name string, mode classdb.VisualShaderVaryingMode, atype classdb.VisualShaderVaryingType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddVarying(gc.String(name), mode, atype)
}
func (self Simple) RemoveVarying(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveVarying(gc.String(name))
}
func (self Simple) HasVarying(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasVarying(gc.String(name)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShader
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Sets the mode of this shader.
*/
//go:nosplit
func (self class) SetMode(mode classdb.ShaderMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_set_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds the specified [param node] to the shader.
*/
//go:nosplit
func (self class) AddNode(atype classdb.VisualShaderType, node object.VisualShaderNode, position gd.Vector2, id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, mmm.Get(node[0].AsPointer())[0])
	callframe.Arg(frame, position)
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_add_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the shader node instance with specified [param type] and [param id].
*/
//go:nosplit
func (self class) GetNode(ctx gd.Lifetime, atype classdb.VisualShaderType, id gd.Int) object.VisualShaderNode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_get_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.VisualShaderNode
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the position of the specified node.
*/
//go:nosplit
func (self class) SetNodePosition(atype classdb.VisualShaderType, id gd.Int, position gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_set_node_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the position of the specified node within the shader graph.
*/
//go:nosplit
func (self class) GetNodePosition(atype classdb.VisualShaderType, id gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_get_node_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the list of all nodes in the shader with the specified type.
*/
//go:nosplit
func (self class) GetNodeList(ctx gd.Lifetime, atype classdb.VisualShaderType) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_get_node_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns next valid node ID that can be added to the shader graph.
*/
//go:nosplit
func (self class) GetValidNodeId(atype classdb.VisualShaderType) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_get_valid_node_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the specified node from the shader.
*/
//go:nosplit
func (self class) RemoveNode(atype classdb.VisualShaderType, id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_remove_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Replaces the specified node with a node of new class type.
*/
//go:nosplit
func (self class) ReplaceNode(atype classdb.VisualShaderType, id gd.Int, new_class gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	callframe.Arg(frame, mmm.Get(new_class))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_replace_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the specified node and port connection exist.
*/
//go:nosplit
func (self class) IsNodeConnection(atype classdb.VisualShaderType, from_node gd.Int, from_port gd.Int, to_node gd.Int, to_port gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, from_node)
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, to_node)
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_is_node_connection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the specified nodes and ports can be connected together.
*/
//go:nosplit
func (self class) CanConnectNodes(atype classdb.VisualShaderType, from_node gd.Int, from_port gd.Int, to_node gd.Int, to_port gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, from_node)
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, to_node)
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_can_connect_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Connects the specified nodes and ports.
*/
//go:nosplit
func (self class) ConnectNodes(atype classdb.VisualShaderType, from_node gd.Int, from_port gd.Int, to_node gd.Int, to_port gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, from_node)
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, to_node)
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_connect_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Connects the specified nodes and ports.
*/
//go:nosplit
func (self class) DisconnectNodes(atype classdb.VisualShaderType, from_node gd.Int, from_port gd.Int, to_node gd.Int, to_port gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, from_node)
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, to_node)
	callframe.Arg(frame, to_port)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_disconnect_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Connects the specified nodes and ports, even if they can't be connected. Such connection is invalid and will not function properly.
*/
//go:nosplit
func (self class) ConnectNodesForced(atype classdb.VisualShaderType, from_node gd.Int, from_port gd.Int, to_node gd.Int, to_port gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, from_node)
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, to_node)
	callframe.Arg(frame, to_port)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_connect_nodes_forced, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the list of connected nodes with the specified type.
*/
//go:nosplit
func (self class) GetNodeConnections(ctx gd.Lifetime, atype classdb.VisualShaderType) gd.ArrayOf[gd.Dictionary] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_get_node_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Dictionary](ret)
}
//go:nosplit
func (self class) SetGraphOffset(offset gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_set_graph_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGraphOffset() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_get_graph_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Attaches the given node to the given frame.
*/
//go:nosplit
func (self class) AttachNodeToFrame(atype classdb.VisualShaderType, id gd.Int, frame_ gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	callframe.Arg(frame, frame_)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_attach_node_to_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Detaches the given node from the frame it is attached to.
*/
//go:nosplit
func (self class) DetachNodeFromFrame(atype classdb.VisualShaderType, id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_detach_node_from_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a new varying value node to the shader.
*/
//go:nosplit
func (self class) AddVarying(name gd.String, mode classdb.VisualShaderVaryingMode, atype classdb.VisualShaderVaryingType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mode)
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_add_varying, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a varying value node with the given [param name]. Prints an error if a node with this name is not found.
*/
//go:nosplit
func (self class) RemoveVarying(name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_remove_varying, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the shader has a varying with the given [param name].
*/
//go:nosplit
func (self class) HasVarying(name gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShader.Bind_has_varying, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShader() Expert { return self[0].AsVisualShader() }


//go:nosplit
func (self Simple) AsVisualShader() Simple { return self[0].AsVisualShader() }


//go:nosplit
func (self class) AsShader() Shader.Expert { return self[0].AsShader() }


//go:nosplit
func (self Simple) AsShader() Shader.Simple { return self[0].AsShader() }


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

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("VisualShader", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Type = classdb.VisualShaderType

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
type VaryingMode = classdb.VisualShaderVaryingMode

const (
/*Varying is passed from [code]Vertex[/code] function to [code]Fragment[/code] and [code]Light[/code] functions.*/
	VaryingModeVertexToFragLight VaryingMode = 0
/*Varying is passed from [code]Fragment[/code] function to [code]Light[/code] function.*/
	VaryingModeFragToLight VaryingMode = 1
/*Represents the size of the [enum VaryingMode] enum.*/
	VaryingModeMax VaryingMode = 2
)
type VaryingType = classdb.VisualShaderVaryingType

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
