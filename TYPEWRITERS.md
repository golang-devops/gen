## Typewriters

This is a list of known open-source typewriters in alphabetical order. 
GoDoc links are to the original implementation that might use `interface{}` instead of a gen type.
Typewriters that would benefit from some improvement are listed at the bottom.
Please add your own by making a pull request. Make sure you have documented the code before submitting a pull request.


`gen add github.com/nickmab/gen/typewriters/container`

```go
// +gen container:"Heap, HeapBy"
type MyType struct{}
```

Implements a strongly-typed heap, based on [golang.org/pkg/container/heap](https://golang.org/pkg/container/heap). A heap is a tree with the property that each node is the minimum-valued node in its subtree. Useful implementation of a priority queue. 


`gen add github.com/clipperhouse/linkedlist`

```go
// +gen list
type MyType struct{}
```

Implements a strongly-typed, doubly-linked list, based on [golang.org/pkg/container/list](https://golang.org/pkg/container/list). 


`gen add github.com/ninibe/atomicmapper/gen`

```go
// +gen atomicmap
type MyType struct{}
```

Atomicmapper is a code generation tool for creating high-performance, scalable, frequently read, but infrequently updated maps of strings to any given `map[string]MyType`. It is based on Go's [atomic.Value read mostly example](https://golang.org/pkg/sync/atomic/#example_Value_readMostly).


`gen add github.com/clipperhouse/ring`

```go
// +gen ring
type MyType struct{}
```

Implements strongly-typed operations on circular lists, based on [golang.org/pkg/container/ring](https://golang.org/pkg/container/ring). 


`gen add github.com/clipperhouse/set`  

```go
// +gen set
type MyType struct{}
```
Implements a strongly-typed unordered set with unique values, based on [github.com/deckarep/golang-set](https://github.com/deckarep/golang-set).


`gen add github.com/jackc/signal`  

```go
// +gen signal
type MyType struct{}
```
Implements the signal pattern. [github.com/jackc/signal](https://github.com/jackc/signal).


`github.com/clipperhouse/slice` `built-in typewriter, no need to install`  

```go
// +gen slice:"Aggregate[T], All, Any, Average, Average[T], Count, Distinct, DisctinctBy, First, GroupBy[T], Max, Max[T], Min, Min[T], MinBy, Select[T], Shuffle, Sort, SortBy, Where"
type MyType struct{}
```
Generates functional convenience methods that will look familiar to users of C#’s LINQ or JavaScript’s Array methods. It is intended to save you some loops, using a “pass a function” pattern. It offers easier ad-hoc sorts. Documentation is available at [clipperhouse.github.io/gen/slice](https://golang-devops.github.io/gen/slice/).


`gen add github.com/svett/gen/stack`

```go
// +gen stack
type MyType struct{}
```
Implements a Stack.


`github.com/clipperhouse/stringer` `built-in typewriter, no need to install`  

```go
// +gen stringer
type Pill int

const (
    Placebo Pill = iota
    Aspirin
    Ibuprofen
    Paracetamol
    Acetaminophen = Paracetamol
)
```
Generates an implementation of the Stringer interface from const variable names, for pretty-printing, based on  [golang.org/x/tools/cmd/stringer](https://golang.org/x/tools/cmd/stringer).



`gen add github.com/ggaaooppeenngg/queue` 

```go
// +gen queue
type MyType struct{}
```
Implements a queue.



Feel free to help these typewriters by making installation easier, improving documentation, writing tests or improving the implementation.


`gen add github.com/michaelsmanley/flags`  

```go
// +gen flags"
type MyType struct{}
```
Convenience functions for manipulating flags, in the bitset sense. 


`gen add github.com/fernandokm/sequences`  

```go
// +gen sequenceGenerator:"Iterator[int64,uint64,*bigInt]"
type MyType struct{}
```
Generates an implementation of a generator of primes, fibonacci or triangular numbers using a given type.


##### Slice extension [![GoDoc](https://godoc.org/github.com/remz/golang-sdk/gen_extras?status.svg)](https://godoc.org/github.com/remz/golang-sdk/gen_extras)
`github.com/remz/golang-sdk/gen_extras` `extends the built-in slice implementation, adding new functions`  

```go
// These are the new functions added by this slice extension:
// +gen slice:"AddOnce, Contains, Cut, Expand, Extend, FillRange, Find, Insert, InsertMultiple, IsEqualTo, MakeCopy, Mapping, Pop, Push, Remove, RemoveValue, ZeroUpTo"
type MyType struct{}
```
Extends the Slice implementation by adding some more functions. Current implementation requires copying files into the Slice implementation folder and rebuilding gen. 

