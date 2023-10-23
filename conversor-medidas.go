package main //package main é obrigatório

// importação de pacotes externos
import (
	"fmt" // pacote que contém funções para formatação de strings
	"os" // pacote que lida com o sistema operacional hospedeiro
	"strconv" // pacote que contém funções para conversão de formatos ex.: str to float
)

func main() {
	// se a quantidade de argumentos informados for menor que três, será retornado um erro ao usuário com a mensagem indicada em PrintIn
	if len(os.Args) < 3 {
		fmt.Printf("Uso: conversor <valores> <unidade>")
		os.Exit(1)
	}

	unidadeOrigem := os.Args[len(os.Args)-1]
	valoresOrigem := os.Args[1 : len(os.Args)-1]

	var unidadeDestino string

	if unidadeOrigem == "celsius" {
		unidadeDestino = "fahrenheit"
	} else if unidadeOrigem == "quilometros" {
		unidadeDestino = "milhas"
	} else {
		fmt.Printf("%s não é uma unidade conhecida!",
			unidadeDestino)
		os.Exit(1)
	}


	for i, v := range valoresOrigem{
		valorOrigem, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf(
				"O valor %s na posição %d" +
					"não é um número válido!\n",
				v, i)
			os.Exit(1)
		}

		var valorDestino float64

		if unidadeOrigem == "celsius" {
			valorDestino = valorOrigem*1.8 + 32
		} else {
			valorDestino = valorOrigem / 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s\n",
			valorOrigem, unidadeOrigem,
			valorDestino, unidadeDestino)
	}
}
