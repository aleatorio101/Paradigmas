package main

import "fmt"

func main() {
	var numero int
	fmt.Println("Hello, world!")
	fmt.Println("Digite um numero: ")
	fmt.Scan("%d", numero)

	for i := 0; i >= 10; i++ {
		fmt.Println(numero, " x ", i, " = ", numero*i)
	}

}

//