package models

import (
	"log"
	"reflect"

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

	rows, err := conn.Query(`SELECT cpf as Cpf_Cliente, id_cartao as Id_Cartao, valor_compra as Valor_Compra FROM milhas WHERE cpf=$1`, cpf)

	if err != nil {
		log.Printf("ERRO QUERY:")
		return nil, err
	}

	defer conn.Close()

	for rows.Next() {
		compra := Compra{}

		s := reflect.ValueOf(&compra).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err := rows.Scan(
			&compra.Cpf_Cliente,
			&compra.ID_Cartao,
			&compra.Valor_Compra,
		)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(compra)
		compras = append(compras, compra)
	}

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
