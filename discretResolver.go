package main

import "fmt"

//declaração das variaveis
var n int
var p string
var metodo int

func main() {
	menu()
	fmt.Scanf("%d", &metodo)
	switch metodo {
	case 1:
		indução()
	case 2:
		recursao()
	default:
	}
}

//Declaração das funções
func menu() {
	fmt.Println("Bem-vindo ao Numeric Resolver")
	fmt.Println("Escolha o metodo que deseja utilizar")
	fmt.Println("1 - Indução Matemática")
	fmt.Println("2 - Recursão")
}

func indução() {
	fmt.Println("Você escolheu o método de Indução Matemática")
	fmt.Print("Digite o valor de n inicial (caso base): ")
	fmt.Scanf("%d", &n)

	if proposicao(n) {
		fmt.Printf("O caso base P(%d) é verdadeiro.\n", n)
	} else {
		fmt.Printf("O caso base P(%d) é falso.\n", n)
		return
	}

	// Verifica a etapa de indução para os próximos valores
	limite := 10 // Definimos um limite máximo para verificação
	inducaoSucesso := true

	for i := n; i < n+limite; i++ {
		if !proposicao(i + 1) {
			fmt.Printf("A indução falhou para P(%d).\n", i+1)
			inducaoSucesso = false
			break
		}
	}

	if inducaoSucesso {
		fmt.Println("A indução foi bem-sucedida para o intervalo especificado.")
	} else {
		fmt.Println("A indução falhou para o intervalo especificado.")
	}
}

// Função recursiva para Fibonacci
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Função recursiva para Fatorial
func fatorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fatorial(n-1)
}

func recursao() {
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
		fmt.Printf("O %d-ésimo número de Fibonacci é: %d\n", n, fibonacci(n))
	case 2:
		fmt.Printf("O fatorial de %d é: %d\n", n, fatorial(n))
	default:
		fmt.Println("Opção inválida")
	}
}

func proposicao(n int) bool {
	// Verifica se 1 + 2 + ... + n = n * (n + 1) / 2
	soma := 0
	for i := 1; i <= n; i++ {
		soma += i
	}
	return soma == n*(n+1)/2
}
