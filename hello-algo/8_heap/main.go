package main

import (
	"___Go/hello-algo/8_heap/heap"
	"fmt"
)

// 堆是一种满足特定条件的完全二叉树，分为两种类型:
// 小顶堆
// 大顶堆

// 许多程序语言提供的是优先队列，这是一种抽象的数据结构，定义为具有优先级排序的队列
// 优先队列 == 堆
func main() {
	maxHeap := &heap.IntHeap{}
	maxHeap.Push(1)
	maxHeap.Push(3)
	maxHeap.Push(2)
	maxHeap.Push(4)
	maxHeap.Push(5)

	fmt.Printf("%d,", maxHeap.Pop())
	fmt.Printf("%d,", maxHeap.Pop())
	fmt.Printf("%d,", maxHeap.Pop())
	fmt.Printf("%d,", maxHeap.Pop())
	fmt.Printf("%d,", maxHeap.Pop())
	fmt.Printf("\n")
}
