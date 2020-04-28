package main

import "fmt"

func main() {

	var current int = 1
	var result int = 2
	var end = 4000000
	var temp int
	var array []int

	for result < end {
		temp = current + result
		current = result
		result = temp

		if (current % 2) == 0 {
			array = append(array, current)
		}
	}

	result = 0
	for _, v := range array {
		result += v
	}

	fmt.Println(result)
	fmt.Println(array)

}
