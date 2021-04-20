/*
https://leetcode.com/problems/reverse-integer/

Example:
Input: x = 123
Output: 321
 */
package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(0))
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
}

func reverse(x int) int {
	// initialise the coefficient
	coefficient := 1

	// if the input is negative, make it positive
	if x < 0 {
		coefficient = -1
		x *= coefficient
	}

	// 32 bit limit for an integer
	limit := 2147483648
	
	reverse := 0
	for x > 0 {
		reverse = (reverse * 10) +  x % 10
		x /= 10
	}

	// check the limit boundaries are satisfied
	if reverse > limit || reverse < -limit {
		return 0
	}

	// if the input was originally negative, this will restore it's initial state
	return reverse * coefficient
}
