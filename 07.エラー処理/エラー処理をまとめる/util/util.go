package util

import (
	"bufio"
	"io"
)

// RuneScanner read runes
type RuneScanner struct {
	err     error
	buf     []rune
	current rune
	idx     int
	done    bool
	reader  io.Reader
}

// Scan read from reader and scan a rune.
func (s *RuneScanner) Scan() bool {
	// はじめに全部がつっと読む.bufioでバッファ読み込みお試し.
	if !s.done {
		reader := bufio.NewReader(s.reader)
		buf := make([]byte, 1024)
		for {
			n, err := reader.Read(buf)
			if n > 0 {
				// byte↔rune変換：[]byte -> string -> []rune
				s.buf = append(s.buf, []rune(string(buf[:n]))...)
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				s.err = err
				return false
			}

		}
	}

	s.current = s.buf[s.idx]
	s.idx++

	return !(len(s.buf) == s.idx)
}

// Rune return current rune
func (s *RuneScanner) Rune() rune {
	return s.current
}

// Err return most recent error
func (s *RuneScanner) Err() error {
	return s.err
}

// NewRuneScanner create RuneScanner
func NewRuneScanner(r io.Reader) *RuneScanner {
	return &RuneScanner{
		buf:    make([]rune, 0, 1024),
		reader: r,
	}
}
