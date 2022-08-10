package models

import (
	"log"

	"github.com/Carlos-Lopes1985/go-rest-api/db"
)

func Insert(compras Compra) (id int64, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO milhas (id_cartao, cpf, valor_compra) VALUES ($1, $2, $3) RETURNING id_milhas`

	err = conn.QueryRow(sql, compras.ID_Cartao, compras.Cpf_Cliente, compras.Valor_Compra).Scan(&id)

	return
}

func FindCpf(cpf string) (compras []Compra, err error) {

	log.Printf("INICIO FINDCPF: %v", cpf)

	conn, err := db.OpenConnection()

	if err != nil {
		log.Printf("ERRO BANCO:")
		return nil, err
	}

	rows := conn.QueryRow(`SELECT * FROM milhas WHERE id_milhas=$1`, 1)

	//if err != nil {
	//	log.Printf("ERRO QUERY:")
	//	return nil, err
	//}

	defer conn.Close()

	log.Printf("ROWS: %v", rows)
	log.Printf("ERROS: %v", err)

	//for rows.Next() {
	var compra Compra

	err = rows.Scan(&compra.ID_Cartao, &compra.Cpf_Cliente,
		&compra.Valor_Compra)

	log.Printf("LOOP QUERY: %v", compra)

	compras = append(compras, compra)

	log.Printf("INICIO FINDCPF: %v", compras)
	//}

	return compras, nil
}

func FindAll() (compras []Compra, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM compra`)

	if err != nil {
		return
	}

	for rows.Next() {
		var compra Compra

		err = rows.Scan(&compra.ID_Cartao, &compra.Cpf_Cliente, &compra.Valor_Compra)

		if err != nil {
			continue
		}

		compras = append(compras, compra)

	}

	return
}
