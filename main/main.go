package main

import (
	"esquemaestrela/internal"
	"fmt"
)

func main() {
	fmt.Println("Inciando, Conectando ao banco de dados...")
	DB, err := internal.ConectaComBancodeDados()
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}
	fmt.Println("Banco de dados conectado, lendo arquivo csv...")
	// Remover espaços e quebras de linha do input
	// Ler o arquivo CSV e obter os dados
	notas, err := internal.ReaderCsv("MICRODADOS_IDD_2022_LGPD_v2.txt")
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}
	fmt.Println("Salvando dados no banco de dados...")
	if err := internal.SaveNotas(notas, DB); err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Println("Notas salvas com sucesso no banco.")
	fmt.Println("Fechando programa.")

}
