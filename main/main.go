package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
	"github.com/sharath/ga-8queens"
)

func ReadFile(name string) ([]string, int) {
	configFile, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("Error: Can't Read Input File")
		os.Exit(-1)
	}
	var states []string
	err = json.Unmarshal(configFile, &states)
	if err != nil {
		fmt.Println("Error: Can't Parse Input File")
	}
	var n int
	fmt.Println("Validating Inputs...")
	for _, t1 := range states {
		for _, t2 := range states {
			if len(t1) != len(t2) {
				fmt.Println("Error: Invalid Input File")
				os.Exit(-1)
			}
		}
		n = len(t1)
	}
	fmt.Printf("Success: Recognized %d-Queens Problem\n", n)
	return states, n
}

func main() {
	var filename string
	if len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		filename = "initial_population.json"
	}
	initialPopulation, n := ReadFile(filename)
	ga_8queens.StartGA(initialPopulation, 5, n, 0.4, true)
}
