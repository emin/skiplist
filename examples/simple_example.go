package examples

import (
	"fmt"
	"github.com/emin/skiplist"
	"io/ioutil"
	"log"
	"regexp"
)

func SimpleUsage() {
	list := skiplist.New(skiplist.StringComparator)
	list.Set("test-key-1", []byte("1-data"))
	list.Set("test-key-2", []byte("2-data"))
	v := list.Get("test-key-2")
	fmt.Println(v)
}

func WordOccurrence() {
	content, err := ioutil.ReadFile("examples/adventures_of_sherlock_holmes.txt")
	if err != nil {
		log.Fatal(err)
	}

	list := skiplist.New(skiplist.StringComparator)
	words := regexp.MustCompile("\\s+").Split(string(content), -1)
	for _, w := range words {
		c := list.Get(w)
		count := 1
		if c != nil {
			count = c.(int) + 1
		}
		list.Set(w, count)
	}

	it := list.Iterator()
	for it.Next() {
		fmt.Printf("%v = %v\n", it.Key(), it.Value())
	}
	fmt.Println(list.KeyCount())
}
