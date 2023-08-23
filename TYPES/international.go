package main

import "fmt"

func ciphSpanish(s string) {
	for _, c := range s {
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			c -= 13
			if c < 'a' || c < 'A' {
				c += 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()
}
