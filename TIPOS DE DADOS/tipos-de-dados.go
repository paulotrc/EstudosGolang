package main

import (
	"errors"
	"fmt"
)

func main() {
	val := 10000000
	var valor int = 100000
	var valor1 int8 = 100
	var valor2 int16 = 1000
	var valor3 int32 = 10000
	var valor4 int64 = 100000

	var valor5 uint = 500000000

	fmt.Println(val, valor, valor1, valor2, valor3, valor4, valor5)

	var valor6 rune = 1000000
	fmt.Println(valor6)

	var valor8 byte = 100
	fmt.Println(valor8)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123999999.45
	fmt.Println(numeroReal2)

	char := 'A'
	fmt.Println(char)

	var booleano bool = true
	fmt.Println(booleano)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
