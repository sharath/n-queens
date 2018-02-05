package ga_8queens

import (
	"math"
	"math/rand"
	"time"
	"fmt"
	"strconv"
)

var size int

func StartGA(initialPopulation []string, iterations int, s int) {
	size = s
	fmt.Println("Started GA. This might take a while...")
	population := initialPopulation
	var solved bool
	var solution string
	var i int
	for !solved {
		i++
		population = GAIteration(population)
		for _, t := range population {
			if Fitness(t) == (size * (size - 1)) {
				solution = t
				solved = true
			}
		}
	}
	fmt.Printf("Solution: \"%s\" in %d Iterations\n", solution, i)
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
func RandomSelection(population []string) string {
	/*// TODO: favor more fit individuals
	a := population[rand.Int()%(len(population)-1)]
	b := population[rand.Int()%(len(population)-1)]
	if Fitness(a) > Fitness(b) {
		return a
	} else {
		return b
	}*/
	return population[rand.Int()%(len(population)-1)]
}

func GAIteration(population []string) []string {
	var newPopulation []string
	threshold := 0.4
	for i := 0; i < len(population); i++ {
		x := RandomSelection(population)
		y := RandomSelection(population)
		child1, child2 := Reproduce(x, y)
		child1 = Mutate(child1, threshold)
		child2 = Mutate(child2, threshold)
		newPopulation = append(newPopulation, child1)
		newPopulation = append(newPopulation, child2)
	}
	return newPopulation
}

//generate a random number c from 1 to len(x) and split into substrings
func Reproduce(individualX string, individualY string) (string, string) {
	rand.Seed(int64(time.Now().Nanosecond()))
	r := rand.Int() % len(individualX)
	n := (r * r) % len(individualX)
	return individualX[n:] + individualY[:n], individualY[n:] + individualX[:n]
}

func Mutate(c string, threshold float64) string {
	rand.Seed(int64(time.Now().Nanosecond()))
	if float64(rand.Uint64()%10000)/10000.0 < threshold {
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
		return n
	} else {
		return c
	}
}
