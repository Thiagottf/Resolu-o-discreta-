package main

import (
	"fmt"
	"time"
)

// Verifica se um número é primo
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Calcula o MDC (máximo divisor comum) usando o algoritmo de Euclides
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Resolve a congruência ax ≡ b (mod m)
// Retorna x e o valor mod m, ou -1 se não houver solução
func solveCongruence(a, b, m int) (int, int) {
	g := gcd(a, m)
	if b%g != 0 {
		return -1, -1 // Sem solução
	}
	// Simplifica a congruência dividindo tudo por g
	a /= g
	b /= g
	m /= g

	x, _ := extendedGCD(a, m)
	x = (x*b%m + m) % m // Garante que x esteja no intervalo [0, m-1]
	return x, m
}

// Algoritmo de Euclides Estendido
func extendedGCD(a, b int) (int, int) {
	if b == 0 {
		return 1, 0
	}
	x1, y1 := extendedGCD(b, a%b)
	x, y := y1, x1-(a/b)*y1
	return x, y
}

// Gera todos os números primos até um limite usando o Crivo de Eratóstenes
func sieve(limit int) []int {
	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}

	for p := 2; p*p <= limit; p++ {
		if isPrime[p] {
			for i := p * p; i <= limit; i += p {
				isPrime[i] = false
			}
		}
	}

	primes := []int{}
	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func TeoriaDosNumeros() {
	var option int

	fmt.Println("Bem-vindo ao Resolver de Teoria dos Números!")
	fmt.Println("Escolha uma funcionalidade:")
	fmt.Println("1 - Verificar se um número é primo")
	fmt.Println("2 - Calcular o MDC de dois números")
	fmt.Println("3 - Resolver uma congruência linear")
	fmt.Println("4 - Gerar números primos até um limite")
	fmt.Println("5 - Voltar ao menu principal")
	fmt.Print("Opção: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		var n int
		fmt.Print("Digite o número para verificar se é primo: ")
		fmt.Scan(&n)
		if isPrime(n) {
			fmt.Printf("%d é primo.\n", n)
		} else {
			fmt.Printf("%d não é primo.\n", n)
		}
	case 2:
		var a, b int
		fmt.Print("Digite dois números para calcular o MDC (a, b): ")
		fmt.Scan(&a, &b)
		fmt.Printf("O MDC de %d e %d é %d.\n", a, b, gcd(a, b))
	case 3:
		var a, b, m int
		fmt.Print("Digite os valores de a, b e m para resolver ax ≡ b (mod m): ")
		fmt.Scan(&a, &b, &m)
		x, mod := solveCongruence(a, b, m)
		if x == -1 {
			fmt.Println("Não há solução para a congruência.")
		} else {
			fmt.Printf("Uma solução para %dx ≡ %d (mod %d) é x ≡ %d (mod %d).\n", a, b, m, x, mod)
		}
	case 4:
		var limit int
		fmt.Print("Digite o limite para gerar números primos: ")
		fmt.Scan(&limit)
		primes := sieve(limit)
		fmt.Printf("Os números primos até %d são: %v\n", limit, primes)
	case 5:
		fmt.Println("CARREGANDO MENU PRINCIPAL...")
		time.Sleep(2 * time.Second)
		Menu()
	default:
		fmt.Println("Opção inválida.")
	}
}
