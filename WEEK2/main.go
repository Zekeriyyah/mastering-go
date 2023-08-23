package main

import "fmt"

func main() {
	response := testIf()
	fmt.Println(response)

	goTypes()
}

func testIf() string {
	a := true
	if !a {
		return "COOL"
	}
	return "Not good"

}

func goTypes() {
	x := []int{23, 54, 67, 90, 84}
	fmt.Println(x)
	fmt.Println(x[2:4])

	y := make(map[string]int)
	y["key"] = 2
	fmt.Println(y)

	if val, ok := y["key"]; ok {
		fmt.Println(val, ok)
	} else {
		fmt.Printf("Not found.\n")
	}
}
