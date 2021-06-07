package skiplist

import (
	"testing"
)

func createTestSkipList() *SkipList {
	l := New(StringComparator)
	l.Set("baaa", "test_b2")
	l.Set("aaab", "test_b1")
	l.Set("aaad", "test_d1")
	l.Set("aaaa", "test_a1")
	l.Set("aaac", "test_c1")
	return l
}

func TestIteratorImpl_Next(t *testing.T) {
	eL := newIterator(New(StringComparator))
	assertEq(t, false, eL.Next())

	l := createTestSkipList()
	it := newIterator(l)
	assertEq(t, true, it.Next())
	assertEq(t, true, it.Next())
	assertEq(t, true, it.Next())
	assertEq(t, true, it.Next())
	assertEq(t, true, it.Next())
	assertEq(t, false, it.Next())
}

func TestIteratorImpl_Key(t *testing.T) {
	l := createTestSkipList()
	it := newIterator(l)
	it.Next()
	assertEq(t, it.Key(), "aaaa")
	it.Next()
	assertEq(t, it.Key(), "aaab")
	it.Next()
	assertEq(t, it.Key(), "aaac")
	it.Next()
	assertEq(t, it.Key(), "aaad")
	it.Next()
	assertEq(t, it.Key(), "baaa")
	it.Next()
	assertEq(t, it.Key(), nil)
}

func TestIteratorImpl_Value(t *testing.T) {
	l := createTestSkipList()
	it := newIterator(l)
	it.Next()
	assertEq(t, it.Value(), "test_a1")
	it.Next()
	assertEq(t, it.Value(), "test_b1")
	it.Next()
	assertEq(t, it.Value(), "test_c1")
	it.Next()
	assertEq(t, it.Value(), "test_d1")
	it.Next()
	assertEq(t, it.Value(), "test_b2")
	it.Next()
	assertEq(t, it.Value(), nil)

}
