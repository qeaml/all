package cio_test

import (
	"os"
	"strings"
	"testing"

	// what we testing
	"github.com/qeaml/all/cio"

	// useful
	. "github.com/qeaml/all/funcs"
)

func TestConcurrentReader(t *testing.T) {
	r := cio.NewReader(strings.NewReader("abc"))
	go ReturnOrPanic(r.Read, make([]byte, 1))
	go ReturnOrPanic(r.Read, make([]byte, 1))
	go ReturnOrPanic(r.Read, make([]byte, 1))
}

func TestConcurrentWriter(t *testing.T) {
	w := cio.NewWriter(os.Stderr)
	go ReturnOrPanic(w.Write, []byte("HELLO "))
	go ReturnOrPanic(w.Write, []byte("WORLD "))
	go ReturnOrPanic(w.Write, []byte("FOO "))
	go ReturnOrPanic(w.Write, []byte("BAR"))
}
