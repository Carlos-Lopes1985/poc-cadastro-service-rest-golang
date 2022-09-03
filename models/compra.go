package models

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/Carlos-Lopes1985/go-rest-api/db"
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

type Cartao struct {
	Code  int
	Name  string
	Valor float32
}

func ReturnValorRedis(id string) (valor float32) {
	conn, err := db.OpenConnectionRedis()

	val, err := conn.HGet("REPOSITORY_ID_CARTAO", id).Result()

	if err != nil {
		fmt.Println(err)
	}

	cartao := Cartao{}
	json.Unmarshal([]byte(val), &cartao)

	fmt.Printf("VALOR... %f", cartao.Valor)
	fmt.Println("")

	var valor_milhas_cartao float32 = cartao.Valor

	return valor_milhas_cartao
}

func ReturnCalculoMilhas(cpf string) (milhas Milhas, err error) {

	var total_compra float32 = 0
	var total_milhas int = 0

	compras, err := FindCpf(cpf)

	log.Printf("Retorno Objeto compras: %v", compras)

	if err != nil {
		log.Printf("Erro Calculo de milhas: %v", err)
		return milhas, err
	}

	for _, element := range compras {
		total_compra = element.Valor_Compra
		s := strconv.Itoa(element.ID_Cartao)
		var retorno_valor_milhas float32 = 0

		retorno_valor_milhas = ReturnValorRedis(s)

		log.Printf("Retorno busca REDIS: %v", retorno_valor_milhas)

		total_milhas += int(total_compra * retorno_valor_milhas)

		log.Printf("Retorno Calculo Milhas: %d", total_milhas)
	}

	log.Printf("Total Milhas: %d", total_milhas)

	milhas = Milhas{
		cpf, total_milhas,
	}

	return milhas, nil
}
