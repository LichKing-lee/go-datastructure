package link

import "strconv"

type linkedList struct {
	Head *node
}

type node struct {
	Value int
	Next  *node
}

func Add(list *linkedList, element int) {
	n := findLastNode(list)
	n.Next = &node{Value: element}
}

func Del(list *linkedList, value int) {
	node := list.Head
	if node.Value == value {
		if node.Next == nil {
			list.Head = nil
			return
		}

		list.Head = node.Next
		return
	}

	for node.Next != nil {
		if node.Next.Value == value {
			node.Next = node.Next.Next
			break
		}
		node = node.Next
	}
}

func findLastNode(list *linkedList) (n *node) {
	for n := list.Head; n.Next != nil; {
		n = n.Next
	}
	return
}

func Print(list *linkedList) (str string) {
	node := list.Head
	if node == nil {
		return "is empty list"
	}

	for {
		str += strconv.Itoa(node.Value)
		node = node.Next

		if node == nil {
			break
		}

		str += " -> "
	}
	return
}

func NewLinkedList(node *node) *linkedList {
	return &linkedList{node}
}

func NewNode(value int, next *node) *node {
	return &node{value, next}
}