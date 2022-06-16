package main

//										fizzbuzz1
import "fmt"

func isDivThree(num int) bool {
	return num%3 == 0
}
func isDivFive(num int) bool {
	return num%5 == 0
}
func isDivThreeFive(num int) bool {
	return isDivThree(num) && isDivFive(num)
}

func main() {
	for i := 1; i <= 100; i++ {
		if isDivThreeFive(i) {
			fmt.Printf("fizzbuzz\n")
		} else if isDivThree(i) {
			fmt.Printf("fizz\n")
		} else if isDivFive(i) {
			fmt.Printf("buzz\n")
		} else {
			fmt.Printf("%d\n", i)
		}

	}
}
