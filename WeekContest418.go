package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
Q1. Maximum Possible Number by Binary Concatenation
*/
func maxGoodNumber(nums []int) int {
	binaries := make([]string, len(nums))
	for i, num := range nums {
		binaries[i] = num2binary2(num)
	}

	sort.Slice(binaries, func(i, j int) bool {
		return binaries[i]+binaries[j] > binaries[j]+binaries[i]
	})

	var result string
	for _, binary := range binaries {
		result += binary
	}
	output := binaryToNum(result)
	return output
}

func num2binary2(num int) string {
	if num == 0 {
		return "0"
	}
	if num == 1 {
		return "1"
	}
	if num%2 == 1 {
		return num2binary2(num/2) + "1"
	}
	return num2binary2(num/2) + "0"
}

func binaryToNum(binaryStr string) int {
	// Convert binary string to integer
	num, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error converting binary to number:", err)
		return 0
	}
	return int(num)
}

/*
Q2. Remove Methods From Project
*/
func remainingMethods(n int, k int, invocations [][]int) []int {
	adjList := make(map[int][]int, n)
	invokedBy := make(map[int][]int, n)

	for _, invocation := range invocations {
		invoker, invoked := invocation[0], invocation[1]
		adjList[invoker] = append(adjList[invoker], invoked)
		invokedBy[invoked] = append(invokedBy[invoked], invoker)
	}

	suspicious := make(map[int]bool)
	var dfs func(int)
	dfs = func(method int) {
		if suspicious[method] {
			return
		}
		suspicious[method] = true
		for _, next := range adjList[method] {
			dfs(next)
		}
	}
	dfs(k)

	for method := range suspicious {
		for _, invoker := range invokedBy[method] {
			if !suspicious[invoker] {
				return makeRange(0, n-1)
			}
		}
	}

	// Step 4: Collect remaining methods
	var remaining []int
	for i := 0; i < n; i++ {
		if !suspicious[i] {
			remaining = append(remaining, i)
		}
	}
	return remaining
}

func makeRange(start, end int) []int {
	r := make([]int, end-start+1)
	for i := range r {
		r[i] = start + i
	}
	return r
}

/*
Q3. Construct 2D Grid Matching Graph Layout
*/
func constructGridLayout(n int, edges [][]int) [][]int {

	return nil
}

/*
Q4. Sorted GCD Pair Queries
*/
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdValues(nums []int, queries []int64) []int {
	n := len(nums)
	gcdMap := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			g := gcd(nums[i], nums[j])
			gcdMap[g]++
		}
	}

	var uniqueGCDs []int
	for g := range gcdMap {
		uniqueGCDs = append(uniqueGCDs, g)
	}

	// Sort unique GCDs
	sort.Ints(uniqueGCDs)

	result := make([]int, 0, len(queries))
	for _, q := range queries {
		index := int(q)
		if index < len(uniqueGCDs) {
			gcdValue := uniqueGCDs[index]
			count := gcdMap[gcdValue]

			for i := 0; i < count; i++ {
				result = append(result, gcdValue)
			}
		}
	}

	return result
}
