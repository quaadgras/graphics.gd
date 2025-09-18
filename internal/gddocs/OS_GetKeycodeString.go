/*
[gdscript]
print(OS.get_keycode_string(KEY_C))                    # Prints "C"
print(OS.get_keycode_string(KEY_ESCAPE))               # Prints "Escape"
print(OS.get_keycode_string(KEY_MASK_SHIFT | KEY_TAB)) # Prints "Shift+Tab"
[/gdscript]
[csharp]
GD.Print(OS.GetKeycodeString(Key.C));                                    // Prints "C"
GD.Print(OS.GetKeycodeString(Key.Escape));                               // Prints "Escape"
GD.Print(OS.GetKeycodeString((Key)KeyModifierMask.MaskShift | Key.Tab)); // Prints "Shift+Tab"
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/OS"
)

func OS_GetKeycodeString() {
	fmt.Println(OS.GetKeycodeString(Input.KeyC))                                   // Prints "C"
	fmt.Println(OS.GetKeycodeString(Input.KeyEscape))                              // Prints "Escape"
	fmt.Println(OS.GetKeycodeString(Input.Key(Input.KeyMaskShift) | Input.KeyTab)) // Prints "Shift+Tab"
}
