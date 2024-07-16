package main

import "fmt"

func two_sum(nums []int, target int) []int {

	var hash_map map[int]int = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		j, ok := hash_map[target-nums[i]]
		if ok {
			result := []int{j, i}
			return result
		}
		hash_map[nums[i]] = i
	}
	result := []int{-1, -1}
	return result
}
func main() {

	var nums []int = []int{2, 7, 11, 15}
	var target int = 9
	fmt.Printf("%v", two_sum(nums, target))

	var nums2 []int = []int{3, 2, 4}
	var target2 int = 6
	fmt.Printf("%v", two_sum(nums2, target2))
}
