# p

```go
func satisfy[I, E any](p func(I) bool) Parser[I, E, I] {
    return func (i []I) Either[[]Error, Tuple2[I, []I]] {
        if len(i) == 0 {
            return Either[[]Error, Tuple2[I, []I]]{left: &[]Error{EndOfInput{}}}
        } else if p(i[0]) {
            return Either[[]Error, Tuple2[I, []I]]{right: &Tuple2[I, []I]{i[0], i[1:]}}
        }
        return Either[[]Error, Tuple2[I, []I]]{left: &[]Error{Unexpected[I]{i[0]}}}
    }
}

func char[I comparable, E any](i I) Parser[I, E, I] {
    return satisfy[I, E](func (i2 I) bool {
        return i == i2
    })
}

func test_char() {
    r := char('e').RunParser("eas") // Right( Tuple2( 'e', "as" ) )
}

////////////
// Either //
////////////

if err != nil {
    return p.Left(&err)
}
return p.Right(&42)

func handle(e Either[error, int]) int {
    if _, ok := e.Left(); ok {
        return 69
    } else if res, ok := e.Right(); ok {
        return *res
    }
    return 0
}
```