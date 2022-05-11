package cio

import (
	"io"
	"sync"
)

type Writer struct {
	w io.Writer
	m sync.Mutex
}

// NewWriter returns a more thread-safe wrapper for the given writer
func NewWriter(target io.Writer) *Writer {
	return &Writer{target, sync.Mutex{}}
}

// Write acquires this wrapper's lock (waiting for it to be unlocked if needed)
// and then writes the given bytes to it's underlying writer
func (w *Writer) Write(p []byte) (n int, err error) {
	w.m.Lock()
	defer w.m.Unlock()
	n, err = w.w.Write(p)
	return
}

type Reader struct {
	r io.Reader
	m sync.Mutex
}

// NewReader returns a more thread-safe wrapper for the given reader
func NewReader(target io.Reader) *Reader {
	return &Reader{target, sync.Mutex{}}
}

// Read acquires this wrapper's lock (waiting for it to be unlocked if needed)
// and then reads from it's underlying reader
func (r *Reader) Read(p []byte) (n int, err error) {
	r.m.Lock()
	defer r.m.Unlock()
	n, err = r.r.Read(p)
	return
}

type ReadWriter struct {
	w io.Writer
	r io.Reader
	m sync.RWMutex
}

// NewReadWriter returns a more thread-safe wrapper for the given reader/writer
// pair
func NewReadWriter(r io.Reader, w io.Writer) *ReadWriter {
	return &ReadWriter{w, r, sync.RWMutex{}}
}

// Write acquires this wrapper's lock (waiting for it to be unlocked if needed)
// and then writes the given bytes to it's underlying writer
func (rw *ReadWriter) Write(p []byte) (n int, err error) {
	rw.m.Lock()
	defer rw.m.Unlock()
	n, err = rw.w.Write(p)
	return
}

// Read acquires this wrapper's lock (waiting for it to be unlocked if needed)
// and then reads from it's underlying reader
func (rw *ReadWriter) Read(p []byte) (n int, err error) {
	rw.m.RLock()
	defer rw.m.RUnlock()
	n, err = rw.r.Read(p)
	return
}
