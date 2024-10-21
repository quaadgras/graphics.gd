package TabBar

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Control"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A control that provides a horizontal bar with tabs. Similar to [TabContainer] but is only in charge of drawing tabs, not interacting with children.

*/
type Simple [1]classdb.TabBar
func (self Simple) SetTabCount(count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTabCount(gd.Int(count))
}
func (self Simple) GetTabCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetTabCount()))
}
func (self Simple) SetCurrentTab(tab_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCurrentTab(gd.Int(tab_idx))
}
func (self Simple) GetCurrentTab() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCurrentTab()))
}
func (self Simple) GetPreviousTab() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPreviousTab()))
}
func (self Simple) SelectPreviousAvailable() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).SelectPreviousAvailable())
}
func (self Simple) SelectNextAvailable() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).SelectNextAvailable())
}
func (self Simple) SetTabTitle(tab_idx int, title string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTabTitle(gd.Int(tab_idx), gc.String(title))
}
func (self Simple) GetTabTitle(tab_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetTabTitle(gc, gd.Int(tab_idx)).String())
}
func (self Simple) SetTabTooltip(tab_idx int, tooltip string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTabTooltip(gd.Int(tab_idx), gc.String(tooltip))
}
func (self Simple) GetTabTooltip(tab_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetTabTooltip(gc, gd.Int(tab_idx)).String())
}
func (self Simple) SetTabTextDirection(tab_idx int, direction classdb.ControlTextDirection) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTabTextDirection(gd.Int(tab_idx), direction)
}
func (self Simple) GetTabTextDirection(tab_idx int) classdb.ControlTextDirection {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ControlTextDirection(Expert(self).GetTabTextDirection(gd.Int(tab_idx)))
}
func (self Simple) SetTabLanguage(tab_idx int, language string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTabLanguage(gd.Int(tab_idx), gc.String(language))
}
func (self Simple) GetTabLanguage(tab_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetTabLanguage(gc, gd.Int(tab_idx)).String())
}
func (self Simple) SetTabIcon(tab_idx int, icon [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTabIcon(gd.Int(tab_idx), icon)
}
func (self Simple) GetTabIcon(tab_idx int) [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetTabIcon(gc, gd.Int(tab_idx)))
}
func (self Simple) SetTabIconMaxWidth(tab_idx int, width int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTabIconMaxWidth(gd.Int(tab_idx), gd.Int(width))
}
func (self Simple) GetTabIconMaxWidth(tab_idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetTabIconMaxWidth(gd.Int(tab_idx))))
}
func (self Simple) SetTabButtonIcon(tab_idx int, icon [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTabButtonIcon(gd.Int(tab_idx), icon)
}
func (self Simple) GetTabButtonIcon(tab_idx int) [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetTabButtonIcon(gc, gd.Int(tab_idx)))
}
func (self Simple) SetTabDisabled(tab_idx int, disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTabDisabled(gd.Int(tab_idx), disabled)
}
func (self Simple) IsTabDisabled(tab_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsTabDisabled(gd.Int(tab_idx)))
}
func (self Simple) SetTabHidden(tab_idx int, hidden bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTabHidden(gd.Int(tab_idx), hidden)
}
func (self Simple) IsTabHidden(tab_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsTabHidden(gd.Int(tab_idx)))
}
func (self Simple) SetTabMetadata(tab_idx int, metadata gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTabMetadata(gd.Int(tab_idx), metadata)
}
func (self Simple) GetTabMetadata(tab_idx int) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetTabMetadata(gc, gd.Int(tab_idx)))
}
func (self Simple) RemoveTab(tab_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveTab(gd.Int(tab_idx))
}
func (self Simple) AddTab(title string, icon [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddTab(gc.String(title), icon)
}
func (self Simple) GetTabIdxAtPoint(point gd.Vector2) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetTabIdxAtPoint(point)))
}
func (self Simple) SetTabAlignment(alignment classdb.TabBarAlignmentMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTabAlignment(alignment)
}
func (self Simple) GetTabAlignment() classdb.TabBarAlignmentMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TabBarAlignmentMode(Expert(self).GetTabAlignment())
}
func (self Simple) SetClipTabs(clip_tabs bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetClipTabs(clip_tabs)
}
func (self Simple) GetClipTabs() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetClipTabs())
}
func (self Simple) GetTabOffset() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetTabOffset()))
}
func (self Simple) GetOffsetButtonsVisible() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetOffsetButtonsVisible())
}
func (self Simple) EnsureTabVisible(idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnsureTabVisible(gd.Int(idx))
}
func (self Simple) GetTabRect(tab_idx int) gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(Expert(self).GetTabRect(gd.Int(tab_idx)))
}
func (self Simple) MoveTab(from int, to int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MoveTab(gd.Int(from), gd.Int(to))
}
func (self Simple) SetTabCloseDisplayPolicy(policy classdb.TabBarCloseButtonDisplayPolicy) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTabCloseDisplayPolicy(policy)
}
func (self Simple) GetTabCloseDisplayPolicy() classdb.TabBarCloseButtonDisplayPolicy {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TabBarCloseButtonDisplayPolicy(Expert(self).GetTabCloseDisplayPolicy())
}
func (self Simple) SetMaxTabWidth(width int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaxTabWidth(gd.Int(width))
}
func (self Simple) GetMaxTabWidth() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMaxTabWidth()))
}
func (self Simple) SetScrollingEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetScrollingEnabled(enabled)
}
func (self Simple) GetScrollingEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetScrollingEnabled())
}
func (self Simple) SetDragToRearrangeEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDragToRearrangeEnabled(enabled)
}
func (self Simple) GetDragToRearrangeEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetDragToRearrangeEnabled())
}
func (self Simple) SetTabsRearrangeGroup(group_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTabsRearrangeGroup(gd.Int(group_id))
}
func (self Simple) GetTabsRearrangeGroup() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetTabsRearrangeGroup()))
}
func (self Simple) SetScrollToSelected(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetScrollToSelected(enabled)
}
func (self Simple) GetScrollToSelected() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetScrollToSelected())
}
func (self Simple) SetSelectWithRmb(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSelectWithRmb(enabled)
}
func (self Simple) GetSelectWithRmb() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetSelectWithRmb())
}
func (self Simple) SetDeselectEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDeselectEnabled(enabled)
}
func (self Simple) GetDeselectEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetDeselectEnabled())
}
func (self Simple) ClearTabs() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearTabs()
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.TabBar
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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
func (self class) SetTabIcon(tab_idx gd.Int, icon object.Texture2D)  {
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
func (self class) GetTabIcon(ctx gd.Lifetime, tab_idx gd.Int) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_tab_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
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
func (self class) SetTabButtonIcon(tab_idx gd.Int, icon object.Texture2D)  {
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
func (self class) GetTabButtonIcon(ctx gd.Lifetime, tab_idx gd.Int) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabBar.Bind_get_tab_button_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
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
func (self class) AddTab(title gd.String, icon object.Texture2D)  {
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

//go:nosplit
func (self class) AsTabBar() Expert { return self[0].AsTabBar() }


//go:nosplit
func (self Simple) AsTabBar() Simple { return self[0].AsTabBar() }


//go:nosplit
func (self class) AsControl() Control.Expert { return self[0].AsControl() }


//go:nosplit
func (self Simple) AsControl() Control.Simple { return self[0].AsControl() }


//go:nosplit
func (self class) AsCanvasItem() CanvasItem.Expert { return self[0].AsCanvasItem() }


//go:nosplit
func (self Simple) AsCanvasItem() CanvasItem.Simple { return self[0].AsCanvasItem() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("TabBar", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
