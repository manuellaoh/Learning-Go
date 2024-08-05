// create a calculator which accepts input in form of strings, converts it to an integer, and then sums it.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i, j string
	var num1, num2 float64

	fmt.Println("Hello there! please input a number : ")
	fmt.Scan(&i)
	fmt.Println("Hello there! please input a number : ")
	fmt.Scan(&j)
	//fmt.Println("your number is:", i + j)

	num1, err := strconv.ParseFloat(i, 64)
	if err != nil {
		fmt.Println("Error converting i to float:", err)
		return
	}

	num2, err2 := strconv.ParseFloat(j, 64)
	if err2 != nil {
		fmt.Println("Error converting j to float:", err2)
		return
	}

	fmt.Println("Your numbers are:", num1, num2)
	sum := num1 + num2
	fmt.Println("Sum:", sum)
}
