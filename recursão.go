package main

import "fmt"

// Função recursiva para Fibonacci
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// Função recursiva para Fatorial
func Fatorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Fatorial(n-1)
}

func Recursao() {
	fmt.Println("Você escolheu o método de Recursão")
	fmt.Println("Escolha a função que deseja resolver:")
	fmt.Println("1 - Fibonacci")
	fmt.Println("2 - Fatorial")
	var escolha int
	fmt.Scanf("%d", &escolha)

	fmt.Print("Digite o valor de n: ")
	fmt.Scanf("%d", &n)

	switch escolha {
	case 1:
		fmt.Printf("O %d-ésimo número de Fibonacci é: %d\n", n, Fibonacci(n))
	case 2:
		fmt.Printf("O fatorial de %d é: %d\n", n, Fatorial(n))
	default:
		fmt.Println("Opção inválida")
	}
}
