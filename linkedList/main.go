package main

import "fmt"

type Node struct { // Структура, обозначающая один квадрат на рисунке (включая его стрелочку)
	data int   // цифра внутри квадрата (значение узла)
	next *Node // стрелка на следующий квадрат в списке (указатель на следующий узел)
}

type LinkedList struct { // структура, которая сама по себе хранит "весь" список
	length int   // длина списка, обновляем ее сами на каждый Pushback(), длину хранить необязательно, она здесь для удобства
	head   *Node // "голова" списка, т.е. первый его элемент. хранить его обязательно, чтобы иметь доступ к узлам списка
	tail   *Node // "хвост" списка, т.е. его последний элемент. хранить его необязательно, он здесь для того, чтобы удобнее было добавлять новые элементы в конец
}

func (l *LinkedList) Pushback(d int) { // метод добавления нового узла в конец списка
	newNode := Node{data: d} // создаем новый узел с нужными данными (квадрат на рисунке)
	if l.head == nil {       // если у нас еще нет "головы", то он был пустым
		// а если он был пустым, у нас в результате добавления будет всего один квадрат, и тогда он будет по определению и "головой" и "хвостом"
		l.head = &newNode // говорим, что "голова" теперь - новый квадрат
		l.tail = &newNode // говорим, что "хвост" теперь - новый квадрат
	} else { // иначе у нас были элементы в списке (т.е. он не был пустым)
		// тогда добавляем новый элемент в конец списка
		l.tail.next = &newNode // рисуем стрелочку от старого хвоста к новому квадрату
		l.tail = &newNode      // назначаем новым хвостом все тот же квадрат
	}
	l.length++ // увеличиваем длину вне зависимости от того, куда добавили новый элемент
}

func (l *LinkedList) At(index int) int { // функция "случайного доступа", которая позволяет взять любой элемент по его индексу (как в массиве)
	// для связного списка взятие по индексу намного менее эффективно, чем для массива (или слайса), но кому-то может пригодиться
	current := l.head             // начинаем с головы
	for i := 1; i <= index; i++ { // итерируемся в цикле на один раз меньше, чем индекс (потому что начали уже с единицы)
		current = current.next // в каждой итерации переходим по стрелочке на один элемент вперед
	}

	return current.data // после цикла current будет указывать на нужный нам элемент, возвращаем цифру внутри этого элемента
}

func (l *LinkedList) Print() {
	for current := l.head; current != nil; current = current.next {
		fmt.Printf("Current element: %d\n", current.data)
	}
}

func main() {
	list := &LinkedList{}
	list.Pushback(15)
	list.Pushback(25)
	list.Pushback(30)
	list.Pushback(40)
	list.Pushback(9)
	list.Print()

	fmt.Printf("Index 2: %d", list.At(2))
}
