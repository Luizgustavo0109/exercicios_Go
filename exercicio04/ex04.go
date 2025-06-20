package main

import (
	"fmt"
	"strings"
)

func main() {
	const quantidadeVetor int = 10
	var nomes [quantidadeVetor]string
	var idades [quantidadeVetor]int

	for i := 0; i < quantidadeVetor; i++ {
		fmt.Printf("Digite o %vº nome da pessoa: ", i+1)
		fmt.Scan(&nomes[i])

		fmt.Printf("Digite a idade de %v: ", nomes[i])
		fmt.Scan(&idades[i])
		fmt.Println()
	}

	var nomesPesquisados []string
	var consulta string

	for {
		fmt.Print("\nDeseja realizar uma consulta? (S/N): ")
		fmt.Scan(&consulta)
		consulta = strings.ToUpper(consulta)

		if consulta != "S" {
			break
		}

		var nomePesquisa string

		for {
			fmt.Print("Digite o nome a ser consultado: ")
			fmt.Scan(&nomePesquisa)

			// Verifica se já foi pesquisado
			jaPesquisado := false
			for _, nome := range nomesPesquisados {
				if nome == nomePesquisa {
					fmt.Println("\nNome já pesquisado anteriormente!")
					jaPesquisado = true
					break
				}
			}

			if !jaPesquisado {
				break
			}
		}

		// Adiciona nome à lista de nomes pesquisados
		nomesPesquisados = append(nomesPesquisados, nomePesquisa)

		// Procura o nome
		encontrado := false
		var ocorrencias int
		for i, nome := range nomes {
			if nome == nomePesquisa {
				fmt.Printf("A pessoa \"%v\" foi encontrada e possui %v anos.\n", nome, idades[i])
				ocorrencias++
				encontrado = true
			}
		}

		if !encontrado {
			fmt.Printf("A pessoa \"%v\" não foi encontrada.\n", nomePesquisa)
		} else {
			fmt.Printf("A pessoa \"%v\" foi encontrada %v vez(es).\n", nomePesquisa, ocorrencias)
		}
	}
}
