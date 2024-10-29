package main

import "fmt"

func main() {
	board := [][]byte{{'a', 'b', 'c'}, {'a', 'e', 'd'}, {'a', 'f', 'g'}}
	words := []string{"abcdefg", "gfedcbaaa", "eaabcdgfa", "befa", "dgc", "ade"}
	fmt.Println(findWords(board, words))
}
