/*
[gdscript]
print(OS.is_keycode_unicode(KEY_G))      # Prints true
print(OS.is_keycode_unicode(KEY_KP_4))   # Prints true
print(OS.is_keycode_unicode(KEY_TAB))    # Prints false
print(OS.is_keycode_unicode(KEY_ESCAPE)) # Prints false
[/gdscript]
[csharp]
GD.Print(OS.IsKeycodeUnicode((long)Key.G));      // Prints True
GD.Print(OS.IsKeycodeUnicode((long)Key.Kp4));    // Prints True
GD.Print(OS.IsKeycodeUnicode((long)Key.Tab));    // Prints False
GD.Print(OS.IsKeycodeUnicode((long)Key.Escape)); // Prints False
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/OS"
)

func OS_IsKeycodeUnicode() {
	fmt.Println(OS.IsKeycodeUnicode(Input.KeyG))      // Prints True
	fmt.Println(OS.IsKeycodeUnicode(Input.KeyKp4))    // Prints True
	fmt.Println(OS.IsKeycodeUnicode(Input.KeyTab))    // Prints False
	fmt.Println(OS.IsKeycodeUnicode(Input.KeyEscape)) // Prints False
}
