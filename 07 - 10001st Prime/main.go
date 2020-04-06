package main

import (
	"fmt"
	"math/big"
)

func main() {
	const n = 10001

	var total int = 0
	var count int64 = 0

	for total < n {
		temp := big.NewInt(count)
		if temp.ProbablyPrime(1) {
			total++
		}
		count++
	}
	fmt.Println(count - 1)
}
