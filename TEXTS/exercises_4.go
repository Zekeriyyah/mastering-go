package main

import (
	"fmt"
	"unicode/utf8"
)

var a = "G"

type ByteSize float64

const (
	_           = iota //Ignore the first iota value
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func exercise() {
	n()
	m()
	n()
	fmt.Println()

	fmt.Printf("%v is KB\n", KB)
	for i := 4; i >= 0; i-- { //It print +Inf division of float64 by zero is encountered
		div := KB / ByteSize(i)
		fmt.Printf("%v/%v is: %v", KB, i, div)
		fmt.Println()
	}

}

func n() { fmt.Print(a) }
func m() {
	a := "O"
	fmt.Print(a)
}

func countChar(s string) {
	bs := []byte(s)
	r := utf8.RuneCount(bs)
	fmt.Printf("The lenth of the string is %d", r)
	fmt.Println()
}
