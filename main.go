package main

import (
	"fmt"
	ds "go-code/src/datastructure"
	lc "go-code/src/leetcode"
)

func main() {
	nums := []int{12, 3, 134, 5, 12, 1, 67, 13, 5, 1, 1}
	fmt.Println(lc.ThreeSum(nums, 18))

	ds.InOrder(nil)
}
