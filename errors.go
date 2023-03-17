package p

/*
Error - sum type
its not sealed. you can create your own error type just implement one empty method

	EndOfInput - if input stream is ended, but expected more
	Unexpected - if current parsing element is unexpected
	CustomError - user custom error that he may want to create
	Empty - empty
*/
type Error interface {
	IsImplement()
}

type EndOfInput struct{}

// Unexpected I - input stream
type Unexpected[T any] struct {
	I T
}

// CustomError E - error
type CustomError[T any] struct {
	E T
}

type Empty struct{}

func (e EndOfInput) IsImplement()     {}
func (e Unexpected[T]) IsImplement()  {}
func (e CustomError[T]) IsImplement() {}
func (e Empty) IsImplement()          {}
