package p

import "fmt"

type Either[L, R any] struct {
	left  *L
	right *R
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
		res any
		side string
	)
	if l, ok := e.Left(); ok {
		res, side = l, "Left"
	} else if r, ok := e.Right(); ok {
		res, side = r, "Right"
	}
	return fmt.Sprintf("%s( %v )", side, res)
}
