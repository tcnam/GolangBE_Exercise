package main

import (
	"fmt"
	"time"
)

type Person struct {
	name          string
	job           string
	date_of_birth time.Time
}

func (per *Person) calculate_age() int {
	var age int = time.Now().Year() - per.date_of_birth.Year()
	return age
}

func (per *Person) check_match_with_job() bool {
	if per.date_of_birth.Year()%len(per.name) == 0 {
		return true
	} else {
		return false
	}
}

func get_char_frequency(s string) map[rune]int {
	var result_map map[rune]int = make(map[rune]int)
	for _, character := range s {
		if _, ok := result_map[character]; ok {
			result_map[character] += 1
		} else {
			result_map[character] = 1
		}
	}
	return result_map
}

func main() {
	var per Person = Person{name: "Nam1", job: "Data engineer", date_of_birth: time.Date(2000, time.Month(2), 15, 0, 0, 0, 0, time.UTC)}
	fmt.Println(per.calculate_age())
	fmt.Println(per.check_match_with_job())

	for key, value := range get_char_frequency("aaaaaaabcbbc") {
		fmt.Printf("%c:%v\n", key, value)
	}
}
