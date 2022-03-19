package slices

// All returns true if all the given predicates return true when called with
// every element of the slice.
func All[T any](s []T, preds ...func(T) bool) bool {
	for _, v := range s {
		for _, pred := range preds {
			if !pred(v) {
				return false
			}
		}
	}
	return true
}

// Any returns true if at least one given predicate returns true when called
// with any element of the slice.
func Any[T any](s []T, preds ...func(T) bool) bool {
	for _, v := range s {
		for _, pred := range preds {
			if pred(v) {
				return true
			}
		}
	}
	return false
}
