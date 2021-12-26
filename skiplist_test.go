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
	l := New()
	assertNotEq(t, nil, l)
}

func TestSet(t *testing.T) {
	l := New()
	k := []byte("1")
	v := []byte("test-val-1")
	l.Set(k, v)
	nV := l.Get(k)
	assertNotEq(t, nil, nV)
	assertEq(t, v, nV)

	v2 := []byte("test-val-2")
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
	l := New()
	for i := 0; i < 10099; i++ {
		k := []byte(fmt.Sprintf("%v", i*10))
		v := []byte(fmt.Sprintf("test-val-%v", i*2))
		l.Set(k, v)
		nV := l.Get(k)
		assertNotEq(t, nil, nV)
		assertEq(t, v, nV)
	}

	for i := 0; i < 10099; i++ {
		k := []byte(fmt.Sprintf("%v", i*10))
		l.Delete(k)
	}

	for i := 0; i < 10099; i++ {
		k := []byte(fmt.Sprintf("%v", i*5))
		v := []byte(fmt.Sprintf("test-val-%v", i*2))
		l.Set(k, v)
		nV := l.Get(k)
		assertNotEq(t, nil, nV)
		assertEq(t, v, nV)
	}

}

func TestDelete(t *testing.T) {

	l := New()

	for i := 0; i < 10000; i++ {
		k := []byte(fmt.Sprintf("test-key-%v", i))
		v := []byte(fmt.Sprintf("test-val-%v", i))
		l.Set(k, v)
	}

	for i := 5000; i > 2000; i-- {
		k := []byte(fmt.Sprintf("test-key-%v", i))
		removed := l.Delete(k)
		assertEq(t, true, removed)
		nV := l.Get(k)
		assertEq(t, nil, nV)
	}

	for i := 9999; i > 8000; i-- {
		k := []byte(fmt.Sprintf("test-key-%v", i))
		removed := l.Delete(k)
		assertEq(t, true, removed)
		nV := l.Get(k)
		assertEq(t, nil, nV)
	}

	for i := 5001; i < 8000; i++ {
		k := []byte(fmt.Sprintf("test-key-%v", i))
		v := []byte(fmt.Sprintf("test-val-%v", i))
		nV := l.Get(k)
		assertEq(t, v, nV)
	}

}

func TestSkipList_KeyCount(t *testing.T) {
	l := New()
	l.Set([]byte("4"), []byte("test"))
	l.Set([]byte("4"), []byte("test"))
	assertEq(t, int64(1), l.KeyCount())
	l.Delete([]byte("4"))
	assertEq(t, int64(0), l.KeyCount())
	l.Delete([]byte("4"))
	assertEq(t, int64(0), l.KeyCount())
	count := 20210
	for i := 0; i < count; i++ {
		k := []byte(fmt.Sprintf("%v", i*10))
		l.Set(k, []byte("test"))
	}
	assertEq(t, int64(count), l.KeyCount())
}

func BenchmarkSetString(b *testing.B) {
	l := New()
	keys := make([][]byte, b.N)
	for n := 0; n < b.N; n++ {
		keys[n] = []byte(fmt.Sprintf("%v", n))
	}
	b.ReportAllocs()
	b.ResetTimer()
	n1 := []byte("s2")
	for n := 0; n < b.N; n++ {
		l.Set(keys[n], n1)
	}
}

func BenchmarkGetString(b *testing.B) {
	l := New()
	for i := 0; i < 1000_000; i++ {
		k := fmt.Sprintf("%v", i)
		v := fmt.Sprintf("val -> %v", i)
		l.Set([]byte(k), []byte(v))
	}
	keys := make([][]byte, b.N)
	for n := 0; n < b.N; n++ {
		keys[n] = []byte(fmt.Sprintf("%v", n))
	}
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		l.Get(keys[n])
	}
}

func BenchmarkGetHash(b *testing.B) {
	l := make(map[string]string)
	for i := 0; i < 1000_000; i++ {
		k := fmt.Sprintf("%v", i)
		v := fmt.Sprintf("val -> %v", i)
		l[k] = v
	}
	a := ""
	void(a)
	keys := make([]string, b.N)
	for n := 0; n < b.N; n++ {
		keys[n] = fmt.Sprintf("%v", n)
	}
	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		a = l[keys[n]]
	}
	void(a)
}
func void(_ string) {
}

func TestSkipList_Iterator(t *testing.T) {
	l := New()
	assertNotEq(t, nil, l.Iterator())
}

func TestSkipList_RawSize(t *testing.T) {
	l := New()
	assertEq(t, int64(0), l.RawSize())
	k := []byte("test")
	v := []byte("abc")
	l.Set(k, v)
	assertEq(t, int64(len(k)+len(v)), l.RawSize())
}
