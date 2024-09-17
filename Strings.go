package main

import (
	"regexp"
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
