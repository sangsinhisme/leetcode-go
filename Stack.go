package main

import "strings"

type stack[T any] struct {
	Push   func(T)
	Pop    func() T
	Length func() int
}

func Stack[T any]() stack[T] {
	slice := make([]T, 0)
	return stack[T]{
		Push: func(i T) {
			slice = append(slice, i)
		},
		Pop: func() T {
			res := slice[len(slice)-1]
			slice = slice[:len(slice)-1]
			return res
		},
		Length: func() int {
			return len(slice)
		},
	}
}

/*
1106. Parsing A Boolean Expression
https://leetcode.com/problems/parsing-a-boolean-expression/description/
*/
func parseBoolExpr(expression string) bool {
	if expression[0] == 't' {
		return true
	}
	if expression[0] == 'f' {
		return false
	}
	exps := expression[2 : len(expression)-1]
	if expression[0] == '!' {
		return !parseBoolExpr(exps)
	}
	exp := splitByComma(exps)
	parser := parseBoolExpr(exp[0])

	if expression[0] == '&' {
		if len(exp) > 1 {
			for i := 1; i < len(exp); i++ {
				parser = parser && parseBoolExpr(exp[i])
			}
		}
		return parser
	} else {
		if len(exp) > 1 {
			for i := 1; i < len(exp); i++ {
				parser = parser || parseBoolExpr(exp[i])
			}
		}
		return parser
	}
}

func splitByComma(s string) []string {
	var result []string
	var current strings.Builder
	level := 0

	for _, char := range s {
		switch char {
		case ',':
			if level == 0 {
				result = append(result, strings.TrimSpace(current.String()))
				current.Reset()
			} else {
				current.WriteRune(char)
			}
		case '(':
			level++
			current.WriteRune(char)
		case ')':
			level--
			current.WriteRune(char)
		default:
			current.WriteRune(char)
		}
	}
	if current.Len() > 0 {
		result = append(result, strings.TrimSpace(current.String()))
	}
	return result
}
