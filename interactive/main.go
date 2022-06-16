package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func read(prompt string) string {
	fmt.Printf(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func main() {
	name := read("type your name: ")
	fmt.Printf("Hello, %s", name)
}
