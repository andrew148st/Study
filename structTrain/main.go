package main

import (
	"fmt"
	"math"
)

type point struct { //объявляем структуру point
	x int //с полями x и y
	y int
}

func (p point) distanceTo(p2 point) int { //объявляем метод distanceTo для вычисления расстояния с переменными типа point
	return int( // возвращает результат формулы, но интовый
		math.Sqrt( // мат. формула для вычисления, данные из переменных p и p2 структуры point
			math.Pow(float64(p2.x-p.x), 2) + math.Pow(float64(p2.y-p.y), 2),
		),
	)
}
func main() {
	var deliverTo point //объявляем переменную типа point
	deliverTo.x = 20    //с конечными координатами
	deliverTo.y = 30
	fmt.Printf("Delivering to: x = %d y =%d\n", deliverTo.x, deliverTo.y)        //выводим их на печать
	deliverFrom := point{y: 50, x: 70}                                           //объявляем переменную типа point с начальными координатами
	fmt.Printf("Delivering from: x = %d y = %d\n", deliverFrom.x, deliverFrom.y) //выводим их на печать
	distance := deliverFrom.distanceTo(deliverTo)                                //переменной присваиваем результат метода
	fmt.Printf("Distance: %d\n", distance)                                       // выводим её
}
