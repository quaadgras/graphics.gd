package CodeHighlighter

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/SyntaxHighlighter"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
By adjusting various properties of this resource, you can change the colors of strings, comments, numbers, and other text patterns inside a [TextEdit] control.

*/
type Simple [1]classdb.CodeHighlighter
func (self Simple) AddKeywordColor(keyword string, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddKeywordColor(gc.String(keyword), color)
}
func (self Simple) RemoveKeywordColor(keyword string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveKeywordColor(gc.String(keyword))
}
func (self Simple) HasKeywordColor(keyword string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasKeywordColor(gc.String(keyword)))
}
func (self Simple) GetKeywordColor(keyword string) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetKeywordColor(gc.String(keyword)))
}
func (self Simple) SetKeywordColors(keywords gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetKeywordColors(keywords)
}
func (self Simple) ClearKeywordColors() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearKeywordColors()
}
func (self Simple) GetKeywordColors() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).GetKeywordColors(gc))
}
func (self Simple) AddMemberKeywordColor(member_keyword string, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddMemberKeywordColor(gc.String(member_keyword), color)
}
func (self Simple) RemoveMemberKeywordColor(member_keyword string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveMemberKeywordColor(gc.String(member_keyword))
}
func (self Simple) HasMemberKeywordColor(member_keyword string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasMemberKeywordColor(gc.String(member_keyword)))
}
func (self Simple) GetMemberKeywordColor(member_keyword string) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetMemberKeywordColor(gc.String(member_keyword)))
}
func (self Simple) SetMemberKeywordColors(member_keyword gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMemberKeywordColors(member_keyword)
}
func (self Simple) ClearMemberKeywordColors() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearMemberKeywordColors()
}
func (self Simple) GetMemberKeywordColors() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).GetMemberKeywordColors(gc))
}
func (self Simple) AddColorRegion(start_key string, end_key string, color gd.Color, line_only bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddColorRegion(gc.String(start_key), gc.String(end_key), color, line_only)
}
func (self Simple) RemoveColorRegion(start_key string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveColorRegion(gc.String(start_key))
}
func (self Simple) HasColorRegion(start_key string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasColorRegion(gc.String(start_key)))
}
func (self Simple) SetColorRegions(color_regions gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetColorRegions(color_regions)
}
func (self Simple) ClearColorRegions() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearColorRegions()
}
func (self Simple) GetColorRegions() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).GetColorRegions(gc))
}
func (self Simple) SetFunctionColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFunctionColor(color)
}
func (self Simple) GetFunctionColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetFunctionColor())
}
func (self Simple) SetNumberColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNumberColor(color)
}
func (self Simple) GetNumberColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetNumberColor())
}
func (self Simple) SetSymbolColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSymbolColor(color)
}
func (self Simple) GetSymbolColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetSymbolColor())
}
func (self Simple) SetMemberVariableColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMemberVariableColor(color)
}
func (self Simple) GetMemberVariableColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetMemberVariableColor())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.CodeHighlighter
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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

//go:nosplit
func (self class) AsCodeHighlighter() Expert { return self[0].AsCodeHighlighter() }


//go:nosplit
func (self Simple) AsCodeHighlighter() Simple { return self[0].AsCodeHighlighter() }


//go:nosplit
func (self class) AsSyntaxHighlighter() SyntaxHighlighter.Expert { return self[0].AsSyntaxHighlighter() }


//go:nosplit
func (self Simple) AsSyntaxHighlighter() SyntaxHighlighter.Simple { return self[0].AsSyntaxHighlighter() }


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
func init() {classdb.Register("CodeHighlighter", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
