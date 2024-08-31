package cmd

import (
	"fmt"

	"markov/internal"
)

func Markov() {
	if *internal.Help_flag {
		fmt.Print(
			"Markov Chain text generator.\n\nUsage:\n  markovchain [-w ,<N>] [-p <S>] [-l <N>]\n  markovchain --help\n\nOptions:\n  --help\tShow this screen.\n  -w N\tNumber of maximum words\n  -p S\tStarting prefix\n  -l N\tPrefix lenth\n",
		)
	}
}
