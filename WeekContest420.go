package main

/*
Q1. Find the Sequence of Strings Appeared on the Screen
*/
func stringSequence(target string) []string {
	var output []string
	var history string

	for _, char := range target {
		var current rune
		current = 'a'
		output = append(output, history+string(current))
		for current != char {
			current += 1
			output = append(output, history+string(current))
		}
		history = history + string(current)
	}
	return output
}

/*
Q2. Count Substrings With K-Frequency Characters I
*/
func numberOfSubstrings(s string, k int) int {
	output := 0
	n := len(s)
	for i := 0; i < n; i++ {
		freq := make(map[rune]int)
		maxFreq := 0
		for j := i; j < n; j++ {
			freq[rune(s[j])]++
			maxFreq = max(maxFreq, freq[rune(s[j])])

			if maxFreq >= k {
				output += n - j
				break
			}
		}
	}
	return output
}

/*
Q3. Minimum Division Operations to Make Array Non-Decreasing
*/
func minOperations(nums []int) int {
	n := len(nums)
	divisions := make([]int, n)
	copy(divisions, nums)

	operations := 0

	for i := n - 2; i >= 0; i-- {
		if divisions[i] > divisions[i+1] {
			divisor := minimalProperDivisor(divisions[i], divisions[i+1])
			if divisor == -1 {
				return -1
			}
			divisions[i] = divisor
			operations++
		}
	}

	return operations
}

func minimalProperDivisor(x int, target int) int {
	if x <= 2 || target < 2 {
		return -1
	}

	for i := target; i >= 2; i-- {
		if x%i == 0 {
			return i
		}
	}
	return -1
}

/*
Q4. Check if DFS Strings Are Palindromes
*/
func buildTree(parent []int) map[int][]int {
	tree := make(map[int][]int)
	for i, p := range parent {
		if p != -1 {
			tree[p] = append(tree[p], i)
		}
	}
	return tree
}

func dfs2(node int, currentStr string, s string, result []string, tree map[int][]int) {
	currentStr += string(s[node])
	result[node] = currentStr
	for _, child := range tree[node] {
		dfs2(child, currentStr, s, result, tree)
	}
}

func findAnswer(parent []int, s string) []bool {
	tree := buildTree(parent)
	result := make([]string, len(s))
	dfs2(0, "", s, result, tree)
	output := make([]bool, len(result))
	for i, elem := range result {
		output[i] = isPalindrome(elem)
	}
	return output
}

func isPalindrome(str string) bool {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}
