package funcs_test

import (
	"errors"
	"fmt"
	"testing"

	// the stuff we're testing
	. "github.com/qeaml/all/funcs"
)

func TestDoOrPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered: %v\n", r)
		}
	}()
	DoOrPanic(func() error {
		return errors.New("why he ourple")
	})
}

func TestGetOrPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered: %v\n", r)
		}
	}()
	GetOrPanic(func() (int, error) {
		return 10, errors.New("something bad happen :(")
	})
}

func TestConsumeOrPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered: %v\n", r)
		}
	}()
	ConsumeOrPanic(func(i int) error {
		_, err := fmt.Println(i)
		return err
	}, 10)
}

func TestReturnOrPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered: %v\n", r)
		}
	}()
	ReturnOrPanic(func(i int) (int, error) {
		return fmt.Println(i)
	}, 123)
}
