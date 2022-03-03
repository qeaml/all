package funcs

// Pipe returns a function that applies pipes the given input through all the
// provided functions. Note that the input and output types must be the same.
func Pipe[T any](fs ...func(T) T) func(T) T {
	if len(fs) < 2 {
		panic("Chain() requires at least two functions")
	}
	return func(i T) T {
		v := i
		for _, f := range fs {
			v = f(v)
		}
		return v
	}
}

// Partial returns the given function but with one argument pre-defined.
func Partial[T1 any, T2 any, R any](t1 T1, f func(T1, T2) R) func(T2) R {
	return func(t2 T2) R {
		return f(t1, t2)
	}
}
