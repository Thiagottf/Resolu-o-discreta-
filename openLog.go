package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func AbrirLog() {
	file, err := os.Open("app.log")
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}
	Menu()
}
