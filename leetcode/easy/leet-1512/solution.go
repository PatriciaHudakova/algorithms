/*
https://leetcode.com/problems/number-of-good-pairs/
Example:
Input: nums = [1,2,3,1,1,3]
Output: 4
Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.
*/
package main

import "fmt"

func main() {
	num := numOfGoodPairs([]int{1, 2, 3, 1, 1, 3})
	fmt.Println(num)
}

func numOfGoodPairs(array []int) int {
	num := 0

	for i, _ := range array {
		for j, _ := range array {
			if array[i] == array[j] && i < j {
				num++
			}
		}
	}
	return num
}
