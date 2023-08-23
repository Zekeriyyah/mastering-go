package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "Who are you to dictate that?"

	// //Prefixes and suffixes
	// fmt.Printf("Expect true: %v", strings.HasPrefix(str, "Who"))
	// fmt.Println()

	// fmt.Printf("Expect false: %v", strings.HasPrefix(str, "@"))
	// fmt.Println()

	// fmt.Printf("Expect true: %v", strings.HasSuffix(str, "?"))
	// fmt.Println()

	// fmt.Printf("Expect false: %v", strings.HasSuffix(str, "3"))
	// fmt.Println()

	var str1 string = "This-is-an-example-of-a-string"
	// fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str1, "Th")
	// fmt.Printf("%t\n", strings.HasPrefix(str1, "Th"))

	// //checking or presence of substring
	// fmt.Printf("%s is part of \"%s\", (T/F): ", "dictate", str)
	// fmt.Printf("%t\n", strings.Contains(str, "dictate"))

	// //Checking for index of firs instance of a substring
	// fmt.Printf("Index of \"%s\" in \"%s\": ", "?", str)
	// fmt.Printf("%d\n", strings.Index(str, "?"))

	// fmt.Println("\nstr: ", str)
	// fmt.Println("LastIndex(str, \"t\"): ", strings.LastIndex(str, "t"))

	// //Replacing substring n= -1 implies replacing all occurrences
	// str2 := strings.Replace(str, "dictate", "tell", 1)
	// fmt.Println("str2: ", str2)

	// //Counting occurrences of a substring
	// fmt.Printf("\"%s\" in \"%s\" are: ", "t", str)
	// fmt.Println(strings.Count(str, "t"))

	// //Repeating a string
	// fmt.Println("strings.Repeat(str, 3): ", strings.Repeat(str, 2))

	// //Changing the  case of string
	// fmt.Println("All lower case: ", strings.ToLower(str))
	// fmt.Println("All upper case: ", strings.ToUpper(str))

	// //Trimming a string
	// fmt.Println("Before strings.TrimSpace(str): ", str)
	// fmt.Println("After strings.TrimSpace(str): ", strings.TrimSpace(str))
	// fmt.Println()
	// fmt.Println("Before strings.Trim(str, \" ?\"): ", str)
	// fmt.Println("After strings.Trim(str, \" ?\"): ", strings.Trim(str, "?"))

	//Splitting and joining strings
	fmt.Println()
	fmt.Println("Splitted str base on whitespace: ", strings.Fields(str))
	fmt.Println("Before splitting str1: ", str1)
	fmt.Println("After splitting base on - : ", strings.Split(str1, "-"))
	fmt.Println()
	fmt.Println("Joining str1 with comma: ", strings.Join(strings.Split(str1, "-"), ", "))

	//Conversion from or to a string
	fmt.Println("Int: ", strconv.IntSize)
	fmt.Println("StringFormof20: ", strconv.Itoa(20)+" dogs have been sold.")
	fmt.Println("20.34tostring: ", strconv.FormatFloat(20.34, 'f', 2, 64)+" is the average mark.")
	num, err := strconv.Atoi("20")
	fmt.Printf("Int: %v  Error: %v\n", num, err)
	f, err := strconv.ParseFloat("43.34", 64)
	fmt.Printf("Float: %v  Error: %v\n", f*2.1, err)
}
