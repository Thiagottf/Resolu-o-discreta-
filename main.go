package main

import "fmt"

var metodo int
var n int

func main() {
	Menu()
	fmt.Scanf("%d", &metodo)
	switch metodo {
	case 1:
		Indução()
	case 2:
		Recursao()
	case 3:
		TeoriaDosNumeros()
	default:
		fmt.Println("SAINDO...")
	}
}
