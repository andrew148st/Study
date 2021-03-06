package main

import "fmt"

//										Тема "Массив и матрица"
func main() {
	arr := [5]int{5, 2, 5, 4, 5} // массив на 5 ячеек, объявлен литералом
	slice := arr[0:2]            // объявляем слайс на ячейки 0 и 1
	slice = append(slice, 2)     // добавляем в слайс ячейку 2 со значением 2 (литералом)
	fmt.Printf("%v\n", arr)      // печатаем массив (он изменился, ячейка 2 тепарь не 5, а 2, как в слайсе)
	fmt.Printf("%v\n", slice)    // печатаем слайс
	fmt.Printf("%v\n", arr[0])   // печатаем ячейку 0 из массива
	fmt.Printf("%v\n", arr[1])   // печатаем ячейку 1 из массива
	fmt.Printf("%v\n", slice[0]) // печатаем ячейку 0 из слайса
	fmt.Printf("%v\n", slice[1]) // печатаем ячейку 1 из слайса

	for i, num := range arr { // выводим на печать массив по форме:
		fmt.Printf("arr[%d] = %d\n", i, num) // "arr [номер ячейки] = [содержимое ячейки]"
	}

	matrix := [10][10]int{}   // объявляем матрицу
	for i := 0; i <= 9; i++ { // цикл по всем рядам
		row := [10]int{}          // создать массив из 10 ячеек, заполненных нулями
		for j := 0; j <= 9; j++ { // цикл по одному ряду
			row[j] = j + 1 // заполняем конкретную ячейку
		}
		matrix[i] = row // присвоить ряд
	}

	fmt.Printf("%v\n", matrix)
}

// попробовать разные математические выражения, например j * 2 и разобраться, почему так
