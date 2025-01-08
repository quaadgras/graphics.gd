// Package JSON provides methods for working with JSON object instances.
package JSON

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
The [JSON] class enables all data types to be converted to and from a JSON string. This is useful for serializing data, e.g. to save to a file or send over the network.
[method stringify] is used to convert any data type into a JSON string.
[method parse] is used to convert any existing JSON data into a [Variant] that can be used within Godot. If successfully parsed, use [member data] to retrieve the [Variant], and use [code]typeof[/code] to check if the Variant's type is what you expect. JSON Objects are converted into a [Dictionary], but JSON data can be used to store [Array]s, numbers, [String]s and even just a boolean.
[b]Example[/b]
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
	    print(data_received) # Prints array
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
- Certain errors, such as invalid Unicode sequences, do not cause a parser error. Instead, the string is cleansed and an error is logged to the console.
*/
type Instance [1]gdclass.JSON

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
func Stringify(data any) string {
	self := Instance{}
	return string(class(self).Stringify(gd.NewVariant(data), gd.NewString(""), true, false).String())
}

/*
Attempts to parse the [param json_string] provided and returns the parsed data. Returns [code]null[/code] if parse failed.
*/
func ParseString(json_string string) any {
	self := Instance{}
	return any(class(self).ParseString(gd.NewString(json_string)).Interface())
}

/*
Attempts to parse the [param json_text] provided.
Returns an [enum Error]. If the parse was successful, it returns [constant OK] and the result can be retrieved using [member data]. If unsuccessful, use [method get_error_line] and [method get_error_message] to identify the source of the failure.
Non-static variant of [method parse_string], if you want custom error handling.
The optional [param keep_text] argument instructs the parser to keep a copy of the original text. This text can be obtained later by using the [method get_parsed_text] function and is used when saving the resource (instead of generating new text from [member data]).
*/
func (self Instance) Parse(json_text string) error {
	return error(class(self).Parse(gd.NewString(json_text), false))
}

/*
Return the text parsed by [method parse] (requires passing [code]keep_text[/code] to [method parse]).
*/
func (self Instance) GetParsedText() string {
	return string(class(self).GetParsedText().String())
}

/*
Returns [code]0[/code] if the last call to [method parse] was successful, or the line number where the parse failed.
*/
func (self Instance) GetErrorLine() int {
	return int(int(class(self).GetErrorLine()))
}

/*
Returns an empty string if the last call to [method parse] was successful, or the error message if it failed.
*/
func (self Instance) GetErrorMessage() string {
	return string(class(self).GetErrorMessage().String())
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
	class(self).SetData(gd.NewVariant(value))
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
func (self class) Stringify(data gd.Variant, indent gd.String, sort_keys bool, full_precision bool) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(data))
	callframe.Arg(frame, pointers.Get(indent))
	callframe.Arg(frame, sort_keys)
	callframe.Arg(frame, full_precision)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSON.Bind_stringify, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Attempts to parse the [param json_string] provided and returns the parsed data. Returns [code]null[/code] if parse failed.
*/
//go:nosplit
func (self class) ParseString(json_string gd.String) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(json_string))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSON.Bind_parse_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
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
func (self class) Parse(json_text gd.String, keep_text bool) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(json_text))
	callframe.Arg(frame, keep_text)
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSON.Bind_parse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetData() gd.Variant {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSON.Bind_get_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetData(data gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(data))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSON.Bind_set_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Return the text parsed by [method parse] (requires passing [code]keep_text[/code] to [method parse]).
*/
//go:nosplit
func (self class) GetParsedText() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSON.Bind_get_parsed_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]0[/code] if the last call to [method parse] was successful, or the line number where the parse failed.
*/
//go:nosplit
func (self class) GetErrorLine() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSON.Bind_get_error_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an empty string if the last call to [method parse] was successful, or the error message if it failed.
*/
//go:nosplit
func (self class) GetErrorMessage() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSON.Bind_get_error_message, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
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

type Error int

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
