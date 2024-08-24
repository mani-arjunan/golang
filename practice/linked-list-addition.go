package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func linkedListAddition(node1, node2 *Node) {
	var resultNode *Node
	// var currentNode *Node
	carry := 0

	for node1 != nil || node2 != nil {
		node1Val := 0
		node2Val := 0

		if node1 != nil {
			node1Val = node1.value
		}

		if node2 != nil {
			node2Val = node2.value
		}
		sum := node1Val + node2Val + carry

		if sum > 9 {
			carry = sum / 10
			sum = sum % 10
		}

		tempNode := &Node{
			value: sum,
			next:  resultNode,
		}

		resultNode = tempNode

		node1 = node1.next
		node2 = node2.next
	}

	resultNode.value = resultNode.value + (carry * 10)
	fmt.Println(resultNode.value)
	fmt.Println(resultNode.next.value)
	fmt.Println(resultNode.next.next.value)
	fmt.Printf("Node 1 => %v \nNode 2 => %v", node1, node2)
}

func main() {
	node1 := &Node{
		value: 1,
		next: &Node{
			value: 9,
			next: &Node{
				value: 9,
				next:  nil,
			},
		},
	}
	node2 := &Node{
		value: 4,
		next: &Node{
			value: 8,
			next: &Node{
				value: 9,
				next:  nil,
			},
		},
	}

	linkedListAddition(node1, node2)
}

// 1 9 9
// 4 8 9
// 5 7 13
