package ga_8queens

import "fmt"

func StartGA(initialPopulation []string) {
	// TODO
	for _, t := range initialPopulation {
		fmt.Println(fitness(t))
	}
}

// should return the number of non-attacking pairs of queens
func fitness(individual string) int {
	var height []int
	for _, t := range individual {
		height = append(height, int(t%48))
	}
	collisions := 0
	// find height collisions
	for i, h := range height {
		for j, c := range height {
			if i != j {
				if h == c || (c-h) == (j-i) {
					collisions++
					//fmt.Println("height")
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
