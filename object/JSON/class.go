package JSON

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
type Simple [1]classdb.JSON
func (self Simple) Stringify(data gd.Variant, indent string, sort_keys bool, full_precision bool) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).Stringify(gc, data, gc.String(indent), sort_keys, full_precision).String())
}
func (self Simple) ParseString(json_string string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).ParseString(gc, gc.String(json_string)))
}
func (self Simple) Parse(json_text string, keep_text bool) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Parse(gc.String(json_text), keep_text))
}
func (self Simple) GetData() gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetData(gc))
}
func (self Simple) SetData(data gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetData(data)
}
func (self Simple) GetParsedText() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetParsedText(gc).String())
}
func (self Simple) GetErrorLine() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetErrorLine()))
}
func (self Simple) GetErrorMessage() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetErrorMessage(gc).String())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.JSON
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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
func (self class) Stringify(ctx gd.Lifetime, data gd.Variant, indent gd.String, sort_keys bool, full_precision bool) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(data))
	callframe.Arg(frame, mmm.Get(indent))
	callframe.Arg(frame, sort_keys)
	callframe.Arg(frame, full_precision)
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.JSON.Bind_stringify, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Attempts to parse the [param json_string] provided and returns the parsed data. Returns [code]null[/code] if parse failed.
*/
//go:nosplit
func (self class) ParseString(ctx gd.Lifetime, json_string gd.String) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(json_string))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.JSON.Bind_parse_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
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
func (self class) Parse(json_text gd.String, keep_text bool) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(json_text))
	callframe.Arg(frame, keep_text)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.JSON.Bind_parse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetData(ctx gd.Lifetime) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.JSON.Bind_get_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetData(data gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.JSON.Bind_set_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Return the text parsed by [method parse] (requires passing [code]keep_text[/code] to [method parse]).
*/
//go:nosplit
func (self class) GetParsedText(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.JSON.Bind_get_parsed_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]0[/code] if the last call to [method parse] was successful, or the line number where the parse failed.
*/
//go:nosplit
func (self class) GetErrorLine() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.JSON.Bind_get_error_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an empty string if the last call to [method parse] was successful, or the error message if it failed.
*/
//go:nosplit
func (self class) GetErrorMessage(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.JSON.Bind_get_error_message, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsJSON() Expert { return self[0].AsJSON() }


//go:nosplit
func (self Simple) AsJSON() Simple { return self[0].AsJSON() }


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
func init() {classdb.Register("JSON", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
