package main

import "fmt"

//							Задание "Алиса"
// функция, которая в зависимости от параметра вернет строку по следующему условию:
// если в параметре name нам передали строку "Alice", мы возвращаем "hello, Alice!"
// иначе возвращаем "hello, stranger!"
// функцию проверяем вызовами в main
func greetAlice(name string) string { // объявляем функцию
	if name == "Alice" { // проверяем на равенство нужной строке
		return "hello, Alice!"
	} else { // пишем условия
		return "hello, stranger!"
	}
}

func main() { // выполняемая часть
	fmt.Printf("result: %s\n", greetAlice("Alice"))   // выполнение разных условий
	fmt.Printf("result: %s\n", greetAlice("Andrew1")) // и соответствующий вывод
}

// add new comment for Git
// one more
