/*
https://leetcode.com/problems/slowest-key/

Input: releaseTimes = [9,29,49,50], keysPressed = "cbcd"
Output: "c"
Explanation: The keypresses were as follows:
Keypress for 'c' had a duration of 9 (pressed at time 0 and released at time 9).
Keypress for 'b' had a duration of 29 - 9 = 20 (pressed at time 9 right after the release of the previous character and released at time 29).
Keypress for 'c' had a duration of 49 - 29 = 20 (pressed at time 29 right after the release of the previous character and released at time 49).
Keypress for 'd' had a duration of 50 - 49 = 1 (pressed at time 49 right after the release of the previous character and released at time 50).
The longest of these was the keypress for 'b' and the second keypress for 'c', both with duration 20.
'c' is lexicographically larger than 'b', so the answer is 'c'.
*/

package main

import "fmt"

func main() {
	byte := slowestKey([]int{1, 2}, "ba")
	fmt.Println(byte)
}

func slowestKey(releaseTimes []int, keysPressed string) byte {
	// create a duration array
	duration := make([]int, len(releaseTimes))

	// populate the duration array with duration times
	for key, value := range releaseTimes {
		if key == 0 {
			duration[0] = value
		} else {
			duration[key] = releaseTimes[key] - releaseTimes[key-1]
		}
	}

	// pick the highest duration time from array
	highestValue := 0
	highestIndex := 0

	for key, value := range duration {
		if value > highestValue {
			highestValue = value
			highestIndex = key
		}
		// check which contender is lexicographically larger
		if value == highestValue {
			contender1 := keysPressed[key]
			contender2 := keysPressed[highestIndex]
			if contender1 >= contender2 {
				highestValue = value
				highestIndex = key
			}
		}
	}

	// check what letter this highest duration corresponds to
	letter := keysPressed[highestIndex]
	return letter
}
