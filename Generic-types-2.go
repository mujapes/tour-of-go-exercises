package main

import "fmt"
// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	fmt.Println(populateLinkedList([]int{2,865,236,864,342,247}))
}

func populateLinkedList[T any](unlinkedSlice []T) []List[T]{
	var linkedList []List[T]
    var prev *List[T] = nil
	for i, v := range unlinkedSlice {
		if i > 0 {
			linkedList = append(linkedList, List[T]{&linkedList[i-1], v})
		} else {
			linkedList = append(linkedList, List[T]{prev, v})
		}
	}
	return linkedList
}
