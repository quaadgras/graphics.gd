// Package Dictionary provides a data structure that holds key-value pairs.
package Dictionary

import (
	"iter"
	"maps"
	"reflect"

	gd "graphics.gd"
	"graphics.gd/variant"
)

// Map is an associative container that contain values referenced by unique keys.
// Dictionaries will preserve the insertion order when adding new entries. In
// other programming languages, this data structure is often referred to as a
// hash map or an associative array.
type Map[K comparable, V any] struct {
	order []K
	value map[K]V
	proxy gd.Dictionary
	fixed bool
}

type Any = gd.Dictionary

func New[K comparable, V any]() Map[K, V] {
	return Map[K, V]{
		value: make(map[K]V),
	}
}

// Index returns the value at the given index in the dictionary.
func (m *Map[K, V]) Index(index K) V { //gd:Dictionary.get Dictionary[]
	if m.proxy != (gd.Dictionary{}) {
		return m.proxy.Index(variant.New(index)).Interface().(V)
	}
	if v, ok := m.value[index]; ok {
		return v
	}
	var zero V
	return zero
}

// SetIndex sets the value at the given index in the dictionary.
func (m *Map[K, V]) SetIndex(index K, value V) {
	if m.proxy != (gd.Dictionary{}) {
		m.proxy.SetIndex(variant.New(index), variant.New(value))
		return
	}
	if _, ok := m.value[index]; ok {
		m.value[index] = value
		return
	}
	m.order = append(m.order, index)
	m.value[index] = value
}

// Clear clears the dictionary, removing all entries from it.
func Clear[K comparable, V any](m *Map[K, V]) { //gd:Dictionary.clear
	if m.proxy != (gd.Dictionary{}) {
		m.proxy.Clear()
		return
	}
	clear(m.order)
	if m.value != nil {
		clear(m.value)
	}
}

// Duplicate creates and returns a new shallow copy of the dictionary.
func Duplicate[K comparable, V any](m *Map[K, V]) Map[K, V] { //gd:Dictionary.duplicate
	if m.proxy != (gd.Dictionary{}) {
		return Map[K, V]{proxy: m.proxy.Duplicate(false)}
	}
	return Map[K, V]{
		order: append([]K(nil), m.order...),
		value: maps.Clone(m.value),
	}
}

// Erase removes the dictionary entry by key, if it exists. Returns true
// if the given key existed in the dictionary, otherwise false.
//
// Note: Do not erase entries while iterating over the dictionary. You can
// iterate over the keys array instead.
func (m *Map[K, V]) Erase(key K) bool { //gd:Dictionary.erase
	if m.proxy != (gd.Dictionary{}) {
		return bool(m.proxy.Erase(variant.New(key)))
	}
	if _, ok := m.value[key]; !ok {
		return false
	}
	delete(m.value, key)
	for i, k := range m.order {
		if k == key {
			m.order = append(m.order[:i], m.order[i+1:]...)
			break
		}
	}
	return true
}

// FindKey finds and returns the first key whose associated value is
// equal to value, or false if it is not found.
func FindKey[K, V comparable](m *Map[K, V], value V) (K, bool) { //gd:Dictionary.find_key
	if m.proxy != (gd.Dictionary{}) {
		return m.proxy.FindKey(variant.New(value)).Interface().(K), true
	}
	for k, v := range m.value {
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
	if m.proxy != (gd.Dictionary{}) {
		return m.proxy.GetOrAdd(variant.New(key), variant.New(value)).Interface().(V)
	}
	if v, ok := m.value[key]; ok {
		return v
	}
	m.order = append(m.order, key)
	m.value[key] = value
	return value
}

// Has returns true if the dictionary contains an entry with the given key.
func (m *Map[K, V]) Has(key K) bool { //gd:Dictionary.has
	if m.proxy != (gd.Dictionary{}) {
		return bool(m.proxy.Has(variant.New(key)))
	}
	_, ok := m.value[key]
	return ok
}

// HasAll returns true if the dictionary contains all the keys.
func (m *Map[K, V]) HasAll(keys ...K) bool { //gd:Dictionary.has_all
	for _, key := range keys {
		if !m.Has(key) {
			return false
		}
	}
	return true
}

// Hash returns a hashed 32-bit integer value representing the dictionary contents.
func (m *Map[K, V]) Hash() uint32 { //gd:Dictionary.hash
	if m.proxy != (gd.Dictionary{}) {
		return uint32(m.proxy.Hash())
	}
	panic("not implemented")
}

// IsEmpty returns true if the dictionary is empty.
func (m *Map[K, V]) IsEmpty() bool { //gd:Dictionary.is_empty
	if m.proxy != (gd.Dictionary{}) {
		return bool(m.proxy.IsEmpty())
	}
	return len(m.value) == 0
}

// IsReadOnly returns true if the dictionary is read-only.
func (m *Map[K, V]) IsReadOnly() bool { //gd:Dictionary.is_read_only
	if m.proxy != (gd.Dictionary{}) {
		return bool(m.proxy.IsReadOnly())
	}
	return m.fixed
}

// Keys returns the list of keys in the dictionary.
func (m *Map[K, V]) Keys() []K { //gd:Dictionary.keys
	if m.proxy != (gd.Dictionary{}) {
		var keys []K
		for _, key := range m.proxy.Keys().Iter() {
			keys = append(keys, key.Interface().(K))
		}
		return keys
	}
	return append([]K(nil), m.order...)
}

// MakeReadOnly makes the dictionary read-only, i.e. disables modification of
// the dictionary's contents. Does not apply to nested content, e.g. content
// of nested dictionaries.
func (m *Map[K, V]) MakeReadOnly() { //gd:Dictionary.make_read_only
	if m.proxy != (gd.Dictionary{}) {
		m.proxy.MakeReadOnly()
	}
	m.fixed = true
}

// Merge adds entries from dictionary to this dictionary. Duplicate
// keys are not copied over.
func (m *Map[K, V]) Merge(dictionary Map[K, V]) { //gd:Dictionary.merge
	if m.proxy != (gd.Dictionary{}) && dictionary.proxy != (gd.Dictionary{}) {
		m.proxy.Merge(dictionary.proxy, false)
		return
	}
	for key, val := range dictionary.Iter() {
		if !m.Has(key) {
			m.SetIndex(key, val)
		}
	}
}

// Merged returns a copy of this dictionary merged with the other dictionary.
func Merged[K comparable, V any](a *Map[K, V], b Map[K, V]) Map[K, V] { //gd:Dictionary.merged
	result := Duplicate(a)
	result.Merge(b)
	return result
}

// RecursivelyEqualTo returns true if the two dictionaries contain the same keys and
// values, inner Dictionary and Array keys and values are compared recursively.
func (m *Map[K, V]) RecursivelyEqualTo(other Map[K, V]) bool { //gd:Dictionary.recursive_equal
	if m.proxy != (gd.Dictionary{}) && other.proxy != (gd.Dictionary{}) {
		return bool(m.proxy.RecursiveEqual(other.proxy, -1))
	}
	if len(m.order) != len(other.order) {
		return false
	}
	return reflect.DeepEqual(m, other)
}

// Size returns the number of entries in the dictionary. Empty dictionaries ({ })
// always return 0. See also [IsEmpty].
func (m *Map[K, V]) Size() int { //gd:Dictionary.size
	if m.proxy != (gd.Dictionary{}) {
		return int(m.proxy.Size())
	}
	return len(m.value)
}

// Values returns the list of values in this dictionary.
func (m *Map[K, V]) Values() []V { //gd:Dictionary.values
	if m.proxy != (gd.Dictionary{}) {
		var values []V
		for _, value := range m.proxy.Values().Iter() {
			values = append(values, value.Interface().(V))
		}
	}
	values := make([]V, len(m.order))
	for i, key := range m.order {
		values[i] = m.value[key]
	}
	return values
}

// Iter returns an iterator for the dictionary.
func (m *Map[K, V]) Iter() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if m.proxy != (gd.Dictionary{}) {
			keys := m.proxy.Keys()
			for _, key := range keys.Iter() {
				if !yield(key.Interface().(K), m.proxy.Index(key).Interface().(V)) {
					return
				}
			}
			return
		}
		for _, key := range m.order {
			if !yield(key, m.value[key]) {
				return
			}
		}
	}
}
