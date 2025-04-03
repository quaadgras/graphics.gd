// Package JSON provides methods for working with JSON object instances.
package JSON

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
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
The [JSON] class enables all data types to be converted to and from a JSON string. This is useful for serializing data, e.g. to save to a file or send over the network.
[method stringify] is used to convert any data type into a JSON string.
[method parse] is used to convert any existing JSON data into a [Variant] that can be used within Godot. If successfully parsed, use [member data] to retrieve the [Variant], and use [method @GlobalScope.typeof] to check if the Variant's type is what you expect. JSON Objects are converted into a [Dictionary], but JSON data can be used to store [Array]s, numbers, [String]s and even just a boolean.
[codeblock]
var data_to_send = ["a", "b", "c"]
var json_string = JSON.stringify(data_to_send)
# Save data
# ...
# Retrieve data
var json = JSON.new()
var error = json.parse(json_string)
if error == OK:

	var data_received = json.data
	if typeof(data_received) == TYPE_ARRAY:
	    print(data_received) # Prints the array.
	else:
	    print("Unexpected data")

else:

	print("JSON Parse Error: ", json.get_error_message(), " in ", json_string, " at line ", json.get_error_line())

[/codeblock]
Alternatively, you can parse strings using the static [method parse_string] method, but it doesn't handle errors.
[codeblock]
var data = JSON.parse_string(json_string) # Returns null if parsing failed.
[/codeblock]
[b]Note:[/b] Both parse methods do not fully comply with the JSON specification:
- Trailing commas in arrays or objects are ignored, instead of causing a parser error.
- New line and tab characters are accepted in string literals, and are treated like their corresponding escape sequences [code]\n[/code] and [code]\t[/code].
- Numbers are parsed using [method String.to_float] which is generally more lax than the JSON specification.
- Certain errors, such as invalid Unicode sequences, do not cause a parser error. Instead, the string is cleaned up and an error is logged to the console.
*/
type Instance [1]gdclass.JSON
type Expanded [1]gdclass.JSON

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsJSON() Instance
}

/*
Converts a [Variant] var to JSON text and returns the result. Useful for serializing data to store or send over the network.
[b]Note:[/b] The JSON specification does not define integer or float types, but only a [i]number[/i] type. Therefore, converting a Variant to JSON text will convert all numerical values to [float] types.
[b]Note:[/b] If [param full_precision] is [code]true[/code], when stringifying floats, the unreliable digits are stringified in addition to the reliable digits to guarantee exact decoding.
The [param indent] parameter controls if and how something is indented; its contents will be used where there should be an indent in the output. Even spaces like [code]"   "[/code] will work. [code]\t[/code] and [code]\n[/code] can also be used for a tab indent, or to make a newline for each indent respectively.
[b]Example output:[/b]
[codeblock]
## JSON.stringify(my_dictionary)
{"name":"my_dictionary","version":"1.0.0","entities":[{"name":"entity_0","value":"value_0"},{"name":"entity_1","value":"value_1"}]}

## JSON.stringify(my_dictionary, "\t")

	{
	    "name": "my_dictionary",
	    "version": "1.0.0",
	    "entities": [
	        {
	            "name": "entity_0",
	            "value": "value_0"
	        },
	        {
	            "name": "entity_1",
	            "value": "value_1"
	        }
	    ]
	}

## JSON.stringify(my_dictionary, "...")
{
..."name": "my_dictionary",
..."version": "1.0.0",
..."entities": [
......{
........."name": "entity_0",
........."value": "value_0"
......},
......{
........."name": "entity_1",
........."value": "value_1"
......}
...]
}
[/codeblock]
*/
func Stringify(data any, indent string, full_precision bool) string { //gd:JSON.stringify
	self := Instance{}
	return string(Advanced(self).Stringify(variant.New(data), String.New(indent), true, full_precision).String())
}

/*
Converts a [Variant] var to JSON text and returns the result. Useful for serializing data to store or send over the network.
[b]Note:[/b] The JSON specification does not define integer or float types, but only a [i]number[/i] type. Therefore, converting a Variant to JSON text will convert all numerical values to [float] types.
[b]Note:[/b] If [param full_precision] is [code]true[/code], when stringifying floats, the unreliable digits are stringified in addition to the reliable digits to guarantee exact decoding.
The [param indent] parameter controls if and how something is indented; its contents will be used where there should be an indent in the output. Even spaces like [code]"   "[/code] will work. [code]\t[/code] and [code]\n[/code] can also be used for a tab indent, or to make a newline for each indent respectively.
[b]Example output:[/b]
[codeblock]
## JSON.stringify(my_dictionary)
{"name":"my_dictionary","version":"1.0.0","entities":[{"name":"entity_0","value":"value_0"},{"name":"entity_1","value":"value_1"}]}

## JSON.stringify(my_dictionary, "\t")

	{
	    "name": "my_dictionary",
	    "version": "1.0.0",
	    "entities": [
	        {
	            "name": "entity_0",
	            "value": "value_0"
	        },
	        {
	            "name": "entity_1",
	            "value": "value_1"
	        }
	    ]
	}

## JSON.stringify(my_dictionary, "...")
{
..."name": "my_dictionary",
..."version": "1.0.0",
..."entities": [
......{
........."name": "entity_0",
........."value": "value_0"
......},
......{
........."name": "entity_1",
........."value": "value_1"
......}
...]
}
[/codeblock]
*/
func StringifyExpanded(data any, indent string, sort_keys bool, full_precision bool) string { //gd:JSON.stringify
	self := Instance{}
	return string(Advanced(self).Stringify(variant.New(data), String.New(indent), sort_keys, full_precision).String())
}

/*
Attempts to parse the [param json_string] provided and returns the parsed data. Returns [code]null[/code] if parse failed.
*/
func ParseString(json_string string) any { //gd:JSON.parse_string
	self := Instance{}
	return any(Advanced(self).ParseString(String.New(json_string)).Interface())
}

/*
Attempts to parse the [param json_text] provided.
Returns an [enum Error]. If the parse was successful, it returns [constant OK] and the result can be retrieved using [member data]. If unsuccessful, use [method get_error_line] and [method get_error_message] to identify the source of the failure.
Non-static variant of [method parse_string], if you want custom error handling.
The optional [param keep_text] argument instructs the parser to keep a copy of the original text. This text can be obtained later by using the [method get_parsed_text] function and is used when saving the resource (instead of generating new text from [member data]).
*/
func (self Instance) Parse(json_text string) error { //gd:JSON.parse
	return error(gd.ToError(Advanced(self).Parse(String.New(json_text), false)))
}

/*
Attempts to parse the [param json_text] provided.
Returns an [enum Error]. If the parse was successful, it returns [constant OK] and the result can be retrieved using [member data]. If unsuccessful, use [method get_error_line] and [method get_error_message] to identify the source of the failure.
Non-static variant of [method parse_string], if you want custom error handling.
The optional [param keep_text] argument instructs the parser to keep a copy of the original text. This text can be obtained later by using the [method get_parsed_text] function and is used when saving the resource (instead of generating new text from [member data]).
*/
func (self Expanded) Parse(json_text string, keep_text bool) error { //gd:JSON.parse
	return error(gd.ToError(Advanced(self).Parse(String.New(json_text), keep_text)))
}

/*
Return the text parsed by [method parse] (requires passing [code]keep_text[/code] to [method parse]).
*/
func (self Instance) GetParsedText() string { //gd:JSON.get_parsed_text
	return string(Advanced(self).GetParsedText().String())
}

/*
Returns [code]0[/code] if the last call to [method parse] was successful, or the line number where the parse failed.
*/
func (self Instance) GetErrorLine() int { //gd:JSON.get_error_line
	return int(int(Advanced(self).GetErrorLine()))
}

/*
Returns an empty string if the last call to [method parse] was successful, or the error message if it failed.
*/
func (self Instance) GetErrorMessage() string { //gd:JSON.get_error_message
	return string(Advanced(self).GetErrorMessage().String())
}

/*
Converts a native engine type to a JSON-compliant value.
By default, objects are ignored for security reasons, unless [param full_objects] is [code]true[/code].
You can convert a native value to a JSON string like this:
[codeblock]
func encode_data(value, full_objects = false):

	return JSON.stringify(JSON.from_native(value, full_objects))

[/codeblock]
*/
func FromNative(v any, full_objects bool) any { //gd:JSON.from_native
	self := Instance{}
	return any(Advanced(self).FromNative(variant.New(v), full_objects).Interface())
}

/*
Converts a native engine type to a JSON-compliant value.
By default, objects are ignored for security reasons, unless [param full_objects] is [code]true[/code].
You can convert a native value to a JSON string like this:
[codeblock]
func encode_data(value, full_objects = false):

	return JSON.stringify(JSON.from_native(value, full_objects))

[/codeblock]
*/
func FromNativeExpanded(v any, full_objects bool) any { //gd:JSON.from_native
	self := Instance{}
	return any(Advanced(self).FromNative(variant.New(v), full_objects).Interface())
}

/*
Converts a JSON-compliant value that was created with [method from_native] back to native engine types.
By default, objects are ignored for security reasons, unless [param allow_objects] is [code]true[/code].
You can convert a JSON string back to a native value like this:
[codeblock]
func decode_data(string, allow_objects = false):

	return JSON.to_native(JSON.parse_string(string), allow_objects)

[/codeblock]
*/
func ToNative(json any, allow_objects bool) any { //gd:JSON.to_native
	self := Instance{}
	return any(Advanced(self).ToNative(variant.New(json), allow_objects).Interface())
}

/*
Converts a JSON-compliant value that was created with [method from_native] back to native engine types.
By default, objects are ignored for security reasons, unless [param allow_objects] is [code]true[/code].
You can convert a JSON string back to a native value like this:
[codeblock]
func decode_data(string, allow_objects = false):

	return JSON.to_native(JSON.parse_string(string), allow_objects)

[/codeblock]
*/
func ToNativeExpanded(json any, allow_objects bool) any { //gd:JSON.to_native
	self := Instance{}
	return any(Advanced(self).ToNative(variant.New(json), allow_objects).Interface())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.JSON

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("JSON"))
	casted := Instance{*(*gdclass.JSON)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Data() any {
	return any(class(self).GetData().Interface())
}

func (self Instance) SetData(value any) {
	class(self).SetData(variant.New(value))
}

/*
Converts a [Variant] var to JSON text and returns the result. Useful for serializing data to store or send over the network.
[b]Note:[/b] The JSON specification does not define integer or float types, but only a [i]number[/i] type. Therefore, converting a Variant to JSON text will convert all numerical values to [float] types.
[b]Note:[/b] If [param full_precision] is [code]true[/code], when stringifying floats, the unreliable digits are stringified in addition to the reliable digits to guarantee exact decoding.
The [param indent] parameter controls if and how something is indented; its contents will be used where there should be an indent in the output. Even spaces like [code]"   "[/code] will work. [code]\t[/code] and [code]\n[/code] can also be used for a tab indent, or to make a newline for each indent respectively.
[b]Example output:[/b]
[codeblock]
## JSON.stringify(my_dictionary)
{"name":"my_dictionary","version":"1.0.0","entities":[{"name":"entity_0","value":"value_0"},{"name":"entity_1","value":"value_1"}]}

## JSON.stringify(my_dictionary, "\t")
{
    "name": "my_dictionary",
    "version": "1.0.0",
    "entities": [
        {
            "name": "entity_0",
            "value": "value_0"
        },
        {
            "name": "entity_1",
            "value": "value_1"
        }
    ]
}

## JSON.stringify(my_dictionary, "...")
{
..."name": "my_dictionary",
..."version": "1.0.0",
..."entities": [
......{
........."name": "entity_0",
........."value": "value_0"
......},
......{
........."name": "entity_1",
........."value": "value_1"
......}
...]
}
[/codeblock]
*/
//go:nosplit
func (self class) Stringify(data variant.Any, indent String.Readable, sort_keys bool, full_precision bool) String.Readable { //gd:JSON.stringify
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(data)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(indent)))
	callframe.Arg(frame, sort_keys)
	callframe.Arg(frame, full_precision)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSON.Bind_stringify, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Attempts to parse the [param json_string] provided and returns the parsed data. Returns [code]null[/code] if parse failed.
*/
//go:nosplit
func (self class) ParseString(json_string String.Readable) variant.Any { //gd:JSON.parse_string
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(json_string)))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSON.Bind_parse_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Attempts to parse the [param json_text] provided.
Returns an [enum Error]. If the parse was successful, it returns [constant OK] and the result can be retrieved using [member data]. If unsuccessful, use [method get_error_line] and [method get_error_message] to identify the source of the failure.
Non-static variant of [method parse_string], if you want custom error handling.
The optional [param keep_text] argument instructs the parser to keep a copy of the original text. This text can be obtained later by using the [method get_parsed_text] function and is used when saving the resource (instead of generating new text from [member data]).
*/
//go:nosplit
func (self class) Parse(json_text String.Readable, keep_text bool) Error.Code { //gd:JSON.parse
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(json_text)))
	callframe.Arg(frame, keep_text)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSON.Bind_parse, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetData() variant.Any { //gd:JSON.get_data
	var frame = callframe.New()
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSON.Bind_get_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetData(data variant.Any) { //gd:JSON.set_data
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(data)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSON.Bind_set_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Return the text parsed by [method parse] (requires passing [code]keep_text[/code] to [method parse]).
*/
//go:nosplit
func (self class) GetParsedText() String.Readable { //gd:JSON.get_parsed_text
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSON.Bind_get_parsed_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]0[/code] if the last call to [method parse] was successful, or the line number where the parse failed.
*/
//go:nosplit
func (self class) GetErrorLine() int64 { //gd:JSON.get_error_line
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSON.Bind_get_error_line, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an empty string if the last call to [method parse] was successful, or the error message if it failed.
*/
//go:nosplit
func (self class) GetErrorMessage() String.Readable { //gd:JSON.get_error_message
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSON.Bind_get_error_message, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Converts a native engine type to a JSON-compliant value.
By default, objects are ignored for security reasons, unless [param full_objects] is [code]true[/code].
You can convert a native value to a JSON string like this:
[codeblock]
func encode_data(value, full_objects = false):
    return JSON.stringify(JSON.from_native(value, full_objects))
[/codeblock]
*/
//go:nosplit
func (self class) FromNative(v variant.Any, full_objects bool) variant.Any { //gd:JSON.from_native
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(v)))
	callframe.Arg(frame, full_objects)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSON.Bind_from_native, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Converts a JSON-compliant value that was created with [method from_native] back to native engine types.
By default, objects are ignored for security reasons, unless [param allow_objects] is [code]true[/code].
You can convert a JSON string back to a native value like this:
[codeblock]
func decode_data(string, allow_objects = false):
    return JSON.to_native(JSON.parse_string(string), allow_objects)
[/codeblock]
*/
//go:nosplit
func (self class) ToNative(json variant.Any, allow_objects bool) variant.Any { //gd:JSON.to_native
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(json)))
	callframe.Arg(frame, allow_objects)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSON.Bind_to_native, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}
func (self class) AsJSON() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsJSON() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("JSON", func(ptr gd.Object) any { return [1]gdclass.JSON{*(*gdclass.JSON)(unsafe.Pointer(&ptr))} })
}
