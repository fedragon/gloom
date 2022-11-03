package gloom

import (
	"errors"

	"github.com/spaolacci/murmur3"
)

// Filter represents a Bloom Filter with `m` elements and `k` number of hashes.
type Filter struct {
	bits *BitArray
	m    uint64
	k    uint64
}

// NewFilter returns a filter with provided parameters, or an error if the inputs are invalid.
func NewFilter(size, numHashes uint64) (*Filter, error) {
	if numHashes == 0 {
		return nil, errors.New("I need at least one hash to work")
	}

	return &Filter{
		bits: NewBitArray(size),
		m:    size,
		k:    numHashes,
	}, nil
}

// Insert inserts a new entry in the bloom filter
func (f *Filter) Insert(value string) error {
	h1, h2 := hash([]byte(value))

	for i := uint64(0); i < f.k; i++ {
		h := nthHash(h1, h2, i, f.m)

		if err := f.bits.Flip(h); err != nil {
			return err
		}
	}

	return nil
}

// Contains checks if the bloom filter (possibly) contains the value
func (f *Filter) Contains(value string) (bool, error) {
	h1, h2 := hash([]byte(value))

	for i := uint64(0); i < f.k; i++ {
		h := nthHash(h1, h2, i, f.m)

		if exists, err := f.bits.At(h); err != nil || !exists {
			return exists, err
		}
	}

	return true, nil
}

func hash(value []byte) (uint64, uint64) {
	return murmur3.Sum128(value)
}

// nthHash implements double hashing (see https://en.wikipedia.org/wiki/Double_hashing), so that I can use a single hash function
func nthHash(h1, h2, n, size uint64) uint64 {
	return (h1 + n*h2) % size
}
