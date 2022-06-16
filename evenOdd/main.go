package main

import "fmt"

func evenOdd(n int) ([]int, []int) { // функция evenOdd, берет на вход одно число, возвращает два слайса, один с четными цифрами, другой с нечетными
	sliceNO := []int{} // объявляем слайс нечетных цифр
	sliceNE := []int{} // объявляем слайс четных цифр
	for n > 0 {        // цикл будет продолжаться пока число не станет равно (или меньше) нуля
		digit := n % 10   // вычленяем последнюю цифру
		if digit%2 == 0 { // если делится без остатка на два
			sliceNE = append(sliceNE, digit) // то добавляем в слайс NE
		} else { // иначе
			sliceNO = append(sliceNO, digit) // добавляем в слайс NO
		}
		n = n / 10 // делим на 10, чтобы перейти к следующей цифре
	}

	return sliceNE, sliceNO // возвращаем два слайса
}

func main() {
	n := 358462
	ne, no := evenOdd(n)
	fmt.Printf("%v %v", ne, no)
}
