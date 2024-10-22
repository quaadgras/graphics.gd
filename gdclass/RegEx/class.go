package RegEx

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
type Go [1]classdb.RegEx

/*
Creates and compiles a new [RegEx] object.
*/
func (self Go) CreateFromString(pattern string) gdclass.RegEx {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.RegEx(class(self).CreateFromString(gc, gc.String(pattern)))
}

/*
This method resets the state of the object, as if it was freshly created. Namely, it unassigns the regular expression of this object.
*/
func (self Go) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Clear()
}

/*
Compiles and assign the search pattern to use. Returns [constant OK] if the compilation is successful. If an error is encountered, details are printed to standard output and an error is returned.
*/
func (self Go) Compile(pattern string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Compile(gc.String(pattern)))
}

/*
Searches the text for the compiled pattern. Returns a [RegExMatch] container of the first matching result if found, otherwise [code]null[/code].
The region to search within can be specified with [param offset] and [param end]. This is useful when searching for another match in the same [param subject] by calling this method again after a previous success. Note that setting these parameters differs from passing over a shortened string. For example, the start anchor [code]^[/code] is not affected by [param offset], and the character before [param offset] will be checked for the word boundary [code]\b[/code].
*/
func (self Go) Search(subject string) gdclass.RegExMatch {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.RegExMatch(class(self).Search(gc, gc.String(subject), gd.Int(0), gd.Int(-1)))
}

/*
Searches the text for the compiled pattern. Returns an array of [RegExMatch] containers for each non-overlapping result. If no results were found, an empty array is returned instead.
The region to search within can be specified with [param offset] and [param end]. This is useful when searching for another match in the same [param subject] by calling this method again after a previous success. Note that setting these parameters differs from passing over a shortened string. For example, the start anchor [code]^[/code] is not affected by [param offset], and the character before [param offset] will be checked for the word boundary [code]\b[/code].
*/
func (self Go) SearchAll(subject string) gd.ArrayOf[gdclass.RegExMatch] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gdclass.RegExMatch](class(self).SearchAll(gc, gc.String(subject), gd.Int(0), gd.Int(-1)))
}

/*
Searches the text for the compiled pattern and replaces it with the specified string. Escapes and backreferences such as [code]$1[/code] and [code]$name[/code] are expanded and resolved. By default, only the first instance is replaced, but it can be changed for all instances (global replacement).
The region to search within can be specified with [param offset] and [param end]. This is useful when searching for another match in the same [param subject] by calling this method again after a previous success. Note that setting these parameters differs from passing over a shortened string. For example, the start anchor [code]^[/code] is not affected by [param offset], and the character before [param offset] will be checked for the word boundary [code]\b[/code].
*/
func (self Go) Sub(subject string, replacement string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).Sub(gc, gc.String(subject), gc.String(replacement), false, gd.Int(0), gd.Int(-1)).String())
}

/*
Returns whether this object has a valid search pattern assigned.
*/
func (self Go) IsValid() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsValid())
}

/*
Returns the original search pattern that was compiled.
*/
func (self Go) GetPattern() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetPattern(gc).String())
}

/*
Returns the number of capturing groups in compiled pattern.
*/
func (self Go) GetGroupCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetGroupCount()))
}

/*
Returns an array of names of named capturing groups in the compiled pattern. They are ordered by appearance.
*/
func (self Go) GetNames() []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetNames(gc).Strings(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RegEx
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RegEx"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Creates and compiles a new [RegEx] object.
*/
//go:nosplit
func (self class) CreateFromString(ctx gd.Lifetime, pattern gd.String) gdclass.RegEx {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(pattern))
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.RegEx.Bind_create_from_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.RegEx
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
This method resets the state of the object, as if it was freshly created. Namely, it unassigns the regular expression of this object.
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RegEx.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Compiles and assign the search pattern to use. Returns [constant OK] if the compilation is successful. If an error is encountered, details are printed to standard output and an error is returned.
*/
//go:nosplit
func (self class) Compile(pattern gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(pattern))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RegEx.Bind_compile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Searches the text for the compiled pattern. Returns a [RegExMatch] container of the first matching result if found, otherwise [code]null[/code].
The region to search within can be specified with [param offset] and [param end]. This is useful when searching for another match in the same [param subject] by calling this method again after a previous success. Note that setting these parameters differs from passing over a shortened string. For example, the start anchor [code]^[/code] is not affected by [param offset], and the character before [param offset] will be checked for the word boundary [code]\b[/code].
*/
//go:nosplit
func (self class) Search(ctx gd.Lifetime, subject gd.String, offset gd.Int, end gd.Int) gdclass.RegExMatch {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(subject))
	callframe.Arg(frame, offset)
	callframe.Arg(frame, end)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RegEx.Bind_search, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.RegExMatch
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Searches the text for the compiled pattern. Returns an array of [RegExMatch] containers for each non-overlapping result. If no results were found, an empty array is returned instead.
The region to search within can be specified with [param offset] and [param end]. This is useful when searching for another match in the same [param subject] by calling this method again after a previous success. Note that setting these parameters differs from passing over a shortened string. For example, the start anchor [code]^[/code] is not affected by [param offset], and the character before [param offset] will be checked for the word boundary [code]\b[/code].
*/
//go:nosplit
func (self class) SearchAll(ctx gd.Lifetime, subject gd.String, offset gd.Int, end gd.Int) gd.ArrayOf[gdclass.RegExMatch] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(subject))
	callframe.Arg(frame, offset)
	callframe.Arg(frame, end)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RegEx.Bind_search_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.RegExMatch](ret)
}
/*
Searches the text for the compiled pattern and replaces it with the specified string. Escapes and backreferences such as [code]$1[/code] and [code]$name[/code] are expanded and resolved. By default, only the first instance is replaced, but it can be changed for all instances (global replacement).
The region to search within can be specified with [param offset] and [param end]. This is useful when searching for another match in the same [param subject] by calling this method again after a previous success. Note that setting these parameters differs from passing over a shortened string. For example, the start anchor [code]^[/code] is not affected by [param offset], and the character before [param offset] will be checked for the word boundary [code]\b[/code].
*/
//go:nosplit
func (self class) Sub(ctx gd.Lifetime, subject gd.String, replacement gd.String, all bool, offset gd.Int, end gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(subject))
	callframe.Arg(frame, mmm.Get(replacement))
	callframe.Arg(frame, all)
	callframe.Arg(frame, offset)
	callframe.Arg(frame, end)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RegEx.Bind_sub, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns whether this object has a valid search pattern assigned.
*/
//go:nosplit
func (self class) IsValid() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RegEx.Bind_is_valid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the original search pattern that was compiled.
*/
//go:nosplit
func (self class) GetPattern(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RegEx.Bind_get_pattern, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the number of capturing groups in compiled pattern.
*/
//go:nosplit
func (self class) GetGroupCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RegEx.Bind_get_group_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an array of names of named capturing groups in the compiled pattern. They are ordered by appearance.
*/
//go:nosplit
func (self class) GetNames(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RegEx.Bind_get_names, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsRegEx() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRegEx() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("RegEx", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
