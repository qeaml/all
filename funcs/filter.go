package funcs

type predicate[T any] func(T) bool

// Is produces true if the value it is given is equal
// to the predefined value
func Is[T comparable](v T) predicate[T] {
	return func(t T) bool {
		return v == t
	}
}

// Isnt is a shorthand for Not(Is(...))
func Isnt[T comparable](v T) predicate[T] {
	return Not(Is(v))
}

// AnyOf produces true if the value it is given is equal
// to any of the predefined values
func AnyOf[T comparable](s ...T) predicate[T] {
	return func(v T) bool {
		for _, e := range s {
			if v == e {
				return true
			}
		}
		return false
	}
}

// Not produces the opposite of the given predicate's output
func Not[T any](f predicate[T]) predicate[T] {
	return func(i T) bool {
		return !f(i)
	}
}

// Or produces true if either of the given predicates return true
func Or[T any](f predicate[T], g predicate[T]) predicate[T] {
	return func(i T) bool {
		return f(i) || g(i)
	}
}

// And produces true if both of the given predicates return true
func And[T any](f predicate[T], g predicate[T]) predicate[T] {
	return func(i T) bool {
		return f(i) && g(i)
	}
}

// Xor produces true if either of the given predicates return true, but produces
// false if both of them are true or both of them are false
func Xor[T any](f predicate[T], g predicate[T]) predicate[T] {
	return func(i T) bool {
		fr := f(i)
		gr := g(i)
		if fr && gr {
			return false
		} else if fr || gr {
			return true
		}
		return false
	}
}

// All produces true only if all provided predicates return true
func All[T any](fs ...predicate[T]) predicate[T] {
	return func(i T) bool {
		for _, f := range fs {
			if !f(i) {
				return false
			}
		}
		return true
	}
}

// Any produces true only if at least one provided predicate produces true
func Any[T any](fs ...predicate[T]) predicate[T] {
	return func(i T) bool {
		for _, f := range fs {
			if f(i) {
				return true
			}
		}
		return false
	}
}
