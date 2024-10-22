package TabContainer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Container"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Arranges child controls into a tabbed view, creating a tab for each one. The active tab's corresponding control is made visible, while all other child controls are hidden. Ignores non-control children.
[b]Note:[/b] The drawing of the clickable tabs is handled by this node; [TabBar] is not needed.

*/
type Go [1]classdb.TabContainer

/*
Returns the number of tabs.
*/
func (self Go) GetTabCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetTabCount()))
}

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
Returns the child [Control] node located at the active tab index.
*/
func (self Go) GetCurrentTabControl() gdclass.Control {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Control(class(self).GetCurrentTabControl(gc))
}

/*
Returns the [TabBar] contained in this container.
[b]Warning:[/b] This is a required internal node, removing and freeing it or editing its tabs may cause a crash. If you wish to edit the tabs, use the methods provided in [TabContainer].
*/
func (self Go) GetTabBar() gdclass.TabBar {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.TabBar(class(self).GetTabBar(gc))
}

/*
Returns the [Control] node from the tab at index [param tab_idx].
*/
func (self Go) GetTabControl(tab_idx int) gdclass.Control {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Control(class(self).GetTabControl(gc, gd.Int(tab_idx)))
}

/*
Sets a custom title for the tab at index [param tab_idx] (tab titles default to the name of the indexed child node). Set it back to the child's name to make the tab default to it again.
*/
func (self Go) SetTabTitle(tab_idx int, title string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabTitle(gd.Int(tab_idx), gc.String(title))
}

/*
Returns the title of the tab at index [param tab_idx]. Tab titles default to the name of the indexed child node, but this can be overridden with [method set_tab_title].
*/
func (self Go) GetTabTitle(tab_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetTabTitle(gc, gd.Int(tab_idx)).String())
}

/*
Sets a custom tooltip text for tab at index [param tab_idx].
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
Sets an icon for the tab at index [param tab_idx].
*/
func (self Go) SetTabIcon(tab_idx int, icon gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabIcon(gd.Int(tab_idx), icon)
}

/*
Returns the [Texture2D] for the tab at index [param tab_idx] or [code]null[/code] if the tab has no [Texture2D].
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
Sets the button icon from the tab at index [param tab_idx].
*/
func (self Go) SetTabButtonIcon(tab_idx int, icon gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabButtonIcon(gd.Int(tab_idx), icon)
}

/*
Returns the button icon from the tab at index [param tab_idx].
*/
func (self Go) GetTabButtonIcon(tab_idx int) gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Texture2D(class(self).GetTabButtonIcon(gc, gd.Int(tab_idx)))
}

/*
Returns the index of the tab at local coordinates [param point]. Returns [code]-1[/code] if the point is outside the control boundaries or if there's no tab at the queried position.
*/
func (self Go) GetTabIdxAtPoint(point gd.Vector2) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetTabIdxAtPoint(point)))
}

/*
Returns the index of the tab tied to the given [param control]. The control must be a child of the [TabContainer].
*/
func (self Go) GetTabIdxFromControl(control gdclass.Control) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetTabIdxFromControl(control)))
}

/*
If set on a [Popup] node instance, a popup menu icon appears in the top-right corner of the [TabContainer] (setting it to [code]null[/code] will make it go away). Clicking it will expand the [Popup] node.
*/
func (self Go) SetPopup(popup gdclass.Node) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPopup(popup)
}

/*
Returns the [Popup] node instance if one has been set already with [method set_popup].
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
func (self Go) GetPopup() gdclass.Popup {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Popup(class(self).GetPopup(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.TabContainer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("TabContainer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) TabAlignment() classdb.TabBarAlignmentMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TabBarAlignmentMode(class(self).GetTabAlignment())
}

func (self Go) SetTabAlignment(value classdb.TabBarAlignmentMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabAlignment(value)
}

func (self Go) CurrentTab() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCurrentTab()))
}

func (self Go) SetCurrentTab(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCurrentTab(gd.Int(value))
}

func (self Go) TabsPosition() classdb.TabContainerTabPosition {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TabContainerTabPosition(class(self).GetTabsPosition())
}

func (self Go) SetTabsPosition(value classdb.TabContainerTabPosition) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabsPosition(value)
}

func (self Go) ClipTabs() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetClipTabs())
}

func (self Go) SetClipTabs(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetClipTabs(value)
}

func (self Go) TabsVisible() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).AreTabsVisible())
}

func (self Go) SetTabsVisible(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabsVisible(value)
}

func (self Go) AllTabsInFront() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsAllTabsInFront())
}

func (self Go) SetAllTabsInFront(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAllTabsInFront(value)
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

func (self Go) UseHiddenTabsForMinSize() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetUseHiddenTabsForMinSize())
}

func (self Go) SetUseHiddenTabsForMinSize(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseHiddenTabsForMinSize(value)
}

func (self Go) TabFocusMode() classdb.ControlFocusMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.ControlFocusMode(class(self).GetTabFocusMode())
}

func (self Go) SetTabFocusMode(value classdb.ControlFocusMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabFocusMode(value)
}

func (self Go) DeselectEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetDeselectEnabled())
}

func (self Go) SetDeselectEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDeselectEnabled(value)
}

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
func (self class) GetCurrentTabControl(ctx gd.Lifetime) gdclass.Control {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_current_tab_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Control
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the [TabBar] contained in this container.
[b]Warning:[/b] This is a required internal node, removing and freeing it or editing its tabs may cause a crash. If you wish to edit the tabs, use the methods provided in [TabContainer].
*/
//go:nosplit
func (self class) GetTabBar(ctx gd.Lifetime) gdclass.TabBar {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_tab_bar, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.TabBar
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the [Control] node from the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) GetTabControl(ctx gd.Lifetime, tab_idx gd.Int) gdclass.Control {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_tab_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Control
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
func (self class) SetTabIcon(tab_idx gd.Int, icon gdclass.Texture2D)  {
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
func (self class) GetTabIcon(ctx gd.Lifetime, tab_idx gd.Int) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_tab_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) SetTabButtonIcon(tab_idx gd.Int, icon gdclass.Texture2D)  {
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
func (self class) GetTabButtonIcon(ctx gd.Lifetime, tab_idx gd.Int) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_tab_button_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
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
func (self class) GetTabIdxFromControl(control gdclass.Control) gd.Int {
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
func (self class) SetPopup(popup gdclass.Node)  {
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
func (self class) GetPopup(ctx gd.Lifetime) gdclass.Popup {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TabContainer.Bind_get_popup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Popup
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
func (self Go) OnActiveTabRearranged(cb func(idx_to int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("active_tab_rearranged"), gc.Callable(cb), 0)
}


func (self Go) OnTabChanged(cb func(tab int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("tab_changed"), gc.Callable(cb), 0)
}


func (self Go) OnTabClicked(cb func(tab int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("tab_clicked"), gc.Callable(cb), 0)
}


func (self Go) OnTabHovered(cb func(tab int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("tab_hovered"), gc.Callable(cb), 0)
}


func (self Go) OnTabSelected(cb func(tab int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("tab_selected"), gc.Callable(cb), 0)
}


func (self Go) OnTabButtonPressed(cb func(tab int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("tab_button_pressed"), gc.Callable(cb), 0)
}


func (self Go) OnPrePopupPressed(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("pre_popup_pressed"), gc.Callable(cb), 0)
}


func (self class) AsTabContainer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTabContainer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsContainer() Container.GD { return *((*Container.GD)(unsafe.Pointer(&self))) }
func (self Go) AsContainer() Container.Go { return *((*Container.Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsContainer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsContainer(), name)
	}
}
func init() {classdb.Register("TabContainer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type TabPosition = classdb.TabContainerTabPosition

const (
/*Places the tab bar at the top.*/
	PositionTop TabPosition = 0
/*Places the tab bar at the bottom. The tab bar's [StyleBox] will be flipped vertically.*/
	PositionBottom TabPosition = 1
/*Represents the size of the [enum TabPosition] enum.*/
	PositionMax TabPosition = 2
)
