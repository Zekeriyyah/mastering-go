package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomDate(n int) {
	for n > 0 {
		randomize()
		n--
		time.Sleep(time.Millisecond) //Just playing with some delay in processing
	}
}

func randomize() {
	var era = "AD"
	year := rand.Intn(181) + 1900 //getting random number from 1900- 2080
	month := rand.Intn(12) + 1
	daysInMonth := 31
	switch month {
	case 2:
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			daysInMonth = 29
		}
		daysInMonth = 28
	case 4, 6, 9, 11:
		daysInMonth = 30
	}
	day := rand.Intn(daysInMonth) + 1
	fmt.Printf("%02d-%02d-%4d %v\n", day, month, year, era)
}
