// Package Dictionary provides a data structure that holds key-value pairs.
package Dictionary

import (
	"iter"
	"reflect"

	"graphics.gd/variant"
)

// Map is an associative container that contain values referenced by unique keys.
// Dictionaries will preserve the insertion order when adding new entries. In
// other programming languages, this data structure is often referred to as a
// hash map or an associative array.
type Map[K comparable, V any] struct {
	state complex128
	proxy Proxy[K, V]
}

type Any = Map[variant.Any, variant.Any]

var Nil Any

func New[K comparable, V any]() Map[K, V] {
	return Map[K, V]{
		proxy: &localFirst[K, V]{
			value: make(map[K]V),
		},
	}
}

// Any returns a dictionary with variant keys and values.
func (m Map[K, V]) Any() Any {
	if m.proxy == nil {
		return Nil
	}
	return m.proxy.Any(m.state)
}

// Index returns the value at the given index in the dictionary.
func (m Map[K, V]) Index(index K) V { //gd:Dictionary.get Dictionary[]
	if m.proxy == nil {
		return [1]V{}[0]
	}
	return m.proxy.Index(m.state, index)
}

// SetIndex sets the value at the given index in the dictionary.
func (m *Map[K, V]) SetIndex(index K, value V) {
	if m.proxy == nil {
		*m = New[K, V]()
	}
	m.proxy.SetIndex(m.state, index, value)
}

// Clear clears the dictionary, removing all entries from it.
func Clear[K comparable, V any](m Map[K, V]) { //gd:Dictionary.clear
	if m.proxy == nil {
		return
	}
	m.proxy.Clear(m.state)
}

// Duplicate creates and returns a new shallow copy of the dictionary.
func Duplicate[K comparable, V any](m Map[K, V]) Map[K, V] { //gd:Dictionary.duplicate
	if m.proxy == nil {
		return Map[K, V]{}
	}
	var clone = New[K, V]()
	for key, value := range m.Iter() {
		clone.SetIndex(key, value)
	}
	return clone
}

// Erase removes the dictionary entry by key, if it exists. Returns true
// if the given key existed in the dictionary, otherwise false.
//
// Note: Do not erase entries while iterating over the dictionary. You can
// iterate over the keys array instead.
func (m Map[K, V]) Erase(key K) bool { //gd:Dictionary.erase
	if m.proxy == nil {
		return false
	}
	return m.proxy.Erase(m.state, key)
}

// FindKey finds and returns the first key whose associated value is
// equal to value, or false if it is not found.
func FindKey[K, V comparable](m Map[K, V], value V) (K, bool) { //gd:Dictionary.find_key
	if m.proxy == nil {
		return [1]K{}[0], false
	}
	for k, v := range m.Iter() {
		if v == value {
			return k, true
		}
	}
	var zero K
	return zero, false
}

// GetOrAdd returns the corresponding value for the given key in the dictionary.
// If the key does not exist, it will be added with the default value.
func (m *Map[K, V]) GetOrAdd(key K, value V) V { //gd:Dictionary.get_or_add
	if m.proxy == nil {
		*m = New[K, V]()
	}
	if v, ok := m.proxy.Lookup(m.state, key); ok {
		return v
	}
	m.proxy.SetIndex(m.state, key, value)
	return value
}

// Has returns true if the dictionary contains an entry with the given key.
func (m Map[K, V]) Has(key K) bool { //gd:Dictionary.has
	if m.proxy == nil {
		return false
	}
	return m.proxy.Has(m.state, key)
}

// HasAll returns true if the dictionary contains all the keys.
func (m Map[K, V]) HasAll(keys ...K) bool { //gd:Dictionary.has_all
	for _, key := range keys {
		if !m.Has(key) {
			return false
		}
	}
	return true
}

// Hash returns a hashed 32-bit integer value representing the dictionary contents.
func (m Map[K, V]) Hash() uint32 { //gd:Dictionary.hash
	if m.proxy == nil {
		return 0
	}
	return m.proxy.Hash(m.state)
}

// IsEmpty returns true if the dictionary is empty.
func (m Map[K, V]) IsEmpty() bool { //gd:Dictionary.is_empty
	if m.proxy == nil {
		return true
	}
	return m.proxy.Len(m.state) == 0
}

// IsReadOnly returns true if the dictionary is read-only.
func (m Map[K, V]) IsReadOnly() bool { //gd:Dictionary.is_read_only
	if m.proxy == nil {
		return false
	}
	return m.proxy.IsReadOnly(m.state)
}

// Keys returns the list of keys in the dictionary.
func (m *Map[K, V]) Keys() []K { //gd:Dictionary.keys
	if m.proxy == nil {
		return nil
	}
	var keys []K
	for key := range m.Iter() {
		keys = append(keys, key)
	}
	return keys
}

// MakeReadOnly makes the dictionary read-only, i.e. disables modification of
// the dictionary's contents. Does not apply to nested content, e.g. content
// of nested dictionaries.
func (m *Map[K, V]) MakeReadOnly() { //gd:Dictionary.make_read_only
	if m.proxy == nil {
		*m = New[K, V]()
	}
	m.proxy.MakeReadOnly(m.state)
}

// Merge adds entries from dictionary to this dictionary. Duplicate
// keys are not copied over.
func (m *Map[K, V]) Merge(dictionary Map[K, V]) { //gd:Dictionary.merge
	if m.proxy == nil {
		*m = New[K, V]()
	}
	for key, val := range dictionary.Iter() {
		if !m.Has(key) {
			m.SetIndex(key, val)
		}
	}
}

// Merged returns a copy of this dictionary merged with the other dictionary.
func Merged[K comparable, V any](a Map[K, V], b Map[K, V]) Map[K, V] { //gd:Dictionary.merged
	result := Duplicate(a)
	result.Merge(b)
	return result
}

// RecursivelyEqualTo returns true if the two dictionaries contain the same keys and
// values, inner Dictionary and Array keys and values are compared recursively.
func (m Map[K, V]) RecursivelyEqualTo(other Map[K, V]) bool { //gd:Dictionary.recursive_equal
	if m.proxy == nil && other.proxy == nil {
		return true
	}
	if m.proxy == nil || other.proxy == nil {
		return false
	}
	for key, value := range m.Iter() {
		if !other.Has(key) {
			return false
		}
		if !reflect.DeepEqual(value, other.Index(key)) {
			return false
		}
	}
	return true
}

// Size returns the number of entries in the dictionary. Empty dictionaries ({ })
// always return 0. See also [IsEmpty].
func (m Map[K, V]) Len() int { //gd:Dictionary.size
	if m.proxy == nil {
		return 0
	}
	return m.proxy.Len(m.state)
}

// Values returns the list of values in this dictionary.
func (m Map[K, V]) Values() []V { //gd:Dictionary.values
	if m.proxy == nil {
		return nil
	}
	var values = make([]V, 0, m.Len())
	for _, value := range m.Iter() {
		values = append(values, value)
	}
	return values
}

// Iter returns an iterator for the dictionary.
func (m *Map[K, V]) Iter() iter.Seq2[K, V] {
	if m.proxy == nil {
		return func(yield func(K, V) bool) {}
	}
	return m.proxy.Iter(m.state)
}
