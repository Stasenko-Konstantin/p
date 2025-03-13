package p

import "fmt"

/*
Either - type from functional languages for error handling. Left often is error, Right - result
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
		res  any = 0
		side string
	)
	if l, ok := e.Left(); ok {
		res, side = *l, "Left"
	} else if r, ok := e.Right(); ok {
		res, side = *r, "Right"
	}
	return fmt.Sprintf("%s( %v )", side, res)
}
