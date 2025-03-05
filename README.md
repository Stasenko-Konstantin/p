# p

inspired by [serokell article](https://serokell.io/blog/parser-combinators-in-haskell)

```go
func satisfy[I, E any](p func(I) bool) Parser[I, E, I] {
	type FR = Either[[]Error, Tuple2[I, []I]]
	type L = []Error
	type R = Tuple2[I, []I]
	return func(i []I) FR {
		if len(i) == 0 {
			return Left[L, R](&L{EndOfInput{}})
		} else if p(i[0]) {
			return Right[L, R](&R{i[0], i[1:]})
		}
		return Left[L, R](&L{Unexpected[I]{i[0]}})
	}
}

func char[I comparable, E any](i I) Parser[I, E, I] {
	return satisfy[I, E](func(i2 I) bool {
		return i == i2
	})
}

func Test(t *testing.T) {
	e, ok := char[rune, string]('e').RunParser([]rune("eas")).Right()
	assert.True(t, ok)
	assert.Equal(t, "e", string(e.E1))
	assert.Equal(t, "as", string(e.E2))
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
