package utils

import "math/rand"

// WeightedPick returns an index chosen from weights.
func WeightedPick(weights []int) int {
	total := 0
	for _, w := range weights {
		total += w
	}
	r := rand.Intn(total)
	acc := 0
	for i, w := range weights {
		acc += w
		if r < acc {
			return i
		}
	}
	return len(weights) - 1
}
