package gloom

import (
	"fmt"
	"strconv"
	"strings"
)

type BitArray []uint64

func NewBitArray(len uint64) BitArray {
	l := len / 64
	if len%64 > 0 {
		l++
	}
	return make(BitArray, l)
}

func (b BitArray) Flip(at int) error {
	if at >= b.Len() {
		return fmt.Errorf("invalid position: %v (out of %v)", at, len(b))
	}
	index, offset := at/64, at%64
	b[index] |= 1 << offset

	return nil
}

func (b BitArray) At(at int) (bool, error) {
	if at >= b.Len() {
		return false, fmt.Errorf("invalid position: %v (out of %v)", at, len(b))
	}

	index, offset := at/64, at%64
	return b[index]&(1<<offset) != 0, nil
}

func (b BitArray) String() string {
	var buf strings.Builder
	for i, v := range b {
		buf.WriteString(strconv.FormatUint(v, 16))
		if i < len(b)-1 {
			buf.WriteString("-")
		}
	}

	return buf.String()
}

func (b BitArray) Len() int {
	return len(b) * 64
}
