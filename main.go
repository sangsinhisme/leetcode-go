package main

import (
	"fmt"
	_ "fmt"
	"math"
)

func minSteps(n int) int {
	var factos int = 0
	for n % 2 == 0 {
		n = n / 2
		factos += 2
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n % i == 0 {
			n = n / i
			factos += i
		}
	}

	if n > 2 {
		factos += n
	}
	return factos
}

func main() {

	var n int = 36
	fmt.Println(minSteps(n))
}
