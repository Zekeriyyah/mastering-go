package main

import (
	"fmt"
	"strings"
)

func hasPrefix(s, pre string) (a bool, err error) {
	for i := 0; i < len(pre); i++ {
		if s[i] == pre[i] {
			i++
		} else {
			return false, fmt.Errorf("--Error: The string \"%s\" has no prefix \"%s\".", s, pre)
		}
	}
	return true, nil
}

func stringProperties() {
	s := "     Don't ever belittle yourself..."
	fmt.Println(s)
	fmt.Println("Index of ever", strings.Index(s, "ever"))
	fmt.Println("Index of el", strings.Index(s, "el"))
	fmt.Println("Index of last instance of el", strings.LastIndex(s, "el"))
	fmt.Println("Index of que", strings.Index(s, "que"))
	fmt.Println("Does the writeup contains ttle? ", strings.Contains(s, "ttle"))

	fmt.Println("Replaced yourself with anyone: ", strings.Replace(s, "yourself", "anyone", 1))
	fmt.Println("a's: ", strings.Count(s, "a"))

	fmt.Println(s)
	fmt.Println(strings.TrimSpace(s))
	fmt.Println(strings.Trim(s, "Don"))

	f := strings.Fields(s) //split string into string slice using whitespace
	for _, val := range f {
		fmt.Println(val)
	}

	s1 := "Abdul-Razaq, Muhammed, Ahmad, Raphael, Edimeh" //split string into slice using seperator as second argument
	f1 := strings.Split(s1, ",")

	for _, val := range f1 {
		fmt.Println(val)
	}

	s2 := strings.Join(f1, "---") //Join strings in slice to form a string seperated with seperator provided as the second argument
	fmt.Println(s2)
}
