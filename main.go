package main

import "fmt"

func main() {
	board := [][]byte{{'a', 'b'}, {'a', 'a'}}
	words := []string{"aba", "baa", "bab", "aaab", "aaa", "aaaa", "aaba"}
	fmt.Println(findWords(board, words))
}
