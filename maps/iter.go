package maps

import "github.com/qeaml/all/funcs"

// MapPairs replaces all the values in the map with the return value of the
// provided function. ALl modifications are done in-place.
func MapPairs[K comparable, V any](m map[K]V, f func(K, V) V) {
	for k, v := range m {
		m[k] = f(k, v)
	}
}

// MapValues is the same as MapPairs, but only provides the value of the pair
// to the provided function
func MapValues[K comparable, V any](m map[K]V, f func(V) V) {
	for k, v := range m {
		m[k] = f(v)
	}
}

// MapKeys is the same as MapPairs, but only provides the key of the pair
// to the provided function
func MapKeys[K comparable, V any](m map[K]V, f func(K) V) {
	for k := range m {
		m[k] = f(k)
	}
}

// IteratePairs calls the given function with every key-value pair in the map.
func IteratePairs[K comparable, V any](m map[K]V, f func(K, V)) {
	for k, v := range m {
		f(k, v)
	}
}

// IterateValues is the same as IteratePairs, but only provides the value of the
// pair to the provided function
func IterateValues[K comparable, V any](m map[K]V, f func(V)) {
	for _, v := range m {
		f(v)
	}
}

// IterateKeys is the same as IteratePairs, but only provides the key of the
// pair to the provided function
func IterateKeys[K comparable, V any](m map[K]V, f func(K)) {
	for k := range m {
		f(k)
	}
}

// FilterPairs creates a copy of the provided map, only containing the keys
// for which the provided function returned true
func FilterPairs[K comparable, V any](m map[K]V, f func(K, V) bool) map[K]V {
	out := map[K]V{}
	for k, v := range m {
		if f(k, v) {
			out[k] = v
		}
	}
	return out
}

// FilterValues is the same as FilterPairs, but only provides the values of the
// pair to the provided function
func FilterValues[K comparable, V any](m map[K]V, f func(V) bool) map[K]V {
	out := map[K]V{}
	for k, v := range m {
		if f(v) {
			out[k] = v
		}
	}
	return out
}

// FilterKeys is the same as FiterPairs, but only provides the key of the
// pair to the provided function
func FilterKeys[K comparable, V any](m map[K]V, f func(K) bool) map[K]V {
	out := map[K]V{}
	for k, v := range m {
		if f(k) {
			out[k] = v
		}
	}
	return out
}

// ContainsKey returns true if the map contains the given key
func ContainsKey[K comparable, V any](m map[K]V, k K) (ok bool) {
	_, ok = m[k]
	return
}

// ContainsValue returns true if the map contains the given value
func ContainsValue[K comparable, V comparable](m map[K]V, v V) bool {
	return len(FilterValues(m, funcs.Is(v))) > 0
}
