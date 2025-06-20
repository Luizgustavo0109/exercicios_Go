package main

import "fmt"

func main() {

	var altura, base, area float32

	fmt.Print("Digite a altura do triângulo: ")
	fmt.Scan(&base)
	fmt.Print("Digite a base do triângulo: ")
	fmt.Scan(&altura)

	area = (base * altura) / 2

	fmt.Printf("A ·rea do triângulo é: %v", area)
}
