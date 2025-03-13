package p

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mkEither(err error) Either[error, int] {
	if err != nil {
		return Left[error, int](&err)
	}
	res := 42
	return Right[error, int](&res)
}

func handle(e Either[error, int]) int {
	if _, ok := e.Left(); ok {
		return 0
	} else if res, ok := e.Right(); ok {
		return *res
	}
	return -1
}

func TestEither(t *testing.T) {
	e1 := mkEither(errors.New("err"))
	assert.Equal(t, 0, handle(e1))
	assert.Equal(t, "Left( err )", e1.String())

	e2 := mkEither(nil)
	assert.Equal(t, 42, handle(e2))
	assert.Equal(t, "Right( 42 )", e2.String())

	e3 := Either[any, any]{}
	assert.Equal(t, "( 0 )", e3.String())
	_, ok := e3.Left()
	assert.False(t, ok)
	_, ok = e3.Right()
	assert.False(t, ok)
}
