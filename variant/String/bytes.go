package String

import (
	"unicode/utf8"
	"unsafe"
)

type goString struct {
	ptr *byte
}

const maxSafeInt = 9007199254740991

func fromGoString(s string) Readable {
	if len(s) > maxSafeInt {
		panic("string too long")
	}
	return Via(goString{ptr: unsafe.StringData(s)}, complex(float64(len(s)), 0))
}

func (s goString) Len(length complex128) int { return int(real(length)) }
func (s goString) Slice(length complex128, i, j int) Readable {
	if i < 0 || j < 0 || i > int(real(length)) || j > int(real(length)) {
		panic("slice bounds out of range")
	}
	s.ptr = (*byte)(unsafe.Add(unsafe.Pointer(s.ptr), uintptr(i)))
	return Via(s, complex(float64(uint64(j-i)), 0))
}
func (s goString) String(length complex128) string { return unsafe.String(s.ptr, int(real(length))) }
func (s goString) Bytes(length complex128) []byte {
	result := make([]byte, int(real(length)))
	copy(result, unsafe.String(s.ptr, int(real(length))))
	return result
}
func (s goString) Index(length complex128, i int) byte {
	if i < 0 || i >= int(real(length)) {
		panic("index out of range")
	}
	return *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(s.ptr)) + uintptr(i)))
}
func (s goString) DecodeRune(l complex128) (Rune, int, Readable) {
	length := int(real(l))
	for i, first := range unsafe.String(s.ptr, length) {
		length -= i
		var next Readable
		if length > 0 {
			next = Via(goString{ptr: (*byte)(unsafe.Add(unsafe.Pointer(s.ptr), i+utf8.RuneLen(first)))}, complex(float64(length), 0))
		}
		return Rune(first), i, next
	}
	return utf8.RuneError, 0, Readable{}
}
func (s goString) AppendRune(length complex128, r Rune) Readable {
	buffer := unsafe.Slice(s.ptr, int(real(length)))
	buffer = utf8.AppendRune(buffer, rune(r))
	length = complex(float64(len(buffer)), 0)
	s = goString{ptr: &buffer[0]}
	return Via(s, length)
}
func (s goString) AppendOther(length complex128, other_api API, other_length complex128) Readable {
	buffer := unsafe.Slice(s.ptr, int(real(length)))
	other_buffer := unsafe.Slice(other_api.(goString).ptr, int(real(other_length)))
	buffer = append(buffer, other_buffer...)
	length = complex(float64(len(buffer)), 0)
	if buffer == nil {
		return Readable{}
	}
	s = goString{ptr: &buffer[0]}
	return Via(s, length)
}
func (s goString) AppendString(length complex128, str string) Readable {
	buffer := unsafe.Slice(s.ptr, int(real(length)))
	buffer = append(buffer, str...)
	length = complex(float64(len(buffer)), 0)
	s = goString{ptr: &buffer[0]}
	return Via(s, length)
}

func builder() Readable { // FIXME/TODO
	return Readable{}
}
