package main

import "fmt"

// Структуры, реализующие SliceMemo будут хранить в себе несколько значений (строковых). Как они это должны делать:
// Внутри структуры храним слайс. Аргумент к Store() (строковый) - это само значение, которое нужно сохранить.
// При этом Load() достает значение по определенному индексу, который принимает так же аргументом.
type SliceMemo interface {
	Store(string)
	Load(int) string
}

type InMemorySliceMemo struct {
	storage []string
}

func (p *InMemorySliceMemo) Store(value string) {
	p.storage = append(p.storage, value)
}

func (p *InMemorySliceMemo) Load(index int) string {
	return p.storage[index]
}

func main() {
	var m SliceMemo = &InMemorySliceMemo{}

	m.Store("stroka")
	m.Store("vtoraya")
	fmt.Printf("reading our slice: %s\n", m.Load(0))
}
