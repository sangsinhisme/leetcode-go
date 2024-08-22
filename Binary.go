package main

import "math"

func findComplement(num int) int {
	return revert2num(num2binary(num, ""), 0)
}

func num2binary(num int, output string) string {
	if num == 0 {
		return output
	}
	if num == 1 {
		return output + "1"
	}
	if num%2 == 1 {
		return num2binary(num/2, output) + "1"
	}
	return num2binary(num/2, output) + "0"
}

func revert2num(num string, output int) int {
	power := len(num)
	if power == 0 {
		return 0
	}
	if power == 1 {
		if num[0] == '0' {
			return output
		}
		return output + 1
	}
	if num[0] == '1' {
		return revert2num(num[1:], output+int(math.Pow(2, float64(power-1))))
	}
	return revert2num(num[1:], output)
}
