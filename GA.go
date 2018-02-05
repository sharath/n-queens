package ga_8queens

import (
	"math"
	"math/rand"
	"time"
	"fmt"
	"strconv"
)

var size int
var threshold float64
var verbose bool

func StartGA(initialPopulation []string, iterations int, s int, th float64, v bool) {
	verbose = v
	rand.Seed(int64(time.Now().Second() * time.Now().Nanosecond()))
	size = s
	threshold = th
	fmt.Println("Started GA. This might take a while...")
	population := initialPopulation
	var solved bool
	var solution string
	var g int
	for !solved {
		g++
		population = GAIteration(population)
		if verbose {
			fmt.Printf("Generation %d\n", g)
		}
		for _, t := range population {
			if Fitness(t) == (size * (size - 1)) {
				solution = t
				solved = true
			}
		}
	}
	fmt.Printf("Solution: \"%s\" in %d Generations\n", solution, g)
}

// should return the number of non-attacking pairs of queens
func Fitness(individual string) int {
	var height []int
	for _, c := range []byte(individual) {
		height = append(height, int(c%48))
	}
	collisions := 0
	// find collisions
	for i := 0; i < len(height); i++ {
		for j := 0; j < len(height); j++ {
			if i != j {
				// if height is same or if the absolute value of the slope is equal to 1, there exists a collision
				if height[i] == height[j] || math.Pow(float64(height[j]-height[i]), 2) == math.Pow(float64(j-i), 2) {
					collisions++
				}
			}
		}
	}
	return (size * (size - 1)) - collisions
}

// choose randomly from population favoring fit individuals
// not sure if i implemented this correctly
func RandomSelection(population []string) string {
	a := population[rand.Int()%len(population)]
	b := population[rand.Int()%len(population)]
	if Fitness(a) > Fitness(b) {
		return a
	} else {
		return b
	}
}

func GAIteration(population []string) []string {
	var newPopulation []string
	for i := 0; i < len(population); i++ {
		x := RandomSelection(population)
		y := RandomSelection(population)
		child1, _ := Reproduce(x, y)
		child1 = Mutate(child1)
		//child2 = Mutate(child2)
		newPopulation = AddIfNotContains(newPopulation, child1)
		//newPopulation = AddIfNotContains(newPopulation, child2)
	}
	return newPopulation
}
func AddIfNotContains(set []string, value string) []string {
	var contains bool
	for _, t := range set {
		if t != value {
			contains = true
		}
	}
	if !contains {
		set = append(set, value)
	}
	return set
}

//generate a random number c from 1 to len(x) and split into substrings
func Reproduce(individualX string, individualY string) (string, string) {
	r := rand.Int() % size
	n := (r * r) % size
	return individualX[n:] + individualY[:n], individualY[n:] + individualX[:n]
}

func Mutate(c string) string {
	if rand.Float64() < threshold {
		r := strconv.Itoa(rand.Int()%size + 1)
		i := rand.Int() % size
		var n string
		for j, ch := range c {
			if j != i {
				n += strconv.Itoa(int(ch % 48))
			} else {
				n += r
			}
		}
		if verbose {
			fmt.Printf("Mutated! \"%s\" -> \"%s\"\n", c, n)
		}
		return n
	} else {
		if verbose {
			fmt.Printf("\"%s\"\n", c)
		}
		return c
	}
}
