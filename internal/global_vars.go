package internal

import "flag"

var (
	Default_prefix string

	Prefix_len   = flag.Int("l", 2, "-l [int value] to change prefix length (2 words by default)")
	Gen_text_len = flag.Int(
		"w",
		100,
		"-w [int value] to change generated text length (100 words by default)",
	)
	Prefix = flag.String(
		"p",
		Default_prefix,
		"-p [string value] to use custom prefix. Must be contained by dictionary you've provided.(First 2 words from dictionary by default)",
	)
)
