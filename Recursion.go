package main

import (
	"fmt"
	"strings"
)

/*
664. Strange Printer
There is a strange printer with the following two special properties:
  - The printer can only print a sequence of the same character each time.
  - At each turn, the printer can print new characters starting from and ending at any place and will cover the original existing characters.

Given a string s, return the minimum number of turns the printer needed to print it.
*/
func strangePrinter(s string) int {
	return helperPrinter(s, []uint8{}, 0)
}

func extractNextString(s string, startChar uint8, memo []uint8) (string, string) {
	if len(s) == 0 {
		return "", ""
	}
	fmt.Println("start:" + s)
	for i := 1; i < len(s); i++ {
		if s[i] != startChar {
			lastIndex := strings.LastIndex(s, s[i:i+1])
			fmt.Println(s[i:lastIndex+1], trimString(s[lastIndex+1:], memo))
			return s[i : lastIndex+1], trimString(s[lastIndex+1:], memo)
		}
	}
	return "", ""
}

func helperPrinter(s string, memo []uint8, step int) int {
	if s == "" {
		return step
	}
	startChar := s[0]
	lastChar := strings.LastIndex(s, s[0:1])

	nextString1, nextString2 := extractNextString(s[:lastChar], startChar, append(memo, startChar))
	return helperPrinter(nextString1, append(memo, startChar), step+1) + helperPrinter(nextString2, append(memo, startChar), 0) + helperPrinter(s[lastChar+1:], memo, 0)
}

func trimString(s string, memo []uint8) string {
	start := 0
	end := len(s) - 1

	for start < len(s) && contains(memo, s[start]) {
		start++
	}
	for end >= 0 && contains(memo, s[end]) {
		end--
	}
	if start > end {
		return ""
	}
	return s[start : end+1]
}

func contains(memo []uint8, char uint8) bool {
	for _, c := range memo {
		if c == char {
			return true
		}
	}
	return false
}
