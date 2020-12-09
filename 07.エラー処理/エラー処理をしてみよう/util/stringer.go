package util

import (
	"fmt"
	"study-go--mercaridoc/07.エラー処理/エラー処理をしてみよう/errors"
)

// Stringer is interface
type Stringer interface {
	String() string
}

type sampleStringer struct {
	v interface{}
}

func (s *sampleStringer) String() string {
	return fmt.Sprintf("%#v\n", s.v)
}

// ToStringer converts to Stringer.
// if nil, return error
func ToStringer(v interface{}) (Stringer, error) {
	if v == nil {
		return nil, errors.NewSampleError("v must not nil.")
	}
	return &sampleStringer{v}, nil
}
