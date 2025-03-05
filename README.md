# p

inspired by [serokell article](https://serokell.io/blog/parser-combinators-in-haskell)

```go
func satisfy[I, E any](f func(I) bool) p.Parser[I, E, I] {
	type L = []p.Error
	type R = p.Tuple2[I, []I]
	type FR = p.Either[L, R]
	return func(i []I) FR {
		if len(i) == 0 {
			return p.Left[L, R](&L{p.EndOfInput{}})
		} else if f(i[0]) {
			return p.Right[L, R](&R{i[0], i[1:]})
		}
		return p.Left[L, R](&L{p.Unexpected[I]{i[0]}})
	}
}

func char[I comparable, E any](i I) p.Parser[I, E, I] {
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
    return p.Left[error, int](&err)
}
return p.Right[error, int](&42)

func handle(e p.Either[error, int]) int {
    if _, ok := e.Left(); ok {
        return 69
    } else if res, ok := e.Right(); ok {
        return *res
    }
    return 0
}
```
