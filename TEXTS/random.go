package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randNum() {
	for i := 0; i < 10; i++ {
		r := rand.Int()
		fmt.Println(r)
	}

	for i := 0; i < 10; i++ {
		r := rand.Intn(8)
		fmt.Println(r)
	}

	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%3.2f\n", 1000*rand.Float64())

	}
}
