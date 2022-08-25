package db

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
)

func OpenConnectionRedis() (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	if err != nil {
		log.Printf("Conexão com Erro! %s", err)
	} else {
		log.Printf("Conexão aberta com sucesso! %s", pong)
	}

	return client, err
}
