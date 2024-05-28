package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	//This is about time pkg of type time.Type which
	//allow to measure and display time
	t1 := time.Now()
	fmt.Println(t1)
	fmt.Printf("%02d.%02d.%4d\n", t1.Day(), t1.Month(), t1.Year())

	t2 := time.Now().UTC()
	fmt.Println("t2", t2)
	//Calculating time
	week = 60 * 60 * 24 * 7 * 1e9
	weekFromNow := t2.Add(2 * week)
	fmt.Println(weekFromNow)

	t3 := time.Now().Local()
	//Formatting time
	fmt.Println(t3.Format(time.RFC822))

	fmt.Println("t2time:", t2.Clock())

}
