package main

import (
	"fmt"
	"os"
)

func soma(num float32, num2 float32) float32 {
	return num + num2
}

func subtracao(num float32, num2 float32) float32 {
	return num - num2
}

func multiplicacao(num float32, num2 float32) float32 {
	return num * num2
}

func divisao(num float32, num2 float32) float32 {
	return num / num2
}

func main() {
	for {
		var opcao int
		var num float32
		var num2 float32
		fmt.Println("Bem-Vindo a Calculadora!")

		fmt.Println("Digite 1 para escolher soma")
		fmt.Println("Digite 2 para escolher subtração")
		fmt.Println("Digite 3 para escolher multiplicação")
		fmt.Println("Digite 4 para escolher divisão")
		fmt.Println("Digite 5 para encerrar a calculadora")

		fmt.Println("Digite aqui a sua escolha: ")
		_, err := fmt.Scanln(&opcao)
		if err != nil {
			fmt.Println("Erro ao ler a entrada:", err)
			os.Exit(1)
		}

		if opcao == 5 {
			os.Exit(0)
		}

		fmt.Println("Digite o primeiro número: ")
		_, err = fmt.Scanln(&num)
		if err != nil {
			fmt.Println("Erro ao ler a entrada:", err)
			os.Exit(1)
		}

		fmt.Println("Digite o segundo número: ")
		_, err = fmt.Scanln(&num2)
		if err != nil {
			fmt.Println("Erro ao ler a entrada:", err)
			os.Exit(1)
		}

		switch opcao {
		case 1:
			fmt.Println("O resultado da soma é: ", soma(num, num2))
		case 2:
			fmt.Println("O resultado da subtração é: ", subtracao(num, num2))
		case 3:
			fmt.Println("O resultado da multiplicação é: ", multiplicacao(num, num2))
		case 4:
			if num2 == 0 {
				fmt.Println("Erro: divisão por zero!")
			} else {
				fmt.Println("O resultado da divisão é: ", divisao(num, num2))
			}
		default:
			fmt.Println("Essa opção é inválida, tente outra!")
		}
	}
}