package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main()  {
	fmt.Println("Hello world")
	fmt.Println("We're learning Go!")
	fmt.Println("the time is", time.Now())
	favorite()
	fmt.Println(add(2, 2))
}

func favorite() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

func add(x int, y int) int {return x + y }

