package main

import (
	"fmt"
	"time"
)

func Menu() {
	fmt.Println("Bem-vindo ao Resolvedor Discreto")
	fmt.Println("Escolha o metodo que deseja utilizar")
	fmt.Println("1 - Indução Matemática")
	fmt.Println("2 - Recursão")
	fmt.Println("3 - Teoria dos Números")
	fmt.Println("4 - Aritmetica Modular")
	fmt.Println("6 - Sair")

	fmt.Scanf("%d", &metodo)
	switch metodo {
	case 1:
		Indução()
	case 2:
		Recursao()
	case 3:
		TeoriaDosNumeros()
	case 4:
		Aritmeticas()
	case 6:
		fmt.Println("SAINDO...")
		time.Sleep(2 * time.Second)
	default:
		fmt.Println("Opção inválida!")
		time.Sleep(1 * time.Second)
		fmt.Println("Por gentileza escolha uma opção válida!")
		Menu()
	}
}
