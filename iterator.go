package skiplist

// A generic Iterator for going over values on specific point on SkipList.
//
// Example usage;
//
//
// 	for iter.Next() {
// 		fmt.Println(iter.Value());
// 	}
//
type Iterator interface {
	Next() bool
	Key() interface{}
	Value() interface{}
}

type iteratorImpl struct {
	list *SkipList
	cur  *node
}

func newIterator(list *SkipList) Iterator {
	return &iteratorImpl{list: list, cur: list.sentinel}
}

func (i *iteratorImpl) Next() bool {
	if i.list.count == 0 {
		return false
	}

	if i.cur != nil {
		if i.cur.Next != nil {
			i.cur = i.cur.Next[0]
		}
	}

	return i.cur != nil
}

// returns the key on the node
func (i *iteratorImpl) Key() interface{} {
	if i.cur != nil {
		return i.cur.Key
	}
	return nil
}

// returns the value on the node
func (i *iteratorImpl) Value() interface{} {
	if i.cur != nil {
		return i.cur.Value
	}
	return nil
}
