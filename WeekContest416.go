package main

import "math"

/*
weekly-contest-416
Q1. Report Spam Message
*/
func reportSpam(message []string, bannedWords []string) bool {
	flag := 0
	spam := make(map[string]bool)
	for _, word := range bannedWords {
		spam[word] = true
	}
	for i := 0; i < len(message) && flag < 3; i++ {
		if spam[message[i]] {
			flag += 1
		}
	}
	return flag > 2
}

/*
weekly-contest-416
Q2. Minimum Number of Seconds to Make Mountain Height Zero
*/
func minNumberOfSeconds(mountainHeight int, workerTimes []int) int {
	canReduce := func(time int) bool {
		totalHeightReduced := 0
		for _, workerTime := range workerTimes {
			low, high := 0, mountainHeight
			for low <= high {
				mid := (low + high) / 2
				reductionTime := workerTime * mid * (mid + 1) / 2
				if reductionTime <= time {
					low = mid + 1
				} else {
					high = mid - 1
				}
			}
			totalHeightReduced += high
			if totalHeightReduced >= mountainHeight {
				return true
			}
		}
		return totalHeightReduced >= mountainHeight
	}

	left, right := 0, math.MaxInt64
	for left < right {
		mid := (left + right) / 2
		if canReduce(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

/*
weekly-contest-416
Q3. Count Substrings That Can Be Rearranged to Contain a String I
*/
func validSubstringCount(word1 string, word2 string) int64 {
	wordDict2 := make(map[rune]int)
	for _, r := range word2 {
		wordDict2[r]++
	}

	len1, len2 := len(word1), len(word2)
	validCount := int64(0)

	// Function to check if the currentDict meets the requirements in wordDict2
	isValid := func(currentDict map[rune]int) bool {
		for r, count := range wordDict2 {
			if currentDict[r] < count {
				return false
			}
		}
		return true
	}

	// Check all possible lengths from len2 to len1
	for length := len2; length <= len1; length++ {
		currentDict := make(map[rune]int)

		// Initialize the first window of the current length
		for i := 0; i < length; i++ {
			currentDict[rune(word1[i])]++
		}

		// Check if the initial window is valid
		if isValid(currentDict) {
			validCount++
		}

		// Slide the window
		for start := 1; start <= len1-length; start++ {
			// Remove the old character
			oldChar := rune(word1[start-1])
			currentDict[oldChar]--

			// Add the new character
			newChar := rune(word1[start+length-1])
			currentDict[newChar]++

			// Check if the current window is valid
			if isValid(currentDict) {
				validCount++
			}
		}
	}

	return validCount
}
