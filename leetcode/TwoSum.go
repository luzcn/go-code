package leetcode

import "sort"

func TwoSum(nums []int, target int) bool {
	var mp = make(map[int]int)

	for i, n := range nums {
		if _, found := mp[target-n]; found {
			return true
		}

		mp[n] = i
	}
	return false
}

func ThreeSum(nums []int, target int) [][]int {

	sort.Ints(nums)
	res := make([][]int, 0)

	i := 0
	for i < len(nums)-1 {

		// remove duplicate
		for i > 0 && nums[i] == nums[i-1] {
			i++
			continue
		}

		var left = i + 1
		var k = len(nums) - 1
		a := nums[i]

		for left < k {
			b := nums[left]
			c := nums[k]
			if b+c == target-a {
				res = append(res, []int{a, b, c})
				left++
				for left < len(nums) && nums[left] == nums[left-1] {
					left++
				}

				k--
				for k > 0 && nums[k] == nums[k+1] {
					k--
				}

				continue
			} else if b+c < target-a {
				left++

			} else {
				k--
			}
		}
		i++
	}
	return res
}
