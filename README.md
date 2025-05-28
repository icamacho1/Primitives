# Description ![main](https://github.com/icamacho1/Primitives/actions/workflows/pr.yml/badge.svg?branch=main)
Primitives is a Go library that provides enhanced functionality for basic data
types (slices, maps, and sets) with a focus on expressive and efficient data manipulation.
The library is built using Go's generics, making it type-safe and
flexible for different data types.

# Main Components
## 1. Slices Package

The slices package provides an enhanced version of Go's slice type with additional utility methods:

```go
    import "github.com/icamacho1/Primitives/pkg/slices"

    // Create a new slice
    s := slices.New("foo", "bar", "baz")

    // Basic operations
    s.Append("qux")           // Add elements
    s.Pop(1)                 // Remove element at index
    first := s.First()       // Get first element
    last := s.Last()         // Get last element
    length := s.Len()        // Get length
    exists := s.Has("foo")   // Check if element exists

    // Transformations
    s.Map(func(v string, _ int) string {
        return v + " modified"
    })

    // Get unique elements
    unique, hasDuplicates := s.Uniq()
```

## 2. Maps Package
The maps package provides an enhanced version of Go's map type with additional utility methods:

```go
    import "github.com/icamacho1/Primitives/pkg/maps"

    // Create a new map
    m := maps.New[string, int]()

    // Basic operations
    m.Add("one", 1)          // Add key-value pair
    value, exists := m.Get("one")  // Get value with existence check
    value = m.MustGet("one")       // Get value (panics if not found)
    exists = m.Has("one")          // Check if key exists
    m.Pop("one")                   // Remove key-value pair

    // Get all keys
    keys := m.Keys()
```

## 3. Sets Package
The sets package provides an ordered set implementation that maintains insertion order:

```go
    import "github.com/icamacho1/Primitives/pkg/sets"

    // Create a new ordered set
    s := sets.NewOrdered("foo", "bar", "baz")

    // Basic operations
    s.Append("qux")          // Add elements
    s.Pop("bar")            // Remove element
    first := s.First()      // Get first element
    last := s.Last()        // Get last element
    exists := s.Has("foo")  // Check if element exists
    element := s.Get(1)     // Get element at index
```
