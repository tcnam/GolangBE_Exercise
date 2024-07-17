package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func get_person_info_from_file(file_name string) []Person {
	var result []Person = make([]Person, 0)
	pwd, _ := os.Getwd()
	myfile, err := os.Open(filepath.Join(pwd, file_name))
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return result
	}

	defer myfile.Close()

	rd := bufio.NewReader(myfile)
	for {
		line, _, err := rd.ReadLine()
		if len(line) > 0 {
			// fmt.Printf("Readline: %q\n", line)
			var person_info []string = strings.Split(string(line), ",")
			var name string = strings.ToUpper(person_info[0])
			var job string = strings.ToLower(person_info[1])
			// fmt.Printf("%T, %T, %T", person_info[0], person_info[1], person_info[len(person_info)-1])
			// for _, val := range person_info {
			// 	fmt.Printf("%T\n", val)
			// }
			var date_info string = strings.Split(person_info[2], "-")
			var dump string = person_info[3]

			var date_of_birth time.Time = time.Date(int(date_info[2]), time.Month(int(date_info[1])), int(date_info[0]), 0, 0, 0, 0, time.UTC)
			result = append(result, Person{name, job, date_of_birth})
			fmt.Printf("%v\n", result)

		}
		if err != nil {
			break
		}
	}
	return result
}

func main() {
	var per Person = Person{name: "Nam1", job: "Data engineer", date_of_birth: time.Date(2000, time.Month(2), 15, 0, 0, 0, 0, time.UTC)}
	fmt.Println(per.calculate_age())
	fmt.Println(per.check_match_with_job())

	for key, value := range get_char_frequency("aaaaaaabcbbc") {
		fmt.Printf("%c:%v\n", key, value)
	}

	get_person_info_from_file("a.txt")
}
