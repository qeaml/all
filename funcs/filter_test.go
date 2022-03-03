package funcs_test

import (
	"testing"

	// Make it easier on ourselves
	. "github.com/qeaml/all/funcs"
)

func isEven(n int) bool {
	return n%2 == 0
}

func isPositive(n int) bool {
	return n >= 0
}

func isBig(n int) bool {
	return n >= 42
}

func TestNot(t *testing.T) {
	gt = t
	isntEven := Not(isEven)
	assert(isntEven(3))
	assert(isntEven(15))
}

func TestOr(t *testing.T) {
	gt = t
	evenOrPositive := Or(isEven, isPositive)
	assert(evenOrPositive(2))
	assert(evenOrPositive(3))
	assert(evenOrPositive(-4))
	assert(Not(evenOrPositive)(-3))
}

func TestAnd(t *testing.T) {
	gt = t
	evenAndPositive := And(isEven, isPositive)
	assert(evenAndPositive(2))
	assert(evenAndPositive(4))
	assert(Not(evenAndPositive)(5))
	assert(Not(evenAndPositive)(-2))
	assert(Not(evenAndPositive)(-7))
}

func TestXor(t *testing.T) {
	gt = t
	either := Xor(isEven, isPositive)
	neither := Not(either)
	assert(either(-2))
	assert(either(3))
	assert(neither(2))
	assert(neither(-3))
}

func TestAll(t *testing.T) {
	gt = t
	criterion := All(isEven, isPositive, isBig)
	ncriterion := Not(criterion)
	assert(criterion(58))
	assert(criterion(444444))
	assert(ncriterion(59))
	assert(ncriterion(40))
	assert(ncriterion(-124))
}

func TestAny(t *testing.T) {
	gt = t
	criterion := Any(isEven, isPositive, isBig)
	ncriterion := Not(criterion)
	assert(criterion(2))
	assert(criterion(3))
	assert(criterion(79))
	assert(ncriterion(-3))
	assert(ncriterion(-777))
}
