/*
[gdscript]
const MAX_15B = 1 << 15
const MAX_16B = 1 << 16

func unsigned16_to_signed(unsigned):
	return (unsigned + MAX_15B) % MAX_16B - MAX_15B

func _ready():
	var f = FileAccess.open("user://file.dat", FileAccess.WRITE_READ)
	f.store_16(-42) # This wraps around and stores 65494 (2^16 - 42).
	f.store_16(121) # In bounds, will store 121.
	f.seek(0) # Go back to start to read the stored value.
	var read1 = f.get_16() # 65494
	var read2 = f.get_16() # 121
	var converted1 = unsigned16_to_signed(read1) # -42
	var converted2 = unsigned16_to_signed(read2) # 121
[/gdscript]
[csharp]
public override void _Ready()
{
	using var f = FileAccess.Open("user://file.dat", FileAccess.ModeFlags.WriteRead);
	f.Store16(unchecked((ushort)-42)); // This wraps around and stores 65494 (2^16 - 42).
	f.Store16(121); // In bounds, will store 121.
	f.Seek(0); // Go back to start to read the stored value.
	ushort read1 = f.Get16(); // 65494
	ushort read2 = f.Get16(); // 121
	short converted1 = (short)read1; // -42
	short converted2 = (short)read2; // 121
}
[/csharp]
*/

package main

import "graphics.gd/classdb/FileAccess"

func ExampleFileAccessStore16() {
	const MAX_15B = 1 << 15
	const MAX_16B = 1 << 16
	unsigned16ToSigned := func(unsigned int) int { return (unsigned+MAX_15B)%MAX_16B - MAX_15B }

	var f = FileAccess.Open("user://file.dat", FileAccess.WriteRead)
	f.Store16(-42)                                  // This wraps around and stores 65494 (2^16 - 42).
	f.Store16(121)                                  // In bounds, will store 121.
	f.SeekTo(0)                                     // Go back to start to read the stored value.
	var read1 = f.Get16()                           // 65494
	var read2 = f.Get16()                           // 121
	var converted1 = unsigned16ToSigned(int(read1)) // -42
	var converted2 = unsigned16ToSigned(int(read2)) // 121
	_, _ = converted1, converted2
}
