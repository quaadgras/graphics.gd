// Package RegEx provides methods for working with RegEx object instances.
package RegEx

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
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
A regular expression (or regex) is a compact language that can be used to recognize strings that follow a specific pattern, such as URLs, email addresses, complete sentences, etc. For example, a regex of [code]ab[0-9][/code] would find any string that is [code]ab[/code] followed by any number from [code]0[/code] to [code]9[/code]. For a more in-depth look, you can easily find various tutorials and detailed explanations on the Internet.
To begin, the RegEx object needs to be compiled with the search pattern using [method compile] before it can be used.
[codeblock]
var regex = RegEx.new()
regex.compile("\\w-(\\d+)")
[/codeblock]
The search pattern must be escaped first for GDScript before it is escaped for the expression. For example, [code]compile("\\d+")[/code] would be read by RegEx as [code]\d+[/code]. Similarly, [code]compile("\"(?:\\\\.|[^\"])*\"")[/code] would be read as [code]"(?:\\.|[^"])*"[/code]. In GDScript, you can also use raw string literals (r-strings). For example, [code]compile(r'"(?:\\.|[^"])*"')[/code] would be read the same.
Using [method search], you can find the pattern within the given text. If a pattern is found, [RegExMatch] is returned and you can retrieve details of the results using methods such as [method RegExMatch.get_string] and [method RegExMatch.get_start].
[codeblock]
var regex = RegEx.new()
regex.compile("\\w-(\\d+)")
var result = regex.search("abc n-0123")
if result:

	print(result.get_string()) # Would print n-0123

[/codeblock]
The results of capturing groups [code]()[/code] can be retrieved by passing the group number to the various methods in [RegExMatch]. Group 0 is the default and will always refer to the entire pattern. In the above example, calling [code]result.get_string(1)[/code] would give you [code]0123[/code].
This version of RegEx also supports named capturing groups, and the names can be used to retrieve the results. If two or more groups have the same name, the name would only refer to the first one with a match.
[codeblock]
var regex = RegEx.new()
regex.compile("d(?<digit>[0-9]+)|x(?<digit>[0-9a-f]+)")
var result = regex.search("the number is x2f")
if result:

	print(result.get_string("digit")) # Would print 2f

[/codeblock]
If you need to process multiple results, [method search_all] generates a list of all non-overlapping results. This can be combined with a [code]for[/code] loop for convenience.
[codeblock]
for result in regex.search_all("d01, d03, d0c, x3f and x42"):

	print(result.get_string("digit"))

# Would print 01 03 0 3f 42
[/codeblock]
[b]Example of splitting a string using a RegEx:[/b]
[codeblock]
var regex = RegEx.new()
regex.compile("\\S+") # Negated whitespace character class.
var results = []
for result in regex.search_all("One  Two \n\tThree"):

	results.push_back(result.get_string())

# The `results` array now contains "One", "Two", "Three".
[/codeblock]
[b]Note:[/b] Godot's regex implementation is based on the [url=https://www.pcre.org/]PCRE2[/url] library. You can view the full pattern reference [url=https://www.pcre.org/current/doc/html/pcre2pattern.html]here[/url].
[b]Tip:[/b] You can use [url=https://regexr.com/]Regexr[/url] to test regular expressions online.
*/
type Instance [1]gdclass.RegEx

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRegEx() Instance
}

/*
Creates and compiles a new [RegEx] object.
*/
func CreateFromString(pattern string) [1]gdclass.RegEx { //gd:RegEx.create_from_string
	self := Instance{}
	return [1]gdclass.RegEx(class(self).CreateFromString(String.New(pattern)))
}

/*
This method resets the state of the object, as if it was freshly created. Namely, it unassigns the regular expression of this object.
*/
func (self Instance) Clear() { //gd:RegEx.clear
	class(self).Clear()
}

/*
Compiles and assign the search pattern to use. Returns [constant OK] if the compilation is successful. If an error is encountered, details are printed to standard output and an error is returned.
*/
func (self Instance) Compile(pattern string) error { //gd:RegEx.compile
	return error(gd.ToError(class(self).Compile(String.New(pattern))))
}

/*
Searches the text for the compiled pattern. Returns a [RegExMatch] container of the first matching result if found, otherwise [code]null[/code].
The region to search within can be specified with [param offset] and [param end]. This is useful when searching for another match in the same [param subject] by calling this method again after a previous success. Note that setting these parameters differs from passing over a shortened string. For example, the start anchor [code]^[/code] is not affected by [param offset], and the character before [param offset] will be checked for the word boundary [code]\b[/code].
*/
func (self Instance) Search(subject string) [1]gdclass.RegExMatch { //gd:RegEx.search
	return [1]gdclass.RegExMatch(class(self).Search(String.New(subject), int64(0), int64(-1)))
}

/*
Searches the text for the compiled pattern. Returns an array of [RegExMatch] containers for each non-overlapping result. If no results were found, an empty array is returned instead.
The region to search within can be specified with [param offset] and [param end]. This is useful when searching for another match in the same [param subject] by calling this method again after a previous success. Note that setting these parameters differs from passing over a shortened string. For example, the start anchor [code]^[/code] is not affected by [param offset], and the character before [param offset] will be checked for the word boundary [code]\b[/code].
*/
func (self Instance) SearchAll(subject string) [][1]gdclass.RegExMatch { //gd:RegEx.search_all
	return [][1]gdclass.RegExMatch(gd.ArrayAs[[][1]gdclass.RegExMatch](gd.InternalArray(class(self).SearchAll(String.New(subject), int64(0), int64(-1)))))
}

/*
Searches the text for the compiled pattern and replaces it with the specified string. Escapes and backreferences such as [code]$1[/code] and [code]$name[/code] are expanded and resolved. By default, only the first instance is replaced, but it can be changed for all instances (global replacement).
The region to search within can be specified with [param offset] and [param end]. This is useful when searching for another match in the same [param subject] by calling this method again after a previous success. Note that setting these parameters differs from passing over a shortened string. For example, the start anchor [code]^[/code] is not affected by [param offset], and the character before [param offset] will be checked for the word boundary [code]\b[/code].
*/
func (self Instance) Sub(subject string, replacement string) string { //gd:RegEx.sub
	return string(class(self).Sub(String.New(subject), String.New(replacement), false, int64(0), int64(-1)).String())
}

/*
Returns whether this object has a valid search pattern assigned.
*/
func (self Instance) IsValid() bool { //gd:RegEx.is_valid
	return bool(class(self).IsValid())
}

/*
Returns the original search pattern that was compiled.
*/
func (self Instance) GetPattern() string { //gd:RegEx.get_pattern
	return string(class(self).GetPattern().String())
}

/*
Returns the number of capturing groups in compiled pattern.
*/
func (self Instance) GetGroupCount() int { //gd:RegEx.get_group_count
	return int(int(class(self).GetGroupCount()))
}

/*
Returns an array of names of named capturing groups in the compiled pattern. They are ordered by appearance.
*/
func (self Instance) GetNames() []string { //gd:RegEx.get_names
	return []string(class(self).GetNames().Strings())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RegEx

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RegEx"))
	casted := Instance{*(*gdclass.RegEx)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Creates and compiles a new [RegEx] object.
*/
//go:nosplit
func (self class) CreateFromString(pattern String.Readable) [1]gdclass.RegEx { //gd:RegEx.create_from_string
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(pattern)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RegEx.Bind_create_from_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.RegEx{gd.PointerWithOwnershipTransferredToGo[gdclass.RegEx](r_ret.Get())}
	frame.Free()
	return ret
}

/*
This method resets the state of the object, as if it was freshly created. Namely, it unassigns the regular expression of this object.
*/
//go:nosplit
func (self class) Clear() { //gd:RegEx.clear
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RegEx.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Compiles and assign the search pattern to use. Returns [constant OK] if the compilation is successful. If an error is encountered, details are printed to standard output and an error is returned.
*/
//go:nosplit
func (self class) Compile(pattern String.Readable) Error.Code { //gd:RegEx.compile
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(pattern)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RegEx.Bind_compile, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Searches the text for the compiled pattern. Returns a [RegExMatch] container of the first matching result if found, otherwise [code]null[/code].
The region to search within can be specified with [param offset] and [param end]. This is useful when searching for another match in the same [param subject] by calling this method again after a previous success. Note that setting these parameters differs from passing over a shortened string. For example, the start anchor [code]^[/code] is not affected by [param offset], and the character before [param offset] will be checked for the word boundary [code]\b[/code].
*/
//go:nosplit
func (self class) Search(subject String.Readable, offset int64, end int64) [1]gdclass.RegExMatch { //gd:RegEx.search
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(subject)))
	callframe.Arg(frame, offset)
	callframe.Arg(frame, end)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RegEx.Bind_search, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.RegExMatch{gd.PointerWithOwnershipTransferredToGo[gdclass.RegExMatch](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Searches the text for the compiled pattern. Returns an array of [RegExMatch] containers for each non-overlapping result. If no results were found, an empty array is returned instead.
The region to search within can be specified with [param offset] and [param end]. This is useful when searching for another match in the same [param subject] by calling this method again after a previous success. Note that setting these parameters differs from passing over a shortened string. For example, the start anchor [code]^[/code] is not affected by [param offset], and the character before [param offset] will be checked for the word boundary [code]\b[/code].
*/
//go:nosplit
func (self class) SearchAll(subject String.Readable, offset int64, end int64) Array.Contains[[1]gdclass.RegExMatch] { //gd:RegEx.search_all
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(subject)))
	callframe.Arg(frame, offset)
	callframe.Arg(frame, end)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RegEx.Bind_search_all, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.RegExMatch]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Searches the text for the compiled pattern and replaces it with the specified string. Escapes and backreferences such as [code]$1[/code] and [code]$name[/code] are expanded and resolved. By default, only the first instance is replaced, but it can be changed for all instances (global replacement).
The region to search within can be specified with [param offset] and [param end]. This is useful when searching for another match in the same [param subject] by calling this method again after a previous success. Note that setting these parameters differs from passing over a shortened string. For example, the start anchor [code]^[/code] is not affected by [param offset], and the character before [param offset] will be checked for the word boundary [code]\b[/code].
*/
//go:nosplit
func (self class) Sub(subject String.Readable, replacement String.Readable, all bool, offset int64, end int64) String.Readable { //gd:RegEx.sub
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(subject)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(replacement)))
	callframe.Arg(frame, all)
	callframe.Arg(frame, offset)
	callframe.Arg(frame, end)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RegEx.Bind_sub, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns whether this object has a valid search pattern assigned.
*/
//go:nosplit
func (self class) IsValid() bool { //gd:RegEx.is_valid
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RegEx.Bind_is_valid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the original search pattern that was compiled.
*/
//go:nosplit
func (self class) GetPattern() String.Readable { //gd:RegEx.get_pattern
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RegEx.Bind_get_pattern, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the number of capturing groups in compiled pattern.
*/
//go:nosplit
func (self class) GetGroupCount() int64 { //gd:RegEx.get_group_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RegEx.Bind_get_group_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an array of names of named capturing groups in the compiled pattern. They are ordered by appearance.
*/
//go:nosplit
func (self class) GetNames() Packed.Strings { //gd:RegEx.get_names
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RegEx.Bind_get_names, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}
func (self class) AsRegEx() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRegEx() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("RegEx", func(ptr gd.Object) any { return [1]gdclass.RegEx{*(*gdclass.RegEx)(unsafe.Pointer(&ptr))} })
}
