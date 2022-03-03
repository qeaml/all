package funcs_test

import (
	"strings"
	"testing"

	// Make it easier on ourselves
	. "github.com/qeaml/all/funcs"
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

func TestPipe(t *testing.T) {
	gt = t
	var id = Pipe(strings.ToLower, strings.TrimSpace)
	assertEquals(id("  HELLO    "), "hello")
	assertEquals(id("Hello"), "hello")
	assertEquals(id(" O LE!"), "o le!")
	assertEquals(id(" - - -  WHAT"), "- - -  what")
}

func TestPartial(t *testing.T) {
	gt = t
	var cmp = Partial("hello", strings.EqualFold)
	assert(cmp("HELLO"))
	assert(cmp("hElLo"))
}
