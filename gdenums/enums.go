package gdenums

type Side int

const (
/*Left side, usually used for [Control] or [StyleBox]-derived classes.*/
	SideLeft Side = 0
/*Top side, usually used for [Control] or [StyleBox]-derived classes.*/
	SideTop Side = 1
/*Right side, usually used for [Control] or [StyleBox]-derived classes.*/
	SideRight Side = 2
/*Bottom side, usually used for [Control] or [StyleBox]-derived classes.*/
	SideBottom Side = 3
)

type Corner int

const (
/*Top-left corner.*/
	CornerTopLeft Corner = 0
/*Top-right corner.*/
	CornerTopRight Corner = 1
/*Bottom-right corner.*/
	CornerBottomRight Corner = 2
/*Bottom-left corner.*/
	CornerBottomLeft Corner = 3
)

type Orientation int

const (
/*General vertical alignment, usually used for [Separator], [ScrollBar], [Slider], etc.*/
	Vertical Orientation = 1
/*General horizontal alignment, usually used for [Separator], [ScrollBar], [Slider], etc.*/
	Horizontal Orientation = 0
)

type ClockDirection int

const (
/*Clockwise rotation. Used by some methods (e.g. [method Image.rotate_90]).*/
	Clockwise ClockDirection = 0
/*Counter-clockwise rotation. Used by some methods (e.g. [method Image.rotate_90]).*/
	Counterclockwise ClockDirection = 1
)

type HorizontalAlignment int

const (
/*Horizontal left alignment, usually for text-derived classes.*/
	HorizontalAlignmentLeft HorizontalAlignment = 0
/*Horizontal center alignment, usually for text-derived classes.*/
	HorizontalAlignmentCenter HorizontalAlignment = 1
/*Horizontal right alignment, usually for text-derived classes.*/
	HorizontalAlignmentRight HorizontalAlignment = 2
/*Expand row to fit width, usually for text-derived classes.*/
	HorizontalAlignmentFill HorizontalAlignment = 3
)

type VerticalAlignment int

const (
/*Vertical top alignment, usually for text-derived classes.*/
	VerticalAlignmentTop VerticalAlignment = 0
/*Vertical center alignment, usually for text-derived classes.*/
	VerticalAlignmentCenter VerticalAlignment = 1
/*Vertical bottom alignment, usually for text-derived classes.*/
	VerticalAlignmentBottom VerticalAlignment = 2
/*Expand rows to fit height, usually for text-derived classes.*/
	VerticalAlignmentFill VerticalAlignment = 3
)

type InlineAlignment int

const (
/*Aligns the top of the inline object (e.g. image, table) to the position of the text specified by [code]INLINE_ALIGNMENT_TO_*[/code] constant.*/
	InlineAlignmentTopTo InlineAlignment = 0
/*Aligns the center of the inline object (e.g. image, table) to the position of the text specified by [code]INLINE_ALIGNMENT_TO_*[/code] constant.*/
	InlineAlignmentCenterTo InlineAlignment = 1
/*Aligns the baseline (user defined) of the inline object (e.g. image, table) to the position of the text specified by [code]INLINE_ALIGNMENT_TO_*[/code] constant.*/
	InlineAlignmentBaselineTo InlineAlignment = 3
/*Aligns the bottom of the inline object (e.g. image, table) to the position of the text specified by [code]INLINE_ALIGNMENT_TO_*[/code] constant.*/
	InlineAlignmentBottomTo InlineAlignment = 2
/*Aligns the position of the inline object (e.g. image, table) specified by [code]INLINE_ALIGNMENT_*_TO[/code] constant to the top of the text.*/
	InlineAlignmentToTop InlineAlignment = 0
/*Aligns the position of the inline object (e.g. image, table) specified by [code]INLINE_ALIGNMENT_*_TO[/code] constant to the center of the text.*/
	InlineAlignmentToCenter InlineAlignment = 4
/*Aligns the position of the inline object (e.g. image, table) specified by [code]INLINE_ALIGNMENT_*_TO[/code] constant to the baseline of the text.*/
	InlineAlignmentToBaseline InlineAlignment = 8
/*Aligns inline object (e.g. image, table) to the bottom of the text.*/
	InlineAlignmentToBottom InlineAlignment = 12
/*Aligns top of the inline object (e.g. image, table) to the top of the text. Equivalent to [code]INLINE_ALIGNMENT_TOP_TO | INLINE_ALIGNMENT_TO_TOP[/code].*/
	InlineAlignmentTop InlineAlignment = 0
/*Aligns center of the inline object (e.g. image, table) to the center of the text. Equivalent to [code]INLINE_ALIGNMENT_CENTER_TO | INLINE_ALIGNMENT_TO_CENTER[/code].*/
	InlineAlignmentCenter InlineAlignment = 5
/*Aligns bottom of the inline object (e.g. image, table) to the bottom of the text. Equivalent to [code]INLINE_ALIGNMENT_BOTTOM_TO | INLINE_ALIGNMENT_TO_BOTTOM[/code].*/
	InlineAlignmentBottom InlineAlignment = 14
/*A bit mask for [code]INLINE_ALIGNMENT_*_TO[/code] alignment constants.*/
	InlineAlignmentImageMask InlineAlignment = 3
/*A bit mask for [code]INLINE_ALIGNMENT_TO_*[/code] alignment constants.*/
	InlineAlignmentTextMask InlineAlignment = 12
)

type EulerOrder int

const (
/*Specifies that Euler angles should be in XYZ order. When composing, the order is X, Y, Z. When decomposing, the order is reversed, first Z, then Y, and X last.*/
	EulerOrderXyz EulerOrder = 0
/*Specifies that Euler angles should be in XZY order. When composing, the order is X, Z, Y. When decomposing, the order is reversed, first Y, then Z, and X last.*/
	EulerOrderXzy EulerOrder = 1
/*Specifies that Euler angles should be in YXZ order. When composing, the order is Y, X, Z. When decomposing, the order is reversed, first Z, then X, and Y last.*/
	EulerOrderYxz EulerOrder = 2
/*Specifies that Euler angles should be in YZX order. When composing, the order is Y, Z, X. When decomposing, the order is reversed, first X, then Z, and Y last.*/
	EulerOrderYzx EulerOrder = 3
/*Specifies that Euler angles should be in ZXY order. When composing, the order is Z, X, Y. When decomposing, the order is reversed, first Y, then X, and Z last.*/
	EulerOrderZxy EulerOrder = 4
/*Specifies that Euler angles should be in ZYX order. When composing, the order is Z, Y, X. When decomposing, the order is reversed, first X, then Y, and Z last.*/
	EulerOrderZyx EulerOrder = 5
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

type KeyModifierMask int

const (
/*Key Code mask.*/
	KeyCodeMask KeyModifierMask = 8388607
/*Modifier key mask.*/
	KeyModifierMaskDefault KeyModifierMask = 532676608
/*Automatically remapped to [constant KEY_META] on macOS and [constant KEY_CTRL] on other platforms, this mask is never set in the actual events, and should be used for key mapping only.*/
	KeyMaskCmdOrCtrl KeyModifierMask = 16777216
/*Shift key mask.*/
	KeyMaskShift KeyModifierMask = 33554432
/*Alt or Option (on macOS) key mask.*/
	KeyMaskAlt KeyModifierMask = 67108864
/*Command (on macOS) or Meta/Windows key mask.*/
	KeyMaskMeta KeyModifierMask = 134217728
/*Control key mask.*/
	KeyMaskCtrl KeyModifierMask = 268435456
/*Keypad key mask.*/
	KeyMaskKpad KeyModifierMask = 536870912
/*Group Switch key mask.*/
	KeyMaskGroupSwitch KeyModifierMask = 1073741824
)

type KeyLocation int

const (
/*Used for keys which only appear once, or when a comparison doesn't need to differentiate the [code]LEFT[/code] and [code]RIGHT[/code] versions.
For example, when using [method InputEvent.is_match], an event which has [constant KEY_LOCATION_UNSPECIFIED] will match any [enum KeyLocation] on the passed event.*/
	KeyLocationUnspecified KeyLocation = 0
/*A key which is to the left of its twin.*/
	KeyLocationLeft KeyLocation = 1
/*A key which is to the right of its twin.*/
	KeyLocationRight KeyLocation = 2
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

type MIDIMessage int

const (
/*Does not correspond to any MIDI message. This is the default value of [member InputEventMIDI.message].*/
	MidiMessageNone MIDIMessage = 0
/*MIDI message sent when a note is released.
[b]Note:[/b] Not all MIDI devices send this message; some may send [constant MIDI_MESSAGE_NOTE_ON] with [member InputEventMIDI.velocity] set to [code]0[/code].*/
	MidiMessageNoteOff MIDIMessage = 8
/*MIDI message sent when a note is pressed.*/
	MidiMessageNoteOn MIDIMessage = 9
/*MIDI message sent to indicate a change in pressure while a note is being pressed down, also called aftertouch.*/
	MidiMessageAftertouch MIDIMessage = 10
/*MIDI message sent when a controller value changes. In a MIDI device, a controller is any input that doesn't play notes. These may include sliders for volume, balance, and panning, as well as switches and pedals. See the [url=https://en.wikipedia.org/wiki/General_MIDI#Controller_events]General MIDI specification[/url] for a small list.*/
	MidiMessageControlChange MIDIMessage = 11
/*MIDI message sent when the MIDI device changes its current instrument (also called [i]program[/i] or [i]preset[/i]).*/
	MidiMessageProgramChange MIDIMessage = 12
/*MIDI message sent to indicate a change in pressure for the whole channel. Some MIDI devices may send this instead of [constant MIDI_MESSAGE_AFTERTOUCH].*/
	MidiMessageChannelPressure MIDIMessage = 13
/*MIDI message sent when the value of the pitch bender changes, usually a wheel on the MIDI device.*/
	MidiMessagePitchBend MIDIMessage = 14
/*MIDI system exclusive (SysEx) message. This type of message is not standardized and it's highly dependent on the MIDI device sending it.
[b]Note:[/b] Getting this message's data from [InputEventMIDI] is not implemented.*/
	MidiMessageSystemExclusive MIDIMessage = 240
/*MIDI message sent every quarter frame to keep connected MIDI devices synchronized. Related to [constant MIDI_MESSAGE_TIMING_CLOCK].
[b]Note:[/b] Getting this message's data from [InputEventMIDI] is not implemented.*/
	MidiMessageQuarterFrame MIDIMessage = 241
/*MIDI message sent to jump onto a new position in the current sequence or song.
[b]Note:[/b] Getting this message's data from [InputEventMIDI] is not implemented.*/
	MidiMessageSongPositionPointer MIDIMessage = 242
/*MIDI message sent to select a sequence or song to play.
[b]Note:[/b] Getting this message's data from [InputEventMIDI] is not implemented.*/
	MidiMessageSongSelect MIDIMessage = 243
/*MIDI message sent to request a tuning calibration. Used on analog synthesizers. Most modern MIDI devices do not need this message.*/
	MidiMessageTuneRequest MIDIMessage = 246
/*MIDI message sent 24 times after [constant MIDI_MESSAGE_QUARTER_FRAME], to keep connected MIDI devices synchronized.*/
	MidiMessageTimingClock MIDIMessage = 248
/*MIDI message sent to start the current sequence or song from the beginning.*/
	MidiMessageStart MIDIMessage = 250
/*MIDI message sent to resume from the point the current sequence or song was paused.*/
	MidiMessageContinue MIDIMessage = 251
/*MIDI message sent to pause the current sequence or song.*/
	MidiMessageStop MIDIMessage = 252
/*MIDI message sent repeatedly while the MIDI device is idle, to tell the receiver that the connection is alive. Most MIDI devices do not send this message.*/
	MidiMessageActiveSensing MIDIMessage = 254
/*MIDI message sent to reset a MIDI device to its default state, as if it was just turned on. It should not be sent when the MIDI device is being turned on.*/
	MidiMessageSystemReset MIDIMessage = 255
)

type Error int

const (
/*Methods that return [enum Error] return [constant OK] when no error occurred.
Since [constant OK] has value 0, and all other error constants are positive integers, it can also be used in boolean checks.
[b]Example:[/b]
[codeblock]
var error = method_that_returns_error()
if error != OK:
    printerr("Failure!")

# Or, alternatively:
if error:
    printerr("Still failing!")
[/codeblock]
[b]Note:[/b] Many functions do not return an error code, but will print error messages to standard output.*/
	Ok Error = 0
/*Generic error.*/
	Failed Error = 1
/*Unavailable error.*/
	ErrUnavailable Error = 2
/*Unconfigured error.*/
	ErrUnconfigured Error = 3
/*Unauthorized error.*/
	ErrUnauthorized Error = 4
/*Parameter range error.*/
	ErrParameterRangeError Error = 5
/*Out of memory (OOM) error.*/
	ErrOutOfMemory Error = 6
/*File: Not found error.*/
	ErrFileNotFound Error = 7
/*File: Bad drive error.*/
	ErrFileBadDrive Error = 8
/*File: Bad path error.*/
	ErrFileBadPath Error = 9
/*File: No permission error.*/
	ErrFileNoPermission Error = 10
/*File: Already in use error.*/
	ErrFileAlreadyInUse Error = 11
/*File: Can't open error.*/
	ErrFileCantOpen Error = 12
/*File: Can't write error.*/
	ErrFileCantWrite Error = 13
/*File: Can't read error.*/
	ErrFileCantRead Error = 14
/*File: Unrecognized error.*/
	ErrFileUnrecognized Error = 15
/*File: Corrupt error.*/
	ErrFileCorrupt Error = 16
/*File: Missing dependencies error.*/
	ErrFileMissingDependencies Error = 17
/*File: End of file (EOF) error.*/
	ErrFileEof Error = 18
/*Can't open error.*/
	ErrCantOpen Error = 19
/*Can't create error.*/
	ErrCantCreate Error = 20
/*Query failed error.*/
	ErrQueryFailed Error = 21
/*Already in use error.*/
	ErrAlreadyInUse Error = 22
/*Locked error.*/
	ErrLocked Error = 23
/*Timeout error.*/
	ErrTimeout Error = 24
/*Can't connect error.*/
	ErrCantConnect Error = 25
/*Can't resolve error.*/
	ErrCantResolve Error = 26
/*Connection error.*/
	ErrConnectionError Error = 27
/*Can't acquire resource error.*/
	ErrCantAcquireResource Error = 28
/*Can't fork process error.*/
	ErrCantFork Error = 29
/*Invalid data error.*/
	ErrInvalidData Error = 30
/*Invalid parameter error.*/
	ErrInvalidParameter Error = 31
/*Already exists error.*/
	ErrAlreadyExists Error = 32
/*Does not exist error.*/
	ErrDoesNotExist Error = 33
/*Database: Read error.*/
	ErrDatabaseCantRead Error = 34
/*Database: Write error.*/
	ErrDatabaseCantWrite Error = 35
/*Compilation failed error.*/
	ErrCompilationFailed Error = 36
/*Method not found error.*/
	ErrMethodNotFound Error = 37
/*Linking failed error.*/
	ErrLinkFailed Error = 38
/*Script failed error.*/
	ErrScriptFailed Error = 39
/*Cycling link (import cycle) error.*/
	ErrCyclicLink Error = 40
/*Invalid declaration error.*/
	ErrInvalidDeclaration Error = 41
/*Duplicate symbol error.*/
	ErrDuplicateSymbol Error = 42
/*Parse error.*/
	ErrParseError Error = 43
/*Busy error.*/
	ErrBusy Error = 44
/*Skip error.*/
	ErrSkip Error = 45
/*Help error. Used internally when passing [code]--version[/code] or [code]--help[/code] as executable options.*/
	ErrHelp Error = 46
/*Bug error, caused by an implementation issue in the method.
[b]Note:[/b] If a built-in method returns this code, please open an issue on [url=https://github.com/godotengine/godot/issues]the GitHub Issue Tracker[/url].*/
	ErrBug Error = 47
/*Printer on fire error (This is an easter egg, no built-in methods return this error code).*/
	ErrPrinterOnFire Error = 48
)

type PropertyHint int

const (
/*The property has no hint for the editor.*/
	PropertyHintNone PropertyHint = 0
/*Hints that an [int] or [float] property should be within a range specified via the hint string [code]"min,max"[/code] or [code]"min,max,step"[/code]. The hint string can optionally include [code]"or_greater"[/code] and/or [code]"or_less"[/code] to allow manual input going respectively above the max or below the min values.
[b]Example:[/b] [code]"-360,360,1,or_greater,or_less"[/code].
Additionally, other keywords can be included: [code]"exp"[/code] for exponential range editing, [code]"radians_as_degrees"[/code] for editing radian angles in degrees (the range values are also in degrees), [code]"degrees"[/code] to hint at an angle and [code]"hide_slider"[/code] to hide the slider.*/
	PropertyHintRange PropertyHint = 1
/*Hints that an [int] or [String] property is an enumerated value to pick in a list specified via a hint string.
The hint string is a comma separated list of names such as [code]"Hello,Something,Else"[/code]. Whitespaces are [b]not[/b] removed from either end of a name. For integer properties, the first name in the list has value 0, the next 1, and so on. Explicit values can also be specified by appending [code]:integer[/code] to the name, e.g. [code]"Zero,One,Three:3,Four,Six:6"[/code].*/
	PropertyHintEnum PropertyHint = 2
/*Hints that a [String] property can be an enumerated value to pick in a list specified via a hint string such as [code]"Hello,Something,Else"[/code].
Unlike [constant PROPERTY_HINT_ENUM], a property with this hint still accepts arbitrary values and can be empty. The list of values serves to suggest possible values.*/
	PropertyHintEnumSuggestion PropertyHint = 3
/*Hints that a [float] property should be edited via an exponential easing function. The hint string can include [code]"attenuation"[/code] to flip the curve horizontally and/or [code]"positive_only"[/code] to exclude in/out easing and limit values to be greater than or equal to zero.*/
	PropertyHintExpEasing PropertyHint = 4
/*Hints that a vector property should allow its components to be linked. For example, this allows [member Vector2.x] and [member Vector2.y] to be edited together.*/
	PropertyHintLink PropertyHint = 5
/*Hints that an [int] property is a bitmask with named bit flags.
The hint string is a comma separated list of names such as [code]"Bit0,Bit1,Bit2,Bit3"[/code]. Whitespaces are [b]not[/b] removed from either end of a name. The first name in the list has value 1, the next 2, then 4, 8, 16 and so on. Explicit values can also be specified by appending [code]:integer[/code] to the name, e.g. [code]"A:4,B:8,C:16"[/code]. You can also combine several flags ([code]"A:4,B:8,AB:12,C:16"[/code]).
[b]Note:[/b] A flag value must be at least [code]1[/code] and at most [code]2 ** 32 - 1[/code].
[b]Note:[/b] Unlike [constant PROPERTY_HINT_ENUM], the previous explicit value is not taken into account. For the hint [code]"A:16,B,C"[/code], A is 16, B is 2, C is 4.*/
	PropertyHintFlags PropertyHint = 6
/*Hints that an [int] property is a bitmask using the optionally named 2D render layers.*/
	PropertyHintLayers2dRender PropertyHint = 7
/*Hints that an [int] property is a bitmask using the optionally named 2D physics layers.*/
	PropertyHintLayers2dPhysics PropertyHint = 8
/*Hints that an [int] property is a bitmask using the optionally named 2D navigation layers.*/
	PropertyHintLayers2dNavigation PropertyHint = 9
/*Hints that an [int] property is a bitmask using the optionally named 3D render layers.*/
	PropertyHintLayers3dRender PropertyHint = 10
/*Hints that an [int] property is a bitmask using the optionally named 3D physics layers.*/
	PropertyHintLayers3dPhysics PropertyHint = 11
/*Hints that an [int] property is a bitmask using the optionally named 3D navigation layers.*/
	PropertyHintLayers3dNavigation PropertyHint = 12
/*Hints that an integer property is a bitmask using the optionally named avoidance layers.*/
	PropertyHintLayersAvoidance PropertyHint = 37
/*Hints that a [String] property is a path to a file. Editing it will show a file dialog for picking the path. The hint string can be a set of filters with wildcards like [code]"*.png,*.jpg"[/code].*/
	PropertyHintFile PropertyHint = 13
/*Hints that a [String] property is a path to a directory. Editing it will show a file dialog for picking the path.*/
	PropertyHintDir PropertyHint = 14
/*Hints that a [String] property is an absolute path to a file outside the project folder. Editing it will show a file dialog for picking the path. The hint string can be a set of filters with wildcards, like [code]"*.png,*.jpg"[/code].*/
	PropertyHintGlobalFile PropertyHint = 15
/*Hints that a [String] property is an absolute path to a directory outside the project folder. Editing it will show a file dialog for picking the path.*/
	PropertyHintGlobalDir PropertyHint = 16
/*Hints that a property is an instance of a [Resource]-derived type, optionally specified via the hint string (e.g. [code]"Texture2D"[/code]). Editing it will show a popup menu of valid resource types to instantiate.*/
	PropertyHintResourceType PropertyHint = 17
/*Hints that a [String] property is text with line breaks. Editing it will show a text input field where line breaks can be typed.*/
	PropertyHintMultilineText PropertyHint = 18
/*Hints that a [String] property is an [Expression].*/
	PropertyHintExpression PropertyHint = 19
/*Hints that a [String] property should show a placeholder text on its input field, if empty. The hint string is the placeholder text to use.*/
	PropertyHintPlaceholderText PropertyHint = 20
/*Hints that a [Color] property should be edited without affecting its transparency ([member Color.a] is not editable).*/
	PropertyHintColorNoAlpha PropertyHint = 21
/*Hints that the property's value is an object encoded as object ID, with its type specified in the hint string. Used by the debugger.*/
	PropertyHintObjectId PropertyHint = 22
/*If a property is [String], hints that the property represents a particular type (class). This allows to select a type from the create dialog. The property will store the selected type as a string.
If a property is [Array], hints the editor how to show elements. The [code]hint_string[/code] must encode nested types using [code]":"[/code] and [code]"/"[/code].
[codeblocks]
[gdscript]
# Array of elem_type.
hint_string = "%d:" % [elem_type]
hint_string = "%d/%d:%s" % [elem_type, elem_hint, elem_hint_string]
# Two-dimensional array of elem_type (array of arrays of elem_type).
hint_string = "%d:%d:" % [TYPE_ARRAY, elem_type]
hint_string = "%d:%d/%d:%s" % [TYPE_ARRAY, elem_type, elem_hint, elem_hint_string]
# Three-dimensional array of elem_type (array of arrays of arrays of elem_type).
hint_string = "%d:%d:%d:" % [TYPE_ARRAY, TYPE_ARRAY, elem_type]
hint_string = "%d:%d:%d/%d:%s" % [TYPE_ARRAY, TYPE_ARRAY, elem_type, elem_hint, elem_hint_string]
[/gdscript]
[csharp]
// Array of elemType.
hintString = $"{elemType:D}:";
hintString = $"{elemType:}/{elemHint:D}:{elemHintString}";
// Two-dimensional array of elemType (array of arrays of elemType).
hintString = $"{Variant.Type.Array:D}:{elemType:D}:";
hintString = $"{Variant.Type.Array:D}:{elemType:D}/{elemHint:D}:{elemHintString}";
// Three-dimensional array of elemType (array of arrays of arrays of elemType).
hintString = $"{Variant.Type.Array:D}:{Variant.Type.Array:D}:{elemType:D}:";
hintString = $"{Variant.Type.Array:D}:{Variant.Type.Array:D}:{elemType:D}/{elemHint:D}:{elemHintString}";
[/csharp]
[/codeblocks]
Examples:
[codeblocks]
[gdscript]
hint_string = "%d:" % [TYPE_INT] # Array of integers.
hint_string = "%d/%d:1,10,1" % [TYPE_INT, PROPERTY_HINT_RANGE] # Array of integers (in range from 1 to 10).
hint_string = "%d/%d:Zero,One,Two" % [TYPE_INT, PROPERTY_HINT_ENUM] # Array of integers (an enum).
hint_string = "%d/%d:Zero,One,Three:3,Six:6" % [TYPE_INT, PROPERTY_HINT_ENUM] # Array of integers (an enum).
hint_string = "%d/%d:*.png" % [TYPE_STRING, PROPERTY_HINT_FILE] # Array of strings (file paths).
hint_string = "%d/%d:Texture2D" % [TYPE_OBJECT, PROPERTY_HINT_RESOURCE_TYPE] # Array of textures.

hint_string = "%d:%d:" % [TYPE_ARRAY, TYPE_FLOAT] # Two-dimensional array of floats.
hint_string = "%d:%d/%d:" % [TYPE_ARRAY, TYPE_STRING, PROPERTY_HINT_MULTILINE_TEXT] # Two-dimensional array of multiline strings.
hint_string = "%d:%d/%d:-1,1,0.1" % [TYPE_ARRAY, TYPE_FLOAT, PROPERTY_HINT_RANGE] # Two-dimensional array of floats (in range from -1 to 1).
hint_string = "%d:%d/%d:Texture2D" % [TYPE_ARRAY, TYPE_OBJECT, PROPERTY_HINT_RESOURCE_TYPE] # Two-dimensional array of textures.
[/gdscript]
[csharp]
hintString = $"{Variant.Type.Int:D}/{PropertyHint.Range:D}:1,10,1"; // Array of integers (in range from 1 to 10).
hintString = $"{Variant.Type.Int:D}/{PropertyHint.Enum:D}:Zero,One,Two"; // Array of integers (an enum).
hintString = $"{Variant.Type.Int:D}/{PropertyHint.Enum:D}:Zero,One,Three:3,Six:6"; // Array of integers (an enum).
hintString = $"{Variant.Type.String:D}/{PropertyHint.File:D}:*.png"; // Array of strings (file paths).
hintString = $"{Variant.Type.Object:D}/{PropertyHint.ResourceType:D}:Texture2D"; // Array of textures.

hintString = $"{Variant.Type.Array:D}:{Variant.Type.Float:D}:"; // Two-dimensional array of floats.
hintString = $"{Variant.Type.Array:D}:{Variant.Type.String:D}/{PropertyHint.MultilineText:D}:"; // Two-dimensional array of multiline strings.
hintString = $"{Variant.Type.Array:D}:{Variant.Type.Float:D}/{PropertyHint.Range:D}:-1,1,0.1"; // Two-dimensional array of floats (in range from -1 to 1).
hintString = $"{Variant.Type.Array:D}:{Variant.Type.Object:D}/{PropertyHint.ResourceType:D}:Texture2D"; // Two-dimensional array of textures.
[/csharp]
[/codeblocks]
[b]Note:[/b] The trailing colon is required for properly detecting built-in types.*/
	PropertyHintTypeString PropertyHint = 23
	PropertyHintNodePathToEditedNode PropertyHint = 24
/*Hints that an object is too big to be sent via the debugger.*/
	PropertyHintObjectTooBig PropertyHint = 25
/*Hints that the hint string specifies valid node types for property of type [NodePath].*/
	PropertyHintNodePathValidTypes PropertyHint = 26
/*Hints that a [String] property is a path to a file. Editing it will show a file dialog for picking the path for the file to be saved at. The dialog has access to the project's directory. The hint string can be a set of filters with wildcards like [code]"*.png,*.jpg"[/code]. See also [member FileDialog.filters].*/
	PropertyHintSaveFile PropertyHint = 27
/*Hints that a [String] property is a path to a file. Editing it will show a file dialog for picking the path for the file to be saved at. The dialog has access to the entire filesystem. The hint string can be a set of filters with wildcards like [code]"*.png,*.jpg"[/code]. See also [member FileDialog.filters].*/
	PropertyHintGlobalSaveFile PropertyHint = 28
	PropertyHintIntIsObjectid PropertyHint = 29
/*Hints that an [int] property is a pointer. Used by GDExtension.*/
	PropertyHintIntIsPointer PropertyHint = 30
/*Hints that a property is an [Array] with the stored type specified in the hint string.*/
	PropertyHintArrayType PropertyHint = 31
/*Hints that a string property is a locale code. Editing it will show a locale dialog for picking language and country.*/
	PropertyHintLocaleId PropertyHint = 32
/*Hints that a dictionary property is string translation map. Dictionary keys are locale codes and, values are translated strings.*/
	PropertyHintLocalizableString PropertyHint = 33
/*Hints that a property is an instance of a [Node]-derived type, optionally specified via the hint string (e.g. [code]"Node2D"[/code]). Editing it will show a dialog for picking a node from the scene.*/
	PropertyHintNodeType PropertyHint = 34
/*Hints that a quaternion property should disable the temporary euler editor.*/
	PropertyHintHideQuaternionEdit PropertyHint = 35
/*Hints that a string property is a password, and every character is replaced with the secret character.*/
	PropertyHintPassword PropertyHint = 36
/*Represents the size of the [enum PropertyHint] enum.*/
	PropertyHintMax PropertyHint = 38
)

type PropertyUsageFlags int

const (
/*The property is not stored, and does not display in the editor. This is the default for non-exported properties.*/
	PropertyUsageNone PropertyUsageFlags = 0
/*The property is serialized and saved in the scene file (default for exported properties).*/
	PropertyUsageStorage PropertyUsageFlags = 2
/*The property is shown in the [EditorInspector] (default for exported properties).*/
	PropertyUsageEditor PropertyUsageFlags = 4
/*The property is excluded from the class reference.*/
	PropertyUsageInternal PropertyUsageFlags = 8
/*The property can be checked in the [EditorInspector].*/
	PropertyUsageCheckable PropertyUsageFlags = 16
/*The property is checked in the [EditorInspector].*/
	PropertyUsageChecked PropertyUsageFlags = 32
/*Used to group properties together in the editor. See [EditorInspector].*/
	PropertyUsageGroup PropertyUsageFlags = 64
/*Used to categorize properties together in the editor.*/
	PropertyUsageCategory PropertyUsageFlags = 128
/*Used to group properties together in the editor in a subgroup (under a group). See [EditorInspector].*/
	PropertyUsageSubgroup PropertyUsageFlags = 256
/*The property is a bitfield, i.e. it contains multiple flags represented as bits.*/
	PropertyUsageClassIsBitfield PropertyUsageFlags = 512
/*The property does not save its state in [PackedScene].*/
	PropertyUsageNoInstanceState PropertyUsageFlags = 1024
/*Editing the property prompts the user for restarting the editor.*/
	PropertyUsageRestartIfChanged PropertyUsageFlags = 2048
/*The property is a script variable which should be serialized and saved in the scene file.*/
	PropertyUsageScriptVariable PropertyUsageFlags = 4096
/*The property value of type [Object] will be stored even if its value is [code]null[/code].*/
	PropertyUsageStoreIfNull PropertyUsageFlags = 8192
/*If this property is modified, all inspector fields will be refreshed.*/
	PropertyUsageUpdateAllIfModified PropertyUsageFlags = 16384
	PropertyUsageScriptDefaultValue PropertyUsageFlags = 32768
/*The property is an enum, i.e. it only takes named integer constants from its associated enumeration.*/
	PropertyUsageClassIsEnum PropertyUsageFlags = 65536
/*If property has [code]nil[/code] as default value, its type will be [Variant].*/
	PropertyUsageNilIsVariant PropertyUsageFlags = 131072
/*The property is an array.*/
	PropertyUsageArray PropertyUsageFlags = 262144
/*When duplicating a resource with [method Resource.duplicate], and this flag is set on a property of that resource, the property should always be duplicated, regardless of the [code]subresources[/code] bool parameter.*/
	PropertyUsageAlwaysDuplicate PropertyUsageFlags = 524288
/*When duplicating a resource with [method Resource.duplicate], and this flag is set on a property of that resource, the property should never be duplicated, regardless of the [code]subresources[/code] bool parameter.*/
	PropertyUsageNeverDuplicate PropertyUsageFlags = 1048576
/*The property is only shown in the editor if modern renderers are supported (the Compatibility rendering method is excluded).*/
	PropertyUsageHighEndGfx PropertyUsageFlags = 2097152
/*The [NodePath] property will always be relative to the scene's root. Mostly useful for local resources.*/
	PropertyUsageNodePathFromSceneRoot PropertyUsageFlags = 4194304
/*Use when a resource is created on the fly, i.e. the getter will always return a different instance. [ResourceSaver] needs this information to properly save such resources.*/
	PropertyUsageResourceNotPersistent PropertyUsageFlags = 8388608
/*Inserting an animation key frame of this property will automatically increment the value, allowing to easily keyframe multiple values in a row.*/
	PropertyUsageKeyingIncrements PropertyUsageFlags = 16777216
	PropertyUsageDeferredSetResource PropertyUsageFlags = 33554432
/*When this property is a [Resource] and base object is a [Node], a resource instance will be automatically created whenever the node is created in the editor.*/
	PropertyUsageEditorInstantiateObject PropertyUsageFlags = 67108864
/*The property is considered a basic setting and will appear even when advanced mode is disabled. Used for project settings.*/
	PropertyUsageEditorBasicSetting PropertyUsageFlags = 134217728
/*The property is read-only in the [EditorInspector].*/
	PropertyUsageReadOnly PropertyUsageFlags = 268435456
/*An export preset property with this flag contains confidential information and is stored separately from the rest of the export preset configuration.*/
	PropertyUsageSecret PropertyUsageFlags = 536870912
/*Default usage (storage and editor).*/
	PropertyUsageDefault PropertyUsageFlags = 6
/*Default usage but without showing the property in the editor (storage).*/
	PropertyUsageNoEditor PropertyUsageFlags = 2
)

type MethodFlags int

const (
/*Flag for a normal method.*/
	MethodFlagNormal MethodFlags = 1
/*Flag for an editor method.*/
	MethodFlagEditor MethodFlags = 2
/*Flag for a constant method.*/
	MethodFlagConst MethodFlags = 4
/*Flag for a virtual method.*/
	MethodFlagVirtual MethodFlags = 8
/*Flag for a method with a variable number of arguments.*/
	MethodFlagVararg MethodFlags = 16
/*Flag for a static method.*/
	MethodFlagStatic MethodFlags = 32
/*Used internally. Allows to not dump core virtual methods (such as [method Object._notification]) to the JSON API.*/
	MethodFlagObjectCore MethodFlags = 64
/*Default method flags (normal).*/
	MethodFlagsDefault MethodFlags = 1
)

type VariantType int

const (
/*Variable is [code]null[/code].*/
	TypeNil VariantType = 0
/*Variable is of type [bool].*/
	TypeBool VariantType = 1
/*Variable is of type [int].*/
	TypeInt VariantType = 2
/*Variable is of type [float].*/
	TypeFloat VariantType = 3
/*Variable is of type [String].*/
	TypeString VariantType = 4
/*Variable is of type [Vector2].*/
	TypeVector2 VariantType = 5
/*Variable is of type [Vector2i].*/
	TypeVector2i VariantType = 6
/*Variable is of type [Rect2].*/
	TypeRect2 VariantType = 7
/*Variable is of type [Rect2i].*/
	TypeRect2i VariantType = 8
/*Variable is of type [Vector3].*/
	TypeVector3 VariantType = 9
/*Variable is of type [Vector3i].*/
	TypeVector3i VariantType = 10
/*Variable is of type [Transform2D].*/
	TypeTransform2d VariantType = 11
/*Variable is of type [Vector4].*/
	TypeVector4 VariantType = 12
/*Variable is of type [Vector4i].*/
	TypeVector4i VariantType = 13
/*Variable is of type [Plane].*/
	TypePlane VariantType = 14
/*Variable is of type [Quaternion].*/
	TypeQuaternion VariantType = 15
/*Variable is of type [AABB].*/
	TypeAabb VariantType = 16
/*Variable is of type [Basis].*/
	TypeBasis VariantType = 17
/*Variable is of type [Transform3D].*/
	TypeTransform3d VariantType = 18
/*Variable is of type [Projection].*/
	TypeProjection VariantType = 19
/*Variable is of type [Color].*/
	TypeColor VariantType = 20
/*Variable is of type [StringName].*/
	TypeStringName VariantType = 21
/*Variable is of type [NodePath].*/
	TypeNodePath VariantType = 22
/*Variable is of type [RID].*/
	TypeRid VariantType = 23
/*Variable is of type [Object].*/
	TypeObject VariantType = 24
/*Variable is of type [Callable].*/
	TypeCallable VariantType = 25
/*Variable is of type [Signal].*/
	TypeSignal VariantType = 26
/*Variable is of type [Dictionary].*/
	TypeDictionary VariantType = 27
/*Variable is of type [Array].*/
	TypeArray VariantType = 28
/*Variable is of type [PackedByteArray].*/
	TypePackedByteArray VariantType = 29
/*Variable is of type [PackedInt32Array].*/
	TypePackedInt32Array VariantType = 30
/*Variable is of type [PackedInt64Array].*/
	TypePackedInt64Array VariantType = 31
/*Variable is of type [PackedFloat32Array].*/
	TypePackedFloat32Array VariantType = 32
/*Variable is of type [PackedFloat64Array].*/
	TypePackedFloat64Array VariantType = 33
/*Variable is of type [PackedStringArray].*/
	TypePackedStringArray VariantType = 34
/*Variable is of type [PackedVector2Array].*/
	TypePackedVector2Array VariantType = 35
/*Variable is of type [PackedVector3Array].*/
	TypePackedVector3Array VariantType = 36
/*Variable is of type [PackedColorArray].*/
	TypePackedColorArray VariantType = 37
/*Variable is of type [PackedVector4Array].*/
	TypePackedVector4Array VariantType = 38
/*Represents the size of the [enum Variant.Type] enum.*/
	TypeMax VariantType = 39
)

type VariantOperator int

const (
/*Equality operator ([code]==[/code]).*/
	OpEqual VariantOperator = 0
/*Inequality operator ([code]!=[/code]).*/
	OpNotEqual VariantOperator = 1
/*Less than operator ([code]<[/code]).*/
	OpLess VariantOperator = 2
/*Less than or equal operator ([code]<=[/code]).*/
	OpLessEqual VariantOperator = 3
/*Greater than operator ([code]>[/code]).*/
	OpGreater VariantOperator = 4
/*Greater than or equal operator ([code]>=[/code]).*/
	OpGreaterEqual VariantOperator = 5
/*Addition operator ([code]+[/code]).*/
	OpAdd VariantOperator = 6
/*Subtraction operator ([code]-[/code]).*/
	OpSubtract VariantOperator = 7
/*Multiplication operator ([code]*[/code]).*/
	OpMultiply VariantOperator = 8
/*Division operator ([code]/[/code]).*/
	OpDivide VariantOperator = 9
/*Unary negation operator ([code]-[/code]).*/
	OpNegate VariantOperator = 10
/*Unary plus operator ([code]+[/code]).*/
	OpPositive VariantOperator = 11
/*Remainder/modulo operator ([code]%[/code]).*/
	OpModule VariantOperator = 12
/*Power operator ([code]**[/code]).*/
	OpPower VariantOperator = 13
/*Left shift operator ([code]<<[/code]).*/
	OpShiftLeft VariantOperator = 14
/*Right shift operator ([code]>>[/code]).*/
	OpShiftRight VariantOperator = 15
/*Bitwise AND operator ([code]&[/code]).*/
	OpBitAnd VariantOperator = 16
/*Bitwise OR operator ([code]|[/code]).*/
	OpBitOr VariantOperator = 17
/*Bitwise XOR operator ([code]^[/code]).*/
	OpBitXor VariantOperator = 18
/*Bitwise NOT operator ([code]~[/code]).*/
	OpBitNegate VariantOperator = 19
/*Logical AND operator ([code]and[/code] or [code]&&[/code]).*/
	OpAnd VariantOperator = 20
/*Logical OR operator ([code]or[/code] or [code]||[/code]).*/
	OpOr VariantOperator = 21
/*Logical XOR operator (not implemented in GDScript).*/
	OpXor VariantOperator = 22
/*Logical NOT operator ([code]not[/code] or [code]![/code]).*/
	OpNot VariantOperator = 23
/*Logical IN operator ([code]in[/code]).*/
	OpIn VariantOperator = 24
/*Represents the size of the [enum Variant.Operator] enum.*/
	OpMax VariantOperator = 25
)

