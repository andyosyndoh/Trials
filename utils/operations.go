package utils

import "sort"

func Unique(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return false
		}
		seen[num] = true
	}
	return true
}

func Sequence(nums []int) {
	copyofnums := make([]int, len(nums))
	copy(copyofnums, nums)
	sort.Ints(copyofnums)
	for i, n := range nums {
		nums[i] = BinarySearch(copyofnums, n)
	}
}

func BinarySearch(nums []int, n int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == n {
			return mid
		}
		if nums[mid] < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1 // n is not in the array.
}


func IntsAreSorted(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]-nums[i+1] != 1 {
			return false
		}
	}
	return true
}



func InOrder(x, y int) bool {
	return y-x == 1
}

func Smallest(nums []int) int {
	min, id := 0x3f3f3f3f, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
			id = i
		}
	}
	return id
}

func Biggest(nums []int) int {
	max, id := -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			id = i
		}
	}
	return id
}