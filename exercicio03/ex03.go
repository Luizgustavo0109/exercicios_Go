package main

import (
	"fmt"
)

func main() {
	var nota, total float64
	var bim int

	for bim = 1; bim <= 4; bim++ {
		for {
			fmt.Printf("Digite a nota do %d° bimestre (entre 0 e 10): ", bim)
			fmt.Scan(&nota)

			if nota >= 0 && nota <= 10 {
				break
			} else {
				fmt.Println("Nota inválida. Digite um valor entre 0 e 10.")
			}
		}
		total += nota
	}

	media := calcularMedia(total, 4)
	fmt.Printf("\nMédia final: %.2f\n", media)
}

// Função que calcula a média
func calcularMedia(soma float64, quantidade int) float64 {
	return soma / float64(quantidade)
}
