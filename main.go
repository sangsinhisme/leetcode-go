package main

import (
	"fmt"
	_ "fmt"
)

func main() {

	fmt.Print(countConsistentStrings("abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}))
}
