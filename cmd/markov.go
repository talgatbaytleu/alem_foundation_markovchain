package cmd

import (
	"fmt"
	"io"
	"strings"

	"markov/internal"
)

func Markov() {
	if *internal.Help_flag {
		fmt.Print(
			"Markov Chain text generator.\n\nUsage:\n  markovchain [-w ,<N>] [-p <S>] [-l <N>]\n  markovchain --help\n\nOptions:\n  --help\tShow this screen.\n  -w N\tNumber of maximum words\n  -p S\tStarting prefix\n  -l N\tPrefix lenth\n",
		)
	}
}

func GenerateMap(prefix_len int) map[string][]string {
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
