package maps_test

import (
	"testing"

	// Make it easier on ourselves
	. "github.com/qeaml/all/maps"
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

func getTestMap() map[string]int {
	return map[string]int{
		"Single":    10,
		"DOUBLE":    2,
		"tRIPLE":    3,
		"QUadRUplE": 4,
	}
}

func TestMapPairs(t *testing.T) {
	gt = t
	themap := getTestMap()
	MapPairs(themap, func(k string, v int) int {
		if k == "Single" {
			return 1
		}
		return v
	})
	assertEquals(themap["Single"], 1)
}

func TestMapKeys(t *testing.T) {
	gt = t
	themap := getTestMap()
	MapKeys(themap, func(k string) int {
		if k == "Single" {
			return 1
		}
		return 99
	})
	assertEquals(themap["Single"], 1)
	assertEquals(themap["DOUBLE"], 99)
}

func TestMapValues(t *testing.T) {
	gt = t
	themap := getTestMap()
	MapValues(themap, func(v int) int {
		return -v
	})
	assertEquals(themap["Single"], -10)
	assertEquals(themap["DOUBLE"], -2)
}

func TestIteratePairs(t *testing.T) {
	gt = t
	themap := getTestMap()
	count := 0
	IteratePairs(themap, func(k string, v int) {
		count++
	})
	assertEquals(count, 4)
}

func TestFilterPairs(t *testing.T) {
	gt = t
	themap := getTestMap()
	themap["Single"] = 1
	newmap := FilterPairs(themap, func(k string, v int) bool {
		return (k[0] >= 'a' && k[0] <= 'z') || v%2 == 0
	})
	assertEquals(len(newmap), 3)
	assertEquals(newmap["tRIPLE"], 3)
	assertEquals(newmap["QUadRUplE"], 4)
}

func TestFilterKeys(t *testing.T) {
	gt = t
	newmap := FilterKeys(getTestMap(), func(k string) bool {
		return (k[0] >= 'a' && k[0] <= 'z')
	})
	assertEquals(len(newmap), 1)
	assertEquals(newmap["tRIPLE"], 3)
}

func TestFilterValues(t *testing.T) {
	gt = t
	themap := getTestMap()
	themap["Single"] = 1
	newmap := FilterValues(themap, func(v int) bool {
		return v%2 == 0
	})
	assertEquals(len(newmap), 2)
	assertEquals(newmap["DOUBLE"], 2)
	assertEquals(newmap["QUadRUplE"], 4)
}
