package gosoln

import (
	"reflect"
	"testing"

	"gotest.tools/assert"
)

func TestDecompressRLEList(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{2, 4, 4, 4}
	actual := DecompressRLEList(input)
	assert.Assert(t, true, reflect.DeepEqual(expected, actual))
}
