package main

import "fmt"

type Person interface { // Объявляем интерфейс под названием Person. Интерфейс - это некий договор, что структуры, которые хотят под него подходить
	// будут реализовывать все методы интерфейса.
	Name() string        // Описание метода Name, который должен будет вернуть строку
	Occupation() string  // Описание метода Occupation, который должен будет вернуть строку
	PrettyPrint() string // Описание метода PrettyPrint, который должен будет вернуть строку
}

type Programmer struct { // Наш первый Person - Programmer. Он станет Person благодаря тому, что реализует все три метода.
	// Нам неважно какие поля, есть в Programmer, чтобы реализовать Person - нужны только методы
	name     string
	language string
}

func (p *Programmer) Name() string { // метод Name, точно такой же, как в спецификации Person
	return p.name // возвращаем поле внутри структуры
}

func (p *Programmer) Occupation() string { // метод Occupation, точно такой же, как в спецификации Person
	return "Programmer" // возвращаем литерал
}

func (p *Programmer) PrettyPrint() string { // метод PrettyPrint, точно такой же, как в спецификации Person
	return fmt.Sprintf("Hello, my name is %s, I work as a coder using the %s language", p.name, p.language) // возвращает форматированную строку
}

func (p *Programmer) Language() string { // метод, которого нет в Person, но это не важно для его имплементации,
	// потому что кроме тех, методов, которые есть в интерфейсе может быть любое количество дополнительных методов
	return p.language
}

// ===================================================================

type Bystander struct{} // структура-затычка, чтобы продемонстрировать минимум для реализации интерфейса

func (b *Bystander) Occupation() string { // метод Name, точно такой же, как в спецификации Person
	return "nothing" // возвращаем затычку
}

func (b *Bystander) Name() string { // метод Occupation, точно такой же, как в спецификации Person
	return "someone" // возвращаем затычку
}

func (b *Bystander) PrettyPrint() string { // метод PrettyPrint, точно такой же, как в спецификации Person
	return "i'm just a simple nobody" // возвращаем затычку
}

func main() {
	var p Person = &Programmer{name: "Vasya", language: "c++"} // явно рассматриваем инстанс структуры Programmer как Person.
	var p2 Person = &Bystander{}                               // так же рассматриваем инстанс Bystander
	fmt.Printf("%s\n", p.PrettyPrint())                        // выводим p
	fmt.Printf("%s\n", p2.PrettyPrint())                       // выводим p2

	p = &Bystander{}                    // показываем, что несмотря на то, что в p раньше был Programmer, теперь может быть Bystander, потому что они оба Person
	fmt.Printf("%s\n", p.PrettyPrint()) // показываем, что там теперь действительно Bystander
}
