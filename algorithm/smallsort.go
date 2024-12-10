package algorithm

import (
	"sort"

	"push/utils"
)

func Smallsort(nums []int) (string, []*utils.Stack ) {
	stack := []*utils.Stack{utils.NewStack(nums), utils.NewStack(make([]int, 0, len(nums)))}
	return SmallInstructions(len(nums), stack), stack
}

func SmallInstructions(length int, stack []*utils.Stack) string {
	half := length / 2
	res := ""
	for {
		a, b := stack[0].Nums, stack[1].Nums
		if utils.IntsAreSorted(a) {
			break
		}
		op := StackAIns(a, half)
		op = matchbothstacks(op, a, b)
		Operation(stack, op)
		res += op + "\n"
	}
	for {
		a, b := stack[0].Nums, stack[1].Nums
		if len(a) == length && utils.IntsAreSorted(a) {
			break
		}
		op := StackBIns(b)
		Operation(stack, op)
		res += op + "\n"
	}
	return res
}



func matchbothstacks(ins string, a, b []int) string {
	if len(b) < 2 || sort.IntsAreSorted(b) {
		return ins
	}
	InsB := StackBIns(b)
	switch ins {
	case "pb":
		if InsB == "sb" {
			if t := len(a) - sAIdx(a); t < 3 {
				if t == 1 {
					return "ss"
				}
				return "ra"
			} else {
				return "sb"
			}
		}
	case "sa":
		if t := len(b) - sBIdx(b); t < 3 {
			if InsB == "sb" {
				return "ss"
			}
			return "rb"
		}
	case "ra":
		if InsB == "rb" || InsB == "sb" && len(b) == 2 {
			return "rr"
		}
	case "rra":
		if InsB == "rrb" || InsB == "sb" && len(b) == 2 {
			return "rrr"
		}
	}
	return ins
}


