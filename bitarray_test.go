package gloom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitArray_NewBitArray_UpTo64(t *testing.T) {
	b := NewBitArray(64)

	b.Flip(63)
	b.Flip(31)
	b.Flip(0)

	assert.Equal(t, "8000000080000001", b.String())
}

func TestBitArray_NewBitArray_MoreThan64(t *testing.T) {
	b := NewBitArray(65)

	b.Flip(64)
	b.Flip(63)
	b.Flip(31)
	b.Flip(0)

	assert.Equal(t, "8000000080000001-1", b.String())
}

func TestBitArray_Flip(t *testing.T) {
	b := NewBitArray(64)

	b.Flip(63)
	b.Flip(31)
	b.Flip(0)

	assert.Equal(t, "8000000080000001", b.String())
}

func TestBitArray_At(t *testing.T) {
	b := NewBitArray(64)

	b.Flip(31)

	assert.True(t, b.At(31))
	assert.False(t, b.At(0))
	assert.False(t, b.At(63))
}
