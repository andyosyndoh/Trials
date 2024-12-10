package algorithm

import "push/utils"

func StackBIns(b []int) string {
	bcopy := make([]int, len(b))
	copy(bcopy, b)
	utils.Sequence(bcopy)
	swap := sBIdx(bcopy)
	mid, last := len(bcopy)/2, len(bcopy)-1
	if swap > 0 {
		if utils.InOrder(b[last], b[last-1]) {
			return "sb"
		} else if swap < mid {
			return "rrb"
		}
	} else {
		maxNumid := utils.Biggest(bcopy)
		if maxNumid == last {
			return "pa"
		} else if maxNumid < mid || utils.InOrder(bcopy[0], bcopy[last]) {
			return "rrb"
		}
	}
	return "rb"
}

func sBIdx(nums []int) int {
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i-1]-nums[i] == 1 {
			return i
		}
	}
	return -1
}