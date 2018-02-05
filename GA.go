package ga_8queens

import (
	"fmt"
	"math"
)

func StartGA(initialPopulation []string) {
	// TODO
	for _, t := range initialPopulation {
		fmt.Printf("fitness for \"%s\": %d\n", t, fitness(t))
	}
}

func addToSet(set []string, a string) bool {
	var contains bool
	for _, b := range set {
		if a == b {
			contains = true
		}
	}
	if !contains {
		set = append(set, a)
	}
	return !contains
}

// should return the number of non-attacking pairs of queens
func fitness(individual string) int {
	var height []int
	for _, t := range individual {
		height = append(height, int(t%48))
	}
	collisions := 0
	// find collisions
	for i := 0; i < len(height); i++ {
		for j := 0; j < len(height); j++ {
			if i != j {
				if height[i] == height[j] || int(math.Abs(float64(height[j]-height[i]))) == int(math.Abs(float64(j-i))) {
					fmt.Printf("collision: (%d, %d), (%d, %d)\n", i+1, height[i], j+1, height[j])
					collisions++
				}
			}
		}
	}
	max := len(height) * (len(height) - 1)
	return max - collisions
}

func GA(population []string, fitness func(string) int) string {
	// TODO
	return ""
}

func Reproduce(individualX string, individualY string) string {
	// TODO
	return ""
}

func Mutate(individual string) {
	// TODO
}
