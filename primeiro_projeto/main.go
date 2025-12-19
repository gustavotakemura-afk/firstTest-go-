package main

import "fmt"

func main() {
	var nome string
	var idade int
	const anoAtual = 2025

	fmt.Print("Digite seu nome: ")
	fmt.Scanln(&nome)

	fmt.Print("Digite sua idade: ")
	fmt.Scanln(&idade)

	anoNas := anoAtual - idade

	fmt.Println("Olá ", nome)
	fmt.Println("Você possui ", idade, "anos")
	fmt.Println("Você nasceu em ", anoNas)
}
