package nio

import (
	"errors"
	"io"
)

var tooUnread = errors.New("nio: Unread() index exceeded backlog size")

// BacklogByteScanner provides an infinitely Unreadable ByteScanner
type BacklogByteScanner struct {
	src     io.ByteScanner
	backlog []byte
	idx     int
}

// ByteBacklog wraps the ByteScanner with a backlog
func ByteBacklog(src io.ByteScanner) *BacklogByteScanner {
	return &BacklogByteScanner{
		src:     src,
		backlog: []byte{},
		idx:     -1,
	}
}

// ReadByte reads a byte, for more info check io.ByteReader
func (s *BacklogByteScanner) ReadByte() (b byte, err error) {
	if s.idx >= 0 {
		b = s.backlog[len(s.backlog)-s.idx]
		s.idx--
		return
	}
	b, err = s.src.ReadByte()
	s.backlog = append(s.backlog, b)
	return
}

// UnreadByte unreads a byte. This returns an error if the amount of Unread()
// calls exceeds the amount of Read() calls.
func (s *BacklogByteScanner) UnreadByte() (err error) {
	s.idx++
	if s.idx >= len(s.backlog) {
		return tooUnread
	}
	return
}

type sizedRune struct {
	Rune rune
	Size int
}

// BacklogRuneScanner provides an infinitely Unreadable RuneScanner
type BacklogRuneScanner struct {
	src     io.RuneScanner
	backlog []sizedRune
	idx     int
}

// Runeacklog wraps the RuneScanner with a backlog
func RuneBacklog(src io.RuneScanner) *BacklogRuneScanner {
	return &BacklogRuneScanner{
		src:     src,
		backlog: []sizedRune{},
		idx:     -1,
	}
}

// ReadRune reads a rune, for more info check io.RuneReader
func (s *BacklogRuneScanner) ReadRune() (r rune, size int, err error) {
	if s.idx >= 0 {
		rn := s.backlog[len(s.backlog)-s.idx]
		r = rn.Rune
		size = rn.Size
		s.idx--
		return
	}
	r, size, err = s.src.ReadRune()
	s.backlog = append(s.backlog, sizedRune{
		Rune: r,
		Size: size,
	})
	return
}

// UnreadRune unreads a rune. This returns an error if the amount of Unread()
// calls exceeds the amount of Read() calls.
func (s *BacklogRuneScanner) UnreadRune() (err error) {
	s.idx++
	if s.idx >= len(s.backlog) {
		return tooUnread
	}
	return
}
