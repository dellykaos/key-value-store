package cache

type linkedList struct {
	head *node
	tail *node
}

func newLinkedList() *linkedList {
	l := &linkedList{
		head: newNode("", nil),
		tail: newNode("", nil),
	}
	l.head.next = l.tail
	l.tail.prev = l.head

	return l
}

func (l *linkedList) delete(n *node) {
	n.next.prev, n.prev.next = n.prev, n.next
}

func (l *linkedList) add(n *node) {
	n.prev = l.tail.prev
	n.next = l.tail
	l.tail.prev.next = n
	l.tail.prev = n
}
