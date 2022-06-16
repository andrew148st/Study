package main

import (
	"fmt"
)

const three = 3
const (
	favoriteNumber = 148
	flatNumber     = 213
)

func main() {
	sum1 := add(2, 3)
	var sum2 int
	sum2 = add(4, 4)
	fmt.Printf("Hello, %d world! %d\n", sum1, sum2)
	sum2 = 148
	fmt.Printf("Hello, %d world! %d\n", sum1, sum2)
	var sum3 int
	sum3 = subtr(256, 128)
	fmt.Printf("Capacity of my flash card is %d GB\n", sum3)
	var sum4 int
	sum4 = multi(26, 2)
	fmt.Printf("I am %d years old\n", sum4)
	var sum5 int
	sum5 = div(5997, 3)
	fmt.Printf("My son was born at %d\n", sum5)
	{
		areEven := isEven(three) && isEven(favoriteNumber)
		fmt.Printf("%t\n", areEven)
		fmt.Printf("My son was born at %d\n", sum5)
	}
	if sum5 == 1999 {
		fmt.Printf("условие выполнено\n")
	} else if sum5 == 1998 {
		fmt.Printf("условие 2 выполнено\n")
	} else {
		fmt.Printf("ничего не выполнено\n")
	}
	switch sum5 {
	case 1999:
		fmt.Printf("условие switch выполнено\n")
	case 1998:
		fmt.Printf("условие 2 выполнено\n")
	default:
		fmt.Printf("nothing done\n")
	}
	for i := 1; i <= 10; i++ {
		if isEven(i) {
			fmt.Printf("число %d четное\n", i)
		} else {
			fmt.Printf("число %d нечетное\n", i)
		}
	}
}
func isEven(num int) bool {
	fmt.Printf("isEven called\n")
	return num%2 == 0
}
func add(num1 int, num2 int) int {
	return num1 + num2
}
func subtr(num1 int, num2 int) int {
	return num1 - num2
}
func multi(num1 int, num2 int) int {
	return num1 * num2
}
func div(num1 int, num2 int) int {
	return num1 / num2
}
