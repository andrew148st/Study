package main

//					тема: "Структура и метод"; задача "вычисление расстояния по координатам"
import (
	"fmt"
	"math"
)

// объявляем, что у нас есть новый тип - структура point
type point struct { // тип "структура" с названием point

	x int // поля структуры
	y int // (внутренние переменные)
}

// объявляем метод
func (p point) distanceTo(p2 point) int { // функция с названием distanceTo с переменной p типа point и переменной p2 типа point
	return int( // возвращает тип int
		math.Sqrt(
			math.Pow(float64(p2.x-p.x), 2) + math.Pow(float64(p2.y-p.y), 2), // матем. формула вычисления расстояния
		),
	)
}

// объявляем обычную функцию, принимающую две точки (для сравнения с методом)
//func distanceTo(p point, p2 point) int {
//	return p.x
//}

func main() { //
	// 				обозначаем точки
	var deliverTo point // объявляем переменную с типом структуры point
	deliverTo.x = 20    // присваиваем переменным внутри конкретно этого point значения
	deliverTo.y = 30
	fmt.Printf("Delivering to: x = %d y = %d\n", deliverTo.x /* берем значение переменной внутри структуры */, deliverTo.y)

	// создание структуры через литералы, код эквивалентен по поведению строкам 25-27
	// deliverFrom := point{10, 10}  <- то же самое, что и следующая строка
	deliverFrom := point{y: 50, x: 70} // но такой метод предпочтительнее по стилю
	fmt.Printf("Delivering from: x = %d y = %d\n", deliverFrom.x /* берем значение переменной внутри структуры */, deliverFrom.y)

	// 								вычисляем расстояние
	// сравниваем метод и обычную функцию
	// distance := distanceTo(deliverFrom, deliverTo) // обычная функция
	distance := deliverFrom.distanceTo(deliverTo) // метод
	fmt.Printf("Distance: %d\n", distance)
}
