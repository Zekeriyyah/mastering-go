package main

import (
	"fmt"
	"time"
)

func time_datePract() {
	fmt.Println(time.Now())

	t := time.Now()
	fmt.Printf("Today's Date: %02d-%02d-%04d\n", t.Day(), t.Month(), t.Year())

	t = time.Now().UTC()
	fmt.Println(t)
}
