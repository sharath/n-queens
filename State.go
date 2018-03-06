package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
)

type State struct {
	state []int
	size int
}

func LoadState(filename string, name string) *State {
	js, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: Can't Read Input File")
		os.Exit(-1)
	}
	fmt.Println("Validating Inputs...")
	temp := make(map[string][]int)
	s := new(State)
	err = json.Unmarshal(js, &temp)
	if err != nil {
		fmt.Println("Error: Can't Parse JSON")
		os.Exit(-1)
	}
	s.state = temp[name]
	s.size = len(s.state)
	if s.size < 4 {
		fmt.Println("Error: Invalid Initial State")
		os.Exit(-1)
	}
	fmt.Printf("Success: Recognized %d-Queens Problem\n", s.size)
	return s
}