package main

import (
	"fmt"
	"strings"
)

func closure() func() int {
	a, b, c := 0, 1, 1
	return func() int {
		a, b, c = b, c, a+b+c
		return c - a - b
	}
}

func tribonacci(num int) {
	fib := closure()
	for i := 0; i < num; i++ {
		fmt.Print(fib(), " ")
	}
}

func variadic(b, c int, a ...interface{}) {
	for _, val := range a {
		switch val.(type) {
		case int:
			fmt.Printf("%d: Int type\n", val)
		case float64:
			fmt.Printf("%f: Float64 type\n", val)
		case string:
			fmt.Printf("%s: String type\n", val)
		default:
			fmt.Println("Value is of unidentified type!!")
		}
	}
}

// func printFactorial(n int) {
// 	for i := 0; i < n; i++ {
// 		fmt.Printf("%d! == %d\n", i, factorial(i))
// 	}
// }

//	func factorial(n big.Int) *big.Int {
//		val := n.Cmp(big.NewInt(0))
//		if val == 0 {
//			return big.NewInt(1)
//		} else {
//			return n * factorial(n-big.NewInt(1))
//		}
//	}
func printrec(n int) {
	if n == 0 {

	} else {
		fmt.Print(n, " ")
		printrec(n - 1)
	}
}

func TestDefer() { //checking the effect of defer keyword
	for i := 0; i < 10; i++ {
		defer fmt.Print(i, " ")
	}
}

func fibonacci(n int) (res int) {

	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func main() {

	fmt.Println(func(a, b int) int { return a + b }(7, 5))
	tribonacci(10)
	fmt.Println()

	variadic(45, 65, 74.3, "animal", 77.3, 65)

	printrec(10)
	fmt.Println()

	fmt.Println()
	//printFactorial(30) //To print factorial of first 30 integers

	TestDefer()
	fmt.Println()

	fmt.Println("fibonacci(10): ", fibonacci(10))

	/*
		Exercise:
		The Map function in the package strings is also a good example of the use of higher order functions,
		like strings.IndexFunc(). Look up its definition in the package documentation and make a little
		test program with a map function that replaces all non-ASCII characters from a string with a ? or
		a space. What do you have to do to delete these characters?
	*/

	modified := func(r rune) rune {
		if r > 255 {
			return '?'
		}
		return r
	}
	var str string
	str = "muhammed😁gmail.com"
	result := strings.Map(modified, str)
	fmt.Println(result)

	type celsius float64
	const degrees = 20
	var temperature celsius = degrees
	temperature += 10
	fmt.Println(temperature)

}
