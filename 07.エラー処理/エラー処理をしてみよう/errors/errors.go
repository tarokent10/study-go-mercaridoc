package errors

import "fmt"

// SampleError is sampleerror
type SampleError struct {
	msg string
}

func (e *SampleError) Error() string {
	return fmt.Sprintf("%s\n", e.msg)
}

// NewSampleError create new error
func NewSampleError(s string) error {
	return &SampleError{s}
}
