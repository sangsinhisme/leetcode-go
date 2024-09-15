package main

import "math"

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
	result := []int{}

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
