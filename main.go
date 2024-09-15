package main

import (
	"fmt"
	"math"
)

func minValidStrings(words []string, target string) int {
	// Build a set of valid prefixes
	prefixSet := make(map[string]struct{})
	for _, word := range words {
		for i := 1; i <= len(word); i++ {
			prefixSet[word[:i]] = struct{}{}
		}
	}

	// Helper function to check if a prefix is valid
	isValidPrefix := func(prefix string) bool {
		_, exists := prefixSet[prefix]
		return exists
	}

	// Initialize dp array where dp[i] represents the minimum number of valid strings to form the substring target[:i]
	dp := make([]int, len(target)+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0 // Base case: 0 valid strings to form an empty target string

	// Fill the dp array
	for i := 1; i <= len(target); i++ {
		for j := 0; j < i; j++ {
			if dp[j] != math.MaxInt32 && isValidPrefix(target[j:i]) {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}

	// If dp[len(target)] is still the maximum value, it means it is not possible to form target
	if dp[len(target)] == math.MaxInt32 {
		return -1
	}
	return dp[len(target)]
}

func main() {
	words := []string{"ab", "bc", "abc", "bcd"}
	target := "abcd"
	fmt.Println(minValidStrings(words, target)) // Output will be the minimum number of valid strings to form the target
}
