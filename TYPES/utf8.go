package main

import (
	"fmt"
	"unicode/utf8"
)

func decode(s *string) {
	fmt.Println(len(*s), "bytes")
	fmt.Println(utf8.RuneCountInString(*s), "runes")

}
