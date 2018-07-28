package goml

import (
	"fmt"
	"sync"
)

type Item interface{}

type Node struct {
	content Item
	next    *Node
}

type linkedList struct {
	head   *Node
	length int32
	lock   sync.RWMutex
}

func (lL *linkedList) Append(content Item) {
	lL.lock.Lock()
	defer lL.lock.Unlock()

	node := &Node{content, nil}

	if lL.length == 0 {
		lL.head = node
	} else {
		last := lL.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = node
	}
	lL.length = lL.length + 1
}

//func (lL *linkedList) String(){

//}
func linkedlist_main() {

	l := linkedList{}
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)

	fmt.Println("Length:", l.length)
	printList(&l)

}

func reverse(l *linkedList) {
	one := l.head
	two := one.next
}

func printList(l *linkedList) {
	list := l.head
	for {
		fmt.Printf("%v ", list.content)
		if list.next != nil {
			list = list.next
		} else {
			break
		}
	}
	fmt.Println()
}
