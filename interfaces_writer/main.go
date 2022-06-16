package main

import "fmt"

type Writer interface {
	Write(string)
}

type ConsoleWriter struct{}

func (p *ConsoleWriter) Write(what string) {
	fmt.Println(what)
}

type SukaConsoleWriter struct {
}

func (p *SukaConsoleWriter) Write(what string) {
	fmt.Printf("%s, Suka\n", what)
}

func main() {
	var w Writer = &SukaConsoleWriter{}

	w.Write("hello world")
	w.Write("poka mir")
	w.Write("ya molodets")
	w.Write("kruto napisal")
	w.Write("dosviduli")
}
