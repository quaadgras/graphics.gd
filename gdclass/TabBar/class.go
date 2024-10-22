package TabBar

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A control that provides a horizontal bar with tabs. Similar to [TabContainer] but is only in charge of drawing tabs, not interacting with children.

*/
type Go [1]classdb.TabBar

/*
Returns the previously active tab index.
*/
func (self Go) GetPreviousTab() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetPreviousTab()))
}

/*
Selects the first available tab with lower index than the currently selected. Returns [code]true[/code] if tab selection changed.
*/
func (self Go) SelectPreviousAvailable() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).SelectPreviousAvailable())
}

/*
Selects the first available tab with greater index than the currently selected. Returns [code]true[/code] if tab selection changed.
*/
func (self Go) SelectNextAvailable() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).SelectNextAvailable())
}

/*
Sets a [param title] for the tab at index [param tab_idx].
*/
func (self Go) SetTabTitle(tab_idx int, title string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabTitle(gd.Int(tab_idx), gc.String(title))
}

/*
Returns the title of the tab at index [param tab_idx].
*/
func (self Go) GetTabTitle(tab_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetTabTitle(gc, gd.Int(tab_idx)).String())
}

/*
Sets a [param tooltip] for tab at index [param tab_idx].
[b]Note:[/b] By default, if the [param tooltip] is empty and the tab text is truncated (not all characters fit into the tab), the title will be displayed as a tooltip. To hide the tooltip, assign [code]" "[/code] as the [param tooltip] text.
*/
func (self Go) SetTabTooltip(tab_idx int, tooltip string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabTooltip(gd.Int(tab_idx), gc.String(tooltip))
}

/*
Returns the tooltip text of the tab at index [param tab_idx].
*/
func (self Go) GetTabTooltip(tab_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetTabTooltip(gc, gd.Int(tab_idx)).String())
}

/*
Sets tab title base writing direction.
*/
func (self Go) SetTabTextDirection(tab_idx int, direction classdb.ControlTextDirection) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabTextDirection(gd.Int(tab_idx), direction)
}

/*
Returns tab title text base writing direction.
*/
func (self Go) GetTabTextDirection(tab_idx int) classdb.ControlTextDirection {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ControlTextDirection(class(self).GetTabTextDirection(gd.Int(tab_idx)))
}

/*
Sets language code of tab title used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
func (self Go) SetTabLanguage(tab_idx int, language string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabLanguage(gd.Int(tab_idx), gc.String(language))
}

/*
Returns tab title language code.
*/
func (self Go) GetTabLanguage(tab_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetTabLanguage(gc, gd.Int(tab_idx)).String())
}

/*
Sets an [param icon] for the tab at index [param tab_idx].
*/
func (self Go) SetTabIcon(tab_idx int, icon gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabIcon(gd.Int(tab_idx), icon)
}

/*
Returns the icon for the tab at index [param tab_idx] or [code]null[/code] if the tab has no icon.
*/
func (self Go) GetTabIcon(tab_idx int) gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Texture2D(class(self).GetTabIcon(gc, gd.Int(tab_idx)))
}

/*
Sets the maximum allowed width of the icon for the tab at index [param tab_idx]. This limit is applied on top of the default size of the icon and on top of [theme_item icon_max_width]. The height is adjusted according to the icon's ratio.
*/
func (self Go) SetTabIconMaxWidth(tab_idx int, width int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabIconMaxWidth(gd.Int(tab_idx), gd.Int(width))
}

/*
Returns the maximum allowed width of the icon for the tab at index [param tab_idx].
*/
func (self Go) GetTabIconMaxWidth(tab_idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetTabIconMaxWidth(gd.Int(tab_idx))))
}

/*
Sets an [param icon] for the button of the tab at index [param tab_idx] (located to the right, before the close button), making it visible and clickable (See [signal tab_button_pressed]). Giving it a [code]null[/code] value will hide the button.
*/
func (self Go) SetTabButtonIcon(tab_idx int, icon gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabButtonIcon(gd.Int(tab_idx), icon)
}

/*
Returns the icon for the right button of the tab at index [param tab_idx] or [code]null[/code] if the right button has no icon.
*/
func (self Go) GetTabButtonIcon(tab_idx int) gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Texture2D(class(self).GetTabButtonIcon(gc, gd.Int(tab_idx)))
}

/*
If [param disabled] is [code]true[/code], disables the tab at index [param tab_idx], making it non-interactable.
*/
func (self Go) SetTabDisabled(tab_idx int, disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabDisabled(gd.Int(tab_idx), disabled)
}

/*
Returns [code]true[/code] if the tab at index [param tab_idx] is disabled.
*/
func (self Go) IsTabDisabled(tab_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsTabDisabled(gd.Int(tab_idx)))
}

/*
If [param hidden] is [code]true[/code], hides the tab at index [param tab_idx], making it disappear from the tab area.
*/
func (self Go) SetTabHidden(tab_idx int, hidden bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabHidden(gd.Int(tab_idx), hidden)
}

/*
Returns [code]true[/code] if the tab at index [param tab_idx] is hidden.
*/
func (self Go) IsTabHidden(tab_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsTabHidden(gd.Int(tab_idx)))
}

/*
Sets the metadata value for the tab at index [param tab_idx], which can be retrieved later using [method get_tab_metadata].
*/
func (self Go) SetTabMetadata(tab_idx int, metadata gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabMetadata(gd.Int(tab_idx), metadata)
}

/*
Returns the metadata value set to the tab at index [param tab_idx] using [method set_tab_metadata]. If no metadata was previously set, returns [code]null[/code] by default.
*/
func (self Go) GetTabMetadata(tab_idx int) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).GetTabMetadata(gc, gd.Int(tab_idx)))
}

/*
Removes the tab at index [param tab_idx].
*/
func (self Go) RemoveTab(tab_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveTab(gd.Int(tab_idx))
}

/*
Adds a new tab.
*/
func (self Go) AddTab() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddTab(gc.String(""), ([1]gdclass.Texture2D{}[0]))
}

/*
Returns the index of the tab at local coordinates [param point]. Returns [code]-1[/code] if the point is outside the control boundaries or if there's no tab at the queried position.
*/
func (self Go) GetTabIdxAtPoint(point gd.Vector2) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetTabIdxAtPoint(point)))
}

/*
Returns the number of hidden tabs offsetted to the left.
*/
func (self Go) GetTabOffset() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetTabOffset()))
}

/*
Returns [code]true[/code] if the offset buttons (the ones that appear when there's not enough space for all tabs) are visible.
*/
func (self Go) GetOffsetButtonsVisible() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetOffsetButtonsVisible())
}

/*
Moves the scroll view to make the tab visible.
*/
func (self Go) EnsureTabVisible(idx int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).EnsureTabVisible(gd.Int(idx))
}

/*
Returns tab [Rect2] with local position and size.
*/
func (self Go) GetTabRect(tab_idx int) gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(class(self).GetTabRect(gd.Int(tab_idx)))
}

/*
Moves a tab from [param from] to [param to].
*/
func (self Go) MoveTab(from int, to int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).MoveTab(gd.Int(from), gd.Int(to))
}

/*
Clears all tabs.
*/
func (self Go) ClearTabs() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearTabs()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.TabBar
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("TabBar"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) CurrentTab() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCurrentTab()))
}

func (self Go) SetCurrentTab(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCurrentTab(gd.Int(value))
}

func (self Go) TabAlignment() classdb.TabBarAlignmentMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TabBarAlignmentMode(class(self).GetTabAlignment())
}

func (self Go) SetTabAlignment(value classdb.TabBarAlignmentMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabAlignment(value)
}

func (self Go) ClipTabs() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetClipTabs())
}

func (self Go) SetClipTabs(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetClipTabs(value)
}

func (self Go) TabCloseDisplayPolicy() classdb.TabBarCloseButtonDisplayPolicy {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TabBarCloseButtonDisplayPolicy(class(self).GetTabCloseDisplayPolicy())
}

func (self Go) SetTabCloseDisplayPolicy(value classdb.TabBarCloseButtonDisplayPolicy) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabCloseDisplayPolicy(value)
}

func (self Go) MaxTabWidth() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMaxTabWidth()))
}

func (self Go) SetMaxTabWidth(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxTabWidth(gd.Int(value))
}

func (self Go) ScrollingEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetScrollingEnabled())
}

func (self Go) SetScrollingEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetScrollingEnabled(value)
}

func (self Go) DragToRearrangeEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetDragToRearrangeEnabled())
}

func (self Go) SetDragToRearrangeEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDragToRearrangeEnabled(value)
}

func (self Go) TabsRearrangeGroup() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetTabsRearrangeGroup()))
}

func (self Go) SetTabsRearrangeGroup(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabsRearrangeGroup(gd.Int(value))
}

func (self Go) ScrollToSelected() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetScrollToSelected())
}

func (self Go) SetScrollToSelected(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetScrollToSelected(value)
}

func (self Go) SelectWithRmb() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetSelectWithRmb())
}

func (self Go) SetSelectWithRmb(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSelectWithRmb(value)
}

func (self Go) DeselectEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetDeselectEnabled())
}

func (self Go) SetDeselectEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDeselectEnabled(value)
}

func (self Go) TabCount() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetTabCount()))
}

func (self Go) SetTabCount(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabCount(gd.Int(value))
}

//go:nosplit
func (self class) SetTabCount(count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_tab_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTabCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_tab_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCurrentTab(tab_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_current_tab, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCurrentTab() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_current_tab, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the previously active tab index.
*/
//go:nosplit
func (self class) GetPreviousTab() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_previous_tab, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Selects the first available tab with lower index than the currently selected. Returns [code]true[/code] if tab selection changed.
*/
//go:nosplit
func (self class) SelectPreviousAvailable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_select_previous_available, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Selects the first available tab with greater index than the currently selected. Returns [code]true[/code] if tab selection changed.
*/
//go:nosplit
func (self class) SelectNextAvailable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_select_next_available, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a [param title] for the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) SetTabTitle(tab_idx gd.Int, title gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, mmm.Get(title))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_tab_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the title of the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) GetTabTitle(ctx gd.Lifetime, tab_idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_tab_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets a [param tooltip] for tab at index [param tab_idx].
[b]Note:[/b] By default, if the [param tooltip] is empty and the tab text is truncated (not all characters fit into the tab), the title will be displayed as a tooltip. To hide the tooltip, assign [code]" "[/code] as the [param tooltip] text.
*/
//go:nosplit
func (self class) SetTabTooltip(tab_idx gd.Int, tooltip gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, mmm.Get(tooltip))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_tab_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the tooltip text of the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) GetTabTooltip(ctx gd.Lifetime, tab_idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_tab_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets tab title base writing direction.
*/
//go:nosplit
func (self class) SetTabTextDirection(tab_idx gd.Int, direction classdb.ControlTextDirection)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_tab_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns tab title text base writing direction.
*/
//go:nosplit
func (self class) GetTabTextDirection(tab_idx gd.Int) classdb.ControlTextDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[classdb.ControlTextDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_tab_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets language code of tab title used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
//go:nosplit
func (self class) SetTabLanguage(tab_idx gd.Int, language gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, mmm.Get(language))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_tab_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns tab title language code.
*/
//go:nosplit
func (self class) GetTabLanguage(ctx gd.Lifetime, tab_idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_tab_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets an [param icon] for the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) SetTabIcon(tab_idx gd.Int, icon gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_tab_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the icon for the tab at index [param tab_idx] or [code]null[/code] if the tab has no icon.
*/
//go:nosplit
func (self class) GetTabIcon(ctx gd.Lifetime, tab_idx gd.Int) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_tab_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the maximum allowed width of the icon for the tab at index [param tab_idx]. This limit is applied on top of the default size of the icon and on top of [theme_item icon_max_width]. The height is adjusted according to the icon's ratio.
*/
//go:nosplit
func (self class) SetTabIconMaxWidth(tab_idx gd.Int, width gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_tab_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the maximum allowed width of the icon for the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) GetTabIconMaxWidth(tab_idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_tab_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets an [param icon] for the button of the tab at index [param tab_idx] (located to the right, before the close button), making it visible and clickable (See [signal tab_button_pressed]). Giving it a [code]null[/code] value will hide the button.
*/
//go:nosplit
func (self class) SetTabButtonIcon(tab_idx gd.Int, icon gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_tab_button_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the icon for the right button of the tab at index [param tab_idx] or [code]null[/code] if the right button has no icon.
*/
//go:nosplit
func (self class) GetTabButtonIcon(ctx gd.Lifetime, tab_idx gd.Int) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_tab_button_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
If [param disabled] is [code]true[/code], disables the tab at index [param tab_idx], making it non-interactable.
*/
//go:nosplit
func (self class) SetTabDisabled(tab_idx gd.Int, disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_tab_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the tab at index [param tab_idx] is disabled.
*/
//go:nosplit
func (self class) IsTabDisabled(tab_idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_is_tab_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param hidden] is [code]true[/code], hides the tab at index [param tab_idx], making it disappear from the tab area.
*/
//go:nosplit
func (self class) SetTabHidden(tab_idx gd.Int, hidden bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, hidden)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_tab_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the tab at index [param tab_idx] is hidden.
*/
//go:nosplit
func (self class) IsTabHidden(tab_idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_is_tab_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the metadata value for the tab at index [param tab_idx], which can be retrieved later using [method get_tab_metadata].
*/
//go:nosplit
func (self class) SetTabMetadata(tab_idx gd.Int, metadata gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, mmm.Get(metadata))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_tab_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the metadata value set to the tab at index [param tab_idx] using [method set_tab_metadata]. If no metadata was previously set, returns [code]null[/code] by default.
*/
//go:nosplit
func (self class) GetTabMetadata(ctx gd.Lifetime, tab_idx gd.Int) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_tab_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Removes the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) RemoveTab(tab_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_remove_tab, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a new tab.
*/
//go:nosplit
func (self class) AddTab(title gd.String, icon gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(title))
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_add_tab, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the index of the tab at local coordinates [param point]. Returns [code]-1[/code] if the point is outside the control boundaries or if there's no tab at the queried position.
*/
//go:nosplit
func (self class) GetTabIdxAtPoint(point gd.Vector2) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_tab_idx_at_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTabAlignment(alignment classdb.TabBarAlignmentMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_tab_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTabAlignment() classdb.TabBarAlignmentMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TabBarAlignmentMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_tab_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetClipTabs(clip_tabs bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, clip_tabs)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_clip_tabs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetClipTabs() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_clip_tabs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of hidden tabs offsetted to the left.
*/
//go:nosplit
func (self class) GetTabOffset() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_tab_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the offset buttons (the ones that appear when there's not enough space for all tabs) are visible.
*/
//go:nosplit
func (self class) GetOffsetButtonsVisible() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_offset_buttons_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Moves the scroll view to make the tab visible.
*/
//go:nosplit
func (self class) EnsureTabVisible(idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_ensure_tab_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns tab [Rect2] with local position and size.
*/
//go:nosplit
func (self class) GetTabRect(tab_idx gd.Int) gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_tab_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Moves a tab from [param from] to [param to].
*/
//go:nosplit
func (self class) MoveTab(from gd.Int, to gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_move_tab, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetTabCloseDisplayPolicy(policy classdb.TabBarCloseButtonDisplayPolicy)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, policy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_tab_close_display_policy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTabCloseDisplayPolicy() classdb.TabBarCloseButtonDisplayPolicy {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TabBarCloseButtonDisplayPolicy](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_tab_close_display_policy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxTabWidth(width gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_max_tab_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxTabWidth() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_max_tab_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetScrollingEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_scrolling_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetScrollingEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_scrolling_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDragToRearrangeEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_drag_to_rearrange_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDragToRearrangeEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_drag_to_rearrange_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTabsRearrangeGroup(group_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, group_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_tabs_rearrange_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTabsRearrangeGroup() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_tabs_rearrange_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetScrollToSelected(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_scroll_to_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetScrollToSelected() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_scroll_to_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSelectWithRmb(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_select_with_rmb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSelectWithRmb() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_select_with_rmb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDeselectEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_set_deselect_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDeselectEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_deselect_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Clears all tabs.
*/
//go:nosplit
func (self class) ClearTabs()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_clear_tabs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnTabSelected(cb func(tab int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("tab_selected"), gc.Callable(cb), 0)
}


func (self Go) OnTabChanged(cb func(tab int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("tab_changed"), gc.Callable(cb), 0)
}


func (self Go) OnTabClicked(cb func(tab int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("tab_clicked"), gc.Callable(cb), 0)
}


func (self Go) OnTabRmbClicked(cb func(tab int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("tab_rmb_clicked"), gc.Callable(cb), 0)
}


func (self Go) OnTabClosePressed(cb func(tab int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("tab_close_pressed"), gc.Callable(cb), 0)
}


func (self Go) OnTabButtonPressed(cb func(tab int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("tab_button_pressed"), gc.Callable(cb), 0)
}


func (self Go) OnTabHovered(cb func(tab int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("tab_hovered"), gc.Callable(cb), 0)
}


func (self Go) OnActiveTabRearranged(cb func(idx_to int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("active_tab_rearranged"), gc.Callable(cb), 0)
}


func (self class) AsTabBar() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTabBar() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsControl(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsControl(), name)
	}
}
func init() {classdb.Register("TabBar", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
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
