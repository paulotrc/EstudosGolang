package main

import "fmt"

func main() {

	fmt.Println(20 * 10)
	soma := somar(10, 20)
	fmt.Println(soma)

	fmt.Println(calculosOperadores(20, 10))

	resSoma, resSubtracao, ResMultiplicacao, _ := calculosOperadores(20, 10)

	fmt.Println(resSoma, resSubtracao, ResMultiplicacao)

	var f = func(texto string) string {
		fmt.Println(texto)
		return texto
	}

	resultado := f("Texto para imprimir.")
	fmt.Println(resultado)

}

func somar(var1 int8, var2 int8) int8 {
	return var1 + var2
}

func calculosOperadores(v1, v2 int16) (int16, int16, int16, int16) {
	soma := v1 + v2
	subtracao := v1 - v2
	multiplicacao := (v1 * v2)
	divisao := v1 / v2
	return soma, subtracao, multiplicacao, divisao
}
