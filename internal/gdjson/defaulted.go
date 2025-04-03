package gdjson

func IsTheDefaultValueZero(value string) bool {
	switch value {
	case "0", `""`, `&""`, `Vector2(0, 0)`, `[]`, `PackedByteArray()`, `Rect2i(0, 0, 0, 0)`, `RID()`,
		`Vector2i(0, 0)`, `Vector3(0, 0, 0)`, `PackedStringArray()`, `Array[RID]([])`, `false`, `{}`, `null`, `Callable()`,
		`PackedVector2Array()`, `PackedInt32Array()`, `Rect2(0, 0, 0, 0)`, `PackedFloat32Array()`, `Array[StringName]([])`, `0.0`:
		return true
	default:
		return false
	}
}
