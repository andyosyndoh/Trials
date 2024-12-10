package main

import (
	"fmt"
	"os"
	"push/algorithm"
	"push/utils"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	input := os.Args[1]
	splittedinput := strings.Split(input, " ")
	nums := []int{}
	for _, numStr := range splittedinput {
		num, err := strconv.Atoi(numStr)
		if err == nil {
			nums = append(nums, num)
		} else {
			fmt.Println("Error")
			os.Exit(0)
		}
	}

	if sort.IntsAreSorted(nums) || !utils.Unique(nums) {
		// fmt.Println("The input array is already sorted in ascending order.")
        return
	}

	utils.Sequence(nums)
	instructions := ""
	var stack []*utils.Stack

	if len(nums) <= 100 {
		instructions, stack = algorithm.Smallsort(nums) 
	}
	fmt.Print(instructions)
	fmt.Println(stack[0].Nums)
}



