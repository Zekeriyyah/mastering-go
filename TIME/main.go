package main

import (
	"fmt"
	"time"
)

func main() {
	//This is about time pkg of type time.Type which
	//allow to measure and display time
	t := time.Now().UTC()
	fmt.Println(t)
}
