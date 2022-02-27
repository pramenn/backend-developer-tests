package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		input    []int64
		expected []string
	}{
		{input: []int64{10, 5, 2}, expected: []string{"1", "Buzz", "3", "Buzz", "Fizz", "Buzz", "7", "Buzz", "9", "FizzBuzz"}},
		{input: []int64{7, 3, -2}, expected: []string{"1", "Buzz", "Fizz", "Buzz", "5", "FizzBuzz", "7"}},
		{input: []int64{10, 0, 0}, expected: []string{}},
		{input: []int64{5, 1, 0}, expected: []string{}},
		{input: []int64{0, 1, 1}, expected: []string{}},
	}

	for _, tc := range testCases {
		output := FizzBuzz(tc.input[0], tc.input[1], tc.input[2])
		assert.Equal(t, output, tc.expected)
	}
}
