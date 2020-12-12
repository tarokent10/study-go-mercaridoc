package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type Stringer interface {
	String() string
}
type SampleError struct {
	error
}

func (e *SampleError) String() string {
	return e.Error()
}

func main() {
	// err := errors.Wrap(mkError("test"), ":wrap")
	err := mkError("test")
	fmt.Println(err)
	if isStringer(err) {
		fmt.Println(errors.Unwrap(err))
	}

}
func isStringer(err error) bool {
	s, ok := err.(Stringer)
	print(s.String())
	return ok
}
func mkError(s string) error {
	return SampleError{errors.New(s)}
}
