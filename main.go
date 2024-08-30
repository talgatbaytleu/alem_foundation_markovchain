package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// cmd.Markov()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Println(input)
	scanner.Scan()
	input = scanner.Text()
	fmt.Println(input)
	scanner.Scan()
	input = scanner.Text()
	fmt.Println(input)
	scanner.Scan()
	input = scanner.Text()
	fmt.Println(input)
	scanner.Scan()
	input = scanner.Text()
	fmt.Println(input)
	scanner.Scan()
	input = scanner.Text()
	fmt.Println(input)
	scanner.Scan()
	input = scanner.Text()
	fmt.Println(input)
}
