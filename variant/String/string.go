// Package String provides a generic string functions.
package String

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/netip"
	"net/url"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
	"unsafe"

	"golang.org/x/text/encoding/unicode/utf32"
	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Int"
)

type Any interface {
	~string | ~[]byte
}

var static = make(map[string]gd.String)

func init() {
	gd.StartupFunctions = append(gd.StartupFunctions, func() {
		for s, v := range static {
			str := gd.Global.Strings.New(s)
			raw, ok := pointers.End(str)
			if ok {
				pointers.Set(v, raw)
			}
		}
	})
}

// New returns an engine-optimised String for use with Advanced functions.
// Can be initialised ahead of time as a global variable.
func New(vals ...any) gd.String { //gd:String() str
	if !gd.Linked {
		str := pointers.Add[gd.String]([1]uintptr{})
		static[fmt.Sprint(vals...)] = str
		return str
	}
	return gd.Global.Strings.New(fmt.Sprint(vals...))
}

// BeginsWith returns true if the string begins with the given text. See also [EndsWith].
func BeginsWith[S Any](prefix S, s S) bool { //gd:String.begins_with
	return strings.HasPrefix(string(s), string(prefix))
}

// Bigrams returns a slice containing the bigrams (pairs of consecutive characters) of this string.
func Bigrams[S Any](s S) []S { //gd:String.bigrams
	if len(s) < 2 {
		return nil
	}
	var r = make([]S, len(s)-1)
	for i := 0; i < len(s)-1; i++ {
		r[i] = s[i : i+2]
	}
	return r
}

// BinaryToInteger converts the string representing a binary number into an int. The string may
// optionally be prefixed with "0b", and an additional - prefix for negative numbers.
func BinaryToInteger[S Any](s S) int { //gd:String.bin_to_int
	s = S(strings.TrimPrefix(string(s), "0b"))
	var r int
	for i := 0; i < len(s); i++ {
		if s[i] == '0' || s[i] == '1' {
			r = r*2 + int(s[i]-'0')
		}
	}
	if strings.HasPrefix(string(s), "-") {
		r = -r
	}
	return r
}

// Escape returns a copy of the string with special characters escaped using the Go language standard.
func Escape[S Any](s S) S { //gd:String.c_escape String.json_escape String.xml_escape
	escaped := strconv.QuoteToASCII(string(s))
	return S(escaped[1 : len(escaped)-1])
}

// Unescape returns a copy of the string with escaped characters replaced by their meanings.
func Unescape[S Any](s S) S { //gd:String.c_unescape String.xml_unescape
	unescaped, _ := strconv.Unquote(`"` + string(s) + `"`)
	return S(unescaped)
}

// Capitalize changes the appearance of the string: replaces underscores (_) with spaces, adds spaces
// before uppercase letters in the middle of a word, converts all letters to lowercase, then converts
// the first one and each one following a space to uppercase.
func Capitalize[S Any](s S) S { //gd:String.capitalize
	pieces := strings.Split(string(s), "_")
	for i, piece := range pieces {
		for i, r := range piece {
			if unicode.IsUpper(r) {
				piece = piece[:i] + " " + piece[i:]
			}
		}
		pieces[i] = strings.Title(strings.ToLower(piece))
	}
	return S(strings.Join(pieces, " "))
}

// Comparison performs a case-insensitive comparison to another string. Returns -1
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
func Comparison[S Any](a, b S) int { //gd:String.nocasecmp_to
	return strings.Compare(strings.ToUpper(string(a)), strings.ToUpper(string(b)))
}

// ComparisonStrict performs a case-sensitive comparison to another string. Returns -1 if less than,
// 1 if greater than, or 0 if equal. "Less than" and "greater than" are determined by the Unicode code
// points of each string, which roughly matches the alphabetical order.
//
// With different string lengths, returns 1 if this string is longer than the to string, or -1 if shorter.
// Note that the length of empty strings is always 0.
//
// To get a bool result from a string comparison, use the == operator instead. See also [Comparison],
// [ComparisonStrictNaturalPrioritizePeriodsAndUnderscores], and [ComparisonStrictNatural].
func ComparisonStrict[S Any](a, b S) int { //gd:String.casecmp_to
	return strings.Compare(string(a), string(b))
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
		m := len(a)
		if n := len(b); n < m {
			m = n
		}
		if m == 0 {
			return 0
		}
		_ = a[m-1]
		_ = b[m-1]
		for i := 0; i < m; i++ {
			ca := a[i]
			cb := b[i]
			if (ca >= '0' && ca <= '9') || (cb >= '0' && cb <= '9') || ca != cb {
				return i
			}
		}
		return m
	}
	digits := func(s S) int {
		for i := 0; i < len(s); i++ {
			c := s[i]
			if c < '0' || c > '9' {
				return i
			}
		}
		return len(s)
	}
	for {
		if p := commonPrefix(a, b); p != 0 {
			a = a[p:]
			b = b[p:]
		}
		if len(a) == 0 {
			if len(b) != 0 {
				return -1
			}
			return 1
		}
		if ia := digits(a); ia > 0 {
			if ib := digits(b); ib > 0 {
				// Both sides have digits.
				an, aerr := strconv.ParseUint(string(a[:ia]), 10, 64)
				bn, berr := strconv.ParseUint(string(b[:ib]), 10, 64)
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
					if ia != len(a) && ib != len(b) {
						a = a[ia:]
						b = b[ib:]
						continue
					}
				}
			}
		}
		return bytes.Compare([]byte(a), []byte(b))
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
	return ComparisonStrictNatural(strings.ToUpper(string(a)), strings.ToUpper(string(b)))
}

// ComparisonStrictNaturalPrioritizePeriodsAndUnderscores like [ComparisonStrictNatural] but prioritizes strings that begin with periods
// (.) and underscores (_) before any other character. Useful when sorting folders or file names.
//
// To get a bool result from a string comparison, use the == operator instead. See also [ComparisonNaturalPrioritizePeriodsAndUnderscores],
// [ComparisonStrictNatural], and [ComparisonStrict].
func ComparisonStrictNaturalPrioritizePeriodsAndUnderscores[S Any](a, b S) int { //gd:String.filecasecmp_to
	aStr := string(a)
	bStr := string(b)
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
	aStr := string(a)
	bStr := string(b)
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
	return strings.Contains(string(s), string(what))
}

// Contains returns true if the string contains what, ignoring case.
// If you need to know where what is within the string, use findn. See also [Contains].
func Contains[S Any](what S, s S) bool { //gd:String.containsn
	return strings.Contains(strings.ToLower(string(s)), strings.ToLower(string(what)))
}

// CountStrict returns the number of occurrences of the case-sensitive substring.
func CountStrict[S Any](what S, s S) int { //gd:String.count
	return strings.Count(string(s), string(what))
}

// Count returns the number of occurrences of the substring, ignoring case.
func Count[D ~string | ~[]byte | rune, S Any](what D, s S) int { //gd:String.countn
	return strings.Count(strings.ToLower(string(s)), strings.ToLower(string(what)))
}

// Dedent returns a copy of the string with indentation (leading tabs and spaces)
// removed. See also [Indent] to add indentation.
func Dedent[S Any](s S) S { //gd:String.dedent
	return S(strings.TrimLeftFunc(string(s), unicode.IsSpace))
}

// EndsWith returns true if the string ends with the given text. See also [BeginsWith].
func EndsWith[S Any](suffix S, s S) bool { //gd:String.ends_with
	return strings.HasSuffix(string(s), string(suffix))
}

// Erase returns a string with the character erased at position.
func Erase[S Any, I Int.Any](position I, s S) S { //gd:String.erase
	return S(string(s[:position]) + string(s[position+1:]))
}

// FindStrict returns the index of the first case-sensitive occurrence of what in this string
// or -1 if not found.
func FindStrict[S Any](what S, s S) int { //gd:String.find
	return strings.Index(string(s), string(what))
}

// Find returns the index of the first occurrence of what in this string, ignoring case.
// Returns -1 if not found.
func Find[S Any](what S, s S) int { //gd:String.findn
	return strings.Index(strings.ToLower(string(s)), strings.ToLower(string(what)))
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
	for i := 0; i < len(format); i++ {
		v := format[i]
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
	return S(result.String())
}

// Directory returns the base directory name if the string is a valid file path.
func Directory[S Any](path S) S { //gd:String.get_base_dir
	return S(filepath.Dir(string(path)))
}

// FileExtension returns the file extension without the leading period (.) if the string is a
// valid file name or path. Otherwise, returns an empty string.
func FileExtension[S Any](path S) S { //gd:String.get_extension
	ext := filepath.Ext(string(path))
	if ext != "" {
		return S(ext[1:])
	}
	return S("")
}

// FileName returns the file name, including the extension, if the string is a valid file path,
func FileName[S Any](path S) S { //gd:String.get_file
	return S(filepath.Base(string(path)))
}

// Extract splits the string using a delimiter and returns the substring at index. Returns
// the original string if delimiter does not occur in the string. Returns an empty string if the
// element does not exist.
//
// This is faster than split, if you only need one substring.
func Extract[S Any, D ~string | ~[]byte | rune, I Int.Any](s S, delimiter D, index I) S { //gd:String.get_slice String.get_slicec
	var n int
	var current S
	for i := 0; i < int(index+1); i++ {
		next := strings.Index(string(s[n:]), string(delimiter))
		if next == -1 {
			return s[n:]
		}
		current = s[n : n+next]
		n += next + 1
	}
	return current
}

// NumExtractions returns the total number of extractable substrings when the string is split with
// the given delimiter (see [Split]).
func NumExtractions[S Any, D ~string | ~[]byte | rune](s S, delimiter D) int { //gd:String.get_slice_count
	return strings.Count(string(s), string(delimiter)) + 1
}

// Hash Returns the 32-bit hash value representing the string's contents.
//
// Note: Strings with equal hash values are not guaranteed to be the same, as a result of hash collisions.
// On the contrary, strings with different hash values are guaranteed to be different.
func Hash[S Any](s S) uint32 { //gd:String.hash
	var hashv uint32 = 5381
	for i := 0; i < len(s); i++ {
		hashv = ((hashv << 5) + hashv) + uint32(s[i])
	}
	return hashv
}

// DecodeHex decodes a hexadecimal string as a []byte.
func DecodeHex[S Any](s S) S { //gd:String.hex_decode
	b, err := hex.DecodeString(string(s))
	if err != nil {
		var zero S
		return zero
	}
	return S(b)
}

// EncodeHex encodes a []byte as a hexadecimal string.
func EncodeHex[S Any](s S) S { //gd:String.hex_encode
	return S(hex.EncodeToString([]byte(s)))
}

// HexToInteger Converts the string representing a hexadecimal number into an int. The
// string may be optionally prefixed with "0x", and an additional - prefix for negative numbers.
func HexToInteger[S Any](s S) int { //gd:String.hex_to_int
	if strings.HasPrefix(string(s), "0x") {
		s = s[2:]
	}
	i, err := strconv.ParseInt(string(s), 16, 0)
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
	return S(strings.ReplaceAll(string(s), "\n", "\n"+string(prefix)))
}

// Insert inserts what at the given position in the string.
func Insert[S Any](what, s S, position int) S { //gd:String.insert
	return S(string(s[:position]) + string(what) + string(s[position:]))
}

// IsEmpty returns true if the string's length is 0 (""). See also length.
func IsEmpty[S Any](s S) bool { //gd:String.is_empty
	return len(s) == 0
}

// IsStrictSubsequenceOf returns true if all characters of this string can be found in text in their original order.
func IsStrictSubsequenceOf[S Any](text, s S) bool { //gd:String.is_subsequence_of
	if len(s) == 0 {
		return true
	}
	n := 0
	for i := 0; i < len(s); i++ {
		for n < len(text) && text[n] != s[i] {
			n++
		}
		if n == len(text) {
			return false
		}
		n++
	}
	return true
}

// IsSubsequenceOf returns true if all characters of this string can be found in text in any order, ignoring case.
func IsSubsequenceOf[S Any](text, s S) bool { //gd:String.is_subsequence_ofn
	if len(s) == 0 {
		return true
	}
	n := 0
	for i := 0; i < len(s); i++ {
		for n < len(text) && unicode.ToLower(rune(text[n])) != unicode.ToLower(rune(s[i])) {
			n++
		}
		if n == len(text) {
			return false
		}
		n++
	}
	return true
}

// IsValidFilename returns true if this string does not contain characters that are not allowed
// in file names (: / \ ? * " | % < >).
func IsValidFilename[S Any](s S) bool { //gd:String.is_valid_filename
	return !strings.ContainsAny(string(s), ":/\\?*\"|%<>")
}

// IsValidFloat returns true if this string represents a valid floating-point number. A valid float
// may contain only digits, one decimal point (.), and the exponent letter (e). It may also be
// prefixed with a positive (+) or negative (-) sign. Any valid integer is also a valid float
// (see [IsValidInt]). See also [ToFloat].
func IsValidFloat[S Any](s S) bool { //gd:String.is_valid_float
	_, err := strconv.ParseFloat(string(s), 64)
	return err == nil
}

// IsValidHexNumber returns true if this string is a valid hexadecimal number. A valid hexadecimal
// number only contains digits or letters A to F (either uppercase or lowercase), and may be prefixed
// with a positive (+) or negative (-) sign.
func IsValidHexNumber[S Any](s S) bool { //gd:String.is_valid_hex_number
	_, err := strconv.ParseInt(string(s), 16, 0)
	return err == nil
}

// IsValidHexColor Returns true if this string is a valid color in hexadecimal HTML notation. The string
// must be a hexadecimal value (see is_valid_hex_number) of either 3, 4, 6 or 8 digits, and may be prefixed
// by a hash sign (#). Other HTML notations for colors, such as names or hsl(), are not considered valid.
func IsValidHexColor[S Any](s S) bool { //gd:String.is_valid_html_color
	if strings.HasPrefix(string(s), "#") {
		s = s[1:]
	}
	switch len(s) {
	case 3, 4, 6, 8:
		return IsValidHexNumber(s)
	default:
		return false
	}
}

// IsValidIdentifier returns true if this string is a valid identifier. A valid identifier may contain only letters,
// digits and underscores (_), and the first character may not be a digit.
func IsValidIdentifier[S Any](s S) bool { //gd:String.is_valid_identifier
	if len(s) == 0 {
		return false
	}
	if unicode.IsDigit(rune(s[0])) {
		return false
	}
	for _, c := range string(s) {
		if !unicode.IsDigit(c) && !unicode.IsLetter(c) && c != '_' {
			return false
		}
	}
	return true
}

// IsValidInt returns true if this string represents a valid integer. A valid integer only contains
// digits, and may be prefixed with a positive (+) or negative (-) sign. See also [ToInt].
func IsValidInt[S Any](s S) bool { //gd:String.is_valid_int
	_, err := strconv.ParseInt(string(s), 10, 0)
	return err == nil
}

// IsValidInternetProtocolAddress returns true if this string represents a well-formatted IPv4 or IPv6 address.
// This method considers reserved IP addresses such as "0.0.0.0" and "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff" as valid.
func IsValidInternetProtocolAddress[S Any](s S) bool { //gd:String.is_valid_ip_address
	_, err := netip.ParseAddr(string(s))
	return err == nil
}

// JoinedWith returns the concatenation of parts' elements, with each element separated by the string calling this method.
// This method is the opposite of [Split].
func JoinedWith[S Any, D ~string | ~[]byte | rune](d D, s ...S) S { //gd:String.join
	if len(s) == 0 {
		return S("")
	}
	var b strings.Builder
	b.WriteString(fmt.Sprint(s[0]))
	for _, v := range s[1:] {
		b.WriteString(fmt.Sprint(d))
		b.WriteString(fmt.Sprint(v))
	}
	return S(b.String())
}

// First, returns the first n characters from the beginning of the string. If length is
// negative, strips the last length characters from the string's end.
func First[S Any, I Int.Any](n I, s S) S { //gd:String.left
	if n < 0 {
		return s[:len(s)+int(n)]
	}
	return s[:int(n)]
}

// Length returns the number of characters in the string.
func Length[S Any](s S) int { //gd:String.length
	return (len(string(s)))
}

// PadPrefix formats the string to be at least min_length long by adding characters
// to the left of the string, if necessary. See also [PadSuffix].
func PadPrefix[S Any](s S, min_length int) S { //gd:String.lpad
	return JoinedWith("", S(strings.Repeat(" ", max(0, min_length-len(s)))), s)
}

// StripPrefix removes a set of characters defined in chars from the string's beginning. See also StripSuffix.
//
// Note: chars is a cutset. Use [TrimPrefix] to remove a whole prefix string, rather than a set of characters.
func StripPrefix[S Any, C ~string | ~[]byte](s S, chars C) S { //gd:String.lstrip
	return S(strings.TrimLeft(string(s), string(chars)))
}

// MatchStrict does a simple expression match (also called "glob" or "globbing"), where * matches zero or more arbitrary
// characters and ? matches any single character except a period (.). An empty string or empty expression always
// evaluates to false.
func MatchStrict[S Any](pattern, text S) bool { //gd:String.match
	ok, _ := path.Match(string(pattern), string(text))
	return ok
}

// Match does a simple case-insensitive expression match, where * matches zero or more arbitrary characters and ? matches
// any single character except a period (.). An empty string or empty expression always evaluates to false.Does a simple
// case-insensitive expression match, where * matches zero or more arbitrary characters and ? matches any single character
// except a period (.). An empty string or empty expression always evaluates to false.
func Match[S Any](pattern, text S) bool { //gd:String.matchn
	ok, _ := path.Match(strings.ToLower(string(pattern)), strings.ToLower(string(text)))
	return ok
}

// MD5 returns the MD5 hash of the string as a byte slice.
func MD5[S Any](s S) []byte { //gd:String.md5_buffer String.md5_text
	sum := md5.Sum([]byte(string(s)))
	return sum[:]
}

// PadDecimals formats the string representing a number to have an exact number of digits
// after the decimal point.
func PadDecimals[S Any](s S, digits int) S { //gd:String.pad_decimals
	if !IsValidFloat(s) {
		return s
	}
	var result strings.Builder
	result.WriteString(string(s))
	result.WriteString(strings.Repeat("0", len(s)-strings.Index(string(s), ".")+digits+1))
	return S(result.String())
}

// PadZeros formats the string representing a number to have an exact number of digits before
// the decimal point.
func PadZeros[S Any](s S, digits int) S { //gd:String.pad_zeros
	if !IsValidFloat(s) {
		return s
	}
	var result strings.Builder
	result.WriteString(strings.Repeat("0", strings.Index(string(s), ".")-digits))
	result.WriteString(string(s))
	return S(result.String())
}

// AddPathElement concatenates file at the end of the string as a subpath, adding / if necessary.
func AddPathElement[S Any](s S, file S) S { //gd:String.path_join
	if strings.HasSuffix(string(s), "/") {
		return JoinedWith("", s, file)
	}
	return JoinedWith("/", s, file)
}

// Repeat repeats this string a number of times. count needs to be greater than 0.
// Otherwise, returns an empty string.
func Repeat[S Any](s S, count int) S { //gd:String.repeat
	if count <= 0 {
		return S("")
	}
	return S(strings.Repeat(string(s), count))
}

// ReplaceStrict replaces all occurrences of what inside the string with the given forwhat.
func ReplaceStrict[S Any](s S, what, forwhat S) S { //gd:String.replace
	return S(strings.ReplaceAll(string(s), string(what), string(forwhat)))
}

// Replace replaces all case-insensitive occurrences of what inside the string with the given
// forwhat.
func Replace[S Any](s S, old, new S) S { //gd:String.replacen
	var n = -1
	if string(old) == string(new) || n == 0 {
		return s // avoid allocation
	}

	// Compute number of replacements.
	if m := Count(s, old); m == 0 {
		return s // avoid allocation
	} else if n < 0 || m < n {
		n = m
	}

	// Apply replacements to buffer.
	var b strings.Builder
	b.Grow(len(s) + n*(len(new)-len(old)))
	start := 0
	for i := 0; i < n; i++ {
		j := start
		if len(old) == 0 {
			if i > 0 {
				_, wid := utf8.DecodeRuneInString(string(s[start:]))
				j += wid
			}
		} else {
			j += strings.Index(string(s[start:]), string(old))
		}
		b.WriteString(string(s[start:j]))
		b.WriteString(string(new))
		start = j + len(old)
	}
	b.WriteString(string(s[start:]))
	return S(b.String())
}

// Reverse returns the copy of this string in reverse order. This operation works on
// unicode codepoints, rather than sequences of codepoints, and may break things like
// compound letters or emojis.
func Reverse[S Any](s S) S { //gd:String.reverse
	runes := []rune(string(s))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return S(string(runes))
}

// FindLastStrict returns the index of the last occurrence of what in this string, or -1 if there are none.
// This method is the reverse of [FindStrict].
func FindLastStrict[S Any](s S, what S) int { //gd:String.rfind
	return strings.LastIndex(string(s), string(what))
}

// FindLast returns the index of the last occurrence of what in this string, or -1 if there are none.
// This method is the reverse of [Find].
func FindLast[S Any](s S, what S) int { //gd:String.rfindn
	return strings.LastIndex(strings.ToLower(string(s)), strings.ToLower(string(what)))
}

// Last returns the last length characters from the end of the string. If length is
// negative, returns the first length characters from the string's beginning.
func Last[S Any](length int, s S) S { //gd:String.right
	if length < 0 {
		return S(string(s)[-length:])
	}
	return S(string(s)[len(s)-length:])
}

// PadSuffix formats the string to be at least min_length long, by adding characters to the right of the
// string, if necessary. See also [PadPrefix].
func PadSuffix[S Any](s S, min_length int) S { //gd:String.rpad
	return S(string(s) + strings.Repeat(string(" "), max(0, min_length-len(s))))
}

// ReverseSplit splits the string using a delimiter and returns an array of the substrings, starting from
// the end of the string. The splits in the returned array appear in the same order as the original string.
func ReverseSplit[S Any](s S, delimiter S) []S { //gd:String.rsplit
	split := ExtractAll(delimiter, s)
	for i, j := 0, len(split)-1; i < j; i, j = i+1, j-1 {
		split[i], split[j] = split[j], split[i]
	}
	return split
}

// StripSuffix removes a set of characters defined in chars from the string's end. See also [StripPrefix].
//
// Note: chars is a cutset. Use [TrimSuffix] to remove an entire suffix, rather than a set of characters.
func StripSuffix[S Any](s S, chars S) S { //gd:String.rstrip
	return S(strings.TrimRight(string(s), string(chars)))
}

// SHA1 returns the SHA-1 hash of the string.
func SHA1[S Any](s S) S { //gd:String.sha1_buffer String.sha1_text
	h := sha1.New()
	h.Write([]byte(string(s)))
	return S(hex.EncodeToString(h.Sum(nil)))
}

// SHA256 returns the SHA-256 hash of the string.
func SHA256[S Any](s S) S { //gd:String.sha256_buffer String.sha256_text
	h := sha256.New()
	h.Write([]byte(string(s)))
	return S(hex.EncodeToString(h.Sum(nil)))
}

// Similarity returns the similarity index (Sorensen-Dice coefficient) of this string compared to another.
// A result of 1.0 means totally similar, while 0.0 means totally dissimilar.
func Similarity[S Any](a, b S) Float.X { //gd:String.similarity
	if bytes.Equal([]byte(a), []byte(b)) {
		return 1.0 // Equal strings are totally similar
	}
	if len(a) < 2 || len(b) < 2 {
		return 0.0 // No way to calculate similarity without a single bigram
	}
	src_bigrams := Bigrams(a)
	tgt_bigrams := Bigrams(b)
	src_size := len(src_bigrams)
	tgt_size := len(tgt_bigrams)
	var sum = Float.X(src_size + tgt_size)
	var inter Float.X = 0
	for i := 0; i < src_size; i++ {
		for j := 0; j < tgt_size; j++ {
			if bytes.Equal([]byte(src_bigrams[i]), []byte(tgt_bigrams[j])) {
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
	return S(filepath.Clean(string(s)))
}

// ExtractAll splits the string using a delimiter and returns an array of the substrings. If delimiter
// is an empty string, each substring will be a single character. This method is the opposite of [JoinedWith].
func ExtractAll[S Any, D ~string | ~[]byte | rune](sep D, s S) []S { //gd:String.split
	explode := func(s S, n int) []S {
		l := utf8.RuneCountInString(string(s))
		if n < 0 || n > l {
			n = l
		}
		a := make([]S, n)
		for i := 0; i < n-1; i++ {
			_, size := utf8.DecodeRuneInString(string(s))
			a[i] = s[:size]
			s = s[size:]
		}
		if n > 0 {
			a[n-1] = s
		}
		return a
	}
	var n = -1
	if n == 0 {
		return nil
	}
	if len(string(sep)) == 0 {
		return explode(s, n)
	}
	if n < 0 {
		n = Count(sep, s) + 1
	}

	if n > len(s)+1 {
		n = len(s) + 1
	}
	a := make([]S, n)
	n--
	i := 0
	for i < n {
		m := strings.Index(string(s), string(sep))
		if m < 0 {
			break
		}
		a[i] = s[:m]
		s = s[m+len(string(sep)):]
		i++
	}
	a[i] = s
	return a[:i+1]
}

// ExtractFloats splits the string into floats by using a delimiter and returns a [Float.X] slice.
func ExtractFloats[S Any, D ~string | ~[]byte | rune](sep D, s S) []Float.X { //gd:String.split_floats
	var floats []Float.X
	for _, val := range ExtractAll(sep, s) {
		f, _ := strconv.ParseFloat(string(val), 64)
		floats = append(floats, Float.X(f))
	}
	return floats
}

// Clean removes all leading and trailing whitespace from the string.
func Clean[S Any](s S) S { //gd:String.strip_edges
	return S(strings.TrimSpace(string(s)))
}

// StripEscapes strips all escape characters from the string. These include all
// non-printable control characters of the first page of the ASCII table (values
// from 0 to 31), such as tabulation (\t) and newline (\n, \r) characters, but not spaces.
func StripEscapes[S Any](s S) S { //gd:String.strip_escapes
	return S(regexp.MustCompile(`[\x00-\x1F]`).ReplaceAllString(string(s), ""))
}

// ToASCII converts the string to an ASCII/Latin-1 encoded byte slice. This method is slightly
// faster than [ToUTF8], but replaces all unsupported characters with spaces.
func ToASCII[S Any](s S) []byte { //gd:String.to_ascii_buffer
	var buf bytes.Buffer
	for _, r := range string(s) {
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
	return string(s)
}

// ToCamelCase returns the string converted to camelCase.
func ToCamelCase[S Any](s S) S { //gd:String.to_camel_case
	return S(strings.ReplaceAll(strings.Title(strings.ToLower(string(s))), " ", ""))
}

// ToFloat converts the string representing a decimal number into a float. This method stops
// on the first non-number character, except the first decimal point (.) and the exponent
// letter (e). See also [IsValidFloat].
func ToFloat[S Any](s S) Float.X { //gd:String.to_float
	l := len(string(s))
	for i := 0; i < len(s); i++ {
		chr := s[i]
		if !unicode.IsDigit(rune(chr)) && chr != '.' && chr != 'e' && chr != 'E' && chr != '-' {
			if i == 0 {
				return 0
			}
			l = i
			break
		}
	}
	if strings.Count(string(s[:l]), ".") > 1 {
		s = s[:strings.LastIndex(string(s[:l]), ".")]
		l = len(string(s))
	}
	f, _ := strconv.ParseFloat(string(s[:l]), 64)
	return Float.X(f)
}

// ToInt converts the string representing an integer number into an int. This method removes any
// non-number character and stops at the first decimal point (.). See also [IsValidInt].
func ToInt[S Any](s S) int { //gd:String.to_int
	n := ""
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) || s[i] == '-' {
			n += string(s[i])
		}
		if s[i] == '.' {
			break
		}
	}
	i, _ := strconv.Atoi(string(n))
	return i
}

// ToLower converts the string to lowercase.
func ToLower[S Any](s S) S { //gd:String.to_lower
	return S(strings.ToLower(string(s)))
}

// ToPascalCase returns the string converted to PascalCase.
func ToPascalCase[S Any](s S) S { //gd:String.to_pascal_case
	return S(strings.ReplaceAll(strings.Title(strings.ToLower(string(s))), " ", ""))
}

// ToSnakeCase Returns the string converted to snake_case.
//
// Note: Numbers followed by a single letter are not separated in the conversion to keep some words (such as "2D") together.
func ToSnakeCase[S Any](s S) S { //gd:String.to_snake_case
	var new_string S
	var start_index = 0
	for i := 1; i < len(s); i++ {
		var is_prev_upper = unicode.IsUpper(rune(s[i-1]))
		var is_prev_lower = unicode.IsLower(rune(s[i-1]))
		var is_prev_digit = unicode.IsDigit(rune(s[i-1]))
		var is_curr_upper = unicode.IsUpper(rune(s[i]))
		var is_curr_lower = unicode.IsLower(rune(s[i]))
		var is_curr_digit = unicode.IsDigit(rune(s[i]))
		var is_next_lower = false
		if i+1 < len(s) {
			is_next_lower = unicode.IsLower(rune(s[i+1]))
		}
		var cond_a = is_prev_lower && is_curr_upper                                     // aA
		var cond_b = (is_prev_upper || is_prev_digit) && is_curr_upper && is_next_lower // AAa, 2Aa
		var cond_c = is_prev_digit && is_curr_lower && is_next_lower                    // 2aa
		var cond_d = (is_prev_upper || is_prev_lower) && is_curr_digit                  // A2, a2
		if cond_a || cond_b || cond_c || cond_d {
			new_string = JoinedWith("", new_string, s[start_index:i], S("_"))
			start_index = i
		}
	}
	new_string = JoinedWith("", new_string, s[start_index:])
	return S(strings.Replace(strings.ToLower(string(new_string)), " ", "_", -1))
}

// ToUpper returns the string converted to UPPERCASE.
func ToUpper[S Any](s S) S { //gd:String.to_upper
	return S(strings.ToUpper(string(s)))
}

// ToUTF8 converts the string to a UTF-8 encoded byte slice. This method is slightly slower than
// [ToASCII], but supports all UTF-8 characters. For most cases, prefer using this method.
func ToUTF8[S Any](s S) []byte { //gd:String.to_utf8_buffer
	return []byte(string(s))
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
	for _, r := range string(s) {
		a, b := utf16.EncodeRune(r)
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
	b, _ := UTF.Bytes([]byte(s))
	return b
}

// UTF32 converts UTF-32 encoded array to string. If the BOM is missing, system endianness
// is assumed. Returns empty string if source array is not valid UTF-32 string. This is
// the inverse of [ToUTF32].
func UTF32[S Any](s S) string { //gd:PackedByteArray.get_string_from_utf32
	var UTF = utf32.UTF32(utf32.LittleEndian, utf32.UseBOM).NewDecoder()
	b, _ := UTF.String(string(s))
	return b
}

// TrimPrefix removes the given prefix from the start of the string, or returns the string unchanged.
func TrimPrefix[S Any](s S, prefix S) S { //gd:String.trim_prefix
	return S(strings.TrimPrefix(string(s), string(prefix)))
}

// TrimSuffix removes the given suffix from the end of the string, or returns the string unchanged.
func TrimSuffix[S Any](s S, suffix S) S { //gd:String.trim_suffix
	return S(strings.TrimSuffix(string(s), string(suffix)))
}

// UnicodeAt returns the Unicode code point at the given index in the string.
func UnicodeAt[S Any](s S, index int) rune { //gd:String.unicode_at
	for i, r := range string(s) {
		if i == index {
			return r
		}
	}
	return 0
}

// StripFilename returns a copy of the string with all characters that are not allowed in [IsValidFilename] replaced with underscores.
func StripFilename[S Any](s S) S { //gd:String.validate_filename
	var result strings.Builder
	for _, r := range string(s) {
		if IsValidFilename(string(r)) {
			result.WriteRune(r)
		} else {
			result.WriteRune('_')
		}
	}
	return S(result.String())
}

// StripNodeName returns a copy of the string with all characters that are not allowed in Node.name (. : @ / " %) replaced with underscores.
func StripNodeName[S Any](s S) S { //gd:String.validate_node_name
	return S(strings.Map(func(r rune) rune {
		switch r {
		case '.', ':', '@', '/', '"', '%':
			return '_'
		}
		return r
	}, string(s)))
}

// StartingFrom returns a slice of the string from the given start index to the end index.
func StartingFrom[S Any](s S, start int) S { //gd:String.substr
	return S(string(s)[start:])
}

// EncodeURI encodes the string to URL-friendly format. This method is meant to properly
// encode the parameters in a URL when sending an HTTP request. See also [DecodeURI].
func EncodeURI[S Any](s S) S { //gd:String.uri_encode
	return S(url.QueryEscape(string(s)))
}

// DecodeURI decodes the string from URL-friendly format. This method is meant to properly
// decode the parameters in a URL when receiving an HTTP request. See also [EncodeURI].
func DecodeURI[S Any](s S) S { //gd:String.uri_decode
	decoded, _ := url.QueryUnescape(string(s))
	return S(decoded)
}
