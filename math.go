package main

import "fmt"

func main() {
	fmt.Println(Somar(10, 15))
}

func Somar(a int, b int) int {
	return a + b
}

func Diminuir(a int, b int) int {
	return a - b
}
