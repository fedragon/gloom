package gloom

import (
	"fmt"
	"strconv"
	"strings"
)

// BitArray represents an array of bits.
type BitArray struct {
	buckets []uint64
	len     uint64
}

// NewBitArray returns a new instance with the requested length.
func NewBitArray(len uint64) *BitArray {
	l := len / 64
	if len%64 > 0 {
		l++
	}

	return &BitArray{
		buckets: make([]uint64, l),
		len:     len,
	}
}

// Flip sets the bit at the `at` position to 1. It returns an error if `at` is greater or equal to the array's length.
func (b *BitArray) Flip(at uint64) error {
	if at >= b.Len() {
		return fmt.Errorf("invalid position: %v (out of %v)", at, b.Len())
	}
	index, offset := at/64, at%64
	b.buckets[index] |= 1 << offset

	return nil
}

// At returns true if the bit at the `at` position is set to 1, false otherwise. It returns an error if `at` is greater or equal to the array's length.
func (b *BitArray) At(at uint64) (bool, error) {
	if at >= b.Len() {
		return false, fmt.Errorf("invalid position: %v (out of %v)", at, b.Len())
	}

	index, offset := at/64, at%64
	return b.buckets[index]&(1<<offset) != 0, nil
}

// Len returns the lenght of the array.
func (b *BitArray) Len() uint64 {
	return b.len
}

func (b *BitArray) String() string {
	var buf strings.Builder
	for i, v := range b.buckets {
		buf.WriteString(strconv.FormatUint(v, 16))
		if i < len(b.buckets)-1 {
			buf.WriteString("-")
		}
	}

	return buf.String()
}
