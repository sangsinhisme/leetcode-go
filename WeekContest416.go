package main

import (
	"math"
)

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
	freWord2 := make([]int, 26)
	count := 0
	for _, char := range word2 {
		if freWord2[char-'a'] == 0 {
			count++
		}
		freWord2[char-'a']++
	}
	var ans int64 = 0
	n := len(word1)
	j := 0
	for i := 0; i < n; i++ {
		k := word1[i] - 'a'
		freWord2[k]--
		if freWord2[k] == 0 {
			count--
		}
		for count == 0 {
			ans += int64(n - i)
			p := word1[j] - 'a'
			freWord2[p]++
			if freWord2[p] == 1 {
				count++
			}
			j++
		}
	}
	return ans
}
