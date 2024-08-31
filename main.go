package main

import (
	"fmt"
	"io"
)

func main() {
	var next_word string
	var prev_word string
	data_map := make(map[string][]string)
	for {
		_, err := fmt.Scan(&next_word)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			break
		}
		data_map[prev_word] = append(data_map[prev_word], next_word)
		prev_word = next_word
	}
	for i, v := range data_map {
		fmt.Print(i + ":")
		fmt.Println(v)
	}
	// cmd.Markov()
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// input := scanner.Text()
	// fmt.Println(input)
	// scanner.Scan()
	// input = scanner.Text()
	// fmt.Println(input)
	// scanner.Scan()
	// input = scanner.Text()
	// fmt.Println(input)
	// scanner.Scan()
	// input = scanner.Text()
	// fmt.Println(input)
	// scanner.Scan()
	// input = scanner.Text()
	// fmt.Println(input)
	// scanner.Scan()
	// input = scanner.Text()
	// fmt.Println(input)
	// scanner.Scan()
	// input = scanner.Text()
	// fmt.Println(input)
}
