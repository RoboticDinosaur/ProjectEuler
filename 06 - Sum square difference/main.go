package main

import "fmt"

func SumOfSquares() int {

	var result int

	for i := 1; i < 101; i++ {
		result += i * i
	}

	return result
}

func SquareOfSums() int {
	var result int
	
	for i := 1; i < 101; i++ {
		result += i
	}

	return result * result
}

func main() {
	fmt.Println(SquareOfSums() - SumOfSquares())
}