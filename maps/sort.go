package maps

import (
	"sort"

	"github.com/qeaml/all/generic"
)

// Sort returns the provided map sorted by keys using the provided comparison
// function. The function is the same as in sort.Slice from the standard lib
func Sort[K comparable, V any](m map[K]V, c func(keys []K, i, j int) bool) map[K]V {
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

// SortKeys returns the provided map sorted by keys in the provided order
func SortKeys[K generic.Ordered, V any](m map[K]V, ascending bool) map[K]V {
	if ascending {
		return Sort(m, func(keys []K, i, j int) bool {
			return keys[i] < keys[j]
		})
	}
	return Sort(m, func(keys []K, i, j int) bool {
		return keys[i] > keys[j]
	})
}

// SortValues returns the provided map sorted by values in the provided order
func SortValues[K comparable, V generic.Ordered](m map[K]V, ascending bool) map[K]V {
	if ascending {
		return Sort(m, func(keys []K, i, j int) bool {
			return m[keys[i]] < m[keys[j]]
		})
	}
	return Sort(m, func(keys []K, i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
}
