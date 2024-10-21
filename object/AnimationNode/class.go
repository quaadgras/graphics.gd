package AnimationNode

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
Base resource for [AnimationTree] nodes. In general, it's not used directly, but you can create custom ones with custom blending formulas.
Inherit this when creating animation nodes mainly for use in [AnimationNodeBlendTree], otherwise [AnimationRootNode] should be used instead.
You can access the time information as read-only parameter which is processed and stored in the previous frame for all nodes except [AnimationNodeOutput].
[b]Note:[/b] If multiple inputs exist in the [AnimationNode], which time information takes precedence depends on the type of [AnimationNode].
[codeblock]
var current_length = $AnimationTree[parameters/AnimationNodeName/current_length]
var current_position = $AnimationTree[parameters/AnimationNodeName/current_position]
var current_delta = $AnimationTree[parameters/AnimationNodeName/current_delta]
[/codeblock]
	// AnimationNode methods that can be overridden by a [Class] that extends it.
	type AnimationNode interface {
		//When inheriting from [AnimationRootNode], implement this virtual method to return all child animation nodes in order as a [code]name: node[/code] dictionary.
		GetChildNodes() gd.Dictionary
		//When inheriting from [AnimationRootNode], implement this virtual method to return a list of the properties on this animation node. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees. Format is similar to [method Object.get_property_list].
		GetParameterList() gd.Array
		//When inheriting from [AnimationRootNode], implement this virtual method to return a child animation node by its [param name].
		GetChildByName(name gd.StringName) object.AnimationNode
		//When inheriting from [AnimationRootNode], implement this virtual method to return the default value of a [param parameter]. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees.
		GetParameterDefaultValue(parameter gd.StringName) gd.Variant
		//When inheriting from [AnimationRootNode], implement this virtual method to return whether the [param parameter] is read-only. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees.
		IsParameterReadOnly(parameter gd.StringName) bool
		//When inheriting from [AnimationRootNode], implement this virtual method to run some code when this animation node is processed. The [param time] parameter is a relative delta, unless [param seek] is [code]true[/code], in which case it is absolute.
		//Here, call the [method blend_input], [method blend_node] or [method blend_animation] functions. You can also use [method get_parameter] and [method set_parameter] to modify local memory.
		//This function should return the delta.
		Process(time gd.Float, seek bool, is_external_seeking bool, test_only bool) gd.Float
		//When inheriting from [AnimationRootNode], implement this virtual method to override the text caption for this animation node.
		GetCaption() gd.String
		//When inheriting from [AnimationRootNode], implement this virtual method to return whether the blend tree editor should display filter editing on this animation node.
		HasFilter() bool
	}

*/
type Simple [1]classdb.AnimationNode
func (Simple) _get_child_nodes(impl func(ptr unsafe.Pointer) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _get_parameter_list(impl func(ptr unsafe.Pointer) gd.Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _get_child_by_name(impl func(ptr unsafe.Pointer, name string) [1]classdb.AnimationNode, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var name = mmm.Let[gd.StringName](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, name.String())
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _get_parameter_default_value(impl func(ptr unsafe.Pointer, parameter string) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var parameter = mmm.Let[gd.StringName](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, parameter.String())
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _is_parameter_read_only(impl func(ptr unsafe.Pointer, parameter string) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var parameter = mmm.Let[gd.StringName](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, parameter.String())
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _process(impl func(ptr unsafe.Pointer, time float64, seek bool, is_external_seeking bool, test_only bool) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var time = gd.UnsafeGet[gd.Float](p_args,0)
		var seek = gd.UnsafeGet[bool](p_args,1)
		var is_external_seeking = gd.UnsafeGet[bool](p_args,2)
		var test_only = gd.UnsafeGet[bool](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, float64(time), seek, is_external_seeking, test_only)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _get_caption(impl func(ptr unsafe.Pointer) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Simple) _has_filter(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (self Simple) AddInput(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).AddInput(gc.String(name)))
}
func (self Simple) RemoveInput(index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveInput(gd.Int(index))
}
func (self Simple) SetInputName(input int, name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).SetInputName(gd.Int(input), gc.String(name)))
}
func (self Simple) GetInputName(input int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetInputName(gc, gd.Int(input)).String())
}
func (self Simple) GetInputCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetInputCount()))
}
func (self Simple) FindInput(name string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).FindInput(gc.String(name))))
}
func (self Simple) SetFilterPath(path gd.NodePath, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFilterPath(path, enable)
}
func (self Simple) IsPathFiltered(path gd.NodePath) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsPathFiltered(path))
}
func (self Simple) SetFilterEnabled(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFilterEnabled(enable)
}
func (self Simple) IsFilterEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsFilterEnabled())
}
func (self Simple) BlendAnimation(animation string, time float64, delta float64, seeked bool, is_external_seeking bool, blend float64, looped_flag classdb.AnimationLoopedFlag) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BlendAnimation(gc.StringName(animation), gd.Float(time), gd.Float(delta), seeked, is_external_seeking, gd.Float(blend), looped_flag)
}
func (self Simple) BlendNode(name string, node [1]classdb.AnimationNode, time float64, seek bool, is_external_seeking bool, blend float64, filter classdb.AnimationNodeFilterAction, sync bool, test_only bool) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).BlendNode(gc.StringName(name), node, gd.Float(time), seek, is_external_seeking, gd.Float(blend), filter, sync, test_only)))
}
func (self Simple) BlendInput(input_index int, time float64, seek bool, is_external_seeking bool, blend float64, filter classdb.AnimationNodeFilterAction, sync bool, test_only bool) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).BlendInput(gd.Int(input_index), gd.Float(time), seek, is_external_seeking, gd.Float(blend), filter, sync, test_only)))
}
func (self Simple) SetParameter(name string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParameter(gc.StringName(name), value)
}
func (self Simple) GetParameter(name string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetParameter(gc, gc.StringName(name)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AnimationNode
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
When inheriting from [AnimationRootNode], implement this virtual method to return all child animation nodes in order as a [code]name: node[/code] dictionary.
*/
func (class) _get_child_nodes(impl func(ptr unsafe.Pointer) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return a list of the properties on this animation node. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees. Format is similar to [method Object.get_property_list].
*/
func (class) _get_parameter_list(impl func(ptr unsafe.Pointer) gd.Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return a child animation node by its [param name].
*/
func (class) _get_child_by_name(impl func(ptr unsafe.Pointer, name gd.StringName) object.AnimationNode, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var name = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, name)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		ctx.End()
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return the default value of a [param parameter]. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees.
*/
func (class) _get_parameter_default_value(impl func(ptr unsafe.Pointer, parameter gd.StringName) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var parameter = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, parameter)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return whether the [param parameter] is read-only. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees.
*/
func (class) _is_parameter_read_only(impl func(ptr unsafe.Pointer, parameter gd.StringName) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var parameter = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, parameter)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to run some code when this animation node is processed. The [param time] parameter is a relative delta, unless [param seek] is [code]true[/code], in which case it is absolute.
Here, call the [method blend_input], [method blend_node] or [method blend_animation] functions. You can also use [method get_parameter] and [method set_parameter] to modify local memory.
This function should return the delta.
*/
func (class) _process(impl func(ptr unsafe.Pointer, time gd.Float, seek bool, is_external_seeking bool, test_only bool) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var time = gd.UnsafeGet[gd.Float](p_args,0)
		var seek = gd.UnsafeGet[bool](p_args,1)
		var is_external_seeking = gd.UnsafeGet[bool](p_args,2)
		var test_only = gd.UnsafeGet[bool](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, time, seek, is_external_seeking, test_only)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to override the text caption for this animation node.
*/
func (class) _get_caption(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return whether the blend tree editor should display filter editing on this animation node.
*/
func (class) _has_filter(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Adds an input to the animation node. This is only useful for animation nodes created for use in an [AnimationNodeBlendTree]. If the addition fails, returns [code]false[/code].
*/
//go:nosplit
func (self class) AddInput(name gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNode.Bind_add_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes an input, call this only when inactive.
*/
//go:nosplit
func (self class) RemoveInput(index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNode.Bind_remove_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the name of the input at the given [param input] index. If the setting fails, returns [code]false[/code].
*/
//go:nosplit
func (self class) SetInputName(input gd.Int, name gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, input)
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNode.Bind_set_input_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets the name of an input by index.
*/
//go:nosplit
func (self class) GetInputName(ctx gd.Lifetime, input gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, input)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNode.Bind_get_input_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Amount of inputs in this animation node, only useful for animation nodes that go into [AnimationNodeBlendTree].
*/
//go:nosplit
func (self class) GetInputCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNode.Bind_get_input_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the input index which corresponds to [param name]. If not found, returns [code]-1[/code].
*/
//go:nosplit
func (self class) FindInput(name gd.String) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNode.Bind_find_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds or removes a path for the filter.
*/
//go:nosplit
func (self class) SetFilterPath(path gd.NodePath, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNode.Bind_set_filter_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the given path is filtered.
*/
//go:nosplit
func (self class) IsPathFiltered(path gd.NodePath) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNode.Bind_is_path_filtered, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFilterEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNode.Bind_set_filter_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFilterEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNode.Bind_is_filter_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Blend an animation by [param blend] amount (name must be valid in the linked [AnimationPlayer]). A [param time] and [param delta] may be passed, as well as whether [param seeked] happened.
A [param looped_flag] is used by internal processing immediately after the loop. See also [enum Animation.LoopedFlag].
*/
//go:nosplit
func (self class) BlendAnimation(animation gd.StringName, time gd.Float, delta gd.Float, seeked bool, is_external_seeking bool, blend gd.Float, looped_flag classdb.AnimationLoopedFlag)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(animation))
	callframe.Arg(frame, time)
	callframe.Arg(frame, delta)
	callframe.Arg(frame, seeked)
	callframe.Arg(frame, is_external_seeking)
	callframe.Arg(frame, blend)
	callframe.Arg(frame, looped_flag)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNode.Bind_blend_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Blend another animation node (in case this animation node contains child animation nodes). This function is only useful if you inherit from [AnimationRootNode] instead, otherwise editors will not display your animation node for addition.
*/
//go:nosplit
func (self class) BlendNode(name gd.StringName, node object.AnimationNode, time gd.Float, seek bool, is_external_seeking bool, blend gd.Float, filter classdb.AnimationNodeFilterAction, sync bool, test_only bool) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(node[0].AsPointer())[0])
	callframe.Arg(frame, time)
	callframe.Arg(frame, seek)
	callframe.Arg(frame, is_external_seeking)
	callframe.Arg(frame, blend)
	callframe.Arg(frame, filter)
	callframe.Arg(frame, sync)
	callframe.Arg(frame, test_only)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNode.Bind_blend_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Blend an input. This is only useful for animation nodes created for an [AnimationNodeBlendTree]. The [param time] parameter is a relative delta, unless [param seek] is [code]true[/code], in which case it is absolute. A filter mode may be optionally passed (see [enum FilterAction] for options).
*/
//go:nosplit
func (self class) BlendInput(input_index gd.Int, time gd.Float, seek bool, is_external_seeking bool, blend gd.Float, filter classdb.AnimationNodeFilterAction, sync bool, test_only bool) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, input_index)
	callframe.Arg(frame, time)
	callframe.Arg(frame, seek)
	callframe.Arg(frame, is_external_seeking)
	callframe.Arg(frame, blend)
	callframe.Arg(frame, filter)
	callframe.Arg(frame, sync)
	callframe.Arg(frame, test_only)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNode.Bind_blend_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a custom parameter. These are used as local memory, because resources can be reused across the tree or scenes.
*/
//go:nosplit
func (self class) SetParameter(name gd.StringName, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNode.Bind_set_parameter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets the value of a parameter. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees.
*/
//go:nosplit
func (self class) GetParameter(ctx gd.Lifetime, name gd.StringName) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNode.Bind_get_parameter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsAnimationNode() Expert { return self[0].AsAnimationNode() }


//go:nosplit
func (self Simple) AsAnimationNode() Simple { return self[0].AsAnimationNode() }


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
	case "_get_child_nodes": return reflect.ValueOf(self._get_child_nodes);
	case "_get_parameter_list": return reflect.ValueOf(self._get_parameter_list);
	case "_get_child_by_name": return reflect.ValueOf(self._get_child_by_name);
	case "_get_parameter_default_value": return reflect.ValueOf(self._get_parameter_default_value);
	case "_is_parameter_read_only": return reflect.ValueOf(self._is_parameter_read_only);
	case "_process": return reflect.ValueOf(self._process);
	case "_get_caption": return reflect.ValueOf(self._get_caption);
	case "_has_filter": return reflect.ValueOf(self._has_filter);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_get_child_nodes": return reflect.ValueOf(self._get_child_nodes);
	case "_get_parameter_list": return reflect.ValueOf(self._get_parameter_list);
	case "_get_child_by_name": return reflect.ValueOf(self._get_child_by_name);
	case "_get_parameter_default_value": return reflect.ValueOf(self._get_parameter_default_value);
	case "_is_parameter_read_only": return reflect.ValueOf(self._is_parameter_read_only);
	case "_process": return reflect.ValueOf(self._process);
	case "_get_caption": return reflect.ValueOf(self._get_caption);
	case "_has_filter": return reflect.ValueOf(self._has_filter);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("AnimationNode", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type FilterAction = classdb.AnimationNodeFilterAction

const (
/*Do not use filtering.*/
	FilterIgnore FilterAction = 0
/*Paths matching the filter will be allowed to pass.*/
	FilterPass FilterAction = 1
/*Paths matching the filter will be discarded.*/
	FilterStop FilterAction = 2
/*Paths matching the filter will be blended (by the blend value).*/
	FilterBlend FilterAction = 3
)
