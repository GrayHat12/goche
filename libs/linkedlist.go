package libs

import "cmp"

type Node[K cmp.Ordered, T any] struct {
	prev *Node[K, T]
	next *Node[K, T]
	key  K
	data T
}

type LinkedList[K cmp.Ordered, T any] struct {
	start   Node[K, T]
	end     Node[K, T]
	size    int
	address map[K]*Node[K, T]
}

func NewLinkedList[K cmp.Ordered, T any]() *LinkedList[K, T] {
	var nilT T
	var nilK K
	return &LinkedList[K, T]{
		start: Node[K, T]{
			prev: nil,
			next: nil,
			data: nilT,
			key:  nilK,
		},
		end: Node[K, T]{
			prev: nil,
			next: nil,
			data: nilT,
			key:  nilK,
		},
		size:    0,
		address: make(map[K]*Node[K, T]),
	}
}

func (ll *LinkedList[K, T]) Add(key K, value T) {
	newNode := Node[K, T]{
		prev: nil,
		next: nil,
		data: value,
		key:  key,
	}
	ll.address[key] = &newNode
	ll.size += 1
	if ll.start.next == nil {
		newNode.prev = &ll.start
		ll.start.next = &newNode
		ll.end.prev = &newNode
		return
	} else {
		if ll.end.prev == nil {
			panic("Something seriously went wrong")
		}
		newNode.next = &ll.end
		ll.end.prev.next = &newNode
		ll.end.prev = &newNode
	}
}

func (ll *LinkedList[K, T]) Get(key K) (T, bool) {
	val, exists := ll.address[key]
	var nilT T
	if exists {
		return val.data, exists
	} else {
		return nilT, exists
	}
}

func (ll *LinkedList[K, T]) Remove(key K) bool {
	node, exists := ll.address[key]
	if !exists {
		return false
	}
	delete(ll.address, key)
	node.prev.next, node.next.prev = node.next, node.prev
	ll.size -= 1
	return true
}

func (ll *LinkedList[K, T]) Size() int {
	return ll.size
}

func (ll *LinkedList[K, T]) RemoveFirst() bool {
	if ll.Size() > 0 {
		return ll.Remove(ll.start.next.key)
	}
	return false
}

func (ll *LinkedList[K, T]) RemoveLast() bool {
	if ll.Size() > 0 {
		return ll.Remove(ll.end.prev.key)
	}
	return false
}
