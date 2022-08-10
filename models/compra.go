package models

import "log"

// Compra
type Compra struct {
	ID_Cartao    int     `json:"id_cartao"`
	Cpf_Cliente  string  `json:"cpf_cliente"`
	Valor_Compra float32 `json:"valor_compra"`
}

type Milhas struct {
	Cpf_Cliente        string `json:"cpf_cliente"`
	Valor_Total_Milhas int    `json:"valor_total_milhas"`
}

// Compras
var compras []Compra

func ReturnCalculoMilhas(cpf string) Milhas {

	compras, err := FindCpf(cpf)

	log.Printf("0003: %v", compras)

	log.Printf("CPF CALCULO MILHAS: %v", cpf)

	if err != nil {
		log.Printf("Erro Calculo de milhas: %v", compras)

	}

	var total_compra float32 = 0
	var total_milhas int = 0

	for _, element := range compras {
		total_compra += element.Valor_Compra

		log.Printf("0001: %f", total_compra)
	}

	total_milhas = int(total_compra / 7)

	log.Printf("0002: %d", total_milhas)

	var milhas = Milhas{
		cpf, total_milhas,
	}

	log.Printf("MILHAS: %v", milhas)

	return milhas
}

/* func ReturnListCompras() []Compra {
	Compras = []Compra{
		Compra{ID_Cartao: 12,
			Cpf_Cliente:  "12345678900",
			Valor_Compra: 200.00},
		Compra{ID_Cartao: 12,
			Cpf_Cliente:  "12345678900",
			Valor_Compra: 4550.37},
	}
	return Compras
}

func ReturnTotalMilhas() int {
	return ReturnCalculoMilhas(ReturnListCompras())
}
*/
