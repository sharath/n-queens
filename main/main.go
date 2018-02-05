package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
	"github.com/sharath/ga-8queens"
)

func ReadFile(name string) []string {
	configFile, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("Error: Can't Read Input File")
		os.Exit(-1)
	}
	var states []string
	json.Unmarshal(configFile, &states)
	return states
}

func main() {
	var filename string
	if len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		filename = "initial_population.json"
	}
	initialPopulation := ReadFile(filename)
	ga_8queens.StartGA(initialPopulation)
}
