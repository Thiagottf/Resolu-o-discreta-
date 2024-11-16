package main

import "fmt"

// Soma modular
func modAdd(a, b, m int) int {
	return (a + b) % m
}

// Subtração modular
func modSub(a, b, m int) int {
	return (a - b + m) % m // Adiciona m para garantir que o resultado seja positivo
}

// Multiplicação modular
func modMul(a, b, m int) int {
	return (a * b) % m
}

// Função para calcular o inverso modular usando o Algoritmo de Euclides Extendido
func modInv(a, m int) int {
	t, newT := 0, 1
	r, newR := m, a

	for newR != 0 {
		quotient := r / newR
		t, newT = newT, t-quotient*newT
		r, newR = newR, r-quotient*newR
	}

	if r > 1 {
		// Inverso não existe
		return -1
	}
	if t < 0 {
		t = t + m
	}
	return t
}

// Divisão modular (usa o inverso modular)
func modDiv(a, b, m int) int {
	invB := modInv(b, m)
	if invB == -1 {
		// Divisão não é possível (sem inverso modular)
		fmt.Println("Inverso modular não existe para", b)
		return -1
	}
	return modMul(a, invB, m)
}

func Aritmeticas() {
	var a, b, m int

	// Leitura dos valores de entrada
	fmt.Print("Digite o valor de a: ")
	fmt.Scan(&a)
	fmt.Print("Digite o valor de b: ")
	fmt.Scan(&b)
	fmt.Print("Digite o valor de m: ")
	fmt.Scan(&m)

	// Operações aritméticas modulares
	soma := modAdd(a, b, m)
	subtracao := modSub(a, b, m)
	multiplicacao := modMul(a, b, m)
	divisao := modDiv(a, b, m)

	// Exibindo os resultados
	fmt.Printf("%d + %d mod %d = %d\n", a, b, m, soma)
	fmt.Printf("%d - %d mod %d = %d\n", a, b, m, subtracao)
	fmt.Printf("%d * %d mod %d = %d\n", a, b, m, multiplicacao)
	if divisao != -1 {
		fmt.Printf("%d / %d mod %d = %d\n", a, b, m, divisao)
	}
}
