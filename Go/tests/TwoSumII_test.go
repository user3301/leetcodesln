package tests

import (
	"gotest.tools/assert"
	"m/v2/questions"
	"testing"
)

func TestTwoSumII(t *testing.T) {
	numbers := []int{2,7,11,15}
	target := 9
	expected := []int{1,2}
	actual := questions.TwoSumII(numbers,target)
	for idx, val := range expected {
		assert.Equal(t, actual[idx], val)
	}
}
