package datastructure

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	var mid int = len(nums) / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])

	return merge(left, right)
}

// merge two sorted array O(m+n)
func merge(left []int, right []int) []int {
	res := make([]int, 0)

	i, j := 0, 0
	for i < len(left) || j < len(right) {
		if i >= len(left) {
			res = append(res, right[j])
			j++
			continue
		}

		if j >= len(right) {
			res = append(res, left[i])
			i++
			continue
		}

		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	return res
}

func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, l int, r int) {
	if l >= r {
		return
	}

	prefix := l - 1
	pivot := nums[r]

	for i := l; i < r; i++ {
		if nums[i] < pivot {
			prefix++

			swap(nums, prefix, i)
		}
	}
	prefix++
	swap(nums, prefix, r)

	quickSort(nums, l, prefix-1)
	quickSort(nums, prefix+1, r)
}

func swap(nums []int, i int, j int) {
	for i < j {
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
		i++
		j--
	}
}
