package main

import (
	"fmt"

	"markov/cmd"
)

func main() {
	a := cmd.GenerateMap(1)
	for i, v := range a {
		fmt.Print(i + ": ")
		fmt.Println(v)
	}
}
