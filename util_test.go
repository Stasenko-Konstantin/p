package p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	tCases := struct {
		len             int
		input           []string
		expectedRunes   []rune
		expectedStrings []string
	}{
		0,
		[]string{
			"", "eas", "as", "aaaeas",
		},
		[]rune{
			'_', 'e', '_', '_',
		},
		[]string{
			"", "as", "as", "aaaeas",
		},
	}
	tCases.len = len(tCases.input)
	if tCases.len != len(tCases.expectedRunes) && tCases.len != len(tCases.expectedStrings) {
		t.Fatalf("Map: lengths of input, runes & strings not equal!")
	}

	r := Map(tCases.input, func(s string) Tuple2[rune, []rune] {
		e := char[rune, string]('e').RunParser([]rune(s))
		if r, ok := e.Right(); ok {
			return *r
		}
		return Tuple2[rune, []rune]{E1: '_', E2: []rune(s)}
	})
	for i := 0; i < tCases.len; i++ {
		assert.Equal(t, tCases.expectedRunes[i], r[i].E1)
		assert.Equal(t, tCases.expectedStrings[i], string(r[i].E2))
	}
}
