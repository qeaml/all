package slices

// Map replaces all values in the given slice with values returned by the
// provided function. All modifications are done in-place.
func Map[T any](s []T, f func(T) T) {
	for i, v := range s {
		s[i] = f(v)
	}
}

// MapPairs is the same as Map, but also provides the index to the function
func MapPairs[T any](s []T, f func(int, T) T) {
	for i, v := range s {
		s[i] = f(i, v)
	}
}

// Iterate runs the given function for each value in the given slice
func Iterate[T any](s []T, f func(T)) {
	for _, v := range s {
		f(v)
	}
}

// IteratePairs is the same as Iterate, but also provides the index to the
// function
func IteratePairs[T any](s []T, f func(int, T)) {
	for i, v := range s {
		f(i, v)
	}
}

// Filter creates a copy of the given slice, only containing the values for
// which the given function returned true (see example)
func Filter[T any](s []T, f func(T) bool) []T {
	out := []T{}
	Iterate(s, func(t T) {
		if f(t) {
			out = append(out, t)
		}
	})
	return out
}

// FilterPairs is the same as Filter, but also provides the index to the
// function
func FilterPairs[T any](s []T, f func(int, T) bool) []T {
	out := []T{}
	IteratePairs(s, func(i int, t T) {
		if f(i, t) {
			out = append(out, t)
		}
	})
	return out
}
