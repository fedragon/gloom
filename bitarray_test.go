package gloom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitArray_NewBitArray_UpTo64(t *testing.T) {
	b := NewBitArray(64)

	assert.NoError(t, b.Flip(63))
	assert.NoError(t, b.Flip(31))
	assert.NoError(t, b.Flip(0))

	assert.Equal(t, "8000000080000001", b.String())
}

func TestBitArray_NewBitArray_MoreThan64(t *testing.T) {
	b := NewBitArray(65)

	assert.NoError(t, b.Flip(64))
	assert.NoError(t, b.Flip(63))
	assert.NoError(t, b.Flip(31))
	assert.NoError(t, b.Flip(0))

	assert.Equal(t, "8000000080000001-1", b.String())
}

func TestBitArray_Flip(t *testing.T) {
	b := NewBitArray(64)

	assert.NoError(t, b.Flip(63))
	assert.NoError(t, b.Flip(31))
	assert.NoError(t, b.Flip(0))

	assert.Equal(t, "8000000080000001", b.String())
}

func TestBitArray_At(t *testing.T) {
	b := NewBitArray(64)

	assert.NoError(t, b.Flip(31))

	if at, err := b.At(31); assert.NoError(t, err) {
		assert.True(t, at)
	}

	if at, err := b.At(0); assert.NoError(t, err) {
		assert.False(t, at)
	}

	if at, err := b.At(63); assert.NoError(t, err) {
		assert.False(t, at)
	}
}
