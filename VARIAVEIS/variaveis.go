package main

import "fmt"

func main() {
	var variavel1 string = "Variavel1"
	fmt.Println(variavel1)

	variavel2 := "Variavel2"
	fmt.Println(variavel2)

	var (
		variavel3 string = "variavel3"
		variavel4 string = "veriavel4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel5", "Variavel6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}
