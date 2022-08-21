package main

import (
	"fmt"

	"github.com/Carlos-Lopes1985/go-rest-api/configs"
	"github.com/Carlos-Lopes1985/go-rest-api/routes"
	"github.com/go-redis/redis"
)

func main() {

	type Cartao struct {
		Code  int
		Name  string
		Value float32
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

	var cartao Cartao

	//p := Cartao{}
	//client.Scan(val, &p)
	fmt.Println("VALOR...")
	fmt.Println(val)

	if err != nil {
		panic(err)
	}

	routes.HandleRequest()

}
