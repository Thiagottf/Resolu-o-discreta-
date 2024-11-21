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
	fmt.Println("5 - Abrir log")
	fmt.Println("6 - Sair")

	fmt.Scanf("%d", &metodo)
	LogEscolha(metodo)
	switch metodo {
	case 1:
		fmt.Println("Indução Matemática selecionada")
		Indução()
	case 2:
		fmt.Println("Recursão selecionada")
		Recursao()
	case 3:
		fmt.Println("Teoria dos Números selecionada")
		TeoriaDosNumeros()
	case 4:
		fmt.Println("Aritmetica Modular selecionada")
		Aritmeticas()
	case 5:
		fmt.Println("Abrindo log...")
		time.Sleep(2 * time.Second)
		AbrirLog()
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
