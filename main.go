package main

import (
	"fmt"
	"list/link"
)

func main() {
	node := link.NewNode(10, nil)
	node.Next = link.NewNode(20, link.NewNode(30, nil))
	list := link.NewLinkedList(node)

	fmt.Println(link.Print(list))

	link.Del(list, 20)
	link.Del(list, 20)
	link.Del(list, 20)

	fmt.Println(link.Print(list))

	link.Del(list, 10)

	fmt.Println(link.Print(list))
}
