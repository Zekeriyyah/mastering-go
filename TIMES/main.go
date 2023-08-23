package main

import (
	"fmt"
	"runtime"
	"time"
)

var week time.Duration

func main() {
	//This is about time pkg of type time.Type which
	//allow to measure and display time
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())

	t = time.Now().UTC()
	fmt.Println(t)
	//Calculating time
	week = 60 * 60 * 24 * 7 * 1e9
	weekFromNow := t.Add(2 * week)
	fmt.Println(weekFromNow)

	t = time.Now().Local()
	//Formatting time
	fmt.Println(t.Format(time.RFC822))

	
}
