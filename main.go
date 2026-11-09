package main

import "fmt"

func twoSum(nums []int, target int) []int {
	num_map := make(map[int]int)
	for i, v := range nums {
		if num_map[v] < 1 {
			num_map[v] = i + 1
		}
		if num_map[target-v] > 0 && i != num_map[target-v]-1 {
			return []int{num_map[target-v] - 1, i}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5, 6}, 6))
}
