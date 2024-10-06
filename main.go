package main

import (
	"fmt"
)

func main() {
	nums := []int{12, 15, 18}
	queries := []int64{0, 1, 2}
	result := gcdValues(nums, queries)
	fmt.Println(result)
}
