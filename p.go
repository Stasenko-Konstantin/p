package p

/*
Parser I - input stream. E - error, A - result of parsing

		func satisfy[I, E any](p func(I) bool) Parser[I, E, I] {
			return func(i []I) Either[[]Error, Tuple2[I, []I]] {
				if len(i) == 0 {
					return Either[[]Error, Tuple2[I, []I]]{left: &[]Error{EndOfInput{}}}
				} else if p(i[0]) {
					return Either[[]Error, Tuple2[I, []I]]{right: &Tuple2[I, []I]{i[0], i[1:]}}
				}
				return Either[[]Error, Tuple2[I, []I]]{left: &[]Error{Unexpected[I]{i[0]}}}
			}
		}

		func char[I comparable, E any](i I) Parser[I, E, I] {
			return satisfy[I, E](func(i2 I) bool {
				return i == i2
			})
		}

		func test_char() {
			r := char('e').RunParser("eas") // Right( Tuple2( 'e', "as" ) )
		}
*/
type Parser[I, E, A any] func([]I) Either[[]Error, Tuple2[A, []I]]

func (p Parser[I, E, A]) RunParser(i []I) Either[[]Error, Tuple2[A, []I]] {
	return p(i)
}
