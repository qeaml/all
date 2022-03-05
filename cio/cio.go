package cio

import (
	"io"
	"sync"
)

type concWriter struct {
	w io.Writer
	m sync.Mutex
}

// NewWriter returns a more thread-safe wrapper for the given writer
func NewWriter(target io.Writer) *concWriter {
	return &concWriter{target, sync.Mutex{}}
}

// Write acquires this wrapper's lock (waiting for it to be unlocked if needed)
// and then writes the given bytes to it's underlying writer
func (w *concWriter) Write(p []byte) (n int, err error) {
	w.m.Lock()
	defer w.m.Unlock()
	n, err = w.w.Write(p)
	return
}

type concReader struct {
	r io.Reader
	m sync.Mutex
}

// NewReader returns a more thread-safe wrapper for the given reader
func NewReader(target io.Reader) *concReader {
	return &concReader{target, sync.Mutex{}}
}

// Read acquires this wrapper's lock (waiting for it to be unlocked if needed)
// and then reads from it's underlying reader
func (r *concReader) Read(p []byte) (n int, err error) {
	r.m.Lock()
	defer r.m.Unlock()
	n, err = r.r.Read(p)
	return
}

type concRW struct {
	w io.Writer
	r io.Reader
	m sync.RWMutex
}

// NewReadWriter returns a more thread-safe wrapper for the given reader/writer
// pair
func NewReadWriter(r io.Reader, w io.Writer) *concRW {
	return &concRW{w, r, sync.RWMutex{}}
}

// Write acquires this wrapper's lock (waiting for it to be unlocked if needed)
// and then writes the given bytes to it's underlying writer
func (rw *concRW) Write(p []byte) (n int, err error) {
	rw.m.Lock()
	defer rw.m.Unlock()
	n, err = rw.w.Write(p)
	return
}

// Read acquires this wrapper's lock (waiting for it to be unlocked if needed)
// and then reads from it's underlying reader
func (rw *concRW) Read(p []byte) (n int, err error) {
	rw.m.RLock()
	defer rw.m.RUnlock()
	n, err = rw.r.Read(p)
	return
}
