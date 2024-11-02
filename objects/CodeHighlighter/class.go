package CodeHighlighter

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/SyntaxHighlighter"
import "grow.graphics/gd/objects/Resource"
import "grow.graphics/gd/variant/Color"
import "grow.graphics/gd/variant/Dictionary"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
By adjusting various properties of this resource, you can change the colors of strings, comments, numbers, and other text patterns inside a [TextEdit] control.
*/
type Instance [1]classdb.CodeHighlighter

/*
Sets the color for a keyword.
The keyword cannot contain any symbols except '_'.
*/
func (self Instance) AddKeywordColor(keyword string, color Color.RGBA) {
	class(self).AddKeywordColor(gd.NewString(keyword), gd.Color(color))
}

/*
Removes the keyword.
*/
func (self Instance) RemoveKeywordColor(keyword string) {
	class(self).RemoveKeywordColor(gd.NewString(keyword))
}

/*
Returns [code]true[/code] if the keyword exists, else [code]false[/code].
*/
func (self Instance) HasKeywordColor(keyword string) bool {
	return bool(class(self).HasKeywordColor(gd.NewString(keyword)))
}

/*
Returns the color for a keyword.
*/
func (self Instance) GetKeywordColor(keyword string) Color.RGBA {
	return Color.RGBA(class(self).GetKeywordColor(gd.NewString(keyword)))
}

/*
Removes all keywords.
*/
func (self Instance) ClearKeywordColors() {
	class(self).ClearKeywordColors()
}

/*
Sets the color for a member keyword.
The member keyword cannot contain any symbols except '_'.
It will not be highlighted if preceded by a '.'.
*/
func (self Instance) AddMemberKeywordColor(member_keyword string, color Color.RGBA) {
	class(self).AddMemberKeywordColor(gd.NewString(member_keyword), gd.Color(color))
}

/*
Removes the member keyword.
*/
func (self Instance) RemoveMemberKeywordColor(member_keyword string) {
	class(self).RemoveMemberKeywordColor(gd.NewString(member_keyword))
}

/*
Returns [code]true[/code] if the member keyword exists, else [code]false[/code].
*/
func (self Instance) HasMemberKeywordColor(member_keyword string) bool {
	return bool(class(self).HasMemberKeywordColor(gd.NewString(member_keyword)))
}

/*
Returns the color for a member keyword.
*/
func (self Instance) GetMemberKeywordColor(member_keyword string) Color.RGBA {
	return Color.RGBA(class(self).GetMemberKeywordColor(gd.NewString(member_keyword)))
}

/*
Removes all member keywords.
*/
func (self Instance) ClearMemberKeywordColors() {
	class(self).ClearMemberKeywordColors()
}

/*
Adds a color region (such as for comments or strings) from [param start_key] to [param end_key]. Both keys should be symbols, and [param start_key] must not be shared with other delimiters.
If [param line_only] is [code]true[/code] or [param end_key] is an empty [String], the region does not carry over to the next line.
*/
func (self Instance) AddColorRegion(start_key string, end_key string, color Color.RGBA) {
	class(self).AddColorRegion(gd.NewString(start_key), gd.NewString(end_key), gd.Color(color), false)
}

/*
Removes the color region that uses that start key.
*/
func (self Instance) RemoveColorRegion(start_key string) {
	class(self).RemoveColorRegion(gd.NewString(start_key))
}

/*
Returns [code]true[/code] if the start key exists, else [code]false[/code].
*/
func (self Instance) HasColorRegion(start_key string) bool {
	return bool(class(self).HasColorRegion(gd.NewString(start_key)))
}

/*
Removes all color regions.
*/
func (self Instance) ClearColorRegions() {
	class(self).ClearColorRegions()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.CodeHighlighter

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CodeHighlighter"))
	return Instance{classdb.CodeHighlighter(object)}
}

func (self Instance) NumberColor() Color.RGBA {
	return Color.RGBA(class(self).GetNumberColor())
}

func (self Instance) SetNumberColor(value Color.RGBA) {
	class(self).SetNumberColor(gd.Color(value))
}

func (self Instance) SymbolColor() Color.RGBA {
	return Color.RGBA(class(self).GetSymbolColor())
}

func (self Instance) SetSymbolColor(value Color.RGBA) {
	class(self).SetSymbolColor(gd.Color(value))
}

func (self Instance) FunctionColor() Color.RGBA {
	return Color.RGBA(class(self).GetFunctionColor())
}

func (self Instance) SetFunctionColor(value Color.RGBA) {
	class(self).SetFunctionColor(gd.Color(value))
}

func (self Instance) MemberVariableColor() Color.RGBA {
	return Color.RGBA(class(self).GetMemberVariableColor())
}

func (self Instance) SetMemberVariableColor(value Color.RGBA) {
	class(self).SetMemberVariableColor(gd.Color(value))
}

func (self Instance) KeywordColors() Dictionary.Any {
	return Dictionary.Any(class(self).GetKeywordColors())
}

func (self Instance) SetKeywordColors(value Dictionary.Any) {
	class(self).SetKeywordColors(value)
}

func (self Instance) MemberKeywordColors() Dictionary.Any {
	return Dictionary.Any(class(self).GetMemberKeywordColors())
}

func (self Instance) SetMemberKeywordColors(value Dictionary.Any) {
	class(self).SetMemberKeywordColors(value)
}

func (self Instance) ColorRegions() Dictionary.Any {
	return Dictionary.Any(class(self).GetColorRegions())
}

func (self Instance) SetColorRegions(value Dictionary.Any) {
	class(self).SetColorRegions(value)
}

/*
Sets the color for a keyword.
The keyword cannot contain any symbols except '_'.
*/
//go:nosplit
func (self class) AddKeywordColor(keyword gd.String, color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(keyword))
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_add_keyword_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the keyword.
*/
//go:nosplit
func (self class) RemoveKeywordColor(keyword gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(keyword))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_remove_keyword_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the keyword exists, else [code]false[/code].
*/
//go:nosplit
func (self class) HasKeywordColor(keyword gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(keyword))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_has_keyword_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the color for a keyword.
*/
//go:nosplit
func (self class) GetKeywordColor(keyword gd.String) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(keyword))
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_get_keyword_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetKeywordColors(keywords gd.Dictionary) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(keywords))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_set_keyword_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes all keywords.
*/
//go:nosplit
func (self class) ClearKeywordColors() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_clear_keyword_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetKeywordColors() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_get_keyword_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the color for a member keyword.
The member keyword cannot contain any symbols except '_'.
It will not be highlighted if preceded by a '.'.
*/
//go:nosplit
func (self class) AddMemberKeywordColor(member_keyword gd.String, color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(member_keyword))
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_add_member_keyword_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the member keyword.
*/
//go:nosplit
func (self class) RemoveMemberKeywordColor(member_keyword gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(member_keyword))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_remove_member_keyword_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the member keyword exists, else [code]false[/code].
*/
//go:nosplit
func (self class) HasMemberKeywordColor(member_keyword gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(member_keyword))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_has_member_keyword_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the color for a member keyword.
*/
//go:nosplit
func (self class) GetMemberKeywordColor(member_keyword gd.String) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(member_keyword))
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_get_member_keyword_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMemberKeywordColors(member_keyword gd.Dictionary) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(member_keyword))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_set_member_keyword_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes all member keywords.
*/
//go:nosplit
func (self class) ClearMemberKeywordColors() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_clear_member_keyword_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMemberKeywordColors() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_get_member_keyword_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Adds a color region (such as for comments or strings) from [param start_key] to [param end_key]. Both keys should be symbols, and [param start_key] must not be shared with other delimiters.
If [param line_only] is [code]true[/code] or [param end_key] is an empty [String], the region does not carry over to the next line.
*/
//go:nosplit
func (self class) AddColorRegion(start_key gd.String, end_key gd.String, color gd.Color, line_only bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(start_key))
	callframe.Arg(frame, pointers.Get(end_key))
	callframe.Arg(frame, color)
	callframe.Arg(frame, line_only)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_add_color_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the color region that uses that start key.
*/
//go:nosplit
func (self class) RemoveColorRegion(start_key gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(start_key))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_remove_color_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the start key exists, else [code]false[/code].
*/
//go:nosplit
func (self class) HasColorRegion(start_key gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(start_key))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_has_color_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColorRegions(color_regions gd.Dictionary) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(color_regions))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_set_color_regions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes all color regions.
*/
//go:nosplit
func (self class) ClearColorRegions() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_clear_color_regions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetColorRegions() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_get_color_regions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFunctionColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_set_function_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFunctionColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_get_function_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNumberColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_set_number_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetNumberColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_get_number_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSymbolColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_set_symbol_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSymbolColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_get_symbol_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMemberVariableColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_set_member_variable_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMemberVariableColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeHighlighter.Bind_get_member_variable_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCodeHighlighter() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCodeHighlighter() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsSyntaxHighlighter() SyntaxHighlighter.Advanced {
	return *((*SyntaxHighlighter.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSyntaxHighlighter() SyntaxHighlighter.Instance {
	return *((*SyntaxHighlighter.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsSyntaxHighlighter(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsSyntaxHighlighter(), name)
	}
}
func init() {
	classdb.Register("CodeHighlighter", func(ptr gd.Object) any { return classdb.CodeHighlighter(ptr) })
}
