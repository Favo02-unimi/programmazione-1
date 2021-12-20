package main

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	str := args[1]
	var nums []int
	for i := 0; i < len(str); i++ {
		n, _ := strconv.Atoi(string(str[i]))
		nums = append(nums, n)
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			Print(nums[i])
		}
	}
}
