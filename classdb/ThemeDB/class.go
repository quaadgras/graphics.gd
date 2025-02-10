// Package ThemeDB provides methods for working with ThemeDB object instances.
package ThemeDB

import "unsafe"
import "sync"
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
This singleton provides access to static information about [Theme] resources used by the engine and by your projects. You can fetch the default engine theme, as well as your project configured theme.
[ThemeDB] also contains fallback values for theme properties.
*/
var self [1]gdclass.ThemeDB
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.ThemeDB)
	self = *(*[1]gdclass.ThemeDB)(unsafe.Pointer(&obj))
}

/*
Returns a reference to the default engine [Theme]. This theme resource is responsible for the out-of-the-box look of [Control] nodes and cannot be overridden.
*/
func GetDefaultTheme() [1]gdclass.Theme { //gd:ThemeDB.get_default_theme
	once.Do(singleton)
	return [1]gdclass.Theme(class(self).GetDefaultTheme())
}

/*
Returns a reference to the custom project [Theme]. This theme resources allows to override the default engine theme for every control node in the project.
To set the project theme, see [member ProjectSettings.gui/theme/custom].
*/
func GetProjectTheme() [1]gdclass.Theme { //gd:ThemeDB.get_project_theme
	once.Do(singleton)
	return [1]gdclass.Theme(class(self).GetProjectTheme())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.ThemeDB

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

func FallbackBaseScale() Float.X {
	once.Do(singleton)
	return Float.X(Float.X(class(self).GetFallbackBaseScale()))
}

func SetFallbackBaseScale(value Float.X) {
	once.Do(singleton)
	class(self).SetFallbackBaseScale(float64(value))
}

func FallbackFont() [1]gdclass.Font {
	once.Do(singleton)
	return [1]gdclass.Font(class(self).GetFallbackFont())
}

func SetFallbackFont(value [1]gdclass.Font) {
	once.Do(singleton)
	class(self).SetFallbackFont(value)
}

func FallbackFontSize() int {
	once.Do(singleton)
	return int(int(class(self).GetFallbackFontSize()))
}

func SetFallbackFontSize(value int) {
	once.Do(singleton)
	class(self).SetFallbackFontSize(int64(value))
}

func FallbackIcon() [1]gdclass.Texture2D {
	once.Do(singleton)
	return [1]gdclass.Texture2D(class(self).GetFallbackIcon())
}

func SetFallbackIcon(value [1]gdclass.Texture2D) {
	once.Do(singleton)
	class(self).SetFallbackIcon(value)
}

func FallbackStylebox() [1]gdclass.StyleBox {
	once.Do(singleton)
	return [1]gdclass.StyleBox(class(self).GetFallbackStylebox())
}

func SetFallbackStylebox(value [1]gdclass.StyleBox) {
	once.Do(singleton)
	class(self).SetFallbackStylebox(value)
}

/*
Returns a reference to the default engine [Theme]. This theme resource is responsible for the out-of-the-box look of [Control] nodes and cannot be overridden.
*/
//go:nosplit
func (self class) GetDefaultTheme() [1]gdclass.Theme { //gd:ThemeDB.get_default_theme
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_get_default_theme, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Theme{gd.PointerWithOwnershipTransferredToGo[gdclass.Theme](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns a reference to the custom project [Theme]. This theme resources allows to override the default engine theme for every control node in the project.
To set the project theme, see [member ProjectSettings.gui/theme/custom].
*/
//go:nosplit
func (self class) GetProjectTheme() [1]gdclass.Theme { //gd:ThemeDB.get_project_theme
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_get_project_theme, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Theme{gd.PointerWithOwnershipTransferredToGo[gdclass.Theme](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFallbackBaseScale(base_scale float64) { //gd:ThemeDB.set_fallback_base_scale
	var frame = callframe.New()
	callframe.Arg(frame, base_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_set_fallback_base_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFallbackBaseScale() float64 { //gd:ThemeDB.get_fallback_base_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_get_fallback_base_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFallbackFont(font [1]gdclass.Font) { //gd:ThemeDB.set_fallback_font
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(font[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_set_fallback_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFallbackFont() [1]gdclass.Font { //gd:ThemeDB.get_fallback_font
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_get_fallback_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Font{gd.PointerWithOwnershipTransferredToGo[gdclass.Font](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFallbackFontSize(font_size int64) { //gd:ThemeDB.set_fallback_font_size
	var frame = callframe.New()
	callframe.Arg(frame, font_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_set_fallback_font_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFallbackFontSize() int64 { //gd:ThemeDB.get_fallback_font_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_get_fallback_font_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFallbackIcon(icon [1]gdclass.Texture2D) { //gd:ThemeDB.set_fallback_icon
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_set_fallback_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFallbackIcon() [1]gdclass.Texture2D { //gd:ThemeDB.get_fallback_icon
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_get_fallback_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFallbackStylebox(stylebox [1]gdclass.StyleBox) { //gd:ThemeDB.set_fallback_stylebox
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(stylebox[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_set_fallback_stylebox, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFallbackStylebox() [1]gdclass.StyleBox { //gd:ThemeDB.get_fallback_stylebox
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ThemeDB.Bind_get_fallback_stylebox, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.StyleBox{gd.PointerWithOwnershipTransferredToGo[gdclass.StyleBox](r_ret.Get())}
	frame.Free()
	return ret
}
func OnFallbackChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("fallback_changed"), gd.NewCallable(cb), 0)
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("ThemeDB", func(ptr gd.Object) any { return [1]gdclass.ThemeDB{*(*gdclass.ThemeDB)(unsafe.Pointer(&ptr))} })
}
