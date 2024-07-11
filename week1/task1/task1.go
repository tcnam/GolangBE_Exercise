package main

import (
	"fmt"
)

func rect_info(height float32, weight float32) (float32, float32) {
	var square float32 = height * weight
	var perimeter float32 = (height + weight) * 2
	return square, perimeter
}

func main() {
	var height float32 = 100.76
	var weight float32 = 4.6
	var square, perimeter float32 = rect_info(height, weight)
	fmt.Println(square, perimeter)
}
