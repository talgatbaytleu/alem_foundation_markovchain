package cmd

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"

	"markov/internal"
)

func Markov() {
	if *internal.Help_flag {
		fmt.Print(
			"Markov Chain text generator.\n\nUsage:\n  markovchain [-w ,<N>] [-p <S>] [-l <N>]\n  markovchain --help\n\nOptions:\n  --help\tShow this screen.\n  -w N\tNumber of maximum words\n  -p S\tStarting prefix\n  -l N\tPrefix lenth\n",
		)
		os.Exit(0)
	}
	if *internal.Gen_text_len < 0 || *internal.Gen_text_len > 10000 {
		fmt.Fprintf(os.Stderr, "Error: Number of words is out of limits [0:10000]\n")
		os.Exit(1)
	}
	if *internal.Prefix_len < 0 || *internal.Prefix_len > 5 {
		fmt.Fprintf(os.Stderr, "Error: Prefix length is out of limits [0:5]\n")
		os.Exit(1)
	}

	data_map := GenerateMap()

	if *internal.Custom_prefix == "" {
		internal.Prefix = internal.Default_prefix
	} else {
		internal.Prefix = *internal.Custom_prefix
	}

	if len(strings.Split(internal.Prefix, " ")) != *internal.Prefix_len {
		fmt.Fprintf(os.Stderr, "Error: p and l flags conflict!!!\n")
		os.Exit(1)
	}

	_, exists := data_map[internal.Prefix]
	if !exists && len(data_map) > 0 {
		fmt.Fprintf(os.Stderr, "Error: Prefix doesn't present in the original text!!!\n")
		os.Exit(1)
	}

	GenerateText(internal.Prefix, data_map)
}

func GenerateMap() map[string][]string {
	res_map := make(map[string][]string)
	// helper_slice := make([]string, 0)
	var floating_prefix string = internal.Default_prefix
	var last_word string
	for {
		_, err := fmt.Scan(&last_word)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			break
		}
		res_map[floating_prefix] = append(res_map[floating_prefix], last_word)
		floating_prefix = strings.Join(
			append(strings.Split(floating_prefix, " "), last_word)[1:],
			" ",
		)
	}
	return res_map
}

func GenerateText(prefix string, data_map map[string][]string) {
	var randomNum int
	var last_word string
	var exists bool
	var amount_of_loops int = *internal.Gen_text_len - *internal.Prefix_len
	fmt.Print(prefix + " ")
	for i := 0; i < amount_of_loops; i++ {
		_, exists = data_map[prefix]
		if exists {
			randomNum = rand.Intn(100) % len(data_map[prefix])
			last_word = data_map[prefix][randomNum]
			fmt.Print(last_word)
			if amount_of_loops-i != 1 {
				fmt.Print(" ")
			}
			prefix = strings.Join(append(strings.Split(prefix, " "), last_word)[1:], " ")
		} else {
			break
		}
	}
	fmt.Println()
}
