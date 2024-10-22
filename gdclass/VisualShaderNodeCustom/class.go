package VisualShaderNodeCustom

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualShaderNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
By inheriting this class you can create a custom [VisualShader] script addon which will be automatically added to the Visual Shader Editor. The [VisualShaderNode]'s behavior is defined by overriding the provided virtual methods.
In order for the node to be registered as an editor addon, you must use the [code]@tool[/code] annotation and provide a [code]class_name[/code] for your custom script. For example:
[codeblock]
@tool
extends VisualShaderNodeCustom
class_name VisualShaderNodeNoise
[/codeblock]
	// VisualShaderNodeCustom methods that can be overridden by a [Class] that extends it.
	type VisualShaderNodeCustom interface {
		//Override this method to define the name of the associated custom node in the Visual Shader Editor's members dialog and graph.
		//Defining this method is [b]optional[/b], but recommended. If not overridden, the node will be named as "Unnamed".
		GetName() string
		//Override this method to define the description of the associated custom node in the Visual Shader Editor's members dialog.
		//Defining this method is [b]optional[/b].
		GetDescription() string
		//Override this method to define the path to the associated custom node in the Visual Shader Editor's members dialog. The path may look like [code]"MyGame/MyFunctions/Noise"[/code].
		//Defining this method is [b]optional[/b]. If not overridden, the node will be filed under the "Addons" category.
		GetCategory() string
		//Override this method to define the return icon of the associated custom node in the Visual Shader Editor's members dialog.
		//Defining this method is [b]optional[/b]. If not overridden, no return icon is shown.
		GetReturnIconType() classdb.VisualShaderNodePortType
		//Override this method to define the number of input ports of the associated custom node.
		//Defining this method is [b]required[/b]. If not overridden, the node has no input ports.
		GetInputPortCount() int
		//Override this method to define the returned type of each input port of the associated custom node (see [enum VisualShaderNode.PortType] for possible types).
		//Defining this method is [b]optional[/b], but recommended. If not overridden, input ports will return the [constant VisualShaderNode.PORT_TYPE_SCALAR] type.
		GetInputPortType(port int) classdb.VisualShaderNodePortType
		//Override this method to define the names of input ports of the associated custom node. The names are used both for the input slots in the editor and as identifiers in the shader code, and are passed in the [code]input_vars[/code] array in [method _get_code].
		//Defining this method is [b]optional[/b], but recommended. If not overridden, input ports are named as [code]"in" + str(port)[/code].
		GetInputPortName(port int) string
		//Override this method to define the default value for the specified input port. Prefer use this over [method VisualShaderNode.set_input_port_default_value].
		//Defining this method is [b]required[/b]. If not overridden, the node has no default values for their input ports.
		GetInputPortDefaultValue(port int) gd.Variant
		//Override this method to define the input port which should be connected by default when this node is created as a result of dragging a connection from an existing node to the empty space on the graph.
		//Defining this method is [b]optional[/b]. If not overridden, the connection will be created to the first valid port.
		GetDefaultInputPort(atype classdb.VisualShaderNodePortType) int
		//Override this method to define the number of output ports of the associated custom node.
		//Defining this method is [b]required[/b]. If not overridden, the node has no output ports.
		GetOutputPortCount() int
		//Override this method to define the returned type of each output port of the associated custom node (see [enum VisualShaderNode.PortType] for possible types).
		//Defining this method is [b]optional[/b], but recommended. If not overridden, output ports will return the [constant VisualShaderNode.PORT_TYPE_SCALAR] type.
		GetOutputPortType(port int) classdb.VisualShaderNodePortType
		//Override this method to define the names of output ports of the associated custom node. The names are used both for the output slots in the editor and as identifiers in the shader code, and are passed in the [code]output_vars[/code] array in [method _get_code].
		//Defining this method is [b]optional[/b], but recommended. If not overridden, output ports are named as [code]"out" + str(port)[/code].
		GetOutputPortName(port int) string
		//Override this method to define the number of the properties.
		//Defining this method is [b]optional[/b].
		GetPropertyCount() int
		//Override this method to define the names of the property of the associated custom node.
		//Defining this method is [b]optional[/b].
		GetPropertyName(index int) string
		//Override this method to define the default index of the property of the associated custom node.
		//Defining this method is [b]optional[/b].
		GetPropertyDefaultIndex(index int) int
		//Override this method to define the options inside the drop-down list property of the associated custom node.
		//Defining this method is [b]optional[/b].
		GetPropertyOptions(index int) []string
		//Override this method to define the actual shader code of the associated custom node. The shader code should be returned as a string, which can have multiple lines (the [code]"""[/code] multiline string construct can be used for convenience).
		//The [param input_vars] and [param output_vars] arrays contain the string names of the various input and output variables, as defined by [code]_get_input_*[/code] and [code]_get_output_*[/code] virtual methods in this class.
		//The output ports can be assigned values in the shader code. For example, [code]return output_vars[0] + " = " + input_vars[0] + ";"[/code].
		//You can customize the generated code based on the shader [param mode] (see [enum Shader.Mode]) and/or [param type] (see [enum VisualShader.Type]).
		//Defining this method is [b]required[/b].
		GetCode(input_vars gd.ArrayOf[gd.String], output_vars gd.ArrayOf[gd.String], mode classdb.ShaderMode, atype classdb.VisualShaderType) string
		//Override this method to add a shader code to the beginning of each shader function (once). The shader code should be returned as a string, which can have multiple lines (the [code]"""[/code] multiline string construct can be used for convenience).
		//If there are multiple custom nodes of different types which use this feature the order of each insertion is undefined.
		//You can customize the generated code based on the shader [param mode] (see [enum Shader.Mode]) and/or [param type] (see [enum VisualShader.Type]).
		//Defining this method is [b]optional[/b].
		GetFuncCode(mode classdb.ShaderMode, atype classdb.VisualShaderType) string
		//Override this method to add shader code on top of the global shader, to define your own standard library of reusable methods, varyings, constants, uniforms, etc. The shader code should be returned as a string, which can have multiple lines (the [code]"""[/code] multiline string construct can be used for convenience).
		//Be careful with this functionality as it can cause name conflicts with other custom nodes, so be sure to give the defined entities unique names.
		//You can customize the generated code based on the shader [param mode] (see [enum Shader.Mode]).
		//Defining this method is [b]optional[/b].
		GetGlobalCode(mode classdb.ShaderMode) string
		//Override this method to enable high-end mark in the Visual Shader Editor's members dialog.
		//Defining this method is [b]optional[/b]. If not overridden, it's [code]false[/code].
		IsHighend() bool
		//Override this method to prevent the node to be visible in the member dialog for the certain [param mode] (see [enum Shader.Mode]) and/or [param type] (see [enum VisualShader.Type]).
		//Defining this method is [b]optional[/b]. If not overridden, it's [code]true[/code].
		IsAvailable(mode classdb.ShaderMode, atype classdb.VisualShaderType) bool
	}

*/
type Go [1]classdb.VisualShaderNodeCustom

/*
Override this method to define the name of the associated custom node in the Visual Shader Editor's members dialog and graph.
Defining this method is [b]optional[/b], but recommended. If not overridden, the node will be named as "Unnamed".
*/
func (Go) _get_name(impl func(ptr unsafe.Pointer) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}

/*
Override this method to define the description of the associated custom node in the Visual Shader Editor's members dialog.
Defining this method is [b]optional[/b].
*/
func (Go) _get_description(impl func(ptr unsafe.Pointer) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}

/*
Override this method to define the path to the associated custom node in the Visual Shader Editor's members dialog. The path may look like [code]"MyGame/MyFunctions/Noise"[/code].
Defining this method is [b]optional[/b]. If not overridden, the node will be filed under the "Addons" category.
*/
func (Go) _get_category(impl func(ptr unsafe.Pointer) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}

/*
Override this method to define the return icon of the associated custom node in the Visual Shader Editor's members dialog.
Defining this method is [b]optional[/b]. If not overridden, no return icon is shown.
*/
func (Go) _get_return_icon_type(impl func(ptr unsafe.Pointer) classdb.VisualShaderNodePortType, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Override this method to define the number of input ports of the associated custom node.
Defining this method is [b]required[/b]. If not overridden, the node has no input ports.
*/
func (Go) _get_input_port_count(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Override this method to define the returned type of each input port of the associated custom node (see [enum VisualShaderNode.PortType] for possible types).
Defining this method is [b]optional[/b], but recommended. If not overridden, input ports will return the [constant VisualShaderNode.PORT_TYPE_SCALAR] type.
*/
func (Go) _get_input_port_type(impl func(ptr unsafe.Pointer, port int) classdb.VisualShaderNodePortType, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var port = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(port))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Override this method to define the names of input ports of the associated custom node. The names are used both for the input slots in the editor and as identifiers in the shader code, and are passed in the [code]input_vars[/code] array in [method _get_code].
Defining this method is [b]optional[/b], but recommended. If not overridden, input ports are named as [code]"in" + str(port)[/code].
*/
func (Go) _get_input_port_name(impl func(ptr unsafe.Pointer, port int) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var port = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(port))
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}

/*
Override this method to define the default value for the specified input port. Prefer use this over [method VisualShaderNode.set_input_port_default_value].
Defining this method is [b]required[/b]. If not overridden, the node has no default values for their input ports.
*/
func (Go) _get_input_port_default_value(impl func(ptr unsafe.Pointer, port int) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var port = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(port))
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}

/*
Override this method to define the input port which should be connected by default when this node is created as a result of dragging a connection from an existing node to the empty space on the graph.
Defining this method is [b]optional[/b]. If not overridden, the connection will be created to the first valid port.
*/
func (Go) _get_default_input_port(impl func(ptr unsafe.Pointer, atype classdb.VisualShaderNodePortType) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var atype = gd.UnsafeGet[classdb.VisualShaderNodePortType](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, atype)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Override this method to define the number of output ports of the associated custom node.
Defining this method is [b]required[/b]. If not overridden, the node has no output ports.
*/
func (Go) _get_output_port_count(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Override this method to define the returned type of each output port of the associated custom node (see [enum VisualShaderNode.PortType] for possible types).
Defining this method is [b]optional[/b], but recommended. If not overridden, output ports will return the [constant VisualShaderNode.PORT_TYPE_SCALAR] type.
*/
func (Go) _get_output_port_type(impl func(ptr unsafe.Pointer, port int) classdb.VisualShaderNodePortType, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var port = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(port))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Override this method to define the names of output ports of the associated custom node. The names are used both for the output slots in the editor and as identifiers in the shader code, and are passed in the [code]output_vars[/code] array in [method _get_code].
Defining this method is [b]optional[/b], but recommended. If not overridden, output ports are named as [code]"out" + str(port)[/code].
*/
func (Go) _get_output_port_name(impl func(ptr unsafe.Pointer, port int) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var port = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(port))
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}

/*
Override this method to define the number of the properties.
Defining this method is [b]optional[/b].
*/
func (Go) _get_property_count(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Override this method to define the names of the property of the associated custom node.
Defining this method is [b]optional[/b].
*/
func (Go) _get_property_name(impl func(ptr unsafe.Pointer, index int) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}

/*
Override this method to define the default index of the property of the associated custom node.
Defining this method is [b]optional[/b].
*/
func (Go) _get_property_default_index(impl func(ptr unsafe.Pointer, index int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Override this method to define the options inside the drop-down list property of the associated custom node.
Defining this method is [b]optional[/b].
*/
func (Go) _get_property_options(impl func(ptr unsafe.Pointer, index int) []string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		gd.UnsafeSet(p_back, mmm.End(gc.PackedStringSlice(ret)))
		gc.End()
	}
}

/*
Override this method to define the actual shader code of the associated custom node. The shader code should be returned as a string, which can have multiple lines (the [code]"""[/code] multiline string construct can be used for convenience).
The [param input_vars] and [param output_vars] arrays contain the string names of the various input and output variables, as defined by [code]_get_input_*[/code] and [code]_get_output_*[/code] virtual methods in this class.
The output ports can be assigned values in the shader code. For example, [code]return output_vars[0] + " = " + input_vars[0] + ";"[/code].
You can customize the generated code based on the shader [param mode] (see [enum Shader.Mode]) and/or [param type] (see [enum VisualShader.Type]).
Defining this method is [b]required[/b].
*/
func (Go) _get_code(impl func(ptr unsafe.Pointer, input_vars gd.ArrayOf[gd.String], output_vars gd.ArrayOf[gd.String], mode classdb.ShaderMode, atype classdb.VisualShaderType) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var input_vars = gd.TypedArray[gd.String](mmm.Let[gd.Array](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0)))
		var output_vars = gd.TypedArray[gd.String](mmm.Let[gd.Array](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1)))
		var mode = gd.UnsafeGet[classdb.ShaderMode](p_args,2)
		var atype = gd.UnsafeGet[classdb.VisualShaderType](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, input_vars, output_vars, mode, atype)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}

/*
Override this method to add a shader code to the beginning of each shader function (once). The shader code should be returned as a string, which can have multiple lines (the [code]"""[/code] multiline string construct can be used for convenience).
If there are multiple custom nodes of different types which use this feature the order of each insertion is undefined.
You can customize the generated code based on the shader [param mode] (see [enum Shader.Mode]) and/or [param type] (see [enum VisualShader.Type]).
Defining this method is [b]optional[/b].
*/
func (Go) _get_func_code(impl func(ptr unsafe.Pointer, mode classdb.ShaderMode, atype classdb.VisualShaderType) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var mode = gd.UnsafeGet[classdb.ShaderMode](p_args,0)
		var atype = gd.UnsafeGet[classdb.VisualShaderType](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, mode, atype)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}

/*
Override this method to add shader code on top of the global shader, to define your own standard library of reusable methods, varyings, constants, uniforms, etc. The shader code should be returned as a string, which can have multiple lines (the [code]"""[/code] multiline string construct can be used for convenience).
Be careful with this functionality as it can cause name conflicts with other custom nodes, so be sure to give the defined entities unique names.
You can customize the generated code based on the shader [param mode] (see [enum Shader.Mode]).
Defining this method is [b]optional[/b].
*/
func (Go) _get_global_code(impl func(ptr unsafe.Pointer, mode classdb.ShaderMode) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var mode = gd.UnsafeGet[classdb.ShaderMode](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, mode)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}

/*
Override this method to enable high-end mark in the Visual Shader Editor's members dialog.
Defining this method is [b]optional[/b]. If not overridden, it's [code]false[/code].
*/
func (Go) _is_highend(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Override this method to prevent the node to be visible in the member dialog for the certain [param mode] (see [enum Shader.Mode]) and/or [param type] (see [enum VisualShader.Type]).
Defining this method is [b]optional[/b]. If not overridden, it's [code]true[/code].
*/
func (Go) _is_available(impl func(ptr unsafe.Pointer, mode classdb.ShaderMode, atype classdb.VisualShaderType) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var mode = gd.UnsafeGet[classdb.ShaderMode](p_args,0)
		var atype = gd.UnsafeGet[classdb.VisualShaderType](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, mode, atype)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Returns the selected index of the drop-down list option within a graph. You may use this function to define the specific behavior in the [method _get_code] or [method _get_global_code].
*/
func (self Go) GetOptionIndex(option int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetOptionIndex(gd.Int(option))))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeCustom
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VisualShaderNodeCustom"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Override this method to define the name of the associated custom node in the Visual Shader Editor's members dialog and graph.
Defining this method is [b]optional[/b], but recommended. If not overridden, the node will be named as "Unnamed".
*/
func (class) _get_name(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Override this method to define the description of the associated custom node in the Visual Shader Editor's members dialog.
Defining this method is [b]optional[/b].
*/
func (class) _get_description(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Override this method to define the path to the associated custom node in the Visual Shader Editor's members dialog. The path may look like [code]"MyGame/MyFunctions/Noise"[/code].
Defining this method is [b]optional[/b]. If not overridden, the node will be filed under the "Addons" category.
*/
func (class) _get_category(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Override this method to define the return icon of the associated custom node in the Visual Shader Editor's members dialog.
Defining this method is [b]optional[/b]. If not overridden, no return icon is shown.
*/
func (class) _get_return_icon_type(impl func(ptr unsafe.Pointer) classdb.VisualShaderNodePortType, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Override this method to define the number of input ports of the associated custom node.
Defining this method is [b]required[/b]. If not overridden, the node has no input ports.
*/
func (class) _get_input_port_count(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Override this method to define the returned type of each input port of the associated custom node (see [enum VisualShaderNode.PortType] for possible types).
Defining this method is [b]optional[/b], but recommended. If not overridden, input ports will return the [constant VisualShaderNode.PORT_TYPE_SCALAR] type.
*/
func (class) _get_input_port_type(impl func(ptr unsafe.Pointer, port gd.Int) classdb.VisualShaderNodePortType, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var port = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, port)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Override this method to define the names of input ports of the associated custom node. The names are used both for the input slots in the editor and as identifiers in the shader code, and are passed in the [code]input_vars[/code] array in [method _get_code].
Defining this method is [b]optional[/b], but recommended. If not overridden, input ports are named as [code]"in" + str(port)[/code].
*/
func (class) _get_input_port_name(impl func(ptr unsafe.Pointer, port gd.Int) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var port = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, port)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Override this method to define the default value for the specified input port. Prefer use this over [method VisualShaderNode.set_input_port_default_value].
Defining this method is [b]required[/b]. If not overridden, the node has no default values for their input ports.
*/
func (class) _get_input_port_default_value(impl func(ptr unsafe.Pointer, port gd.Int) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var port = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, port)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Override this method to define the input port which should be connected by default when this node is created as a result of dragging a connection from an existing node to the empty space on the graph.
Defining this method is [b]optional[/b]. If not overridden, the connection will be created to the first valid port.
*/
func (class) _get_default_input_port(impl func(ptr unsafe.Pointer, atype classdb.VisualShaderNodePortType) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var atype = gd.UnsafeGet[classdb.VisualShaderNodePortType](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, atype)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Override this method to define the number of output ports of the associated custom node.
Defining this method is [b]required[/b]. If not overridden, the node has no output ports.
*/
func (class) _get_output_port_count(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Override this method to define the returned type of each output port of the associated custom node (see [enum VisualShaderNode.PortType] for possible types).
Defining this method is [b]optional[/b], but recommended. If not overridden, output ports will return the [constant VisualShaderNode.PORT_TYPE_SCALAR] type.
*/
func (class) _get_output_port_type(impl func(ptr unsafe.Pointer, port gd.Int) classdb.VisualShaderNodePortType, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var port = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, port)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Override this method to define the names of output ports of the associated custom node. The names are used both for the output slots in the editor and as identifiers in the shader code, and are passed in the [code]output_vars[/code] array in [method _get_code].
Defining this method is [b]optional[/b], but recommended. If not overridden, output ports are named as [code]"out" + str(port)[/code].
*/
func (class) _get_output_port_name(impl func(ptr unsafe.Pointer, port gd.Int) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var port = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, port)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Override this method to define the number of the properties.
Defining this method is [b]optional[/b].
*/
func (class) _get_property_count(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Override this method to define the names of the property of the associated custom node.
Defining this method is [b]optional[/b].
*/
func (class) _get_property_name(impl func(ptr unsafe.Pointer, index gd.Int) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Override this method to define the default index of the property of the associated custom node.
Defining this method is [b]optional[/b].
*/
func (class) _get_property_default_index(impl func(ptr unsafe.Pointer, index gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Override this method to define the options inside the drop-down list property of the associated custom node.
Defining this method is [b]optional[/b].
*/
func (class) _get_property_options(impl func(ptr unsafe.Pointer, index gd.Int) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Override this method to define the actual shader code of the associated custom node. The shader code should be returned as a string, which can have multiple lines (the [code]"""[/code] multiline string construct can be used for convenience).
The [param input_vars] and [param output_vars] arrays contain the string names of the various input and output variables, as defined by [code]_get_input_*[/code] and [code]_get_output_*[/code] virtual methods in this class.
The output ports can be assigned values in the shader code. For example, [code]return output_vars[0] + " = " + input_vars[0] + ";"[/code].
You can customize the generated code based on the shader [param mode] (see [enum Shader.Mode]) and/or [param type] (see [enum VisualShader.Type]).
Defining this method is [b]required[/b].
*/
func (class) _get_code(impl func(ptr unsafe.Pointer, input_vars gd.ArrayOf[gd.String], output_vars gd.ArrayOf[gd.String], mode classdb.ShaderMode, atype classdb.VisualShaderType) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var input_vars = gd.TypedArray[gd.String](mmm.Let[gd.Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0)))
		var output_vars = gd.TypedArray[gd.String](mmm.Let[gd.Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1)))
		var mode = gd.UnsafeGet[classdb.ShaderMode](p_args,2)
		var atype = gd.UnsafeGet[classdb.VisualShaderType](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, input_vars, output_vars, mode, atype)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Override this method to add a shader code to the beginning of each shader function (once). The shader code should be returned as a string, which can have multiple lines (the [code]"""[/code] multiline string construct can be used for convenience).
If there are multiple custom nodes of different types which use this feature the order of each insertion is undefined.
You can customize the generated code based on the shader [param mode] (see [enum Shader.Mode]) and/or [param type] (see [enum VisualShader.Type]).
Defining this method is [b]optional[/b].
*/
func (class) _get_func_code(impl func(ptr unsafe.Pointer, mode classdb.ShaderMode, atype classdb.VisualShaderType) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var mode = gd.UnsafeGet[classdb.ShaderMode](p_args,0)
		var atype = gd.UnsafeGet[classdb.VisualShaderType](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, mode, atype)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Override this method to add shader code on top of the global shader, to define your own standard library of reusable methods, varyings, constants, uniforms, etc. The shader code should be returned as a string, which can have multiple lines (the [code]"""[/code] multiline string construct can be used for convenience).
Be careful with this functionality as it can cause name conflicts with other custom nodes, so be sure to give the defined entities unique names.
You can customize the generated code based on the shader [param mode] (see [enum Shader.Mode]).
Defining this method is [b]optional[/b].
*/
func (class) _get_global_code(impl func(ptr unsafe.Pointer, mode classdb.ShaderMode) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var mode = gd.UnsafeGet[classdb.ShaderMode](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, mode)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Override this method to enable high-end mark in the Visual Shader Editor's members dialog.
Defining this method is [b]optional[/b]. If not overridden, it's [code]false[/code].
*/
func (class) _is_highend(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Override this method to prevent the node to be visible in the member dialog for the certain [param mode] (see [enum Shader.Mode]) and/or [param type] (see [enum VisualShader.Type]).
Defining this method is [b]optional[/b]. If not overridden, it's [code]true[/code].
*/
func (class) _is_available(impl func(ptr unsafe.Pointer, mode classdb.ShaderMode, atype classdb.VisualShaderType) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var mode = gd.UnsafeGet[classdb.ShaderMode](p_args,0)
		var atype = gd.UnsafeGet[classdb.VisualShaderType](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, mode, atype)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Returns the selected index of the drop-down list option within a graph. You may use this function to define the specific behavior in the [method _get_code] or [method _get_global_code].
*/
//go:nosplit
func (self class) GetOptionIndex(option gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, option)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeCustom.Bind_get_option_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeCustom() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeCustom() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNode() VisualShaderNode.GD { return *((*VisualShaderNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNode() VisualShaderNode.Go { return *((*VisualShaderNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_name": return reflect.ValueOf(self._get_name);
	case "_get_description": return reflect.ValueOf(self._get_description);
	case "_get_category": return reflect.ValueOf(self._get_category);
	case "_get_return_icon_type": return reflect.ValueOf(self._get_return_icon_type);
	case "_get_input_port_count": return reflect.ValueOf(self._get_input_port_count);
	case "_get_input_port_type": return reflect.ValueOf(self._get_input_port_type);
	case "_get_input_port_name": return reflect.ValueOf(self._get_input_port_name);
	case "_get_input_port_default_value": return reflect.ValueOf(self._get_input_port_default_value);
	case "_get_default_input_port": return reflect.ValueOf(self._get_default_input_port);
	case "_get_output_port_count": return reflect.ValueOf(self._get_output_port_count);
	case "_get_output_port_type": return reflect.ValueOf(self._get_output_port_type);
	case "_get_output_port_name": return reflect.ValueOf(self._get_output_port_name);
	case "_get_property_count": return reflect.ValueOf(self._get_property_count);
	case "_get_property_name": return reflect.ValueOf(self._get_property_name);
	case "_get_property_default_index": return reflect.ValueOf(self._get_property_default_index);
	case "_get_property_options": return reflect.ValueOf(self._get_property_options);
	case "_get_code": return reflect.ValueOf(self._get_code);
	case "_get_func_code": return reflect.ValueOf(self._get_func_code);
	case "_get_global_code": return reflect.ValueOf(self._get_global_code);
	case "_is_highend": return reflect.ValueOf(self._is_highend);
	case "_is_available": return reflect.ValueOf(self._is_available);
	default: return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_get_name": return reflect.ValueOf(self._get_name);
	case "_get_description": return reflect.ValueOf(self._get_description);
	case "_get_category": return reflect.ValueOf(self._get_category);
	case "_get_return_icon_type": return reflect.ValueOf(self._get_return_icon_type);
	case "_get_input_port_count": return reflect.ValueOf(self._get_input_port_count);
	case "_get_input_port_type": return reflect.ValueOf(self._get_input_port_type);
	case "_get_input_port_name": return reflect.ValueOf(self._get_input_port_name);
	case "_get_input_port_default_value": return reflect.ValueOf(self._get_input_port_default_value);
	case "_get_default_input_port": return reflect.ValueOf(self._get_default_input_port);
	case "_get_output_port_count": return reflect.ValueOf(self._get_output_port_count);
	case "_get_output_port_type": return reflect.ValueOf(self._get_output_port_type);
	case "_get_output_port_name": return reflect.ValueOf(self._get_output_port_name);
	case "_get_property_count": return reflect.ValueOf(self._get_property_count);
	case "_get_property_name": return reflect.ValueOf(self._get_property_name);
	case "_get_property_default_index": return reflect.ValueOf(self._get_property_default_index);
	case "_get_property_options": return reflect.ValueOf(self._get_property_options);
	case "_get_code": return reflect.ValueOf(self._get_code);
	case "_get_func_code": return reflect.ValueOf(self._get_func_code);
	case "_get_global_code": return reflect.ValueOf(self._get_global_code);
	case "_is_highend": return reflect.ValueOf(self._is_highend);
	case "_is_available": return reflect.ValueOf(self._is_available);
	default: return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}
func init() {classdb.Register("VisualShaderNodeCustom", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
