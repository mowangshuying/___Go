package main

import (
	"fmt"
	"math/rand"
)

// 1.初始化数组
func initarray() {
	// var arr[5] int
	// nums := []int{1,3,2,4,5}
}

// 2.随机访问元素
func randomAcess(nums []int) int {
	index := rand.Intn(len(nums))
	return nums[index]
}

// 3.插入元素
// 在index位置插入num.
func insert(nums []int, index int, num int) {
	for i:= len(nums) - 1; i > index; i-- {
		nums[i] = nums[i-1]
	}
	nums[index] = num
}

// 4.删除元素
func remove(nums []int, index int) {
	for i:=index; i < len(nums) - 1; i++ {
		nums[i] = nums[i+1]
	}
}

// 5.遍历数组
func traverse(nums []int) {
	// count := 0
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d", nums[i])
		if i != (len(nums)-1) {
			fmt.Printf("/")
		} else {
			fmt.Printf("\n")
		}	
	}
}

// 6.查找元素
func find(nums []int, target int) int {
	index := -1
	for i:=0; i < len(nums); i++ {
		if nums[i] == target {
			index = i
			break
		}
	}
	return index
}

// 7.扩容数组
func extend(nums []int, enlarge int) []int {
	res := make([]int, len(nums)+enlarge)
	for i,num := range nums {
		res[i] = num
	}
	return res
}

func main() {
	// fmt.Printf("--- start")
	
	// init array.
	nums :=[]int {1,3,2,4,5}
	traverse(nums)
	
	n0 := randomAcess(nums)
	n1 := randomAcess(nums)
	fmt.Printf("n0:%d, n1:=%d\n", n0, n1) 
	
	i1 := find(nums, 5)
	fmt.Printf("i:%d v:%d\n", i1, 5)
}