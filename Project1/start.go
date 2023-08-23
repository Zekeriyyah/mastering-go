package main

// import fmt package
import (
	"fmt"
	"math/rand"
	"time"
)

func randNum() {
	num := rand.Intn(345000001) + 56000000 //To generate random number between 56e6 to 401e6
	fmt.Println("Random Number1 is", num)

	num = rand.Intn(345000001) + 56000000
	fmt.Println("Random Number2 is", num)
}

func countDown() {
	count := 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
	fmt.Println("Loop terminated")
}

func infLoop() { //To test how to randomly break from an infinite loop
	degrees := 0
	for {
		fmt.Println(degrees)
		degrees++
		if degrees >= 360 {
			degrees = 0
			if rand.Intn(2) == 0 {
				break
			}
		}
	}
}

// Implementing a countdown where every second there's 1 in 100 chance of failure
func countDown2() {
	count := 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		if rand.Intn(100) == 0 {
			break
		}
		count--
	}
	if count == 0 {
		fmt.Println("LiftOff!")
	} else {
		fmt.Println("Launch Failed!")
	}

}

func main() {

	//Working on big package
	var distance float64 = 99e18
	fmt.Println("Distance is ", distance)

	//How long does it take to get the mars
	const lightSpeed = 299792 //km/s
	var dist_ = 56000000      //km
	fmt.Println(dist_/lightSpeed, "Seconds.")

	dist_ = 401000000
	fmt.Println(dist_/lightSpeed, "Seconds.")

	//QuickCh2.3 ques: How many days would it take to reach mars
	var speed = 100800  //km/h
	const hrPerDay = 24 //h
	var dist = 96300000 //km
	fmt.Printf("It will take %v days.\n", (dist/speed)/hrPerDay)

	//QC Ques2.6: Asking computer to generate random number
	randNum()

	//Using for loop in counterdown timer
	//countDown()

	//QC QUES3.6: Printing angle within the bound of 0 to 360 degrees
	//infLoop()
	//countDown2()
	randomDate(100)
}
