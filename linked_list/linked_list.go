package linked_list

import (
	"fmt"
)

type node struct {
	next *node
	key  int
}

type LinkedList struct {
	head *node
}

func (l *LinkedList) Prepend(key int) {
	n := &node{next: l.head, key: key}
	l.head = n
	fmt.Printf("SUCCESS: prepend %d", key)
}

func (l *LinkedList) RemoveHead() {
	head := l.head
	if head == nil {
		fmt.Println("list is empty")
	} else {
		key := l.head.key
		l.head = l.head.next
		fmt.Println("removedHead' key is %d", key)
	}
}

func (l *LinkedList) IsExist(num int) bool {
	node := l.head
	for node != nil {
		if node.key == num {
			return true
		}
		node = node.next
	}

	return false
}

func (l *LinkedList) Remove(num int) {
	node := l.head
	cnt := 0
	for {
		if node.next == nil {
			break
		}
		if node.key == node.next.key {
			node.next = node.next.next
			cnt++
		} else {
			node = node.next
		}
	}

	fmt.Printf("%d node deleted", cnt)
}

func (l *LinkedList) GetLength() {
	cnt := 0
	node := l.head
	for node != nil {
		cnt++
		node = node.next
	}

	fmt.Printf("length is %d", cnt)
}

func (l *LinkedList) PrintAll() {
	node := l.head
	list := []int{}
	for node != nil {
		list = append(list, node.key)
		node = node.next
	}

	for i, key := range list {
		fmt.Println(i, key)
	}
}
