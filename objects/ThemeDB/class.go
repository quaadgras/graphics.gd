package ThemeDB

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This singleton provides access to static information about [Theme] resources used by the engine and by your projects. You can fetch the default engine theme, as well as your project configured theme.
[ThemeDB] also contains fallback values for theme properties.
*/
var self objects.ThemeDB
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.ThemeDB)
	self = *(*objects.ThemeDB)(unsafe.Pointer(&obj))
}

/*
Returns a reference to the default engine [Theme]. This theme resource is responsible for the out-of-the-box look of [Control] nodes and cannot be overridden.
*/
func GetDefaultTheme() objects.Theme {
	once.Do(singleton)
	return objects.Theme(class(self).GetDefaultTheme())
}

/*
Returns a reference to the custom project [Theme]. This theme resources allows to override the default engine theme for every control node in the project.
To set the project theme, see [member ProjectSettings.gui/theme/custom].
*/
func GetProjectTheme() objects.Theme {
	once.Do(singleton)
	return objects.Theme(class(self).GetProjectTheme())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]classdb.ThemeDB

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

func FallbackBaseScale() Float.X {
	return Float.X(Float.X(class(self).GetFallbackBaseScale()))
}

func SetFallbackBaseScale(value Float.X) {
	class(self).SetFallbackBaseScale(gd.Float(value))
}

func FallbackFont() objects.Font {
	return objects.Font(class(self).GetFallbackFont())
}

func SetFallbackFont(value objects.Font) {
	class(self).SetFallbackFont(value)
}

func FallbackFontSize() int {
	return int(int(class(self).GetFallbackFontSize()))
}

func SetFallbackFontSize(value int) {
	class(self).SetFallbackFontSize(gd.Int(value))
}

func FallbackIcon() objects.Texture2D {
	return objects.Texture2D(class(self).GetFallbackIcon())
}

func SetFallbackIcon(value objects.Texture2D) {
	class(self).SetFallbackIcon(value)
}

func FallbackStylebox() objects.StyleBox {
	return objects.StyleBox(class(self).GetFallbackStylebox())
}

func SetFallbackStylebox(value objects.StyleBox) {
	class(self).SetFallbackStylebox(value)
}

/*
Returns a reference to the default engine [Theme]. This theme resource is responsible for the out-of-the-box look of [Control] nodes and cannot be overridden.
*/
//go:nosplit
func (self class) GetDefaultTheme() objects.Theme {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_get_default_theme, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Theme{gd.PointerWithOwnershipTransferredToGo[classdb.Theme](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns a reference to the custom project [Theme]. This theme resources allows to override the default engine theme for every control node in the project.
To set the project theme, see [member ProjectSettings.gui/theme/custom].
*/
//go:nosplit
func (self class) GetProjectTheme() objects.Theme {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_get_project_theme, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Theme{gd.PointerWithOwnershipTransferredToGo[classdb.Theme](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFallbackBaseScale(base_scale gd.Float) {
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
func (self class) SetFallbackFont(font objects.Font) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(font[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_set_fallback_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFallbackFont() objects.Font {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_get_fallback_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Font{gd.PointerWithOwnershipTransferredToGo[classdb.Font](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFallbackFontSize(font_size gd.Int) {
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
func (self class) SetFallbackIcon(icon objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_set_fallback_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFallbackIcon() objects.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_get_fallback_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{gd.PointerWithOwnershipTransferredToGo[classdb.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFallbackStylebox(stylebox objects.StyleBox) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(stylebox[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_set_fallback_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFallbackStylebox() objects.StyleBox {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_get_fallback_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.StyleBox{gd.PointerWithOwnershipTransferredToGo[classdb.StyleBox](r_ret.Get())}
	frame.Free()
	return ret
}
func OnFallbackChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("fallback_changed"), gd.NewCallable(cb), 0)
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("ThemeDB", func(ptr gd.Object) any { return [1]classdb.ThemeDB{*(*classdb.ThemeDB)(unsafe.Pointer(&ptr))} })
}
