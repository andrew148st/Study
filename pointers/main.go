package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	a := 5       // объявляем переменную a
	print_(&a)   // передаем адрес на переменную a в функцию print_(), адрес берется оператором &
	fmt.Print(a) // показываем, что a изменилась

	for _, value := range arr {
		print_(&value)
	}

	fmt.Print(arr)
}

func print_(address1 *int) {
	fmt.Printf("%d\n", *address1) // выводим значение по адресу address1
	*address1 = 4                 // изменяем значение по адресу address1
}
