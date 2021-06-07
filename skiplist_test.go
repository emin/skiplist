package skiplist

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func equalCheck(expected interface{}, actual interface{}) bool {
	actualV := actual
	if (reflect.ValueOf(actual).Kind() == reflect.Slice ||
		reflect.ValueOf(actual).Kind() == reflect.Ptr) &&
		reflect.ValueOf(actual).IsNil() {
		actualV = nil
	}
	return reflect.DeepEqual(expected, actualV)
}

func assertEq(t *testing.T, expected interface{}, actual interface{}) {
	if !equalCheck(expected, actual) {
		_, file, line, ok := runtime.Caller(1)
		fileLoc := ""
		if ok {
			fileLoc = fmt.Sprintf("\nFailed at: %v:%v", file, line)
		}
		t.Errorf("%v\nExpected: %+v\n, Actual: %+v\n", fileLoc, expected, actual)
	}
}

func assertNotEq(t *testing.T, expected interface{}, actual interface{}) {
	if equalCheck(expected, actual) {
		_, file, line, ok := runtime.Caller(1)
		fileLoc := ""
		if ok {
			fileLoc = fmt.Sprintf("\nFailed at: %v:%v", file, line)
		}
		t.Errorf("%v\nExpected: %+v\n, Actual: %+v\n", fileLoc, expected, actual)
	}
}

func TestNew(t *testing.T) {
	l := New(StringComparator)
	assertNotEq(t, nil, l)
}

func TestSet(t *testing.T) {
	l := New(StringComparator)
	k := "1"
	v := "test-val-1"
	l.Set(k, v)
	nV := l.Get(k)
	assertNotEq(t, nil, nV)
	assertEq(t, v, nV)

	v2 := "test-val-2"
	l.Set(k, v2)

	nV = l.Get(k)
	assertNotEq(t, nil, nV)
	assertEq(t, v2, nV)

	l.Delete(k)
	nV = l.Get(k)
	assertEq(t, nil, nV)

	l.Set(k, v)
	nV = l.Get(k)
	assertNotEq(t, nil, nV)
	assertEq(t, v, nV)

	l.Set(k, v2)
	nV = l.Get(k)
	assertNotEq(t, nil, nV)
	assertEq(t, v2, nV)

	l.Delete(k)
	nV = l.Get(k)
	assertEq(t, nil, nV)

}

func TestGet(t *testing.T) {
	l := New(StringComparator)
	for i := 0; i < 10099; i++ {
		k := fmt.Sprintf("test-key-%v", i)
		v := fmt.Sprintf("test-val-%v", i*2)
		l.Set(k, v)
		nV := l.Get(k)
		assertNotEq(t, nil, nV)
		assertEq(t, v, nV)
	}

}

func TestDelete(t *testing.T) {
	l := New(Int64Comparator)

	for i := 0; i < 100; i++ {
		var k int64 = 13
		l.Set(k, []byte("1"))
		r := l.Delete(k)
		assertEq(t, true, r)
		r = l.Delete(k)
		assertEq(t, false, r)
		assertEq(t, nil, l.Get(k))
		l.Set(k, []byte("1"))
		assertEq(t, []byte("1"), l.Get(k))
	}

	l = New(ByteSliceComparator)

	for i := 0; i < 10000; i++ {
		k := []byte(fmt.Sprintf("test-key-%v", i))
		v := []byte(fmt.Sprintf("test-val-%v", i))
		l.Set(k, v)
	}

	for i := 0; i < 5000; i++ {
		k := []byte(fmt.Sprintf("test-key-%v", i))
		removed := l.Delete(k)
		assertEq(t, true, removed)
		nV := l.Get(k)
		assertEq(t, nil, nV)
	}

	for i := 5000; i < 10000; i++ {
		k := []byte(fmt.Sprintf("test-key-%v", i))
		v := []byte(fmt.Sprintf("test-val-%v", i))
		nV := l.Get(k)
		assertEq(t, v, nV)
	}

}

func TestSkipList_KeyCount(t *testing.T) {
	l := New(IntComparator)
	l.Set(4, []byte("test"))
	l.Set(4, []byte("test"))
	assertEq(t, int64(1), l.KeyCount())
	l.Delete(4)
	assertEq(t, int64(0), l.KeyCount())
	l.Delete(4)
	assertEq(t, int64(0), l.KeyCount())
	count := 20210
	for i := 0; i < count; i++ {
		l.Set(i*10, []byte("test"))
	}
	assertEq(t, int64(count), l.KeyCount())
}

func BenchmarkSet(b *testing.B) {
	l := New(ByteSliceComparator)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		v := []byte(fmt.Sprintf("%v", n))
		l.Set(v, v)
	}
}

func BenchmarkGetByteSlice(b *testing.B) {
	l := New(ByteSliceComparator)
	for i := 0; i < 10000000; i++ {
		k := fmt.Sprintf("%v", i)
		v := fmt.Sprintf("val -> %v", i)
		l.Set([]byte(k), []byte(v))
	}
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		v := []byte(fmt.Sprintf("%v", n))
		l.Get(v)
	}
}

func BenchmarkGetString(b *testing.B) {
	l := New(StringComparator)
	for i := 0; i < 10000000; i++ {
		k := fmt.Sprintf("%v", i)
		v := fmt.Sprintf("val -> %v", i)
		l.Set(k, []byte(v))
	}
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		v := fmt.Sprintf("%v", n)
		l.Get(v)
	}
}

func BenchmarkGetInt(b *testing.B) {
	l := New(IntComparator)
	for i := 0; i < 10000000; i++ {
		v := fmt.Sprintf("val -> %v", i)
		l.Set(i, []byte(v))
	}
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		l.Get(n)
	}
}

func BenchmarkGetHash(b *testing.B) {
	l := make(map[string]string)
	for i := 0; i < 10000000; i++ {
		k := fmt.Sprintf("%v", i)
		v := fmt.Sprintf("val -> %v", i)
		l[k] = v
	}
	a := ""
	void(a)
	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		v := fmt.Sprintf("%v", n)
		a = l[v]
	}
}
func void(_ string) {
}

func TestSkipList_Iterator(t *testing.T) {
	l := New(ByteSliceComparator)
	assertNotEq(t, nil, l.Iterator())
}
