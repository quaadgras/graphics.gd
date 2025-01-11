// Package TabContainer provides methods for working with TabContainer object instances.
package TabContainer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Container"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
Arranges child controls into a tabbed view, creating a tab for each one. The active tab's corresponding control is made visible, while all other child controls are hidden. Ignores non-control children.
[b]Note:[/b] The drawing of the clickable tabs is handled by this node; [TabBar] is not needed.
*/
type Instance [1]gdclass.TabContainer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTabContainer() Instance
}

/*
Returns the number of tabs.
*/
func (self Instance) GetTabCount() int {
	return int(int(class(self).GetTabCount()))
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
Returns the child [Control] node located at the active tab index.
*/
func (self Instance) GetCurrentTabControl() [1]gdclass.Control {
	return [1]gdclass.Control(class(self).GetCurrentTabControl())
}

/*
Returns the [TabBar] contained in this container.
[b]Warning:[/b] This is a required internal node, removing and freeing it or editing its tabs may cause a crash. If you wish to edit the tabs, use the methods provided in [TabContainer].
*/
func (self Instance) GetTabBar() [1]gdclass.TabBar {
	return [1]gdclass.TabBar(class(self).GetTabBar())
}

/*
Returns the [Control] node from the tab at index [param tab_idx].
*/
func (self Instance) GetTabControl(tab_idx int) [1]gdclass.Control {
	return [1]gdclass.Control(class(self).GetTabControl(gd.Int(tab_idx)))
}

/*
Sets a custom title for the tab at index [param tab_idx] (tab titles default to the name of the indexed child node). Set it back to the child's name to make the tab default to it again.
*/
func (self Instance) SetTabTitle(tab_idx int, title string) {
	class(self).SetTabTitle(gd.Int(tab_idx), gd.NewString(title))
}

/*
Returns the title of the tab at index [param tab_idx]. Tab titles default to the name of the indexed child node, but this can be overridden with [method set_tab_title].
*/
func (self Instance) GetTabTitle(tab_idx int) string {
	return string(class(self).GetTabTitle(gd.Int(tab_idx)).String())
}

/*
Sets a custom tooltip text for tab at index [param tab_idx].
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
Sets an icon for the tab at index [param tab_idx].
*/
func (self Instance) SetTabIcon(tab_idx int, icon [1]gdclass.Texture2D) {
	class(self).SetTabIcon(gd.Int(tab_idx), icon)
}

/*
Returns the [Texture2D] for the tab at index [param tab_idx] or [code]null[/code] if the tab has no [Texture2D].
*/
func (self Instance) GetTabIcon(tab_idx int) [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTabIcon(gd.Int(tab_idx)))
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
Sets the button icon from the tab at index [param tab_idx].
*/
func (self Instance) SetTabButtonIcon(tab_idx int, icon [1]gdclass.Texture2D) {
	class(self).SetTabButtonIcon(gd.Int(tab_idx), icon)
}

/*
Returns the button icon from the tab at index [param tab_idx].
*/
func (self Instance) GetTabButtonIcon(tab_idx int) [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTabButtonIcon(gd.Int(tab_idx)))
}

/*
Returns the index of the tab at local coordinates [param point]. Returns [code]-1[/code] if the point is outside the control boundaries or if there's no tab at the queried position.
*/
func (self Instance) GetTabIdxAtPoint(point Vector2.XY) int {
	return int(int(class(self).GetTabIdxAtPoint(gd.Vector2(point))))
}

/*
Returns the index of the tab tied to the given [param control]. The control must be a child of the [TabContainer].
*/
func (self Instance) GetTabIdxFromControl(control [1]gdclass.Control) int {
	return int(int(class(self).GetTabIdxFromControl(control)))
}

/*
If set on a [Popup] node instance, a popup menu icon appears in the top-right corner of the [TabContainer] (setting it to [code]null[/code] will make it go away). Clicking it will expand the [Popup] node.
*/
func (self Instance) SetPopup(popup [1]gdclass.Node) {
	class(self).SetPopup(popup)
}

/*
Returns the [Popup] node instance if one has been set already with [method set_popup].
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
func (self Instance) GetPopup() [1]gdclass.Popup {
	return [1]gdclass.Popup(class(self).GetPopup())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TabContainer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TabContainer"))
	casted := Instance{*(*gdclass.TabContainer)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) TabAlignment() gdclass.TabBarAlignmentMode {
	return gdclass.TabBarAlignmentMode(class(self).GetTabAlignment())
}

func (self Instance) SetTabAlignment(value gdclass.TabBarAlignmentMode) {
	class(self).SetTabAlignment(value)
}

func (self Instance) CurrentTab() int {
	return int(int(class(self).GetCurrentTab()))
}

func (self Instance) SetCurrentTab(value int) {
	class(self).SetCurrentTab(gd.Int(value))
}

func (self Instance) TabsPosition() gdclass.TabContainerTabPosition {
	return gdclass.TabContainerTabPosition(class(self).GetTabsPosition())
}

func (self Instance) SetTabsPosition(value gdclass.TabContainerTabPosition) {
	class(self).SetTabsPosition(value)
}

func (self Instance) ClipTabs() bool {
	return bool(class(self).GetClipTabs())
}

func (self Instance) SetClipTabs(value bool) {
	class(self).SetClipTabs(value)
}

func (self Instance) TabsVisible() bool {
	return bool(class(self).AreTabsVisible())
}

func (self Instance) SetTabsVisible(value bool) {
	class(self).SetTabsVisible(value)
}

func (self Instance) AllTabsInFront() bool {
	return bool(class(self).IsAllTabsInFront())
}

func (self Instance) SetAllTabsInFront(value bool) {
	class(self).SetAllTabsInFront(value)
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

func (self Instance) UseHiddenTabsForMinSize() bool {
	return bool(class(self).GetUseHiddenTabsForMinSize())
}

func (self Instance) SetUseHiddenTabsForMinSize(value bool) {
	class(self).SetUseHiddenTabsForMinSize(value)
}

func (self Instance) TabFocusMode() gdclass.ControlFocusMode {
	return gdclass.ControlFocusMode(class(self).GetTabFocusMode())
}

func (self Instance) SetTabFocusMode(value gdclass.ControlFocusMode) {
	class(self).SetTabFocusMode(value)
}

func (self Instance) DeselectEnabled() bool {
	return bool(class(self).GetDeselectEnabled())
}

func (self Instance) SetDeselectEnabled(value bool) {
	class(self).SetDeselectEnabled(value)
}

/*
Returns the number of tabs.
*/
//go:nosplit
func (self class) GetTabCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCurrentTab(tab_idx gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_current_tab, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCurrentTab() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_current_tab, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_previous_tab, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_select_previous_available, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_select_next_available, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the child [Control] node located at the active tab index.
*/
//go:nosplit
func (self class) GetCurrentTabControl() [1]gdclass.Control {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_current_tab_control, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Control{gd.PointerMustAssertInstanceID[gdclass.Control](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the [TabBar] contained in this container.
[b]Warning:[/b] This is a required internal node, removing and freeing it or editing its tabs may cause a crash. If you wish to edit the tabs, use the methods provided in [TabContainer].
*/
//go:nosplit
func (self class) GetTabBar() [1]gdclass.TabBar {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_bar, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TabBar{gd.PointerLifetimeBoundTo[gdclass.TabBar](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the [Control] node from the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) GetTabControl(tab_idx gd.Int) [1]gdclass.Control {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_control, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Control{gd.PointerMustAssertInstanceID[gdclass.Control](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTabAlignment(alignment gdclass.TabBarAlignmentMode) {
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tab_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTabAlignment() gdclass.TabBarAlignmentMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TabBarAlignmentMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTabsPosition(tabs_position gdclass.TabContainerTabPosition) {
	var frame = callframe.New()
	callframe.Arg(frame, tabs_position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tabs_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTabsPosition() gdclass.TabContainerTabPosition {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TabContainerTabPosition](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tabs_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetClipTabs(clip_tabs bool) {
	var frame = callframe.New()
	callframe.Arg(frame, clip_tabs)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_clip_tabs, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetClipTabs() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_clip_tabs, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTabsVisible(visible bool) {
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tabs_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) AreTabsVisible() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_are_tabs_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllTabsInFront(is_front bool) {
	var frame = callframe.New()
	callframe.Arg(frame, is_front)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_all_tabs_in_front, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsAllTabsInFront() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_is_all_tabs_in_front, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets a custom title for the tab at index [param tab_idx] (tab titles default to the name of the indexed child node). Set it back to the child's name to make the tab default to it again.
*/
//go:nosplit
func (self class) SetTabTitle(tab_idx gd.Int, title gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, pointers.Get(title))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tab_title, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the title of the tab at index [param tab_idx]. Tab titles default to the name of the indexed child node, but this can be overridden with [method set_tab_title].
*/
//go:nosplit
func (self class) GetTabTitle(tab_idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_title, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets a custom tooltip text for tab at index [param tab_idx].
[b]Note:[/b] By default, if the [param tooltip] is empty and the tab text is truncated (not all characters fit into the tab), the title will be displayed as a tooltip. To hide the tooltip, assign [code]" "[/code] as the [param tooltip] text.
*/
//go:nosplit
func (self class) SetTabTooltip(tab_idx gd.Int, tooltip gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, pointers.Get(tooltip))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tab_tooltip, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_tooltip, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets an icon for the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) SetTabIcon(tab_idx gd.Int, icon [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tab_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [Texture2D] for the tab at index [param tab_idx] or [code]null[/code] if the tab has no [Texture2D].
*/
//go:nosplit
func (self class) GetTabIcon(tab_idx gd.Int) [1]gdclass.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
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
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tab_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
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
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tab_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_is_tab_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tab_hidden, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_is_tab_hidden, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tab_metadata, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the metadata value set to the tab at index [param tab_idx] using [method set_tab_metadata]. If no metadata was previously set, returns [code]null[/code] by default.
*/
//go:nosplit
func (self class) GetTabMetadata(tab_idx gd.Int) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[variantPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_metadata, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the button icon from the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) SetTabButtonIcon(tab_idx gd.Int, icon [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tab_button_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the button icon from the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) GetTabButtonIcon(tab_idx gd.Int) [1]gdclass.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_button_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the index of the tab at local coordinates [param point]. Returns [code]-1[/code] if the point is outside the control boundaries or if there's no tab at the queried position.
*/
//go:nosplit
func (self class) GetTabIdxAtPoint(point gd.Vector2) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_idx_at_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the index of the tab tied to the given [param control]. The control must be a child of the [TabContainer].
*/
//go:nosplit
func (self class) GetTabIdxFromControl(control [1]gdclass.Control) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(control[0])[0])
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_idx_from_control, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If set on a [Popup] node instance, a popup menu icon appears in the top-right corner of the [TabContainer] (setting it to [code]null[/code] will make it go away). Clicking it will expand the [Popup] node.
*/
//go:nosplit
func (self class) SetPopup(popup [1]gdclass.Node) {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(popup[0].AsObject()[0]))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_popup, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [Popup] node instance if one has been set already with [method set_popup].
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
//go:nosplit
func (self class) GetPopup() [1]gdclass.Popup {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_popup, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Popup{gd.PointerMustAssertInstanceID[gdclass.Popup](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDragToRearrangeEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_drag_to_rearrange_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDragToRearrangeEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_drag_to_rearrange_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTabsRearrangeGroup(group_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, group_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tabs_rearrange_group, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTabsRearrangeGroup() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tabs_rearrange_group, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseHiddenTabsForMinSize(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_use_hidden_tabs_for_min_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUseHiddenTabsForMinSize() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_use_hidden_tabs_for_min_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTabFocusMode(focus_mode gdclass.ControlFocusMode) {
	var frame = callframe.New()
	callframe.Arg(frame, focus_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tab_focus_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTabFocusMode() gdclass.ControlFocusMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.ControlFocusMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_focus_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDeselectEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_deselect_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDeselectEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_deselect_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnActiveTabRearranged(cb func(idx_to int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("active_tab_rearranged"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTabChanged(cb func(tab int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("tab_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTabClicked(cb func(tab int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("tab_clicked"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTabHovered(cb func(tab int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("tab_hovered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTabSelected(cb func(tab int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("tab_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTabButtonPressed(cb func(tab int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("tab_button_pressed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPrePopupPressed(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("pre_popup_pressed"), gd.NewCallable(cb), 0)
}

func (self class) AsTabContainer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTabContainer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsContainer() Container.Advanced {
	return *((*Container.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsContainer() Container.Instance {
	return *((*Container.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(Container.Advanced(self.AsContainer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Container.Instance(self.AsContainer()), name)
	}
}
func init() {
	gdclass.Register("TabContainer", func(ptr gd.Object) any {
		return [1]gdclass.TabContainer{*(*gdclass.TabContainer)(unsafe.Pointer(&ptr))}
	})
}

type TabPosition = gdclass.TabContainerTabPosition

const (
	/*Places the tab bar at the top.*/
	PositionTop TabPosition = 0
	/*Places the tab bar at the bottom. The tab bar's [StyleBox] will be flipped vertically.*/
	PositionBottom TabPosition = 1
	/*Represents the size of the [enum TabPosition] enum.*/
	PositionMax TabPosition = 2
)
