/*
https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/

Input: nums = [8,1,2,2,3]
Output: [4,0,1,1,3]
Explanation:
For nums[0]=8 there exist four smaller numbers than it (1, 2, 2 and 3).
For nums[1]=1 does not exist any smaller number than it.
For nums[2]=2 there exist one smaller number than it (1).
For nums[3]=2 there exist one smaller number than it (1).
For nums[4]=3 there exist three smaller numbers than it (1, 2 and 2).
 */
package main

import "fmt"

func main() {
	array := smallerNumbersThanCurrent([]int{8,1,2,2,3})
	fmt.Println(array)
}

func smallerNumbersThanCurrent(nums []int) []int {
	var array []int

	for i, _ := range nums {
		count := 0
		for j, _ := range nums {
			if i != j && nums[j] < nums[i] {
				count++
			}
		}
		array = append(array, count)
	}

	return array
}
