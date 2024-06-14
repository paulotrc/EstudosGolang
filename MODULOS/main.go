package main

import (
	"fmt"
	"modulos/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main do módulo")
	auxiliar.Escrever()
	resp := checkmail.ValidateFormat("teste___@gmail.com")
	fmt.Println(resp)
}
