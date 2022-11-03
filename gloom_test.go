package gloom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter_Insert(t *testing.T) {
	testCases := []struct {
		name        string
		filterSize  uint64
		numHashes   uint64
		errExpected bool
	}{
		{
			"m = 64, k = 0",
			64,
			0,
			true,
		},
		{
			"m = 64, k = 1",
			64,
			1,
			false,
		},
		{
			"m = 64, k = 7",
			64,
			7,
			false,
		},
		{
			"m = 256, k = 3",
			256,
			3,
			false,
		},
		{
			"m = 256, k = 5",
			256,
			5,
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			f, err := NewFilter(tc.filterSize, tc.numHashes)
			if tc.errExpected {
				assert.Error(t, err)
				return
			}

			if assert.NoError(t, f.Insert("hello")) {
				if res, err := f.Contains("hello"); assert.NoError(t, err) {
					assert.True(t, res)
				}
				if res, err := f.Contains("world"); assert.NoError(t, err) {
					assert.False(t, res)
				}
			}
		})
	}
}
