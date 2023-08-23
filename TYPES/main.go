package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	p := 0.1
	p += 0.2
	fmt.Println(p == 0.3)

	fmt.Printf("Type %T for %[1]v\n", p)

	//testing limit
	future := time.Unix(21212111122, 0)
	fmt.Println(future)
	//Experimenting with floating-point numbers
	var piggyBal float64
	balance(&piggyBal)

	//Experimenting big number handling techniques
	val := big.NewInt(299792)
	fmt.Printf("The type of %v is %[1]T\n", val)

	val1 := new(big.Int)
	val1.SetString("24000000000000000000", 10)
	fmt.Printf("The type of %v is %[1]T\n", val1)

	second := new(big.Int)
	second.Div(val1, val)
	fmt.Println("Time in seconds is ", second)

	secPerDay := big.NewInt(24 * 3600)

	days := new(big.Int)
	days.Div(second, secPerDay)
	fmt.Println("Which is", days, "days.")

	//Using const keyword to declare a very big integer
	const speed = 29979234

	const dist = 236000000000000000
	const timeCal = dist / speed / 3600 / 24 / 365.25
	fmt.Printf("Dist is %v light years\n", timeCal)

	//Experimenting with string type
	msg := "shalom"
	for _, char := range msg {
		fmt.Printf("%c\n", char)
	}

	c := 'g'
	fmt.Printf("c is %c\n", c-'a'+'A')

	message := "uv vagreangvbany fcnpr fgngvba" //Cipher

	for i := 0; i < len(message); i++ { //For loop to decipher
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c += 13
			if c > 'z' {
				c -= 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()
	fmt.Println("==============================")

	//Calling decipher function to decipher string of uppercase and lowercase combination
	quote := "L fdph, L vdz, L frqtxhuhg."
	decipher(quote)

	question := "¿Cómo estás?"
	decode(&question)

	//Calling ciphSpanish() to encode spanish in ROT13
	s := "Hola Estación Espacial Internacional"
	ciphSpanish(s)

}

func balance(bal *float64) {
	nickel := 0.05
	dimes := 0.10
	quarters := 0.25

	for *bal-20.00 < 0 {
		switch gift := rand.Intn(3); gift {
		case 0:
			*bal += nickel
		case 1:
			*bal += dimes
		case 2:
			*bal += quarters
		}
		fmt.Printf("The Piggy Bank Balance: %.2f\n", *bal)
	}
}
