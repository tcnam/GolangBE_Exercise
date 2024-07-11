package main

import (
	"fmt"
)

func even_str_or_not(str string) bool {
	var length int = len(str)
	if length%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var str_odd string = "abc"
	var str_even string = "ab"
	fmt.Println(even_str_or_not(str_odd))
	fmt.Println(even_str_or_not(str_even))
}
