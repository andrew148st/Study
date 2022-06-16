package fizzbuzz

// 									fizzbuzz0
import "fmt"

func isDivThree(n int) bool {
	return n%3 == 0
}

func isDivFive(n int) bool {
	return n%5 == 0
}

func isDivThreeFive(n int) bool {
	return isDivThree(n) && isDivFive(n)
}

// написать функцию, которая вернет fizz для чисел, делящихся без остатка на 3, buzz, для делящихся без остатка на 5,
// fizzbuzz для чисел, которые делятся без остатка и на 5 и на 3
// и само число, превращенное в строку в любом другом случае
func fizzBuzz(n int) string {
	if isDivThreeFive(n) {
		return "fizzbuzz"
	} else if isDivThree(n) {
		return "fizz"
	} else if isDivFive(n) {
		return "buzz"
	} else {
		return fmt.Sprintf("%d", n)
	}
}

func main() {
}
