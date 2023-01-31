package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// fmt.Println("HELLO WORLD")
	rand.Seed(time.Now().UnixNano())
	m := rand.Intn(100000)
	fmt.Println(m)
}
