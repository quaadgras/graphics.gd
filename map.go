package gd

import (
	gd "grow.graphics/gd/internal"
	"runtime.link/mmm"
)

/*
Map is an associative container that contain values referenced by unique keys. Map will
preserve the insertion order when adding new entries. K and V should be godot types.
*/
type Map[K, V any] Dictionary

// NewMap returns a new map with the given lifetime. K and V should be godot types.
func NewMap[K, V any](lt Lifetime) Map[K, V] {
	return Map[K, V](lt.Dictionary())
}

// Clear clears the map, removing all entries from it.
func (m Map[K, V]) Clear() { Dictionary(m).Clear() }

// Duplicate creates and returns a new copy of the map. If deep is true,
// inner Dictionary and Array keys and values are also copied, recursively.
func (m Map[K, V]) Duplicate(lt Lifetime, deep bool) Map[K, V] {
	return Map[K, V](Dictionary(m).Duplicate(lt, deep))
}

// Erase removes the map entry by key, if it exists. Returns true if the given
// key existed in the map, otherwise false.
//
// Note: Do not erase entries while iterating over the map. You can iterate
// over the keys array instead.
func (m Map[K, V]) Erase(key K) {
	tmp := gd.NewLifetime(mmm.API(Dictionary(m)))
	defer tmp.End()
	Dictionary(m).Erase(tmp.Variant(key))
}

// FindKey finds and returns the first key whose associated value is equal to value,
// or null if it is not found.
//
// Note: null is also a valid key. If inside the dictionary, find_key may give
// misleading results.
func (m Map[K, V]) FindKey(lt Lifetime, value V) K {
	tmp := gd.NewLifetime(mmm.API(Dictionary(m)))
	defer tmp.End()
	return Dictionary(m).FindKey(tmp, tmp.Variant(value)).Interface(lt).(K)
}

// Get returns the corresponding value for the given key in the map. If the key
// does not exist, returns default, or null if the parameter is omitted.
func (m Map[K, V]) Get(lt Lifetime, key K) V {
	tmp := gd.NewLifetime(mmm.API(Dictionary(m)))
	defer tmp.End()
	return Dictionary(m).Get(tmp, tmp.Variant(key), gd.Variant{}).Interface(lt).(V)
}

// Set sets the value for the given key in the map. If the key does not exist,
// it is added to the map.
func (m Map[K, V]) Set(key K, value V) {
	tmp := gd.NewLifetime(mmm.API(Dictionary(m)))
	defer tmp.End()
	Dictionary(m).SetIndex(tmp.Variant(key), tmp.Variant(value))
}

// Has returns true if the map contains an entry with the given key.
func (m Map[K, V]) Has(key K) bool {
	tmp := gd.NewLifetime(mmm.API(Dictionary(m)))
	defer tmp.End()
	return Dictionary(m).Has(tmp.Variant(key))
}

// HasAll returns true if the map contains all keys in the given keys array.
func (m Map[K, V]) HasAll(keys ArrayOf[K]) bool {
	tmp := gd.NewLifetime(mmm.API(Dictionary(m)))
	defer tmp.End()
	return Dictionary(m).HasAll(keys.Array())
}

// Hash returns a hashed 32-bit integer value representing the map's contents.
func (m Map[K, V]) Hash() int64 {
	return Dictionary(m).Hash()
}

// IsEmpty returns true if the map is empty (its length is 0). See also [Map.Len].
func (m Map[K, V]) IsEmpty() bool {
	return Dictionary(m).IsEmpty()
}

// IsReadOnly returns true if the map is read-only. See [Map.MakeReadOnly].
func (m Map[K, V]) IsReadOnly() bool {
	return Dictionary(m).IsReadOnly()
}

// Keys returns the list of keys in the map.
func (m Map[K, V]) Keys(lt Lifetime) ArrayOf[K] {
	return gd.TypedArray[K](Dictionary(m).Keys(lt))
}

// MakeReadOnly makes the map read-only, i.e. disables modification of the map's contents.
// Does not apply to nested content, e.g. content of nested maps.
func (m Map[K, V]) MakeReadOnly() {
	Dictionary(m).MakeReadOnly()
}

// Merge adds entries from another map to this map. By default, duplicate keys are not copied
// over, unless overwrite is true.
func (m Map[K, V]) Merge(with Map[K, V], overwrite bool) {
	tmp := gd.NewLifetime(mmm.API(Dictionary(m)))
	defer tmp.End()
	Dictionary(m).Merge(Dictionary(with), overwrite)
}

// Len returns the number of entries in the dictionary. Empty maps always return 0. See also [Map.IsEmpty].
func (m Map[K, V]) Len() int {
	return int(Dictionary(m).Size())
}

// Values returns the list of values in this dictionary.
func (m Map[K, V]) Values(lt Lifetime) ArrayOf[V] {
	return gd.TypedArray[V](Dictionary(m).Values(lt))
}
