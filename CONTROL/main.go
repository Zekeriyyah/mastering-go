package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	orig := "ABC"
	orig = "44"
	anInt, err := strconv.Atoi(orig)

	if err != nil {
		fmt.Printf("An error occurred during string to int conversion: %v", err)
		os.Exit(1)
	}
	fmt.Printf("The integer form of %s is: %d\n", orig, anInt)

	// f, err := os.Open("my_cards")
	// if err != nil {
	// 	fmt.Printf("Error in Opening file: %v\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Println("Content of the file:\n", f)
	//switch-case control structure

	k := 6
	switch k {
	case 4:
		fmt.Println("Was <= 4")
		fallthrough
	case 5:
		fmt.Println("Was <= 5")
		fallthrough
	case 6:
		fmt.Println("Was <= 6")
		fallthrough
	case 7:
		fmt.Println("Was <= 7")
		fallthrough
	case 8:
		fmt.Println("Was <= 8")
		fallthrough
	default:
		fmt.Println("It's nothing")
	}

	exercises()

}

func exercises() {
	//Exercise 5.4
	for i := 0; i < 15; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	//Exercise 5.4-2
	i := 0
LABEL:
	if i < 15 {
		fmt.Printf("%d ", i)
		i += 1
		goto LABEL
	}
	fmt.Println()

	//Exercies 5.5
	for i := 0; i < 25; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}
	//Using string concatenation
	s := "G"
	for len(s) < 25 {
		fmt.Println(s)
		s += "G"
	}

	//To show bitwise complement of integers from 0 to 10
	for i := 0; i <= 10; i++ {
		j := ^i
		fmt.Printf("%d == %b", i, j)
		fmt.Println()
	}
	//Checking for multiples while printing index
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("fizzBuzz ")
		case i%3 == 0:
			fmt.Print("fizz ")
		case i%5 == 0:
			fmt.Print("Buzz ")
		default:
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	//Printing a rectangle of width 20 and height 10
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			if i == 0 || i == 9 {
				fmt.Print("=")
			} else if j == 0 || j == 19 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
	//Infinite loop
	// for {
	// 	t := time.Now()
	// 	if t.Minute() == 56 {
	// 		break
	// 	}
	// 	fmt.Print("000000000   ")
	// }
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}
	// for i := 0; ; i++ {
	// 	fmt.Println("Value of i is now: ", i)
	// }
	s = ""
	for s != "aaaaaa" {
		fmt.Println("Value of s: ", s)
		s += "a"
	}
	i = 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println("Value of i is: ", i)
		i++
	}
	fmt.Println("Statement after the for loop")
}
