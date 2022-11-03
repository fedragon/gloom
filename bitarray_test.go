package gloom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitArray_NewBitArray_UpTo64(t *testing.T) {
	var capacity uint64 = 64
	b := NewBitArray(capacity)

	assert.Equal(t, capacity, b.Len())
}

func TestBitArray_NewBitArray_MoreThan64(t *testing.T) {
	var capacity uint64 = 65
	b := NewBitArray(capacity)

	assert.Equal(t, capacity, b.Len())
}

func TestBitArray_Flip(t *testing.T) {
	b := NewBitArray(244)

	assert.NoError(t, b.Flip(243))
	assert.NoError(t, b.Flip(121))
	assert.NoError(t, b.Flip(66))
	assert.NoError(t, b.Flip(31))
	assert.NoError(t, b.Flip(0))

	assert.Equal(t, "80000001-200000000000004-0-8000000000000", b.String())
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
