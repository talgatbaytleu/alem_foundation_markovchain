package main

import (
	"markov/cmd"
	"markov/internal"
)

func main() {
	a := cmd.GenerateMap()
	cmd.GenerateText(internal.Default_prefix, a)
}
