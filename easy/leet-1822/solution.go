/*
https://leetcode.com/problems/sign-of-the-product-of-an-array/
Input: nums = [-1,-2,-3,-4,3,2,1]
Output: 1
Explanation: The product of all values in the array is 144, and signFunc(144) = 1
 */
package main

import "fmt"

func main() {
	fmt.Println(arraySign([]int{-1,-2,-3,-4,3,2,1}))
	fmt.Println(arraySignPreferred([]int{-1,-2,-3,-4,3,2,1}))

	// overflow use case
	fmt.Println(arraySign([]int{9,72,34,29,-49,-22,-77,-17,-66,-75,-44,-30,-24}))
}

/*
As per requirement, we evaluate the product value
 */
func arraySign(nums []int) int {
	product := 1
	for _, value := range nums {
		product = product * value

		// manually check for overflow (int32), not great
		if product > 2147483647 {
			product = product - 2147483647
		}

		if product < -2147483647 {
			product = product + 2147483647
		}
	}

	switch {
	case product == 0:
		return 0
	case product < 0:
		return -1
	default:
		return 1
	}
}

/*
My preferred approach would have been focusing on the signs rather than the actual product value,
also prevents overflow
*/
func arraySignPreferred(nums []int) int {
	negCounter := 0

	for _, value := range nums {
		if value < 0 {
			negCounter++
		}
		if value == 0 {
			return 0
		}
	}

	// if the amount of -ve integers are odd then the overall value must be negative
	if negCounter % 2 == 0 {
		return 1
	}

	return -1
}