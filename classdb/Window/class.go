// Package Window provides methods for working with Window object instances.
package Window

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/classdb/Viewport"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Vector2i"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Rect2i"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
A node that creates a window. The window can either be a native system window or embedded inside another [Window] (see [member Viewport.gui_embed_subwindows]).
At runtime, [Window]s will not close automatically when requested. You need to handle it manually using the [signal close_requested] signal (this applies both to pressing the close button and clicking outside of a popup).

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=Window)
*/
type Instance [1]gdclass.Window

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsWindow() Instance
}
type Interface interface {
	//Virtual method to be implemented by the user. Overrides the value returned by [method get_contents_minimum_size].
	GetContentsMinimumSize() Vector2.XY
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetContentsMinimumSize() (_ Vector2.XY) { return }

/*
Virtual method to be implemented by the user. Overrides the value returned by [method get_contents_minimum_size].
*/
func (Instance) _get_contents_minimum_size(impl func(ptr unsafe.Pointer) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
Returns the ID of the window.
*/
func (self Instance) GetWindowId() int {
	return int(int(class(self).GetWindowId()))
}

/*
Centers a native window on the current screen and an embedded window on its embedder [Viewport].
*/
func (self Instance) MoveToCenter() {
	class(self).MoveToCenter()
}

/*
Resets the size to the minimum size, which is the max of [member min_size] and (if [member wrap_controls] is enabled) [method get_contents_minimum_size]. This is equivalent to calling [code]set_size(Vector2i())[/code] (or any size below the minimum).
*/
func (self Instance) ResetSize() {
	class(self).ResetSize()
}

/*
Returns the window's position including its border.
[b]Note:[/b] If [member visible] is [code]false[/code], this method returns the same value as [member position].
*/
func (self Instance) GetPositionWithDecorations() Vector2i.XY {
	return Vector2i.XY(class(self).GetPositionWithDecorations())
}

/*
Returns the window's size including its border.
[b]Note:[/b] If [member visible] is [code]false[/code], this method returns the same value as [member size].
*/
func (self Instance) GetSizeWithDecorations() Vector2i.XY {
	return Vector2i.XY(class(self).GetSizeWithDecorations())
}

/*
Returns [code]true[/code] if the window can be maximized (the maximize button is enabled).
*/
func (self Instance) IsMaximizeAllowed() bool {
	return bool(class(self).IsMaximizeAllowed())
}

/*
Tells the OS that the [Window] needs an attention. This makes the window stand out in some way depending on the system, e.g. it might blink on the task bar.
*/
func (self Instance) RequestAttention() {
	class(self).RequestAttention()
}

/*
Causes the window to grab focus, allowing it to receive user input.
*/
func (self Instance) MoveToForeground() {
	class(self).MoveToForeground()
}

/*
Hides the window. This is not the same as minimized state. Hidden window can't be interacted with and needs to be made visible with [method show].
*/
func (self Instance) Hide() {
	class(self).Hide()
}

/*
Makes the [Window] appear. This enables interactions with the [Window] and doesn't change any of its property other than visibility (unlike e.g. [method popup]).
*/
func (self Instance) Show() {
	class(self).Show()
}

/*
If [param unparent] is [code]true[/code], the window is automatically unparented when going invisible.
[b]Note:[/b] Make sure to keep a reference to the node, otherwise it will be orphaned. You also need to manually call [method Node.queue_free] to free the window if it's not parented.
*/
func (self Instance) SetUnparentWhenInvisible(unparent bool) {
	class(self).SetUnparentWhenInvisible(unparent)
}

/*
Returns whether the window is being drawn to the screen.
*/
func (self Instance) CanDraw() bool {
	return bool(class(self).CanDraw())
}

/*
Returns [code]true[/code] if the window is focused.
*/
func (self Instance) HasFocus() bool {
	return bool(class(self).HasFocus())
}

/*
Causes the window to grab focus, allowing it to receive user input.
*/
func (self Instance) GrabFocus() {
	class(self).GrabFocus()
}

/*
If [param active] is [code]true[/code], enables system's native IME (Input Method Editor).
*/
func (self Instance) SetImeActive(active bool) {
	class(self).SetImeActive(active)
}

/*
Moves IME to the given position.
*/
func (self Instance) SetImePosition(position Vector2i.XY) {
	class(self).SetImePosition(gd.Vector2i(position))
}

/*
Returns [code]true[/code] if the window is currently embedded in another window.
*/
func (self Instance) IsEmbedded() bool {
	return bool(class(self).IsEmbedded())
}

/*
Returns the combined minimum size from the child [Control] nodes of the window. Use [method child_controls_changed] to update it when child nodes have changed.
The value returned by this method can be overridden with [method _get_contents_minimum_size].
*/
func (self Instance) GetContentsMinimumSize() Vector2.XY {
	return Vector2.XY(class(self).GetContentsMinimumSize())
}

/*
Enables font oversampling. This makes fonts look better when they are scaled up.
*/
func (self Instance) SetUseFontOversampling(enable bool) {
	class(self).SetUseFontOversampling(enable)
}

/*
Returns [code]true[/code] if font oversampling is enabled. See [method set_use_font_oversampling].
*/
func (self Instance) IsUsingFontOversampling() bool {
	return bool(class(self).IsUsingFontOversampling())
}

/*
Requests an update of the [Window] size to fit underlying [Control] nodes.
*/
func (self Instance) ChildControlsChanged() {
	class(self).ChildControlsChanged()
}

/*
Prevents [code]*_theme_*_override[/code] methods from emitting [constant NOTIFICATION_THEME_CHANGED] until [method end_bulk_theme_override] is called.
*/
func (self Instance) BeginBulkThemeOverride() {
	class(self).BeginBulkThemeOverride()
}

/*
Ends a bulk theme override update. See [method begin_bulk_theme_override].
*/
func (self Instance) EndBulkThemeOverride() {
	class(self).EndBulkThemeOverride()
}

/*
Creates a local override for a theme icon with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_icon_override].
See also [method get_theme_icon].
*/
func (self Instance) AddThemeIconOverride(name string, texture [1]gdclass.Texture2D) {
	class(self).AddThemeIconOverride(gd.NewStringName(name), texture)
}

/*
Creates a local override for a theme [StyleBox] with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_stylebox_override].
See also [method get_theme_stylebox] and [method Control.add_theme_stylebox_override] for more details.
*/
func (self Instance) AddThemeStyleboxOverride(name string, stylebox [1]gdclass.StyleBox) {
	class(self).AddThemeStyleboxOverride(gd.NewStringName(name), stylebox)
}

/*
Creates a local override for a theme [Font] with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_font_override].
See also [method get_theme_font].
*/
func (self Instance) AddThemeFontOverride(name string, font [1]gdclass.Font) {
	class(self).AddThemeFontOverride(gd.NewStringName(name), font)
}

/*
Creates a local override for a theme font size with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_font_size_override].
See also [method get_theme_font_size].
*/
func (self Instance) AddThemeFontSizeOverride(name string, font_size int) {
	class(self).AddThemeFontSizeOverride(gd.NewStringName(name), gd.Int(font_size))
}

/*
Creates a local override for a theme [Color] with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_color_override].
See also [method get_theme_color] and [method Control.add_theme_color_override] for more details.
*/
func (self Instance) AddThemeColorOverride(name string, color Color.RGBA) {
	class(self).AddThemeColorOverride(gd.NewStringName(name), gd.Color(color))
}

/*
Creates a local override for a theme constant with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_constant_override].
See also [method get_theme_constant].
*/
func (self Instance) AddThemeConstantOverride(name string, constant int) {
	class(self).AddThemeConstantOverride(gd.NewStringName(name), gd.Int(constant))
}

/*
Removes a local override for a theme icon with the specified [param name] previously added by [method add_theme_icon_override] or via the Inspector dock.
*/
func (self Instance) RemoveThemeIconOverride(name string) {
	class(self).RemoveThemeIconOverride(gd.NewStringName(name))
}

/*
Removes a local override for a theme [StyleBox] with the specified [param name] previously added by [method add_theme_stylebox_override] or via the Inspector dock.
*/
func (self Instance) RemoveThemeStyleboxOverride(name string) {
	class(self).RemoveThemeStyleboxOverride(gd.NewStringName(name))
}

/*
Removes a local override for a theme [Font] with the specified [param name] previously added by [method add_theme_font_override] or via the Inspector dock.
*/
func (self Instance) RemoveThemeFontOverride(name string) {
	class(self).RemoveThemeFontOverride(gd.NewStringName(name))
}

/*
Removes a local override for a theme font size with the specified [param name] previously added by [method add_theme_font_size_override] or via the Inspector dock.
*/
func (self Instance) RemoveThemeFontSizeOverride(name string) {
	class(self).RemoveThemeFontSizeOverride(gd.NewStringName(name))
}

/*
Removes a local override for a theme [Color] with the specified [param name] previously added by [method add_theme_color_override] or via the Inspector dock.
*/
func (self Instance) RemoveThemeColorOverride(name string) {
	class(self).RemoveThemeColorOverride(gd.NewStringName(name))
}

/*
Removes a local override for a theme constant with the specified [param name] previously added by [method add_theme_constant_override] or via the Inspector dock.
*/
func (self Instance) RemoveThemeConstantOverride(name string) {
	class(self).RemoveThemeConstantOverride(gd.NewStringName(name))
}

/*
Returns an icon from the first matching [Theme] in the tree if that [Theme] has an icon item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
func (self Instance) GetThemeIcon(name string) [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetThemeIcon(gd.NewStringName(name), gd.NewStringName("")))
}

/*
Returns a [StyleBox] from the first matching [Theme] in the tree if that [Theme] has a stylebox item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
func (self Instance) GetThemeStylebox(name string) [1]gdclass.StyleBox {
	return [1]gdclass.StyleBox(class(self).GetThemeStylebox(gd.NewStringName(name), gd.NewStringName("")))
}

/*
Returns a [Font] from the first matching [Theme] in the tree if that [Theme] has a font item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
func (self Instance) GetThemeFont(name string) [1]gdclass.Font {
	return [1]gdclass.Font(class(self).GetThemeFont(gd.NewStringName(name), gd.NewStringName("")))
}

/*
Returns a font size from the first matching [Theme] in the tree if that [Theme] has a font size item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
func (self Instance) GetThemeFontSize(name string) int {
	return int(int(class(self).GetThemeFontSize(gd.NewStringName(name), gd.NewStringName(""))))
}

/*
Returns a [Color] from the first matching [Theme] in the tree if that [Theme] has a color item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for more details.
*/
func (self Instance) GetThemeColor(name string) Color.RGBA {
	return Color.RGBA(class(self).GetThemeColor(gd.NewStringName(name), gd.NewStringName("")))
}

/*
Returns a constant from the first matching [Theme] in the tree if that [Theme] has a constant item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for more details.
*/
func (self Instance) GetThemeConstant(name string) int {
	return int(int(class(self).GetThemeConstant(gd.NewStringName(name), gd.NewStringName(""))))
}

/*
Returns [code]true[/code] if there is a local override for a theme icon with the specified [param name] in this [Control] node.
See [method add_theme_icon_override].
*/
func (self Instance) HasThemeIconOverride(name string) bool {
	return bool(class(self).HasThemeIconOverride(gd.NewStringName(name)))
}

/*
Returns [code]true[/code] if there is a local override for a theme [StyleBox] with the specified [param name] in this [Control] node.
See [method add_theme_stylebox_override].
*/
func (self Instance) HasThemeStyleboxOverride(name string) bool {
	return bool(class(self).HasThemeStyleboxOverride(gd.NewStringName(name)))
}

/*
Returns [code]true[/code] if there is a local override for a theme [Font] with the specified [param name] in this [Control] node.
See [method add_theme_font_override].
*/
func (self Instance) HasThemeFontOverride(name string) bool {
	return bool(class(self).HasThemeFontOverride(gd.NewStringName(name)))
}

/*
Returns [code]true[/code] if there is a local override for a theme font size with the specified [param name] in this [Control] node.
See [method add_theme_font_size_override].
*/
func (self Instance) HasThemeFontSizeOverride(name string) bool {
	return bool(class(self).HasThemeFontSizeOverride(gd.NewStringName(name)))
}

/*
Returns [code]true[/code] if there is a local override for a theme [Color] with the specified [param name] in this [Control] node.
See [method add_theme_color_override].
*/
func (self Instance) HasThemeColorOverride(name string) bool {
	return bool(class(self).HasThemeColorOverride(gd.NewStringName(name)))
}

/*
Returns [code]true[/code] if there is a local override for a theme constant with the specified [param name] in this [Control] node.
See [method add_theme_constant_override].
*/
func (self Instance) HasThemeConstantOverride(name string) bool {
	return bool(class(self).HasThemeConstantOverride(gd.NewStringName(name)))
}

/*
Returns [code]true[/code] if there is a matching [Theme] in the tree that has an icon item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
func (self Instance) HasThemeIcon(name string) bool {
	return bool(class(self).HasThemeIcon(gd.NewStringName(name), gd.NewStringName("")))
}

/*
Returns [code]true[/code] if there is a matching [Theme] in the tree that has a stylebox item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
func (self Instance) HasThemeStylebox(name string) bool {
	return bool(class(self).HasThemeStylebox(gd.NewStringName(name), gd.NewStringName("")))
}

/*
Returns [code]true[/code] if there is a matching [Theme] in the tree that has a font item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
func (self Instance) HasThemeFont(name string) bool {
	return bool(class(self).HasThemeFont(gd.NewStringName(name), gd.NewStringName("")))
}

/*
Returns [code]true[/code] if there is a matching [Theme] in the tree that has a font size item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
func (self Instance) HasThemeFontSize(name string) bool {
	return bool(class(self).HasThemeFontSize(gd.NewStringName(name), gd.NewStringName("")))
}

/*
Returns [code]true[/code] if there is a matching [Theme] in the tree that has a color item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
func (self Instance) HasThemeColor(name string) bool {
	return bool(class(self).HasThemeColor(gd.NewStringName(name), gd.NewStringName("")))
}

/*
Returns [code]true[/code] if there is a matching [Theme] in the tree that has a constant item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
func (self Instance) HasThemeConstant(name string) bool {
	return bool(class(self).HasThemeConstant(gd.NewStringName(name), gd.NewStringName("")))
}

/*
Returns the default base scale value from the first matching [Theme] in the tree if that [Theme] has a valid [member Theme.default_base_scale] value.
See [method Control.get_theme_color] for details.
*/
func (self Instance) GetThemeDefaultBaseScale() Float.X {
	return Float.X(Float.X(class(self).GetThemeDefaultBaseScale()))
}

/*
Returns the default font from the first matching [Theme] in the tree if that [Theme] has a valid [member Theme.default_font] value.
See [method Control.get_theme_color] for details.
*/
func (self Instance) GetThemeDefaultFont() [1]gdclass.Font {
	return [1]gdclass.Font(class(self).GetThemeDefaultFont())
}

/*
Returns the default font size value from the first matching [Theme] in the tree if that [Theme] has a valid [member Theme.default_font_size] value.
See [method Control.get_theme_color] for details.
*/
func (self Instance) GetThemeDefaultFontSize() int {
	return int(int(class(self).GetThemeDefaultFontSize()))
}

/*
Sets layout direction and text writing direction. Right-to-left layouts are necessary for certain languages (e.g. Arabic and Hebrew).
*/
func (self Instance) SetLayoutDirection(direction gdclass.WindowLayoutDirection) {
	class(self).SetLayoutDirection(direction)
}

/*
Returns layout direction and text writing direction.
*/
func (self Instance) GetLayoutDirection() gdclass.WindowLayoutDirection {
	return gdclass.WindowLayoutDirection(class(self).GetLayoutDirection())
}

/*
Returns [code]true[/code] if layout is right-to-left.
*/
func (self Instance) IsLayoutRtl() bool {
	return bool(class(self).IsLayoutRtl())
}

/*
Shows the [Window] and makes it transient (see [member transient]). If [param rect] is provided, it will be set as the [Window]'s size. Fails if called on the main window.
If [member ProjectSettings.display/window/subwindows/embed_subwindows] is [code]true[/code] (single-window mode), [param rect]'s coordinates are global and relative to the main window's top-left corner (excluding window decorations). If [param rect]'s position coordinates are negative, the window will be located outside the main window and may not be visible as a result.
If [member ProjectSettings.display/window/subwindows/embed_subwindows] is [code]false[/code] (multi-window mode), [param rect]'s coordinates are global and relative to the top-left corner of the leftmost screen. If [param rect]'s position coordinates are negative, the window will be placed at the top-left corner of the screen.
[b]Note:[/b] [param rect] must be in global coordinates if specified.
*/
func (self Instance) Popup() {
	class(self).Popup(gd.Rect2i(gd.NewRect2i(0, 0, 0, 0)))
}

/*
Popups the [Window] with a position shifted by parent [Window]'s position. If the [Window] is embedded, has the same effect as [method popup].
*/
func (self Instance) PopupOnParent(parent_rect Rect2i.PositionSize) {
	class(self).PopupOnParent(gd.Rect2i(parent_rect))
}

/*
Popups the [Window] at the center of the current screen, with optionally given minimum size. If the [Window] is embedded, it will be centered in the parent [Viewport] instead.
[b]Note:[/b] Calling it with the default value of [param minsize] is equivalent to calling it with [member size].
*/
func (self Instance) PopupCentered() {
	class(self).PopupCentered(gd.Vector2i(gd.Vector2i{0, 0}))
}

/*
If [Window] is embedded, popups the [Window] centered inside its embedder and sets its size as a [param ratio] of embedder's size.
If [Window] is a native window, popups the [Window] centered inside the screen of its parent [Window] and sets its size as a [param ratio] of the screen size.
*/
func (self Instance) PopupCenteredRatio() {
	class(self).PopupCenteredRatio(gd.Float(0.8))
}

/*
Popups the [Window] centered inside its parent [Window]. [param fallback_ratio] determines the maximum size of the [Window], in relation to its parent.
[b]Note:[/b] Calling it with the default value of [param minsize] is equivalent to calling it with [member size].
*/
func (self Instance) PopupCenteredClamped() {
	class(self).PopupCenteredClamped(gd.Vector2i(gd.Vector2i{0, 0}), gd.Float(0.75))
}

/*
Attempts to parent this dialog to the last exclusive window relative to [param from_node], and then calls [method Window.popup] on it. The dialog must have no current parent, otherwise the method fails.
See also [method set_unparent_when_invisible] and [method Node.get_last_exclusive_window].
*/
func (self Instance) PopupExclusive(from_node [1]gdclass.Node) {
	class(self).PopupExclusive(from_node, gd.Rect2i(gd.NewRect2i(0, 0, 0, 0)))
}

/*
Attempts to parent this dialog to the last exclusive window relative to [param from_node], and then calls [method Window.popup_on_parent] on it. The dialog must have no current parent, otherwise the method fails.
See also [method set_unparent_when_invisible] and [method Node.get_last_exclusive_window].
*/
func (self Instance) PopupExclusiveOnParent(from_node [1]gdclass.Node, parent_rect Rect2i.PositionSize) {
	class(self).PopupExclusiveOnParent(from_node, gd.Rect2i(parent_rect))
}

/*
Attempts to parent this dialog to the last exclusive window relative to [param from_node], and then calls [method Window.popup_centered] on it. The dialog must have no current parent, otherwise the method fails.
See also [method set_unparent_when_invisible] and [method Node.get_last_exclusive_window].
*/
func (self Instance) PopupExclusiveCentered(from_node [1]gdclass.Node) {
	class(self).PopupExclusiveCentered(from_node, gd.Vector2i(gd.Vector2i{0, 0}))
}

/*
Attempts to parent this dialog to the last exclusive window relative to [param from_node], and then calls [method Window.popup_centered_ratio] on it. The dialog must have no current parent, otherwise the method fails.
See also [method set_unparent_when_invisible] and [method Node.get_last_exclusive_window].
*/
func (self Instance) PopupExclusiveCenteredRatio(from_node [1]gdclass.Node) {
	class(self).PopupExclusiveCenteredRatio(from_node, gd.Float(0.8))
}

/*
Attempts to parent this dialog to the last exclusive window relative to [param from_node], and then calls [method Window.popup_centered_clamped] on it. The dialog must have no current parent, otherwise the method fails.
See also [method set_unparent_when_invisible] and [method Node.get_last_exclusive_window].
*/
func (self Instance) PopupExclusiveCenteredClamped(from_node [1]gdclass.Node) {
	class(self).PopupExclusiveCenteredClamped(from_node, gd.Vector2i(gd.Vector2i{0, 0}), gd.Float(0.75))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Window

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Window"))
	casted := Instance{*(*gdclass.Window)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Mode() gdclass.WindowMode {
	return gdclass.WindowMode(class(self).GetMode())
}

func (self Instance) SetMode(value gdclass.WindowMode) {
	class(self).SetMode(value)
}

func (self Instance) Title() string {
	return string(class(self).GetTitle().String())
}

func (self Instance) SetTitle(value string) {
	class(self).SetTitle(gd.NewString(value))
}

func (self Instance) InitialPosition() gdclass.WindowWindowInitialPosition {
	return gdclass.WindowWindowInitialPosition(class(self).GetInitialPosition())
}

func (self Instance) SetInitialPosition(value gdclass.WindowWindowInitialPosition) {
	class(self).SetInitialPosition(value)
}

func (self Instance) Position() Vector2i.XY {
	return Vector2i.XY(class(self).GetPosition())
}

func (self Instance) SetPosition(value Vector2i.XY) {
	class(self).SetPosition(gd.Vector2i(value))
}

func (self Instance) Size() Vector2i.XY {
	return Vector2i.XY(class(self).GetSize())
}

func (self Instance) SetSize(value Vector2i.XY) {
	class(self).SetSize(gd.Vector2i(value))
}

func (self Instance) CurrentScreen() int {
	return int(int(class(self).GetCurrentScreen()))
}

func (self Instance) SetCurrentScreen(value int) {
	class(self).SetCurrentScreen(gd.Int(value))
}

func (self Instance) MousePassthroughPolygon() []Vector2.XY {
	return []Vector2.XY(class(self).GetMousePassthroughPolygon().AsSlice())
}

func (self Instance) SetMousePassthroughPolygon(value []Vector2.XY) {
	class(self).SetMousePassthroughPolygon(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&value))))
}

func (self Instance) Visible() bool {
	return bool(class(self).IsVisible())
}

func (self Instance) SetVisible(value bool) {
	class(self).SetVisible(value)
}

func (self Instance) WrapControls() bool {
	return bool(class(self).IsWrappingControls())
}

func (self Instance) SetWrapControls(value bool) {
	class(self).SetWrapControls(value)
}

func (self Instance) Transient() bool {
	return bool(class(self).IsTransient())
}

func (self Instance) SetTransient(value bool) {
	class(self).SetTransient(value)
}

func (self Instance) TransientToFocused() bool {
	return bool(class(self).IsTransientToFocused())
}

func (self Instance) SetTransientToFocused(value bool) {
	class(self).SetTransientToFocused(value)
}

func (self Instance) Exclusive() bool {
	return bool(class(self).IsExclusive())
}

func (self Instance) SetExclusive(value bool) {
	class(self).SetExclusive(value)
}

func (self Instance) Unresizable() bool {
	return bool(class(self).GetFlag(0))
}

func (self Instance) SetUnresizable(value bool) {
	class(self).SetFlag(0, value)
}

func (self Instance) Borderless() bool {
	return bool(class(self).GetFlag(1))
}

func (self Instance) SetBorderless(value bool) {
	class(self).SetFlag(1, value)
}

func (self Instance) AlwaysOnTop() bool {
	return bool(class(self).GetFlag(2))
}

func (self Instance) SetAlwaysOnTop(value bool) {
	class(self).SetFlag(2, value)
}

func (self Instance) Transparent() bool {
	return bool(class(self).GetFlag(3))
}

func (self Instance) SetTransparent(value bool) {
	class(self).SetFlag(3, value)
}

func (self Instance) Unfocusable() bool {
	return bool(class(self).GetFlag(4))
}

func (self Instance) SetUnfocusable(value bool) {
	class(self).SetFlag(4, value)
}

func (self Instance) PopupWindow() bool {
	return bool(class(self).GetFlag(5))
}

func (self Instance) SetPopupWindow(value bool) {
	class(self).SetFlag(5, value)
}

func (self Instance) ExtendToTitle() bool {
	return bool(class(self).GetFlag(6))
}

func (self Instance) SetExtendToTitle(value bool) {
	class(self).SetFlag(6, value)
}

func (self Instance) MousePassthrough() bool {
	return bool(class(self).GetFlag(7))
}

func (self Instance) SetMousePassthrough(value bool) {
	class(self).SetFlag(7, value)
}

func (self Instance) ForceNative() bool {
	return bool(class(self).GetForceNative())
}

func (self Instance) SetForceNative(value bool) {
	class(self).SetForceNative(value)
}

func (self Instance) MinSize() Vector2i.XY {
	return Vector2i.XY(class(self).GetMinSize())
}

func (self Instance) SetMinSize(value Vector2i.XY) {
	class(self).SetMinSize(gd.Vector2i(value))
}

func (self Instance) MaxSize() Vector2i.XY {
	return Vector2i.XY(class(self).GetMaxSize())
}

func (self Instance) SetMaxSize(value Vector2i.XY) {
	class(self).SetMaxSize(gd.Vector2i(value))
}

func (self Instance) KeepTitleVisible() bool {
	return bool(class(self).GetKeepTitleVisible())
}

func (self Instance) SetKeepTitleVisible(value bool) {
	class(self).SetKeepTitleVisible(value)
}

func (self Instance) ContentScaleSize() Vector2i.XY {
	return Vector2i.XY(class(self).GetContentScaleSize())
}

func (self Instance) SetContentScaleSize(value Vector2i.XY) {
	class(self).SetContentScaleSize(gd.Vector2i(value))
}

func (self Instance) ContentScaleMode() gdclass.WindowContentScaleMode {
	return gdclass.WindowContentScaleMode(class(self).GetContentScaleMode())
}

func (self Instance) SetContentScaleMode(value gdclass.WindowContentScaleMode) {
	class(self).SetContentScaleMode(value)
}

func (self Instance) ContentScaleAspect() gdclass.WindowContentScaleAspect {
	return gdclass.WindowContentScaleAspect(class(self).GetContentScaleAspect())
}

func (self Instance) SetContentScaleAspect(value gdclass.WindowContentScaleAspect) {
	class(self).SetContentScaleAspect(value)
}

func (self Instance) ContentScaleStretch() gdclass.WindowContentScaleStretch {
	return gdclass.WindowContentScaleStretch(class(self).GetContentScaleStretch())
}

func (self Instance) SetContentScaleStretch(value gdclass.WindowContentScaleStretch) {
	class(self).SetContentScaleStretch(value)
}

func (self Instance) ContentScaleFactor() Float.X {
	return Float.X(Float.X(class(self).GetContentScaleFactor()))
}

func (self Instance) SetContentScaleFactor(value Float.X) {
	class(self).SetContentScaleFactor(gd.Float(value))
}

func (self Instance) AutoTranslate() bool {
	return bool(class(self).IsAutoTranslating())
}

func (self Instance) SetAutoTranslate(value bool) {
	class(self).SetAutoTranslate(value)
}

func (self Instance) Theme() [1]gdclass.Theme {
	return [1]gdclass.Theme(class(self).GetTheme())
}

func (self Instance) SetTheme(value [1]gdclass.Theme) {
	class(self).SetTheme(value)
}

func (self Instance) ThemeTypeVariation() string {
	return string(class(self).GetThemeTypeVariation().String())
}

func (self Instance) SetThemeTypeVariation(value string) {
	class(self).SetThemeTypeVariation(gd.NewStringName(value))
}

/*
Virtual method to be implemented by the user. Overrides the value returned by [method get_contents_minimum_size].
*/
func (class) _get_contents_minimum_size(impl func(ptr unsafe.Pointer) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

//go:nosplit
func (self class) SetTitle(title gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(title))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_title, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTitle() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_title, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the ID of the window.
*/
//go:nosplit
func (self class) GetWindowId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_window_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInitialPosition(initial_position gdclass.WindowWindowInitialPosition) {
	var frame = callframe.New()
	callframe.Arg(frame, initial_position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_initial_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInitialPosition() gdclass.WindowWindowInitialPosition {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.WindowWindowInitialPosition](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_initial_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCurrentScreen(index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_current_screen, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCurrentScreen() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_current_screen, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPosition(position gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPosition() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Centers a native window on the current screen and an embedded window on its embedder [Viewport].
*/
//go:nosplit
func (self class) MoveToCenter() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_move_to_center, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetSize(size gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSize() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Resets the size to the minimum size, which is the max of [member min_size] and (if [member wrap_controls] is enabled) [method get_contents_minimum_size]. This is equivalent to calling [code]set_size(Vector2i())[/code] (or any size below the minimum).
*/
//go:nosplit
func (self class) ResetSize() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_reset_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the window's position including its border.
[b]Note:[/b] If [member visible] is [code]false[/code], this method returns the same value as [member position].
*/
//go:nosplit
func (self class) GetPositionWithDecorations() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_position_with_decorations, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_size_with_decorations, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxSize(max_size gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, max_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_max_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxSize() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_max_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinSize(min_size gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, min_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_min_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinSize() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_min_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMode(mode gdclass.WindowMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMode() gdclass.WindowMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.WindowMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets a specified window flag.
*/
//go:nosplit
func (self class) SetFlag(flag gdclass.WindowFlags, enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_flag, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the [param flag] is set.
*/
//go:nosplit
func (self class) GetFlag(flag gdclass.WindowFlags) bool {
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_flag, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the window can be maximized (the maximize button is enabled).
*/
//go:nosplit
func (self class) IsMaximizeAllowed() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_is_maximize_allowed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Tells the OS that the [Window] needs an attention. This makes the window stand out in some way depending on the system, e.g. it might blink on the task bar.
*/
//go:nosplit
func (self class) RequestAttention() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_request_attention, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Causes the window to grab focus, allowing it to receive user input.
*/
//go:nosplit
func (self class) MoveToForeground() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_move_to_foreground, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetVisible(visible bool) {
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsVisible() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_is_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Hides the window. This is not the same as minimized state. Hidden window can't be interacted with and needs to be made visible with [method show].
*/
//go:nosplit
func (self class) Hide() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_hide, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Makes the [Window] appear. This enables interactions with the [Window] and doesn't change any of its property other than visibility (unlike e.g. [method popup]).
*/
//go:nosplit
func (self class) Show() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_show, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetTransient(transient bool) {
	var frame = callframe.New()
	callframe.Arg(frame, transient)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_transient, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsTransient() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_is_transient, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTransientToFocused(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_transient_to_focused, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsTransientToFocused() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_is_transient_to_focused, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExclusive(exclusive bool) {
	var frame = callframe.New()
	callframe.Arg(frame, exclusive)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_exclusive, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsExclusive() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_is_exclusive, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param unparent] is [code]true[/code], the window is automatically unparented when going invisible.
[b]Note:[/b] Make sure to keep a reference to the node, otherwise it will be orphaned. You also need to manually call [method Node.queue_free] to free the window if it's not parented.
*/
//go:nosplit
func (self class) SetUnparentWhenInvisible(unparent bool) {
	var frame = callframe.New()
	callframe.Arg(frame, unparent)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_unparent_when_invisible, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether the window is being drawn to the screen.
*/
//go:nosplit
func (self class) CanDraw() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_can_draw, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the window is focused.
*/
//go:nosplit
func (self class) HasFocus() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_has_focus, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Causes the window to grab focus, allowing it to receive user input.
*/
//go:nosplit
func (self class) GrabFocus() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_grab_focus, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
If [param active] is [code]true[/code], enables system's native IME (Input Method Editor).
*/
//go:nosplit
func (self class) SetImeActive(active bool) {
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_ime_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Moves IME to the given position.
*/
//go:nosplit
func (self class) SetImePosition(position gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_ime_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the window is currently embedded in another window.
*/
//go:nosplit
func (self class) IsEmbedded() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_is_embedded, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_contents_minimum_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetForceNative(force_native bool) {
	var frame = callframe.New()
	callframe.Arg(frame, force_native)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_force_native, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetForceNative() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_force_native, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetContentScaleSize(size gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_content_scale_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetContentScaleSize() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_content_scale_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetContentScaleMode(mode gdclass.WindowContentScaleMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_content_scale_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetContentScaleMode() gdclass.WindowContentScaleMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.WindowContentScaleMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_content_scale_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetContentScaleAspect(aspect gdclass.WindowContentScaleAspect) {
	var frame = callframe.New()
	callframe.Arg(frame, aspect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_content_scale_aspect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetContentScaleAspect() gdclass.WindowContentScaleAspect {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.WindowContentScaleAspect](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_content_scale_aspect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetContentScaleStretch(stretch gdclass.WindowContentScaleStretch) {
	var frame = callframe.New()
	callframe.Arg(frame, stretch)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_content_scale_stretch, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetContentScaleStretch() gdclass.WindowContentScaleStretch {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.WindowContentScaleStretch](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_content_scale_stretch, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetKeepTitleVisible(title_visible bool) {
	var frame = callframe.New()
	callframe.Arg(frame, title_visible)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_keep_title_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetKeepTitleVisible() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_keep_title_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetContentScaleFactor(factor gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, factor)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_content_scale_factor, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetContentScaleFactor() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_content_scale_factor, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Enables font oversampling. This makes fonts look better when they are scaled up.
*/
//go:nosplit
func (self class) SetUseFontOversampling(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_use_font_oversampling, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if font oversampling is enabled. See [method set_use_font_oversampling].
*/
//go:nosplit
func (self class) IsUsingFontOversampling() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_is_using_font_oversampling, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMousePassthroughPolygon(polygon gd.PackedVector2Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(polygon))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_mouse_passthrough_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMousePassthroughPolygon() gd.PackedVector2Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_mouse_passthrough_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWrapControls(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_wrap_controls, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsWrappingControls() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_is_wrapping_controls, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Requests an update of the [Window] size to fit underlying [Control] nodes.
*/
//go:nosplit
func (self class) ChildControlsChanged() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_child_controls_changed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetTheme(theme [1]gdclass.Theme) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(theme[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_theme, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTheme() [1]gdclass.Theme {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_theme, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Theme{gd.PointerWithOwnershipTransferredToGo[gdclass.Theme](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetThemeTypeVariation(theme_type gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_theme_type_variation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetThemeTypeVariation() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_theme_type_variation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Prevents [code]*_theme_*_override[/code] methods from emitting [constant NOTIFICATION_THEME_CHANGED] until [method end_bulk_theme_override] is called.
*/
//go:nosplit
func (self class) BeginBulkThemeOverride() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_begin_bulk_theme_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Ends a bulk theme override update. See [method begin_bulk_theme_override].
*/
//go:nosplit
func (self class) EndBulkThemeOverride() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_end_bulk_theme_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Creates a local override for a theme icon with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_icon_override].
See also [method get_theme_icon].
*/
//go:nosplit
func (self class) AddThemeIconOverride(name gd.StringName, texture [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_add_theme_icon_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Creates a local override for a theme [StyleBox] with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_stylebox_override].
See also [method get_theme_stylebox] and [method Control.add_theme_stylebox_override] for more details.
*/
//go:nosplit
func (self class) AddThemeStyleboxOverride(name gd.StringName, stylebox [1]gdclass.StyleBox) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(stylebox[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_add_theme_stylebox_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Creates a local override for a theme [Font] with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_font_override].
See also [method get_theme_font].
*/
//go:nosplit
func (self class) AddThemeFontOverride(name gd.StringName, font [1]gdclass.Font) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(font[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_add_theme_font_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Creates a local override for a theme font size with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_font_size_override].
See also [method get_theme_font_size].
*/
//go:nosplit
func (self class) AddThemeFontSizeOverride(name gd.StringName, font_size gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, font_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_add_theme_font_size_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Creates a local override for a theme [Color] with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_color_override].
See also [method get_theme_color] and [method Control.add_theme_color_override] for more details.
*/
//go:nosplit
func (self class) AddThemeColorOverride(name gd.StringName, color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_add_theme_color_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Creates a local override for a theme constant with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_constant_override].
See also [method get_theme_constant].
*/
//go:nosplit
func (self class) AddThemeConstantOverride(name gd.StringName, constant gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, constant)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_add_theme_constant_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a local override for a theme icon with the specified [param name] previously added by [method add_theme_icon_override] or via the Inspector dock.
*/
//go:nosplit
func (self class) RemoveThemeIconOverride(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_remove_theme_icon_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a local override for a theme [StyleBox] with the specified [param name] previously added by [method add_theme_stylebox_override] or via the Inspector dock.
*/
//go:nosplit
func (self class) RemoveThemeStyleboxOverride(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_remove_theme_stylebox_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a local override for a theme [Font] with the specified [param name] previously added by [method add_theme_font_override] or via the Inspector dock.
*/
//go:nosplit
func (self class) RemoveThemeFontOverride(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_remove_theme_font_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a local override for a theme font size with the specified [param name] previously added by [method add_theme_font_size_override] or via the Inspector dock.
*/
//go:nosplit
func (self class) RemoveThemeFontSizeOverride(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_remove_theme_font_size_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a local override for a theme [Color] with the specified [param name] previously added by [method add_theme_color_override] or via the Inspector dock.
*/
//go:nosplit
func (self class) RemoveThemeColorOverride(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_remove_theme_color_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a local override for a theme constant with the specified [param name] previously added by [method add_theme_constant_override] or via the Inspector dock.
*/
//go:nosplit
func (self class) RemoveThemeConstantOverride(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_remove_theme_constant_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an icon from the first matching [Theme] in the tree if that [Theme] has an icon item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
//go:nosplit
func (self class) GetThemeIcon(name gd.StringName, theme_type gd.StringName) [1]gdclass.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_theme_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns a [StyleBox] from the first matching [Theme] in the tree if that [Theme] has a stylebox item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
//go:nosplit
func (self class) GetThemeStylebox(name gd.StringName, theme_type gd.StringName) [1]gdclass.StyleBox {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_theme_stylebox, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.StyleBox{gd.PointerWithOwnershipTransferredToGo[gdclass.StyleBox](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns a [Font] from the first matching [Theme] in the tree if that [Theme] has a font item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
//go:nosplit
func (self class) GetThemeFont(name gd.StringName, theme_type gd.StringName) [1]gdclass.Font {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_theme_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Font{gd.PointerWithOwnershipTransferredToGo[gdclass.Font](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns a font size from the first matching [Theme] in the tree if that [Theme] has a font size item with the specified [param name] and [param theme_type].
See [method Control.get_theme_color] for details.
*/
//go:nosplit
func (self class) GetThemeFontSize(name gd.StringName, theme_type gd.StringName) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_theme_font_size, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_theme_color, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_theme_constant, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_has_theme_icon_override, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_has_theme_stylebox_override, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_has_theme_font_override, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_has_theme_font_size_override, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_has_theme_color_override, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_has_theme_constant_override, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_has_theme_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_has_theme_stylebox, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_has_theme_font, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_has_theme_font_size, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_has_theme_color, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_has_theme_constant, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_theme_default_base_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the default font from the first matching [Theme] in the tree if that [Theme] has a valid [member Theme.default_font] value.
See [method Control.get_theme_color] for details.
*/
//go:nosplit
func (self class) GetThemeDefaultFont() [1]gdclass.Font {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_theme_default_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Font{gd.PointerWithOwnershipTransferredToGo[gdclass.Font](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the default font size value from the first matching [Theme] in the tree if that [Theme] has a valid [member Theme.default_font_size] value.
See [method Control.get_theme_color] for details.
*/
//go:nosplit
func (self class) GetThemeDefaultFontSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_theme_default_font_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets layout direction and text writing direction. Right-to-left layouts are necessary for certain languages (e.g. Arabic and Hebrew).
*/
//go:nosplit
func (self class) SetLayoutDirection(direction gdclass.WindowLayoutDirection) {
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_layout_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns layout direction and text writing direction.
*/
//go:nosplit
func (self class) GetLayoutDirection() gdclass.WindowLayoutDirection {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.WindowLayoutDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_get_layout_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if layout is right-to-left.
*/
//go:nosplit
func (self class) IsLayoutRtl() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_is_layout_rtl, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoTranslate(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_set_auto_translate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsAutoTranslating() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_is_auto_translating, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) Popup(rect gd.Rect2i) {
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_popup, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Popups the [Window] with a position shifted by parent [Window]'s position. If the [Window] is embedded, has the same effect as [method popup].
*/
//go:nosplit
func (self class) PopupOnParent(parent_rect gd.Rect2i) {
	var frame = callframe.New()
	callframe.Arg(frame, parent_rect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_popup_on_parent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Popups the [Window] at the center of the current screen, with optionally given minimum size. If the [Window] is embedded, it will be centered in the parent [Viewport] instead.
[b]Note:[/b] Calling it with the default value of [param minsize] is equivalent to calling it with [member size].
*/
//go:nosplit
func (self class) PopupCentered(minsize gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, minsize)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_popup_centered, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
If [Window] is embedded, popups the [Window] centered inside its embedder and sets its size as a [param ratio] of embedder's size.
If [Window] is a native window, popups the [Window] centered inside the screen of its parent [Window] and sets its size as a [param ratio] of the screen size.
*/
//go:nosplit
func (self class) PopupCenteredRatio(ratio gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_popup_centered_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Popups the [Window] centered inside its parent [Window]. [param fallback_ratio] determines the maximum size of the [Window], in relation to its parent.
[b]Note:[/b] Calling it with the default value of [param minsize] is equivalent to calling it with [member size].
*/
//go:nosplit
func (self class) PopupCenteredClamped(minsize gd.Vector2i, fallback_ratio gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, minsize)
	callframe.Arg(frame, fallback_ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_popup_centered_clamped, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Attempts to parent this dialog to the last exclusive window relative to [param from_node], and then calls [method Window.popup] on it. The dialog must have no current parent, otherwise the method fails.
See also [method set_unparent_when_invisible] and [method Node.get_last_exclusive_window].
*/
//go:nosplit
func (self class) PopupExclusive(from_node [1]gdclass.Node, rect gd.Rect2i) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(from_node[0])[0])
	callframe.Arg(frame, rect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_popup_exclusive, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Attempts to parent this dialog to the last exclusive window relative to [param from_node], and then calls [method Window.popup_on_parent] on it. The dialog must have no current parent, otherwise the method fails.
See also [method set_unparent_when_invisible] and [method Node.get_last_exclusive_window].
*/
//go:nosplit
func (self class) PopupExclusiveOnParent(from_node [1]gdclass.Node, parent_rect gd.Rect2i) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(from_node[0])[0])
	callframe.Arg(frame, parent_rect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_popup_exclusive_on_parent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Attempts to parent this dialog to the last exclusive window relative to [param from_node], and then calls [method Window.popup_centered] on it. The dialog must have no current parent, otherwise the method fails.
See also [method set_unparent_when_invisible] and [method Node.get_last_exclusive_window].
*/
//go:nosplit
func (self class) PopupExclusiveCentered(from_node [1]gdclass.Node, minsize gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(from_node[0])[0])
	callframe.Arg(frame, minsize)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_popup_exclusive_centered, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Attempts to parent this dialog to the last exclusive window relative to [param from_node], and then calls [method Window.popup_centered_ratio] on it. The dialog must have no current parent, otherwise the method fails.
See also [method set_unparent_when_invisible] and [method Node.get_last_exclusive_window].
*/
//go:nosplit
func (self class) PopupExclusiveCenteredRatio(from_node [1]gdclass.Node, ratio gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(from_node[0])[0])
	callframe.Arg(frame, ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_popup_exclusive_centered_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Attempts to parent this dialog to the last exclusive window relative to [param from_node], and then calls [method Window.popup_centered_clamped] on it. The dialog must have no current parent, otherwise the method fails.
See also [method set_unparent_when_invisible] and [method Node.get_last_exclusive_window].
*/
//go:nosplit
func (self class) PopupExclusiveCenteredClamped(from_node [1]gdclass.Node, minsize gd.Vector2i, fallback_ratio gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(from_node[0])[0])
	callframe.Arg(frame, minsize)
	callframe.Arg(frame, fallback_ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Window.Bind_popup_exclusive_centered_clamped, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnWindowInput(cb func(event [1]gdclass.InputEvent)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("window_input"), gd.NewCallable(cb), 0)
}

func (self Instance) OnFilesDropped(cb func(files []string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("files_dropped"), gd.NewCallable(cb), 0)
}

func (self Instance) OnMouseEntered(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("mouse_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnMouseExited(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("mouse_exited"), gd.NewCallable(cb), 0)
}

func (self Instance) OnFocusEntered(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("focus_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnFocusExited(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("focus_exited"), gd.NewCallable(cb), 0)
}

func (self Instance) OnCloseRequested(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("close_requested"), gd.NewCallable(cb), 0)
}

func (self Instance) OnGoBackRequested(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("go_back_requested"), gd.NewCallable(cb), 0)
}

func (self Instance) OnVisibilityChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("visibility_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAboutToPopup(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("about_to_popup"), gd.NewCallable(cb), 0)
}

func (self Instance) OnThemeChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("theme_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnDpiChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("dpi_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTitlebarChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("titlebar_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsWindow() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsWindow() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsViewport() Viewport.Advanced {
	return *((*Viewport.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsViewport() Viewport.Instance {
	return *((*Viewport.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_contents_minimum_size":
		return reflect.ValueOf(self._get_contents_minimum_size)
	default:
		return gd.VirtualByName(Viewport.Advanced(self.AsViewport()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_contents_minimum_size":
		return reflect.ValueOf(self._get_contents_minimum_size)
	default:
		return gd.VirtualByName(Viewport.Instance(self.AsViewport()), name)
	}
}
func init() {
	gdclass.Register("Window", func(ptr gd.Object) any { return [1]gdclass.Window{*(*gdclass.Window)(unsafe.Pointer(&ptr))} })
}

type Mode = gdclass.WindowMode

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

type Flags = gdclass.WindowFlags

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

type ContentScaleMode = gdclass.WindowContentScaleMode

const (
	/*The content will not be scaled to match the [Window]'s size.*/
	ContentScaleModeDisabled ContentScaleMode = 0
	/*The content will be rendered at the target size. This is more performance-expensive than [constant CONTENT_SCALE_MODE_VIEWPORT], but provides better results.*/
	ContentScaleModeCanvasItems ContentScaleMode = 1
	/*The content will be rendered at the base size and then scaled to the target size. More performant than [constant CONTENT_SCALE_MODE_CANVAS_ITEMS], but results in pixelated image.*/
	ContentScaleModeViewport ContentScaleMode = 2
)

type ContentScaleAspect = gdclass.WindowContentScaleAspect

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

type ContentScaleStretch = gdclass.WindowContentScaleStretch

const (
	/*The content will be stretched according to a fractional factor. This fills all the space available in the window, but allows "pixel wobble" to occur due to uneven pixel scaling.*/
	ContentScaleStretchFractional ContentScaleStretch = 0
	/*The content will be stretched only according to an integer factor, preserving sharp pixels. This may leave a black background visible on the window's edges depending on the window size.*/
	ContentScaleStretchInteger ContentScaleStretch = 1
)

type LayoutDirection = gdclass.WindowLayoutDirection

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

type WindowInitialPosition = gdclass.WindowWindowInitialPosition

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
