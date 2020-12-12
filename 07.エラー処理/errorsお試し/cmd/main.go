package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type sampleError struct {
	s string
}

func (e *sampleError) Error() string {
	return e.s
}
func (e *sampleError) String() string {
	return e.Error()
}

func main() {
	err := errors.Wrap(mkError("test"), ":wrap") // Wrap()時にはwithStackされる
	if isStringer(errors.Cause(err)) {
		// e := errors.Unwrap(err) //Unwrapはerr.Unwarp()を呼び出す。詳細情報用。Causeを返す訳ではない
		fmt.Printf("%+v\n", err)
	}

}
func isStringer(err error) bool {
	_, ok := err.(fmt.Stringer)
	return ok
}
func mkError(s string) error {
	return errors.WithStack(&sampleError{s}) // stacktrace付与
}
