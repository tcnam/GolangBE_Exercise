package main

import "fmt"

func two_sum(nums []int, target int) []int {
	mymap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		j, ok := mymap[target-nums[i]]
		if ok {
			result := []int{j, i}
			return result
		}
		mymap[nums[i]] = i
	}
	result := []int{-1, -1}
	return result
}
func main() {
	var nums []int = []int{3, 2, 7, 11, 15}
	var target int = 9
	for val := range two_sum(nums, target) {
		fmt.Println(val)
	}

	var nums2 []int = []int{3, 2, 4}
	var target2 int = 6
	for val2 := range two_sum(nums2, target2) {
		fmt.Println(val2)
	}

}
