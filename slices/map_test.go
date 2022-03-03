package slices_test

import (
	"testing"

	// Make it easier on ourselves
	. "github.com/qeaml/all/slices"
)

func TestMap(t *testing.T) {
	gt = t
	theslice := []int{7, 9, 2, 1}
	Map(theslice, func(i int) int {
		return i * 2
	})
	assertEquals(theslice[0], 14)
	assertEquals(theslice[1], 18)
	assertEquals(theslice[2], 4)
	assertEquals(theslice[3], 2)
}

func TestMapPairs(t *testing.T) {
	gt = t
	theslice := []int{7, 9, 2, 1}
	MapPairs(theslice, func(i int, v int) int {
		return i + v
	})
	assertEquals(theslice[0], 7)
	assertEquals(theslice[1], 10)
	assertEquals(theslice[2], 4)
	assertEquals(theslice[3], 4)
}

func TestFilter(t *testing.T) {
	gt = t
	theslice := []int{7, 9, 2, 1}
	newslice := Filter(theslice, func(i int) bool {
		return i%2 == 0
	})
	assertEquals(len(newslice), 1)
	assertEquals(newslice[0], 2)
}

func TestFilterPairs(t *testing.T) {
	gt = t
	theslice := []int{7, 9, 2, 1}
	newslice := FilterPairs(theslice, func(i int, v int) bool {
		return i%2 == 0 || v%2 == 0
	})
	assertEquals(len(newslice), 2)
	assertEquals(newslice[0], 7)
	assertEquals(newslice[1], 2)
}
