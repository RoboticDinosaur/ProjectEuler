package main

import "fmt"

// SumOfSquares returns the sum of the squares.
func SumOfSquares() int {

	var result int

	for i := 1; i < 101; i++ {
		result += i * i
	}

	return result
}

// SquareOfSums retuns the square of the sums.
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
