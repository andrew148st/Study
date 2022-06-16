package main

// функция squareRange: принимает на вход два значения типа int: start и finish. возвращает слайс, с квадратами
//всех чисел, входящих в промежуток между start и finish (включительно с обеих сторон), т.е. при вводе (1, 3),
//возвращаем [1, 4, 9]

import "fmt"

func squareRange(start int, finish int) []int { // объявляем функцию, которая выполняет задание
	sliceSquare := []int{}             // объявляем слайс с итоговыми данными //
	for i := start; i <= finish; i++ { // запускаем цикл по входным данным
		square := i * i                           // вычисляем квадрат текущего числа в цикле
		sliceSquare = append(sliceSquare, square) // добавляем его в слайс с итоговыми данными
	}
	return sliceSquare //возвращаем итоговый слайс
}
func main() {
	start := 8                                     // присваиваем входные значания
	finish := 15                                   // присваиваем входные значания
	fmt.Printf("%d\n", squareRange(start, finish)) //выводим на печать результат работы функции
}
