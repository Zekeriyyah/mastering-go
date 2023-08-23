package main

import "fmt"

func decipher(a string) {
	for _, char := range a {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			char -= 3
			if char < 'A' || (char > 'Z' && char < 'a') {
				char += 26
			}
		}

		fmt.Printf("%c", char)
	}
	fmt.Println()
}
