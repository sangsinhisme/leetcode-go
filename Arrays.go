package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

/*
650. 2 Keys Keyboard
There is only one character 'A' on the screen of a notepad. You can perform one of two operations on this notepad for each step:

Copy All: You can copy all the characters present on the screen (a partial copy is not allowed).
Paste: You can paste the characters which are copied last time.
Given an integer n, return the minimum number of operations to get the character 'A' exactly n times on the screen.
. . .
*/

func minSteps(n int) int {
	var factors = 0
	for n%2 == 0 {
		n = n / 2
		factors += 2
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			n = n / i
			factors += i
		}
	}

	if n > 2 {
		factors += n
	}
	return factors
}

func getSneakyNumbers(nums []int) []int {
	freqMap := frequency(nums)
	var result []int

	// Loop through the frequency map and select numbers that have a frequency of 2
	for num, count := range freqMap {
		if count == 2 {
			result = append(result, num)
		}
	}

	return result
}

func frequency(nums []int) map[int]int {
	freqMap := make(map[int]int)

	// Loop through each number in the slice and update its frequency count
	for _, num := range nums {
		freqMap[num]++
	}

	return freqMap
}

func findMinDifference(timePoints []string) int {
	var helper func(time string) int
	helper = func(time string) int {
		parser := strings.Split(time, ":")
		hour, _ := strconv.Atoi(parser[0])
		minute, _ := strconv.Atoi(parser[1])
		return hour*60 + minute
	}
	n := len(timePoints)

	times := make([]int, n)
	for i, time := range timePoints {
		times[i] = helper(time)
	}
	sort.Ints(times)
	minDiff := 1440 + times[0] - times[n-1]
	for i := 1; i < n; i++ {
		minDiff = min(minDiff, times[i]-times[i-1])
	}
	return minDiff
}

/*
386. Lexicographical Numbers
Given an integer n, return all the numbers in the range [1, n] sorted in lexicographical order.
You must write an algorithm that runs in O(n) time and uses O(1) extra space.
*/
func lexicalOrder(n int) []int {
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	sort.Slice(nums, func(i, j int) bool {
		num1 := strconv.Itoa(nums[i])
		num2 := strconv.Itoa(nums[j])
		if num1[0] == num2[0] {
			return num1 < num2
		}
		concat1, _ := strconv.Atoi(num1 + num2)
		concat2, _ := strconv.Atoi(num2 + num1)
		return concat1 < concat2
	})
	return nums
}

/*
3043. Find the Length of the Longest Common Prefix
You are given two arrays with positive integers arr1 and arr2.
A prefix of a positive integer is an integer formed by one or more of its digits, starting from its leftmost digit.
Return the length of the longest common prefix among all pairs. If no common prefix exists among them, return 0.
*/
func longestCommonPrefix(arr1 []int, arr2 []int) int {
	prefix1 := make(map[int]bool)
	for _, num := range arr1 {
		for num > 0 {
			prefix1[num] = true
			num = num / 10
		}
	}
	maxPrefix := 0
	for _, num := range arr2 {
		for num > 0 {
			if prefix1[num] {
				length := lenLoop(num)
				if length > maxPrefix {
					maxPrefix = length
				}
				break
			}
			num = num / 10
		}
	}
	return maxPrefix
}

func lenLoop(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}
