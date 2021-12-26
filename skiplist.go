/*
This is a skip list implementation in Go.
Skip list is a probabilistic data structure which stores elements in order.
It provides O(log n) insert and search complexity. (For more: https://en.wikipedia.org/wiki/Skip_list )
*/
package skiplist

import (
	"bytes"
	"math/rand"
)

const maxLevel = 32

// SkipList implementation
type SkipList struct {
	highestLevel int
	sentinel     *node
	count        int64
	size         int64
}

type node struct {
	Key   []byte
	Value []byte
	Next  []*node
}

// Returns a new SkipList
func New() *SkipList {
	l := &SkipList{
		highestLevel: 0,
	}
	root := make([]*node, maxLevel)
	l.sentinel = &node{
		Next: root,
	}
	return l
}

func flipCoin() bool {
	return rand.Intn(2) == 1
}

// Return value for given key
// if key does not exists, it returns nil
func (list *SkipList) Get(key []byte) []byte {
	h := list.highestLevel
	cur := list.sentinel

	for ; h >= 0; h-- {
		for cur != nil {
			if cur.Next[h] != nil {
				if cmp := bytes.Compare(cur.Next[h].Key, key); cmp >= 0 {
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

	return nil
}

// Deletes key from SkipList
func (list *SkipList) Delete(key []byte) bool {
	h := list.highestLevel
	cur := list.sentinel

	removed := false
	for ; h >= 0; h-- {
		for cur != nil {
			if cur.Next[h] != nil {
				if cmp := bytes.Compare(cur.Next[h].Key, key); cmp >= 0 {
					if cmp == 0 {
						removed = true
						cur.Next[h] = cur.Next[h].Next[h]
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
func (list *SkipList) Set(key []byte, value []byte) {
	h := list.highestLevel
	cur := list.sentinel

	stack := make([]*node, 0, list.highestLevel)

	for ; h >= 0; h-- {
		for cur != nil {
			if h >= len(cur.Next) || cur.Next[h] == nil || bytes.Compare(cur.Next[h].Key, key) >= 0 {
				stack = append(stack, cur)
				break
			} else {
				cur = cur.Next[h]
			}
		}
	}

	if cur.Next[0] != nil && bytes.Compare(cur.Next[0].Key, key) == 0 {
		list.size -= int64(len(cur.Next[0].Value))
		cur.Next[0].Value = value
		list.size += int64(len(value))
	} else {
		list.count++
		list.size += int64(len(value) + len(key))
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

			newNode.Next = append(newNode.Next, nextNode)
			lNode.Next[j] = newNode
		}
	}
}

// Returns how many keys currently in SkipList
func (list *SkipList) KeyCount() int64 {
	return list.count
}

// Returns an iterator for SkipList
func (list *SkipList) Iterator() Iterator {
	return newIterator(list)
}

func (list *SkipList) RawSize() int64 {
	return list.size
}
