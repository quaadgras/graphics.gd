// Package String provides a generic string functions.
package String

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"iter"
	"net/netip"
	"net/url"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
	"unique"
	"unsafe"

	"golang.org/x/text/encoding/unicode/utf32"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Int"
)

// Readable string containing human-readable characters in an implementation-specific encoding.
type Readable Generic

// String returns a copy of the Readable as a raw string.
func (utf Readable) String() string {
	if utf.api == nil {
		return ""
	}
	return utf.api.String(utf.buf)
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (utf Readable) MarshalText() ([]byte, error) {
	return []byte(utf.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (utf *Readable) UnmarshalText(text []byte) error {
	*utf = New(string(text))
	return nil
}

// Size returns the number of bytes of data stored in the Readable.
func (utf Readable) Size() int {
	if utf.api == nil {
		return 0
	}
	return utf.api.Len(utf.buf)
}

// Generic can be used as a ~T parameter that accept any [Readable]-like type.
type Generic = struct {
	_   [0]func()
	buf complex128
	api API
}

// Comparable string optimized for use as a key, such that comparisons
// are efficient.
type Comparable = struct {
	ptr *Readable
}

// Name will be replaced with [Comparable] in Go 1.24
type Name Generic

// String implements the [fmt.Stringer] interface.
func (name Name) String() string {
	return Readable(name).String()
}

func MakeComparable[T Any](s T) Comparable {
	panic("not implemented")
}

// New returns a new [Readable] string concatenated from the given values.
func New(val ...any) Readable { //gd:String() str
	if len(val) == 0 {
		return Readable{}
	}
	if len(val) == 1 {
		switch v := val[0].(type) {
		case Rune:
			return fromGoString(string(rune(v)))
		case string:
			return fromGoString(v)
		case Readable:
			return v
		case []byte:
			return fromGoString(string(v))
		}
	}
	return newAs[Readable](val...)
}

func newAs[T Any](val ...any) T {
	return As[T](fromGoString(fmt.Sprint(val...))) // FIXME optimise
}

// API required to implement a [Readable] string.
type API interface {
	Len(complex128) int
	Slice(complex128, int, int) Readable
	String(complex128) string
	Index(complex128, int) byte
	DecodeRune(complex128) (Rune, int, Readable)
	AppendRune(complex128, Rune) Readable
	AppendOther(complex128, API, complex128) Readable
	AppendString(complex128, string) Readable
	CompareOther(complex128, API, complex128) int
}

// Proxy can be used to transform the underlying encoding and implementation of a [Readable]
// string into a specific type. The alloc function should return a new [Implementation] of
// the desired type, and the internal state. This may be cached in the existing implementation
// and, if so, will be passed to the check function so that it can decide whether to reuse
// the cache or not.
//
// The uint64 can be used to avoid unnecessary allocations, such that T can be zero-sized if
// the implementation does not require more that 64bits of state.
func Proxy[S ~Generic, T API](s S, reuse func(T, complex128) bool, alloc func() (T, complex128)) (T, complex128) {
	utf := Readable(s)
	if utf.api == nil {
		return alloc()
	}
	already, ok := utf.api.(T)
	if ok && reuse(already, utf.buf) {
		return already, utf.buf
	}
	b, state := alloc()
	appended := b.AppendString(state, As[string](s))
	b, state = appended.api.(T), appended.buf
	/*for _, c := range Runes(utf) {
	appended := b.AppendRune(state, c)
	b, state = appended.api.(T), appended.buf
	}*/
	return b, state
}

func Via(impl API, state complex128) Readable {
	return Readable{api: impl, buf: state}
}

type Rune rune

type Any interface {
	~string | ~[]byte | ~Generic | ~Comparable | unique.Handle[string] | Rune | []rune
}

func As[T, S Any](s S) T {
	uni := New(s)
	rtype := reflect.TypeFor[T]()
	switch {
	case rtype == reflect.TypeFor[string]():
		c := uni.api.String(uni.buf)
		return *(*T)(unsafe.Pointer(&c))
	case rtype.Kind() == reflect.Slice && rtype.Elem().Kind() == reflect.Uint8:
		b := []byte(uni.api.String(uni.buf))
		return *(*T)(unsafe.Pointer(&b))
	case rtype.ConvertibleTo(reflect.TypeOf(Readable{})):
		u := uni
		return *(*T)(unsafe.Pointer(&u))
	case rtype.ConvertibleTo(reflect.TypeOf(Comparable{})):
		c := MakeComparable(uni)
		return *(*T)(unsafe.Pointer(&c))
	default:
		panic("unreachable")
	}
}

// Append returns a new string with the given string appended to the end.
func Append[S Any, B Any](a S, b B) S { //gd:String.append
	s1, s2 := New(a), New(b)
	if s2.api == nil {
		return a
	}
	if s1.api == nil {
		s1 = Readable{
			api: goString{},
		}
	}
	if reflect.TypeOf(s1.api) == reflect.TypeOf(s2.api) {
		return As[S](s1.api.AppendOther(s1.buf, s2.api, s2.buf))
	}
	return As[S](s1.api.AppendString(s1.buf, s2.api.String(s2.buf)))
}

// Index returns the byte at the given index in the string. If the index is out of range, this
// function will panic. See also [Slice].
func Index[S Any](s S, i int) byte {
	utf := New(s)
	return utf.api.Index(utf.buf, i)
}

// HasPrefix returns true if the string begins with the given text. See also [HasSuffix].
func HasPrefix[S, P Any](s S, prefix P) bool { //gd:String.begins_with
	return Length(prefix) <= Length(s) && Comparison(Slice(s, 0, Length(prefix)), prefix) == 0
}

// HasSuffix returns true if the string ends with the given text. See also [HasSuffix].
func HasSuffix[S, B Any](s S, suffix B) bool { //gd:String.ends_with
	return Length(suffix) <= Length(s) && Comparison(Slice(s, Length(s)-Length(suffix), Length(s)), suffix) == 0
}

// Bigrams returns a sequence of bigrams (pairs of consecutive characters) in this string.
func Bigrams[S Any](s S) iter.Seq[S] { //gd:String.bigrams
	return func(yield func(S) bool) {
		if Length(s) < 2 {
			return
		}
		for i := 0; i < Length(s)-1; i++ {
			if !yield(Slice(s, i, i+2)) {
				return
			}
		}
	}
}

// BinaryToInteger converts the string representing a binary number into an int. The string may
// optionally be prefixed with "0b", and an additional - prefix for negative numbers.
func BinaryToInteger[S Any](s S) int { //gd:String.bin_to_int
	s = TrimPrefix(s, "0b")
	var r int
	for i := 0; i < Length(s); i++ {
		b := Index(s, i)
		if b == '0' || b == '1' {
			r = r*2 + int(b-'0')
		}
	}
	if HasPrefix(s, "-") {
		r = -r
	}
	return r
}

// Escape returns a copy of the string with special characters escaped using the Go language standard.
func Escape[S Any](s S) S { //gd:String.c_escape String.json_escape String.xml_escape
	escaped := strconv.QuoteToASCII(As[string](s))
	return As[S](escaped[1 : len(escaped)-1])
}

// Unescape returns a copy of the string with escaped characters replaced by their meanings.
func Unescape[S Any](s S) S { //gd:String.c_unescape String.xml_unescape
	unescaped, _ := strconv.Unquote(`"` + As[string](s) + `"`)
	return As[S](unescaped)
}

// Runes returns a sequence of runes in this string and their indices.
func Runes[S Any](s S) iter.Seq2[int, Rune] { //gd:String.chars
	return func(yield func(int, Rune) bool) {
		utf := New(s)
		for i := 0; i < utf.api.Len(utf.buf); i++ {
			r, _, next := utf.api.DecodeRune(utf.buf)
			if r == utf8.RuneError {
				break
			}
			if !yield(i, r) {
				return
			}
			utf = next
		}
	}
}

// firstRune returns the first rune in the string and its width in bytes.
func firstRune[S Any](s S) (Rune, bool) {
	utf := New(s)
	r, _, _ := utf.api.DecodeRune(utf.buf)
	return r, r != utf8.RuneError
}

// Title returns a copy of the string s with all Unicode letters that begin words
// mapped to their Unicode title case. Is not always reliable in the presence of
// punctuation.
func Title[S Any](s S) S {
	// Use a closure here to remember state.
	// Hackish but effective. Depends on Map scanning in order and calling
	// the closure once per rune.
	var prev Rune = ' '
	return Map(
		func(r Rune) Rune {
			if isSeparator(rune(prev)) {
				prev = r
				return Rune(unicode.ToTitle(rune(r)))
			}
			prev = r
			return r
		},
		s,
	)
}

// Map returns a copy of the string s with all its characters modified
// according to the mapping function. If mapping returns a negative value, the character is
// dropped from the string with no replacement.
func Map[S Any](mapping func(Rune) Rune, s S) S {
	var (
		buf Readable = builder()
	)
	for i, c := range Runes(s) {
		r := mapping(c)
		if r == c || c == utf8.RuneError {
			continue
		}
		width := utf8.RuneLen(rune(c))
		buf = Append(buf, Slice(s, 0, i))
		if r >= 0 {
			buf = Append(buf, Rune(r))
		}
		s = Slice(s, i+width, Length(s))
		break
	}
	if buf.api == nil {
		return s // Fast path for unchanged input
	}
	for _, c := range Runes(s) {
		r := mapping(c)
		if r >= 0 {
			buf = Append(buf, New(r))
		}
	}
	return As[S](buf)
}

// Capitalize changes the appearance of the string: replaces underscores (_) with spaces, adds spaces
// before uppercase letters in the middle of a word, converts all letters to lowercase, then converts
// the first one and each one following a space to uppercase.
func Capitalize[S Any](s S) S { //gd:String.capitalize
	pieces := slices.Collect(Splits(s, "_"))
	for i, piece := range pieces {
		for i, r := range Runes(piece) {
			if unicode.IsUpper(rune(r)) {
				piece = As[S](New(Slice(piece, 0, i), " ", Slice(piece, i, Length(piece))))
			}
		}
		pieces[i] = Title(ToLower(piece))
	}
	return JoinedWith(" ", pieces...)
}

// ComparisonIgnoreCasing performs a case-insensitive comparison to another string. Returns -1
// if less than, 1 if greater than, or 0 if equal. "Less than" or "greater than"
// are determined by the Unicode code points of each string, which roughly matches
// the alphabetical order. Internally, lowercase characters are converted to
// uppercase for the comparison.
//
// With different string lengths, returns 1 if this string is longer than the to
// string, or -1 if shorter. Note that the length of empty strings is always 0.
//
// To get a bool result from a string comparison, use the == operator instead. See
// also [ComparisonStrict], [ComparisonStrictNaturalPrioritizePeriodsAndUnderscores],
// and [ComparisonNatural].
func ComparisonIgnoreCasing[A, B Any](a A, b B) int { //gd:String.nocasecmp_to
	return strings.Compare(strings.ToUpper(As[string](a)), strings.ToUpper(As[string](b)))
}

// Comparison performs a case-sensitive comparison to another string. Returns -1 if less than,
// 1 if greater than, or 0 if equal. "Less than" and "greater than" are determined by the Unicode code
// points of each string, which roughly matches the alphabetical order.
//
// With different string lengths, returns 1 if this string is longer than the to string, or -1 if shorter.
// Note that the length of empty strings is always 0.
//
// To get a bool result from a string comparison, use the == operator instead. See also [Comparison],
// [ComparisonStrictNaturalPrioritizePeriodsAndUnderscores], and [ComparisonStrictNatural].
func Comparison[A, B Any](a A, b B) int { //gd:String.casecmp_to
	if IsEmpty(a) && IsEmpty(b) {
		return 0
	}
	if IsEmpty(a) {
		return -1
	}
	if IsEmpty(b) {
		return 1
	}
	for i := 0; i < Length(a) && i < Length(b); i++ {
		if i >= Length(a) {
			return -1
		} else if i >= Length(b) {
			return 1
		}
		ac, bc := Index(a, i), Index(b, i)
		if ac < bc {
			return -1
		} else if ac > bc {
			return 1
		}
	}
	return 0
}

// ComparisonStrictNatural performs a case-sensitive, natural order comparison to another string. Returns -1
// if less than, 1 if greater than, or 0 if equal. "Less than" or "greater than" are determined by the Unicode
// code points of each string, which roughly matches the alphabetical order.
//
// When used for sorting, natural order comparison orders sequences of numbers by the combined value of
// each digit as is often expected, instead of the single digit's value. A sorted sequence of numbered
// strings will be ["1", "2", "3", ...], not ["1", "10", "2", "3", ...].
//
// With different string lengths, returns 1 if this string is longer than the to string, or -1 if shorter.
// Note that the length of empty strings is always 0.
//
// To get a bool result from a string comparison, use the == operator instead. See also [ComparisonNatural],
// [ComparisonNaturalPrioritizePeriodsAndUnderscores], and [Comparison].
func ComparisonStrictNatural[S Any](a, b S) int { //gd:String.naturalnocasecmp_to
	commonPrefix := func(a, b S) int {
		m := Length(a)
		if n := Length(b); n < m {
			m = n
		}
		if m == 0 {
			return 0
		}
		for i := 0; i < m; i++ {
			ca := Index(a, i)
			cb := Index(b, i)
			if (ca >= '0' && ca <= '9') || (cb >= '0' && cb <= '9') || ca != cb {
				return i
			}
		}
		return m
	}
	digits := func(s S) int {
		for i := 0; i < Length(s); i++ {
			c := Index(s, i)
			if c < '0' || c > '9' {
				return i
			}
		}
		return Length(s)
	}
	for {
		if p := commonPrefix(a, b); p != 0 {
			a = Slice(a, p, Length(a))
			b = Slice(b, p, Length(b))
		}
		if Length(a) == 0 {
			if Length(b) != 0 {
				return -1
			}
			return 1
		}
		if ia := digits(a); ia > 0 {
			if ib := digits(b); ib > 0 {
				// Both sides have digits.
				an, aerr := strconv.ParseUint(As[string](Slice(a, 0, ia)), 10, 64)
				bn, berr := strconv.ParseUint(As[string](Slice(b, 0, ib)), 10, 64)
				if aerr == nil && berr == nil {
					if an != bn {
						if an < bn {
							return -1
						}
						return 1
					}
					// Semantically the same digits, e.g. "00" == "0", "01" == "1". In
					// this case, only continue processing if there's trailing data on
					// both sides, otherwise do lexical comparison.
					if ia != Length(a) && ib != Length(b) {
						a = Slice(a, ia, Length(a))
						b = Slice(b, ib, Length(b))
						continue
					}
				}
			}
		}
		return Comparison(a, b)
	}
}

// ComparisonNatural performs a case-insensitive, natural order comparison to another string.
// Returns -1 if less than, 1 if greater than, or 0 if equal. "Less than" or "greater than"
// are determined by the Unicode code points of each string, which roughly matches the
// alphabetical order. Internally, lowercase characters are converted to uppercase for the comparison.
//
// When used for sorting, natural order comparison orders sequences of numbers by the combined
// value of each digit as is often expected, instead of the single digit's value. A sorted sequence
// of numbered strings will be ["1", "2", "3", ...], not ["1", "10", "2", "3", ...].
//
// With different string lengths, returns 1 if this string is longer than the to string, or -1 if
// shorter. Note that the length of empty strings is always 0.
//
// To get a bool result from a string comparison, use the == operator instead. See also
// [ComparisonStrictNatural], [ComparisonStrictNaturalPrioritizePeriodsAndUnderscores], and [Comparison].
func ComparisonNatural[S Any](a, b S) int { //gd:String.naturalcasecmp_to
	return ComparisonStrictNatural(strings.ToUpper(As[string](a)), strings.ToUpper(As[string](b)))
}

// ComparisonStrictNaturalPrioritizePeriodsAndUnderscores like [ComparisonStrictNatural] but prioritizes strings that begin with periods
// (.) and underscores (_) before any other character. Useful when sorting folders or file names.
//
// To get a bool result from a string comparison, use the == operator instead. See also [ComparisonNaturalPrioritizePeriodsAndUnderscores],
// [ComparisonStrictNatural], and [ComparisonStrict].
func ComparisonStrictNaturalPrioritizePeriodsAndUnderscores[S Any](a, b S) int { //gd:String.filecasecmp_to
	aStr := As[string](a)
	bStr := As[string](b)
	for i := 0; i < len(aStr) && i < len(bStr); i++ {
		if aStr[i] == '_' && bStr[i] != '_' {
			return -1
		}
		if aStr[i] == '.' && bStr[i] != '.' {
			return -1
		}
		if aStr[i] != '_' && bStr[i] == '_' {
			return 1
		}
		if aStr[i] != '.' && bStr[i] == '.' {
			return 1
		}
	}
	return ComparisonStrictNatural(a, b)
}

// ComparisonNaturalPrioritizePeriodsAndUnderscores like naturalnocasecmp_to but prioritizes strings that begin with
// periods (.) and underscores (_) before any other character. Useful when sorting folders or file names.
//
// To get a bool result from a string comparison, use the == operator instead. See also [ComparisonStrictNaturalPrioritizePeriodsAndUnderscores],
// [ComparisonNatural], and [Comparison].
func ComparisonNaturalPrioritizePeriodsAndUnderscores[S Any](a, b S) int { //gd:String.filenocasecmp_to
	aStr := As[string](a)
	bStr := As[string](b)
	for i := 0; i < len(aStr) && i < len(bStr); i++ {
		if aStr[i] == '_' && bStr[i] != '_' {
			return -1
		}
		if aStr[i] == '.' && bStr[i] != '.' {
			return -1
		}
		if aStr[i] != '_' && bStr[i] == '_' {
			return 1
		}
		if aStr[i] != '.' && bStr[i] == '.' {
			return 1
		}
	}
	return ComparisonNatural(a, b)
}

// Character returns a single Unicode character from the decimal char. You may use unicodelookup.com or
// unicode.org as points of reference.
func Character[I Int.Any](chr I) string { //gd:String.chr
	return string(rune(chr))
}

// ContainsStrict returns true if the string contains case-sensitive what.
// If you need to know where what is within the string, use [Find]. See also [Contains].
func ContainsStrict[S Any](what S, s S) bool { //gd:String.contains
	return strings.Contains(As[string](s), As[string](what))
}

// Contains returns true if the string contains what, ignoring case.
// If you need to know where what is within the string, use findn. See also [Contains].
func Contains[A, B Any](what A, s B) bool { //gd:String.containsn
	return strings.Contains(strings.ToLower(As[string](s)), strings.ToLower(As[string](what)))
}

// CountStrict returns the number of occurrences of the case-sensitive substring.
func CountStrict[S Any](what S, s S) int { //gd:String.count
	return strings.Count(As[string](s), As[string](what))
}

// Count returns the number of occurrences of the substring, ignoring case.
func Count[D Any, S Any](what D, s S) int { //gd:String.countn
	return strings.Count(strings.ToLower(As[string](s)), strings.ToLower(As[string](what)))
}

// Dedent returns a copy of the string with indentation (leading tabs and spaces)
// removed. See also [Indent] to add indentation.
func Dedent[S Any](s S) S { //gd:String.dedent
	return As[S](strings.TrimLeftFunc(As[string](s), unicode.IsSpace))
}

// Erase returns a string with the character erased at position.
func Erase[S Any, I Int.Any](position I, s S) S { //gd:String.erase
	return As[S](As[string](Slice(s, 0, position)) + As[string](Slice(s, position+1, Length(s))))
}

// FindStrict returns the index of the first case-sensitive occurrence of what in this string
// or -1 if not found.
func FindStrict[S Any](what S, s S) int { //gd:String.find
	return strings.Index(As[string](s), As[string](what))
}

// FindIndex returns the index of the first occurrence of what in this string, ignoring case.
// Returns -1 if not found.
func FindIndex[S, V Any](s S, what V) int { //gd:String.findn
	return strings.Index(strings.ToLower(As[string](s)), strings.ToLower(As[string](what))) // FIXME slow
}

// Decimal converts a float to a string representation of a decimal number, with
// the number of decimal places specified in decimals.
//
// The string representation may only have up to 14 significant digits, with digits
// before the decimal point having priority over digits after.
//
// Trailing zeros are not included in the string. The last digit is rounded, not truncated.
func Decimal[F Float.Any](number F) string { //gd:String.num
	return strconv.FormatFloat(float64(number), 'g', -1, 64)
}

// Integer converts the given number to a string representation.
func Integer[I Int.Any](number I) string { //gd:String.num_int64
	return strconv.Itoa(int(number))
}

// Natural converts the given unsigned int to a string representation, with the given base.
func Natural[I ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~uintptr](number I) string { //gd:String.num_uint64
	return strconv.FormatUint(uint64(number), 10)
}

// Scientific converts the given number to a string representation, in scientific notation.
func Scientific[F Float.Any](number F) string { //gd:String.num_scientific
	return strconv.FormatFloat(float64(number), 'e', -1, 64)
}

// Format formats the string by replacing all occurrences of placeholder with the elements of values.
//
// Any underscores in placeholder will be replaced with the corresponding keys in advance. Array
// elements use their index as keys.
//
// Note: The replacement of placeholders is not done all at once, instead each placeholder is replaced
// in the order they are passed, this means that if one of the replacement strings contains a key it
// will also be replaced. This can be very powerful, but can also cause unexpected results if you are
// not careful. If you do not need to perform replacement in the replacement strings, make sure your
// replacements do not contain placeholders to ensure reliable results.
func Format[S Any](format S, args ...any) S { //gd:String.format
	var result strings.Builder
	var search strings.Builder
	var braces bool
	for i := 0; i < Length(format); i++ {
		v := Index(format, i)
		switch v {
		case '{':
			braces = true
			continue
		case '}':
			if braces {
				key, err := strconv.Atoi(search.String())
				if err == nil {
					result.WriteString(fmt.Sprint(args[key]))
				} else {
					for _, arg := range args {
						rvalue := reflect.ValueOf(arg)
						switch rvalue.Kind() {
						case reflect.Map:
							if rvalue.Type().Key().Kind() == reflect.String {
								if rvalue.MapIndex(reflect.ValueOf(search.String()).Convert(rvalue.Type().Key())).IsValid() {
									result.WriteString(fmt.Sprint(rvalue.MapIndex(reflect.ValueOf(search.String())).Interface()))
								}
							}
						case reflect.Struct:
							if rvalue.FieldByName(search.String()).IsValid() {
								result.WriteString(fmt.Sprint(rvalue.FieldByName(search.String()).Interface()))
							}
						}
					}
				}
				search.Reset()
				braces = false
				continue
			}
			fallthrough
		default:
			if braces {
				search.WriteByte(v)
			} else {
				result.WriteByte(v)
			}
		}
	}
	return As[S](result.String())
}

// Extract splits the string using a delimiter and returns the substring at index. Returns
// the original string if delimiter does not occur in the string. Returns an empty string if the
// element does not exist.
//
// This is faster than split, if you only need one substring.
func Extract[S Any, D Any, I Int.Any](s S, delimiter D, index I) S { //gd:String.get_slice String.get_slicec
	var n int
	var current S
	for i := 0; i < int(index+1); i++ {
		next := FindIndex(Slice(s, n, Length(s)), delimiter)
		if next == -1 {
			return Slice(s, n, Length(s))
		}
		current = Slice(s, n, n+next)
		n += next + 1
	}
	return current
}

// NumExtractions returns the total number of extractable substrings when the string is split with
// the given delimiter (see [Splits]).
func NumExtractions[S Any, D Any](s S, delimiter D) int { //gd:String.get_slice_count
	return strings.Count(As[string](s), As[string](delimiter)) + 1
}

// Hash Returns the 32-bit hash value representing the string's contents.
//
// Note: Strings with equal hash values are not guaranteed to be the same, as a result of hash collisions.
// On the contrary, strings with different hash values are guaranteed to be different.
func Hash[S Any](s S) uint32 { //gd:String.hash
	var hashv uint32 = 5381
	for i := 0; i < Length(s); i++ {
		hashv = ((hashv << 5) + hashv) + uint32(Index(s, i))
	}
	return hashv
}

// DecodeHex decodes a hexadecimal string as a []byte.
func DecodeHex[S Any](s S) S { //gd:String.hex_decode
	b, err := hex.DecodeString(As[string](s))
	if err != nil {
		var zero S
		return zero
	}
	return As[S](b)
}

// EncodeHex encodes a []byte as a hexadecimal string.
func EncodeHex[S Any](s S) S { //gd:String.hex_encode
	return As[S](hex.EncodeToString(As[[]byte](s)))
}

// HexToInteger Converts the string representing a hexadecimal number into an int. The
// string may be optionally prefixed with "0x", and an additional - prefix for negative numbers.
func HexToInteger[S Any](s S) int { //gd:String.hex_to_int
	if strings.HasPrefix(As[string](s), "0x") {
		s = Slice(s, 2, Length(s))
	}
	i, err := strconv.ParseInt(As[string](s), 16, 0)
	if err != nil {
		return 0
	}
	return int(i)
}

// HumanReadableSize converts size which represents a number of bytes into a human-readable form.
//
// The result is in IEC prefix format, which may end in either "B", "KiB", "MiB", "GiB", "TiB", "PiB", or "EiB".
func HumanReadableSize(size int) string { //gd:String.humanize_size
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size; n >= unit && exp < 6; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(size)/float64(div), "KMGTPE"[exp])
}

// Indent indents every line of the string with the given prefix. Empty lines are not indented.
// See also [Dedent] to remove indentation.
//
// For example, the string can be indented with two tabulations using "\t\t", or four spaces using "    ".
func Indent[S Any](s, prefix S) S { //gd:String.indent
	return As[S](strings.ReplaceAll(As[string](s), "\n", "\n"+As[string](prefix)))
}

// Insert inserts what at the given position in the string.
func Insert[S Any](what, s S, position int) S { //gd:String.insert
	return As[S](As[string](Slice(s, 0, position)) + As[string](what) + As[string](Slice(s, position, Length(s))))
}

// IsEmpty returns true if the string's length is 0 (""). See also length.
func IsEmpty[S Any](s S) bool { //gd:String.is_empty
	return Length(s) == 0
}

// IsStrictSubsequenceOf returns true if all characters of this string can be found in text in their original order.
func IsStrictSubsequenceOf[S Any](text, s S) bool { //gd:String.is_subsequence_of
	if Length(s) == 0 {
		return true
	}
	n := 0
	for i := 0; i < Length(s); i++ {
		for n < Length(text) && Index(text, n) != Index(s, i) {
			n++
		}
		if n == Length(text) {
			return false
		}
		n++
	}
	return true
}

// IsSubsequenceOf returns true if all characters of this string can be found in text in any order, ignoring case.
func IsSubsequenceOf[S Any](text, s S) bool { //gd:String.is_subsequence_ofn
	if Length(s) == 0 {
		return true
	}
	n := 0
	for i := 0; i < Length(s); i++ {
		for n < Length(text) && unicode.ToLower(rune(Index(text, n))) != unicode.ToLower(rune(Index(s, i))) {
			n++
		}
		if n == Length(text) {
			return false
		}
		n++
	}
	return true
}

// IsValidFilename returns true if this string does not contain characters that are not allowed
// in file names (: / \ ? * " | % < >).
func IsValidFilename[S Any](s S) bool { //gd:String.is_valid_filename
	return !strings.ContainsAny(As[string](s), ":/\\?*\"|%<>")
}

// IsValidFloat returns true if this string represents a valid floating-point number. A valid float
// may contain only digits, one decimal point (.), and the exponent letter (e). It may also be
// prefixed with a positive (+) or negative (-) sign. Any valid integer is also a valid float
// (see [IsValidInt]). See also [ToFloat].
func IsValidFloat[S Any](s S) bool { //gd:String.is_valid_float
	_, err := strconv.ParseFloat(As[string](s), 64)
	return err == nil
}

// IsValidHexNumber returns true if this string is a valid hexadecimal number. A valid hexadecimal
// number only contains digits or letters A to F (either uppercase or lowercase), and may be prefixed
// with a positive (+) or negative (-) sign.
func IsValidHexNumber[S Any](s S) bool { //gd:String.is_valid_hex_number
	_, err := strconv.ParseInt(As[string](s), 16, 0)
	return err == nil
}

// IsValidHexColor Returns true if this string is a valid color in hexadecimal HTML notation. The string
// must be a hexadecimal value (see is_valid_hex_number) of either 3, 4, 6 or 8 digits, and may be prefixed
// by a hash sign (#). Other HTML notations for colors, such as names or hsl(), are not considered valid.
func IsValidHexColor[S Any](s S) bool { //gd:String.is_valid_html_color
	if strings.HasPrefix(As[string](s), "#") {
		s = Slice(s, 1, Length(s))
	}
	switch Length(s) {
	case 3, 4, 6, 8:
		return IsValidHexNumber(s)
	default:
		return false
	}
}

// IsValidIdentifier returns true if this string is a valid identifier. A valid identifier may contain only letters,
// digits and underscores (_), and the first character may not be a digit.
func IsValidIdentifier[S Any](s S) bool { //gd:String.is_valid_identifier
	if Length(s) == 0 {
		return false
	}
	if unicode.IsDigit(rune(Index(s, 0))) {
		return false
	}
	for _, c := range Runes(s) {
		if !unicode.IsDigit(rune(c)) && !unicode.IsLetter(rune(c)) && c != '_' {
			return false
		}
	}
	return true
}

// IsValidInt returns true if this string represents a valid integer. A valid integer only contains
// digits, and may be prefixed with a positive (+) or negative (-) sign. See also [ToInt].
func IsValidInt[S Any](s S) bool { //gd:String.is_valid_int
	_, err := strconv.ParseInt(As[string](s), 10, 0)
	return err == nil
}

// IsValidInternetProtocolAddress returns true if this string represents a well-formatted IPv4 or IPv6 address.
// This method considers reserved IP addresses such as "0.0.0.0" and "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff" as valid.
func IsValidInternetProtocolAddress[S Any](s S) bool { //gd:String.is_valid_ip_address
	_, err := netip.ParseAddr(As[string](s))
	return err == nil
}

// JoinedWith returns the concatenation of parts' elements, with each element separated by the string calling this method.
// This method is the opposite of [Splits].
func JoinedWith[S, D Any](d D, s ...S) S { //gd:String.join
	if len(s) == 0 {
		return [1]S{}[0]
	}
	var b Readable = builder()
	b = Append(b, New(fmt.Sprint(s[0])))
	for _, v := range s[1:] {
		b = Append(b, d)
		b = Append(b, v)
	}
	return As[S](b)
}

// First, returns the first n characters from the beginning of the string. If length is
// negative, strips the last length characters from the string's end.
func First[S Any, I Int.Any](n I, s S) S { //gd:String.left
	if n < 0 {
		return Slice(s, 0, Length(s)+int(n))
	}
	return Slice(s, 0, n)
}

// Length returns the number of characters in the string.
func Length[S Any](s S) int { //gd:String.length
	utf := New(s)
	if utf.api == nil {
		return 0
	}
	return utf.api.Len(utf.buf)
}

// PadPrefix formats the string to be at least min_length long by adding characters
// to the left of the string, if necessary. See also [PadSuffix].
func PadPrefix[S Any](s S, min_length int) S { //gd:String.lpad
	return JoinedWith("", As[S](strings.Repeat(" ", max(0, min_length-Length(s)))), s)
}

// StripPrefix removes a set of characters defined in chars from the string's beginning. See also StripSuffix.
//
// Note: chars is a cutset. Use [TrimPrefix] to remove a whole prefix string, rather than a set of characters.
func StripPrefix[S Any, C ~string | ~[]byte](s S, chars C) S { //gd:String.lstrip
	return As[S](strings.TrimLeft(As[string](s), As[string](chars)))
}

// MatchStrict does a simple expression match (also called "glob" or "globbing"), where * matches zero or more arbitrary
// characters and ? matches any single character except a period (.). An empty string or empty expression always
// evaluates to false.
func MatchStrict[S Any](pattern, text S) bool { //gd:String.match
	ok, _ := path.Match(As[string](pattern), As[string](text))
	return ok
}

// Match does a simple case-insensitive expression match, where * matches zero or more arbitrary characters and ? matches
// any single character except a period (.). An empty string or empty expression always evaluates to false.Does a simple
// case-insensitive expression match, where * matches zero or more arbitrary characters and ? matches any single character
// except a period (.). An empty string or empty expression always evaluates to false.
func Match[S Any](pattern, text S) bool { //gd:String.matchn
	ok, _ := path.Match(strings.ToLower(As[string](pattern)), strings.ToLower(As[string](text)))
	return ok
}

// MD5 returns the MD5 hash of the string as a byte slice.
func MD5[S Any](s S) []byte { //gd:String.md5_buffer String.md5_text
	sum := md5.Sum([]byte(As[string](s)))
	return sum[:]
}

// PadDecimals formats the string representing a number to have an exact number of digits
// after the decimal point.
func PadDecimals[S Any](s S, digits int) S { //gd:String.pad_decimals
	if !IsValidFloat(s) {
		return s
	}
	var result Readable = builder()
	result = Append(result, New(s))
	result = Append(result, New(strings.Repeat("0", Length(s)-FindIndex(s, ".")+digits+1)))
	return As[S](result)
}

// PadZeros formats the string representing a number to have an exact number of digits before
// the decimal point.
func PadZeros[S Any](s S, digits int) S { //gd:String.pad_zeros
	if !IsValidFloat(s) {
		return s
	}
	var result Readable = builder()
	result = Append(result, New(strings.Repeat("0", FindIndex(s, ".")-digits)))
	result = Append(result, New(s))
	return As[S](result)
}

// AddPathElement concatenates file at the end of the string as a subpath, adding / if necessary.
func AddPathElement[S Any](s S, file S) S { //gd:String.path_join
	if HasSuffix(s, "/") {
		return JoinedWith("", s, file)
	}
	return JoinedWith("/", s, file)
}

// Repeat repeats this string a number of times. count needs to be greater than 0.
// Otherwise, returns an empty string.
func Repeat[S Any](s S, count int) S { //gd:String.repeat
	if count <= 0 {
		return [1]S{}[0]
	}
	return As[S](strings.Repeat(As[string](s), count))
}

// ReplaceStrict replaces all occurrences of what inside the string with the given forwhat.
func ReplaceStrict[S Any](s S, what, forwhat S) S { //gd:String.replace
	return As[S](strings.Replace(As[string](s), As[string](what), As[string](forwhat), -1))
}

// Replace replaces all case-insensitive occurrences of what inside the string with the given
// forwhat.
func Replace[S Any, T Any](s S, old, new T) S { //gd:String.replacen
	var n = -1
	if Comparison(old, new) == 0 || n == 0 {
		return s // avoid allocation
	}
	// Compute number of replacements.
	if m := Count(old, s); m == 0 {
		return s // avoid allocation
	} else if n < 0 || m < n {
		n = m
	}
	// Apply replacements to buffer.
	var b = builder()
	//b.Grow(len(s) + n*(len(new)-len(old)))
	start := 0
	for i := 0; i < n; i++ {
		j := start
		if Length(old) == 0 {
			if i > 0 {
				_, wid := utf8.DecodeRuneInString(As[string](Slice(s, start, Length(s))))
				j += wid
			}
		} else {
			j += FindIndex(Slice(s, start, Length(s)), old)
		}
		b = Append(b, New(Slice(s, start, j)))
		b = Append(b, New(new))
		start = j + Length(old)
	}
	b = Append(b, New(Slice(s, start, Length(s))))
	return As[S](b)
}

// Reverse returns the copy of this string in reverse order. This operation works on
// unicode codepoints, rather than sequences of codepoints, and may break things like
// compound letters or emojis.
func Reverse[S Any](s S) S { //gd:String.reverse
	runes := []rune(As[string](s))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return As[S](New(runes))
}

// FindLastStrict returns the index of the last occurrence of what in this string, or -1 if there are none.
// This method is the reverse of [FindStrict].
func FindLastStrict[S Any](s S, what S) int { //gd:String.rfind
	return strings.LastIndex(As[string](s), As[string](what))
}

// FindLast returns the index of the last occurrence of what in this string, or -1 if there are none.
// This method is the reverse of [Find].
func FindLast[S Any, W Any](s S, what W) int { //gd:String.rfindn
	return strings.LastIndex(strings.ToLower(As[string](s)), strings.ToLower(As[string](what)))
}

// Last returns the last length characters from the end of the string. If length is
// negative, returns the first length characters from the string's beginning.
func Last[S Any](length int, s S) S { //gd:String.right
	if length < 0 {
		return Slice(s, -length, Length(s))
	}
	return Slice(s, Length(s)-length, Length(s))
}

// PadSuffix formats the string to be at least min_length long, by adding characters to the right of the
// string, if necessary. See also [PadPrefix].
func PadSuffix[S Any](s S, min_length int) S { //gd:String.rpad
	return As[S](As[string](s) + strings.Repeat(As[string](" "), max(0, min_length-Length(s))))
}

// ReverseSplit splits the string using a delimiter and returns an array of the substrings, starting from
// the end of the string. The splits in the returned array appear in the same order as the original string.
func ReverseSplit[S Any](s S, delimiter S) []S { //gd:String.rsplit
	split := slices.Collect(Splits(s, delimiter))
	for i, j := 0, len(split)-1; i < j; i, j = i+1, j-1 {
		split[i], split[j] = split[j], split[i]
	}
	return split
}

// StripSuffix removes a set of characters defined in chars from the string's end. See also [StripPrefix].
//
// Note: chars is a cutset. Use [TrimSuffix] to remove an entire suffix, rather than a set of characters.
func StripSuffix[S Any](s S, chars S) S { //gd:String.rstrip
	return As[S](strings.TrimRight(As[string](s), As[string](chars)))
}

// SHA1 returns the SHA-1 hash of the string.
func SHA1[S Any](s S) S { //gd:String.sha1_buffer String.sha1_text
	h := sha1.New()
	h.Write([]byte(As[string](s)))
	return As[S](hex.EncodeToString(h.Sum(nil)))
}

// SHA256 returns the SHA-256 hash of the string.
func SHA256[S Any](s S) S { //gd:String.sha256_buffer String.sha256_text
	h := sha256.New()
	h.Write([]byte(As[string](s)))
	return As[S](hex.EncodeToString(h.Sum(nil)))
}

// Similarity returns the similarity index (Sorensen-Dice coefficient) of this string compared to another.
// A result of 1.0 means totally similar, while 0.0 means totally dissimilar.
func Similarity[S Any](a, b S) Float.X { //gd:String.similarity
	if Comparison(a, b) == 0 {
		return 1.0 // Equal strings are totally similar
	}
	if Length(a) < 2 || Length(b) < 2 {
		return 0.0 // No way to calculate similarity without a single bigram
	}
	src_bigrams := slices.Collect(Bigrams(a))
	tgt_bigrams := slices.Collect(Bigrams(b))
	src_size := len(src_bigrams)
	tgt_size := len(tgt_bigrams)
	var sum = Float.X(src_size + tgt_size)
	var inter Float.X = 0
	for i := 0; i < src_size; i++ {
		for j := 0; j < tgt_size; j++ {
			if Comparison(src_bigrams[i], tgt_bigrams[j]) == 0 {
				inter++
				break
			}
		}
	}
	return (2.0 * inter) / sum
}

// SimplifyPath converts the string into a canonical path if the string is a valid file path. This
// is the shortest possible path, without "./", and all the unnecessary ".." and "/".
func SimplifyPath[S Any](s S) S { //gd:String.simplify_path
	return As[S](filepath.Clean(As[string](s)))
}

// Splits returns a sequence of substrings by splitting the string using a delimiter. If the delimiter
// is an empty string, each substring will be a single character. This method is the opposite of [JoinedWith].
func Splits[S Any, B Any](s S, sep B) iter.Seq[S] { //gd:String.split
	if Length(sep) == 0 {
		return func(yield func(S) bool) {
			for _, char := range As[string](s) {
				if !yield(As[S](New(string(char)))) {
					return
				}
			}
		}
	}
	return func(yield func(S) bool) {
		for {
			m := FindIndex(s, sep)
			if m < 0 {
				yield(s)
				break
			}
			if !yield(Slice(s, 0, m)) {
				return
			}
			s = Slice(s, m+Length(sep), Length(s))
		}
	}
}

// ExtractFloats splits the string into floats by using a delimiter and returns a [Float.X] slice.
func ExtractFloats[S Any, D Any](sep D, s S) []Float.X { //gd:String.split_floats
	var floats []Float.X
	for val := range Splits(s, sep) {
		f, _ := strconv.ParseFloat(As[string](val), 64)
		floats = append(floats, Float.X(f))
	}
	return floats
}

// Clean removes all leading and trailing whitespace from the string.
func Clean[S Any](s S) S { //gd:String.strip_edges
	return As[S](strings.TrimSpace(As[string](s)))
}

// StripEscapes strips all escape characters from the string. These include all
// non-printable control characters of the first page of the ASCII table (values
// from 0 to 31), such as tabulation (\t) and newline (\n, \r) characters, but not spaces.
func StripEscapes[S Any](s S) S { //gd:String.strip_escapes
	return As[S](regexp.MustCompile(`[\x00-\x1F]`).ReplaceAllString(As[string](s), ""))
}

// ToASCII converts the string to an ASCII/Latin-1 encoded byte slice. This method is slightly
// faster than [ToUTF8], but replaces all unsupported characters with spaces.
func ToASCII[S Any](s S) []byte { //gd:String.to_ascii_buffer
	var buf bytes.Buffer
	for _, r := range As[string](s) {
		if r > 255 {
			buf.WriteRune(' ')
		} else {
			buf.WriteByte(byte(r))
		}
	}
	return buf.Bytes()
}

// ASCII converts ASCII/Latin-1 encoded array to string. Fast alternative to get_string_from_utf8 if the
// content is ASCII/Latin-1 only. Unlike the UTF-8 function this function maps every byte to a character
// in the array. Multibyte sequences will not be interpreted correctly. For parsing user input always use
// [UTF8]. This is the inverse of [ToASCII].
func ASCII[S Any](s S) string { //gd:PackedByteArray.get_string_from_ascii
	return As[string](s)
}

// ToCamelCase returns the string converted to camelCase.
func ToCamelCase[S Any](s S) S { //gd:String.to_camel_case
	return As[S](strings.ReplaceAll(strings.Title(strings.ToLower(As[string](s))), " ", ""))
}

// ToFloat converts the string representing a decimal number into a float. This method stops
// on the first non-number character, except the first decimal point (.) and the exponent
// letter (e). See also [IsValidFloat].
func ToFloat[S Any](s S) Float.X { //gd:String.to_float
	l := Length(s)
	for i := 0; i < l; i++ {
		chr := Index(s, i)
		if !unicode.IsDigit(rune(chr)) && chr != '.' && chr != 'e' && chr != 'E' && chr != '-' {
			if i == 0 {
				return 0
			}
			l = i
			break
		}
	}
	if Count(".", Slice(s, 0, l)) > 1 {
		s = Slice(s, 0, FindLast(Slice(s, 0, l), "."))
		l = Length(s)
	}
	f, _ := strconv.ParseFloat(As[string](Slice(s, 0, l)), 64)
	return Float.X(f)
}

// ToInt converts the string representing an integer number into an int. This method removes any
// non-number character and stops at the first decimal point (.). See also [IsValidInt].
func ToInt[S Any](s S) int { //gd:String.to_int
	n := ""
	for i := 0; i < Length(s); i++ {
		if unicode.IsDigit(rune(Index(s, i))) || Index(s, i) == '-' {
			n += string(Index(s, i))
		}
		if Index(s, i) == '.' {
			break
		}
	}
	i, _ := strconv.Atoi(string(n))
	return i
}

// ToLower converts the string to lowercase.
func ToLower[S Any](s S) S { //gd:String.to_lower
	return As[S](strings.ToLower(As[string](s)))
}

// ToPascalCase returns the string converted to PascalCase.
func ToPascalCase[S Any](s S) S { //gd:String.to_pascal_case
	return As[S](strings.ReplaceAll(Title(strings.ToLower(As[string](s))), " ", ""))
}

// ToSnakeCase Returns the string converted to snake_case.
//
// Note: Numbers followed by a single letter are not separated in the conversion to keep some words (such as "2D") together.
func ToSnakeCase[S Any](s S) S { //gd:String.to_snake_case
	var new_string S
	var start_index = 0
	for i := 1; i < Length(s); i++ {
		var is_prev_upper = unicode.IsUpper(rune(Index(s, i-1)))
		var is_prev_lower = unicode.IsLower(rune(Index(s, i-1)))
		var is_prev_digit = unicode.IsDigit(rune(Index(s, i-1)))
		var is_curr_upper = unicode.IsUpper(rune(Index(s, i)))
		var is_curr_lower = unicode.IsLower(rune(Index(s, i)))
		var is_curr_digit = unicode.IsDigit(rune(Index(s, i)))
		var is_next_lower = false
		if i+1 < Length(s) {
			is_next_lower = unicode.IsLower(rune(Index(s, i+1)))
		}
		var cond_a = is_prev_lower && is_curr_upper                                     // aA
		var cond_b = (is_prev_upper || is_prev_digit) && is_curr_upper && is_next_lower // AAa, 2Aa
		var cond_c = is_prev_digit && is_curr_lower && is_next_lower                    // 2aa
		var cond_d = (is_prev_upper || is_prev_lower) && is_curr_digit                  // A2, a2
		if cond_a || cond_b || cond_c || cond_d {
			new_string = JoinedWith("", new_string, Slice(s, start_index, i), As[S](New("_")))
			start_index = i
		}
	}
	new_string = JoinedWith("", new_string, Slice(s, start_index, Length(s)))
	return As[S](strings.Replace(strings.ToLower(As[string](new_string)), " ", "_", -1))
}

// ToUpper returns the string converted to UPPERCASE.
func ToUpper[S Any](s S) S { //gd:String.to_upper
	return As[S](strings.ToUpper(As[string](s)))
}

// ToUTF8 converts the string to a UTF-8 encoded byte slice. This method is slightly slower than
// [ToASCII], but supports all UTF-8 characters. For most cases, prefer using this method.
func ToUTF8[S Any](s S) []byte { //gd:String.to_utf8_buffer
	return As[[]byte](s)
}

// UTF8 converts UTF-8 encoded array to string. Slower than [ASCII] but supports UTF-8 encoded
// data. Use this function if you are unsure about the source of the data. For user input this function should
// always be preferred. Returns empty string if source array is not valid UTF-8 string. This is the inverse
// of [ToUTF8].
func UTF8[S Any](s []byte) string { //gd:PackedByteArray.get_string_from_utf8
	return string(s)
}

// ToUTF16 converts the string to a UTF-16.
func ToUTF16[S Any](s S) []byte { //gd:String.to_utf16_buffer String.to_wchar_buffer
	var result bytes.Buffer
	for _, r := range Runes(s) {
		a, b := utf16.EncodeRune(rune(r))
		result.WriteByte(byte(a))
		result.WriteByte(byte(b))
	}
	return result.Bytes()
}

// UTF16 converts UTF-16 encoded array to String. If the BOM is missing, system endianness
// is assumed. Returns empty string if source array is not valid UTF-16 string. This is
// the inverse of [ToUTF16].
func UTF16[S Any](s []byte) string { //gd:PackedByteArray.get_string_from_utf16 PackedByteArray.get_string_from_wchar
	if len(s)%2 != 0 {
		return ""
	}
	return string(utf16.Decode(unsafe.Slice((*uint16)(unsafe.Pointer(&s[0])), len(s)/2)))
}

// ToUTF32 converts the string to a UTF-32.
func ToUTF32[S Any](s S) []byte { //gd:String.to_utf32_buffer
	var UTF = utf32.UTF32(utf32.LittleEndian, utf32.UseBOM).NewEncoder()
	b, _ := UTF.Bytes(As[[]byte](s))
	return b
}

// UTF32 converts UTF-32 encoded array to string. If the BOM is missing, system endianness
// is assumed. Returns empty string if source array is not valid UTF-32 string. This is
// the inverse of [ToUTF32].
func UTF32[S Any](s S) string { //gd:PackedByteArray.get_string_from_utf32
	var UTF = utf32.UTF32(utf32.LittleEndian, utf32.UseBOM).NewDecoder()
	b, _ := UTF.String(As[string](s))
	return b
}

// TrimPrefix removes the given prefix from the start of the string, or returns the string unchanged.
func TrimPrefix[S, P Any](s S, prefix P) S { //gd:String.trim_prefix
	if HasPrefix(s, prefix) {
		return Slice(s, Length(prefix), Length(s))
	}
	return s
}

// TrimSuffix removes the given suffix from the end of the string, or returns the string unchanged.
func TrimSuffix[S, P Any](s S, suffix P) S { //gd:String.trim_suffix
	return As[S](strings.TrimSuffix(As[string](s), As[string](suffix)))
}

// UnicodeAt returns the Unicode code point at the given index in the string.
func UnicodeAt[S Any](s S, index int) Rune { //gd:String.unicode_at
	for i, r := range Runes(s) {
		if i == index {
			return r
		}
	}
	return 0
}

// StripFilename returns a copy of the string with all characters that are not allowed in [IsValidFilename] replaced with underscores.
func StripFilename[S Any](s S) S { //gd:String.validate_filename
	var result = builder()
	for _, r := range Runes(s) {
		if IsValidFilename(string(r)) {
			result = Append(result, New(r))
		} else {
			result = Append(result, New(Rune('_')))
		}
	}
	return As[S](result)
}

// StripNodeName returns a copy of the string with all characters that are not allowed in Node.name (. : @ / " %) replaced with underscores.
func StripNodeName[S Any](s S) S { //gd:String.validate_node_name
	return S(Map(func(r Rune) Rune {
		switch r {
		case '.', ':', '@', '/', '"', '%':
			return '_'
		}
		return r
	}, s))
}

// StartingFrom returns a slice of the string from the given start index to the end index.
func StartingFrom[S Any](s S, start int) S { //gd:String.substr
	return Slice(s, start, Length(s))
}

// Slice returns a slice of the string from the start index to the end index.
func Slice[S Any, A, B Int.Any](s S, start A, close B) S { //gd:String.substrn
	uni := New(s)
	return As[S](uni.api.Slice(uni.buf, int(start), int(close)))
}

// EncodeURI encodes the string to URL-friendly format. This method is meant to properly
// encode the parameters in a URL when sending an HTTP request. See also [DecodeURI].
func EncodeURI[S Any](s S) S { //gd:String.uri_encode
	return As[S](url.QueryEscape(As[string](s)))
}

// DecodeURI decodes the string from URL-friendly format. This method is meant to properly
// decode the parameters in a URL when receiving an HTTP request. See also [EncodeURI].
func DecodeURI[S Any](s S) S { //gd:String.uri_decode
	decoded, _ := url.QueryUnescape(As[string](s))
	return As[S](decoded)
}

// isSeparator reports whether the rune could mark a word boundary.
// TODO: update when package unicode captures more of the properties.
func isSeparator(r rune) bool {
	// ASCII alphanumerics and underscore are not separators
	if r <= 0x7F {
		switch {
		case '0' <= r && r <= '9':
			return false
		case 'a' <= r && r <= 'z':
			return false
		case 'A' <= r && r <= 'Z':
			return false
		case r == '_':
			return false
		}
		return true
	}
	// Letters and digits are not separators
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return false
	}
	// Otherwise, all we can do for now is treat spaces as separators.
	return unicode.IsSpace(r)
}
