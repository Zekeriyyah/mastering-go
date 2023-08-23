package main

import (
	"fmt"
	"math"
)

var a string

func main() {
	// goos := os.Getenv("GOOS")
	// fmt.Println("The operating system is: ", goos)

	// path := os.Getenv("PATH")
	// fmt.Println("The path is: ", path)

	a := "G"
	fmt.Print(a)
	f1()

	fmt.Printf("The value of a = G is: %t\n", a == "G")

	var x int16 = 2
	var y int32 = 5
	y = int32(x)
	fmt.Println(y)
	z := float64(3474834.277)
	fmt.Printf("%2.4g is the code.\n", z)

	m := -30
	conv_value, err := Uint8FromInt(m) //converting from int to unsigned integer
	fmt.Println(conv_value, "--Error: ", err)

	n := 203765678.4
	fmt.Printf("int32 form of %f is %d.\n", n, IntFromFloat64(n))

	type Rope string
	var q Rope = "Thick"
	fmt.Println(q)

	s := "come home..."
	value, errValue := hasPrefix(s, "Come")
	fmt.Println(value, "  ", errValue)
	fmt.Println()
	fmt.Println("\n================STRING=============")
	stringProperties()

	fmt.Println("\n============================TIME&DATE==========")

	time_datePract() //run all my practice on time and date package in time_datepkg.go file

}

func f1() {
	a := "O"
	fmt.Print(a)
	f2()
}
func f2() {
	fmt.Println(a)
}

func Uint8FromInt(n int) (val uint8, err error) { //Convert int to uint8
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the range of uint8.", n)
}

func IntFromFloat64(n float64) int {
	if n >= math.MinInt && n <= math.MaxInt {
		whole, fraction := math.Modf(n)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("ERROR:-%g is out range of type int32.", n))
}
