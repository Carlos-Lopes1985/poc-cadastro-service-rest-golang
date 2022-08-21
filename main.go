package main

import (
	"encoding/json"
	"fmt"

	"github.com/Carlos-Lopes1985/go-rest-api/configs"
	"github.com/Carlos-Lopes1985/go-rest-api/routes"
	"github.com/go-redis/redis"
)

func main() {

	type Cartao struct {
		Code  int     `json:"code"`
		Name  string  `json:"name"`
		Value float32 `json:"valor"`
	}

	fmt.Println("Iniciando o servidor Rest com Go...")

	err := configs.Load()

	fmt.Println("Iniciando REDIS...")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("name", "Elliot", 0).Err()

	if err != nil {
		fmt.Println(err)
	}

	val, err := client.HGet("REPOSITORY_ID_CARTAO", "1233").Result()
	if err != nil {
		fmt.Println(err)
	}

	/*
		type Request struct {
			Operation string      `json:"operation"`
			Key string            `json:"key"`
			Value string          `json:"value"`
		}

		func main() {
			s := string(`{"operation": "get", "key": "example"}`)
			data := Request{}
			json.Unmarshal([]byte(s), &data)
			fmt.Printf("Operation: %s", data.Operation)
		}
	*/

	cartao := Cartao{}
	json.Unmarshal([]byte(val), &cartao)

	//p := Cartao{}
	//client.Scan(val, &p)
	fmt.Printf("VALOR... %f", cartao.Value)
	fmt.Println("")

	if err != nil {
		panic(err)
	}

	routes.HandleRequest()

}
