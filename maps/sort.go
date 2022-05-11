package maps

import (
	"sort"

	"github.com/qeaml/all/generic"
)

// Sort returns the provided map sorted by keys using the provided comparison
// function. The function is the same as in sort.Slice from the standard lib
func Sort[K generic.Ordered, V any](m map[K]V, c func(keys []K, i, j int) bool) map[K]V {
	keys := []K{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return c(keys, i, j)
	})
	out := map[K]V{}
	for _, k := range keys {
		out[k] = m[k]
	}
	return out
}

// SortDesc returns the provided map sorted by keys in descending order
func SortDesc[K generic.Ordered, V any](m map[K]V) map[K]V {
	return Sort(m, func(keys []K, i, j int) bool {
		return keys[i] > keys[j]
	})
}

// SortAsc returns the provided map sorted by keys in ascending order
func SortAsc[K generic.Ordered, V any](m map[K]V) map[K]V {
	return Sort(m, func(keys []K, i, j int) bool {
		return keys[i] < keys[j]
	})
}
