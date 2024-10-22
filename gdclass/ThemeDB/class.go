package ThemeDB

import "unsafe"
import "sync"
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
This singleton provides access to static information about [Theme] resources used by the engine and by your projects. You can fetch the default engine theme, as well as your project configured theme.
[ThemeDB] also contains fallback values for theme properties.

*/
var self gdclass.ThemeDB
var once sync.Once
func singleton() {
	gc := gd.Static
	obj := gc.API.Object.GetSingleton(gc, gc.API.Singletons.ThemeDB)
	self = *(*gdclass.ThemeDB)(unsafe.Pointer(&obj))
}

/*
Returns a reference to the default engine [Theme]. This theme resource is responsible for the out-of-the-box look of [Control] nodes and cannot be overridden.
*/
func GetDefaultTheme() gdclass.Theme {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.Theme(class(self).GetDefaultTheme(gc))
}

/*
Returns a reference to the custom project [Theme]. This theme resources allows to override the default engine theme for every control node in the project.
To set the project theme, see [member ProjectSettings.gui/theme/custom].
*/
func GetProjectTheme() gdclass.Theme {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.Theme(class(self).GetProjectTheme(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func GD() class { once.Do(singleton); return self }
type class [1]classdb.ThemeDB
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

func FallbackBaseScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFallbackBaseScale()))
}

func SetFallbackBaseScale(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFallbackBaseScale(gd.Float(value))
}

func FallbackFont() gdclass.Font {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Font(class(self).GetFallbackFont(gc))
}

func SetFallbackFont(value gdclass.Font) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFallbackFont(value)
}

func FallbackFontSize() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetFallbackFontSize()))
}

func SetFallbackFontSize(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFallbackFontSize(gd.Int(value))
}

func FallbackIcon() gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Texture2D(class(self).GetFallbackIcon(gc))
}

func SetFallbackIcon(value gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFallbackIcon(value)
}

func FallbackStylebox() gdclass.StyleBox {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.StyleBox(class(self).GetFallbackStylebox(gc))
}

func SetFallbackStylebox(value gdclass.StyleBox) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFallbackStylebox(value)
}

/*
Returns a reference to the default engine [Theme]. This theme resource is responsible for the out-of-the-box look of [Control] nodes and cannot be overridden.
*/
//go:nosplit
func (self class) GetDefaultTheme(ctx gd.Lifetime) gdclass.Theme {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ThemeDB.Bind_get_default_theme, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Theme
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns a reference to the custom project [Theme]. This theme resources allows to override the default engine theme for every control node in the project.
To set the project theme, see [member ProjectSettings.gui/theme/custom].
*/
//go:nosplit
func (self class) GetProjectTheme(ctx gd.Lifetime) gdclass.Theme {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ThemeDB.Bind_get_project_theme, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Theme
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
func (self class) SetFallbackFont(font gdclass.Font)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(font[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ThemeDB.Bind_set_fallback_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFallbackFont(ctx gd.Lifetime) gdclass.Font {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ThemeDB.Bind_get_fallback_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Font
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
func (self class) SetFallbackIcon(icon gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ThemeDB.Bind_set_fallback_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFallbackIcon(ctx gd.Lifetime) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ThemeDB.Bind_get_fallback_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFallbackStylebox(stylebox gdclass.StyleBox)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(stylebox[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ThemeDB.Bind_set_fallback_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFallbackStylebox(ctx gd.Lifetime) gdclass.StyleBox {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ThemeDB.Bind_get_fallback_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.StyleBox
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func OnFallbackChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("fallback_changed"), gc.Callable(cb), 0)
}


func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("ThemeDB", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
