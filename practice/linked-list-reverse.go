package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

// POINTER RETURN TYPE IMPLEMENTATION
// func reverseLinkedList(node *Node) *Node {
//   var reverseNode *Node
// 	current := node
//
// 	for current != nil {
//
// 		temp := Node{
// 			value: current.value,
// 			next:  reverseNode,
// 		}
//
// 		reverseNode = &temp
//
// 		current = current.next
//
// 	}
//
// 	return reverseNode
// }

// Node RETURN TYPE IMPLEMENTATION
func reverseLinkedList(node *Node) Node {
	var prevPtr *Node = nil
	current := node

	for current != nil {
		temp := Node{value: current.value, next: prevPtr}

		prevPtr = &temp
		current = current.next
	}

	return *prevPtr
}

func main() {
	fmt.Printf("Reversing a LinkedList")
	node := reverseLinkedList(&Node{
		value: 1,
		next: &Node{
			value: 2,
			next: &Node{
				value: 3,
				next:  nil,
			},
		},
	})

	fmt.Printf("%+v", node)
	// fmt.Printf("%v", node.
	fmt.Printf("%v", node.next.value)
	fmt.Printf("%v", node.next.next.value)
	fmt.Printf("%v", node.next.next.next)
	// fmt.Printf("%v", node.next.next.next.next.next.value)
}
