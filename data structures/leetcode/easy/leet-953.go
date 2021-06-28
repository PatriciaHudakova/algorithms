/*
https://leetcode.com/problems/verifying-an-alien-dictionary/

Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
Output: true
Explanation: As 'h' comes before 'l' in this language, then the sequence is sorted.
*/
package main

import "fmt"

func main() {
	result := isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz")
	fmt.Println(result)
}

func isAlienSorted(words []string, order string) bool {
	// base case
	if len(words) < 2 {
		return true
	}

	// create an empty map
	hash := make(map[byte]int)

	// create a map that maps lexicographical order to an order in string
	for i := 0; i < len(order); i++ {
		hash[order[i]] = i
		fmt.Println(hash)
	}

	for i := 0; i < len(words)-1; i++ {
		pointer := 0
		word1, word2 := words[i], words[i+1]

		// while the pointer points to a populated character
		for pointer < len(word1) && pointer < len(word2) {
			if hash[word1[pointer]] > hash[word2[pointer]] {
				return false
			}
			if hash[word1[pointer]] < hash[word2[pointer]] {
				break
			} else {
				pointer = pointer + 1
			}
		}

		// if the second word is shorter than first while equating in all previous characters
		if pointer < len(word1) && pointer >= len(word2) {
			return false
		}
	}
	return true
}
