package models

import (
	"log"
)

type Compra struct {
	ID_Cartao    int     `json:"id_cartao"`
	Cpf_Cliente  string  `json:"cpf_cliente"`
	Valor_Compra float32 `json:"valor_compra"`
}

type Milhas struct {
	Cpf_Cliente        string `json:"cpf_cliente"`
	Valor_Total_Milhas int    `json:"valor_total_milhas"`
}

func ReturnCalculoMilhas(cpf string) (milhas Milhas, err error) {

	var total_compra float32 = 0
	var total_milhas int = 0

	compras, err := FindCpf(cpf)

	if err != nil {
		log.Printf("Erro Calculo de milhas: %v", err)
		return milhas, err
	}

	for _, element := range compras {
		total_compra += element.Valor_Compra
	}

	total_milhas = int(total_compra / 7)

	log.Printf("Total Milhas: %d", total_milhas)

	milhas = Milhas{
		cpf, total_milhas,
	}

	return milhas, nil
}
