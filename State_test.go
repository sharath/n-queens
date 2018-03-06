package main

import (
	"testing"
	"fmt"
)

func TestLoadState(t *testing.T) {
	{
		s := LoadState("initial.json", "8")
		if s == nil {
			t.FailNow()
		}
		if len(s.state) != 8 {
			fmt.Println(len(s.state))
			t.FailNow()
		}
		for _, i := range s.state {
			if i != 1 {
				fmt.Println(i)
				t.FailNow()
			}
		}
	}
	{
		s := LoadState("initial.json", "4")
		if s == nil {
			t.FailNow()
		}
		if len(s.state) != 4 {
			fmt.Println(len(s.state))
			t.FailNow()
		}
		for _, i := range s.state {
			if i != 1 {
				fmt.Println(i)
				t.FailNow()
			}
		}
	}
}
