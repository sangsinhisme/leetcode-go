package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	rolls := []int{1, 2, 3, 4}
	fmt.Println(missingRolls(rolls, 6, 4))
}
