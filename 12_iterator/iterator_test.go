package iterator

import "testing"

func TestIterator(t *testing.T) {
	aggregate := NewNumbers(10, 20)

	IteratorPrint(aggregate.Iterator())
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10

	aggregateArr := NewArrays([]string{
		"hello",
		"world",
	})

	IteratorPrint(aggregateArr.Iterator())
}
