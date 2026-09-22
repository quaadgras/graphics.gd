// IOKit special key types for graphics.gd's macOS SDK (the data1 of
// NSEventTypeSystemDefined events carries one in its top byte).
//
// Written from the public documentation of the API. Not derived from
// Apple's SDK headers.
#ifndef GD_IOKIT_EV_KEYMAP_H
#define GD_IOKIT_EV_KEYMAP_H

#define NX_KEYTYPE_SOUND_UP 0
#define NX_KEYTYPE_SOUND_DOWN 1
#define NX_KEYTYPE_BRIGHTNESS_UP 2
#define NX_KEYTYPE_BRIGHTNESS_DOWN 3
#define NX_KEYTYPE_CAPS_LOCK 4
#define NX_KEYTYPE_HELP 5
#define NX_POWER_KEY 6
#define NX_KEYTYPE_MUTE 7
#define NX_UP_ARROW_KEY 8
#define NX_DOWN_ARROW_KEY 9
#define NX_KEYTYPE_NUM_LOCK 10
#define NX_KEYTYPE_CONTRAST_UP 11
#define NX_KEYTYPE_CONTRAST_DOWN 12
#define NX_KEYTYPE_LAUNCH_PANEL 13
#define NX_KEYTYPE_EJECT 14
#define NX_KEYTYPE_VIDMIRROR 15
#define NX_KEYTYPE_PLAY 16
#define NX_KEYTYPE_NEXT 17
#define NX_KEYTYPE_PREVIOUS 18
#define NX_KEYTYPE_FAST 19
#define NX_KEYTYPE_REWIND 20
#define NX_KEYTYPE_ILLUMINATION_UP 21
#define NX_KEYTYPE_ILLUMINATION_DOWN 22
#define NX_KEYTYPE_ILLUMINATION_TOGGLE 23
#define NX_NUMSPECIALKEYS 24

#endif
