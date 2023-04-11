package main

import "fmt"

//Crie uma função que receba dois parâmetros (a e b) e retorne a divisão entre eles. Caso o segundo parâmetro seja zero, retorne um erro.

func main() {
	var a, b int
	fmt.Println("Escolha um valor de a")
	fmt.Scan(&a)
	fmt.Println("Escolha um valor de b")
	fmt.Scan(&b)
	resultado, err := divisão(a, b)
	if err != nil {
		fmt.Printf("Ocorreu um erro ao dividir a e b: %s\n", err)
		return
	}
	fmt.Println(resultado)

}

func divisão(a int, b int) (int, error) {

	if b == 0 {
		return 0, fmt.Errorf("É... vai dar 0")
	}
	return a / b, nil

}
