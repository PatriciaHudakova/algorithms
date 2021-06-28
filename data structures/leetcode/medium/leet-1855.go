/*
https://leetcode.com/problems/maximum-distance-between-a-pair-of-values/

Input: nums1 = [55,30,5,4,2], nums2 = [100,20,10,10,5]
Output: 2
Explanation: The valid pairs are (0,0), (2,2), (2,3), (2,4), (3,3), (3,4), and (4,4).
The maximum distance is 2 with pair (2,4).
*/
package main

import "fmt"

func main() {
	nums1, nums2 := []int{55, 30, 5, 4, 2}, []int{100, 20, 10, 10, 5}

	max := maxDistance(nums1, nums2)

	fmt.Println(max)
}

func maxDistance(nums1 []int, nums2 []int) int {
	maxDistance := 0

	for i, v := range nums1 {
		for j, w := range nums2 {
			if i <= j && v <= w {
				distance := j - i
				if distance > maxDistance {
					maxDistance = distance
				}
			}
		}
	}

	return maxDistance
}
