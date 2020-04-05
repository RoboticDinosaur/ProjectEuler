package main

/* 03 - Largest prime factor
 * -------------------------
 *
 * The prime factors of 13195 are 5, 7, 13 and 29.
 * What is the largest prime factor of the number 600851475143 ?
 */

import (
	"fmt"
)

// PrimeFactors - Takes in a value and creates an array of Prime
// Factor numbers.
//
// Converted from C++ with a little help from another solution.
// It will need to be updated at a later date.
func PrimeFactors(n int64) (a []int64) {

	for n%2 == 0 {
		a = append(a, 2)
		n = n / 2
	}

	var i int64

	for i = 3; i*i <= n; i += 2 {
		for n%i == 0 {
			a = append(a, i)
			n = n / i
		}
	}

	if n > 2 {
		a = append(a, n)

	}
	return
}

func main() {
	fmt.Println(PrimeFactors(600851475143))
}
