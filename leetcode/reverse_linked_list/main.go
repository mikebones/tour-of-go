package main

func main() {

}

// linked list is a linear data strcture by references in memory
// not their physical locatio in memory
// no contigious locations like arryas
// arrays are fixed size contiguous memory store
// linked lists are linked u sing points

// a linked list is a collection of nodes

type Node struct {
	value int // data
	next *Node // reference // initializes with nil
}

type LinkedList struct {
	len int
	head *Node // reference the top of the linked list
	// follow the rest of the list down
}

// change the last node with a value of nil
// to the new memory location of the node

func (list *LinkedList) Insert(val int) {
	n := Node{}
	n.value = val
	if list.len == 0 {
		list.head = &n
		list.len++
		return
	}
	nextNodePointer := list.head
	for i := 0 ; i < list.len ; i ++ {
		if nextNodePointer.next == nil {
			nextNodePointer.next = &n
			list.len++
			return
		}
		nextNodePointer = nextNodePointer.next
	}
}

func (list *LinkedList) InsertAt(pos int, val int) {
		
}


// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

//singlylinked list
func reverseList(head *ListNode) *ListNode {
	var prevPtr *ListNode = nil
	var currentPtr *ListNode = head
	for (currentPtr != nil) {
		
		//  head.Next the previous pointer
		
		// head.Next = prevPtr
		// head = currentPtr
		// prevPtr = currentPtr
		// currentPtr = currentPtr.Next
		next := currentPtr.Next
		currentPtr.Next = prevPtr
		prevPtr = currentPtr
		currentPtr = next
	}
	return currentPtr
	// find nil
	// hold previous listnode
	// pointer to current
}

// time O(n)
// memorry O(1)
