# gloom

Bloom filters implementation in Go.

## Context

Bloom filters are probabilistic data structures with a predetermined size and false positives' rate. They are space efficient and have a fixed `O(k)` cost to add an item or check if it's present, where `k` is the number of hashing functions used by the filter.

## Usage

```go
import "github.com/fedragon/gloom"

// create a bloom filter which can store 10.000 elements with a 1% rate of false positives
bf := gloom.NewFilter(10000, 0.01)

// insert a value in the filter
if err := bf.Insert("a"); err != nil {
	return err
}

...

// is this value in the filter?
if ok, err := bf.Contains("b"); err != nil {
	return err
	
	// if so, do something
	if ok {
		...
	}
}
```
