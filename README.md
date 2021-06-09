## Skip List Implementation in Go

[![Build Status](https://travis-ci.com/emin/skiplist.svg?branch=master)](https://travis-ci.com/emin/skiplist)
[![codecov](https://codecov.io/gh/emin/skiplist/branch/master/graph/badge.svg?token=G3GTH9KRRN)](https://codecov.io/gh/emin/skiplist)
[![Dependency Status](https://img.shields.io/badge/dependencies-none-green)](https://github.com/emin/skiplist/blob/master/go.mod)

### What is it ?
Skip list is a probabilistic data structure which keeps its elements in order. It behaves like sorted map. 
It provides O(log n) insertion, deletion and search time. For more look at [Wikipedia](https://en.wikipedia.org/wiki/Skip_list)


#### Simple Usage

```go
package main

import (
	"fmt"
	"github.com/emin/skiplist"
)

func main(){
    // use string keys only
    list := skiplist.New(skiplist.StringComparator)
    list.Set("key-1", []byte("byte slice data"))
    list.Set("key-2", "string data")
    list.Set("key-3", 999)
    v := list.Get("key-2")
    fmt.Println(v)
    if list.Delete("key-2") {
        fmt.Println("key-2 is deleted")
    }
}
```

#### Iterating over skip list

```go
package main

import (
	"fmt"
	"github.com/emin/skiplist"
)

func main(){
    // use int keys only
    list := skiplist.New(skiplist.IntComparator)
    list.Set(1, "data-1")
    list.Set(3, "data-3")
    list.Set(2, "data-2")
    
    it := list.Iterator()
    for it.Next() {
    	fmt.Printf("%v = %v \n", it.Key(), it.Value())
    }
}
```

For more examples, look at `examples/` folder in the project.


#### Skiplist Interface

```go
New(comparator func(interface{}, interface{}) int) *SkipList
Get(key interface{}) interface{}
Delete(key interface{}) bool
Set(key, value interface{})
KeyCount() int64
Iterator() Iterator

```

For more, look at [![Go Reference](https://pkg.go.dev/badge/github.com/emin/skiplist.svg)](https://pkg.go.dev/github.com/emin/skiplist)

#### Benchmarks

Run on M1 Macbook Air 8GB model.
```
go test -bench=.
goos: darwin
goarch: arm64
pkg: github.com/emin/skiplist
BenchmarkSetString-8   	 1655217	       760.4 ns/op	     605 B/op	       9 allocs/op
BenchmarkSetInt-8      	 2185052	       590.1 ns/op	     608 B/op	      10 allocs/op
BenchmarkGetString-8   	 3935841	       305.6 ns/op	      16 B/op	       1 allocs/op
BenchmarkGetInt-8      	 8904025	       160.9 ns/op	       7 B/op	       0 allocs/op
BenchmarkGetHash-8     	20032552	        51.57 ns/op	       0 B/op	       0 allocs/op
```
At the last row `BenchmarkGetHash-8` is `map[string]string`, I compared the numbers with hashtable performance while developing the code.
Decided to left it as something to compare for you. Don't forget hashtable has O(1) access, skip list is O(log n).

#### License

MIT License