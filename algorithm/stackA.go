package algorithm

import "push/utils"


func StackAIns(a []int, half int) string {
	swap := sAIdx(a)
	mid, last := len(a)/2, len(a)-1
	if a[last] < half {
		return "pb"
	}
	if swap > 0 {
		if utils.InOrder(a[last-1], a[last]) {
			return "sa"
		} else if swap < mid {
			return "rra"
		}
	} else {
		minNumId := utils.Smallest(a)
		if minNumId == last {
			return "pb"
		} else if minNumId < mid || utils.InOrder(a[last], a[0]) {
			return "rra"
		}
	}
	return "ra"
}

func sAIdx(nums []int) int {
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i]-nums[i-1] == 1 {
			return i
		}
	}
	return -1
}