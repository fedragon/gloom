package gloom

import (
	"strconv"
	"strings"
)

type BitArray []uint64

func NewBitArray(len int) BitArray {
	l := len / 64
	if len%64 > 0 {
		l++
	}
	return make(BitArray, l)
}

func (b BitArray) Flip(at int) {
	index, offset := at/64, at%64
	b[index] |= 1 << offset
}

func (b BitArray) At(at int) bool {
	index, offset := at/64, at%64
	return b[index]&(1<<offset) != 0
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
