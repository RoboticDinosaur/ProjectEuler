package main

import (
	"fmt"
	"sort"
)

/* 04 - Largest palindrome product
 * -------------------------
 *
 * A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
 *
 * Find the largest palindrome made from the product of two 3-digit numbers.
 */

// CheckPalindrome - Returns true or false if the number is a palindrome.
func CheckPalindrome(number int) bool {
	var remainder int
	var reverse int = 0
	var temp = number

	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number /= 10

		if number == 0 {
			break
		}
	}
	if temp == reverse {
		return true
	}

	return false

}

// RemoveDuplicates - Removes duplicate numbers from the array
func RemoveDuplicates(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	var one = 999
	var two = 999

	var array []int

	for one > 0 {
		for two > 0 {
			if CheckPalindrome(one*two) == true {
				array = append(array, one*two)
			}
			two--
		}
		one--
		two = 999
	}

	sort.Ints(array)

	fmt.Println(array[len(array)-1])
}
