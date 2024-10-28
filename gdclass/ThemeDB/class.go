package ThemeDB

import "unsafe"
import "sync"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
This singleton provides access to static information about [Theme] resources used by the engine and by your projects. You can fetch the default engine theme, as well as your project configured theme.
[ThemeDB] also contains fallback values for theme properties.

*/
var self gdclass.ThemeDB
var once sync.Once
func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.ThemeDB)
	self = *(*gdclass.ThemeDB)(unsafe.Pointer(&obj))
}

/*
Returns a reference to the default engine [Theme]. This theme resource is responsible for the out-of-the-box look of [Control] nodes and cannot be overridden.
*/
func GetDefaultTheme() gdclass.Theme {
	once.Do(singleton)
	return gdclass.Theme(class(self).GetDefaultTheme())
}

/*
Returns a reference to the custom project [Theme]. This theme resources allows to override the default engine theme for every control node in the project.
To set the project theme, see [member ProjectSettings.gui/theme/custom].
*/
func GetProjectTheme() gdclass.Theme {
	once.Do(singleton)
	return gdclass.Theme(class(self).GetProjectTheme())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func GD() class { once.Do(singleton); return self }
type class [1]classdb.ThemeDB
func (self class) AsObject() gd.Object { return self[0].AsObject() }

func FallbackBaseScale() float64 {
		return float64(float64(class(self).GetFallbackBaseScale()))
}

func SetFallbackBaseScale(value float64) {
	class(self).SetFallbackBaseScale(gd.Float(value))
}

func FallbackFont() gdclass.Font {
		return gdclass.Font(class(self).GetFallbackFont())
}

func SetFallbackFont(value gdclass.Font) {
	class(self).SetFallbackFont(value)
}

func FallbackFontSize() int {
		return int(int(class(self).GetFallbackFontSize()))
}

func SetFallbackFontSize(value int) {
	class(self).SetFallbackFontSize(gd.Int(value))
}

func FallbackIcon() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetFallbackIcon())
}

func SetFallbackIcon(value gdclass.Texture2D) {
	class(self).SetFallbackIcon(value)
}

func FallbackStylebox() gdclass.StyleBox {
		return gdclass.StyleBox(class(self).GetFallbackStylebox())
}

func SetFallbackStylebox(value gdclass.StyleBox) {
	class(self).SetFallbackStylebox(value)
}

/*
Returns a reference to the default engine [Theme]. This theme resource is responsible for the out-of-the-box look of [Control] nodes and cannot be overridden.
*/
//go:nosplit
func (self class) GetDefaultTheme() gdclass.Theme {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_get_default_theme, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Theme{classdb.Theme(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Returns a reference to the custom project [Theme]. This theme resources allows to override the default engine theme for every control node in the project.
To set the project theme, see [member ProjectSettings.gui/theme/custom].
*/
//go:nosplit
func (self class) GetProjectTheme() gdclass.Theme {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_get_project_theme, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Theme{classdb.Theme(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFallbackBaseScale(base_scale gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, base_scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_set_fallback_base_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFallbackBaseScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_get_fallback_base_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFallbackFont(font gdclass.Font)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(font[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_set_fallback_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFallbackFont() gdclass.Font {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_get_fallback_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Font{classdb.Font(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFallbackFontSize(font_size gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, font_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_set_fallback_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFallbackFontSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_get_fallback_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFallbackIcon(icon gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(icon[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_set_fallback_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFallbackIcon() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_get_fallback_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFallbackStylebox(stylebox gdclass.StyleBox)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(stylebox[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_set_fallback_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFallbackStylebox() gdclass.StyleBox {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_get_fallback_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.StyleBox{classdb.StyleBox(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func OnFallbackChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("fallback_changed"), gd.NewCallable(cb), 0)
}


func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("ThemeDB", func(ptr gd.Object) any { return classdb.ThemeDB(ptr) })}
