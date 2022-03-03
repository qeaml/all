package slices

import (
	"constraints"
	"sort"

	"github.com/qeaml/all/funcs"
)

// Contains uses the builtin equal operator to check if an element is present
// in the slice
func Contains[V comparable](s []V, v V) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

// ContainsComparator uses the provided comparison function to check if an
// element is present in the slice
func ContainsComparator[V any](s []V, v V, c func(V, V) bool) bool {
	pc := funcs.Partial(v, c)
	for _, e := range s {
		if pc(e) {
			return true
		}
	}
	return false
}

func Reverse[V any](s []V) {
	clone := []V{}
	Iterate(s, func(v V) {
		clone = append(clone, v)
	})
	MapPairs(s, func(i int, v V) V {
		return clone[len(s)-i-1]
	})
}

// SortDesc sorts the slice in descending order
func SortDesc[V constraints.Ordered](s []V) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
}

// SortAsc sorts the slice in ascending order
func SortAsc[V constraints.Ordered](s []V) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}
