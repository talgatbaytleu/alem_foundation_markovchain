package internal

import (
	"flag"
	"fmt"
)

var (
	Help_flag    = flag.Bool("help", false, "--help to see Usage Information")
	Prefix_len   = flag.Int("l", 2, "-l [int value] to change prefix length (2 words by default)")
	Gen_text_len = flag.Int(
		"w",
		100,
		"-w [int value] to change generated text length (100 words by default)",
	)
	Default_prefix string
	Prefix         = flag.String(
		"p",
		Default_prefix,
		"-p [string value] to use custom prefix. Must be contained by dictionary you've provided.(First 2 words from dictionary by default)",
	)
)

func InitDefaultPrefix(prefix_len int) string {
	var res string
	var current_word string
	for i := 0; i < prefix_len; i++ {
		fmt.Scan(&current_word)
		res += current_word
	}
	return res
}
