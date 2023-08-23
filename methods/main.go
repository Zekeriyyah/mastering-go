package main

import "fmt"

type celsius float64
type farenheit float64
type kelvin float64

func main() {
	var c celsius = 127
	var f farenheit = 45
	var k kelvin = 345

	//var f farenheit = 20
	// warmUp := 10.0
	// c += celsius(warmUp)
	// fmt.Println(c)

	// fmt.Print(c, " o C is ", CelsiusToKelvin(c), " o K\n")

	//Converting between temperatures using methods

	fmt.Print(c, " oC is ", c.farenheit(), " oF\n")
	fmt.Print(c, " oC is ", c.kelvin(), " K\n")
	fmt.Println()

	fmt.Print(f, " oF is ", f.celsius(), " oC\n")
	fmt.Print(f, "oF is ", f.kelvin(), " K\n")
	fmt.Println()

	fmt.Print(k, " K is ", k.celsius(), " oC\n")
	fmt.Print(k, " K is ", k.farenheit(), " oF\n")

}

func CelsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}
