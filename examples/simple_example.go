package examples

import (
	"fmt"

	"github.com/emin/skiplist"
)

func SimpleUsage() {
	list := skiplist.New()
	list.Set([]byte("test-key-1"), []byte("1-data"))
	list.Set([]byte("test-key-2"), []byte("2-data"))
	v := list.Get([]byte("test-key-2"))
	fmt.Println(v)
}
