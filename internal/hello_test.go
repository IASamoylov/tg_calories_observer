package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	t.Parallel()

	result := SayHello("Alice")

	assert.Equal(t, "Hi, Alice", result)
}

type Test struct {
	in  int
	out string
}

var tests = []Test{
	{-1, "negative"},
	{5, "small"},
	{1000, "enormous"},
	{100, "huge"},
}

func TestSize(t *testing.T) {
	t.Parallel()

	for i, test := range tests {
		size := Size(test.in)
		if size != test.out {
			t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
		}
	}
}
