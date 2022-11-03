# gloom

Bloom filters implementation in Go.

## Context

Bloom filters are probabilistic data structures with a predetermined size and false positives' rate. They are very space efficient and have a fixed `O(k)` cost to add an item or check if it's present, where `k` is the number of hashing functions used by the filter.

## Usage

```go
import "github.com/fedragon/gloom"

bf := gloom.NewFilter(10000, 0.01 )
if err := bf.Insert("value"); err != nil {
	return err
}

if ok, err := bf.Contains("value"); err != nil {
	return err
}
```