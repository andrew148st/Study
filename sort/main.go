package main

//											Сортировка "Пузырьковая"
import "fmt"

func main() {
	inputSlice := []int{1, 3, 645, 124, 2, 2, 3, 4, 12} // создаем слайс, который будем сортировать
	outputSlice := bubbleSort(inputSlice)               // получаем в результате вызова функции сортировки новый слайс
	fmt.Printf("%v\n", outputSlice)                     // выводим на экран
}

func bubbleSort(slice []int) []int { // функция bubbleSort принимает в кач-ве аргумента слайс int, возвращает отсортированный новый слайс int
	for range slice { // проходимся во внешнем цикле ровно столько раз, сколько элементов, потому что один проход по всему массиву позволяет "всплыть" только одному элемента
		for j := range slice { // проходимся по каждой паре элементов в массиве
			if j == len(slice)-1 { // если элемент последний - у него нет пары, пропускаем его
				continue
			}
			if slice[j] > slice[j+1] { // если элемент слева больше элемента справа -> нужно поменять их местами
				temp := slice[j+1]    // сохраняем во временную перемменную значение правого элемента, потому что мы будем его затирать в самом массиве
				slice[j+1] = slice[j] // затираем правый элемент левым
				slice[j] = temp       // затираем левый элемент копией правого
			}
		}
	}
	return slice // возвращаем измененный слайс
}