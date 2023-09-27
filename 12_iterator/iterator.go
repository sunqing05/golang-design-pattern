package iterator

import "fmt"

type Aggregate interface {
	Iterator() Iterator
}

type Iterator interface {
	First()
	IsDone() bool
	Next() interface{}
}

type Numbers struct {
	start, end int
}

func NewNumbers(start, end int) *Numbers {
	return &Numbers{
		start: start,
		end:   end,
	}
}

func (n *Numbers) Iterator() Iterator {
	return &NumbersIterator{
		numbers: n,
		next:    n.start,
	}
}

type NumbersIterator struct {
	numbers *Numbers
	next    int
}

func (i *NumbersIterator) First() {
	i.next = i.numbers.start
}

func (i *NumbersIterator) IsDone() bool {
	return i.next > i.numbers.end
}

func (i *NumbersIterator) Next() interface{} {
	if !i.IsDone() {
		next := i.next
		i.next++
		return next
	}
	return nil
}

func IteratorPrint(i Iterator) {
	for i.First(); !i.IsDone(); {
		c := i.Next()
		fmt.Printf("%#v\n", c)
	}
}

type Arrays struct {
	Items []string
	Index int
}

func NewArrays(s []string) *Arrays {
	return &Arrays{
		Items: s,
		Index: 0,
	}
}

func (n *Arrays) Iterator() Iterator {
	return &ArraysIterator{
		items: n,
		next:  n.Index,
	}
}

type ArraysIterator struct {
	items *Arrays
	next  int
}

func (i *ArraysIterator) First() {
	i.next = 0
}

func (i *ArraysIterator) IsDone() bool {
	return i.next >= len(i.items.Items)
}

func (i *ArraysIterator) Next() interface{} {
	if !i.IsDone() {
		next := i.next
		i.next++
		return i.items.Items[next]
	}
	return nil
}
