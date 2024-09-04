package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	commands := []int{4, -1, 4, -2, 4}
	var obstacles [][]int
	fmt.Println(robotSim(commands, obstacles))
}
