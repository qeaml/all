package funcs

type filter[T any] func(T) bool

func Is[T comparable](v T) filter[T] {
	return func(t T) bool {
		return v == t
	}
}

// Isnt is a shorthand for Not(Is(...))
func Isnt[T comparable](v T) filter[T] {
	return Not(Is(v))
}

// Not returns the opposite of the given function
func Not[T any](f filter[T]) filter[T] {
	return func(i T) bool {
		return !f(i)
	}
}

// Or returns true if either of the given functions return true
func Or[T any](f filter[T], g filter[T]) filter[T] {
	return func(i T) bool {
		return f(i) || g(i)
	}
}

// And returns true if both of the given functions return true
func And[T any](f filter[T], g filter[T]) filter[T] {
	return func(i T) bool {
		return f(i) && g(i)
	}
}

// Xor returns true if either of the given functions return true, but returns
// false if both of them are true or both of them are false
func Xor[T any](f filter[T], g filter[T]) filter[T] {
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

// All returns true only if all provided functions return true
func All[T any](fs ...filter[T]) filter[T] {
	return func(i T) bool {
		for _, f := range fs {
			if !f(i) {
				return false
			}
		}
		return true
	}
}

// Any returns true only if at least one provided function returns true
func Any[T any](fs ...filter[T]) filter[T] {
	return func(i T) bool {
		for _, f := range fs {
			if f(i) {
				return true
			}
		}
		return false
	}
}
