package DisplayServer

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[DisplayServer] handles everything related to window management. It is separated from [OS] as a single operating system may support multiple display servers.
[b]Headless mode:[/b] Starting the engine with the [code]--headless[/code] [url=$DOCS_URL/tutorials/editor/command_line_tutorial.html]command line argument[/url] disables all rendering and window management functions. Most functions from [DisplayServer] will return dummy values in this case.

*/
type Simple [1]classdb.DisplayServer
func (self Simple) HasFeature(feature classdb.DisplayServerFeature) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasFeature(feature))
}
func (self Simple) GetName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetName(gc).String())
}
func (self Simple) HelpSetSearchCallbacks(search_callback gd.Callable, action_callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).HelpSetSearchCallbacks(search_callback, action_callback)
}
func (self Simple) GlobalMenuSetPopupCallbacks(menu_root string, open_callback gd.Callable, close_callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalMenuSetPopupCallbacks(gc.String(menu_root), open_callback, close_callback)
}
func (self Simple) GlobalMenuAddSubmenuItem(menu_root string, label string, submenu string, index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GlobalMenuAddSubmenuItem(gc.String(menu_root), gc.String(label), gc.String(submenu), gd.Int(index))))
}
func (self Simple) GlobalMenuAddItem(menu_root string, label string, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GlobalMenuAddItem(gc.String(menu_root), gc.String(label), callback, key_callback, tag, accelerator, gd.Int(index))))
}
func (self Simple) GlobalMenuAddCheckItem(menu_root string, label string, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GlobalMenuAddCheckItem(gc.String(menu_root), gc.String(label), callback, key_callback, tag, accelerator, gd.Int(index))))
}
func (self Simple) GlobalMenuAddIconItem(menu_root string, icon [1]classdb.Texture2D, label string, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GlobalMenuAddIconItem(gc.String(menu_root), icon, gc.String(label), callback, key_callback, tag, accelerator, gd.Int(index))))
}
func (self Simple) GlobalMenuAddIconCheckItem(menu_root string, icon [1]classdb.Texture2D, label string, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GlobalMenuAddIconCheckItem(gc.String(menu_root), icon, gc.String(label), callback, key_callback, tag, accelerator, gd.Int(index))))
}
func (self Simple) GlobalMenuAddRadioCheckItem(menu_root string, label string, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GlobalMenuAddRadioCheckItem(gc.String(menu_root), gc.String(label), callback, key_callback, tag, accelerator, gd.Int(index))))
}
func (self Simple) GlobalMenuAddIconRadioCheckItem(menu_root string, icon [1]classdb.Texture2D, label string, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GlobalMenuAddIconRadioCheckItem(gc.String(menu_root), icon, gc.String(label), callback, key_callback, tag, accelerator, gd.Int(index))))
}
func (self Simple) GlobalMenuAddMultistateItem(menu_root string, label string, max_states int, default_state int, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GlobalMenuAddMultistateItem(gc.String(menu_root), gc.String(label), gd.Int(max_states), gd.Int(default_state), callback, key_callback, tag, accelerator, gd.Int(index))))
}
func (self Simple) GlobalMenuAddSeparator(menu_root string, index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GlobalMenuAddSeparator(gc.String(menu_root), gd.Int(index))))
}
func (self Simple) GlobalMenuGetItemIndexFromText(menu_root string, text string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GlobalMenuGetItemIndexFromText(gc.String(menu_root), gc.String(text))))
}
func (self Simple) GlobalMenuGetItemIndexFromTag(menu_root string, tag gd.Variant) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GlobalMenuGetItemIndexFromTag(gc.String(menu_root), tag)))
}
func (self Simple) GlobalMenuIsItemChecked(menu_root string, idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GlobalMenuIsItemChecked(gc.String(menu_root), gd.Int(idx)))
}
func (self Simple) GlobalMenuIsItemCheckable(menu_root string, idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GlobalMenuIsItemCheckable(gc.String(menu_root), gd.Int(idx)))
}
func (self Simple) GlobalMenuIsItemRadioCheckable(menu_root string, idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GlobalMenuIsItemRadioCheckable(gc.String(menu_root), gd.Int(idx)))
}
func (self Simple) GlobalMenuGetItemCallback(menu_root string, idx int) gd.Callable {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Callable(Expert(self).GlobalMenuGetItemCallback(gc, gc.String(menu_root), gd.Int(idx)))
}
func (self Simple) GlobalMenuGetItemKeyCallback(menu_root string, idx int) gd.Callable {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Callable(Expert(self).GlobalMenuGetItemKeyCallback(gc, gc.String(menu_root), gd.Int(idx)))
}
func (self Simple) GlobalMenuGetItemTag(menu_root string, idx int) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GlobalMenuGetItemTag(gc, gc.String(menu_root), gd.Int(idx)))
}
func (self Simple) GlobalMenuGetItemText(menu_root string, idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GlobalMenuGetItemText(gc, gc.String(menu_root), gd.Int(idx)).String())
}
func (self Simple) GlobalMenuGetItemSubmenu(menu_root string, idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GlobalMenuGetItemSubmenu(gc, gc.String(menu_root), gd.Int(idx)).String())
}
func (self Simple) GlobalMenuGetItemAccelerator(menu_root string, idx int) gd.Key {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Key(Expert(self).GlobalMenuGetItemAccelerator(gc.String(menu_root), gd.Int(idx)))
}
func (self Simple) GlobalMenuIsItemDisabled(menu_root string, idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GlobalMenuIsItemDisabled(gc.String(menu_root), gd.Int(idx)))
}
func (self Simple) GlobalMenuIsItemHidden(menu_root string, idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GlobalMenuIsItemHidden(gc.String(menu_root), gd.Int(idx)))
}
func (self Simple) GlobalMenuGetItemTooltip(menu_root string, idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GlobalMenuGetItemTooltip(gc, gc.String(menu_root), gd.Int(idx)).String())
}
func (self Simple) GlobalMenuGetItemState(menu_root string, idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GlobalMenuGetItemState(gc.String(menu_root), gd.Int(idx))))
}
func (self Simple) GlobalMenuGetItemMaxStates(menu_root string, idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GlobalMenuGetItemMaxStates(gc.String(menu_root), gd.Int(idx))))
}
func (self Simple) GlobalMenuGetItemIcon(menu_root string, idx int) [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GlobalMenuGetItemIcon(gc, gc.String(menu_root), gd.Int(idx)))
}
func (self Simple) GlobalMenuGetItemIndentationLevel(menu_root string, idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GlobalMenuGetItemIndentationLevel(gc.String(menu_root), gd.Int(idx))))
}
func (self Simple) GlobalMenuSetItemChecked(menu_root string, idx int, checked bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalMenuSetItemChecked(gc.String(menu_root), gd.Int(idx), checked)
}
func (self Simple) GlobalMenuSetItemCheckable(menu_root string, idx int, checkable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalMenuSetItemCheckable(gc.String(menu_root), gd.Int(idx), checkable)
}
func (self Simple) GlobalMenuSetItemRadioCheckable(menu_root string, idx int, checkable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalMenuSetItemRadioCheckable(gc.String(menu_root), gd.Int(idx), checkable)
}
func (self Simple) GlobalMenuSetItemCallback(menu_root string, idx int, callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalMenuSetItemCallback(gc.String(menu_root), gd.Int(idx), callback)
}
func (self Simple) GlobalMenuSetItemHoverCallbacks(menu_root string, idx int, callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalMenuSetItemHoverCallbacks(gc.String(menu_root), gd.Int(idx), callback)
}
func (self Simple) GlobalMenuSetItemKeyCallback(menu_root string, idx int, key_callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalMenuSetItemKeyCallback(gc.String(menu_root), gd.Int(idx), key_callback)
}
func (self Simple) GlobalMenuSetItemTag(menu_root string, idx int, tag gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalMenuSetItemTag(gc.String(menu_root), gd.Int(idx), tag)
}
func (self Simple) GlobalMenuSetItemText(menu_root string, idx int, text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalMenuSetItemText(gc.String(menu_root), gd.Int(idx), gc.String(text))
}
func (self Simple) GlobalMenuSetItemSubmenu(menu_root string, idx int, submenu string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalMenuSetItemSubmenu(gc.String(menu_root), gd.Int(idx), gc.String(submenu))
}
func (self Simple) GlobalMenuSetItemAccelerator(menu_root string, idx int, keycode gd.Key) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalMenuSetItemAccelerator(gc.String(menu_root), gd.Int(idx), keycode)
}
func (self Simple) GlobalMenuSetItemDisabled(menu_root string, idx int, disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalMenuSetItemDisabled(gc.String(menu_root), gd.Int(idx), disabled)
}
func (self Simple) GlobalMenuSetItemHidden(menu_root string, idx int, hidden bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalMenuSetItemHidden(gc.String(menu_root), gd.Int(idx), hidden)
}
func (self Simple) GlobalMenuSetItemTooltip(menu_root string, idx int, tooltip string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalMenuSetItemTooltip(gc.String(menu_root), gd.Int(idx), gc.String(tooltip))
}
func (self Simple) GlobalMenuSetItemState(menu_root string, idx int, state int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalMenuSetItemState(gc.String(menu_root), gd.Int(idx), gd.Int(state))
}
func (self Simple) GlobalMenuSetItemMaxStates(menu_root string, idx int, max_states int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalMenuSetItemMaxStates(gc.String(menu_root), gd.Int(idx), gd.Int(max_states))
}
func (self Simple) GlobalMenuSetItemIcon(menu_root string, idx int, icon [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalMenuSetItemIcon(gc.String(menu_root), gd.Int(idx), icon)
}
func (self Simple) GlobalMenuSetItemIndentationLevel(menu_root string, idx int, level int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalMenuSetItemIndentationLevel(gc.String(menu_root), gd.Int(idx), gd.Int(level))
}
func (self Simple) GlobalMenuGetItemCount(menu_root string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GlobalMenuGetItemCount(gc.String(menu_root))))
}
func (self Simple) GlobalMenuRemoveItem(menu_root string, idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalMenuRemoveItem(gc.String(menu_root), gd.Int(idx))
}
func (self Simple) GlobalMenuClear(menu_root string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalMenuClear(gc.String(menu_root))
}
func (self Simple) GlobalMenuGetSystemMenuRoots() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).GlobalMenuGetSystemMenuRoots(gc))
}
func (self Simple) TtsIsSpeaking() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).TtsIsSpeaking())
}
func (self Simple) TtsIsPaused() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).TtsIsPaused())
}
func (self Simple) TtsGetVoices() gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Dictionary](Expert(self).TtsGetVoices(gc))
}
func (self Simple) TtsGetVoicesForLanguage(language string) gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).TtsGetVoicesForLanguage(gc, gc.String(language)))
}
func (self Simple) TtsSpeak(text string, voice string, volume int, pitch float64, rate float64, utterance_id int, interrupt bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).TtsSpeak(gc.String(text), gc.String(voice), gd.Int(volume), gd.Float(pitch), gd.Float(rate), gd.Int(utterance_id), interrupt)
}
func (self Simple) TtsPause() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).TtsPause()
}
func (self Simple) TtsResume() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).TtsResume()
}
func (self Simple) TtsStop() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).TtsStop()
}
func (self Simple) TtsSetUtteranceCallback(event classdb.DisplayServerTTSUtteranceEvent, callable gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).TtsSetUtteranceCallback(event, callable)
}
func (self Simple) IsDarkModeSupported() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDarkModeSupported())
}
func (self Simple) IsDarkMode() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDarkMode())
}
func (self Simple) GetAccentColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetAccentColor())
}
func (self Simple) GetBaseColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetBaseColor())
}
func (self Simple) SetSystemThemeChangeCallback(callable gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSystemThemeChangeCallback(callable)
}
func (self Simple) MouseSetMode(mouse_mode classdb.DisplayServerMouseMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MouseSetMode(mouse_mode)
}
func (self Simple) MouseGetMode() classdb.DisplayServerMouseMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.DisplayServerMouseMode(Expert(self).MouseGetMode())
}
func (self Simple) WarpMouse(position gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WarpMouse(position)
}
func (self Simple) MouseGetPosition() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).MouseGetPosition())
}
func (self Simple) MouseGetButtonState() gd.MouseButtonMask {
	gc := gd.GarbageCollector(); _ = gc
	return gd.MouseButtonMask(Expert(self).MouseGetButtonState())
}
func (self Simple) ClipboardSet(clipboard string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClipboardSet(gc.String(clipboard))
}
func (self Simple) ClipboardGet() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).ClipboardGet(gc).String())
}
func (self Simple) ClipboardGetImage() [1]classdb.Image {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Image(Expert(self).ClipboardGetImage(gc))
}
func (self Simple) ClipboardHas() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).ClipboardHas())
}
func (self Simple) ClipboardHasImage() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).ClipboardHasImage())
}
func (self Simple) ClipboardSetPrimary(clipboard_primary string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClipboardSetPrimary(gc.String(clipboard_primary))
}
func (self Simple) ClipboardGetPrimary() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).ClipboardGetPrimary(gc).String())
}
func (self Simple) GetDisplayCutouts() gd.ArrayOf[gd.Rect2] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Rect2](Expert(self).GetDisplayCutouts(gc))
}
func (self Simple) GetDisplaySafeArea() gd.Rect2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2i(Expert(self).GetDisplaySafeArea())
}
func (self Simple) GetScreenCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetScreenCount()))
}
func (self Simple) GetPrimaryScreen() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPrimaryScreen()))
}
func (self Simple) GetKeyboardFocusScreen() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetKeyboardFocusScreen()))
}
func (self Simple) GetScreenFromRect(rect gd.Rect2) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetScreenFromRect(rect)))
}
func (self Simple) ScreenGetPosition(screen int) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).ScreenGetPosition(gd.Int(screen)))
}
func (self Simple) ScreenGetSize(screen int) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).ScreenGetSize(gd.Int(screen)))
}
func (self Simple) ScreenGetUsableRect(screen int) gd.Rect2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2i(Expert(self).ScreenGetUsableRect(gd.Int(screen)))
}
func (self Simple) ScreenGetDpi(screen int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).ScreenGetDpi(gd.Int(screen))))
}
func (self Simple) ScreenGetScale(screen int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).ScreenGetScale(gd.Int(screen))))
}
func (self Simple) IsTouchscreenAvailable() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsTouchscreenAvailable())
}
func (self Simple) ScreenGetMaxScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).ScreenGetMaxScale()))
}
func (self Simple) ScreenGetRefreshRate(screen int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).ScreenGetRefreshRate(gd.Int(screen))))
}
func (self Simple) ScreenGetPixel(position gd.Vector2i) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).ScreenGetPixel(position))
}
func (self Simple) ScreenGetImage(screen int) [1]classdb.Image {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Image(Expert(self).ScreenGetImage(gc, gd.Int(screen)))
}
func (self Simple) ScreenSetOrientation(orientation classdb.DisplayServerScreenOrientation, screen int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ScreenSetOrientation(orientation, gd.Int(screen))
}
func (self Simple) ScreenGetOrientation(screen int) classdb.DisplayServerScreenOrientation {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.DisplayServerScreenOrientation(Expert(self).ScreenGetOrientation(gd.Int(screen)))
}
func (self Simple) ScreenSetKeepOn(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ScreenSetKeepOn(enable)
}
func (self Simple) ScreenIsKeptOn() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).ScreenIsKeptOn())
}
func (self Simple) GetWindowList() gd.PackedInt32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt32Array(Expert(self).GetWindowList(gc))
}
func (self Simple) GetWindowAtScreenPosition(position gd.Vector2i) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetWindowAtScreenPosition(position)))
}
func (self Simple) WindowGetNativeHandle(handle_type classdb.DisplayServerHandleType, window_id int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).WindowGetNativeHandle(handle_type, gd.Int(window_id))))
}
func (self Simple) WindowGetActivePopup() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).WindowGetActivePopup()))
}
func (self Simple) WindowSetPopupSafeRect(window int, rect gd.Rect2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowSetPopupSafeRect(gd.Int(window), rect)
}
func (self Simple) WindowGetPopupSafeRect(window int) gd.Rect2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2i(Expert(self).WindowGetPopupSafeRect(gd.Int(window)))
}
func (self Simple) WindowSetTitle(title string, window_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowSetTitle(gc.String(title), gd.Int(window_id))
}
func (self Simple) WindowGetTitleSize(title string, window_id int) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).WindowGetTitleSize(gc.String(title), gd.Int(window_id)))
}
func (self Simple) WindowSetMousePassthrough(region gd.PackedVector2Array, window_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowSetMousePassthrough(region, gd.Int(window_id))
}
func (self Simple) WindowGetCurrentScreen(window_id int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).WindowGetCurrentScreen(gd.Int(window_id))))
}
func (self Simple) WindowSetCurrentScreen(screen int, window_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowSetCurrentScreen(gd.Int(screen), gd.Int(window_id))
}
func (self Simple) WindowGetPosition(window_id int) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).WindowGetPosition(gd.Int(window_id)))
}
func (self Simple) WindowGetPositionWithDecorations(window_id int) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).WindowGetPositionWithDecorations(gd.Int(window_id)))
}
func (self Simple) WindowSetPosition(position gd.Vector2i, window_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowSetPosition(position, gd.Int(window_id))
}
func (self Simple) WindowGetSize(window_id int) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).WindowGetSize(gd.Int(window_id)))
}
func (self Simple) WindowSetSize(size gd.Vector2i, window_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowSetSize(size, gd.Int(window_id))
}
func (self Simple) WindowSetRectChangedCallback(callback gd.Callable, window_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowSetRectChangedCallback(callback, gd.Int(window_id))
}
func (self Simple) WindowSetWindowEventCallback(callback gd.Callable, window_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowSetWindowEventCallback(callback, gd.Int(window_id))
}
func (self Simple) WindowSetInputEventCallback(callback gd.Callable, window_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowSetInputEventCallback(callback, gd.Int(window_id))
}
func (self Simple) WindowSetInputTextCallback(callback gd.Callable, window_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowSetInputTextCallback(callback, gd.Int(window_id))
}
func (self Simple) WindowSetDropFilesCallback(callback gd.Callable, window_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowSetDropFilesCallback(callback, gd.Int(window_id))
}
func (self Simple) WindowGetAttachedInstanceId(window_id int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).WindowGetAttachedInstanceId(gd.Int(window_id))))
}
func (self Simple) WindowGetMaxSize(window_id int) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).WindowGetMaxSize(gd.Int(window_id)))
}
func (self Simple) WindowSetMaxSize(max_size gd.Vector2i, window_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowSetMaxSize(max_size, gd.Int(window_id))
}
func (self Simple) WindowGetMinSize(window_id int) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).WindowGetMinSize(gd.Int(window_id)))
}
func (self Simple) WindowSetMinSize(min_size gd.Vector2i, window_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowSetMinSize(min_size, gd.Int(window_id))
}
func (self Simple) WindowGetSizeWithDecorations(window_id int) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).WindowGetSizeWithDecorations(gd.Int(window_id)))
}
func (self Simple) WindowGetMode(window_id int) classdb.DisplayServerWindowMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.DisplayServerWindowMode(Expert(self).WindowGetMode(gd.Int(window_id)))
}
func (self Simple) WindowSetMode(mode classdb.DisplayServerWindowMode, window_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowSetMode(mode, gd.Int(window_id))
}
func (self Simple) WindowSetFlag(flag classdb.DisplayServerWindowFlags, enabled bool, window_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowSetFlag(flag, enabled, gd.Int(window_id))
}
func (self Simple) WindowGetFlag(flag classdb.DisplayServerWindowFlags, window_id int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).WindowGetFlag(flag, gd.Int(window_id)))
}
func (self Simple) WindowSetWindowButtonsOffset(offset gd.Vector2i, window_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowSetWindowButtonsOffset(offset, gd.Int(window_id))
}
func (self Simple) WindowGetSafeTitleMargins(window_id int) gd.Vector3i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3i(Expert(self).WindowGetSafeTitleMargins(gd.Int(window_id)))
}
func (self Simple) WindowRequestAttention(window_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowRequestAttention(gd.Int(window_id))
}
func (self Simple) WindowMoveToForeground(window_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowMoveToForeground(gd.Int(window_id))
}
func (self Simple) WindowIsFocused(window_id int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).WindowIsFocused(gd.Int(window_id)))
}
func (self Simple) WindowCanDraw(window_id int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).WindowCanDraw(gd.Int(window_id)))
}
func (self Simple) WindowSetTransient(window_id int, parent_window_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowSetTransient(gd.Int(window_id), gd.Int(parent_window_id))
}
func (self Simple) WindowSetExclusive(window_id int, exclusive bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowSetExclusive(gd.Int(window_id), exclusive)
}
func (self Simple) WindowSetImeActive(active bool, window_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowSetImeActive(active, gd.Int(window_id))
}
func (self Simple) WindowSetImePosition(position gd.Vector2i, window_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowSetImePosition(position, gd.Int(window_id))
}
func (self Simple) WindowSetVsyncMode(vsync_mode classdb.DisplayServerVSyncMode, window_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WindowSetVsyncMode(vsync_mode, gd.Int(window_id))
}
func (self Simple) WindowGetVsyncMode(window_id int) classdb.DisplayServerVSyncMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.DisplayServerVSyncMode(Expert(self).WindowGetVsyncMode(gd.Int(window_id)))
}
func (self Simple) WindowIsMaximizeAllowed(window_id int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).WindowIsMaximizeAllowed(gd.Int(window_id)))
}
func (self Simple) WindowMaximizeOnTitleDblClick() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).WindowMaximizeOnTitleDblClick())
}
func (self Simple) WindowMinimizeOnTitleDblClick() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).WindowMinimizeOnTitleDblClick())
}
func (self Simple) ImeGetSelection() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).ImeGetSelection())
}
func (self Simple) ImeGetText() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).ImeGetText(gc).String())
}
func (self Simple) VirtualKeyboardShow(existing_text string, position gd.Rect2, atype classdb.DisplayServerVirtualKeyboardType, max_length int, cursor_start int, cursor_end int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).VirtualKeyboardShow(gc.String(existing_text), position, atype, gd.Int(max_length), gd.Int(cursor_start), gd.Int(cursor_end))
}
func (self Simple) VirtualKeyboardHide() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).VirtualKeyboardHide()
}
func (self Simple) VirtualKeyboardGetHeight() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).VirtualKeyboardGetHeight()))
}
func (self Simple) CursorSetShape(shape classdb.DisplayServerCursorShape) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CursorSetShape(shape)
}
func (self Simple) CursorGetShape() classdb.DisplayServerCursorShape {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.DisplayServerCursorShape(Expert(self).CursorGetShape())
}
func (self Simple) CursorSetCustomImage(cursor [1]classdb.Resource, shape classdb.DisplayServerCursorShape, hotspot gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CursorSetCustomImage(cursor, shape, hotspot)
}
func (self Simple) GetSwapCancelOk() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetSwapCancelOk())
}
func (self Simple) EnableForStealingFocus(process_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnableForStealingFocus(gd.Int(process_id))
}
func (self Simple) DialogShow(title string, description string, buttons gd.PackedStringArray, callback gd.Callable) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).DialogShow(gc.String(title), gc.String(description), buttons, callback))
}
func (self Simple) DialogInputText(title string, description string, existing_text string, callback gd.Callable) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).DialogInputText(gc.String(title), gc.String(description), gc.String(existing_text), callback))
}
func (self Simple) FileDialogShow(title string, current_directory string, filename string, show_hidden bool, mode classdb.DisplayServerFileDialogMode, filters gd.PackedStringArray, callback gd.Callable) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).FileDialogShow(gc.String(title), gc.String(current_directory), gc.String(filename), show_hidden, mode, filters, callback))
}
func (self Simple) FileDialogWithOptionsShow(title string, current_directory string, root string, filename string, show_hidden bool, mode classdb.DisplayServerFileDialogMode, filters gd.PackedStringArray, options gd.ArrayOf[gd.Dictionary], callback gd.Callable) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).FileDialogWithOptionsShow(gc.String(title), gc.String(current_directory), gc.String(root), gc.String(filename), show_hidden, mode, filters, options, callback))
}
func (self Simple) KeyboardGetLayoutCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).KeyboardGetLayoutCount()))
}
func (self Simple) KeyboardGetCurrentLayout() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).KeyboardGetCurrentLayout()))
}
func (self Simple) KeyboardSetCurrentLayout(index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).KeyboardSetCurrentLayout(gd.Int(index))
}
func (self Simple) KeyboardGetLayoutLanguage(index int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).KeyboardGetLayoutLanguage(gc, gd.Int(index)).String())
}
func (self Simple) KeyboardGetLayoutName(index int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).KeyboardGetLayoutName(gc, gd.Int(index)).String())
}
func (self Simple) KeyboardGetKeycodeFromPhysical(keycode gd.Key) gd.Key {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Key(Expert(self).KeyboardGetKeycodeFromPhysical(keycode))
}
func (self Simple) KeyboardGetLabelFromPhysical(keycode gd.Key) gd.Key {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Key(Expert(self).KeyboardGetLabelFromPhysical(keycode))
}
func (self Simple) ProcessEvents() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ProcessEvents()
}
func (self Simple) ForceProcessAndDropEvents() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ForceProcessAndDropEvents()
}
func (self Simple) SetNativeIcon(filename string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNativeIcon(gc.String(filename))
}
func (self Simple) SetIcon(image [1]classdb.Image) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetIcon(image)
}
func (self Simple) CreateStatusIndicator(icon [1]classdb.Texture2D, tooltip string, callback gd.Callable) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).CreateStatusIndicator(icon, gc.String(tooltip), callback)))
}
func (self Simple) StatusIndicatorSetIcon(id int, icon [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).StatusIndicatorSetIcon(gd.Int(id), icon)
}
func (self Simple) StatusIndicatorSetTooltip(id int, tooltip string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).StatusIndicatorSetTooltip(gd.Int(id), gc.String(tooltip))
}
func (self Simple) StatusIndicatorSetMenu(id int, menu_rid gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).StatusIndicatorSetMenu(gd.Int(id), menu_rid)
}
func (self Simple) StatusIndicatorSetCallback(id int, callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).StatusIndicatorSetCallback(gd.Int(id), callback)
}
func (self Simple) StatusIndicatorGetRect(id int) gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(Expert(self).StatusIndicatorGetRect(gd.Int(id)))
}
func (self Simple) DeleteStatusIndicator(id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DeleteStatusIndicator(gd.Int(id))
}
func (self Simple) TabletGetDriverCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).TabletGetDriverCount()))
}
func (self Simple) TabletGetDriverName(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).TabletGetDriverName(gc, gd.Int(idx)).String())
}
func (self Simple) TabletGetCurrentDriver() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).TabletGetCurrentDriver(gc).String())
}
func (self Simple) TabletSetCurrentDriver(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).TabletSetCurrentDriver(gc.String(name))
}
func (self Simple) IsWindowTransparencyAvailable() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsWindowTransparencyAvailable())
}
func (self Simple) RegisterAdditionalOutput(obj gd.Object) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RegisterAdditionalOutput(obj)
}
func (self Simple) UnregisterAdditionalOutput(obj gd.Object) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UnregisterAdditionalOutput(obj)
}
func (self Simple) HasAdditionalOutputs() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasAdditionalOutputs())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.DisplayServer
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns [code]true[/code] if the specified [param feature] is supported by the current [DisplayServer], [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) HasFeature(feature classdb.DisplayServerFeature) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, feature)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_has_feature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the name of the [DisplayServer] currently in use. Most operating systems only have a single [DisplayServer], but Linux has access to more than one [DisplayServer] (currently X11 and Wayland).
The names of built-in display servers are [code]Windows[/code], [code]macOS[/code], [code]X11[/code] (Linux), [code]Wayland[/code] (Linux), [code]Android[/code], [code]iOS[/code], [code]web[/code] (HTML5), and [code]headless[/code] (when started with the [code]--headless[/code] [url=$DOCS_URL/tutorials/editor/command_line_tutorial.html]command line argument[/url]).
*/
//go:nosplit
func (self class) GetName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_get_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets native help system search callbacks.
[param search_callback] has the following arguments: [code]String search_string, int result_limit[/code] and return a [Dictionary] with "key, display name" pairs for the search results. Called when the user enters search terms in the [code]Help[/code] menu.
[param action_callback] has the following arguments: [code]String key[/code]. Called when the user selects a search result in the [code]Help[/code] menu.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) HelpSetSearchCallbacks(search_callback gd.Callable, action_callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(search_callback))
	callframe.Arg(frame, mmm.Get(action_callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_help_set_search_callbacks, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Registers callables to emit when the menu is respectively about to show or closed. Callback methods should have zero arguments.
*/
//go:nosplit
func (self class) GlobalMenuSetPopupCallbacks(menu_root gd.String, open_callback gd.Callable, close_callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, mmm.Get(open_callback))
	callframe.Arg(frame, mmm.Get(close_callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_set_popup_callbacks, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds an item that will act as a submenu of the global menu [param menu_root]. The [param submenu] argument is the ID of the global menu root that will be shown when the item is clicked.
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
//go:nosplit
func (self class) GlobalMenuAddSubmenuItem(menu_root gd.String, label gd.String, submenu gd.String, index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, mmm.Get(submenu))
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_add_submenu_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a new item with text [param label] to the global menu with ID [param menu_root].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
//go:nosplit
func (self class) GlobalMenuAddItem(menu_root gd.String, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, mmm.Get(callback))
	callframe.Arg(frame, mmm.Get(key_callback))
	callframe.Arg(frame, mmm.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_add_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a new checkable item with text [param label] to the global menu with ID [param menu_root].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
//go:nosplit
func (self class) GlobalMenuAddCheckItem(menu_root gd.String, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, mmm.Get(callback))
	callframe.Arg(frame, mmm.Get(key_callback))
	callframe.Arg(frame, mmm.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_add_check_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a new item with text [param label] and icon [param icon] to the global menu with ID [param menu_root].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
//go:nosplit
func (self class) GlobalMenuAddIconItem(menu_root gd.String, icon object.Texture2D, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, mmm.Get(callback))
	callframe.Arg(frame, mmm.Get(key_callback))
	callframe.Arg(frame, mmm.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_add_icon_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a new checkable item with text [param label] and icon [param icon] to the global menu with ID [param menu_root].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
//go:nosplit
func (self class) GlobalMenuAddIconCheckItem(menu_root gd.String, icon object.Texture2D, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, mmm.Get(callback))
	callframe.Arg(frame, mmm.Get(key_callback))
	callframe.Arg(frame, mmm.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_add_icon_check_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a new radio-checkable item with text [param label] to the global menu with ID [param menu_root].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] Radio-checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method global_menu_set_item_checked] for more info on how to control it.
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
//go:nosplit
func (self class) GlobalMenuAddRadioCheckItem(menu_root gd.String, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, mmm.Get(callback))
	callframe.Arg(frame, mmm.Get(key_callback))
	callframe.Arg(frame, mmm.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_add_radio_check_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a new radio-checkable item with text [param label] and icon [param icon] to the global menu with ID [param menu_root].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] Radio-checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method global_menu_set_item_checked] for more info on how to control it.
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
//go:nosplit
func (self class) GlobalMenuAddIconRadioCheckItem(menu_root gd.String, icon object.Texture2D, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, mmm.Get(callback))
	callframe.Arg(frame, mmm.Get(key_callback))
	callframe.Arg(frame, mmm.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_add_icon_radio_check_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a new item with text [param label] to the global menu with ID [param menu_root].
Contrarily to normal binary items, multistate items can have more than two states, as defined by [param max_states]. Each press or activate of the item will increase the state by one. The default value is defined by [param default_state].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] By default, there's no indication of the current item state, it should be changed manually.
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
//go:nosplit
func (self class) GlobalMenuAddMultistateItem(menu_root gd.String, label gd.String, max_states gd.Int, default_state gd.Int, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, max_states)
	callframe.Arg(frame, default_state)
	callframe.Arg(frame, mmm.Get(callback))
	callframe.Arg(frame, mmm.Get(key_callback))
	callframe.Arg(frame, mmm.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_add_multistate_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a separator between items to the global menu with ID [param menu_root]. Separators also occupy an index.
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
//go:nosplit
func (self class) GlobalMenuAddSeparator(menu_root gd.String, index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_add_separator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the index of the item with the specified [param text]. Indices are automatically assigned to each item by the engine, and cannot be set manually.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemIndexFromText(menu_root gd.String, text gd.String) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, mmm.Get(text))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_get_item_index_from_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the index of the item with the specified [param tag]. Indices are automatically assigned to each item by the engine, and cannot be set manually.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemIndexFromTag(menu_root gd.String, tag gd.Variant) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, mmm.Get(tag))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_get_item_index_from_tag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the item at index [param idx] is checked.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuIsItemChecked(menu_root gd.String, idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_is_item_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the item at index [param idx] is checkable in some way, i.e. if it has a checkbox or radio button.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuIsItemCheckable(menu_root gd.String, idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_is_item_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the item at index [param idx] has radio button-style checkability.
[b]Note:[/b] This is purely cosmetic; you must add the logic for checking/unchecking items in radio groups.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuIsItemRadioCheckable(menu_root gd.String, idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_is_item_radio_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the callback of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemCallback(ctx gd.Lifetime, menu_root gd.String, idx gd.Int) gd.Callable {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_get_item_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Callable](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the callback of the item accelerator at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemKeyCallback(ctx gd.Lifetime, menu_root gd.String, idx gd.Int) gd.Callable {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_get_item_key_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Callable](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the metadata of the specified item, which might be of any type. You can set it with [method global_menu_set_item_tag], which provides a simple way of assigning context data to items.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemTag(ctx gd.Lifetime, menu_root gd.String, idx gd.Int) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_get_item_tag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the text of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemText(ctx gd.Lifetime, menu_root gd.String, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_get_item_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the submenu ID of the item at index [param idx]. See [method global_menu_add_submenu_item] for more info on how to add a submenu.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemSubmenu(ctx gd.Lifetime, menu_root gd.String, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_get_item_submenu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the accelerator of the item at index [param idx]. Accelerators are special combinations of keys that activate the item, no matter which control is focused.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemAccelerator(menu_root gd.String, idx gd.Int) gd.Key {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Key](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_get_item_accelerator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the item at index [param idx] is disabled. When it is disabled it can't be selected, or its action invoked.
See [method global_menu_set_item_disabled] for more info on how to disable an item.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuIsItemDisabled(menu_root gd.String, idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_is_item_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the item at index [param idx] is hidden.
See [method global_menu_set_item_hidden] for more info on how to hide an item.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuIsItemHidden(menu_root gd.String, idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_is_item_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the tooltip associated with the specified index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemTooltip(ctx gd.Lifetime, menu_root gd.String, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_get_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the state of a multistate item. See [method global_menu_add_multistate_item] for details.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemState(menu_root gd.String, idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_get_item_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns number of states of a multistate item. See [method global_menu_add_multistate_item] for details.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemMaxStates(menu_root gd.String, idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_get_item_max_states, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the icon of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemIcon(ctx gd.Lifetime, menu_root gd.String, idx gd.Int) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_get_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the horizontal offset of the item at the given [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemIndentationLevel(menu_root gd.String, idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_get_item_indentation_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the checkstate status of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemChecked(menu_root gd.String, idx gd.Int, checked bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, checked)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_set_item_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets whether the item at index [param idx] has a checkbox. If [code]false[/code], sets the type of the item to plain text.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemCheckable(menu_root gd.String, idx gd.Int, checkable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, checkable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_set_item_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the type of the item at the specified index [param idx] to radio button. If [code]false[/code], sets the type of the item to plain text.
[b]Note:[/b] This is purely cosmetic; you must add the logic for checking/unchecking items in radio groups.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemRadioCheckable(menu_root gd.String, idx gd.Int, checkable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, checkable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_set_item_radio_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the callback of the item at index [param idx]. Callback is emitted when an item is pressed.
[b]Note:[/b] The [param callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemCallback(menu_root gd.String, idx gd.Int, callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_set_item_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the callback of the item at index [param idx]. The callback is emitted when an item is hovered.
[b]Note:[/b] The [param callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemHoverCallbacks(menu_root gd.String, idx gd.Int, callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_set_item_hover_callbacks, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the callback of the item at index [param idx]. Callback is emitted when its accelerator is activated.
[b]Note:[/b] The [param key_callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemKeyCallback(menu_root gd.String, idx gd.Int, key_callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(key_callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_set_item_key_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the metadata of an item, which may be of any type. You can later get it with [method global_menu_get_item_tag], which provides a simple way of assigning context data to items.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemTag(menu_root gd.String, idx gd.Int, tag gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(tag))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_set_item_tag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the text of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemText(menu_root gd.String, idx gd.Int, text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_set_item_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the submenu of the item at index [param idx]. The submenu is the ID of a global menu root that would be shown when the item is clicked.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemSubmenu(menu_root gd.String, idx gd.Int, submenu gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(submenu))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_set_item_submenu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the accelerator of the item at index [param idx]. [param keycode] can be a single [enum Key], or a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemAccelerator(menu_root gd.String, idx gd.Int, keycode gd.Key)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, keycode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_set_item_accelerator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Enables/disables the item at index [param idx]. When it is disabled, it can't be selected and its action can't be invoked.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemDisabled(menu_root gd.String, idx gd.Int, disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_set_item_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Hides/shows the item at index [param idx]. When it is hidden, an item does not appear in a menu and its action cannot be invoked.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemHidden(menu_root gd.String, idx gd.Int, hidden bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, hidden)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_set_item_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [String] tooltip of the item at the specified index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemTooltip(menu_root gd.String, idx gd.Int, tooltip gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(tooltip))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_set_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the state of a multistate item. See [method global_menu_add_multistate_item] for details.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemState(menu_root gd.String, idx gd.Int, state gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, state)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_set_item_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets number of state of a multistate item. See [method global_menu_add_multistate_item] for details.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemMaxStates(menu_root gd.String, idx gd.Int, max_states gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, max_states)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_set_item_max_states, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Replaces the [Texture2D] icon of the specified [param idx].
[b]Note:[/b] This method is implemented only on macOS.
[b]Note:[/b] This method is not supported by macOS "_dock" menu items.
*/
//go:nosplit
func (self class) GlobalMenuSetItemIcon(menu_root gd.String, idx gd.Int, icon object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_set_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the horizontal offset of the item at the given [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemIndentationLevel(menu_root gd.String, idx gd.Int, level gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, level)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_set_item_indentation_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns number of items in the global menu with ID [param menu_root].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemCount(menu_root gd.String) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_get_item_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the item at index [param idx] from the global menu [param menu_root].
[b]Note:[/b] The indices of items after the removed item will be shifted by one.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuRemoveItem(menu_root gd.String, idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_remove_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all items from the global menu with ID [param menu_root].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
//go:nosplit
func (self class) GlobalMenuClear(menu_root gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu_root))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns Dictionary of supported system menu IDs and names.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetSystemMenuRoots(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_global_menu_get_system_menu_roots, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the synthesizer is generating speech, or have utterance waiting in the queue.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
//go:nosplit
func (self class) TtsIsSpeaking() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_tts_is_speaking, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the synthesizer is in a paused state.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
//go:nosplit
func (self class) TtsIsPaused() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_tts_is_paused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an [Array] of voice information dictionaries.
Each [Dictionary] contains two [String] entries:
- [code]name[/code] is voice name.
- [code]id[/code] is voice identifier.
- [code]language[/code] is language code in [code]lang_Variant[/code] format. The [code]lang[/code] part is a 2 or 3-letter code based on the ISO-639 standard, in lowercase. The [code skip-lint]Variant[/code] part is an engine-dependent string describing country, region or/and dialect.
Note that Godot depends on system libraries for text-to-speech functionality. These libraries are installed by default on Windows and macOS, but not on all Linux distributions. If they are not present, this method will return an empty list. This applies to both Godot users on Linux, as well as end-users on Linux running Godot games that use text-to-speech.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
//go:nosplit
func (self class) TtsGetVoices(ctx gd.Lifetime) gd.ArrayOf[gd.Dictionary] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_tts_get_voices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Dictionary](ret)
}
/*
Returns an [PackedStringArray] of voice identifiers for the [param language].
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
//go:nosplit
func (self class) TtsGetVoicesForLanguage(ctx gd.Lifetime, language gd.String) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(language))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_tts_get_voices_for_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Adds an utterance to the queue. If [param interrupt] is [code]true[/code], the queue is cleared first.
- [param voice] identifier is one of the [code]"id"[/code] values returned by [method tts_get_voices] or one of the values returned by [method tts_get_voices_for_language].
- [param volume] ranges from [code]0[/code] (lowest) to [code]100[/code] (highest).
- [param pitch] ranges from [code]0.0[/code] (lowest) to [code]2.0[/code] (highest), [code]1.0[/code] is default pitch for the current voice.
- [param rate] ranges from [code]0.1[/code] (lowest) to [code]10.0[/code] (highest), [code]1.0[/code] is a normal speaking rate. Other values act as a percentage relative.
- [param utterance_id] is passed as a parameter to the callback functions.
[b]Note:[/b] On Windows and Linux (X11/Wayland), utterance [param text] can use SSML markup. SSML support is engine and voice dependent. If the engine does not support SSML, you should strip out all XML markup before calling [method tts_speak].
[b]Note:[/b] The granularity of pitch, rate, and volume is engine and voice dependent. Values may be truncated.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
//go:nosplit
func (self class) TtsSpeak(text gd.String, voice gd.String, volume gd.Int, pitch gd.Float, rate gd.Float, utterance_id gd.Int, interrupt bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, mmm.Get(voice))
	callframe.Arg(frame, volume)
	callframe.Arg(frame, pitch)
	callframe.Arg(frame, rate)
	callframe.Arg(frame, utterance_id)
	callframe.Arg(frame, interrupt)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_tts_speak, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Puts the synthesizer into a paused state.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
//go:nosplit
func (self class) TtsPause()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_tts_pause, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Resumes the synthesizer if it was paused.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
//go:nosplit
func (self class) TtsResume()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_tts_resume, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stops synthesis in progress and removes all utterances from the queue.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Linux), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
//go:nosplit
func (self class) TtsStop()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_tts_stop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a callback, which is called when the utterance has started, finished, canceled or reached a text boundary.
- [constant TTS_UTTERANCE_STARTED], [constant TTS_UTTERANCE_ENDED], and [constant TTS_UTTERANCE_CANCELED] callable's method should take one [int] parameter, the utterance ID.
- [constant TTS_UTTERANCE_BOUNDARY] callable's method should take two [int] parameters, the index of the character and the utterance ID.
[b]Note:[/b] The granularity of the boundary callbacks is engine dependent.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
//go:nosplit
func (self class) TtsSetUtteranceCallback(event classdb.DisplayServerTTSUtteranceEvent, callable gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, event)
	callframe.Arg(frame, mmm.Get(callable))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_tts_set_utterance_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if OS supports dark mode.
[b]Note:[/b] This method is implemented on Android, iOS, macOS, Windows, and Linux (X11/Wayland).
*/
//go:nosplit
func (self class) IsDarkModeSupported() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_is_dark_mode_supported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if OS is using dark mode.
[b]Note:[/b] This method is implemented on Android, iOS, macOS, Windows, and Linux (X11/Wayland).
*/
//go:nosplit
func (self class) IsDarkMode() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_is_dark_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns OS theme accent color. Returns [code]Color(0, 0, 0, 0)[/code], if accent color is unknown.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) GetAccentColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_get_accent_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the OS theme base color (default control background). Returns [code]Color(0, 0, 0, 0)[/code] if the base color is unknown.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) GetBaseColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_get_base_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [param callable] that should be called when system theme settings are changed. Callback method should have zero arguments.
[b]Note:[/b] This method is implemented on Android, iOS, macOS, Windows, and Linux (X11/Wayland).
*/
//go:nosplit
func (self class) SetSystemThemeChangeCallback(callable gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(callable))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_set_system_theme_change_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the current mouse mode. See also [method mouse_get_mode].
*/
//go:nosplit
func (self class) MouseSetMode(mouse_mode classdb.DisplayServerMouseMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mouse_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_mouse_set_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the current mouse mode. See also [method mouse_set_mode].
*/
//go:nosplit
func (self class) MouseGetMode() classdb.DisplayServerMouseMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.DisplayServerMouseMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_mouse_get_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the mouse cursor position to the given [param position] relative to an origin at the upper left corner of the currently focused game Window Manager window.
[b]Note:[/b] [method warp_mouse] is only supported on Windows, macOS, and Linux (X11/Wayland). It has no effect on Android, iOS, and Web.
*/
//go:nosplit
func (self class) WarpMouse(position gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_warp_mouse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the mouse cursor's current position in screen coordinates.
*/
//go:nosplit
func (self class) MouseGetPosition() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_mouse_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the current state of mouse buttons (whether each button is pressed) as a bitmask. If multiple mouse buttons are pressed at the same time, the bits are added together. Equivalent to [method Input.get_mouse_button_mask].
*/
//go:nosplit
func (self class) MouseGetButtonState() gd.MouseButtonMask {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.MouseButtonMask](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_mouse_get_button_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the user's clipboard content to the given string.
*/
//go:nosplit
func (self class) ClipboardSet(clipboard gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(clipboard))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_clipboard_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the user's clipboard as a string if possible.
*/
//go:nosplit
func (self class) ClipboardGet(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_clipboard_get, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the user's clipboard as an image if possible.
[b]Note:[/b] This method uses the copied pixel data, e.g. from a image editing software or a web browser, not an image file copied from file explorer.
*/
//go:nosplit
func (self class) ClipboardGetImage(ctx gd.Lifetime) object.Image {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_clipboard_get_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a text content on the user's clipboard.
*/
//go:nosplit
func (self class) ClipboardHas() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_clipboard_has, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is an image content on the user's clipboard.
*/
//go:nosplit
func (self class) ClipboardHasImage() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_clipboard_has_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the user's [url=https://unix.stackexchange.com/questions/139191/whats-the-difference-between-primary-selection-and-clipboard-buffer]primary[/url] clipboard content to the given string. This is the clipboard that is set when the user selects text in any application, rather than when pressing [kbd]Ctrl + C[/kbd]. The clipboard data can then be pasted by clicking the middle mouse button in any application that supports the primary clipboard mechanism.
[b]Note:[/b] This method is only implemented on Linux (X11/Wayland).
*/
//go:nosplit
func (self class) ClipboardSetPrimary(clipboard_primary gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(clipboard_primary))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_clipboard_set_primary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the user's [url=https://unix.stackexchange.com/questions/139191/whats-the-difference-between-primary-selection-and-clipboard-buffer]primary[/url] clipboard as a string if possible. This is the clipboard that is set when the user selects text in any application, rather than when pressing [kbd]Ctrl + C[/kbd]. The clipboard data can then be pasted by clicking the middle mouse button in any application that supports the primary clipboard mechanism.
[b]Note:[/b] This method is only implemented on Linux (X11/Wayland).
*/
//go:nosplit
func (self class) ClipboardGetPrimary(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_clipboard_get_primary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an [Array] of [Rect2], each of which is the bounding rectangle for a display cutout or notch. These are non-functional areas on edge-to-edge screens used by cameras and sensors. Returns an empty array if the device does not have cutouts. See also [method get_display_safe_area].
[b]Note:[/b] Currently only implemented on Android. Other platforms will return an empty array even if they do have display cutouts or notches.
*/
//go:nosplit
func (self class) GetDisplayCutouts(ctx gd.Lifetime) gd.ArrayOf[gd.Rect2] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_get_display_cutouts, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Rect2](ret)
}
/*
Returns the unobscured area of the display where interactive controls should be rendered. See also [method get_display_cutouts].
*/
//go:nosplit
func (self class) GetDisplaySafeArea() gd.Rect2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_get_display_safe_area, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of displays available.
*/
//go:nosplit
func (self class) GetScreenCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_get_screen_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns index of the primary screen.
*/
//go:nosplit
func (self class) GetPrimaryScreen() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_get_primary_screen, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the index of the screen containing the window with the keyboard focus, or the primary screen if there's no focused window.
*/
//go:nosplit
func (self class) GetKeyboardFocusScreen() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_get_keyboard_focus_screen, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns index of the screen which contains specified rectangle.
*/
//go:nosplit
func (self class) GetScreenFromRect(rect gd.Rect2) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_get_screen_from_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the screen's top-left corner position in pixels. On multi-monitor setups, the screen position is relative to the virtual desktop area. On multi-monitor setups with different screen resolutions or orientations, the origin may be located outside any display like this:
[codeblock lang=text]
* (0, 0)        +-------+
                |       |
+-------------+ |       |
|             | |       |
|             | |       |
+-------------+ +-------+
[/codeblock]
See also [method screen_get_size].
[b]Note:[/b] On Linux (Wayland) this method always returns [code](0, 0)[/code].
*/
//go:nosplit
func (self class) ScreenGetPosition(screen gd.Int) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_screen_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the screen's size in pixels. See also [method screen_get_position] and [method screen_get_usable_rect].
*/
//go:nosplit
func (self class) ScreenGetSize(screen gd.Int) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_screen_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the portion of the screen that is not obstructed by a status bar in pixels. See also [method screen_get_size].
*/
//go:nosplit
func (self class) ScreenGetUsableRect(screen gd.Int) gd.Rect2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[gd.Rect2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_screen_get_usable_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the dots per inch density of the specified screen. If [param screen] is [constant SCREEN_OF_MAIN_WINDOW] (the default value), a screen with the main window will be used.
[b]Note:[/b] On macOS, returned value is inaccurate if fractional display scaling mode is used.
[b]Note:[/b] On Android devices, the actual screen densities are grouped into six generalized densities:
[codeblock lang=text]
   ldpi - 120 dpi
   mdpi - 160 dpi
   hdpi - 240 dpi
  xhdpi - 320 dpi
 xxhdpi - 480 dpi
xxxhdpi - 640 dpi
[/codeblock]
[b]Note:[/b] This method is implemented on Android, Linux (X11/Wayland), macOS and Windows. Returns [code]72[/code] on unsupported platforms.
*/
//go:nosplit
func (self class) ScreenGetDpi(screen gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_screen_get_dpi, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the scale factor of the specified screen by index.
[b]Note:[/b] On macOS, the returned value is [code]2.0[/code] for hiDPI (Retina) screens, and [code]1.0[/code] for all other cases.
[b]Note:[/b] On Linux (Wayland), the returned value is accurate only when [param screen] is [constant SCREEN_OF_MAIN_WINDOW]. Due to API limitations, passing a direct index will return a rounded-up integer, if the screen has a fractional scale (e.g. [code]1.25[/code] would get rounded up to [code]2.0[/code]).
[b]Note:[/b] This method is implemented only on macOS and Linux (Wayland).
*/
//go:nosplit
func (self class) ScreenGetScale(screen gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_screen_get_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if touch events are available (Android or iOS), the capability is detected on the Web platform or if [member ProjectSettings.input_devices/pointing/emulate_touch_from_mouse] is [code]true[/code].
*/
//go:nosplit
func (self class) IsTouchscreenAvailable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_is_touchscreen_available, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the greatest scale factor of all screens.
[b]Note:[/b] On macOS returned value is [code]2.0[/code] if there is at least one hiDPI (Retina) screen in the system, and [code]1.0[/code] in all other cases.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) ScreenGetMaxScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_screen_get_max_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the current refresh rate of the specified screen. If [param screen] is [constant SCREEN_OF_MAIN_WINDOW] (the default value), a screen with the main window will be used.
[b]Note:[/b] Returns [code]-1.0[/code] if the DisplayServer fails to find the refresh rate for the specified screen. On Web, [method screen_get_refresh_rate] will always return [code]-1.0[/code] as there is no way to retrieve the refresh rate on that platform.
To fallback to a default refresh rate if the method fails, try:
[codeblock]
var refresh_rate = DisplayServer.screen_get_refresh_rate()
if refresh_rate < 0:
    refresh_rate = 60.0
[/codeblock]
*/
//go:nosplit
func (self class) ScreenGetRefreshRate(screen gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_screen_get_refresh_rate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns color of the display pixel at the [param position].
[b]Note:[/b] This method is implemented on Linux (X11), macOS, and Windows.
[b]Note:[/b] On macOS, this method requires "Screen Recording" permission, if permission is not granted it will return desktop wallpaper color.
*/
//go:nosplit
func (self class) ScreenGetPixel(position gd.Vector2i) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_screen_get_pixel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns screenshot of the [param screen].
[b]Note:[/b] This method is implemented on Linux (X11), macOS, and Windows.
[b]Note:[/b] On macOS, this method requires "Screen Recording" permission, if permission is not granted it will return desktop wallpaper color.
*/
//go:nosplit
func (self class) ScreenGetImage(ctx gd.Lifetime, screen gd.Int) object.Image {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_screen_get_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the [param screen]'s [param orientation]. See also [method screen_get_orientation].
[b]Note:[/b] On iOS, this method has no effect if [member ProjectSettings.display/window/handheld/orientation] is not set to [constant SCREEN_SENSOR].
*/
//go:nosplit
func (self class) ScreenSetOrientation(orientation classdb.DisplayServerScreenOrientation, screen gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, orientation)
	callframe.Arg(frame, screen)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_screen_set_orientation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [param screen]'s current orientation. See also [method screen_set_orientation].
[b]Note:[/b] This method is implemented on Android and iOS.
*/
//go:nosplit
func (self class) ScreenGetOrientation(screen gd.Int) classdb.DisplayServerScreenOrientation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[classdb.DisplayServerScreenOrientation](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_screen_get_orientation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the screen should never be turned off by the operating system's power-saving measures. See also [method screen_is_kept_on].
*/
//go:nosplit
func (self class) ScreenSetKeepOn(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_screen_set_keep_on, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the screen should never be turned off by the operating system's power-saving measures. See also [method screen_set_keep_on].
*/
//go:nosplit
func (self class) ScreenIsKeptOn() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_screen_is_kept_on, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the list of Godot window IDs belonging to this process.
[b]Note:[/b] Native dialogs are not included in this list.
*/
//go:nosplit
func (self class) GetWindowList(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_get_window_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the ID of the window at the specified screen [param position] (in pixels). On multi-monitor setups, the screen position is relative to the virtual desktop area. On multi-monitor setups with different screen resolutions or orientations, the origin may be located outside any display like this:
[codeblock lang=text]
* (0, 0)        +-------+
                |       |
+-------------+ |       |
|             | |       |
|             | |       |
+-------------+ +-------+
[/codeblock]
*/
//go:nosplit
func (self class) GetWindowAtScreenPosition(position gd.Vector2i) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_get_window_at_screen_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns internal structure pointers for use in plugins.
[b]Note:[/b] This method is implemented on Android, Linux (X11/Wayland), macOS, and Windows.
*/
//go:nosplit
func (self class) WindowGetNativeHandle(handle_type classdb.DisplayServerHandleType, window_id gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, handle_type)
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_get_native_handle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns ID of the active popup window, or [constant INVALID_WINDOW_ID] if there is none.
*/
//go:nosplit
func (self class) WindowGetActivePopup() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_get_active_popup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the bounding box of control, or menu item that was used to open the popup window, in the screen coordinate system. Clicking this area will not auto-close this popup.
*/
//go:nosplit
func (self class) WindowSetPopupSafeRect(window gd.Int, rect gd.Rect2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, window)
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_set_popup_safe_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the bounding box of control, or menu item that was used to open the popup window, in the screen coordinate system.
*/
//go:nosplit
func (self class) WindowGetPopupSafeRect(window gd.Int) gd.Rect2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, window)
	var r_ret = callframe.Ret[gd.Rect2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_get_popup_safe_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the title of the given window to [param title].
[b]Note:[/b] It's recommended to change this value using [member Window.title] instead.
[b]Note:[/b] Avoid changing the window title every frame, as this can cause performance issues on certain window managers. Try to change the window title only a few times per second at most.
*/
//go:nosplit
func (self class) WindowSetTitle(title gd.String, window_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(title))
	callframe.Arg(frame, window_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_set_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the estimated window title bar size (including text and window buttons) for the window specified by [param window_id] (in pixels). This method does not change the window title.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) WindowGetTitleSize(title gd.String, window_id gd.Int) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(title))
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_get_title_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a polygonal region of the window which accepts mouse events. Mouse events outside the region will be passed through.
Passing an empty array will disable passthrough support (all mouse events will be intercepted by the window, which is the default behavior).
[codeblocks]
[gdscript]
# Set region, using Path2D node.
DisplayServer.window_set_mouse_passthrough($Path2D.curve.get_baked_points())

# Set region, using Polygon2D node.
DisplayServer.window_set_mouse_passthrough($Polygon2D.polygon)

# Reset region to default.
DisplayServer.window_set_mouse_passthrough([])
[/gdscript]
[csharp]
// Set region, using Path2D node.
DisplayServer.WindowSetMousePassthrough(GetNode<Path2D>("Path2D").Curve.GetBakedPoints());

// Set region, using Polygon2D node.
DisplayServer.WindowSetMousePassthrough(GetNode<Polygon2D>("Polygon2D").Polygon);

// Reset region to default.
DisplayServer.WindowSetMousePassthrough(new Vector2[] {});
[/csharp]
[/codeblocks]
[b]Note:[/b] On Windows, the portion of a window that lies outside the region is not drawn, while on Linux (X11) and macOS it is.
[b]Note:[/b] This method is implemented on Linux (X11), macOS and Windows.
*/
//go:nosplit
func (self class) WindowSetMousePassthrough(region gd.PackedVector2Array, window_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(region))
	callframe.Arg(frame, window_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_set_mouse_passthrough, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the screen the window specified by [param window_id] is currently positioned on. If the screen overlaps multiple displays, the screen where the window's center is located is returned. See also [method window_set_current_screen].
*/
//go:nosplit
func (self class) WindowGetCurrentScreen(window_id gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_get_current_screen, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Moves the window specified by [param window_id] to the specified [param screen]. See also [method window_get_current_screen].
*/
//go:nosplit
func (self class) WindowSetCurrentScreen(screen gd.Int, window_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	callframe.Arg(frame, window_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_set_current_screen, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the position of the client area of the given window on the screen.
*/
//go:nosplit
func (self class) WindowGetPosition(window_id gd.Int) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the position of the given window on the screen including the borders drawn by the operating system. See also [method window_get_position].
*/
//go:nosplit
func (self class) WindowGetPositionWithDecorations(window_id gd.Int) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_get_position_with_decorations, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the position of the given window to [param position]. On multi-monitor setups, the screen position is relative to the virtual desktop area. On multi-monitor setups with different screen resolutions or orientations, the origin may be located outside any display like this:
[codeblock lang=text]
* (0, 0)        +-------+
                |       |
+-------------+ |       |
|             | |       |
|             | |       |
+-------------+ +-------+
[/codeblock]
See also [method window_get_position] and [method window_set_size].
[b]Note:[/b] It's recommended to change this value using [member Window.position] instead.
[b]Note:[/b] On Linux (Wayland): this method is a no-op.
*/
//go:nosplit
func (self class) WindowSetPosition(position gd.Vector2i, window_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, window_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_set_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the size of the window specified by [param window_id] (in pixels), excluding the borders drawn by the operating system. This is also called the "client area". See also [method window_get_size_with_decorations], [method window_set_size] and [method window_get_position].
*/
//go:nosplit
func (self class) WindowGetSize(window_id gd.Int) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the size of the given window to [param size] (in pixels). See also [method window_get_size] and [method window_get_position].
[b]Note:[/b] It's recommended to change this value using [member Window.size] instead.
*/
//go:nosplit
func (self class) WindowSetSize(size gd.Vector2i, window_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, window_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [param callback] that will be called when the window specified by [param window_id] is moved or resized.
[b]Warning:[/b] Advanced users only! Adding such a callback to a [Window] node will override its default implementation, which can introduce bugs.
*/
//go:nosplit
func (self class) WindowSetRectChangedCallback(callback gd.Callable, window_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(callback))
	callframe.Arg(frame, window_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_set_rect_changed_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [param callback] that will be called when an event occurs in the window specified by [param window_id].
[b]Warning:[/b] Advanced users only! Adding such a callback to a [Window] node will override its default implementation, which can introduce bugs.
*/
//go:nosplit
func (self class) WindowSetWindowEventCallback(callback gd.Callable, window_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(callback))
	callframe.Arg(frame, window_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_set_window_event_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [param callback] that should be called when any [InputEvent] is sent to the window specified by [param window_id].
[b]Warning:[/b] Advanced users only! Adding such a callback to a [Window] node will override its default implementation, which can introduce bugs.
*/
//go:nosplit
func (self class) WindowSetInputEventCallback(callback gd.Callable, window_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(callback))
	callframe.Arg(frame, window_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_set_input_event_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [param callback] that should be called when text is entered using the virtual keyboard to the window specified by [param window_id].
[b]Warning:[/b] Advanced users only! Adding such a callback to a [Window] node will override its default implementation, which can introduce bugs.
*/
//go:nosplit
func (self class) WindowSetInputTextCallback(callback gd.Callable, window_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(callback))
	callframe.Arg(frame, window_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_set_input_text_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [param callback] that should be called when files are dropped from the operating system's file manager to the window specified by [param window_id]. [param callback] should take one [PackedStringArray] argument, which is the list of dropped files.
[b]Warning:[/b] Advanced users only! Adding such a callback to a [Window] node will override its default implementation, which can introduce bugs.
[b]Note:[/b] This method is implemented on Windows, macOS, Linux (X11/Wayland), and Web.
*/
//go:nosplit
func (self class) WindowSetDropFilesCallback(callback gd.Callable, window_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(callback))
	callframe.Arg(frame, window_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_set_drop_files_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [method Object.get_instance_id] of the [Window] the [param window_id] is attached to.
*/
//go:nosplit
func (self class) WindowGetAttachedInstanceId(window_id gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_get_attached_instance_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the window's maximum size (in pixels). See also [method window_set_max_size].
*/
//go:nosplit
func (self class) WindowGetMaxSize(window_id gd.Int) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_get_max_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the maximum size of the window specified by [param window_id] in pixels. Normally, the user will not be able to drag the window to make it larger than the specified size. See also [method window_get_max_size].
[b]Note:[/b] It's recommended to change this value using [member Window.max_size] instead.
[b]Note:[/b] Using third-party tools, it is possible for users to disable window geometry restrictions and therefore bypass this limit.
*/
//go:nosplit
func (self class) WindowSetMaxSize(max_size gd.Vector2i, window_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_size)
	callframe.Arg(frame, window_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_set_max_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the window's minimum size (in pixels). See also [method window_set_min_size].
*/
//go:nosplit
func (self class) WindowGetMinSize(window_id gd.Int) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_get_min_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the minimum size for the given window to [param min_size] in pixels. Normally, the user will not be able to drag the window to make it smaller than the specified size. See also [method window_get_min_size].
[b]Note:[/b] It's recommended to change this value using [member Window.min_size] instead.
[b]Note:[/b] By default, the main window has a minimum size of [code]Vector2i(64, 64)[/code]. This prevents issues that can arise when the window is resized to a near-zero size.
[b]Note:[/b] Using third-party tools, it is possible for users to disable window geometry restrictions and therefore bypass this limit.
*/
//go:nosplit
func (self class) WindowSetMinSize(min_size gd.Vector2i, window_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, min_size)
	callframe.Arg(frame, window_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_set_min_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the size of the window specified by [param window_id] (in pixels), including the borders drawn by the operating system. See also [method window_get_size].
*/
//go:nosplit
func (self class) WindowGetSizeWithDecorations(window_id gd.Int) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_get_size_with_decorations, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the mode of the given window.
*/
//go:nosplit
func (self class) WindowGetMode(window_id gd.Int) classdb.DisplayServerWindowMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[classdb.DisplayServerWindowMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_get_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets window mode for the given window to [param mode]. See [enum WindowMode] for possible values and how each mode behaves.
[b]Note:[/b] Setting the window to full screen forcibly sets the borderless flag to [code]true[/code], so make sure to set it back to [code]false[/code] when not wanted.
*/
//go:nosplit
func (self class) WindowSetMode(mode classdb.DisplayServerWindowMode, window_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	callframe.Arg(frame, window_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_set_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Enables or disables the given window's given [param flag]. See [enum WindowFlags] for possible values and their behavior.
*/
//go:nosplit
func (self class) WindowSetFlag(flag classdb.DisplayServerWindowFlags, enabled bool, window_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, enabled)
	callframe.Arg(frame, window_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_set_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the current value of the given window's [param flag].
*/
//go:nosplit
func (self class) WindowGetFlag(flag classdb.DisplayServerWindowFlags, window_id gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_get_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
When [constant WINDOW_FLAG_EXTEND_TO_TITLE] flag is set, set offset to the center of the first titlebar button.
[b]Note:[/b] This flag is implemented only on macOS.
*/
//go:nosplit
func (self class) WindowSetWindowButtonsOffset(offset gd.Vector2i, window_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	callframe.Arg(frame, window_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_set_window_buttons_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns left margins ([code]x[/code]), right margins ([code]y[/code]) and height ([code]z[/code]) of the title that are safe to use (contains no buttons or other elements) when [constant WINDOW_FLAG_EXTEND_TO_TITLE] flag is set.
*/
//go:nosplit
func (self class) WindowGetSafeTitleMargins(window_id gd.Int) gd.Vector3i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Vector3i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_get_safe_title_margins, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Makes the window specified by [param window_id] request attention, which is materialized by the window title and taskbar entry blinking until the window is focused. This usually has no visible effect if the window is currently focused. The exact behavior varies depending on the operating system.
*/
//go:nosplit
func (self class) WindowRequestAttention(window_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_request_attention, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Moves the window specified by [param window_id] to the foreground, so that it is visible over other windows.
*/
//go:nosplit
func (self class) WindowMoveToForeground(window_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_move_to_foreground, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the window specified by [param window_id] is focused.
*/
//go:nosplit
func (self class) WindowIsFocused(window_id gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_is_focused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if anything can be drawn in the window specified by [param window_id], [code]false[/code] otherwise. Using the [code]--disable-render-loop[/code] command line argument or a headless build will return [code]false[/code].
*/
//go:nosplit
func (self class) WindowCanDraw(window_id gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_can_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets window transient parent. Transient window is will be destroyed with its transient parent and will return focus to their parent when closed. The transient window is displayed on top of a non-exclusive full-screen parent window. Transient windows can't enter full-screen mode.
[b]Note:[/b] It's recommended to change this value using [member Window.transient] instead.
[b]Note:[/b] The behavior might be different depending on the platform.
*/
//go:nosplit
func (self class) WindowSetTransient(window_id gd.Int, parent_window_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	callframe.Arg(frame, parent_window_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_set_transient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If set to [code]true[/code], this window will always stay on top of its parent window, parent window will ignore input while this window is opened.
[b]Note:[/b] On macOS, exclusive windows are confined to the same space (virtual desktop or screen) as the parent window.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) WindowSetExclusive(window_id gd.Int, exclusive bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	callframe.Arg(frame, exclusive)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_set_exclusive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets whether [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] should be enabled for the window specified by [param window_id]. See also [method window_set_ime_position].
*/
//go:nosplit
func (self class) WindowSetImeActive(active bool, window_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, active)
	callframe.Arg(frame, window_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_set_ime_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the position of the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] popup for the specified [param window_id]. Only effective if [method window_set_ime_active] was set to [code]true[/code] for the specified [param window_id].
*/
//go:nosplit
func (self class) WindowSetImePosition(position gd.Vector2i, window_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, window_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_set_ime_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the V-Sync mode of the given window. See also [member ProjectSettings.display/window/vsync/vsync_mode].
See [enum DisplayServer.VSyncMode] for possible values and how they affect the behavior of your application.
Depending on the platform and used renderer, the engine will fall back to [constant VSYNC_ENABLED] if the desired mode is not supported.
[b]Note:[/b] V-Sync modes other than [constant VSYNC_ENABLED] are only supported in the Forward+ and Mobile rendering methods, not Compatibility.
*/
//go:nosplit
func (self class) WindowSetVsyncMode(vsync_mode classdb.DisplayServerVSyncMode, window_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, vsync_mode)
	callframe.Arg(frame, window_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_set_vsync_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the V-Sync mode of the given window.
*/
//go:nosplit
func (self class) WindowGetVsyncMode(window_id gd.Int) classdb.DisplayServerVSyncMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[classdb.DisplayServerVSyncMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_get_vsync_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the given window can be maximized (the maximize button is enabled).
*/
//go:nosplit
func (self class) WindowIsMaximizeAllowed(window_id gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_is_maximize_allowed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code], if double-click on a window title should maximize it.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) WindowMaximizeOnTitleDblClick() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_maximize_on_title_dbl_click, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code], if double-click on a window title should minimize it.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) WindowMinimizeOnTitleDblClick() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_window_minimize_on_title_dbl_click, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the text selection in the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] composition string, with the [Vector2i]'s [code]x[/code] component being the caret position and [code]y[/code] being the length of the selection.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) ImeGetSelection() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_ime_get_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the composition string contained within the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] window.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) ImeGetText(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_ime_get_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Shows the virtual keyboard if the platform has one.
[param existing_text] parameter is useful for implementing your own [LineEdit] or [TextEdit], as it tells the virtual keyboard what text has already been typed (the virtual keyboard uses it for auto-correct and predictions).
[param position] parameter is the screen space [Rect2] of the edited text.
[param type] parameter allows configuring which type of virtual keyboard to show.
[param max_length] limits the number of characters that can be entered if different from [code]-1[/code].
[param cursor_start] can optionally define the current text cursor position if [param cursor_end] is not set.
[param cursor_start] and [param cursor_end] can optionally define the current text selection.
[b]Note:[/b] This method is implemented on Android, iOS and Web.
*/
//go:nosplit
func (self class) VirtualKeyboardShow(existing_text gd.String, position gd.Rect2, atype classdb.DisplayServerVirtualKeyboardType, max_length gd.Int, cursor_start gd.Int, cursor_end gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(existing_text))
	callframe.Arg(frame, position)
	callframe.Arg(frame, atype)
	callframe.Arg(frame, max_length)
	callframe.Arg(frame, cursor_start)
	callframe.Arg(frame, cursor_end)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_virtual_keyboard_show, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Hides the virtual keyboard if it is shown, does nothing otherwise.
*/
//go:nosplit
func (self class) VirtualKeyboardHide()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_virtual_keyboard_hide, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the on-screen keyboard's height in pixels. Returns 0 if there is no keyboard or if it is currently hidden.
*/
//go:nosplit
func (self class) VirtualKeyboardGetHeight() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_virtual_keyboard_get_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the default mouse cursor shape. The cursor's appearance will vary depending on the user's operating system and mouse cursor theme. See also [method cursor_get_shape] and [method cursor_set_custom_image].
*/
//go:nosplit
func (self class) CursorSetShape(shape classdb.DisplayServerCursorShape)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_cursor_set_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the default mouse cursor shape set by [method cursor_set_shape].
*/
//go:nosplit
func (self class) CursorGetShape() classdb.DisplayServerCursorShape {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.DisplayServerCursorShape](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_cursor_get_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a custom mouse cursor image for the given [param shape]. This means the user's operating system and mouse cursor theme will no longer influence the mouse cursor's appearance.
[param cursor] can be either a [Texture2D] or an [Image], and it should not be larger than 256×256 to display correctly. Optionally, [param hotspot] can be set to offset the image's position relative to the click point. By default, [param hotspot] is set to the top-left corner of the image. See also [method cursor_set_shape].
*/
//go:nosplit
func (self class) CursorSetCustomImage(cursor object.Resource, shape classdb.DisplayServerCursorShape, hotspot gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(cursor[0].AsPointer())[0])
	callframe.Arg(frame, shape)
	callframe.Arg(frame, hotspot)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_cursor_set_custom_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if positions of [b]OK[/b] and [b]Cancel[/b] buttons are swapped in dialogs. This is enabled by default on Windows to follow interface conventions, and be toggled by changing [member ProjectSettings.gui/common/swap_cancel_ok].
[b]Note:[/b] This doesn't affect native dialogs such as the ones spawned by [method DisplayServer.dialog_show].
*/
//go:nosplit
func (self class) GetSwapCancelOk() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_get_swap_cancel_ok, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Allows the [param process_id] PID to steal focus from this window. In other words, this disables the operating system's focus stealing protection for the specified PID.
[b]Note:[/b] This method is implemented only on Windows.
*/
//go:nosplit
func (self class) EnableForStealingFocus(process_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, process_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_enable_for_stealing_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Shows a text dialog which uses the operating system's native look-and-feel. [param callback] should accept a single [int] parameter which corresponds to the index of the pressed button.
[b]Note:[/b] This method is implemented if the display server has the [constant FEATURE_NATIVE_DIALOG] feature. Supported platforms include macOS and Windows.
*/
//go:nosplit
func (self class) DialogShow(title gd.String, description gd.String, buttons gd.PackedStringArray, callback gd.Callable) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(title))
	callframe.Arg(frame, mmm.Get(description))
	callframe.Arg(frame, mmm.Get(buttons))
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_dialog_show, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Shows a text input dialog which uses the operating system's native look-and-feel. [param callback] should accept a single [String] parameter which contains the text field's contents.
[b]Note:[/b] This method is implemented if the display server has the [constant FEATURE_NATIVE_DIALOG_INPUT] feature. Supported platforms include macOS and Windows.
*/
//go:nosplit
func (self class) DialogInputText(title gd.String, description gd.String, existing_text gd.String, callback gd.Callable) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(title))
	callframe.Arg(frame, mmm.Get(description))
	callframe.Arg(frame, mmm.Get(existing_text))
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_dialog_input_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Displays OS native dialog for selecting files or directories in the file system.
Each filter string in the [param filters] array should be formatted like this: [code]*.txt,*.doc;Text Files[/code]. The description text of the filter is optional and can be omitted. See also [member FileDialog.filters].
Callbacks have the following arguments: [code]status: bool, selected_paths: PackedStringArray, selected_filter_index: int[/code].
[b]Note:[/b] This method is implemented if the display server has the [constant FEATURE_NATIVE_DIALOG_FILE] feature. Supported platforms include Linux (X11/Wayland), Windows, and macOS.
[b]Note:[/b] [param current_directory] might be ignored.
[b]Note:[/b] On Linux, [param show_hidden] is ignored.
[b]Note:[/b] On macOS, native file dialogs have no title.
[b]Note:[/b] On macOS, sandboxed apps will save security-scoped bookmarks to retain access to the opened folders across multiple sessions. Use [method OS.get_granted_permissions] to get a list of saved bookmarks.
*/
//go:nosplit
func (self class) FileDialogShow(title gd.String, current_directory gd.String, filename gd.String, show_hidden bool, mode classdb.DisplayServerFileDialogMode, filters gd.PackedStringArray, callback gd.Callable) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(title))
	callframe.Arg(frame, mmm.Get(current_directory))
	callframe.Arg(frame, mmm.Get(filename))
	callframe.Arg(frame, show_hidden)
	callframe.Arg(frame, mode)
	callframe.Arg(frame, mmm.Get(filters))
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_file_dialog_show, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Displays OS native dialog for selecting files or directories in the file system with additional user selectable options.
Each filter string in the [param filters] array should be formatted like this: [code]*.txt,*.doc;Text Files[/code]. The description text of the filter is optional and can be omitted. See also [member FileDialog.filters].
[param options] is array of [Dictionary]s with the following keys:
- [code]"name"[/code] - option's name [String].
- [code]"values"[/code] - [PackedStringArray] of values. If empty, boolean option (check box) is used.
- [code]"default"[/code] - default selected option index ([int]) or default boolean value ([bool]).
Callbacks have the following arguments: [code]status: bool, selected_paths: PackedStringArray, selected_filter_index: int, selected_option: Dictionary[/code].
[b]Note:[/b] This method is implemented if the display server has the [constant FEATURE_NATIVE_DIALOG_FILE] feature. Supported platforms include Linux (X11/Wayland), Windows, and macOS.
[b]Note:[/b] [param current_directory] might be ignored.
[b]Note:[/b] On Linux (X11), [param show_hidden] is ignored.
[b]Note:[/b] On macOS, native file dialogs have no title.
[b]Note:[/b] On macOS, sandboxed apps will save security-scoped bookmarks to retain access to the opened folders across multiple sessions. Use [method OS.get_granted_permissions] to get a list of saved bookmarks.
*/
//go:nosplit
func (self class) FileDialogWithOptionsShow(title gd.String, current_directory gd.String, root gd.String, filename gd.String, show_hidden bool, mode classdb.DisplayServerFileDialogMode, filters gd.PackedStringArray, options gd.ArrayOf[gd.Dictionary], callback gd.Callable) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(title))
	callframe.Arg(frame, mmm.Get(current_directory))
	callframe.Arg(frame, mmm.Get(root))
	callframe.Arg(frame, mmm.Get(filename))
	callframe.Arg(frame, show_hidden)
	callframe.Arg(frame, mode)
	callframe.Arg(frame, mmm.Get(filters))
	callframe.Arg(frame, mmm.Get(options))
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_file_dialog_with_options_show, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of keyboard layouts.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
//go:nosplit
func (self class) KeyboardGetLayoutCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_keyboard_get_layout_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns active keyboard layout index.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS, and Windows.
*/
//go:nosplit
func (self class) KeyboardGetCurrentLayout() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_keyboard_get_current_layout, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the active keyboard layout.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
//go:nosplit
func (self class) KeyboardSetCurrentLayout(index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_keyboard_set_current_layout, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the ISO-639/BCP-47 language code of the keyboard layout at position [param index].
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
//go:nosplit
func (self class) KeyboardGetLayoutLanguage(ctx gd.Lifetime, index gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_keyboard_get_layout_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the localized name of the keyboard layout at position [param index].
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
//go:nosplit
func (self class) KeyboardGetLayoutName(ctx gd.Lifetime, index gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_keyboard_get_layout_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Converts a physical (US QWERTY) [param keycode] to one in the active keyboard layout.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
//go:nosplit
func (self class) KeyboardGetKeycodeFromPhysical(keycode gd.Key) gd.Key {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, keycode)
	var r_ret = callframe.Ret[gd.Key](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_keyboard_get_keycode_from_physical, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Converts a physical (US QWERTY) [param keycode] to localized label printed on the key in the active keyboard layout.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
//go:nosplit
func (self class) KeyboardGetLabelFromPhysical(keycode gd.Key) gd.Key {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, keycode)
	var r_ret = callframe.Ret[gd.Key](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_keyboard_get_label_from_physical, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Perform window manager processing, including input flushing. See also [method force_process_and_drop_events], [method Input.flush_buffered_events] and [member Input.use_accumulated_input].
*/
//go:nosplit
func (self class) ProcessEvents()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_process_events, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Forces window manager processing while ignoring all [InputEvent]s. See also [method process_events].
[b]Note:[/b] This method is implemented on Windows and macOS.
*/
//go:nosplit
func (self class) ForceProcessAndDropEvents()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_force_process_and_drop_events, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the window icon (usually displayed in the top-left corner) in the operating system's [i]native[/i] format. The file at [param filename] must be in [code].ico[/code] format on Windows or [code].icns[/code] on macOS. By using specially crafted [code].ico[/code] or [code].icns[/code] icons, [method set_native_icon] allows specifying different icons depending on the size the icon is displayed at. This size is determined by the operating system and user preferences (including the display scale factor). To use icons in other formats, use [method set_icon] instead.
[b]Note:[/b] Requires support for [constant FEATURE_NATIVE_ICON].
*/
//go:nosplit
func (self class) SetNativeIcon(filename gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(filename))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_set_native_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the window icon (usually displayed in the top-left corner) with an [Image]. To use icons in the operating system's native format, use [method set_native_icon] instead.
[b]Note:[/b] Requires support for [constant FEATURE_ICON].
*/
//go:nosplit
func (self class) SetIcon(image object.Image)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(image[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_set_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a new application status indicator with the specified icon, tooltip, and activation callback.
[param callback] should take two arguments: the pressed mouse button (one of the [enum MouseButton] constants) and the click position in screen coordinates (a [Vector2i]).
*/
//go:nosplit
func (self class) CreateStatusIndicator(icon object.Texture2D, tooltip gd.String, callback gd.Callable) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(tooltip))
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_create_status_indicator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the application status indicator icon.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) StatusIndicatorSetIcon(id gd.Int, icon object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_status_indicator_set_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the application status indicator tooltip.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) StatusIndicatorSetTooltip(id gd.Int, tooltip gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, mmm.Get(tooltip))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_status_indicator_set_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the application status indicator native popup menu.
[b]Note:[/b] On macOS, the menu is activated by any mouse button. Its activation callback is [i]not[/i] triggered.
[b]Note:[/b] On Windows, the menu is activated by the right mouse button, selecting the status icon and pressing [kbd]Shift + F10[/kbd], or the applications key. The menu's activation callback for the other mouse buttons is still triggered.
[b]Note:[/b] Native popup is only supported if [NativeMenu] supports the [constant NativeMenu.FEATURE_POPUP_MENU] feature.
*/
//go:nosplit
func (self class) StatusIndicatorSetMenu(id gd.Int, menu_rid gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, menu_rid)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_status_indicator_set_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the application status indicator activation callback. [param callback] should take two arguments: [int] mouse button index (one of [enum MouseButton] values) and [Vector2i] click position in screen coordinates.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) StatusIndicatorSetCallback(id gd.Int, callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_status_indicator_set_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the rectangle for the given status indicator [param id] in screen coordinates. If the status indicator is not visible, returns an empty [Rect2].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) StatusIndicatorGetRect(id gd.Int) gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_status_indicator_get_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the application status indicator.
*/
//go:nosplit
func (self class) DeleteStatusIndicator(id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_delete_status_indicator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the total number of available tablet drivers.
[b]Note:[/b] This method is implemented only on Windows.
*/
//go:nosplit
func (self class) TabletGetDriverCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_tablet_get_driver_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the tablet driver name for the given index.
[b]Note:[/b] This method is implemented only on Windows.
*/
//go:nosplit
func (self class) TabletGetDriverName(ctx gd.Lifetime, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_tablet_get_driver_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns current active tablet driver name.
[b]Note:[/b] This method is implemented only on Windows.
*/
//go:nosplit
func (self class) TabletGetCurrentDriver(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_tablet_get_current_driver, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Set active tablet driver name.
Supported drivers:
- [code]winink[/code]: Windows Ink API, default (Windows 8.1+ required).
- [code]wintab[/code]: Wacom Wintab API (compatible device driver required).
- [code]dummy[/code]: Dummy driver, tablet input is disabled.
[b]Note:[/b] This method is implemented only on Windows.
*/
//go:nosplit
func (self class) TabletSetCurrentDriver(name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_tablet_set_current_driver, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the window background can be made transparent. This method returns [code]false[/code] if [member ProjectSettings.display/window/per_pixel_transparency/allowed] is set to [code]false[/code], or if transparency is not supported by the renderer or OS compositor.
*/
//go:nosplit
func (self class) IsWindowTransparencyAvailable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_is_window_transparency_available, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Registers an [Object] which represents an additional output that will be rendered too, beyond normal windows. The [Object] is only used as an identifier, which can be later passed to [method unregister_additional_output].
This can be used to prevent Godot from skipping rendering when no normal windows are visible.
*/
//go:nosplit
func (self class) RegisterAdditionalOutput(obj gd.Object)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(obj.AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_register_additional_output, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Unregisters an [Object] representing an additional output, that was registered via [method register_additional_output].
*/
//go:nosplit
func (self class) UnregisterAdditionalOutput(obj gd.Object)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(obj.AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_unregister_additional_output, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if any additional outputs have been registered via [method register_additional_output].
*/
//go:nosplit
func (self class) HasAdditionalOutputs() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DisplayServer.Bind_has_additional_outputs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsDisplayServer() Expert { return self[0].AsDisplayServer() }


//go:nosplit
func (self Simple) AsDisplayServer() Simple { return self[0].AsDisplayServer() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("DisplayServer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Feature = classdb.DisplayServerFeature

const (
/*Display server supports global menu. This allows the application to display its menu items in the operating system's top bar. [b]macOS[/b]*/
	FeatureGlobalMenu Feature = 0
/*Display server supports multiple windows that can be moved outside of the main window. [b]Windows, macOS, Linux (X11)[/b]*/
	FeatureSubwindows Feature = 1
/*Display server supports touchscreen input. [b]Windows, Linux (X11), Android, iOS, Web[/b]*/
	FeatureTouchscreen Feature = 2
/*Display server supports mouse input. [b]Windows, macOS, Linux (X11/Wayland), Android, Web[/b]*/
	FeatureMouse Feature = 3
/*Display server supports warping mouse coordinates to keep the mouse cursor constrained within an area, but looping when one of the edges is reached. [b]Windows, macOS, Linux (X11/Wayland)[/b]*/
	FeatureMouseWarp Feature = 4
/*Display server supports setting and getting clipboard data. See also [constant FEATURE_CLIPBOARD_PRIMARY]. [b]Windows, macOS, Linux (X11/Wayland), Android, iOS, Web[/b]*/
	FeatureClipboard Feature = 5
/*Display server supports popping up a virtual keyboard when requested to input text without a physical keyboard. [b]Android, iOS, Web[/b]*/
	FeatureVirtualKeyboard Feature = 6
/*Display server supports setting the mouse cursor shape to be different from the default. [b]Windows, macOS, Linux (X11/Wayland), Android, Web[/b]*/
	FeatureCursorShape Feature = 7
/*Display server supports setting the mouse cursor shape to a custom image. [b]Windows, macOS, Linux (X11/Wayland), Web[/b]*/
	FeatureCustomCursorShape Feature = 8
/*Display server supports spawning text dialogs using the operating system's native look-and-feel. See [method dialog_show]. [b]Windows, macOS[/b]*/
	FeatureNativeDialog Feature = 9
/*Display server supports [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url], which is commonly used for inputting Chinese/Japanese/Korean text. This is handled by the operating system, rather than by Godot. [b]Windows, macOS, Linux (X11)[/b]*/
	FeatureIme Feature = 10
/*Display server supports windows can use per-pixel transparency to make windows behind them partially or fully visible. [b]Windows, macOS, Linux (X11/Wayland)[/b]*/
	FeatureWindowTransparency Feature = 11
/*Display server supports querying the operating system's display scale factor. This allows for [i]reliable[/i] automatic hiDPI display detection, as opposed to guessing based on the screen resolution and reported display DPI (which can be unreliable due to broken monitor EDID). [b]Windows, Linux (Wayland), macOS[/b]*/
	FeatureHidpi Feature = 12
/*Display server supports changing the window icon (usually displayed in the top-left corner). [b]Windows, macOS, Linux (X11)[/b]*/
	FeatureIcon Feature = 13
/*Display server supports changing the window icon (usually displayed in the top-left corner). [b]Windows, macOS[/b]*/
	FeatureNativeIcon Feature = 14
/*Display server supports changing the screen orientation. [b]Android, iOS[/b]*/
	FeatureOrientation Feature = 15
/*Display server supports V-Sync status can be changed from the default (which is forced to be enabled platforms not supporting this feature). [b]Windows, macOS, Linux (X11/Wayland)[/b]*/
	FeatureSwapBuffers Feature = 16
/*Display server supports Primary clipboard can be used. This is a different clipboard from [constant FEATURE_CLIPBOARD]. [b]Linux (X11/Wayland)[/b]*/
	FeatureClipboardPrimary Feature = 18
/*Display server supports text-to-speech. See [code]tts_*[/code] methods. [b]Windows, macOS, Linux (X11/Wayland), Android, iOS, Web[/b]*/
	FeatureTextToSpeech Feature = 19
/*Display server supports expanding window content to the title. See [constant WINDOW_FLAG_EXTEND_TO_TITLE]. [b]macOS[/b]*/
	FeatureExtendToTitle Feature = 20
/*Display server supports reading screen pixels. See [method screen_get_pixel].*/
	FeatureScreenCapture Feature = 21
/*Display server supports application status indicators.*/
	FeatureStatusIndicator Feature = 22
/*Display server supports native help system search callbacks. See [method help_set_search_callbacks].*/
	FeatureNativeHelp Feature = 23
/*Display server supports spawning text input dialogs using the operating system's native look-and-feel. See [method dialog_input_text]. [b]Windows, macOS[/b]*/
	FeatureNativeDialogInput Feature = 24
/*Display server supports spawning dialogs for selecting files or directories using the operating system's native look-and-feel. See [method file_dialog_show] and [method file_dialog_with_options_show]. [b]Windows, macOS, Linux (X11/Wayland)[/b]*/
	FeatureNativeDialogFile Feature = 25
)
type MouseMode = classdb.DisplayServerMouseMode

const (
/*Makes the mouse cursor visible if it is hidden.*/
	MouseModeVisible MouseMode = 0
/*Makes the mouse cursor hidden if it is visible.*/
	MouseModeHidden MouseMode = 1
/*Captures the mouse. The mouse will be hidden and its position locked at the center of the window manager's window.
[b]Note:[/b] If you want to process the mouse's movement in this mode, you need to use [member InputEventMouseMotion.relative].*/
	MouseModeCaptured MouseMode = 2
/*Confines the mouse cursor to the game window, and make it visible.*/
	MouseModeConfined MouseMode = 3
/*Confines the mouse cursor to the game window, and make it hidden.*/
	MouseModeConfinedHidden MouseMode = 4
)
type ScreenOrientation = classdb.DisplayServerScreenOrientation

const (
/*Default landscape orientation.*/
	ScreenLandscape ScreenOrientation = 0
/*Default portrait orientation.*/
	ScreenPortrait ScreenOrientation = 1
/*Reverse landscape orientation (upside down).*/
	ScreenReverseLandscape ScreenOrientation = 2
/*Reverse portrait orientation (upside down).*/
	ScreenReversePortrait ScreenOrientation = 3
/*Automatic landscape orientation (default or reverse depending on sensor).*/
	ScreenSensorLandscape ScreenOrientation = 4
/*Automatic portrait orientation (default or reverse depending on sensor).*/
	ScreenSensorPortrait ScreenOrientation = 5
/*Automatic landscape or portrait orientation (default or reverse depending on sensor).*/
	ScreenSensor ScreenOrientation = 6
)
type VirtualKeyboardType = classdb.DisplayServerVirtualKeyboardType

const (
/*Default text virtual keyboard.*/
	KeyboardTypeDefault VirtualKeyboardType = 0
/*Multiline virtual keyboard.*/
	KeyboardTypeMultiline VirtualKeyboardType = 1
/*Virtual number keypad, useful for PIN entry.*/
	KeyboardTypeNumber VirtualKeyboardType = 2
/*Virtual number keypad, useful for entering fractional numbers.*/
	KeyboardTypeNumberDecimal VirtualKeyboardType = 3
/*Virtual phone number keypad.*/
	KeyboardTypePhone VirtualKeyboardType = 4
/*Virtual keyboard with additional keys to assist with typing email addresses.*/
	KeyboardTypeEmailAddress VirtualKeyboardType = 5
/*Virtual keyboard for entering a password. On most platforms, this should disable autocomplete and autocapitalization.
[b]Note:[/b] This is not supported on Web. Instead, this behaves identically to [constant KEYBOARD_TYPE_DEFAULT].*/
	KeyboardTypePassword VirtualKeyboardType = 6
/*Virtual keyboard with additional keys to assist with typing URLs.*/
	KeyboardTypeUrl VirtualKeyboardType = 7
)
type CursorShape = classdb.DisplayServerCursorShape

const (
/*Arrow cursor shape. This is the default when not pointing anything that overrides the mouse cursor, such as a [LineEdit] or [TextEdit].*/
	CursorArrow CursorShape = 0
/*I-beam cursor shape. This is used by default when hovering a control that accepts text input, such as [LineEdit] or [TextEdit].*/
	CursorIbeam CursorShape = 1
/*Pointing hand cursor shape. This is used by default when hovering a [LinkButton] or a URL tag in a [RichTextLabel].*/
	CursorPointingHand CursorShape = 2
/*Crosshair cursor. This is intended to be displayed when the user needs precise aim over an element, such as a rectangle selection tool or a color picker.*/
	CursorCross CursorShape = 3
/*Wait cursor. On most cursor themes, this displays a spinning icon [i]besides[/i] the arrow. Intended to be used for non-blocking operations (when the user can do something else at the moment). See also [constant CURSOR_BUSY].*/
	CursorWait CursorShape = 4
/*Wait cursor. On most cursor themes, this [i]replaces[/i] the arrow with a spinning icon. Intended to be used for blocking operations (when the user can't do anything else at the moment). See also [constant CURSOR_WAIT].*/
	CursorBusy CursorShape = 5
/*Dragging hand cursor. This is displayed during drag-and-drop operations. See also [constant CURSOR_CAN_DROP].*/
	CursorDrag CursorShape = 6
/*"Can drop" cursor. This is displayed during drag-and-drop operations if hovering over a [Control] that can accept the drag-and-drop event. On most cursor themes, this displays a dragging hand with an arrow symbol besides it. See also [constant CURSOR_DRAG].*/
	CursorCanDrop CursorShape = 7
/*Forbidden cursor. This is displayed during drag-and-drop operations if the hovered [Control] can't accept the drag-and-drop event.*/
	CursorForbidden CursorShape = 8
/*Vertical resize cursor. Intended to be displayed when the hovered [Control] can be vertically resized using the mouse. See also [constant CURSOR_VSPLIT].*/
	CursorVsize CursorShape = 9
/*Horizontal resize cursor. Intended to be displayed when the hovered [Control] can be horizontally resized using the mouse. See also [constant CURSOR_HSPLIT].*/
	CursorHsize CursorShape = 10
/*Secondary diagonal resize cursor (top-right/bottom-left). Intended to be displayed when the hovered [Control] can be resized on both axes at once using the mouse.*/
	CursorBdiagsize CursorShape = 11
/*Main diagonal resize cursor (top-left/bottom-right). Intended to be displayed when the hovered [Control] can be resized on both axes at once using the mouse.*/
	CursorFdiagsize CursorShape = 12
/*Move cursor. Intended to be displayed when the hovered [Control] can be moved using the mouse.*/
	CursorMove CursorShape = 13
/*Vertical split cursor. This is displayed when hovering a [Control] with splits that can be vertically resized using the mouse, such as [VSplitContainer]. On some cursor themes, this cursor may have the same appearance as [constant CURSOR_VSIZE].*/
	CursorVsplit CursorShape = 14
/*Horizontal split cursor. This is displayed when hovering a [Control] with splits that can be horizontally resized using the mouse, such as [HSplitContainer]. On some cursor themes, this cursor may have the same appearance as [constant CURSOR_HSIZE].*/
	CursorHsplit CursorShape = 15
/*Help cursor. On most cursor themes, this displays a question mark icon instead of the mouse cursor. Intended to be used when the user has requested help on the next element that will be clicked.*/
	CursorHelp CursorShape = 16
/*Represents the size of the [enum CursorShape] enum.*/
	CursorMax CursorShape = 17
)
type FileDialogMode = classdb.DisplayServerFileDialogMode

const (
/*The native file dialog allows selecting one, and only one file.*/
	FileDialogModeOpenFile FileDialogMode = 0
/*The native file dialog allows selecting multiple files.*/
	FileDialogModeOpenFiles FileDialogMode = 1
/*The native file dialog only allows selecting a directory, disallowing the selection of any file.*/
	FileDialogModeOpenDir FileDialogMode = 2
/*The native file dialog allows selecting one file or directory.*/
	FileDialogModeOpenAny FileDialogMode = 3
/*The native file dialog will warn when a file exists.*/
	FileDialogModeSaveFile FileDialogMode = 4
)
type WindowMode = classdb.DisplayServerWindowMode

const (
/*Windowed mode, i.e. [Window] doesn't occupy the whole screen (unless set to the size of the screen).*/
	WindowModeWindowed WindowMode = 0
/*Minimized window mode, i.e. [Window] is not visible and available on window manager's window list. Normally happens when the minimize button is pressed.*/
	WindowModeMinimized WindowMode = 1
/*Maximized window mode, i.e. [Window] will occupy whole screen area except task bar and still display its borders. Normally happens when the maximize button is pressed.*/
	WindowModeMaximized WindowMode = 2
/*Full screen mode with full multi-window support.
Full screen window covers the entire display area of a screen and has no decorations. The display's video mode is not changed.
[b]On Windows:[/b] Multi-window full-screen mode has a 1px border of the [member ProjectSettings.rendering/environment/defaults/default_clear_color] color.
[b]On macOS:[/b] A new desktop is used to display the running project.
[b]Note:[/b] Regardless of the platform, enabling full screen will change the window size to match the monitor's size. Therefore, make sure your project supports [url=$DOCS_URL/tutorials/rendering/multiple_resolutions.html]multiple resolutions[/url] when enabling full screen mode.*/
	WindowModeFullscreen WindowMode = 3
/*A single window full screen mode. This mode has less overhead, but only one window can be open on a given screen at a time (opening a child window or application switching will trigger a full screen transition).
Full screen window covers the entire display area of a screen and has no border or decorations. The display's video mode is not changed.
[b]On Windows:[/b] Depending on video driver, full screen transition might cause screens to go black for a moment.
[b]On macOS:[/b] A new desktop is used to display the running project. Exclusive full screen mode prevents Dock and Menu from showing up when the mouse pointer is hovering the edge of the screen.
[b]On Linux (X11):[/b] Exclusive full screen mode bypasses compositor.
[b]Note:[/b] Regardless of the platform, enabling full screen will change the window size to match the monitor's size. Therefore, make sure your project supports [url=$DOCS_URL/tutorials/rendering/multiple_resolutions.html]multiple resolutions[/url] when enabling full screen mode.*/
	WindowModeExclusiveFullscreen WindowMode = 4
)
type WindowFlags = classdb.DisplayServerWindowFlags

const (
/*The window can't be resized by dragging its resize grip. It's still possible to resize the window using [method window_set_size]. This flag is ignored for full screen windows.*/
	WindowFlagResizeDisabled WindowFlags = 0
/*The window do not have native title bar and other decorations. This flag is ignored for full-screen windows.*/
	WindowFlagBorderless WindowFlags = 1
/*The window is floating on top of all other windows. This flag is ignored for full-screen windows.*/
	WindowFlagAlwaysOnTop WindowFlags = 2
/*The window background can be transparent.
[b]Note:[/b] This flag has no effect if [method is_window_transparency_available] returns [code]false[/code].
[b]Note:[/b] Transparency support is implemented on Linux (X11/Wayland), macOS, and Windows, but availability might vary depending on GPU driver, display manager, and compositor capabilities.*/
	WindowFlagTransparent WindowFlags = 3
/*The window can't be focused. No-focus window will ignore all input, except mouse clicks.*/
	WindowFlagNoFocus WindowFlags = 4
/*Window is part of menu or [OptionButton] dropdown. This flag can't be changed when the window is visible. An active popup window will exclusively receive all input, without stealing focus from its parent. Popup windows are automatically closed when uses click outside it, or when an application is switched. Popup window must have transient parent set (see [method window_set_transient]).*/
	WindowFlagPopup WindowFlags = 5
/*Window content is expanded to the full size of the window. Unlike borderless window, the frame is left intact and can be used to resize the window, title bar is transparent, but have minimize/maximize/close buttons.
Use [method window_set_window_buttons_offset] to adjust minimize/maximize/close buttons offset.
Use [method window_get_safe_title_margins] to determine area under the title bar that is not covered by decorations.
[b]Note:[/b] This flag is implemented only on macOS.*/
	WindowFlagExtendToTitle WindowFlags = 6
/*All mouse events are passed to the underlying window of the same application.*/
	WindowFlagMousePassthrough WindowFlags = 7
/*Max value of the [enum WindowFlags].*/
	WindowFlagMax WindowFlags = 8
)
type WindowEvent = classdb.DisplayServerWindowEvent

const (
/*Sent when the mouse pointer enters the window.*/
	WindowEventMouseEnter WindowEvent = 0
/*Sent when the mouse pointer exits the window.*/
	WindowEventMouseExit WindowEvent = 1
/*Sent when the window grabs focus.*/
	WindowEventFocusIn WindowEvent = 2
/*Sent when the window loses focus.*/
	WindowEventFocusOut WindowEvent = 3
/*Sent when the user has attempted to close the window (e.g. close button is pressed).*/
	WindowEventCloseRequest WindowEvent = 4
/*Sent when the device "Back" button is pressed.
[b]Note:[/b] This event is implemented only on Android.*/
	WindowEventGoBackRequest WindowEvent = 5
/*Sent when the window is moved to the display with different DPI, or display DPI is changed.
[b]Note:[/b] This flag is implemented only on macOS.*/
	WindowEventDpiChange WindowEvent = 6
/*Sent when the window title bar decoration is changed (e.g. [constant WINDOW_FLAG_EXTEND_TO_TITLE] is set or window entered/exited full screen mode).
[b]Note:[/b] This flag is implemented only on macOS.*/
	WindowEventTitlebarChange WindowEvent = 7
)
type VSyncMode = classdb.DisplayServerVSyncMode

const (
/*No vertical synchronization, which means the engine will display frames as fast as possible (tearing may be visible). Framerate is unlimited (regardless of [member Engine.max_fps]).*/
	VsyncDisabled VSyncMode = 0
/*Default vertical synchronization mode, the image is displayed only on vertical blanking intervals (no tearing is visible). Framerate is limited by the monitor refresh rate (regardless of [member Engine.max_fps]).*/
	VsyncEnabled VSyncMode = 1
/*Behaves like [constant VSYNC_DISABLED] when the framerate drops below the screen's refresh rate to reduce stuttering (tearing may be visible). Otherwise, vertical synchronization is enabled to avoid tearing. Framerate is limited by the monitor refresh rate (regardless of [member Engine.max_fps]). Behaves like [constant VSYNC_ENABLED] when using the Compatibility rendering method.*/
	VsyncAdaptive VSyncMode = 2
/*Displays the most recent image in the queue on vertical blanking intervals, while rendering to the other images (no tearing is visible). Framerate is unlimited (regardless of [member Engine.max_fps]).
Although not guaranteed, the images can be rendered as fast as possible, which may reduce input lag (also called "Fast" V-Sync mode). [constant VSYNC_MAILBOX] works best when at least twice as many frames as the display refresh rate are rendered. Behaves like [constant VSYNC_ENABLED] when using the Compatibility rendering method.*/
	VsyncMailbox VSyncMode = 3
)
type HandleType = classdb.DisplayServerHandleType

const (
/*Display handle:
- Linux (X11): [code]X11::Display*[/code] for the display.
- Android: [code]EGLDisplay[/code] for the display.*/
	DisplayHandle HandleType = 0
/*Window handle:
- Windows: [code]HWND[/code] for the window.
- Linux (X11): [code]X11::Window*[/code] for the window.
- macOS: [code]NSWindow*[/code] for the window.
- iOS: [code]UIViewController*[/code] for the view controller.
- Android: [code]jObject[/code] for the activity.*/
	WindowHandle HandleType = 1
/*Window view:
- Windows: [code]HDC[/code] for the window (only with the GL Compatibility renderer).
- macOS: [code]NSView*[/code] for the window main view.
- iOS: [code]UIView*[/code] for the window main view.*/
	WindowView HandleType = 2
/*OpenGL context (only with the GL Compatibility renderer):
- Windows: [code]HGLRC[/code] for the window (native GL), or [code]EGLContext[/code] for the window (ANGLE).
- Linux (X11): [code]GLXContext*[/code] for the window.
- macOS: [code]NSOpenGLContext*[/code] for the window (native GL), or [code]EGLContext[/code] for the window (ANGLE).
- Android: [code]EGLContext[/code] for the window.*/
	OpenglContext HandleType = 3
)
type TTSUtteranceEvent = classdb.DisplayServerTTSUtteranceEvent

const (
/*Utterance has begun to be spoken.*/
	TtsUtteranceStarted TTSUtteranceEvent = 0
/*Utterance was successfully finished.*/
	TtsUtteranceEnded TTSUtteranceEvent = 1
/*Utterance was canceled, or TTS service was unable to process it.*/
	TtsUtteranceCanceled TTSUtteranceEvent = 2
/*Utterance reached a word or sentence boundary.*/
	TtsUtteranceBoundary TTSUtteranceEvent = 3
)
