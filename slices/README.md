# slices

## SliceDedupe

Designed as a _generic_ function, so requires Go 1.18 or newer.

Expects to support: `byte`, `float32`, `float64`, `int`, `int16`, `int64`, `int8`, `rune`, `string`, `uint`, `uint16`, `uint32`, `uint64`, and `uintptr`.

Have written tests against: `float32`, `float64`, `int`, `int64`, `string`, `uint`, and `uint64`.

```go
import (
    "github.com/northwood-labs/golang-utils/slices"
)

func main() {
    mixed := []string{"woof", "moo", "quack", "quack", "bok", "bok", "baaah"}

    // Sorting is part of the dedupe process.
    sortedDeduped := slices.SliceDedupe(mixed)

    //=> sortedDeduped = []string{"baaah", "bok", "moo", "quack", "woof"}
}
```

## StringSliceToHashmap

```go
import (
    "github.com/northwood-labs/golang-utils/slices"
)

func main() {
    listOfThings := []string{"baaah", "bok", "moo", "quack", "woof"}

    // Flip the values into keys.
    lookupMap := slices.StringSliceToHashmap(listOfThings)

    if _, ok := lookupMap["moo"]; ok {
        // Do the thing.
    }
}
```
