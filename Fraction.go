package main

import (
	"regexp"
	"strconv"
	"strings"
)

/*
592. Fraction Addition and Subtraction
Given a string expression representing an expression of fraction addition and subtraction,
return the calculation result in string format.

The final result should be an irreducible fraction. If your final result is an integer, change it to the format of a
fraction that has a denominator 1. So in this case, 2 should be converted to 2/1.
*/
func fractionAddition(expression string) string {
	var numerator = 0
	var denominator = 1
	if expression[0] != '-' {
		expression = "+" + expression
	}
	r, _ := regexp.Compile("[+-][^+-]*")
	fraction := r.FindAllString(expression, -1)
	for i := range fraction {
		var flag = 1
		if fraction[i][0] == '-' {
			flag = -1
		}
		split := strings.Split(fraction[i][1:], "/")
		i, _ := strconv.Atoi(split[0])
		j, _ := strconv.Atoi(split[1])
		numerator = numerator*j + flag*i*denominator
		denominator = denominator * j
		if numerator == 0 {
			denominator = 1
		} else {
			gcdCurr := gcd(numerator, denominator)
			numerator = numerator / gcdCurr
			denominator = denominator / gcdCurr
		}
	}
	return strconv.Itoa(numerator) + "/" + strconv.Itoa(denominator)
}

func gcd2(a int, b int) int {
	if a < 0 {
		return gcd(-a, b)
	}
	if a < b {
		return gcd(b, a)
	}
	if b == 0 {
		return a
	}
	return gcd(a%b, b)
}
