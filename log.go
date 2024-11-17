package main

import (
	"log"
	"os"
)

var infoLogger *log.Logger
var errorLogger *log.Logger

func init() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogInput(input string) {
	infoLogger.Println("Input:", input)
}

func LogEscolha(choice int) {
	infoLogger.Println("Choice:", choice)
}

func LogResposta(response string) {
	infoLogger.Println("Response:", response)
}
