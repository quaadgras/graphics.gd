package Window

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Viewport"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A node that creates a window. The window can either be a native system window or embedded inside another [Window] (see [member Viewport.gui_embed_subwindows]).
At runtime, [Window]s will not close automatically when requested. You need to handle it manually using the [signal close_requested] signal (this applies both to pressing the close button and clicking outside of a popup).
	// Window methods that can be overridden by a [Class] that extends it.
	type Window interface {
		//Virtual method to be implemented by the user. Overrides the value returned by [method get_contents_minimum_size].
		GetContentsMinimumSize() gd.Vector2
	}

*/
type Simple [1]classdb.Window
func (self Simple) SetTitle(title string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTitle(gc.String(title))
}
func (self Simple) GetTitle() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetTitle(gc).String())
}
func (self Simple) GetWindowId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetWindowId()))
}
func (self Simple) SetInitialPosition(initial_position classdb.WindowWindowInitialPosition) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInitialPosition(initial_position)
}
func (self Simple) GetInitialPosition() classdb.WindowWindowInitialPosition {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.WindowWindowInitialPosition(Expert(self).GetInitialPosition())
}
func (self Simple) SetCurrentScreen(index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCurrentScreen(gd.Int(index))
}
func (self Simple) GetCurrentScreen() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCurrentScreen()))
}
func (self Simple) SetPosition(position gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPosition(position)
}
func (self Simple) GetPosition() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetPosition())
}
func (self Simple) MoveToCenter() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MoveToCenter()
}
func (self Simple) SetSize(size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSize(size)
}
func (self Simple) GetSize() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetSize())
}
func (self Simple) ResetSize() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ResetSize()
}
func (self Simple) GetPositionWithDecorations() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetPositionWithDecorations())
}
func (self Simple) GetSizeWithDecorations() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetSizeWithDecorations())
}
func (self Simple) SetMaxSize(max_size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaxSize(max_size)
}
func (self Simple) GetMaxSize() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetMaxSize())
}
func (self Simple) SetMinSize(min_size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMinSize(min_size)
}
func (self Simple) GetMinSize() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetMinSize())
}
func (self Simple) SetMode(mode classdb.WindowMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMode(mode)
}
func (self Simple) GetMode() classdb.WindowMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.WindowMode(Expert(self).GetMode())
}
func (self Simple) SetFlag(flag classdb.WindowFlags, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFlag(flag, enabled)
}
func (self Simple) GetFlag(flag classdb.WindowFlags) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetFlag(flag))
}
func (self Simple) IsMaximizeAllowed() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsMaximizeAllowed())
}
func (self Simple) RequestAttention() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RequestAttention()
}
func (self Simple) MoveToForeground() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MoveToForeground()
}
func (self Simple) SetVisible(visible bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVisible(visible)
}
func (self Simple) IsVisible() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsVisible())
}
func (self Simple) Hide() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Hide()
}
func (self Simple) Show() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Show()
}
func (self Simple) SetTransient(transient bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTransient(transient)
}
func (self Simple) IsTransient() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsTransient())
}
func (self Simple) SetTransientToFocused(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTransientToFocused(enable)
}
func (self Simple) IsTransientToFocused() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsTransientToFocused())
}
func (self Simple) SetExclusive(exclusive bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetExclusive(exclusive)
}
func (self Simple) IsExclusive() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsExclusive())
}
func (self Simple) SetUnparentWhenInvisible(unparent bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUnparentWhenInvisible(unparent)
}
func (self Simple) CanDraw() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).CanDraw())
}
func (self Simple) HasFocus() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasFocus())
}
func (self Simple) GrabFocus() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GrabFocus()
}
func (self Simple) SetImeActive(active bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetImeActive(active)
}
func (self Simple) SetImePosition(position gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetImePosition(position)
}
func (self Simple) IsEmbedded() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsEmbedded())
}
func (self Simple) GetContentsMinimumSize() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetContentsMinimumSize())
}
func (self Simple) SetForceNative(force_native bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetForceNative(force_native)
}
func (self Simple) GetForceNative() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetForceNative())
}
func (self Simple) SetContentScaleSize(size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetContentScaleSize(size)
}
func (self Simple) GetContentScaleSize() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetContentScaleSize())
}
func (self Simple) SetContentScaleMode(mode classdb.WindowContentScaleMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetContentScaleMode(mode)
}
func (self Simple) GetContentScaleMode() classdb.WindowContentScaleMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.WindowContentScaleMode(Expert(self).GetContentScaleMode())
}
func (self Simple) SetContentScaleAspect(aspect classdb.WindowContentScaleAspect) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetContentScaleAspect(aspect)
}
func (self Simple) GetContentScaleAspect() classdb.WindowContentScaleAspect {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.WindowContentScaleAspect(Expert(self).GetContentScaleAspect())
}
func (self Simple) SetContentScaleStretch(stretch classdb.WindowContentScaleStretch) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetContentScaleStretch(stretch)
}
func (self Simple) GetContentScaleStretch() classdb.WindowContentScaleStretch {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.WindowContentScaleStretch(Expert(self).GetContentScaleStretch())
}
func (self Simple) SetKeepTitleVisible(title_visible bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetKeepTitleVisible(title_visible)
}
func (self Simple) GetKeepTitleVisible() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetKeepTitleVisible())
}
func (self Simple) SetContentScaleFactor(factor float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetContentScaleFactor(gd.Float(factor))
}
func (self Simple) GetContentScaleFactor() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetContentScaleFactor()))
}
func (self Simple) SetUseFontOversampling(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseFontOversampling(enable)
}
func (self Simple) IsUsingFontOversampling() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsUsingFontOversampling())
}
func (self Simple) SetMousePassthroughPolygon(polygon gd.PackedVector2Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMousePassthroughPolygon(polygon)
}
func (self Simple) GetMousePassthroughPolygon() gd.PackedVector2Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector2Array(Expert(self).GetMousePassthroughPolygon(gc))
}
func (self Simple) SetWrapControls(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetWrapControls(enable)
}
func (self Simple) IsWrappingControls() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsWrappingControls())
}
func (self Simple) ChildControlsChanged() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ChildControlsChanged()
}
func (self Simple) SetTheme(theme [1]classdb.Theme) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTheme(theme)
}
func (self Simple) GetTheme() [1]classdb.Theme {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Theme(Expert(self).GetTheme(gc))
}
func (self Simple) SetThemeTypeVariation(theme_type string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetThemeTypeVariation(gc.StringName(theme_type))
}
func (self Simple) GetThemeTypeVariation() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetThemeTypeVariation(gc).String())
}
func (self Simple) BeginBulkThemeOverride() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BeginBulkThemeOverride()
}
func (self Simple) EndBulkThemeOverride() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EndBulkThemeOverride()
}
func (self Simple) AddThemeIconOverride(name string, texture [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddThemeIconOverride(gc.StringName(name), texture)
}
func (self Simple) AddThemeStyleboxOverride(name string, stylebox [1]classdb.StyleBox) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddThemeStyleboxOverride(gc.StringName(name), stylebox)
}
func (self Simple) AddThemeFontOverride(name string, font [1]classdb.Font) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddThemeFontOverride(gc.StringName(name), font)
}
func (self Simple) AddThemeFontSizeOverride(name string, font_size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddThemeFontSizeOverride(gc.StringName(name), gd.Int(font_size))
}
func (self Simple) AddThemeColorOverride(name string, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddThemeColorOverride(gc.StringName(name), color)
}
func (self Simple) AddThemeConstantOverride(name string, constant int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddThemeConstantOverride(gc.StringName(name), gd.Int(constant))
}
func (self Simple) RemoveThemeIconOverride(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveThemeIconOverride(gc.StringName(name))
}
func (self Simple) RemoveThemeStyleboxOverride(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveThemeStyleboxOverride(gc.StringName(name))
}
func (self Simple) RemoveThemeFontOverride(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveThemeFontOverride(gc.StringName(name))
}
func (self Simple) RemoveThemeFontSizeOverride(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveThemeFontSizeOverride(gc.StringName(name))
}
func (self Simple) RemoveThemeColorOverride(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveThemeColorOverride(gc.StringName(name))
}
func (self Simple) RemoveThemeConstantOverride(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveThemeConstantOverride(gc.StringName(name))
}
func (self Simple) GetThemeIcon(name string, theme_type string) [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetThemeIcon(gc, gc.StringName(name), gc.StringName(theme_type)))
}
func (self Simple) GetThemeStylebox(name string, theme_type string) [1]classdb.StyleBox {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.StyleBox(Expert(self).GetThemeStylebox(gc, gc.StringName(name), gc.StringName(theme_type)))
}
func (self Simple) GetThemeFont(name string, theme_type string) [1]classdb.Font {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Font(Expert(self).GetThemeFont(gc, gc.StringName(name), gc.StringName(theme_type)))
}
func (self Simple) GetThemeFontSize(name string, theme_type string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetThemeFontSize(gc.StringName(name), gc.StringName(theme_type))))
}
func (self Simple) GetThemeColor(name string, theme_type string) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetThemeColor(gc.StringName(name), gc.StringName(theme_type)))
}
func (self Simple) GetThemeConstant(name string, theme_type string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetThemeConstant(gc.StringName(name), gc.StringName(theme_type))))
}
func (self Simple) HasThemeIconOverride(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeIconOverride(gc.StringName(name)))
}
func (self Simple) HasThemeStyleboxOverride(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeStyleboxOverride(gc.StringName(name)))
}
func (self Simple) HasThemeFontOverride(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeFontOverride(gc.StringName(name)))
}
func (self Simple) HasThemeFontSizeOverride(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeFontSizeOverride(gc.StringName(name)))
}
func (self Simple) HasThemeColorOverride(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeColorOverride(gc.StringName(name)))
}
func (self Simple) HasThemeConstantOverride(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeConstantOverride(gc.StringName(name)))
}
func (self Simple) HasThemeIcon(name string, theme_type string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeIcon(gc.StringName(name), gc.StringName(theme_type)))
}
func (self Simple) HasThemeStylebox(name string, theme_type string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeStylebox(gc.StringName(name), gc.StringName(theme_type)))
}
func (self Simple) HasThemeFont(name string, theme_type string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeFont(gc.StringName(name), gc.StringName(theme_type)))
}
func (self Simple) HasThemeFontSize(name string, theme_type string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeFontSize(gc.StringName(name), gc.StringName(theme_type)))
}
func (self Simple) HasThemeColor(name string, theme_type string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeColor(gc.StringName(name), gc.StringName(theme_type)))
}
func (self Simple) HasThemeConstant(name string, theme_type string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeConstant(gc.StringName(name), gc.StringName(theme_type)))
}
func (self Simple) GetThemeDefaultBaseScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetThemeDefaultBaseScale()))
}
func (self Simple) GetThemeDefaultFont() [1]classdb.Font {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Font(Expert(self).GetThemeDefaultFont(gc))
}
func (self Simple) GetThemeDefaultFontSize() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetThemeDefaultFontSize()))
}
func (self Simple) SetLayoutDirection(direction classdb.WindowLayoutDirection) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLayoutDirection(direction)
}
func (self Simple) GetLayoutDirection() classdb.WindowLayoutDirection {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.WindowLayoutDirection(Expert(self).GetLayoutDirection())
}
func (self Simple) IsLayoutRtl() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsLayoutRtl())
}
func (self Simple) SetAutoTranslate(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAutoTranslate(enable)
}
func (self Simple) IsAutoTranslating() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAutoTranslating())
}
func (self Simple) Popup(rect gd.Rect2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Popup(rect)
}
func (self Simple) PopupOnParent(parent_rect gd.Rect2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PopupOnParent(parent_rect)
}
func (self Simple) PopupCentered(minsize gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PopupCentered(minsize)
}
func (self Simple) PopupCenteredRatio(ratio float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PopupCenteredRatio(gd.Float(ratio))
}
func (self Simple) PopupCenteredClamped(minsize gd.Vector2i, fallback_ratio float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PopupCenteredClamped(minsize, gd.Float(fallback_ratio))
}
func (self Simple) PopupExclusive(from_node [1]classdb.Node, rect gd.Rect2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PopupExclusive(from_node, rect)
}
func (self Simple) PopupExclusiveOnParent(from_node [1]classdb.Node, parent_rect gd.Rect2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PopupExclusiveOnParent(from_node, parent_rect)
}
func (self Simple) PopupExclusiveCentered(from_node [1]classdb.Node, minsize gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PopupExclusiveCentered(from_node, minsize)
}
func (self Simple) PopupExclusiveCenteredRatio(from_node [1]classdb.Node, ratio float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PopupExclusiveCenteredRatio(from_node, gd.Float(ratio))
}
func (self Simple) PopupExclusiveCenteredClamped(from_node [1]classdb.Node, minsize gd.Vector2i, fallback_ratio float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PopupExclusiveCenteredClamped(from_node, minsize, gd.Float(fallback_ratio))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Window
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Virtual method to be implemented by the user. Overrides the value returned by [method get_contents_minimum_size].
*/
func (class) _get_contents_minimum_size(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

//go:nosplit
func (self class) SetTitle(title gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(title))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTitle(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the ID of the window.
*/
//go:nosplit
func (self class) GetWindowId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_window_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInitialPosition(initial_position classdb.WindowWindowInitialPosition)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, initial_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_initial_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInitialPosition() classdb.WindowWindowInitialPosition {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.WindowWindowInitialPosition](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_initial_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCurrentScreen(index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_current_screen, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCurrentScreen() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_current_screen, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPosition(position gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPosition() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Centers a native window on the current screen and an embedded window on its embedder [Viewport].
*/
//go:nosplit
func (self class) MoveToCenter()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_move_to_center, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetSize(size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSize() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Resets the size to the minimum size, which is the max of [member min_size] and (if [member wrap_controls] is enabled) [method get_contents_minimum_size]. This is equivalent to calling [code]set_size(Vector2i())[/code] (or any size below the minimum).
*/
//go:nosplit
func (self class) ResetSize()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_reset_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the window's position including its border.
[b]Note:[/b] If [member visible] is [code]false[/code], this method returns the same value as [member position].
*/
//go:nosplit
func (self class) GetPositionWithDecorations() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_position_with_decorations, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the window's size including its border.
[b]Note:[/b] If [member visible] is [code]false[/code], this method returns the same value as [member size].
*/
//go:nosplit
func (self class) GetSizeWithDecorations() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_size_with_decorations, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxSize(max_size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_max_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxSize() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_max_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMinSize(min_size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, min_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_min_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMinSize() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_min_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMode(mode classdb.WindowMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMode() classdb.WindowMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.WindowMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a specified window flag.
*/
//go:nosplit
func (self class) SetFlag(flag classdb.WindowFlags, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the [param flag] is set.
*/
//go:nosplit
func (self class) GetFlag(flag classdb.WindowFlags) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the window can be maximized (the maximize button is enabled).
*/
//go:nosplit
func (self class) IsMaximizeAllowed() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_is_maximize_allowed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Tells the OS that the [Window] needs an attention. This makes the window stand out in some way depending on the system, e.g. it might blink on the task bar.
*/
//go:nosplit
func (self class) RequestAttention()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_request_attention, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Causes the window to grab focus, allowing it to receive user input.
*/
//go:nosplit
func (self class) MoveToForeground()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_move_to_foreground, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetVisible(visible bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsVisible() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_is_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Hides the window. This is not the same as minimized state. Hidden window can't be interacted with and needs to be made visible with [method show].
*/
//go:nosplit
func (self class) Hide()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_hide, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Makes the [Window] appear. This enables interactions with the [Window] and doesn't change any of its property other than visibility (unlike e.g. [method popup]).
*/
//go:nosplit
func (self class) Show()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_show, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetTransient(transient bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, transient)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_transient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsTransient() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_is_transient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTransientToFocused(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_transient_to_focused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsTransientToFocused() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_is_transient_to_focused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetExclusive(exclusive bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, exclusive)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_exclusive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsExclusive() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_is_exclusive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param unparent] is [code]true[/code], the window is automatically unparented when going invisible.
[b]Note:[/b] Make sure to keep a reference to the node, otherwise it will be orphaned. You also need to manually call [method Node.queue_free] to free the window if it's not parented.
*/
//go:nosplit
func (self class) SetUnparentWhenInvisible(unparent bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, unparent)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_unparent_when_invisible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the window is being drawn to the screen.
*/
//go:nosplit
func (self class) CanDraw() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_can_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the window is focused.
*/
//go:nosplit
func (self class) HasFocus() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_has_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Causes the window to grab focus, allowing it to receive user input.
*/
//go:nosplit
func (self class) GrabFocus()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_grab_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [param active] is [code]true[/code], enables system's native IME (Input Method Editor).
*/
//go:nosplit
func (self class) SetImeActive(active bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_ime_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Moves IME to the given position.
*/
//go:nosplit
func (self class) SetImePosition(position gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_ime_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the window is currently embedded in another window.
*/
//go:nosplit
func (self class) IsEmbedded() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_is_embedded, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the combined minimum size from the child [Control] nodes of the window. Use [method child_controls_changed] to update it when child nodes have changed.
The value returned by this method can be overridden with [method _get_contents_minimum_size].
*/
//go:nosplit
func (self class) GetContentsMinimumSize() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_contents_minimum_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetForceNative(force_native bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, force_native)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_force_native, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetForceNative() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_force_native, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetContentScaleSize(size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_content_scale_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetContentScaleSize() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_content_scale_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetContentScaleMode(mode classdb.WindowContentScaleMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_content_scale_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetContentScaleMode() classdb.WindowContentScaleMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.WindowContentScaleMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_content_scale_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetContentScaleAspect(aspect classdb.WindowContentScaleAspect)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, aspect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_content_scale_aspect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetContentScaleAspect() classdb.WindowContentScaleAspect {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.WindowContentScaleAspect](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_content_scale_aspect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetContentScaleStretch(stretch classdb.WindowContentScaleStretch)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stretch)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_content_scale_stretch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetContentScaleStretch() classdb.WindowContentScaleStretch {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.WindowContentScaleStretch](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_content_scale_stretch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetKeepTitleVisible(title_visible bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, title_visible)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_keep_title_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetKeepTitleVisible() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_keep_title_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetContentScaleFactor(factor gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, factor)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_content_scale_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetContentScaleFactor() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_content_scale_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Enables font oversampling. This makes fonts look better when they are scaled up.
*/
//go:nosplit
func (self class) SetUseFontOversampling(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_use_font_oversampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if font oversampling is enabled. See [method set_use_font_oversampling].
*/
//go:nosplit
func (self class) IsUsingFontOversampling() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_is_using_font_oversampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMousePassthroughPolygon(polygon gd.PackedVector2Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(polygon))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_mouse_passthrough_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMousePassthroughPolygon(ctx gd.Lifetime) gd.PackedVector2Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_mouse_passthrough_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector2Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWrapControls(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_wrap_controls, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsWrappingControls() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_is_wrapping_controls, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Requests an update of the [Window] size to fit underlying [Control] nodes.
*/
//go:nosplit
func (self class) ChildControlsChanged()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_child_controls_changed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetTheme(theme object.Theme)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(theme[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_theme, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTheme(ctx gd.Lifetime) object.Theme {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_theme, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Theme
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetThemeTypeVariation(theme_type gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_theme_type_variation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetThemeTypeVariation(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_theme_type_variation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Prevents [code]*_theme_*_override[/code] methods from emitting [constant NOTIFICATION_THEME_CHANGED] until [method end_bulk_theme_override] is called.
*/
//go:nosplit
func (self class) BeginBulkThemeOverride()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_begin_bulk_theme_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Ends a bulk theme override update. See [method begin_bulk_theme_override].
*/
//go:nosplit
func (self class) EndBulkThemeOverride()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_end_bulk_theme_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a local override for a theme icon with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_icon_override].
See also [method get_theme_icon].
*/
//go:nosplit
func (self class) AddThemeIconOverride(name gd.StringName, texture object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_add_theme_icon_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a local override for a theme [StyleBox] with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_stylebox_override].
See also [method get_theme_stylebox] and [method Control.add_theme_stylebox_override] for more details.
*/
//go:nosplit
func (self class) AddThemeStyleboxOverride(name gd.StringName, stylebox object.StyleBox)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(stylebox[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_add_theme_stylebox_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a local override for a theme [Font] with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_font_override].
See also [method get_theme_font].
*/
//go:nosplit
func (self class) AddThemeFontOverride(name gd.StringName, font object.Font)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(font[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_add_theme_font_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a local override for a theme font size with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_font_size_override].
See also [method get_theme_font_size].
*/
//go:nosplit
func (self class) AddThemeFontSizeOverride(name gd.StringName, font_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, font_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_add_theme_font_size_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a local override for a theme [Color] with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_color_override].
See also [method get_theme_color] and [method Control.add_theme_color_override] for more details.
*/
//go:nosplit
func (self class) AddThemeColorOverride(name gd.StringName, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_add_theme_color_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a local override for a theme constant with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_constant_override].
See also [method get_theme_constant].
*/
//go:nosplit
func (self class) AddThemeConstantOverride(name gd.StringName, constant gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, constant)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_add_theme_constant_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a local override for a theme icon with the specified [param name] previously added by [method add_theme_icon_override] or via the Inspector dock.
*/
//go:nosplit
func (self class) RemoveThemeIconOverride(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_remove_theme_icon_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a local override for a theme [StyleBox] with the specified [param name] previously added by [method add_theme_stylebox_override] or via the Inspector dock.
*/
//go:nosplit
func (self class) RemoveThemeStyleboxOverride(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_remove_theme_stylebox_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a local override for a theme [Font] with the specified [param name] previously added by [method add_theme_font_override] or via the Inspector dock.
*/
//go:nosplit
func (self class) RemoveThemeFontOverride(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_remove_theme_font_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a local override for a theme font size with the specified [param name] previously added by [method add_theme_font_size_override] or via the Inspector dock.
*/
//go:nosplit
func (self class) RemoveThemeFontSizeOverride(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_remove_theme_font_size_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a local override for a theme [Color] with the specified [param name] previously added by [method add_theme_color_override] or via the Inspector dock.
*/
//go:nosplit
func (self class) RemoveThemeColorOverride(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_remove_theme_color_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a local override for a theme constant with the specified [param name] previously added by [method add_theme_constant_override] or via the Inspector dock.
*/
//go:nosplit
func (self class) RemoveThemeConstantOverride(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_remove_theme_constant_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an icon from the first matching [Theme] in the tree if that [Theme] has an icon item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
//go:nosplit
func (self class) GetThemeIcon(ctx gd.Lifetime, name gd.StringName, theme_type gd.StringName) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_theme_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns a [StyleBox] from the first matching [Theme] in the tree if that [Theme] has a stylebox item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
//go:nosplit
func (self class) GetThemeStylebox(ctx gd.Lifetime, name gd.StringName, theme_type gd.StringName) object.StyleBox {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_theme_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.StyleBox
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns a [Font] from the first matching [Theme] in the tree if that [Theme] has a font item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
//go:nosplit
func (self class) GetThemeFont(ctx gd.Lifetime, name gd.StringName, theme_type gd.StringName) object.Font {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_theme_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Font
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns a font size from the first matching [Theme] in the tree if that [Theme] has a font size item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
//go:nosplit
func (self class) GetThemeFontSize(name gd.StringName, theme_type gd.StringName) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_theme_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a [Color] from the first matching [Theme] in the tree if that [Theme] has a color item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for more details.
*/
//go:nosplit
func (self class) GetThemeColor(name gd.StringName, theme_type gd.StringName) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_theme_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a constant from the first matching [Theme] in the tree if that [Theme] has a constant item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for more details.
*/
//go:nosplit
func (self class) GetThemeConstant(name gd.StringName, theme_type gd.StringName) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_theme_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a local override for a theme icon with the specified [param name] in this [Control] node.
See [method add_theme_icon_override].
*/
//go:nosplit
func (self class) HasThemeIconOverride(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_has_theme_icon_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a local override for a theme [StyleBox] with the specified [param name] in this [Control] node.
See [method add_theme_stylebox_override].
*/
//go:nosplit
func (self class) HasThemeStyleboxOverride(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_has_theme_stylebox_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a local override for a theme [Font] with the specified [param name] in this [Control] node.
See [method add_theme_font_override].
*/
//go:nosplit
func (self class) HasThemeFontOverride(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_has_theme_font_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a local override for a theme font size with the specified [param name] in this [Control] node.
See [method add_theme_font_size_override].
*/
//go:nosplit
func (self class) HasThemeFontSizeOverride(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_has_theme_font_size_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a local override for a theme [Color] with the specified [param name] in this [Control] node.
See [method add_theme_color_override].
*/
//go:nosplit
func (self class) HasThemeColorOverride(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_has_theme_color_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a local override for a theme constant with the specified [param name] in this [Control] node.
See [method add_theme_constant_override].
*/
//go:nosplit
func (self class) HasThemeConstantOverride(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_has_theme_constant_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a matching [Theme] in the tree that has an icon item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
//go:nosplit
func (self class) HasThemeIcon(name gd.StringName, theme_type gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_has_theme_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a matching [Theme] in the tree that has a stylebox item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
//go:nosplit
func (self class) HasThemeStylebox(name gd.StringName, theme_type gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_has_theme_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a matching [Theme] in the tree that has a font item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
//go:nosplit
func (self class) HasThemeFont(name gd.StringName, theme_type gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_has_theme_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a matching [Theme] in the tree that has a font size item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
//go:nosplit
func (self class) HasThemeFontSize(name gd.StringName, theme_type gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_has_theme_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a matching [Theme] in the tree that has a color item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
//go:nosplit
func (self class) HasThemeColor(name gd.StringName, theme_type gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_has_theme_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a matching [Theme] in the tree that has a constant item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
//go:nosplit
func (self class) HasThemeConstant(name gd.StringName, theme_type gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_has_theme_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the default base scale value from the first matching [Theme] in the tree if that [Theme] has a valid [member Theme.default_base_scale] value.
See [method Control.get_theme_color] for details.
*/
//go:nosplit
func (self class) GetThemeDefaultBaseScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_theme_default_base_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the default font from the first matching [Theme] in the tree if that [Theme] has a valid [member Theme.default_font] value.
See [method Control.get_theme_color] for details.
*/
//go:nosplit
func (self class) GetThemeDefaultFont(ctx gd.Lifetime) object.Font {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_theme_default_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Font
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the default font size value from the first matching [Theme] in the tree if that [Theme] has a valid [member Theme.default_font_size] value.
See [method Control.get_theme_color] for details.
*/
//go:nosplit
func (self class) GetThemeDefaultFontSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_theme_default_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets layout direction and text writing direction. Right-to-left layouts are necessary for certain languages (e.g. Arabic and Hebrew).
*/
//go:nosplit
func (self class) SetLayoutDirection(direction classdb.WindowLayoutDirection)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_layout_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns layout direction and text writing direction.
*/
//go:nosplit
func (self class) GetLayoutDirection() classdb.WindowLayoutDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.WindowLayoutDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_get_layout_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if layout is right-to-left.
*/
//go:nosplit
func (self class) IsLayoutRtl() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_is_layout_rtl, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoTranslate(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_set_auto_translate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAutoTranslating() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_is_auto_translating, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Shows the [Window] and makes it transient (see [member transient]). If [param rect] is provided, it will be set as the [Window]'s size. Fails if called on the main window.
If [member ProjectSettings.display/window/subwindows/embed_subwindows] is [code]true[/code] (single-window mode), [param rect]'s coordinates are global and relative to the main window's top-left corner (excluding window decorations). If [param rect]'s position coordinates are negative, the window will be located outside the main window and may not be visible as a result.
If [member ProjectSettings.display/window/subwindows/embed_subwindows] is [code]false[/code] (multi-window mode), [param rect]'s coordinates are global and relative to the top-left corner of the leftmost screen. If [param rect]'s position coordinates are negative, the window will be placed at the top-left corner of the screen.
[b]Note:[/b] [param rect] must be in global coordinates if specified.
*/
//go:nosplit
func (self class) Popup(rect gd.Rect2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_popup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Popups the [Window] with a position shifted by parent [Window]'s position. If the [Window] is embedded, has the same effect as [method popup].
*/
//go:nosplit
func (self class) PopupOnParent(parent_rect gd.Rect2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, parent_rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_popup_on_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Popups the [Window] at the center of the current screen, with optionally given minimum size. If the [Window] is embedded, it will be centered in the parent [Viewport] instead.
[b]Note:[/b] Calling it with the default value of [param minsize] is equivalent to calling it with [member size].
*/
//go:nosplit
func (self class) PopupCentered(minsize gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, minsize)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_popup_centered, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [Window] is embedded, popups the [Window] centered inside its embedder and sets its size as a [param ratio] of embedder's size.
If [Window] is a native window, popups the [Window] centered inside the screen of its parent [Window] and sets its size as a [param ratio] of the screen size.
*/
//go:nosplit
func (self class) PopupCenteredRatio(ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_popup_centered_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Popups the [Window] centered inside its parent [Window]. [param fallback_ratio] determines the maximum size of the [Window], in relation to its parent.
[b]Note:[/b] Calling it with the default value of [param minsize] is equivalent to calling it with [member size].
*/
//go:nosplit
func (self class) PopupCenteredClamped(minsize gd.Vector2i, fallback_ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, minsize)
	callframe.Arg(frame, fallback_ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_popup_centered_clamped, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Attempts to parent this dialog to the last exclusive window relative to [param from_node], and then calls [method Window.popup] on it. The dialog must have no current parent, otherwise the method fails.
See also [method set_unparent_when_invisible] and [method Node.get_last_exclusive_window].
*/
//go:nosplit
func (self class) PopupExclusive(from_node object.Node, rect gd.Rect2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(from_node[0].AsPointer())[0])
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_popup_exclusive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Attempts to parent this dialog to the last exclusive window relative to [param from_node], and then calls [method Window.popup_on_parent] on it. The dialog must have no current parent, otherwise the method fails.
See also [method set_unparent_when_invisible] and [method Node.get_last_exclusive_window].
*/
//go:nosplit
func (self class) PopupExclusiveOnParent(from_node object.Node, parent_rect gd.Rect2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(from_node[0].AsPointer())[0])
	callframe.Arg(frame, parent_rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_popup_exclusive_on_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Attempts to parent this dialog to the last exclusive window relative to [param from_node], and then calls [method Window.popup_centered] on it. The dialog must have no current parent, otherwise the method fails.
See also [method set_unparent_when_invisible] and [method Node.get_last_exclusive_window].
*/
//go:nosplit
func (self class) PopupExclusiveCentered(from_node object.Node, minsize gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(from_node[0].AsPointer())[0])
	callframe.Arg(frame, minsize)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_popup_exclusive_centered, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Attempts to parent this dialog to the last exclusive window relative to [param from_node], and then calls [method Window.popup_centered_ratio] on it. The dialog must have no current parent, otherwise the method fails.
See also [method set_unparent_when_invisible] and [method Node.get_last_exclusive_window].
*/
//go:nosplit
func (self class) PopupExclusiveCenteredRatio(from_node object.Node, ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(from_node[0].AsPointer())[0])
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_popup_exclusive_centered_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Attempts to parent this dialog to the last exclusive window relative to [param from_node], and then calls [method Window.popup_centered_clamped] on it. The dialog must have no current parent, otherwise the method fails.
See also [method set_unparent_when_invisible] and [method Node.get_last_exclusive_window].
*/
//go:nosplit
func (self class) PopupExclusiveCenteredClamped(from_node object.Node, minsize gd.Vector2i, fallback_ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(from_node[0].AsPointer())[0])
	callframe.Arg(frame, minsize)
	callframe.Arg(frame, fallback_ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Window.Bind_popup_exclusive_centered_clamped, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsWindow() Expert { return self[0].AsWindow() }


//go:nosplit
func (self Simple) AsWindow() Simple { return self[0].AsWindow() }


//go:nosplit
func (self class) AsViewport() Viewport.Expert { return self[0].AsViewport() }


//go:nosplit
func (self Simple) AsViewport() Viewport.Simple { return self[0].AsViewport() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_contents_minimum_size": return reflect.ValueOf(self._get_contents_minimum_size);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("Window", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Mode = classdb.WindowMode

const (
/*Windowed mode, i.e. [Window] doesn't occupy the whole screen (unless set to the size of the screen).*/
	ModeWindowed Mode = 0
/*Minimized window mode, i.e. [Window] is not visible and available on window manager's window list. Normally happens when the minimize button is pressed.*/
	ModeMinimized Mode = 1
/*Maximized window mode, i.e. [Window] will occupy whole screen area except task bar and still display its borders. Normally happens when the maximize button is pressed.*/
	ModeMaximized Mode = 2
/*Full screen mode with full multi-window support.
Full screen window covers the entire display area of a screen and has no decorations. The display's video mode is not changed.
[b]On Windows:[/b] Multi-window full-screen mode has a 1px border of the [member ProjectSettings.rendering/environment/defaults/default_clear_color] color.
[b]On macOS:[/b] A new desktop is used to display the running project.
[b]Note:[/b] Regardless of the platform, enabling full screen will change the window size to match the monitor's size. Therefore, make sure your project supports [url=$DOCS_URL/tutorials/rendering/multiple_resolutions.html]multiple resolutions[/url] when enabling full screen mode.*/
	ModeFullscreen Mode = 3
/*A single window full screen mode. This mode has less overhead, but only one window can be open on a given screen at a time (opening a child window or application switching will trigger a full screen transition).
Full screen window covers the entire display area of a screen and has no border or decorations. The display's video mode is not changed.
[b]On Windows:[/b] Depending on video driver, full screen transition might cause screens to go black for a moment.
[b]On macOS:[/b] A new desktop is used to display the running project. Exclusive full screen mode prevents Dock and Menu from showing up when the mouse pointer is hovering the edge of the screen.
[b]On Linux (X11):[/b] Exclusive full screen mode bypasses compositor.
[b]Note:[/b] Regardless of the platform, enabling full screen will change the window size to match the monitor's size. Therefore, make sure your project supports [url=$DOCS_URL/tutorials/rendering/multiple_resolutions.html]multiple resolutions[/url] when enabling full screen mode.*/
	ModeExclusiveFullscreen Mode = 4
)
type Flags = classdb.WindowFlags

const (
/*The window can't be resized by dragging its resize grip. It's still possible to resize the window using [member size]. This flag is ignored for full screen windows. Set with [member unresizable].*/
	FlagResizeDisabled Flags = 0
/*The window do not have native title bar and other decorations. This flag is ignored for full-screen windows. Set with [member borderless].*/
	FlagBorderless Flags = 1
/*The window is floating on top of all other windows. This flag is ignored for full-screen windows. Set with [member always_on_top].*/
	FlagAlwaysOnTop Flags = 2
/*The window background can be transparent. Set with [member transparent].
[b]Note:[/b] This flag has no effect if either [member ProjectSettings.display/window/per_pixel_transparency/allowed], or the window's [member Viewport.transparent_bg] is set to [code]false[/code].*/
	FlagTransparent Flags = 3
/*The window can't be focused. No-focus window will ignore all input, except mouse clicks. Set with [member unfocusable].*/
	FlagNoFocus Flags = 4
/*Window is part of menu or [OptionButton] dropdown. This flag can't be changed when the window is visible. An active popup window will exclusively receive all input, without stealing focus from its parent. Popup windows are automatically closed when uses click outside it, or when an application is switched. Popup window must have transient parent set (see [member transient]).
[b]Note:[/b] This flag has no effect in embedded windows (unless said window is a [Popup]).*/
	FlagPopup Flags = 5
/*Window content is expanded to the full size of the window. Unlike borderless window, the frame is left intact and can be used to resize the window, title bar is transparent, but have minimize/maximize/close buttons. Set with [member extend_to_title].
[b]Note:[/b] This flag is implemented only on macOS.
[b]Note:[/b] This flag has no effect in embedded windows.*/
	FlagExtendToTitle Flags = 6
/*All mouse events are passed to the underlying window of the same application.
[b]Note:[/b] This flag has no effect in embedded windows.*/
	FlagMousePassthrough Flags = 7
/*Max value of the [enum Flags].*/
	FlagMax Flags = 8
)
type ContentScaleMode = classdb.WindowContentScaleMode

const (
/*The content will not be scaled to match the [Window]'s size.*/
	ContentScaleModeDisabled ContentScaleMode = 0
/*The content will be rendered at the target size. This is more performance-expensive than [constant CONTENT_SCALE_MODE_VIEWPORT], but provides better results.*/
	ContentScaleModeCanvasItems ContentScaleMode = 1
/*The content will be rendered at the base size and then scaled to the target size. More performant than [constant CONTENT_SCALE_MODE_CANVAS_ITEMS], but results in pixelated image.*/
	ContentScaleModeViewport ContentScaleMode = 2
)
type ContentScaleAspect = classdb.WindowContentScaleAspect

const (
/*The aspect will be ignored. Scaling will simply stretch the content to fit the target size.*/
	ContentScaleAspectIgnore ContentScaleAspect = 0
/*The content's aspect will be preserved. If the target size has different aspect from the base one, the image will be centered and black bars will appear on left and right sides.*/
	ContentScaleAspectKeep ContentScaleAspect = 1
/*The content can be expanded vertically. Scaling horizontally will result in keeping the width ratio and then black bars on left and right sides.*/
	ContentScaleAspectKeepWidth ContentScaleAspect = 2
/*The content can be expanded horizontally. Scaling vertically will result in keeping the height ratio and then black bars on top and bottom sides.*/
	ContentScaleAspectKeepHeight ContentScaleAspect = 3
/*The content's aspect will be preserved. If the target size has different aspect from the base one, the content will stay in the top-left corner and add an extra visible area in the stretched space.*/
	ContentScaleAspectExpand ContentScaleAspect = 4
)
type ContentScaleStretch = classdb.WindowContentScaleStretch

const (
/*The content will be stretched according to a fractional factor. This fills all the space available in the window, but allows "pixel wobble" to occur due to uneven pixel scaling.*/
	ContentScaleStretchFractional ContentScaleStretch = 0
/*The content will be stretched only according to an integer factor, preserving sharp pixels. This may leave a black background visible on the window's edges depending on the window size.*/
	ContentScaleStretchInteger ContentScaleStretch = 1
)
type LayoutDirection = classdb.WindowLayoutDirection

const (
/*Automatic layout direction, determined from the parent window layout direction.*/
	LayoutDirectionInherited LayoutDirection = 0
/*Automatic layout direction, determined from the current locale.*/
	LayoutDirectionLocale LayoutDirection = 1
/*Left-to-right layout direction.*/
	LayoutDirectionLtr LayoutDirection = 2
/*Right-to-left layout direction.*/
	LayoutDirectionRtl LayoutDirection = 3
)
type WindowInitialPosition = classdb.WindowWindowInitialPosition

const (
/*Initial window position is determined by [member position].*/
	WindowInitialPositionAbsolute WindowInitialPosition = 0
/*Initial window position is the center of the primary screen.*/
	WindowInitialPositionCenterPrimaryScreen WindowInitialPosition = 1
/*Initial window position is the center of the main window screen.*/
	WindowInitialPositionCenterMainWindowScreen WindowInitialPosition = 2
/*Initial window position is the center of [member current_screen] screen.*/
	WindowInitialPositionCenterOtherScreen WindowInitialPosition = 3
/*Initial window position is the center of the screen containing the mouse pointer.*/
	WindowInitialPositionCenterScreenWithMouseFocus WindowInitialPosition = 4
/*Initial window position is the center of the screen containing the window with the keyboard focus.*/
	WindowInitialPositionCenterScreenWithKeyboardFocus WindowInitialPosition = 5
)
