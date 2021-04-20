/*
https://leetcode.com/problems/running-sum-of-1d-array/

Example:
Input: nums = [1,2,3,4]
Output: [1,3,6,10]
Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].
 */
package main

import "fmt"

func main() {
	sumArray := runningSum([]int{1,2,3,4})
	fmt.Println(sumArray)
}

func runningSum(array []int) []int {
	var sumArray []int
	for i, entry := range array {
		if i == 0 {
			sumArray = append(sumArray, entry)
		} else {
			sum := array[i] + sumArray[i-1]
			sumArray = append(sumArray, sum)
		}
	}

	return sumArray
}
