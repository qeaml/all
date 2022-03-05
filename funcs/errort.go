package funcs

// DoOrPanic calls the given function, panicking if it returns an error.
// This isn't a good solution! You should always handle your errors properly
func DoOrPanic(f func() error) {
	if err := f(); err != nil {
		panic(err)
	}
}

// GetOrPanic calls the given function and forwards it's return value, panicking
// if it returns an error.
// This isn't a good solution! You should always handle your errors properly
func GetOrPanic[T any](f func() (T, error)) T {
	if t, err := f(); err != nil {
		panic(err)
	} else {
		return t
	}
}

// ConsumeOrPanic calls the given function with the provided argument, panicking
// if it returns an error.
// This isn't a good solution! You should always handle your errors properly
func ConsumeOrPanic[T any](f func(T) error, t T) {
	if err := f(t); err != nil {
		panic(err)
	}
}

// ReturnOrPanic calls the given function with the provided argument and
// forwards it's return value, panicking if it returns an error.
// This isn't a good solution! You should always handle your errors properly
func ReturnOrPanic[T any, R any](f func(T) (R, error), t T) R {
	if r, err := f(t); err != nil {
		panic(err)
	} else {
		return r
	}
}
