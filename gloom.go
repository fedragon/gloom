package gloom

import (
	"errors"
	"math"

	"github.com/spaolacci/murmur3"
)

// Filter represents a Bloom Filter with `m` elements and `k` number of hashes.
type Filter struct {
	elems *BitArray
	m     uint64 // size
	k     uint64 // number of hashing functions
}

// NewFilter returns a bloom filter which can store up to `n` elements with an approximate false positives' rate `p`.
// It returns an error if any of the inputs is invalid.
// Empirical tests show that, given a desired rate of 0.05, the actual rate is 0.01 <= x <= 0.08
func NewFilter(n uint64, p float64) (*Filter, error) {
	if n == 0 {
		return nil, errors.New("n cannot be zero")
	}

	if p <= 0 {
		return nil, errors.New("p must be greater than zero")
	}

	// source: https://en.wikipedia.org/wiki/Bloom_filter#Optimal_number_of_hash_functions
	m := math.Abs((float64(n) * math.Log(p)) / math.Pow(math.Ln2, 2))
	k := math.Abs(math.Log2(p))

	return newFilter(uint64(math.Ceil(m)), uint64(math.Ceil(k)))
}

func newFilter(size, numHashes uint64) (*Filter, error) {
	if size == 0 {
		return nil, errors.New("size cannot be zero")
	}

	if numHashes == 0 {
		return nil, errors.New("numHashes cannot be zero")
	}

	return &Filter{
		elems: NewBitArray(size),
		m:     size,
		k:     numHashes,
	}, nil
}

// Insert inserts a new entry in the bloom filter
func (f *Filter) Insert(value string) error {
	h1, h2 := hash([]byte(value))

	for i := uint64(0); i < f.k; i++ {
		h := nthHash(h1, h2, i, f.m)

		if err := f.elems.Flip(h); err != nil {
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

		if exists, err := f.elems.At(h); err != nil || !exists {
			return false, err
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
