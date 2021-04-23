/*
https://leetcode.com/problems/sum-of-all-odd-length-subarrays/
Input: arr = [1,4,2,5,3]
Output: 58
Explanation: The odd-length subarrays of arr and their sums are:
[1] = 1/
[4] = 4
[2] = 2
[5] = 5
[3] = 3
[1,4,2] = 7/
[4,2,5] = 11
[2,5,3] = 10
[1,4,2,5,3] = 15/
If we add all these together we get 1 + 4 + 2 + 5 + 3 + 7 + 11 + 10 + 15 = 58
 */

package main

import "fmt"

func main() {
	sum1 := sumOddLengthSubarrays([]int{1,4,2,5,3})
	sum2 := sumOddLengthSubarrays([]int{6,9,14,5,3,8,7,12,13,1})
	fmt.Println(sum1, sum2)
}

func sumOddLengthSubarrays(arr []int) int {
	sum := 0

	for i, value := range arr {
		sum += ((len(arr) - i) * (i + 1) + 1) / 2 * value
	}
	return sum
}