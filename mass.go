package main

import (
	"fmt"
)

func activate(sum float64) int {
	if sum > 0 {
		return 1
	}

	return 0
}

func main() {
	inputs := []float64{12, 4}
	weights := []float64{0.5,-1}
	sum := 0.0

	for i := 0; i < len(inputs); i++ {
		sum += inputs[i] * weights[i]
	}

	output := activate(sum)

	fmt.Println(inputs, weights, sum, output)
}
