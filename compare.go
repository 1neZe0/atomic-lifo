package main

import (
	"atomic-lifo/atomiclifo"
	"atomic-lifo/standartlifo"
	"fmt"
	"time"
)

var total_items = 1000

func testAtomic() {
	start := time.Now()
	q := atomiclifo.NewQueue()
	for i := 0; i < total_items; i++ {
		q.Push(&atomiclifo.Node{Value: i})
	}

	for i := 0; i < total_items; i++ {
		popNode := q.Pop()
		if popNode != nil {
			//fmt.Println(popNode.Value)
			_ = popNode
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %s\n", elapsed)
}

func testStandard() {
	start := time.Now()
	q := standardlifo.Stack{}
	for i := 0; i < total_items; i++ {
		q.Push(i)
	}

	for i := 0; i < total_items; i++ {
		popNode := q.Pop()
		if popNode != 0 {
			//fmt.Println(popNode)
			_ = popNode
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %s\n", elapsed)
}

func main() {
	testAtomic()
	testStandard()

}
