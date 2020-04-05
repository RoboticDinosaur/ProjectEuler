package main

import "fmt"

// gcd - get the Greatest Common Divisor via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

// lcm - fint the Least Common Multiple via gcm
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	result := 1
	for j := 2; j <= 20; j++ {
		result = lcm(result, j)
	}

	fmt.Println(result)
}
