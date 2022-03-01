package funcs

// Not returns the opposite of the given function
func Not[T any](f func(T) bool) func(T) bool {
	return func(i T) bool {
		return !f(i)
	}
}

// Or returns true if either of the given functions return true
func Or[T any](f func(T) bool, g func(T) bool) func(T) bool {
	return func(i T) bool {
		return f(i) || g(i)
	}
}

// And returns true if both of the given functions return true
func And[T any](f func(T) bool, g func(T) bool) func(T) bool {
	return func(i T) bool {
		return f(i) && g(i)
	}
}

// Xor returns true if either of the given functions return true, but returns
// false if both of them are true or both of them are false
func Xor[T any](f func(T) bool, g func(T) bool) func(T) bool {
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
func All[T any](fs ...func(T) bool) func(T) bool {
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
func Any[T any](fs ...func(T) bool) func(T) bool {
	return func(i T) bool {
		for _, f := range fs {
			if f(i) {
				return true
			}
		}
		return false
	}
}
