package Input

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Vector3"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
The [Input] singleton handles key presses, mouse buttons and movement, gamepads, and input actions. Actions and their events can be set in the [b]Input Map[/b] tab in [b]Project > Project Settings[/b], or with the [InputMap] class.
[b]Note:[/b] [Input]'s methods reflect the global input state and are not affected by [method Control.accept_event] or [method Viewport.set_input_as_handled], as those methods only deal with the way input is propagated in the [SceneTree].
*/
var self objects.Input
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.Input)
	self = *(*objects.Input)(unsafe.Pointer(&obj))
}

/*
Returns [code]true[/code] if any action, key, joypad button, or mouse button is being pressed. This will also return [code]true[/code] if any action is simulated via code by calling [method action_press].
*/
func IsAnythingPressed() bool {
	once.Do(singleton)
	return bool(class(self).IsAnythingPressed())
}

/*
Returns [code]true[/code] if you are pressing the Latin key in the current keyboard layout. You can pass a [enum Key] constant.
[method is_key_pressed] is only recommended over [method is_physical_key_pressed] in non-game applications. This ensures that shortcut keys behave as expected depending on the user's keyboard layout, as keyboard shortcuts are generally dependent on the keyboard layout in non-game applications. If in doubt, use [method is_physical_key_pressed].
[b]Note:[/b] Due to keyboard ghosting, [method is_key_pressed] may return [code]false[/code] even if one of the action's keys is pressed. See [url=$DOCS_URL/tutorials/inputs/input_examples.html#keyboard-events]Input examples[/url] in the documentation for more information.
*/
func IsKeyPressed(keycode Key) bool {
	once.Do(singleton)
	return bool(class(self).IsKeyPressed(keycode))
}

/*
Returns [code]true[/code] if you are pressing the key in the physical location on the 101/102-key US QWERTY keyboard. You can pass a [enum Key] constant.
[method is_physical_key_pressed] is recommended over [method is_key_pressed] for in-game actions, as it will make [kbd]W[/kbd]/[kbd]A[/kbd]/[kbd]S[/kbd]/[kbd]D[/kbd] layouts work regardless of the user's keyboard layout. [method is_physical_key_pressed] will also ensure that the top row number keys work on any keyboard layout. If in doubt, use [method is_physical_key_pressed].
[b]Note:[/b] Due to keyboard ghosting, [method is_physical_key_pressed] may return [code]false[/code] even if one of the action's keys is pressed. See [url=$DOCS_URL/tutorials/inputs/input_examples.html#keyboard-events]Input examples[/url] in the documentation for more information.
*/
func IsPhysicalKeyPressed(keycode Key) bool {
	once.Do(singleton)
	return bool(class(self).IsPhysicalKeyPressed(keycode))
}

/*
Returns [code]true[/code] if you are pressing the key with the [param keycode] printed on it. You can pass a [enum Key] constant or any Unicode character code.
*/
func IsKeyLabelPressed(keycode Key) bool {
	once.Do(singleton)
	return bool(class(self).IsKeyLabelPressed(keycode))
}

/*
Returns [code]true[/code] if you are pressing the mouse button specified with [enum MouseButton].
*/
func IsMouseButtonPressed(button MouseButton) bool {
	once.Do(singleton)
	return bool(class(self).IsMouseButtonPressed(button))
}

/*
Returns [code]true[/code] if you are pressing the joypad button (see [enum JoyButton]).
*/
func IsJoyButtonPressed(device int, button JoyButton) bool {
	once.Do(singleton)
	return bool(class(self).IsJoyButtonPressed(gd.Int(device), button))
}

/*
Returns [code]true[/code] if you are pressing the action event.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
[b]Note:[/b] Due to keyboard ghosting, [method is_action_pressed] may return [code]false[/code] even if one of the action's keys is pressed. See [url=$DOCS_URL/tutorials/inputs/input_examples.html#keyboard-events]Input examples[/url] in the documentation for more information.
*/
func IsActionPressed(action string) bool {
	once.Do(singleton)
	return bool(class(self).IsActionPressed(gd.NewStringName(action), false))
}

/*
Returns [code]true[/code] when the user has [i]started[/i] pressing the action event in the current frame or physics tick. It will only return [code]true[/code] on the frame or tick that the user pressed down the button.
This is useful for code that needs to run only once when an action is pressed, instead of every frame while it's pressed.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
[b]Note:[/b] Returning [code]true[/code] does not imply that the action is [i]still[/i] pressed. An action can be pressed and released again rapidly, and [code]true[/code] will still be returned so as not to miss input.
[b]Note:[/b] Due to keyboard ghosting, [method is_action_just_pressed] may return [code]false[/code] even if one of the action's keys is pressed. See [url=$DOCS_URL/tutorials/inputs/input_examples.html#keyboard-events]Input examples[/url] in the documentation for more information.
[b]Note:[/b] During input handling (e.g. [method Node._input]), use [method InputEvent.is_action_pressed] instead to query the action state of the current event.
*/
func IsActionJustPressed(action string) bool {
	once.Do(singleton)
	return bool(class(self).IsActionJustPressed(gd.NewStringName(action), false))
}

/*
Returns [code]true[/code] when the user [i]stops[/i] pressing the action event in the current frame or physics tick. It will only return [code]true[/code] on the frame or tick that the user releases the button.
[b]Note:[/b] Returning [code]true[/code] does not imply that the action is [i]still[/i] not pressed. An action can be released and pressed again rapidly, and [code]true[/code] will still be returned so as not to miss input.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
[b]Note:[/b] During input handling (e.g. [method Node._input]), use [method InputEvent.is_action_released] instead to query the action state of the current event.
*/
func IsActionJustReleased(action string) bool {
	once.Do(singleton)
	return bool(class(self).IsActionJustReleased(gd.NewStringName(action), false))
}

/*
Returns a value between 0 and 1 representing the intensity of the given action. In a joypad, for example, the further away the axis (analog sticks or L2, R2 triggers) is from the dead zone, the closer the value will be to 1. If the action is mapped to a control that has no axis such as the keyboard, the value returned will be 0 or 1.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
func GetActionStrength(action string) Float.X {
	once.Do(singleton)
	return Float.X(Float.X(class(self).GetActionStrength(gd.NewStringName(action), false)))
}

/*
Returns a value between 0 and 1 representing the raw intensity of the given action, ignoring the action's deadzone. In most cases, you should use [method get_action_strength] instead.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
func GetActionRawStrength(action string) Float.X {
	once.Do(singleton)
	return Float.X(Float.X(class(self).GetActionRawStrength(gd.NewStringName(action), false)))
}

/*
Get axis input by specifying two actions, one negative and one positive.
This is a shorthand for writing [code]Input.get_action_strength("positive_action") - Input.get_action_strength("negative_action")[/code].
*/
func GetAxis(negative_action string, positive_action string) Float.X {
	once.Do(singleton)
	return Float.X(Float.X(class(self).GetAxis(gd.NewStringName(negative_action), gd.NewStringName(positive_action))))
}

/*
Gets an input vector by specifying four actions for the positive and negative X and Y axes.
This method is useful when getting vector input, such as from a joystick, directional pad, arrows, or WASD. The vector has its length limited to 1 and has a circular deadzone, which is useful for using vector input as movement.
By default, the deadzone is automatically calculated from the average of the action deadzones. However, you can override the deadzone to be whatever you want (on the range of 0 to 1).
*/
func GetVector(negative_x string, positive_x string, negative_y string, positive_y string) Vector2.XY {
	once.Do(singleton)
	return Vector2.XY(class(self).GetVector(gd.NewStringName(negative_x), gd.NewStringName(positive_x), gd.NewStringName(negative_y), gd.NewStringName(positive_y), gd.Float(-1.0)))
}

/*
Adds a new mapping entry (in SDL2 format) to the mapping database. Optionally update already connected devices.
*/
func AddJoyMapping(mapping string) {
	once.Do(singleton)
	class(self).AddJoyMapping(gd.NewString(mapping), false)
}

/*
Removes all mappings from the internal database that match the given GUID.
*/
func RemoveJoyMapping(guid string) {
	once.Do(singleton)
	class(self).RemoveJoyMapping(gd.NewString(guid))
}

/*
Returns [code]true[/code] if the system knows the specified device. This means that it sets all button and axis indices. Unknown joypads are not expected to match these constants, but you can still retrieve events from them.
*/
func IsJoyKnown(device int) bool {
	once.Do(singleton)
	return bool(class(self).IsJoyKnown(gd.Int(device)))
}

/*
Returns the current value of the joypad axis at given index (see [enum JoyAxis]).
*/
func GetJoyAxis(device int, axis JoyAxis) Float.X {
	once.Do(singleton)
	return Float.X(Float.X(class(self).GetJoyAxis(gd.Int(device), axis)))
}

/*
Returns the name of the joypad at the specified device index, e.g. [code]PS4 Controller[/code]. Godot uses the [url=https://github.com/gabomdq/SDL_GameControllerDB]SDL2 game controller database[/url] to determine gamepad names.
*/
func GetJoyName(device int) string {
	once.Do(singleton)
	return string(class(self).GetJoyName(gd.Int(device)).String())
}

/*
Returns an SDL2-compatible device GUID on platforms that use gamepad remapping, e.g. [code]030000004c050000c405000000010000[/code]. Returns [code]"Default Gamepad"[/code] otherwise. Godot uses the [url=https://github.com/gabomdq/SDL_GameControllerDB]SDL2 game controller database[/url] to determine gamepad names and mappings based on this GUID.
*/
func GetJoyGuid(device int) string {
	once.Do(singleton)
	return string(class(self).GetJoyGuid(gd.Int(device)).String())
}

/*
Returns a dictionary with extra platform-specific information about the device, e.g. the raw gamepad name from the OS or the Steam Input index.
On Windows the dictionary contains the following fields:
[code]xinput_index[/code]: The index of the controller in the XInput system.
On Linux:
[code]raw_name[/code]: The name of the controller as it came from the OS, before getting renamed by the godot controller database.
[code]vendor_id[/code]: The USB vendor ID of the device.
[code]product_id[/code]: The USB product ID of the device.
[code]steam_input_index[/code]: The Steam Input gamepad index, if the device is not a Steam Input device this key won't be present.
*/
func GetJoyInfo(device int) Dictionary.Any {
	once.Do(singleton)
	return Dictionary.Any(class(self).GetJoyInfo(gd.Int(device)))
}

/*
Queries whether an input device should be ignored or not. Devices can be ignored by setting the environment variable [code]SDL_GAMECONTROLLER_IGNORE_DEVICES[/code]. Read the [url=https://wiki.libsdl.org/SDL2]SDL documentation[/url] for more information.
[b]Note:[/b] Some 3rd party tools can contribute to the list of ignored devices. For example, [i]SteamInput[/i] creates virtual devices from physical devices for remapping purposes. To avoid handling the same input device twice, the original device is added to the ignore list.
*/
func ShouldIgnoreDevice(vendor_id int, product_id int) bool {
	once.Do(singleton)
	return bool(class(self).ShouldIgnoreDevice(gd.Int(vendor_id), gd.Int(product_id)))
}

/*
Returns an [Array] containing the device IDs of all currently connected joypads.
*/
func GetConnectedJoypads() gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).GetConnectedJoypads())
}

/*
Returns the strength of the joypad vibration: x is the strength of the weak motor, and y is the strength of the strong motor.
*/
func GetJoyVibrationStrength(device int) Vector2.XY {
	once.Do(singleton)
	return Vector2.XY(class(self).GetJoyVibrationStrength(gd.Int(device)))
}

/*
Returns the duration of the current vibration effect in seconds.
*/
func GetJoyVibrationDuration(device int) Float.X {
	once.Do(singleton)
	return Float.X(Float.X(class(self).GetJoyVibrationDuration(gd.Int(device))))
}

/*
Starts to vibrate the joypad. Joypads usually come with two rumble motors, a strong and a weak one. [param weak_magnitude] is the strength of the weak motor (between 0 and 1) and [param strong_magnitude] is the strength of the strong motor (between 0 and 1). [param duration] is the duration of the effect in seconds (a duration of 0 will try to play the vibration indefinitely). The vibration can be stopped early by calling [method stop_joy_vibration].
[b]Note:[/b] Not every hardware is compatible with long effect durations; it is recommended to restart an effect if it has to be played for more than a few seconds.
[b]Note:[/b] For macOS, vibration is only supported in macOS 11 and later.
*/
func StartJoyVibration(device int, weak_magnitude Float.X, strong_magnitude Float.X) {
	once.Do(singleton)
	class(self).StartJoyVibration(gd.Int(device), gd.Float(weak_magnitude), gd.Float(strong_magnitude), gd.Float(0))
}

/*
Stops the vibration of the joypad started with [method start_joy_vibration].
*/
func StopJoyVibration(device int) {
	once.Do(singleton)
	class(self).StopJoyVibration(gd.Int(device))
}

/*
Vibrate the handheld device for the specified duration in milliseconds.
[param amplitude] is the strength of the vibration, as a value between [code]0.0[/code] and [code]1.0[/code]. If set to [code]-1.0[/code], the default vibration strength of the device is used.
[b]Note:[/b] This method is implemented on Android, iOS, and Web. It has no effect on other platforms.
[b]Note:[/b] For Android, [method vibrate_handheld] requires enabling the [code]VIBRATE[/code] permission in the export preset. Otherwise, [method vibrate_handheld] will have no effect.
[b]Note:[/b] For iOS, specifying the duration is only supported in iOS 13 and later.
[b]Note:[/b] For Web, the amplitude cannot be changed.
[b]Note:[/b] Some web browsers such as Safari and Firefox for Android do not support [method vibrate_handheld].
*/
func VibrateHandheld() {
	once.Do(singleton)
	class(self).VibrateHandheld(gd.Int(500), gd.Float(-1.0))
}

/*
Returns the gravity in m/s² of the device's accelerometer sensor, if the device has one. Otherwise, the method returns [constant Vector3.ZERO].
[b]Note:[/b] This method only works on Android and iOS. On other platforms, it always returns [constant Vector3.ZERO].
*/
func GetGravity() Vector3.XYZ {
	once.Do(singleton)
	return Vector3.XYZ(class(self).GetGravity())
}

/*
Returns the acceleration in m/s² of the device's accelerometer sensor, if the device has one. Otherwise, the method returns [constant Vector3.ZERO].
Note this method returns an empty [Vector3] when running from the editor even when your device has an accelerometer. You must export your project to a supported device to read values from the accelerometer.
[b]Note:[/b] This method only works on Android and iOS. On other platforms, it always returns [constant Vector3.ZERO].
*/
func GetAccelerometer() Vector3.XYZ {
	once.Do(singleton)
	return Vector3.XYZ(class(self).GetAccelerometer())
}

/*
Returns the magnetic field strength in micro-Tesla for all axes of the device's magnetometer sensor, if the device has one. Otherwise, the method returns [constant Vector3.ZERO].
[b]Note:[/b] This method only works on Android and iOS. On other platforms, it always returns [constant Vector3.ZERO].
*/
func GetMagnetometer() Vector3.XYZ {
	once.Do(singleton)
	return Vector3.XYZ(class(self).GetMagnetometer())
}

/*
Returns the rotation rate in rad/s around a device's X, Y, and Z axes of the gyroscope sensor, if the device has one. Otherwise, the method returns [constant Vector3.ZERO].
[b]Note:[/b] This method only works on Android and iOS. On other platforms, it always returns [constant Vector3.ZERO].
*/
func GetGyroscope() Vector3.XYZ {
	once.Do(singleton)
	return Vector3.XYZ(class(self).GetGyroscope())
}

/*
Sets the gravity value of the accelerometer sensor. Can be used for debugging on devices without a hardware sensor, for example in an editor on a PC.
[b]Note:[/b] This value can be immediately overwritten by the hardware sensor value on Android and iOS.
*/
func SetGravity(value Vector3.XYZ) {
	once.Do(singleton)
	class(self).SetGravity(gd.Vector3(value))
}

/*
Sets the acceleration value of the accelerometer sensor. Can be used for debugging on devices without a hardware sensor, for example in an editor on a PC.
[b]Note:[/b] This value can be immediately overwritten by the hardware sensor value on Android and iOS.
*/
func SetAccelerometer(value Vector3.XYZ) {
	once.Do(singleton)
	class(self).SetAccelerometer(gd.Vector3(value))
}

/*
Sets the value of the magnetic field of the magnetometer sensor. Can be used for debugging on devices without a hardware sensor, for example in an editor on a PC.
[b]Note:[/b] This value can be immediately overwritten by the hardware sensor value on Android and iOS.
*/
func SetMagnetometer(value Vector3.XYZ) {
	once.Do(singleton)
	class(self).SetMagnetometer(gd.Vector3(value))
}

/*
Sets the value of the rotation rate of the gyroscope sensor. Can be used for debugging on devices without a hardware sensor, for example in an editor on a PC.
[b]Note:[/b] This value can be immediately overwritten by the hardware sensor value on Android and iOS.
*/
func SetGyroscope(value Vector3.XYZ) {
	once.Do(singleton)
	class(self).SetGyroscope(gd.Vector3(value))
}

/*
Returns the last mouse velocity. To provide a precise and jitter-free velocity, mouse velocity is only calculated every 0.1s. Therefore, mouse velocity will lag mouse movements.
*/
func GetLastMouseVelocity() Vector2.XY {
	once.Do(singleton)
	return Vector2.XY(class(self).GetLastMouseVelocity())
}

/*
Returns the last mouse velocity in screen coordinates. To provide a precise and jitter-free velocity, mouse velocity is only calculated every 0.1s. Therefore, mouse velocity will lag mouse movements.
*/
func GetLastMouseScreenVelocity() Vector2.XY {
	once.Do(singleton)
	return Vector2.XY(class(self).GetLastMouseScreenVelocity())
}

/*
Returns mouse buttons as a bitmask. If multiple mouse buttons are pressed at the same time, the bits are added together. Equivalent to [method DisplayServer.mouse_get_button_state].
*/
func GetMouseButtonMask() MouseButtonMask {
	once.Do(singleton)
	return MouseButtonMask(class(self).GetMouseButtonMask())
}

/*
Sets the mouse position to the specified vector, provided in pixels and relative to an origin at the upper left corner of the currently focused Window Manager game window.
Mouse position is clipped to the limits of the screen resolution, or to the limits of the game window if [enum MouseMode] is set to [constant MOUSE_MODE_CONFINED] or [constant MOUSE_MODE_CONFINED_HIDDEN].
[b]Note:[/b] [method warp_mouse] is only supported on Windows, macOS and Linux. It has no effect on Android, iOS and Web.
*/
func WarpMouse(position Vector2.XY) {
	once.Do(singleton)
	class(self).WarpMouse(gd.Vector2(position))
}

/*
This will simulate pressing the specified action.
The strength can be used for non-boolean actions, it's ranged between 0 and 1 representing the intensity of the given action.
[b]Note:[/b] This method will not cause any [method Node._input] calls. It is intended to be used with [method is_action_pressed] and [method is_action_just_pressed]. If you want to simulate [code]_input[/code], use [method parse_input_event] instead.
*/
func ActionPress(action string) {
	once.Do(singleton)
	class(self).ActionPress(gd.NewStringName(action), gd.Float(1.0))
}

/*
If the specified action is already pressed, this will release it.
*/
func ActionRelease(action string) {
	once.Do(singleton)
	class(self).ActionRelease(gd.NewStringName(action))
}

/*
Sets the default cursor shape to be used in the viewport instead of [constant CURSOR_ARROW].
[b]Note:[/b] If you want to change the default cursor shape for [Control]'s nodes, use [member Control.mouse_default_cursor_shape] instead.
[b]Note:[/b] This method generates an [InputEventMouseMotion] to update cursor immediately.
*/
func SetDefaultCursorShape() {
	once.Do(singleton)
	class(self).SetDefaultCursorShape(0)
}

/*
Returns the currently assigned cursor shape (see [enum CursorShape]).
*/
func GetCurrentCursorShape() classdb.InputCursorShape {
	once.Do(singleton)
	return classdb.InputCursorShape(class(self).GetCurrentCursorShape())
}

/*
Sets a custom mouse cursor image, which is only visible inside the game window. The hotspot can also be specified. Passing [code]null[/code] to the image parameter resets to the system cursor. See [enum CursorShape] for the list of shapes.
[param image] can be either [Texture2D] or [Image] and its size must be lower than or equal to 256×256. To avoid rendering issues, sizes lower than or equal to 128×128 are recommended.
[param hotspot] must be within [param image]'s size.
[b]Note:[/b] [AnimatedTexture]s aren't supported as custom mouse cursors. If using an [AnimatedTexture], only the first frame will be displayed.
[b]Note:[/b] The [b]Lossless[/b], [b]Lossy[/b] or [b]Uncompressed[/b] compression modes are recommended. The [b]Video RAM[/b] compression mode can be used, but it will be decompressed on the CPU, which means loading times are slowed down and no memory is saved compared to lossless modes.
[b]Note:[/b] On the web platform, the maximum allowed cursor image size is 128×128. Cursor images larger than 32×32 will also only be displayed if the mouse cursor image is entirely located within the page for [url=https://chromestatus.com/feature/5825971391299584]security reasons[/url].
*/
func SetCustomMouseCursor(image objects.Resource) {
	once.Do(singleton)
	class(self).SetCustomMouseCursor(image, 0, gd.Vector2(gd.Vector2{0, 0}))
}

/*
Feeds an [InputEvent] to the game. Can be used to artificially trigger input events from code. Also generates [method Node._input] calls.
[b]Example:[/b]
[codeblocks]
[gdscript]
var cancel_event = InputEventAction.new()
cancel_event.action = "ui_cancel"
cancel_event.pressed = true
Input.parse_input_event(cancel_event)
[/gdscript]
[csharp]
var cancelEvent = new InputEventAction();
cancelEvent.Action = "ui_cancel";
cancelEvent.Pressed = true;
Input.ParseInputEvent(cancelEvent);
[/csharp]
[/codeblocks]
[b]Note:[/b] Calling this function has no influence on the operating system. So for example sending an [InputEventMouseMotion] will not move the OS mouse cursor to the specified position (use [method warp_mouse] instead) and sending [kbd]Alt/Cmd + Tab[/kbd] as [InputEventKey] won't toggle between active windows.
*/
func ParseInputEvent(event objects.InputEvent) {
	once.Do(singleton)
	class(self).ParseInputEvent(event)
}

/*
Sends all input events which are in the current buffer to the game loop. These events may have been buffered as a result of accumulated input ([member use_accumulated_input]) or agile input flushing ([member ProjectSettings.input_devices/buffering/agile_event_flushing]).
The engine will already do this itself at key execution points (at least once per frame). However, this can be useful in advanced cases where you want precise control over the timing of event handling.
*/
func FlushBufferedEvents() {
	once.Do(singleton)
	class(self).FlushBufferedEvents()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]classdb.Input

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

func MouseMode() classdb.InputMouseMode {
	return classdb.InputMouseMode(class(self).GetMouseMode())
}

func SetMouseMode(value classdb.InputMouseMode) {
	class(self).SetMouseMode(value)
}

func UseAccumulatedInput() bool {
	return bool(class(self).IsUsingAccumulatedInput())
}

func SetUseAccumulatedInput(value bool) {
	class(self).SetUseAccumulatedInput(value)
}

func EmulateMouseFromTouch() bool {
	return bool(class(self).IsEmulatingMouseFromTouch())
}

func SetEmulateMouseFromTouch(value bool) {
	class(self).SetEmulateMouseFromTouch(value)
}

func EmulateTouchFromMouse() bool {
	return bool(class(self).IsEmulatingTouchFromMouse())
}

func SetEmulateTouchFromMouse(value bool) {
	class(self).SetEmulateTouchFromMouse(value)
}

/*
Returns [code]true[/code] if any action, key, joypad button, or mouse button is being pressed. This will also return [code]true[/code] if any action is simulated via code by calling [method action_press].
*/
//go:nosplit
func (self class) IsAnythingPressed() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_is_anything_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if you are pressing the Latin key in the current keyboard layout. You can pass a [enum Key] constant.
[method is_key_pressed] is only recommended over [method is_physical_key_pressed] in non-game applications. This ensures that shortcut keys behave as expected depending on the user's keyboard layout, as keyboard shortcuts are generally dependent on the keyboard layout in non-game applications. If in doubt, use [method is_physical_key_pressed].
[b]Note:[/b] Due to keyboard ghosting, [method is_key_pressed] may return [code]false[/code] even if one of the action's keys is pressed. See [url=$DOCS_URL/tutorials/inputs/input_examples.html#keyboard-events]Input examples[/url] in the documentation for more information.
*/
//go:nosplit
func (self class) IsKeyPressed(keycode Key) bool {
	var frame = callframe.New()
	callframe.Arg(frame, keycode)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_is_key_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if you are pressing the key in the physical location on the 101/102-key US QWERTY keyboard. You can pass a [enum Key] constant.
[method is_physical_key_pressed] is recommended over [method is_key_pressed] for in-game actions, as it will make [kbd]W[/kbd]/[kbd]A[/kbd]/[kbd]S[/kbd]/[kbd]D[/kbd] layouts work regardless of the user's keyboard layout. [method is_physical_key_pressed] will also ensure that the top row number keys work on any keyboard layout. If in doubt, use [method is_physical_key_pressed].
[b]Note:[/b] Due to keyboard ghosting, [method is_physical_key_pressed] may return [code]false[/code] even if one of the action's keys is pressed. See [url=$DOCS_URL/tutorials/inputs/input_examples.html#keyboard-events]Input examples[/url] in the documentation for more information.
*/
//go:nosplit
func (self class) IsPhysicalKeyPressed(keycode Key) bool {
	var frame = callframe.New()
	callframe.Arg(frame, keycode)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_is_physical_key_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if you are pressing the key with the [param keycode] printed on it. You can pass a [enum Key] constant or any Unicode character code.
*/
//go:nosplit
func (self class) IsKeyLabelPressed(keycode Key) bool {
	var frame = callframe.New()
	callframe.Arg(frame, keycode)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_is_key_label_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if you are pressing the mouse button specified with [enum MouseButton].
*/
//go:nosplit
func (self class) IsMouseButtonPressed(button MouseButton) bool {
	var frame = callframe.New()
	callframe.Arg(frame, button)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_is_mouse_button_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if you are pressing the joypad button (see [enum JoyButton]).
*/
//go:nosplit
func (self class) IsJoyButtonPressed(device gd.Int, button JoyButton) bool {
	var frame = callframe.New()
	callframe.Arg(frame, device)
	callframe.Arg(frame, button)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_is_joy_button_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if you are pressing the action event.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
[b]Note:[/b] Due to keyboard ghosting, [method is_action_pressed] may return [code]false[/code] even if one of the action's keys is pressed. See [url=$DOCS_URL/tutorials/inputs/input_examples.html#keyboard-events]Input examples[/url] in the documentation for more information.
*/
//go:nosplit
func (self class) IsActionPressed(action gd.StringName, exact_match bool) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	callframe.Arg(frame, exact_match)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_is_action_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] when the user has [i]started[/i] pressing the action event in the current frame or physics tick. It will only return [code]true[/code] on the frame or tick that the user pressed down the button.
This is useful for code that needs to run only once when an action is pressed, instead of every frame while it's pressed.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
[b]Note:[/b] Returning [code]true[/code] does not imply that the action is [i]still[/i] pressed. An action can be pressed and released again rapidly, and [code]true[/code] will still be returned so as not to miss input.
[b]Note:[/b] Due to keyboard ghosting, [method is_action_just_pressed] may return [code]false[/code] even if one of the action's keys is pressed. See [url=$DOCS_URL/tutorials/inputs/input_examples.html#keyboard-events]Input examples[/url] in the documentation for more information.
[b]Note:[/b] During input handling (e.g. [method Node._input]), use [method InputEvent.is_action_pressed] instead to query the action state of the current event.
*/
//go:nosplit
func (self class) IsActionJustPressed(action gd.StringName, exact_match bool) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	callframe.Arg(frame, exact_match)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_is_action_just_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] when the user [i]stops[/i] pressing the action event in the current frame or physics tick. It will only return [code]true[/code] on the frame or tick that the user releases the button.
[b]Note:[/b] Returning [code]true[/code] does not imply that the action is [i]still[/i] not pressed. An action can be released and pressed again rapidly, and [code]true[/code] will still be returned so as not to miss input.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
[b]Note:[/b] During input handling (e.g. [method Node._input]), use [method InputEvent.is_action_released] instead to query the action state of the current event.
*/
//go:nosplit
func (self class) IsActionJustReleased(action gd.StringName, exact_match bool) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	callframe.Arg(frame, exact_match)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_is_action_just_released, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a value between 0 and 1 representing the intensity of the given action. In a joypad, for example, the further away the axis (analog sticks or L2, R2 triggers) is from the dead zone, the closer the value will be to 1. If the action is mapped to a control that has no axis such as the keyboard, the value returned will be 0 or 1.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
//go:nosplit
func (self class) GetActionStrength(action gd.StringName, exact_match bool) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	callframe.Arg(frame, exact_match)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_get_action_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a value between 0 and 1 representing the raw intensity of the given action, ignoring the action's deadzone. In most cases, you should use [method get_action_strength] instead.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
//go:nosplit
func (self class) GetActionRawStrength(action gd.StringName, exact_match bool) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	callframe.Arg(frame, exact_match)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_get_action_raw_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Get axis input by specifying two actions, one negative and one positive.
This is a shorthand for writing [code]Input.get_action_strength("positive_action") - Input.get_action_strength("negative_action")[/code].
*/
//go:nosplit
func (self class) GetAxis(negative_action gd.StringName, positive_action gd.StringName) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(negative_action))
	callframe.Arg(frame, pointers.Get(positive_action))
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_get_axis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets an input vector by specifying four actions for the positive and negative X and Y axes.
This method is useful when getting vector input, such as from a joystick, directional pad, arrows, or WASD. The vector has its length limited to 1 and has a circular deadzone, which is useful for using vector input as movement.
By default, the deadzone is automatically calculated from the average of the action deadzones. However, you can override the deadzone to be whatever you want (on the range of 0 to 1).
*/
//go:nosplit
func (self class) GetVector(negative_x gd.StringName, positive_x gd.StringName, negative_y gd.StringName, positive_y gd.StringName, deadzone gd.Float) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(negative_x))
	callframe.Arg(frame, pointers.Get(positive_x))
	callframe.Arg(frame, pointers.Get(negative_y))
	callframe.Arg(frame, pointers.Get(positive_y))
	callframe.Arg(frame, deadzone)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_get_vector, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a new mapping entry (in SDL2 format) to the mapping database. Optionally update already connected devices.
*/
//go:nosplit
func (self class) AddJoyMapping(mapping gd.String, update_existing bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mapping))
	callframe.Arg(frame, update_existing)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_add_joy_mapping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes all mappings from the internal database that match the given GUID.
*/
//go:nosplit
func (self class) RemoveJoyMapping(guid gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(guid))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_remove_joy_mapping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the system knows the specified device. This means that it sets all button and axis indices. Unknown joypads are not expected to match these constants, but you can still retrieve events from them.
*/
//go:nosplit
func (self class) IsJoyKnown(device gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, device)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_is_joy_known, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the current value of the joypad axis at given index (see [enum JoyAxis]).
*/
//go:nosplit
func (self class) GetJoyAxis(device gd.Int, axis JoyAxis) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, device)
	callframe.Arg(frame, axis)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_get_joy_axis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the name of the joypad at the specified device index, e.g. [code]PS4 Controller[/code]. Godot uses the [url=https://github.com/gabomdq/SDL_GameControllerDB]SDL2 game controller database[/url] to determine gamepad names.
*/
//go:nosplit
func (self class) GetJoyName(device gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, device)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_get_joy_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an SDL2-compatible device GUID on platforms that use gamepad remapping, e.g. [code]030000004c050000c405000000010000[/code]. Returns [code]"Default Gamepad"[/code] otherwise. Godot uses the [url=https://github.com/gabomdq/SDL_GameControllerDB]SDL2 game controller database[/url] to determine gamepad names and mappings based on this GUID.
*/
//go:nosplit
func (self class) GetJoyGuid(device gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, device)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_get_joy_guid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a dictionary with extra platform-specific information about the device, e.g. the raw gamepad name from the OS or the Steam Input index.
On Windows the dictionary contains the following fields:
[code]xinput_index[/code]: The index of the controller in the XInput system.
On Linux:
[code]raw_name[/code]: The name of the controller as it came from the OS, before getting renamed by the godot controller database.
[code]vendor_id[/code]: The USB vendor ID of the device.
[code]product_id[/code]: The USB product ID of the device.
[code]steam_input_index[/code]: The Steam Input gamepad index, if the device is not a Steam Input device this key won't be present.
*/
//go:nosplit
func (self class) GetJoyInfo(device gd.Int) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, device)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_get_joy_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Queries whether an input device should be ignored or not. Devices can be ignored by setting the environment variable [code]SDL_GAMECONTROLLER_IGNORE_DEVICES[/code]. Read the [url=https://wiki.libsdl.org/SDL2]SDL documentation[/url] for more information.
[b]Note:[/b] Some 3rd party tools can contribute to the list of ignored devices. For example, [i]SteamInput[/i] creates virtual devices from physical devices for remapping purposes. To avoid handling the same input device twice, the original device is added to the ignore list.
*/
//go:nosplit
func (self class) ShouldIgnoreDevice(vendor_id gd.Int, product_id gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, vendor_id)
	callframe.Arg(frame, product_id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_should_ignore_device, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an [Array] containing the device IDs of all currently connected joypads.
*/
//go:nosplit
func (self class) GetConnectedJoypads() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_get_connected_joypads, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the strength of the joypad vibration: x is the strength of the weak motor, and y is the strength of the strong motor.
*/
//go:nosplit
func (self class) GetJoyVibrationStrength(device gd.Int) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, device)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_get_joy_vibration_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the duration of the current vibration effect in seconds.
*/
//go:nosplit
func (self class) GetJoyVibrationDuration(device gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, device)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_get_joy_vibration_duration, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Starts to vibrate the joypad. Joypads usually come with two rumble motors, a strong and a weak one. [param weak_magnitude] is the strength of the weak motor (between 0 and 1) and [param strong_magnitude] is the strength of the strong motor (between 0 and 1). [param duration] is the duration of the effect in seconds (a duration of 0 will try to play the vibration indefinitely). The vibration can be stopped early by calling [method stop_joy_vibration].
[b]Note:[/b] Not every hardware is compatible with long effect durations; it is recommended to restart an effect if it has to be played for more than a few seconds.
[b]Note:[/b] For macOS, vibration is only supported in macOS 11 and later.
*/
//go:nosplit
func (self class) StartJoyVibration(device gd.Int, weak_magnitude gd.Float, strong_magnitude gd.Float, duration gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, device)
	callframe.Arg(frame, weak_magnitude)
	callframe.Arg(frame, strong_magnitude)
	callframe.Arg(frame, duration)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_start_joy_vibration, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Stops the vibration of the joypad started with [method start_joy_vibration].
*/
//go:nosplit
func (self class) StopJoyVibration(device gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, device)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_stop_joy_vibration, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Vibrate the handheld device for the specified duration in milliseconds.
[param amplitude] is the strength of the vibration, as a value between [code]0.0[/code] and [code]1.0[/code]. If set to [code]-1.0[/code], the default vibration strength of the device is used.
[b]Note:[/b] This method is implemented on Android, iOS, and Web. It has no effect on other platforms.
[b]Note:[/b] For Android, [method vibrate_handheld] requires enabling the [code]VIBRATE[/code] permission in the export preset. Otherwise, [method vibrate_handheld] will have no effect.
[b]Note:[/b] For iOS, specifying the duration is only supported in iOS 13 and later.
[b]Note:[/b] For Web, the amplitude cannot be changed.
[b]Note:[/b] Some web browsers such as Safari and Firefox for Android do not support [method vibrate_handheld].
*/
//go:nosplit
func (self class) VibrateHandheld(duration_ms gd.Int, amplitude gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, duration_ms)
	callframe.Arg(frame, amplitude)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_vibrate_handheld, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the gravity in m/s² of the device's accelerometer sensor, if the device has one. Otherwise, the method returns [constant Vector3.ZERO].
[b]Note:[/b] This method only works on Android and iOS. On other platforms, it always returns [constant Vector3.ZERO].
*/
//go:nosplit
func (self class) GetGravity() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_get_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the acceleration in m/s² of the device's accelerometer sensor, if the device has one. Otherwise, the method returns [constant Vector3.ZERO].
Note this method returns an empty [Vector3] when running from the editor even when your device has an accelerometer. You must export your project to a supported device to read values from the accelerometer.
[b]Note:[/b] This method only works on Android and iOS. On other platforms, it always returns [constant Vector3.ZERO].
*/
//go:nosplit
func (self class) GetAccelerometer() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_get_accelerometer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the magnetic field strength in micro-Tesla for all axes of the device's magnetometer sensor, if the device has one. Otherwise, the method returns [constant Vector3.ZERO].
[b]Note:[/b] This method only works on Android and iOS. On other platforms, it always returns [constant Vector3.ZERO].
*/
//go:nosplit
func (self class) GetMagnetometer() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_get_magnetometer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the rotation rate in rad/s around a device's X, Y, and Z axes of the gyroscope sensor, if the device has one. Otherwise, the method returns [constant Vector3.ZERO].
[b]Note:[/b] This method only works on Android and iOS. On other platforms, it always returns [constant Vector3.ZERO].
*/
//go:nosplit
func (self class) GetGyroscope() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_get_gyroscope, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the gravity value of the accelerometer sensor. Can be used for debugging on devices without a hardware sensor, for example in an editor on a PC.
[b]Note:[/b] This value can be immediately overwritten by the hardware sensor value on Android and iOS.
*/
//go:nosplit
func (self class) SetGravity(value gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_set_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the acceleration value of the accelerometer sensor. Can be used for debugging on devices without a hardware sensor, for example in an editor on a PC.
[b]Note:[/b] This value can be immediately overwritten by the hardware sensor value on Android and iOS.
*/
//go:nosplit
func (self class) SetAccelerometer(value gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_set_accelerometer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the value of the magnetic field of the magnetometer sensor. Can be used for debugging on devices without a hardware sensor, for example in an editor on a PC.
[b]Note:[/b] This value can be immediately overwritten by the hardware sensor value on Android and iOS.
*/
//go:nosplit
func (self class) SetMagnetometer(value gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_set_magnetometer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the value of the rotation rate of the gyroscope sensor. Can be used for debugging on devices without a hardware sensor, for example in an editor on a PC.
[b]Note:[/b] This value can be immediately overwritten by the hardware sensor value on Android and iOS.
*/
//go:nosplit
func (self class) SetGyroscope(value gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_set_gyroscope, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the last mouse velocity. To provide a precise and jitter-free velocity, mouse velocity is only calculated every 0.1s. Therefore, mouse velocity will lag mouse movements.
*/
//go:nosplit
func (self class) GetLastMouseVelocity() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_get_last_mouse_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the last mouse velocity in screen coordinates. To provide a precise and jitter-free velocity, mouse velocity is only calculated every 0.1s. Therefore, mouse velocity will lag mouse movements.
*/
//go:nosplit
func (self class) GetLastMouseScreenVelocity() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_get_last_mouse_screen_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns mouse buttons as a bitmask. If multiple mouse buttons are pressed at the same time, the bits are added together. Equivalent to [method DisplayServer.mouse_get_button_state].
*/
//go:nosplit
func (self class) GetMouseButtonMask() MouseButtonMask {
	var frame = callframe.New()
	var r_ret = callframe.Ret[MouseButtonMask](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_get_mouse_button_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMouseMode(mode classdb.InputMouseMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_set_mouse_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMouseMode() classdb.InputMouseMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.InputMouseMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_get_mouse_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the mouse position to the specified vector, provided in pixels and relative to an origin at the upper left corner of the currently focused Window Manager game window.
Mouse position is clipped to the limits of the screen resolution, or to the limits of the game window if [enum MouseMode] is set to [constant MOUSE_MODE_CONFINED] or [constant MOUSE_MODE_CONFINED_HIDDEN].
[b]Note:[/b] [method warp_mouse] is only supported on Windows, macOS and Linux. It has no effect on Android, iOS and Web.
*/
//go:nosplit
func (self class) WarpMouse(position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_warp_mouse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
This will simulate pressing the specified action.
The strength can be used for non-boolean actions, it's ranged between 0 and 1 representing the intensity of the given action.
[b]Note:[/b] This method will not cause any [method Node._input] calls. It is intended to be used with [method is_action_pressed] and [method is_action_just_pressed]. If you want to simulate [code]_input[/code], use [method parse_input_event] instead.
*/
//go:nosplit
func (self class) ActionPress(action gd.StringName, strength gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	callframe.Arg(frame, strength)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_action_press, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
If the specified action is already pressed, this will release it.
*/
//go:nosplit
func (self class) ActionRelease(action gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_action_release, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the default cursor shape to be used in the viewport instead of [constant CURSOR_ARROW].
[b]Note:[/b] If you want to change the default cursor shape for [Control]'s nodes, use [member Control.mouse_default_cursor_shape] instead.
[b]Note:[/b] This method generates an [InputEventMouseMotion] to update cursor immediately.
*/
//go:nosplit
func (self class) SetDefaultCursorShape(shape classdb.InputCursorShape) {
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_set_default_cursor_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the currently assigned cursor shape (see [enum CursorShape]).
*/
//go:nosplit
func (self class) GetCurrentCursorShape() classdb.InputCursorShape {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.InputCursorShape](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_get_current_cursor_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets a custom mouse cursor image, which is only visible inside the game window. The hotspot can also be specified. Passing [code]null[/code] to the image parameter resets to the system cursor. See [enum CursorShape] for the list of shapes.
[param image] can be either [Texture2D] or [Image] and its size must be lower than or equal to 256×256. To avoid rendering issues, sizes lower than or equal to 128×128 are recommended.
[param hotspot] must be within [param image]'s size.
[b]Note:[/b] [AnimatedTexture]s aren't supported as custom mouse cursors. If using an [AnimatedTexture], only the first frame will be displayed.
[b]Note:[/b] The [b]Lossless[/b], [b]Lossy[/b] or [b]Uncompressed[/b] compression modes are recommended. The [b]Video RAM[/b] compression mode can be used, but it will be decompressed on the CPU, which means loading times are slowed down and no memory is saved compared to lossless modes.
[b]Note:[/b] On the web platform, the maximum allowed cursor image size is 128×128. Cursor images larger than 32×32 will also only be displayed if the mouse cursor image is entirely located within the page for [url=https://chromestatus.com/feature/5825971391299584]security reasons[/url].
*/
//go:nosplit
func (self class) SetCustomMouseCursor(image objects.Resource, shape classdb.InputCursorShape, hotspot gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(image[0])[0])
	callframe.Arg(frame, shape)
	callframe.Arg(frame, hotspot)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_set_custom_mouse_cursor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Feeds an [InputEvent] to the game. Can be used to artificially trigger input events from code. Also generates [method Node._input] calls.
[b]Example:[/b]
[codeblocks]
[gdscript]
var cancel_event = InputEventAction.new()
cancel_event.action = "ui_cancel"
cancel_event.pressed = true
Input.parse_input_event(cancel_event)
[/gdscript]
[csharp]
var cancelEvent = new InputEventAction();
cancelEvent.Action = "ui_cancel";
cancelEvent.Pressed = true;
Input.ParseInputEvent(cancelEvent);
[/csharp]
[/codeblocks]
[b]Note:[/b] Calling this function has no influence on the operating system. So for example sending an [InputEventMouseMotion] will not move the OS mouse cursor to the specified position (use [method warp_mouse] instead) and sending [kbd]Alt/Cmd + Tab[/kbd] as [InputEventKey] won't toggle between active windows.
*/
//go:nosplit
func (self class) ParseInputEvent(event objects.InputEvent) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(event[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_parse_input_event, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetUseAccumulatedInput(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_set_use_accumulated_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingAccumulatedInput() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_is_using_accumulated_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sends all input events which are in the current buffer to the game loop. These events may have been buffered as a result of accumulated input ([member use_accumulated_input]) or agile input flushing ([member ProjectSettings.input_devices/buffering/agile_event_flushing]).
The engine will already do this itself at key execution points (at least once per frame). However, this can be useful in advanced cases where you want precise control over the timing of event handling.
*/
//go:nosplit
func (self class) FlushBufferedEvents() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_flush_buffered_events, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetEmulateMouseFromTouch(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_set_emulate_mouse_from_touch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsEmulatingMouseFromTouch() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_is_emulating_mouse_from_touch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmulateTouchFromMouse(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_set_emulate_touch_from_mouse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsEmulatingTouchFromMouse() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Input.Bind_is_emulating_touch_from_mouse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func OnJoyConnectionChanged(cb func(device int, connected bool)) {
	self[0].AsObject().Connect(gd.NewStringName("joy_connection_changed"), gd.NewCallable(cb), 0)
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("Input", func(ptr gd.Object) any { return [1]classdb.Input{*(*classdb.Input)(unsafe.Pointer(&ptr))} })
}

type MouseModeValue = classdb.InputMouseMode

const (
	/*Makes the mouse cursor visible if it is hidden.*/
	MouseModeVisible MouseModeValue = 0
	/*Makes the mouse cursor hidden if it is visible.*/
	MouseModeHidden MouseModeValue = 1
	/*Captures the mouse. The mouse will be hidden and its position locked at the center of the window manager's window.
	  [b]Note:[/b] If you want to process the mouse's movement in this mode, you need to use [member InputEventMouseMotion.relative].*/
	MouseModeCaptured MouseModeValue = 2
	/*Confines the mouse cursor to the game window, and make it visible.*/
	MouseModeConfined MouseModeValue = 3
	/*Confines the mouse cursor to the game window, and make it hidden.*/
	MouseModeConfinedHidden MouseModeValue = 4
)

type CursorShape = classdb.InputCursorShape

const (
	/*Arrow cursor. Standard, default pointing cursor.*/
	CursorArrow CursorShape = 0
	/*I-beam cursor. Usually used to show where the text cursor will appear when the mouse is clicked.*/
	CursorIbeam CursorShape = 1
	/*Pointing hand cursor. Usually used to indicate the pointer is over a link or other interactable item.*/
	CursorPointingHand CursorShape = 2
	/*Cross cursor. Typically appears over regions in which a drawing operation can be performed or for selections.*/
	CursorCross CursorShape = 3
	/*Wait cursor. Indicates that the application is busy performing an operation, and that it cannot be used during the operation (e.g. something is blocking its main thread).*/
	CursorWait CursorShape = 4
	/*Busy cursor. Indicates that the application is busy performing an operation, and that it is still usable during the operation.*/
	CursorBusy CursorShape = 5
	/*Drag cursor. Usually displayed when dragging something.
	  [b]Note:[/b] Windows lacks a dragging cursor, so [constant CURSOR_DRAG] is the same as [constant CURSOR_MOVE] for this platform.*/
	CursorDrag CursorShape = 6
	/*Can drop cursor. Usually displayed when dragging something to indicate that it can be dropped at the current position.*/
	CursorCanDrop CursorShape = 7
	/*Forbidden cursor. Indicates that the current action is forbidden (for example, when dragging something) or that the control at a position is disabled.*/
	CursorForbidden CursorShape = 8
	/*Vertical resize mouse cursor. A double-headed vertical arrow. It tells the user they can resize the window or the panel vertically.*/
	CursorVsize CursorShape = 9
	/*Horizontal resize mouse cursor. A double-headed horizontal arrow. It tells the user they can resize the window or the panel horizontally.*/
	CursorHsize CursorShape = 10
	/*Window resize mouse cursor. The cursor is a double-headed arrow that goes from the bottom left to the top right. It tells the user they can resize the window or the panel both horizontally and vertically.*/
	CursorBdiagsize CursorShape = 11
	/*Window resize mouse cursor. The cursor is a double-headed arrow that goes from the top left to the bottom right, the opposite of [constant CURSOR_BDIAGSIZE]. It tells the user they can resize the window or the panel both horizontally and vertically.*/
	CursorFdiagsize CursorShape = 12
	/*Move cursor. Indicates that something can be moved.*/
	CursorMove CursorShape = 13
	/*Vertical split mouse cursor. On Windows, it's the same as [constant CURSOR_VSIZE].*/
	CursorVsplit CursorShape = 14
	/*Horizontal split mouse cursor. On Windows, it's the same as [constant CURSOR_HSIZE].*/
	CursorHsplit CursorShape = 15
	/*Help cursor. Usually a question mark.*/
	CursorHelp CursorShape = 16
)

type JoyAxis int

const (
	/*An invalid game controller axis.*/
	JoyAxisInvalid JoyAxis = -1
	/*Game controller left joystick x-axis.*/
	JoyAxisLeftX JoyAxis = 0
	/*Game controller left joystick y-axis.*/
	JoyAxisLeftY JoyAxis = 1
	/*Game controller right joystick x-axis.*/
	JoyAxisRightX JoyAxis = 2
	/*Game controller right joystick y-axis.*/
	JoyAxisRightY JoyAxis = 3
	/*Game controller left trigger axis.*/
	JoyAxisTriggerLeft JoyAxis = 4
	/*Game controller right trigger axis.*/
	JoyAxisTriggerRight JoyAxis = 5
	/*The number of SDL game controller axes.*/
	JoyAxisSdlMax JoyAxis = 6
	/*The maximum number of game controller axes: OpenVR supports up to 5 Joysticks making a total of 10 axes.*/
	JoyAxisMax JoyAxis = 10
)

type JoyButton int

const (
	/*An invalid game controller button.*/
	JoyButtonInvalid JoyButton = -1
	/*Game controller SDL button A. Corresponds to the bottom action button: Sony Cross, Xbox A, Nintendo B.*/
	JoyButtonA JoyButton = 0
	/*Game controller SDL button B. Corresponds to the right action button: Sony Circle, Xbox B, Nintendo A.*/
	JoyButtonB JoyButton = 1
	/*Game controller SDL button X. Corresponds to the left action button: Sony Square, Xbox X, Nintendo Y.*/
	JoyButtonX JoyButton = 2
	/*Game controller SDL button Y. Corresponds to the top action button: Sony Triangle, Xbox Y, Nintendo X.*/
	JoyButtonY JoyButton = 3
	/*Game controller SDL back button. Corresponds to the Sony Select, Xbox Back, Nintendo - button.*/
	JoyButtonBack JoyButton = 4
	/*Game controller SDL guide button. Corresponds to the Sony PS, Xbox Home button.*/
	JoyButtonGuide JoyButton = 5
	/*Game controller SDL start button. Corresponds to the Sony Options, Xbox Menu, Nintendo + button.*/
	JoyButtonStart JoyButton = 6
	/*Game controller SDL left stick button. Corresponds to the Sony L3, Xbox L/LS button.*/
	JoyButtonLeftStick JoyButton = 7
	/*Game controller SDL right stick button. Corresponds to the Sony R3, Xbox R/RS button.*/
	JoyButtonRightStick JoyButton = 8
	/*Game controller SDL left shoulder button. Corresponds to the Sony L1, Xbox LB button.*/
	JoyButtonLeftShoulder JoyButton = 9
	/*Game controller SDL right shoulder button. Corresponds to the Sony R1, Xbox RB button.*/
	JoyButtonRightShoulder JoyButton = 10
	/*Game controller D-pad up button.*/
	JoyButtonDpadUp JoyButton = 11
	/*Game controller D-pad down button.*/
	JoyButtonDpadDown JoyButton = 12
	/*Game controller D-pad left button.*/
	JoyButtonDpadLeft JoyButton = 13
	/*Game controller D-pad right button.*/
	JoyButtonDpadRight JoyButton = 14
	/*Game controller SDL miscellaneous button. Corresponds to Xbox share button, PS5 microphone button, Nintendo Switch capture button.*/
	JoyButtonMisc1 JoyButton = 15
	/*Game controller SDL paddle 1 button.*/
	JoyButtonPaddle1 JoyButton = 16
	/*Game controller SDL paddle 2 button.*/
	JoyButtonPaddle2 JoyButton = 17
	/*Game controller SDL paddle 3 button.*/
	JoyButtonPaddle3 JoyButton = 18
	/*Game controller SDL paddle 4 button.*/
	JoyButtonPaddle4 JoyButton = 19
	/*Game controller SDL touchpad button.*/
	JoyButtonTouchpad JoyButton = 20
	/*The number of SDL game controller buttons.*/
	JoyButtonSdlMax JoyButton = 21
	/*The maximum number of game controller buttons supported by the engine. The actual limit may be lower on specific platforms:
	  - [b]Android:[/b] Up to 36 buttons.
	  - [b]Linux:[/b] Up to 80 buttons.
	  - [b]Windows[/b] and [b]macOS:[/b] Up to 128 buttons.*/
	JoyButtonMax JoyButton = 128
)

type Key int

const (
	/*Enum value which doesn't correspond to any key. This is used to initialize [enum Key] properties with a generic state.*/
	KeyNone Key = 0
	/*Keycodes with this bit applied are non-printable.*/
	KeySpecial Key = 4194304
	/*Escape key.*/
	KeyEscape Key = 4194305
	/*Tab key.*/
	KeyTab Key = 4194306
	/*Shift + Tab key.*/
	KeyBacktab Key = 4194307
	/*Backspace key.*/
	KeyBackspace Key = 4194308
	/*Return key (on the main keyboard).*/
	KeyEnter Key = 4194309
	/*Enter key on the numeric keypad.*/
	KeyKpEnter Key = 4194310
	/*Insert key.*/
	KeyInsert Key = 4194311
	/*Delete key.*/
	KeyDelete Key = 4194312
	/*Pause key.*/
	KeyPause Key = 4194313
	/*Print Screen key.*/
	KeyPrint Key = 4194314
	/*System Request key.*/
	KeySysreq Key = 4194315
	/*Clear key.*/
	KeyClear Key = 4194316
	/*Home key.*/
	KeyHome Key = 4194317
	/*End key.*/
	KeyEnd Key = 4194318
	/*Left arrow key.*/
	KeyLeft Key = 4194319
	/*Up arrow key.*/
	KeyUp Key = 4194320
	/*Right arrow key.*/
	KeyRight Key = 4194321
	/*Down arrow key.*/
	KeyDown Key = 4194322
	/*Page Up key.*/
	KeyPageup Key = 4194323
	/*Page Down key.*/
	KeyPagedown Key = 4194324
	/*Shift key.*/
	KeyShift Key = 4194325
	/*Control key.*/
	KeyCtrl Key = 4194326
	/*Meta key.*/
	KeyMeta Key = 4194327
	/*Alt key.*/
	KeyAlt Key = 4194328
	/*Caps Lock key.*/
	KeyCapslock Key = 4194329
	/*Num Lock key.*/
	KeyNumlock Key = 4194330
	/*Scroll Lock key.*/
	KeyScrolllock Key = 4194331
	/*F1 key.*/
	KeyF1 Key = 4194332
	/*F2 key.*/
	KeyF2 Key = 4194333
	/*F3 key.*/
	KeyF3 Key = 4194334
	/*F4 key.*/
	KeyF4 Key = 4194335
	/*F5 key.*/
	KeyF5 Key = 4194336
	/*F6 key.*/
	KeyF6 Key = 4194337
	/*F7 key.*/
	KeyF7 Key = 4194338
	/*F8 key.*/
	KeyF8 Key = 4194339
	/*F9 key.*/
	KeyF9 Key = 4194340
	/*F10 key.*/
	KeyF10 Key = 4194341
	/*F11 key.*/
	KeyF11 Key = 4194342
	/*F12 key.*/
	KeyF12 Key = 4194343
	/*F13 key.*/
	KeyF13 Key = 4194344
	/*F14 key.*/
	KeyF14 Key = 4194345
	/*F15 key.*/
	KeyF15 Key = 4194346
	/*F16 key.*/
	KeyF16 Key = 4194347
	/*F17 key.*/
	KeyF17 Key = 4194348
	/*F18 key.*/
	KeyF18 Key = 4194349
	/*F19 key.*/
	KeyF19 Key = 4194350
	/*F20 key.*/
	KeyF20 Key = 4194351
	/*F21 key.*/
	KeyF21 Key = 4194352
	/*F22 key.*/
	KeyF22 Key = 4194353
	/*F23 key.*/
	KeyF23 Key = 4194354
	/*F24 key.*/
	KeyF24 Key = 4194355
	/*F25 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF25 Key = 4194356
	/*F26 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF26 Key = 4194357
	/*F27 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF27 Key = 4194358
	/*F28 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF28 Key = 4194359
	/*F29 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF29 Key = 4194360
	/*F30 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF30 Key = 4194361
	/*F31 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF31 Key = 4194362
	/*F32 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF32 Key = 4194363
	/*F33 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF33 Key = 4194364
	/*F34 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF34 Key = 4194365
	/*F35 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF35 Key = 4194366
	/*Multiply (*) key on the numeric keypad.*/
	KeyKpMultiply Key = 4194433
	/*Divide (/) key on the numeric keypad.*/
	KeyKpDivide Key = 4194434
	/*Subtract (-) key on the numeric keypad.*/
	KeyKpSubtract Key = 4194435
	/*Period (.) key on the numeric keypad.*/
	KeyKpPeriod Key = 4194436
	/*Add (+) key on the numeric keypad.*/
	KeyKpAdd Key = 4194437
	/*Number 0 on the numeric keypad.*/
	KeyKp0 Key = 4194438
	/*Number 1 on the numeric keypad.*/
	KeyKp1 Key = 4194439
	/*Number 2 on the numeric keypad.*/
	KeyKp2 Key = 4194440
	/*Number 3 on the numeric keypad.*/
	KeyKp3 Key = 4194441
	/*Number 4 on the numeric keypad.*/
	KeyKp4 Key = 4194442
	/*Number 5 on the numeric keypad.*/
	KeyKp5 Key = 4194443
	/*Number 6 on the numeric keypad.*/
	KeyKp6 Key = 4194444
	/*Number 7 on the numeric keypad.*/
	KeyKp7 Key = 4194445
	/*Number 8 on the numeric keypad.*/
	KeyKp8 Key = 4194446
	/*Number 9 on the numeric keypad.*/
	KeyKp9 Key = 4194447
	/*Context menu key.*/
	KeyMenu Key = 4194370
	/*Hyper key. (On Linux/X11 only).*/
	KeyHyper Key = 4194371
	/*Help key.*/
	KeyHelp Key = 4194373
	/*Media back key. Not to be confused with the Back button on an Android device.*/
	KeyBack Key = 4194376
	/*Media forward key.*/
	KeyForward Key = 4194377
	/*Media stop key.*/
	KeyStop Key = 4194378
	/*Media refresh key.*/
	KeyRefresh Key = 4194379
	/*Volume down key.*/
	KeyVolumedown Key = 4194380
	/*Mute volume key.*/
	KeyVolumemute Key = 4194381
	/*Volume up key.*/
	KeyVolumeup Key = 4194382
	/*Media play key.*/
	KeyMediaplay Key = 4194388
	/*Media stop key.*/
	KeyMediastop Key = 4194389
	/*Previous song key.*/
	KeyMediaprevious Key = 4194390
	/*Next song key.*/
	KeyMedianext Key = 4194391
	/*Media record key.*/
	KeyMediarecord Key = 4194392
	/*Home page key.*/
	KeyHomepage Key = 4194393
	/*Favorites key.*/
	KeyFavorites Key = 4194394
	/*Search key.*/
	KeySearch Key = 4194395
	/*Standby key.*/
	KeyStandby Key = 4194396
	/*Open URL / Launch Browser key.*/
	KeyOpenurl Key = 4194397
	/*Launch Mail key.*/
	KeyLaunchmail Key = 4194398
	/*Launch Media key.*/
	KeyLaunchmedia Key = 4194399
	/*Launch Shortcut 0 key.*/
	KeyLaunch0 Key = 4194400
	/*Launch Shortcut 1 key.*/
	KeyLaunch1 Key = 4194401
	/*Launch Shortcut 2 key.*/
	KeyLaunch2 Key = 4194402
	/*Launch Shortcut 3 key.*/
	KeyLaunch3 Key = 4194403
	/*Launch Shortcut 4 key.*/
	KeyLaunch4 Key = 4194404
	/*Launch Shortcut 5 key.*/
	KeyLaunch5 Key = 4194405
	/*Launch Shortcut 6 key.*/
	KeyLaunch6 Key = 4194406
	/*Launch Shortcut 7 key.*/
	KeyLaunch7 Key = 4194407
	/*Launch Shortcut 8 key.*/
	KeyLaunch8 Key = 4194408
	/*Launch Shortcut 9 key.*/
	KeyLaunch9 Key = 4194409
	/*Launch Shortcut A key.*/
	KeyLauncha Key = 4194410
	/*Launch Shortcut B key.*/
	KeyLaunchb Key = 4194411
	/*Launch Shortcut C key.*/
	KeyLaunchc Key = 4194412
	/*Launch Shortcut D key.*/
	KeyLaunchd Key = 4194413
	/*Launch Shortcut E key.*/
	KeyLaunche Key = 4194414
	/*Launch Shortcut F key.*/
	KeyLaunchf Key = 4194415
	/*"Globe" key on Mac / iPad keyboard.*/
	KeyGlobe Key = 4194416
	/*"On-screen keyboard" key on iPad keyboard.*/
	KeyKeyboard Key = 4194417
	/*英数 key on Mac keyboard.*/
	KeyJisEisu Key = 4194418
	/*かな key on Mac keyboard.*/
	KeyJisKana Key = 4194419
	/*Unknown key.*/
	KeyUnknown Key = 8388607
	/*Space key.*/
	KeySpace Key = 32
	/*! key.*/
	KeyExclam Key = 33
	/*" key.*/
	KeyQuotedbl Key = 34
	/*# key.*/
	KeyNumbersign Key = 35
	/*$ key.*/
	KeyDollar Key = 36
	/*% key.*/
	KeyPercent Key = 37
	/*& key.*/
	KeyAmpersand Key = 38
	/*' key.*/
	KeyApostrophe Key = 39
	/*( key.*/
	KeyParenleft Key = 40
	/*) key.*/
	KeyParenright Key = 41
	/** key.*/
	KeyAsterisk Key = 42
	/*+ key.*/
	KeyPlus Key = 43
	/*, key.*/
	KeyComma Key = 44
	/*- key.*/
	KeyMinus Key = 45
	/*. key.*/
	KeyPeriod Key = 46
	/*/ key.*/
	KeySlash Key = 47
	/*Number 0 key.*/
	Key0 Key = 48
	/*Number 1 key.*/
	Key1 Key = 49
	/*Number 2 key.*/
	Key2 Key = 50
	/*Number 3 key.*/
	Key3 Key = 51
	/*Number 4 key.*/
	Key4 Key = 52
	/*Number 5 key.*/
	Key5 Key = 53
	/*Number 6 key.*/
	Key6 Key = 54
	/*Number 7 key.*/
	Key7 Key = 55
	/*Number 8 key.*/
	Key8 Key = 56
	/*Number 9 key.*/
	Key9 Key = 57
	/*: key.*/
	KeyColon Key = 58
	/*; key.*/
	KeySemicolon Key = 59
	/*< key.*/
	KeyLess Key = 60
	/*= key.*/
	KeyEqual Key = 61
	/*> key.*/
	KeyGreater Key = 62
	/*? key.*/
	KeyQuestion Key = 63
	/*@ key.*/
	KeyAt Key = 64
	/*A key.*/
	KeyA Key = 65
	/*B key.*/
	KeyB Key = 66
	/*C key.*/
	KeyC Key = 67
	/*D key.*/
	KeyD Key = 68
	/*E key.*/
	KeyE Key = 69
	/*F key.*/
	KeyF Key = 70
	/*G key.*/
	KeyG Key = 71
	/*H key.*/
	KeyH Key = 72
	/*I key.*/
	KeyI Key = 73
	/*J key.*/
	KeyJ Key = 74
	/*K key.*/
	KeyK Key = 75
	/*L key.*/
	KeyL Key = 76
	/*M key.*/
	KeyM Key = 77
	/*N key.*/
	KeyN Key = 78
	/*O key.*/
	KeyO Key = 79
	/*P key.*/
	KeyP Key = 80
	/*Q key.*/
	KeyQ Key = 81
	/*R key.*/
	KeyR Key = 82
	/*S key.*/
	KeyS Key = 83
	/*T key.*/
	KeyT Key = 84
	/*U key.*/
	KeyU Key = 85
	/*V key.*/
	KeyV Key = 86
	/*W key.*/
	KeyW Key = 87
	/*X key.*/
	KeyX Key = 88
	/*Y key.*/
	KeyY Key = 89
	/*Z key.*/
	KeyZ Key = 90
	/*[ key.*/
	KeyBracketleft Key = 91
	/*\ key.*/
	KeyBackslash Key = 92
	/*] key.*/
	KeyBracketright Key = 93
	/*^ key.*/
	KeyAsciicircum Key = 94
	/*_ key.*/
	KeyUnderscore Key = 95
	/*` key.*/
	KeyQuoteleft Key = 96
	/*{ key.*/
	KeyBraceleft Key = 123
	/*| key.*/
	KeyBar Key = 124
	/*} key.*/
	KeyBraceright Key = 125
	/*~ key.*/
	KeyAsciitilde Key = 126
	/*¥ key.*/
	KeyYen Key = 165
	/*§ key.*/
	KeySection Key = 167
)

type MouseButton int

const (
	/*Enum value which doesn't correspond to any mouse button. This is used to initialize [enum MouseButton] properties with a generic state.*/
	MouseButtonNone MouseButton = 0
	/*Primary mouse button, usually assigned to the left button.*/
	MouseButtonLeft MouseButton = 1
	/*Secondary mouse button, usually assigned to the right button.*/
	MouseButtonRight MouseButton = 2
	/*Middle mouse button.*/
	MouseButtonMiddle MouseButton = 3
	/*Mouse wheel scrolling up.*/
	MouseButtonWheelUp MouseButton = 4
	/*Mouse wheel scrolling down.*/
	MouseButtonWheelDown MouseButton = 5
	/*Mouse wheel left button (only present on some mice).*/
	MouseButtonWheelLeft MouseButton = 6
	/*Mouse wheel right button (only present on some mice).*/
	MouseButtonWheelRight MouseButton = 7
	/*Extra mouse button 1. This is sometimes present, usually to the sides of the mouse.*/
	MouseButtonXbutton1 MouseButton = 8
	/*Extra mouse button 2. This is sometimes present, usually to the sides of the mouse.*/
	MouseButtonXbutton2 MouseButton = 9
)

type MouseButtonMask int

const (
	/*Primary mouse button mask, usually for the left button.*/
	MouseButtonMaskLeft MouseButtonMask = 1
	/*Secondary mouse button mask, usually for the right button.*/
	MouseButtonMaskRight MouseButtonMask = 2
	/*Middle mouse button mask.*/
	MouseButtonMaskMiddle MouseButtonMask = 4
	/*Extra mouse button 1 mask.*/
	MouseButtonMaskMbXbutton1 MouseButtonMask = 128
	/*Extra mouse button 2 mask.*/
	MouseButtonMaskMbXbutton2 MouseButtonMask = 256
)
