package skiplist

import (
	"testing"
)

func createTestSkipList() *SkipList {
	l := New()
	l.Set([]byte("baaa"), []byte("test_b2"))
	l.Set([]byte("aaab"), []byte("test_b1"))
	l.Set([]byte("aaad"), []byte("test_d1"))
	l.Set([]byte("aaaa"), []byte("test_a1"))
	l.Set([]byte("aaac"), []byte("test_c1"))
	return l
}

func TestIteratorImpl_Next(t *testing.T) {
	eL := newIterator(New())
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
	assertEq(t, it.Key(), []byte("aaaa"))
	it.Next()
	assertEq(t, it.Key(), []byte("aaab"))
	it.Next()
	assertEq(t, it.Key(), []byte("aaac"))
	it.Next()
	assertEq(t, it.Key(), []byte("aaad"))
	it.Next()
	assertEq(t, it.Key(), []byte("baaa"))
	it.Next()
	assertEq(t, nil, it.Key())
}

func TestIteratorImpl_Value(t *testing.T) {
	l := createTestSkipList()
	it := newIterator(l)
	it.Next()
	assertEq(t, it.Value(), []byte("test_a1"))
	it.Next()
	assertEq(t, it.Value(), []byte("test_b1"))
	it.Next()
	assertEq(t, it.Value(), []byte("test_c1"))
	it.Next()
	assertEq(t, it.Value(), []byte("test_d1"))
	it.Next()
	assertEq(t, it.Value(), []byte("test_b2"))
	it.Next()
	assertEq(t, nil, it.Value())

}
