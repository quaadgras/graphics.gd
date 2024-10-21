package ThemeDB

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This singleton provides access to static information about [Theme] resources used by the engine and by your projects. You can fetch the default engine theme, as well as your project configured theme.
[ThemeDB] also contains fallback values for theme properties.

*/
type Simple [1]classdb.ThemeDB
func (self Simple) GetDefaultTheme() [1]classdb.Theme {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Theme(Expert(self).GetDefaultTheme(gc))
}
func (self Simple) GetProjectTheme() [1]classdb.Theme {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Theme(Expert(self).GetProjectTheme(gc))
}
func (self Simple) SetFallbackBaseScale(base_scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFallbackBaseScale(gd.Float(base_scale))
}
func (self Simple) GetFallbackBaseScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFallbackBaseScale()))
}
func (self Simple) SetFallbackFont(font [1]classdb.Font) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFallbackFont(font)
}
func (self Simple) GetFallbackFont() [1]classdb.Font {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Font(Expert(self).GetFallbackFont(gc))
}
func (self Simple) SetFallbackFontSize(font_size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFallbackFontSize(gd.Int(font_size))
}
func (self Simple) GetFallbackFontSize() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetFallbackFontSize()))
}
func (self Simple) SetFallbackIcon(icon [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFallbackIcon(icon)
}
func (self Simple) GetFallbackIcon() [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetFallbackIcon(gc))
}
func (self Simple) SetFallbackStylebox(stylebox [1]classdb.StyleBox) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFallbackStylebox(stylebox)
}
func (self Simple) GetFallbackStylebox() [1]classdb.StyleBox {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.StyleBox(Expert(self).GetFallbackStylebox(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ThemeDB
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns a reference to the default engine [Theme]. This theme resource is responsible for the out-of-the-box look of [Control] nodes and cannot be overridden.
*/
//go:nosplit
func (self class) GetDefaultTheme(ctx gd.Lifetime) object.Theme {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ThemeDB.Bind_get_default_theme, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Theme
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns a reference to the custom project [Theme]. This theme resources allows to override the default engine theme for every control node in the project.
To set the project theme, see [member ProjectSettings.gui/theme/custom].
*/
//go:nosplit
func (self class) GetProjectTheme(ctx gd.Lifetime) object.Theme {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ThemeDB.Bind_get_project_theme, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Theme
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFallbackBaseScale(base_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, base_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ThemeDB.Bind_set_fallback_base_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFallbackBaseScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ThemeDB.Bind_get_fallback_base_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFallbackFont(font object.Font)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(font[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ThemeDB.Bind_set_fallback_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFallbackFont(ctx gd.Lifetime) object.Font {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ThemeDB.Bind_get_fallback_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Font
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFallbackFontSize(font_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ThemeDB.Bind_set_fallback_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFallbackFontSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ThemeDB.Bind_get_fallback_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFallbackIcon(icon object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ThemeDB.Bind_set_fallback_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFallbackIcon(ctx gd.Lifetime) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ThemeDB.Bind_get_fallback_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFallbackStylebox(stylebox object.StyleBox)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(stylebox[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ThemeDB.Bind_set_fallback_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFallbackStylebox(ctx gd.Lifetime) object.StyleBox {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ThemeDB.Bind_get_fallback_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.StyleBox
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsThemeDB() Expert { return self[0].AsThemeDB() }


//go:nosplit
func (self Simple) AsThemeDB() Simple { return self[0].AsThemeDB() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("ThemeDB", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
