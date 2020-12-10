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
	eof     bool
	reader  io.Reader
}

// Scan read from reader and scan a rune.
func (s *RuneScanner) Scan() bool {
	// はじめに全部がつっと読む.bufioでバッファ読み込みお試し.
	if !s.done {
		reader := bufio.NewReader(s.reader)
		rbuf := make([]byte, 1024)
		wbuf := make([]byte, 0, 1024)
		pos := 0
		for {
			n, err := reader.Read(rbuf)
			if n > 0 {
				wbuf = append(wbuf, rbuf[:n]...)
				pos += n
			}
			if err != nil {
				if err == io.EOF {
					s.done = true
					break
				}
				s.err = err
				return false
			}
		}
		s.buf = []rune(string(wbuf[:pos]))
	}
	if s.eof {
		return false
	}

	s.current = s.buf[s.idx]
	s.idx++

	if s.idx == len(s.buf) {
		s.eof = true
	}
	return (s.idx <= len(s.buf))
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
