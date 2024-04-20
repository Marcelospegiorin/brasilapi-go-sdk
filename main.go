package main

import (
	"brasilapi/banks"
	"fmt"
)

func main() {
	banksResult, err := banks.GetAll()
	if err != nil {
		fmt.Println("Erro ao obter bancos:", err)
		return
	}

	for _, bank := range banksResult {
		fmt.Printf("Banco: %s - Código: %d\n", bank.Name, bank.Code)
	}
}
