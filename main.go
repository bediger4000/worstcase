package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Node struct {
	Data uint
	Next *Node
}

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	hd := presortedList(n, false)

	Print(hd)

	node := createWorstCase(hd)
	Print(node)

	nl := recursiveMergeSort(node)
	Print(nl)
	if sz, sorted := isSorted(nl); sorted {
		fmt.Printf("List of length %d is sorted\n", sz)
		if sz != n {
			fmt.Printf("Sorted list not same length (%d) as unsorted (%d)\n", sz, n)
		}
	} else {
		fmt.Printf("List of length %d is NOT sorted\n", sz)
		if sz != n {
			fmt.Printf("Sorted list not same length (%d) as unsorted (%d)\n", sz, n)
		}
	}
}

func createWorstCase(node *Node) *Node {
	if node.Next == nil {
		return node
	}

	left, right := alternateSplit(node)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	if left.Next == nil && right.Next == nil {
		if left.Data > right.Data {
			left.Next = right
			return left
		}
		right.Next = left
		return right
	}

	left = createWorstCase(left)
	right = createWorstCase(right)

	tail := &left
	for (*tail).Next != nil {
		tail = &(*tail).Next
	}

	(*tail).Next = right

	return left
}

func alternateSplit(node *Node) (*Node, *Node) {
	left := &Node{}
	right := &Node{}

	tL, tR := left, right

	for node != nil {
		tL.Next = node
		node = node.Next
		tL = tL.Next

		if node != nil {
			tR.Next = node
			node = node.Next
			tR = tR.Next
		}
	}

	tL.Next = nil
	tR.Next = nil

	return left.Next, right.Next
}

func split(head *Node) (*Node, *Node) {
	// Setting rabbit and turtle like this means we split an
	// odd-length-list (head) into lists of length n (right)
	// and n+1 (left).
	rabbit, turtle := head, &head

	for rabbit != nil {
		turtle = &(*turtle).Next
		if rabbit = rabbit.Next; rabbit != nil {
			rabbit = rabbit.Next
		}
	}

	right := *turtle
	*turtle = nil

	return head, right
}

func presortedList(n int, _ bool) *Node {

	var head *Node

	for i := n; i > 0; i-- {
		head = &Node{
			Data: uint(i),
			Next: head,
		}
	}

	return head
}

func Print(list *Node) {
	for node := list; node != nil; node = node.Next {
		fmt.Printf("%d -> ", node.Data)
	}
	fmt.Println()
}

func recursiveMergeSort(head *Node) *Node {
	if head.Next == nil {
		// single node list is sorted by definition
		return head
	}

	fmt.Printf("enter recursiveMergeSort: ")
	Print(head)

	// because of recursion bottoming out at a 1-long-list,
	// head points to a list of at least 2 elements.

	// Setting rabbit and turtle like this means we split an
	// odd-length-list (head) into lists of length n (right)
	// and n+1 (left).
	rabbit, turtle := head.Next, &head

	for rabbit != nil {
		turtle = &(*turtle).Next
		if rabbit = rabbit.Next; rabbit != nil {
			rabbit = rabbit.Next
		}
	}

	right := *turtle
	*turtle = nil

	left := recursiveMergeSort(head)
	right = recursiveMergeSort(right)

	fmt.Printf("\nmerging lists\n\tleft: ")
	Print(left)
	fmt.Printf("\tright: ")
	Print(right)

	/*
		// Set h, t variables so that the loop doing the merge
		// does not have to have a "if h == nil" check every iteration.
		x := &right
		if left.Data < right.Data {
			x = &left
		}

		h, t := *x, *x
		*x = (*x).Next
	*/

	dummy := &Node{}
	h, t := dummy, dummy

	// left and right are either equal in length, or right is one
	// node longer, but the "<" check might take more from one list
	// than the other. Have to check both for nil.
	for left != nil && right != nil {
		n := &right
		if left.Data < right.Data {
			n = &left
		}
		t.Next = *n
		*n = (*n).Next
		t = t.Next
		// At the end of this for-loop, t.Next ends up being nil
		// because of the left/right list splitting.
	}

	fmt.Printf("merged list: ")
	Print(h.Next)
	if left != nil {
		fmt.Printf("left list : ")
		Print(left)
	}
	if right != nil {
		fmt.Printf("right list : ")
		Print(right)
	}

	// Either left or right are nil. If left == nil,
	// assigning nil to t.Next is no issue.
	t.Next = left
	if right != nil {
		// but if right is nil, can't assign nil to t.Next,
		// because left was non-nil.
		t.Next = right
	}

	return h.Next
}

func isSorted(head *Node) (int, bool) {
	if head == nil {
		return 0, true
	}
	if head.Next == nil {
		return 1, true
	}
	var sz int
	for ; head.Next != nil; head = head.Next {
		sz++
		if head.Data > head.Next.Data {
			return sz, false
		}
	}
	sz++ // for-loop checks head.Next, count final element on list
	return sz, true
}
