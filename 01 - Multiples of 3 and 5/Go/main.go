package main

import "fmt"

func main() {

	var number int = 1000
	var result int = 0

	for i := 1; i < number; i++ {
		if (i % 3) == 0 {
			result = result + i
		} else if (i % 5) == 0 {
			result = result + i
		}
	}

	fmt.Println(result)

}
