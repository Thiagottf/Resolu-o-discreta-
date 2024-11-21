package main

import "fmt"

func Inducao(proposicao func(int) bool) {
	var n int
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

func proposicaoSoma(n int) bool {
	// Verifica se 1 + 2 + ... + n = n * (n + 1) / 2
	soma := 0
	for i := 1; i <= n; i++ {
		soma += i
	}
	return soma == n*(n+1)/2
}

func proposicaoQuadrado(n int) bool {
	// Verifica se 1^2 + 2^2 + ... + n^2 = n * (n + 1) * (2n + 1) / 6
	soma := 0
	for i := 1; i <= n; i++ {
		soma += i * i
	}
	return soma == n*(n+1)*(2*n+1)/6
}

func Indução() {
	var escolha int
	fmt.Println("Escolha a proposição para verificar:")
	fmt.Println("1 - Soma dos primeiros n números naturais")
	fmt.Println("2 - Soma dos quadrados dos primeiros n números naturais")
	fmt.Print("Opção: ")
	fmt.Scanf("%d", &escolha)

	switch escolha {
	case 1:
		Inducao(proposicaoSoma)
	case 2:
		Inducao(proposicaoQuadrado)
	default:
		fmt.Println("Opção inválida.")
	}
}
