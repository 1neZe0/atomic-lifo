package atomiclifo

import (
	"sync/atomic"
	"unsafe"
)

func NewQueue() *Queue {
	node := &Node{}
	return &Queue{head: node, tail: node}
}

type Node struct {
	Value int
	Next  *Node
}

type Queue struct {
	head, tail *Node
}

func (q *Queue) Push(n *Node) {
	for {
		tail := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&q.tail)))
		next := (*Node)(tail).Next
		if tail == atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&q.tail))) {
			if next == nil {
				if atomic.CompareAndSwapPointer(
					(*unsafe.Pointer)(unsafe.Pointer(&(*Node)(tail).Next)),
					nil,
					unsafe.Pointer(n),
				) {
					atomic.CompareAndSwapPointer(
						(*unsafe.Pointer)(unsafe.Pointer(&q.tail)),
						tail,
						unsafe.Pointer(n),
					)
					return
				}
			} else {
				atomic.CompareAndSwapPointer(
					(*unsafe.Pointer)(unsafe.Pointer(&q.tail)),
					tail,
					unsafe.Pointer(next),
				)
			}
		}
	}
}

func (q *Queue) Pop() *Node {
	for {
		head := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&q.head)))
		tail := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&q.tail)))
		next := (*Node)(head).Next
		if head == atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&q.head))) {
			if head == tail {
				if next == nil {
					return nil
				}
				atomic.CompareAndSwapPointer(
					(*unsafe.Pointer)(unsafe.Pointer(&q.tail)),
					tail,
					unsafe.Pointer(next),
				)
			} else {
				value := next.Value
				if atomic.CompareAndSwapPointer(
					(*unsafe.Pointer)(unsafe.Pointer(&q.head)),
					head,
					unsafe.Pointer(next),
				) {
					return &Node{Value: value}
				}
			}
		}
	}
}
