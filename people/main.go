package main

import "fmt"

type person struct { // объявляем тип: структура person
	name  string  // которое содержит 2 поля: строку name
	money float64 // и число money
}

func (p person) isRicher(p2 person) bool { // объявляем метод (функцию) с аргументами p и p2 типа структура person
	return p.money > p2.money // возвращает булевый результат сравнения p.money и p2.money
}
func (p person) prettyPrint() string { // метод с одним аргументом p типа person
	return fmt.Sprintf("My name is %s, I have a lot of money - %f $", p.name, p.money) // возвращает строку
}
func (p *person) addMoney(m float64) { // метод с аргументом p - указателем на person и аргументом m
	p.money += m // возвращает оригинал значения money плюс m
}

func main() {
	ivan := person{name: "Ivan", money: 200}                 // присваиваем значения полям структуры
	vasya := person{name: "Vasya"}                           // и делаем из них переменные vasya и ivan
	fmt.Printf("%s\n", ivan.prettyPrint())                   // выводим результат функции prettyPrint с переменной ivan
	fmt.Printf("%s\n", vasya.prettyPrint())                  //выводим результат функции prettyPrint с переменной vasya
	fmt.Printf("Is ivan richer? %v\n", ivan.isRicher(vasya)) // выводим результат функции isRicher

	vasya.addMoney(300)                                      // выполняем метод addMoney с переменными vasya и m
	fmt.Printf("%s\n", ivan.prettyPrint())                   // выводим результат функции prettyPrint с переменной ivan
	fmt.Printf("%s\n", vasya.prettyPrint())                  //выводим результат функции prettyPrint с переменной vasya
	fmt.Printf("Is ivan richer? %v\n", ivan.isRicher(vasya)) // выводим результат функции isRicher
}
