package slices_test

import (
	"strings"
	"testing"

	"github.com/qeaml/all/funcs"
	"github.com/qeaml/all/slices"
)

var gt *testing.T

func assert(what bool) {
	if !what {
		gt.Fatal("assertion fail")
	}
}

func assertEquals[T comparable](a T, b T) {
	if a != b {
		gt.Fatalf("assertion fail: %v != %v", a, b)
	}
}

func TestContains(t *testing.T) {
	gt = t
	theslice := []string{"one", "two", "three", "four"}
	cond := funcs.Partial(theslice, slices.Contains[string])
	assert(cond("one"))
	assert(cond("two"))
	assert(cond("three"))
	assert(cond("four"))
}

func TestContainsComparator(t *testing.T) {
	gt = t
	theslice := []string{"One", "tWo", "thRee", "fouR"}
	assert(slices.ContainsComparator(theslice, "one", strings.EqualFold))
	assert(slices.ContainsComparator(theslice, "two", strings.EqualFold))
	assert(slices.ContainsComparator(theslice, "three", strings.EqualFold))
	assert(slices.ContainsComparator(theslice, "four", strings.EqualFold))
}

func TestReverse(t *testing.T) {
	gt = t
	theslice := []int{22, 678, 9, 500}
	slices.Reverse(theslice)
	assertEquals(theslice[0], 500)
	assertEquals(theslice[1], 9)
	assertEquals(theslice[2], 678)
	assertEquals(theslice[3], 22)
	slices.Reverse(theslice)
	assertEquals(theslice[3], 500)
	assertEquals(theslice[2], 9)
	assertEquals(theslice[1], 678)
	assertEquals(theslice[0], 22)
}

func TestSortDesc(t *testing.T) {
	gt = t
	theslice := []int{22, 678, 9, 500}
	slices.SortAsc(theslice)
	assertEquals(theslice[0], 9)
	assertEquals(theslice[1], 22)
	assertEquals(theslice[2], 500)
	assertEquals(theslice[3], 678)
}

func TestSortAsc(t *testing.T) {
	gt = t
	theslice := []int{22, 678, 9, 500}
	slices.SortDesc(theslice)
	assertEquals(theslice[3], 9)
	assertEquals(theslice[2], 22)
	assertEquals(theslice[1], 500)
	assertEquals(theslice[0], 678)
}
