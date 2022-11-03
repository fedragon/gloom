package gloom

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestFilter_NewFilter(t *testing.T) {
	testCases := []struct {
		name        string
		size        uint64
		errorRate   float64
		errExpected bool
	}{
		{
			"returns an error when size is zero",
			0,
			0.05,
			true,
		},
		{
			"returns an error when error rate is negative",
			0,
			-0.05,
			true,
		},
		{
			"returns an error when error rate is zero",
			0,
			0,
			true,
		},
		{
			"succeeds otherwise",
			10,
			0.05,
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewFilter(tc.size, tc.errorRate)
			if tc.errExpected {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestFilter_Insert(t *testing.T) {
	testCases := []struct {
		name        string
		filterSize  uint64
		numHashes   uint64
		errExpected bool
	}{
		{
			"returns an error when numHashes is zero",
			0,
			1,
			true,
		},
		{
			"returns an error when numHashes is zero",
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
			f, err := newFilter(tc.filterSize, tc.numHashes)
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

func TestFilter_PositivesRate(t *testing.T) {
	n := 1000
	p := 0.1
	tolerance := 0.04

	for x := 0; x < 100; x++ {
		var falsePositives float64 = 0
		f, err := NewFilter(uint64(n), p)
		assert.NoError(t, err)

		members := make(map[string]struct{})

		for i := 0; i < n; i++ {
			value := randString(30)
			members[value] = struct{}{}

			assert.NoError(t, f.Insert(value))
			if contains, err := f.Contains(value); assert.NoError(t, err) {
				assert.True(t, contains)
			}
		}

		total := 0
		for i := 0; i < n; i++ {
			value := randString(30)

			if _, exists := members[value]; exists {
				continue
			}

			total++

			if contains, err := f.Contains(value); assert.NoError(t, err) {
				if contains {
					falsePositives++
				}
			}
		}

		assert.LessOrEqual(t, falsePositives/float64(total), p+tolerance)
	}
}

func randString(size int) string {
	pool := "abcdefghijklmnopqrstuvwyz0123456789"

	buf := strings.Builder{}
	for i := 0; i < size; i++ {
		n := rand.Intn(len(pool))
		buf.WriteRune(rune(pool[n]))
	}

	return buf.String()
}
