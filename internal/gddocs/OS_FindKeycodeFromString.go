/*
[gdscript]
print(OS.find_keycode_from_string("C"))         # Prints 67 (KEY_C)
print(OS.find_keycode_from_string("Escape"))    # Prints 4194305 (KEY_ESCAPE)
print(OS.find_keycode_from_string("Shift+Tab")) # Prints 37748738 (KEY_MASK_SHIFT | KEY_TAB)
print(OS.find_keycode_from_string("Unknown"))   # Prints 0 (KEY_NONE)
[/gdscript]
[csharp]
GD.Print(OS.FindKeycodeFromString("C"));         // Prints C (Key.C)
GD.Print(OS.FindKeycodeFromString("Escape"));    // Prints Escape (Key.Escape)
GD.Print(OS.FindKeycodeFromString("Shift+Tab")); // Prints 37748738 (KeyModifierMask.MaskShift | Key.Tab)
GD.Print(OS.FindKeycodeFromString("Unknown"));   // Prints None (Key.None)
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/OS"
)

func OS_FindKeycodeFromString() {
	fmt.Println(OS.FindKeycodeFromString("C"))         // Prints 67 (KEY_C)
	fmt.Println(OS.FindKeycodeFromString("Escape"))    // Prints 4194305 (KEY_ESCAPE)
	fmt.Println(OS.FindKeycodeFromString("Shift+Tab")) // Prints 37748738 (KEY_MASK_SHIFT | KEY_TAB)
	fmt.Println(OS.FindKeycodeFromString("Unknown"))   // Prints 0 (KEY_NONE)
}
