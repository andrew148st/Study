package main

import "fmt"

func linearSearch(inputSlice []int, value int) int { //объявляем функцию поиска с аргументами: исходный слайс и индекс искомого
	for i := range inputSlice { // запускаем цикл поиска по слайсу
		if inputSlice[i] == value { // если совпадает с искомым
			return i // возвращаем индекс искомого
		}
	}
	return -1 // если дошли сюда, возвращаем -1
}
func linearCount(inputSlice []int, value int) int { //
	count := 0                  // создаём переменную под "счётчик"
	for i := range inputSlice { // создаём цикл
		if inputSlice[i] == value {
			count++ //добавлем единицу к счётчику
		}
	}
	return count // возвращаем значение счётчика
}

func main() {
	input1 := []int{25, 17, 54, 85, 17, 695, 8, 17, 4875, 17} // заполняем исходный слайс литералом
	fmt.Printf("found: %d\n", linearSearch(input1, 17))       // выводим на печать результат работы фунции
	fmt.Printf("found: %d", linearCount(input1, 17))          // выводим на печать результат работы фунции
}
