## SkipList Implementation in Go

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
    // use string keys only
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

