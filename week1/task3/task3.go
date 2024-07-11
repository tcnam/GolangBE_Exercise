package main

import (
	"fmt"
)

func input_slice() []float32 {
	var size int
	fmt.Println("Enter size of the slices: ")
	fmt.Scanln(&size)
	slice := make([]float32, size, size*2)
	for i := range size {
		fmt.Printf("Enter the value for slice[%d]: ", i)
		fmt.Scanln(&slice[i])
	}
	return slice
}

func merge_sort(slice []float32) []float32 {
	if len(slice) <= 1 {
		return slice
	}
	var mid int = int(len(slice) / 2)
	var left_slice []float32 = slice[:mid]
	var right_slice []float32 = slice[mid:]

	var sorted_left_slice []float32 = merge_sort(left_slice)
	var sorted_right_slice []float32 = merge_sort(right_slice)

	return merge(sorted_left_slice, sorted_right_slice)
}

func merge(left_slice []float32, right_slice []float32) []float32 {
	var result_slice []float32
	var left_ind, right_ind int = 0, 0

	for left_ind < len(left_slice) && right_ind < len(right_slice) {
		if left_slice[left_ind] < right_slice[right_ind] {
			result_slice = append(result_slice, left_slice[left_ind])
			left_ind += 1
		} else {
			result_slice = append(result_slice, right_slice[right_ind])
			right_ind += 1
		}
	}

	for left_ind < len(left_slice) {
		result_slice = append(result_slice, left_slice[left_ind])
		left_ind += 1
	}
	for right_ind < len(right_slice) {
		result_slice = append(result_slice, right_slice[right_ind])
		right_ind += 1
	}
	return result_slice
}

func slice_info(slice []float32) (float32, float32, float32, float32) {

	var sum_val float32 = 0.0
	var new_slice []float32 = merge_sort(slice)
	for ind, val := range new_slice {
		sum_val += val
		fmt.Printf("slice[%d]: %.2f\n", ind, val)
	}
	return new_slice[0], new_slice[len(new_slice)-1], sum_val, sum_val / float32(len(new_slice))
}

func main() {
	var slice []float32 = input_slice()
	var min_val, max_val, sum_val, avg_val float32 = slice_info(slice)
	fmt.Printf("min_val: %.2f\nmax_val: %.2f\nsum_val: %.2f\navg_val: %.2f\n", min_val, max_val, sum_val, avg_val)
}
