package TabContainer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Container"
import "grow.graphics/gd/object/Control"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Arranges child controls into a tabbed view, creating a tab for each one. The active tab's corresponding control is made visible, while all other child controls are hidden. Ignores non-control children.
[b]Note:[/b] The drawing of the clickable tabs is handled by this node; [TabBar] is not needed.

*/
type Simple [1]classdb.TabContainer
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
func (self Simple) GetCurrentTabControl() [1]classdb.Control {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Control(Expert(self).GetCurrentTabControl(gc))
}
func (self Simple) GetTabBar() [1]classdb.TabBar {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TabBar(Expert(self).GetTabBar(gc))
}
func (self Simple) GetTabControl(tab_idx int) [1]classdb.Control {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Control(Expert(self).GetTabControl(gc, gd.Int(tab_idx)))
}
func (self Simple) SetTabAlignment(alignment classdb.TabBarAlignmentMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTabAlignment(alignment)
}
func (self Simple) GetTabAlignment() classdb.TabBarAlignmentMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TabBarAlignmentMode(Expert(self).GetTabAlignment())
}
func (self Simple) SetTabsPosition(tabs_position classdb.TabContainerTabPosition) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTabsPosition(tabs_position)
}
func (self Simple) GetTabsPosition() classdb.TabContainerTabPosition {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TabContainerTabPosition(Expert(self).GetTabsPosition())
}
func (self Simple) SetClipTabs(clip_tabs bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetClipTabs(clip_tabs)
}
func (self Simple) GetClipTabs() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetClipTabs())
}
func (self Simple) SetTabsVisible(visible bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTabsVisible(visible)
}
func (self Simple) AreTabsVisible() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).AreTabsVisible())
}
func (self Simple) SetAllTabsInFront(is_front bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAllTabsInFront(is_front)
}
func (self Simple) IsAllTabsInFront() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAllTabsInFront())
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
func (self Simple) SetTabButtonIcon(tab_idx int, icon [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTabButtonIcon(gd.Int(tab_idx), icon)
}
func (self Simple) GetTabButtonIcon(tab_idx int) [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetTabButtonIcon(gc, gd.Int(tab_idx)))
}
func (self Simple) GetTabIdxAtPoint(point gd.Vector2) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetTabIdxAtPoint(point)))
}
func (self Simple) GetTabIdxFromControl(control [1]classdb.Control) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetTabIdxFromControl(control)))
}
func (self Simple) SetPopup(popup [1]classdb.Node) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPopup(popup)
}
func (self Simple) GetPopup() [1]classdb.Popup {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Popup(Expert(self).GetPopup(gc))
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
func (self Simple) SetUseHiddenTabsForMinSize(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseHiddenTabsForMinSize(enabled)
}
func (self Simple) GetUseHiddenTabsForMinSize() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetUseHiddenTabsForMinSize())
}
func (self Simple) SetTabFocusMode(focus_mode classdb.ControlFocusMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTabFocusMode(focus_mode)
}
func (self Simple) GetTabFocusMode() classdb.ControlFocusMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ControlFocusMode(Expert(self).GetTabFocusMode())
}
func (self Simple) SetDeselectEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDeselectEnabled(enabled)
}
func (self Simple) GetDeselectEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetDeselectEnabled())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.TabContainer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the number of tabs.
*/
//go:nosplit
func (self class) GetTabCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_tab_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_set_current_tab, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCurrentTab() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_current_tab, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_previous_tab, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_select_previous_available, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_select_next_available, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the child [Control] node located at the active tab index.
*/
//go:nosplit
func (self class) GetCurrentTabControl(ctx gd.Lifetime) object.Control {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_current_tab_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Control
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the [TabBar] contained in this container.
[b]Warning:[/b] This is a required internal node, removing and freeing it or editing its tabs may cause a crash. If you wish to edit the tabs, use the methods provided in [TabContainer].
*/
//go:nosplit
func (self class) GetTabBar(ctx gd.Lifetime) object.TabBar {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_tab_bar, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TabBar
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the [Control] node from the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) GetTabControl(ctx gd.Lifetime, tab_idx gd.Int) object.Control {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_tab_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Control
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTabAlignment(alignment classdb.TabBarAlignmentMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_set_tab_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTabAlignment() classdb.TabBarAlignmentMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TabBarAlignmentMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_tab_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTabsPosition(tabs_position classdb.TabContainerTabPosition)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tabs_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_set_tabs_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTabsPosition() classdb.TabContainerTabPosition {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TabContainerTabPosition](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_tabs_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_set_clip_tabs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetClipTabs() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_clip_tabs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTabsVisible(visible bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_set_tabs_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) AreTabsVisible() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_are_tabs_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAllTabsInFront(is_front bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, is_front)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_set_all_tabs_in_front, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAllTabsInFront() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_is_all_tabs_in_front, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a custom title for the tab at index [param tab_idx] (tab titles default to the name of the indexed child node). Set it back to the child's name to make the tab default to it again.
*/
//go:nosplit
func (self class) SetTabTitle(tab_idx gd.Int, title gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, mmm.Get(title))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_set_tab_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the title of the tab at index [param tab_idx]. Tab titles default to the name of the indexed child node, but this can be overridden with [method set_tab_title].
*/
//go:nosplit
func (self class) GetTabTitle(ctx gd.Lifetime, tab_idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_tab_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets a custom tooltip text for tab at index [param tab_idx].
[b]Note:[/b] By default, if the [param tooltip] is empty and the tab text is truncated (not all characters fit into the tab), the title will be displayed as a tooltip. To hide the tooltip, assign [code]" "[/code] as the [param tooltip] text.
*/
//go:nosplit
func (self class) SetTabTooltip(tab_idx gd.Int, tooltip gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, mmm.Get(tooltip))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_set_tab_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_tab_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets an icon for the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) SetTabIcon(tab_idx gd.Int, icon object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_set_tab_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Texture2D] for the tab at index [param tab_idx] or [code]null[/code] if the tab has no [Texture2D].
*/
//go:nosplit
func (self class) GetTabIcon(ctx gd.Lifetime, tab_idx gd.Int) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_tab_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_set_tab_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_tab_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_set_tab_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_is_tab_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_set_tab_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_is_tab_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_set_tab_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_tab_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the button icon from the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) SetTabButtonIcon(tab_idx gd.Int, icon object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_set_tab_button_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the button icon from the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) GetTabButtonIcon(ctx gd.Lifetime, tab_idx gd.Int) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_tab_button_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_tab_idx_at_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the index of the tab tied to the given [param control]. The control must be a child of the [TabContainer].
*/
//go:nosplit
func (self class) GetTabIdxFromControl(control object.Control) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(control[0].AsPointer())[0])
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_tab_idx_from_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If set on a [Popup] node instance, a popup menu icon appears in the top-right corner of the [TabContainer] (setting it to [code]null[/code] will make it go away). Clicking it will expand the [Popup] node.
*/
//go:nosplit
func (self class) SetPopup(popup object.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(popup[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_set_popup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Popup] node instance if one has been set already with [method set_popup].
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
//go:nosplit
func (self class) GetPopup(ctx gd.Lifetime) object.Popup {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_popup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Popup
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDragToRearrangeEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_set_drag_to_rearrange_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDragToRearrangeEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_drag_to_rearrange_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_set_tabs_rearrange_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTabsRearrangeGroup() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_tabs_rearrange_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseHiddenTabsForMinSize(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_set_use_hidden_tabs_for_min_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUseHiddenTabsForMinSize() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_use_hidden_tabs_for_min_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTabFocusMode(focus_mode classdb.ControlFocusMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, focus_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_set_tab_focus_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTabFocusMode() classdb.ControlFocusMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ControlFocusMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_tab_focus_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_set_deselect_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDeselectEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_deselect_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsTabContainer() Expert { return self[0].AsTabContainer() }


//go:nosplit
func (self Simple) AsTabContainer() Simple { return self[0].AsTabContainer() }


//go:nosplit
func (self class) AsContainer() Container.Expert { return self[0].AsContainer() }


//go:nosplit
func (self Simple) AsContainer() Container.Simple { return self[0].AsContainer() }


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

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("TabContainer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type TabPosition = classdb.TabContainerTabPosition

const (
/*Places the tab bar at the top.*/
	PositionTop TabPosition = 0
/*Places the tab bar at the bottom. The tab bar's [StyleBox] will be flipped vertically.*/
	PositionBottom TabPosition = 1
/*Represents the size of the [enum TabPosition] enum.*/
	PositionMax TabPosition = 2
)
