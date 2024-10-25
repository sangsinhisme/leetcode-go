package main

import "slices"

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

	memo := make(map[int]int)

	for i := n - 2; i >= 0; i-- {
		if divisions[i] > divisions[i+1] {
			divisor := memo[divisions[i]]
			if divisor == 0 {
				divisor = divisions[i] / minimalProperDivisor(divisions[i])
				memo[divisions[i]] = divisor
			}
			if divisor < 0 || divisor > divisions[i+1] {
				return -1
			}
			divisions[i] = divisor
			operations++
		}
	}

	return operations
}

func minimalProperDivisor(x int) int {
	for i := x / 2; i > 1; i-- {
		if x%i == 0 {
			return i
		}
	}
	return -1
}

/*
Q4. Check if DFS Strings Are Palindromes
*/
func findAnswer(pa []int, a string) (ans []bool) {
	n := len(pa)
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		v := pa[w]
		g[v] = append(g[v], w)
	}

	nodes := make([]struct{ l, r int }, len(g))
	dfn := -1
	var s []byte
	var buildDFN func(int) int
	buildDFN = func(v int) (size int) {
		slices.Sort(g[v])
		en := dfn + 1
		for _, w := range g[v] {
			sz := buildDFN(w)
			size += sz
		}
		dfn++
		nodes[v].l = en
		nodes[v].r = nodes[v].l + size
		size++
		s = append(s, a[v])
		return
	}
	buildDFN(0)

	t := append(make([]byte, 0, len(s)*2+3), '^')
	for _, c := range s {
		t = append(t, '#', c)
	}
	t = append(t, '#', '$')

	halfLen := make([]int, len(t)-2)
	halfLen[1] = 1
	boxM, boxR := 0, 0
	for i := 2; i < len(halfLen); i++ {
		hl := 1
		if i < boxR {
			hl = min(halfLen[boxM*2-i], boxR-i)
		}
		for t[i-hl] == t[i+hl] {
			hl++
			boxM, boxR = i, i+hl
		}
		halfLen[i] = hl
	}

	isP := func(l, r int) bool { return halfLen[l+r+2] > r-l+1 }

	for _, p := range nodes {
		res := isP(p.l, p.r)
		ans = append(ans, res)
	}

	return
}
