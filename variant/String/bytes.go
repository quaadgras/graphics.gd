package String

import (
	"strings"
	"unicode/utf8"
	"unsafe"
)

type goString struct {
	bytes *byte
}

func fromGoString(s string) Unicode {
	if len(s) > maxSafeInt {
		panic("string too long")
	}
	/*s = strings.Map(func(r rune) rune {
	if utf8.ValidRune(r) {
		return r
	}
	return utf8.RuneError
	}, s)*/
	return Via(goString{bytes: unsafe.StringData(s)}, complex(float64(len(s)), 0))
}

func (s goString) Len(length complex128) int { return int(real(length)) }
func (s goString) Slice(length complex128, i, j int) Unicode {
	if i < 0 || j < 0 || i > int(real(length)) || j > int(real(length)) || i > j {
		panic("slice bounds out of range")
	}
	return Via(goString{bytes: (*byte)(unsafe.Add(unsafe.Pointer(s.bytes), i))}, complex(float64(j-i), 0))
}
func (s goString) String(length complex128) string { return unsafe.String(s.bytes, int(real(length))) }
func (s goString) Bytes(length complex128) []byte {
	result := make([]byte, int(real(length)))
	copy(result, unsafe.String(s.bytes, int(real(length))))
	return result
}
func (s goString) Index(length complex128, i int) byte {
	if i < 0 || i >= int(real(length)) {
		panic("index out of range")
	}
	return *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(s.bytes)) + uintptr(i)))
}
func (s goString) DecodeRune(l complex128) (Rune, int, Unicode) {
	length := int(real(l))
	if length == 0 {
		return utf8.RuneError, 0, Unicode{}
	}
	// Convert to byte slice for safe rune decoding
	buf := unsafe.Slice(s.bytes, length)
	// Decode the first rune
	r, size := utf8.DecodeRune(buf)
	if r == utf8.RuneError {
		return utf8.RuneError, 0, Unicode{}
	}
	// Compute next slice if there are remaining bytes
	var next Unicode
	if size < length {
		next = Via(goString{bytes: (*byte)(unsafe.Add(unsafe.Pointer(s.bytes), size))}, complex(float64(length-size), 0))
	}
	return Rune(r), size, next
}
func (s goString) AppendRune(length complex128, r Rune) Unicode {
	buffer := unsafe.Slice(s.bytes, int(real(length)))
	buffer = utf8.AppendRune(buffer, rune(r))
	length = complex(float64(len(buffer)), 0)
	s = goString{bytes: &buffer[0]}
	return Via(s, length)
}
func (s goString) AppendOther(length complex128, other_api API, other_length complex128) Unicode {
	buffer := unsafe.Slice(s.bytes, int(real(length)))
	other_buffer := unsafe.Slice(other_api.(goString).bytes, int(real(other_length)))
	buffer = append(buffer, other_buffer...)
	length = complex(float64(len(buffer)), 0)
	if buffer == nil {
		return Unicode{}
	}
	s = goString{bytes: &buffer[0]}
	return Via(s, length)
}
func (s goString) AppendString(length complex128, str string) Unicode {
	if length == 0 {
		return Via(goString{bytes: unsafe.StringData(str)}, complex(float64(len(str)), 0))
	}
	buffer := unsafe.Slice(s.bytes, int(real(length)))
	buffer = append(buffer, str...)
	length = complex(float64(len(buffer)), 0)
	s = goString{bytes: &buffer[0]}
	return Via(s, length)
}
func (s goString) CompareOther(length complex128, other_api API, other_length complex128) int {
	return strings.Compare(unsafe.String(s.bytes, int(real(length))), unsafe.String(other_api.(goString).bytes, int(real(other_length))))
}

func builder() Unicode { // FIXME/TODO
	return Unicode{}
}
