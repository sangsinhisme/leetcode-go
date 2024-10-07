package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

/*
*
564. Find the Closest Palindrome
Given a string n representing an integer, return the closest integer (not including itself), which is a palindrome.
If there is a tie, return the smaller one.
The closest is defined as the absolute difference minimized between two integers.
Example 1:

	Input: n = "123"
	Output: "121"

Example 2:

	Input: n = "1"
	Output: "0"
*/
func nearestPalindromic(n string) string {
	length := len(n)
	num, _ := strconv.Atoi(n)
	if length == 1 {
		return strconv.Itoa(num - 1)
	}
	if n[0] == '1' {
		if n[1:] == strings.Repeat("0", length-1) {
			return strconv.Itoa(num - 1)
		}
	}
	if n[0] == '9' {
		if n[1:] == strings.Repeat("9", length-1) {
			return strconv.Itoa(num + 2)
		}
	}
	nR := reverse(n)
	if n == nR {
		if n == "11" {
			return "9"
		}
		palindDown := palindrome(num - power10(length/2))
		palindUp := palindrome(num + power10(length/2))
		cPalindDown, _ := strconv.Atoi(palindDown)
		cPalindUp, _ := strconv.Atoi(palindUp)
		cPalindDown = abs(num - cPalindDown)
		cPalindUp = abs(num - cPalindUp)
		if cPalindDown <= cPalindUp {
			return palindDown
		}
		return palindUp
	}
	palindDown := palindrome(num - power10(length/2))
	palindUp := palindrome(num + power10(length/2))
	palindMiddle := palindrome(num)
	cPalindDown, _ := strconv.Atoi(palindDown)
	cPalindDown = abs(num - cPalindDown)
	cPalindMiddle, _ := strconv.Atoi(palindMiddle)
	cPalindMiddle = abs(num - cPalindMiddle)
	cPalindUp, _ := strconv.Atoi(palindUp)
	cPalindUp = abs(num - cPalindUp)
	if cPalindDown <= cPalindUp {
		if cPalindDown <= cPalindMiddle {
			return palindDown
		}
		return palindMiddle
	}
	if cPalindMiddle <= cPalindUp {
		if cPalindMiddle < cPalindDown {
			return palindMiddle
		}
		return palindDown
	}
	return palindUp
}
func power10(n int) int {
	power, _ := strconv.Atoi("1" + strings.Repeat("0", n))
	return power
}
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func palindrome(num int) string {
	s := strconv.Itoa(num)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			runes[j] = runes[i]
		}
	}
	return string(runes)
}

/*
1684. Count the Number of Consistent Strings
You are given a string allowed consisting of distinct characters and an array of strings words.
A string is consistent if all characters in the string appear in the string allowed.

Return the number of consistent strings in the array words.
*/
func countConsistentStrings(allowed string, words []string) int {
	r, _ := regexp.Compile("[" + allowed + "]+")
	allow := 0
	for _, word := range words {
		regex := r.FindAllString(word, -1)
		if len(regex) == 1 && len(regex[0]) == len(word) {
			allow += 1
		}
	}
	return allow
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

/*
884. Uncommon Words from Two Sentences
A sentence is a string of single-space separated words where each word consists only of lowercase letters.
A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.
Given two sentences s1 and s2, return a list of all the uncommon words. You may return the answer in any order.
*/
func uncommonFromSentences(s1 string, s2 string) (res []string) {
	freqMap := make(map[string]int)
	words1 := strings.Split(s1, " ")
	words2 := strings.Split(s2, " ")
	for _, word := range words1 {
		freqMap[word]++
	}
	for _, word := range words2 {
		freqMap[word]++
	}
	for k, v := range freqMap {
		if v == 1 {
			res = append(res, k)
		}
	}
	return
}

/*
179. Largest Number
Given a list of non-negative integers nums, arrange them such that they form the largest number and return it.
Since the result may be very large, so you need to return a string instead of an integer.
Example 1:
Input: nums = [10,2]
Output: "210"
*/
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		num1 := strconv.Itoa(nums[i])
		num2 := strconv.Itoa(nums[j])
		concat1, _ := strconv.Atoi(num1 + num2)
		concat2, _ := strconv.Atoi(num2 + num1)
		return concat1 > concat2
	})
	output := ""
	for _, elem := range nums {
		output += strconv.Itoa(elem)
	}
	length := len(output)
	trimmed := strings.TrimLeft(output[:length-1], "0")
	result := trimmed + output[length-1:]

	return result
}

/*
214. Shortest Palindrome
You are given a string s. You can convert s to a palindrome by adding characters in front of it.
Return the shortest palindrome you can find by performing this transformation
*/
func shortestPalindrome(s string) string {
	reverse := reverse(s)
	i := 0
	n := len(s)
	flag := true
	for flag && i < n {
		if !strings.Contains(reverse, s[:i+1]) {
			flag = false
		} else {
			i++
		}
	}
	flag = true
	for flag {
		fmt.Println(reverse[n-i:])
		fmt.Println(s[:i+1])
		if reverse[n-i:] == s[:i] {
			flag = false
		} else {
			i--
		}
	}
	return reverse[:n-i] + s
}

/*
3043. Find the Length of the Longest Common Prefix
You are given two arrays with positive integers arr1 and arr2.
A prefix of a positive integer is an integer formed by one or more of its digits, starting from its leftmost digit.
Return the length of the longest common prefix among all pairs. If no common prefix exists among them, return 0.
*/
func longestCommonPrefix2(arr1 []int, arr2 []int) int {
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

func sumPrefixScores(words []string) []int {
	prefixes := make(map[string][]string)
	for _, word := range words {
		var temp []string
		for i := range word {
			temp = append(temp, word[:len(word)-i])
		}
		prefixes[word] = temp
	}
	freqMap := make(map[string]int)
	for _, value := range words {
		for _, word := range prefixes[value] {
			freqMap[word]++
		}
	}
	var output []int
	for _, value := range words {
		total := 0
		for _, word := range prefixes[value] {
			total += freqMap[word]
		}
		output = append(output, total)
	}
	return output
}

/*
567. Permutation in String
https://leetcode.com/problems/permutation-in-string/description/
*/
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	freq1 := make([]int, 26)
	freq2 := make([]int, 26)
	for i := range s1 {
		freq1[s1[i]-'a']++
		freq2[s2[i]-'a']++
	}
	for i := 0; i < len(s2)-len(s1); i++ {
		if matches(freq1, freq2) {
			return true
		}
		freq2[s2[i+len(s1)]-'a']++
		freq2[s2[i]-'a']--
	}
	return matches(freq1, freq2)

}

func matches(freq1 []int, freq2 []int) bool {
	for i := 0; i < 26; i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}
	return true
}

func minLength(s string) int {
	for strings.Contains(s, "AB") || strings.Contains(s, "CD") {
		s = strings.Replace(s, "AB", "", -1)
		s = strings.Replace(s, "CD", "", -1)
	}
	return len(s)
}
