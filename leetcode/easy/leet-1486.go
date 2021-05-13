/*
https://leetcode.com/problems/xor-operation-in-an-array/

Input: n = 5, start = 0
Output: 8
Explanation: Array nums is equal to [0, 2, 4, 6, 8] where (0 ^ 2 ^ 4 ^ 6 ^ 8) = 8.
Where "^" corresponds to bitwise XOR operator.
*/

package main

import "fmt"

func main() {
	result := xorOperation(5, 0)
	fmt.Println(result)
}

func xorOperation(n int, start int) int {
	// create an empty array
	nums := make([]int, n)

	// populate the array with values from given formula
	for key, _ := range nums {
		nums[key] = start + 2*key
	}

	// calculate XOR using values from the array
	result := nums[0]

	for key, _ := range nums {
		if key == len(nums)-1 {
			return result
		}
		result = result ^ nums[key+1]
	}

	return result
}
