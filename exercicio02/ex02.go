package main

import (
	"fmt"

	"math"
)

func main() {

	var angulo float64
	fmt.Println("Digite o ângulo desejado: ")
	fmt.Scan(&angulo)

	fmt.Println("Cosseno do ângulo:", math.Cos(angulo))

	fmt.Println("Seno do ângulo:", math.Sin(angulo))
}
