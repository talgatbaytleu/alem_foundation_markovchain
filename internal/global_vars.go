package internal

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	Help_flag    = flag.Bool("help", false, "--help to see Usage Information")
	Prefix_len   = flag.Int("l", 2, "-l [int value] to change prefix length (2 words by default)")
	Gen_text_len = flag.Int(
		"w",
		100,
		"-w [int value] to change generated text length (100 words by default)",
	)
	Custom_prefix = flag.String(
		"p",
		"",
		"-p [string value] to use custom prefix. Must be contained by dictionary you've provided.(First 2 words from dictionary by default)",
	)
	Parsed         = IfFlagsParsed()
	Default_prefix = InitDefaultPrefix(*Prefix_len)
	Prefix         string
)

func InitDefaultPrefix(prefix_len int) string {
	var res string
	var current_word string
	var err error
	for i := 0; i < prefix_len; i++ {
		_, err = fmt.Scan(&current_word)
		if err == io.EOF {
			fmt.Fprintf(os.Stderr, "The length of provided text is less than prefix length\n")
			os.Exit(1)
		}
		res += current_word + " "
	}
	res = res[:len(res)-1]
	return res
}

func IfFlagsParsed() bool {
	flag.Parse()
	return true
}
