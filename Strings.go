package main

import (
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

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
