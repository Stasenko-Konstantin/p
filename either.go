package p

import "fmt"

/*
Either - type from functional languages for error handling. Left often is error, Right - result

	if err != nil {
		return p.Left(&err)
	}
	return p.Right(&42)

	func handle(e Either[error, int]) int {
		if _, ok := e.Left(); ok {
			return 0
		} else if res, ok := e.Right(); ok {
			return *res
		}
	}
*/
type Either[L, R any] struct {
	left  *L
	right *R
}

func Left[L, R any](left *L) Either[L, R] {
	return Either[L, R]{left: left}
}

func Right[L, R any](right *R) Either[L, R] {
	return Either[L, R]{right: right}
}

func (e Either[L, _]) Left() (left *L, ok bool) {
	if e.left != nil {
		return e.left, true
	}
	return nil, false
}

func (e Either[_, R]) Right() (right *R, ok bool) {
	if e.right != nil {
		return e.right, true
	}
	return nil, false
}

func (e Either[L, R]) String() string {
	var (
		res  any
		side string
	)
	if l, ok := e.Left(); ok {
		res, side = l, "Left"
	} else if r, ok := e.Right(); ok {
		res, side = r, "Right"
	}
	return fmt.Sprintf("%s( %v )", side, res)
}
