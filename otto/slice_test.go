package otto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringInSlice(t *testing.T) {
	res := InSlice("foo", []string{"bar", "foo", "slice", "pizza"})

	assert.True(t, res, "Result should be true as 'foo' is in the slice given")
}

func TestStringNotInSlice(t *testing.T) {
	res := InSlice("foo", []string{"bar", "slice", "pizza"})

	assert.False(t, res, "Result should not be true as 'foo' is not in the slice given")
}

func TestStringInSubString(t *testing.T) {
	res := InSlice("foo", []string{"bar", "slice foo bar", "pizza foo"})

	assert.True(t, res, "Result should not be true as 'foo' is not in the slice given")
}

func TestStringLowerCase(t *testing.T) {
	res := InSlice("foo", []string{"bar", "slice", "pizza", "FOO"})

	assert.True(t, res, "Result should not be true as 'foo' is not in the slice given")
}

func TestRemoveFromSlice(t *testing.T) {
	res := RemoveFromSlice([]string{"bar", "slice", "pizza", "FOO"}, 3)

	assert.Equal(t, []string{"bar", "slice", "pizza"}, res, "Result should not have the last element, which is FOO")
}
