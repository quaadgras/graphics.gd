package TabBar

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Control"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Rect2"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A control that provides a horizontal bar with tabs. Similar to [TabContainer] but is only in charge of drawing tabs, not interacting with children.
*/
type Instance [1]classdb.TabBar
type Any interface {
	gd.IsClass
	AsTabBar() Instance
}

/*
Returns the previously active tab index.
*/
func (self Instance) GetPreviousTab() int {
	return int(int(class(self).GetPreviousTab()))
}

/*
Selects the first available tab with lower index than the currently selected. Returns [code]true[/code] if tab selection changed.
*/
func (self Instance) SelectPreviousAvailable() bool {
	return bool(class(self).SelectPreviousAvailable())
}

/*
Selects the first available tab with greater index than the currently selected. Returns [code]true[/code] if tab selection changed.
*/
func (self Instance) SelectNextAvailable() bool {
	return bool(class(self).SelectNextAvailable())
}

/*
Sets a [param title] for the tab at index [param tab_idx].
*/
func (self Instance) SetTabTitle(tab_idx int, title string) {
	class(self).SetTabTitle(gd.Int(tab_idx), gd.NewString(title))
}

/*
Returns the title of the tab at index [param tab_idx].
*/
func (self Instance) GetTabTitle(tab_idx int) string {
	return string(class(self).GetTabTitle(gd.Int(tab_idx)).String())
}

/*
Sets a [param tooltip] for tab at index [param tab_idx].
[b]Note:[/b] By default, if the [param tooltip] is empty and the tab text is truncated (not all characters fit into the tab), the title will be displayed as a tooltip. To hide the tooltip, assign [code]" "[/code] as the [param tooltip] text.
*/
func (self Instance) SetTabTooltip(tab_idx int, tooltip string) {
	class(self).SetTabTooltip(gd.Int(tab_idx), gd.NewString(tooltip))
}

/*
Returns the tooltip text of the tab at index [param tab_idx].
*/
func (self Instance) GetTabTooltip(tab_idx int) string {
	return string(class(self).GetTabTooltip(gd.Int(tab_idx)).String())
}

/*
Sets tab title base writing direction.
*/
func (self Instance) SetTabTextDirection(tab_idx int, direction classdb.ControlTextDirection) {
	class(self).SetTabTextDirection(gd.Int(tab_idx), direction)
}

/*
Returns tab title text base writing direction.
*/
func (self Instance) GetTabTextDirection(tab_idx int) classdb.ControlTextDirection {
	return classdb.ControlTextDirection(class(self).GetTabTextDirection(gd.Int(tab_idx)))
}

/*
Sets language code of tab title used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
func (self Instance) SetTabLanguage(tab_idx int, language string) {
	class(self).SetTabLanguage(gd.Int(tab_idx), gd.NewString(language))
}

/*
Returns tab title language code.
*/
func (self Instance) GetTabLanguage(tab_idx int) string {
	return string(class(self).GetTabLanguage(gd.Int(tab_idx)).String())
}

/*
Sets an [param icon] for the tab at index [param tab_idx].
*/
func (self Instance) SetTabIcon(tab_idx int, icon objects.Texture2D) {
	class(self).SetTabIcon(gd.Int(tab_idx), icon)
}

/*
Returns the icon for the tab at index [param tab_idx] or [code]null[/code] if the tab has no icon.
*/
func (self Instance) GetTabIcon(tab_idx int) objects.Texture2D {
	return objects.Texture2D(class(self).GetTabIcon(gd.Int(tab_idx)))
}

/*
Sets the maximum allowed width of the icon for the tab at index [param tab_idx]. This limit is applied on top of the default size of the icon and on top of [theme_item icon_max_width]. The height is adjusted according to the icon's ratio.
*/
func (self Instance) SetTabIconMaxWidth(tab_idx int, width int) {
	class(self).SetTabIconMaxWidth(gd.Int(tab_idx), gd.Int(width))
}

/*
Returns the maximum allowed width of the icon for the tab at index [param tab_idx].
*/
func (self Instance) GetTabIconMaxWidth(tab_idx int) int {
	return int(int(class(self).GetTabIconMaxWidth(gd.Int(tab_idx))))
}

/*
Sets an [param icon] for the button of the tab at index [param tab_idx] (located to the right, before the close button), making it visible and clickable (See [signal tab_button_pressed]). Giving it a [code]null[/code] value will hide the button.
*/
func (self Instance) SetTabButtonIcon(tab_idx int, icon objects.Texture2D) {
	class(self).SetTabButtonIcon(gd.Int(tab_idx), icon)
}

/*
Returns the icon for the right button of the tab at index [param tab_idx] or [code]null[/code] if the right button has no icon.
*/
func (self Instance) GetTabButtonIcon(tab_idx int) objects.Texture2D {
	return objects.Texture2D(class(self).GetTabButtonIcon(gd.Int(tab_idx)))
}

/*
If [param disabled] is [code]true[/code], disables the tab at index [param tab_idx], making it non-interactable.
*/
func (self Instance) SetTabDisabled(tab_idx int, disabled bool) {
	class(self).SetTabDisabled(gd.Int(tab_idx), disabled)
}

/*
Returns [code]true[/code] if the tab at index [param tab_idx] is disabled.
*/
func (self Instance) IsTabDisabled(tab_idx int) bool {
	return bool(class(self).IsTabDisabled(gd.Int(tab_idx)))
}

/*
If [param hidden] is [code]true[/code], hides the tab at index [param tab_idx], making it disappear from the tab area.
*/
func (self Instance) SetTabHidden(tab_idx int, hidden bool) {
	class(self).SetTabHidden(gd.Int(tab_idx), hidden)
}

/*
Returns [code]true[/code] if the tab at index [param tab_idx] is hidden.
*/
func (self Instance) IsTabHidden(tab_idx int) bool {
	return bool(class(self).IsTabHidden(gd.Int(tab_idx)))
}

/*
Sets the metadata value for the tab at index [param tab_idx], which can be retrieved later using [method get_tab_metadata].
*/
func (self Instance) SetTabMetadata(tab_idx int, metadata any) {
	class(self).SetTabMetadata(gd.Int(tab_idx), gd.NewVariant(metadata))
}

/*
Returns the metadata value set to the tab at index [param tab_idx] using [method set_tab_metadata]. If no metadata was previously set, returns [code]null[/code] by default.
*/
func (self Instance) GetTabMetadata(tab_idx int) any {
	return any(class(self).GetTabMetadata(gd.Int(tab_idx)).Interface())
}

/*
Removes the tab at index [param tab_idx].
*/
func (self Instance) RemoveTab(tab_idx int) {
	class(self).RemoveTab(gd.Int(tab_idx))
}

/*
Adds a new tab.
*/
func (self Instance) AddTab() {
	class(self).AddTab(gd.NewString(""), [1]objects.Texture2D{}[0])
}

/*
Returns the index of the tab at local coordinates [param point]. Returns [code]-1[/code] if the point is outside the control boundaries or if there's no tab at the queried position.
*/
func (self Instance) GetTabIdxAtPoint(point Vector2.XY) int {
	return int(int(class(self).GetTabIdxAtPoint(gd.Vector2(point))))
}

/*
Returns the number of hidden tabs offsetted to the left.
*/
func (self Instance) GetTabOffset() int {
	return int(int(class(self).GetTabOffset()))
}

/*
Returns [code]true[/code] if the offset buttons (the ones that appear when there's not enough space for all tabs) are visible.
*/
func (self Instance) GetOffsetButtonsVisible() bool {
	return bool(class(self).GetOffsetButtonsVisible())
}

/*
Moves the scroll view to make the tab visible.
*/
func (self Instance) EnsureTabVisible(idx int) {
	class(self).EnsureTabVisible(gd.Int(idx))
}

/*
Returns tab [Rect2] with local position and size.
*/
func (self Instance) GetTabRect(tab_idx int) Rect2.PositionSize {
	return Rect2.PositionSize(class(self).GetTabRect(gd.Int(tab_idx)))
}

/*
Moves a tab from [param from] to [param to].
*/
func (self Instance) MoveTab(from int, to int) {
	class(self).MoveTab(gd.Int(from), gd.Int(to))
}

/*
Clears all tabs.
*/
func (self Instance) ClearTabs() {
	class(self).ClearTabs()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.TabBar

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TabBar"))
	return Instance{classdb.TabBar(object)}
}

func (self Instance) CurrentTab() int {
	return int(int(class(self).GetCurrentTab()))
}

func (self Instance) SetCurrentTab(value int) {
	class(self).SetCurrentTab(gd.Int(value))
}

func (self Instance) TabAlignment() classdb.TabBarAlignmentMode {
	return classdb.TabBarAlignmentMode(class(self).GetTabAlignment())
}

func (self Instance) SetTabAlignment(value classdb.TabBarAlignmentMode) {
	class(self).SetTabAlignment(value)
}

func (self Instance) ClipTabs() bool {
	return bool(class(self).GetClipTabs())
}

func (self Instance) SetClipTabs(value bool) {
	class(self).SetClipTabs(value)
}

func (self Instance) TabCloseDisplayPolicy() classdb.TabBarCloseButtonDisplayPolicy {
	return classdb.TabBarCloseButtonDisplayPolicy(class(self).GetTabCloseDisplayPolicy())
}

func (self Instance) SetTabCloseDisplayPolicy(value classdb.TabBarCloseButtonDisplayPolicy) {
	class(self).SetTabCloseDisplayPolicy(value)
}

func (self Instance) MaxTabWidth() int {
	return int(int(class(self).GetMaxTabWidth()))
}

func (self Instance) SetMaxTabWidth(value int) {
	class(self).SetMaxTabWidth(gd.Int(value))
}

func (self Instance) ScrollingEnabled() bool {
	return bool(class(self).GetScrollingEnabled())
}

func (self Instance) SetScrollingEnabled(value bool) {
	class(self).SetScrollingEnabled(value)
}

func (self Instance) DragToRearrangeEnabled() bool {
	return bool(class(self).GetDragToRearrangeEnabled())
}

func (self Instance) SetDragToRearrangeEnabled(value bool) {
	class(self).SetDragToRearrangeEnabled(value)
}

func (self Instance) TabsRearrangeGroup() int {
	return int(int(class(self).GetTabsRearrangeGroup()))
}

func (self Instance) SetTabsRearrangeGroup(value int) {
	class(self).SetTabsRearrangeGroup(gd.Int(value))
}

func (self Instance) ScrollToSelected() bool {
	return bool(class(self).GetScrollToSelected())
}

func (self Instance) SetScrollToSelected(value bool) {
	class(self).SetScrollToSelected(value)
}

func (self Instance) SelectWithRmb() bool {
	return bool(class(self).GetSelectWithRmb())
}

func (self Instance) SetSelectWithRmb(value bool) {
	class(self).SetSelectWithRmb(value)
}

func (self Instance) DeselectEnabled() bool {
	return bool(class(self).GetDeselectEnabled())
}

func (self Instance) SetDeselectEnabled(value bool) {
	class(self).SetDeselectEnabled(value)
}

func (self Instance) TabCount() int {
	return int(int(class(self).GetTabCount()))
}

func (self Instance) SetTabCount(value int) {
	class(self).SetTabCount(gd.Int(value))
}

//go:nosplit
func (self class) SetTabCount(count gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_tab_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTabCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_tab_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCurrentTab(tab_idx gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_current_tab, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCurrentTab() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_current_tab, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the previously active tab index.
*/
//go:nosplit
func (self class) GetPreviousTab() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_previous_tab, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Selects the first available tab with lower index than the currently selected. Returns [code]true[/code] if tab selection changed.
*/
//go:nosplit
func (self class) SelectPreviousAvailable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_select_previous_available, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Selects the first available tab with greater index than the currently selected. Returns [code]true[/code] if tab selection changed.
*/
//go:nosplit
func (self class) SelectNextAvailable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_select_next_available, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets a [param title] for the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) SetTabTitle(tab_idx gd.Int, title gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, pointers.Get(title))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_tab_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the title of the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) GetTabTitle(tab_idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_tab_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets a [param tooltip] for tab at index [param tab_idx].
[b]Note:[/b] By default, if the [param tooltip] is empty and the tab text is truncated (not all characters fit into the tab), the title will be displayed as a tooltip. To hide the tooltip, assign [code]" "[/code] as the [param tooltip] text.
*/
//go:nosplit
func (self class) SetTabTooltip(tab_idx gd.Int, tooltip gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, pointers.Get(tooltip))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_tab_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the tooltip text of the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) GetTabTooltip(tab_idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_tab_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets tab title base writing direction.
*/
//go:nosplit
func (self class) SetTabTextDirection(tab_idx gd.Int, direction classdb.ControlTextDirection) {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_tab_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns tab title text base writing direction.
*/
//go:nosplit
func (self class) GetTabTextDirection(tab_idx gd.Int) classdb.ControlTextDirection {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[classdb.ControlTextDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_tab_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets language code of tab title used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
//go:nosplit
func (self class) SetTabLanguage(tab_idx gd.Int, language gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, pointers.Get(language))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_tab_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns tab title language code.
*/
//go:nosplit
func (self class) GetTabLanguage(tab_idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_tab_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets an [param icon] for the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) SetTabIcon(tab_idx gd.Int, icon objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_tab_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the icon for the tab at index [param tab_idx] or [code]null[/code] if the tab has no icon.
*/
//go:nosplit
func (self class) GetTabIcon(tab_idx gd.Int) objects.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_tab_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Sets the maximum allowed width of the icon for the tab at index [param tab_idx]. This limit is applied on top of the default size of the icon and on top of [theme_item icon_max_width]. The height is adjusted according to the icon's ratio.
*/
//go:nosplit
func (self class) SetTabIconMaxWidth(tab_idx gd.Int, width gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_tab_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the maximum allowed width of the icon for the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) GetTabIconMaxWidth(tab_idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_tab_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets an [param icon] for the button of the tab at index [param tab_idx] (located to the right, before the close button), making it visible and clickable (See [signal tab_button_pressed]). Giving it a [code]null[/code] value will hide the button.
*/
//go:nosplit
func (self class) SetTabButtonIcon(tab_idx gd.Int, icon objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_tab_button_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the icon for the right button of the tab at index [param tab_idx] or [code]null[/code] if the right button has no icon.
*/
//go:nosplit
func (self class) GetTabButtonIcon(tab_idx gd.Int) objects.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_tab_button_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
If [param disabled] is [code]true[/code], disables the tab at index [param tab_idx], making it non-interactable.
*/
//go:nosplit
func (self class) SetTabDisabled(tab_idx gd.Int, disabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_tab_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the tab at index [param tab_idx] is disabled.
*/
//go:nosplit
func (self class) IsTabDisabled(tab_idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_is_tab_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param hidden] is [code]true[/code], hides the tab at index [param tab_idx], making it disappear from the tab area.
*/
//go:nosplit
func (self class) SetTabHidden(tab_idx gd.Int, hidden bool) {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, hidden)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_tab_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the tab at index [param tab_idx] is hidden.
*/
//go:nosplit
func (self class) IsTabHidden(tab_idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_is_tab_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the metadata value for the tab at index [param tab_idx], which can be retrieved later using [method get_tab_metadata].
*/
//go:nosplit
func (self class) SetTabMetadata(tab_idx gd.Int, metadata gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, pointers.Get(metadata))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_tab_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the metadata value set to the tab at index [param tab_idx] using [method set_tab_metadata]. If no metadata was previously set, returns [code]null[/code] by default.
*/
//go:nosplit
func (self class) GetTabMetadata(tab_idx gd.Int) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_tab_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Removes the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) RemoveTab(tab_idx gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_remove_tab, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a new tab.
*/
//go:nosplit
func (self class) AddTab(title gd.String, icon objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(title))
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_add_tab, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the index of the tab at local coordinates [param point]. Returns [code]-1[/code] if the point is outside the control boundaries or if there's no tab at the queried position.
*/
//go:nosplit
func (self class) GetTabIdxAtPoint(point gd.Vector2) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_tab_idx_at_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTabAlignment(alignment classdb.TabBarAlignmentMode) {
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_tab_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTabAlignment() classdb.TabBarAlignmentMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TabBarAlignmentMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_tab_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetClipTabs(clip_tabs bool) {
	var frame = callframe.New()
	callframe.Arg(frame, clip_tabs)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_clip_tabs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetClipTabs() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_clip_tabs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of hidden tabs offsetted to the left.
*/
//go:nosplit
func (self class) GetTabOffset() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_tab_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the offset buttons (the ones that appear when there's not enough space for all tabs) are visible.
*/
//go:nosplit
func (self class) GetOffsetButtonsVisible() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_offset_buttons_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Moves the scroll view to make the tab visible.
*/
//go:nosplit
func (self class) EnsureTabVisible(idx gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_ensure_tab_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns tab [Rect2] with local position and size.
*/
//go:nosplit
func (self class) GetTabRect(tab_idx gd.Int) gd.Rect2 {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_tab_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Moves a tab from [param from] to [param to].
*/
//go:nosplit
func (self class) MoveTab(from gd.Int, to gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_move_tab, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetTabCloseDisplayPolicy(policy classdb.TabBarCloseButtonDisplayPolicy) {
	var frame = callframe.New()
	callframe.Arg(frame, policy)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_tab_close_display_policy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTabCloseDisplayPolicy() classdb.TabBarCloseButtonDisplayPolicy {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TabBarCloseButtonDisplayPolicy](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_tab_close_display_policy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxTabWidth(width gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_max_tab_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxTabWidth() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_max_tab_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScrollingEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_scrolling_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetScrollingEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_scrolling_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDragToRearrangeEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_drag_to_rearrange_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDragToRearrangeEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_drag_to_rearrange_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTabsRearrangeGroup(group_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, group_id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_tabs_rearrange_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTabsRearrangeGroup() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_tabs_rearrange_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScrollToSelected(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_scroll_to_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetScrollToSelected() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_scroll_to_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSelectWithRmb(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_select_with_rmb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSelectWithRmb() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_select_with_rmb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDeselectEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_set_deselect_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDeselectEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_get_deselect_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Clears all tabs.
*/
//go:nosplit
func (self class) ClearTabs() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabBar.Bind_clear_tabs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Instance) OnTabSelected(cb func(tab int)) {
	self[0].AsObject().Connect(gd.NewStringName("tab_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTabChanged(cb func(tab int)) {
	self[0].AsObject().Connect(gd.NewStringName("tab_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTabClicked(cb func(tab int)) {
	self[0].AsObject().Connect(gd.NewStringName("tab_clicked"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTabRmbClicked(cb func(tab int)) {
	self[0].AsObject().Connect(gd.NewStringName("tab_rmb_clicked"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTabClosePressed(cb func(tab int)) {
	self[0].AsObject().Connect(gd.NewStringName("tab_close_pressed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTabButtonPressed(cb func(tab int)) {
	self[0].AsObject().Connect(gd.NewStringName("tab_button_pressed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTabHovered(cb func(tab int)) {
	self[0].AsObject().Connect(gd.NewStringName("tab_hovered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnActiveTabRearranged(cb func(idx_to int)) {
	self[0].AsObject().Connect(gd.NewStringName("active_tab_rearranged"), gd.NewCallable(cb), 0)
}

func (self class) AsTabBar() Advanced          { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTabBar() Instance       { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.Advanced { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsControl(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsControl(), name)
	}
}
func init() { classdb.Register("TabBar", func(ptr gd.Object) any { return classdb.TabBar(ptr) }) }

type AlignmentMode = classdb.TabBarAlignmentMode

const (
	/*Places tabs to the left.*/
	AlignmentLeft AlignmentMode = 0
	/*Places tabs in the middle.*/
	AlignmentCenter AlignmentMode = 1
	/*Places tabs to the right.*/
	AlignmentRight AlignmentMode = 2
	/*Represents the size of the [enum AlignmentMode] enum.*/
	AlignmentMax AlignmentMode = 3
)

type CloseButtonDisplayPolicy = classdb.TabBarCloseButtonDisplayPolicy

const (
	/*Never show the close buttons.*/
	CloseButtonShowNever CloseButtonDisplayPolicy = 0
	/*Only show the close button on the currently active tab.*/
	CloseButtonShowActiveOnly CloseButtonDisplayPolicy = 1
	/*Show the close button on all tabs.*/
	CloseButtonShowAlways CloseButtonDisplayPolicy = 2
	/*Represents the size of the [enum CloseButtonDisplayPolicy] enum.*/
	CloseButtonMax CloseButtonDisplayPolicy = 3
)
