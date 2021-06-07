/*
This is a skip list implementation in Go.
Skip list is a probabilistic data structure which stores elements in order.
It provides O(log n) insert and search complexity. (For more: https://en.wikipedia.org/wiki/Skip_list )
*/
package skiplist

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
)

const maxLevel = 32

// SkipList implementation
type SkipList struct {
	highestLevel int
	sentinel     *node
	count        int64
	comparator   func(interface{}, interface{}) int
}

type node struct {
	Key   interface{}
	Value interface{}
	Next  []*node
}

// ready to use comparators for common types
var ByteSliceComparator = func(a, b interface{}) int { return bytes.Compare(a.([]byte), b.([]byte)) }
var StringComparator = func(a, b interface{}) int { return strings.Compare(a.(string), b.(string)) }
var IntComparator = func(a, b interface{}) int { return a.(int) - b.(int) }
var Int64Comparator = func(a, b interface{}) int { return int(a.(int64) - b.(int64)) }

// Returns a new SkipList
// Requires comparison function for the key type which will be used in skip list
//
// Example comparator;
//
//	 func(a, b interface{}) int {
//			return bytes.Compare(a.([]byte), b.([]byte))
// 	}
//
func New(comparator func(interface{}, interface{}) int) *SkipList {
	l := &SkipList{
		highestLevel: 0,
		comparator:   comparator,
	}
	l.sentinel = &node{
		Next: make([]*node, maxLevel),
	}
	return l
}

func flipCoin() bool {
	return rand.Intn(2) == 1
}

// Return value for given key
// if key does not exists, it returns nil
func (list *SkipList) Get(key interface{}) interface{} {
	h := list.highestLevel
	cur := list.sentinel

	for ; h >= 0; h-- {
		for cur != nil {
			if cur.Next[h] != nil {
				if cmp := list.comparator(cur.Next[h].Key, key); cmp >= 0 {
					if cmp == 0 {
						return cur.Next[h].Value
					}
					break
				}
				cur = cur.Next[h]
			} else {
				break
			}
		}
	}

	if cur != nil && cur.Key != nil && list.comparator(cur.Key, key) == 0 {
		return cur.Value
	}

	return nil
}

// Deletes key from SkipList
func (list *SkipList) Delete(key interface{}) bool {
	h := list.highestLevel
	cur := list.sentinel

	removed := false
	for ; h >= 0; h-- {
		for cur != nil {
			if cur.Next[h] != nil {
				if cmp := list.comparator(cur.Next[h].Key, key); cmp >= 0 {
					if cmp == 0 {
						removed = true
						if len(cur.Next[h].Next) > h {
							cur.Next[h] = cur.Next[h].Next[h]
						} else {
							cur.Next[h] = nil
						}
						if cur == list.sentinel && cur.Next[h] == nil && h > 0 {
							list.highestLevel--
						}

					}
					break
				}
				cur = cur.Next[h]
			} else {
				break
			}

		}
	}
	if removed {
		list.count--
		return true
	}

	return false
}

// Sets value for key in SkipList
func (list *SkipList) Set(key, value interface{}) {
	h := list.highestLevel
	cur := list.sentinel

	stack := make([]*node, 0)

	for ; h >= 0; h-- {
		for cur != nil {
			if h >= len(cur.Next) || cur.Next[h] == nil || list.comparator(cur.Next[h].Key, key) >= 0 {
				stack = append(stack, cur)
				break
			} else {
				cur = cur.Next[h]
			}
		}
	}

	if cur.Next[0] != nil && list.comparator(cur.Next[0].Key, key) == 0 {
		cur.Next[0].Value = value
	} else {
		list.count++
		newNode := &node{
			Key:   key,
			Value: value,
			Next:  make([]*node, 1),
		}
		newNode.Next[0] = cur.Next[0]
		cur.Next[0] = newNode

		for j := 1; flipCoin() && j < maxLevel; j++ {

			if j > list.highestLevel {
				list.highestLevel++
			}
			lNode := list.sentinel
			if j < len(stack) {
				lNode = stack[len(stack)-j-1]
			}

			var nextNode *node = nil
			if j < len(lNode.Next) {
				nextNode = lNode.Next[j]
			}

			if j >= len(newNode.Next) {
				newNode.Next = append(newNode.Next, nextNode)
			} else {
				newNode.Next[j] = nextNode
			}

			if j >= len(lNode.Next) {
				lNode.Next = append(lNode.Next, newNode)
			} else {
				lNode.Next[j] = newNode
			}

		}
	}
}

// Returns how many keys currently in SkipList
func (list *SkipList) KeyCount() int64 {
	return list.count
}

// This is for debug purposes, don't use it with big sizes
func (list *SkipList) debugPrint() {
	fmt.Printf("============level = %v================\n", list.highestLevel)

	for h := 0; h <= list.highestLevel; h++ {
		cur := list.sentinel.Next[h]
		for cur != nil {
			fmt.Printf("%v ", cur.Key)
			cur = cur.Next[h]
		}
		fmt.Println("")
	}
	//fmt.Println("=============================")
}

// Returns an iterator for SkipList
func (list *SkipList) Iterator() Iterator {
	return newIterator(list)
}
