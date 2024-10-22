package CodeHighlighter

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/SyntaxHighlighter"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
By adjusting various properties of this resource, you can change the colors of strings, comments, numbers, and other text patterns inside a [TextEdit] control.

*/
type Go [1]classdb.CodeHighlighter

/*
Sets the color for a keyword.
The keyword cannot contain any symbols except '_'.
*/
func (self Go) AddKeywordColor(keyword string, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddKeywordColor(gc.String(keyword), color)
}

/*
Removes the keyword.
*/
func (self Go) RemoveKeywordColor(keyword string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveKeywordColor(gc.String(keyword))
}

/*
Returns [code]true[/code] if the keyword exists, else [code]false[/code].
*/
func (self Go) HasKeywordColor(keyword string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasKeywordColor(gc.String(keyword)))
}

/*
Returns the color for a keyword.
*/
func (self Go) GetKeywordColor(keyword string) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(class(self).GetKeywordColor(gc.String(keyword)))
}

/*
Removes all keywords.
*/
func (self Go) ClearKeywordColors() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearKeywordColors()
}

/*
Sets the color for a member keyword.
The member keyword cannot contain any symbols except '_'.
It will not be highlighted if preceded by a '.'.
*/
func (self Go) AddMemberKeywordColor(member_keyword string, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddMemberKeywordColor(gc.String(member_keyword), color)
}

/*
Removes the member keyword.
*/
func (self Go) RemoveMemberKeywordColor(member_keyword string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveMemberKeywordColor(gc.String(member_keyword))
}

/*
Returns [code]true[/code] if the member keyword exists, else [code]false[/code].
*/
func (self Go) HasMemberKeywordColor(member_keyword string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasMemberKeywordColor(gc.String(member_keyword)))
}

/*
Returns the color for a member keyword.
*/
func (self Go) GetMemberKeywordColor(member_keyword string) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(class(self).GetMemberKeywordColor(gc.String(member_keyword)))
}

/*
Removes all member keywords.
*/
func (self Go) ClearMemberKeywordColors() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearMemberKeywordColors()
}

/*
Adds a color region (such as for comments or strings) from [param start_key] to [param end_key]. Both keys should be symbols, and [param start_key] must not be shared with other delimiters.
If [param line_only] is [code]true[/code] or [param end_key] is an empty [String], the region does not carry over to the next line.
*/
func (self Go) AddColorRegion(start_key string, end_key string, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddColorRegion(gc.String(start_key), gc.String(end_key), color, false)
}

/*
Removes the color region that uses that start key.
*/
func (self Go) RemoveColorRegion(start_key string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveColorRegion(gc.String(start_key))
}

/*
Returns [code]true[/code] if the start key exists, else [code]false[/code].
*/
func (self Go) HasColorRegion(start_key string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasColorRegion(gc.String(start_key)))
}

/*
Removes all color regions.
*/
func (self Go) ClearColorRegions() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearColorRegions()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.CodeHighlighter
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("CodeHighlighter"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) NumberColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetNumberColor())
}

func (self Go) SetNumberColor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNumberColor(value)
}

func (self Go) SymbolColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetSymbolColor())
}

func (self Go) SetSymbolColor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSymbolColor(value)
}

func (self Go) FunctionColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetFunctionColor())
}

func (self Go) SetFunctionColor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFunctionColor(value)
}

func (self Go) MemberVariableColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetMemberVariableColor())
}

func (self Go) SetMemberVariableColor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMemberVariableColor(value)
}

func (self Go) KeywordColors() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Dictionary(class(self).GetKeywordColors(gc))
}

func (self Go) SetKeywordColors(value gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetKeywordColors(value)
}

func (self Go) MemberKeywordColors() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Dictionary(class(self).GetMemberKeywordColors(gc))
}

func (self Go) SetMemberKeywordColors(value gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMemberKeywordColors(value)
}

func (self Go) ColorRegions() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Dictionary(class(self).GetColorRegions(gc))
}

func (self Go) SetColorRegions(value gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetColorRegions(value)
}

/*
Sets the color for a keyword.
The keyword cannot contain any symbols except '_'.
*/
//go:nosplit
func (self class) AddKeywordColor(keyword gd.String, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(keyword))
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_add_keyword_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the keyword.
*/
//go:nosplit
func (self class) RemoveKeywordColor(keyword gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(keyword))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_remove_keyword_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the keyword exists, else [code]false[/code].
*/
//go:nosplit
func (self class) HasKeywordColor(keyword gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(keyword))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_has_keyword_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the color for a keyword.
*/
//go:nosplit
func (self class) GetKeywordColor(keyword gd.String) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(keyword))
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_get_keyword_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetKeywordColors(keywords gd.Dictionary)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(keywords))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_set_keyword_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all keywords.
*/
//go:nosplit
func (self class) ClearKeywordColors()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_clear_keyword_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetKeywordColors(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_get_keyword_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the color for a member keyword.
The member keyword cannot contain any symbols except '_'.
It will not be highlighted if preceded by a '.'.
*/
//go:nosplit
func (self class) AddMemberKeywordColor(member_keyword gd.String, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(member_keyword))
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_add_member_keyword_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the member keyword.
*/
//go:nosplit
func (self class) RemoveMemberKeywordColor(member_keyword gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(member_keyword))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_remove_member_keyword_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the member keyword exists, else [code]false[/code].
*/
//go:nosplit
func (self class) HasMemberKeywordColor(member_keyword gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(member_keyword))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_has_member_keyword_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the color for a member keyword.
*/
//go:nosplit
func (self class) GetMemberKeywordColor(member_keyword gd.String) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(member_keyword))
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_get_member_keyword_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMemberKeywordColors(member_keyword gd.Dictionary)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(member_keyword))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_set_member_keyword_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all member keywords.
*/
//go:nosplit
func (self class) ClearMemberKeywordColors()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_clear_member_keyword_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMemberKeywordColors(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_get_member_keyword_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Adds a color region (such as for comments or strings) from [param start_key] to [param end_key]. Both keys should be symbols, and [param start_key] must not be shared with other delimiters.
If [param line_only] is [code]true[/code] or [param end_key] is an empty [String], the region does not carry over to the next line.
*/
//go:nosplit
func (self class) AddColorRegion(start_key gd.String, end_key gd.String, color gd.Color, line_only bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(start_key))
	callframe.Arg(frame, mmm.Get(end_key))
	callframe.Arg(frame, color)
	callframe.Arg(frame, line_only)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_add_color_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the color region that uses that start key.
*/
//go:nosplit
func (self class) RemoveColorRegion(start_key gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(start_key))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_remove_color_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the start key exists, else [code]false[/code].
*/
//go:nosplit
func (self class) HasColorRegion(start_key gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(start_key))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_has_color_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetColorRegions(color_regions gd.Dictionary)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(color_regions))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_set_color_regions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all color regions.
*/
//go:nosplit
func (self class) ClearColorRegions()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_clear_color_regions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColorRegions(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_get_color_regions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFunctionColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_set_function_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFunctionColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_get_function_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNumberColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_set_number_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNumberColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_get_number_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSymbolColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_set_symbol_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSymbolColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_get_symbol_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMemberVariableColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_set_member_variable_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMemberVariableColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeHighlighter.Bind_get_member_variable_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCodeHighlighter() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCodeHighlighter() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsSyntaxHighlighter() SyntaxHighlighter.GD { return *((*SyntaxHighlighter.GD)(unsafe.Pointer(&self))) }
func (self Go) AsSyntaxHighlighter() SyntaxHighlighter.Go { return *((*SyntaxHighlighter.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsSyntaxHighlighter(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsSyntaxHighlighter(), name)
	}
}
func init() {classdb.Register("CodeHighlighter", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
