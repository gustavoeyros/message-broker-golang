package main

import (
	"database/sql"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gustavoeyros/message-broker-golang/internal/infra/repository"
	"github.com/gustavoeyros/message-broker-golang/internal/infra/repository/akafka"
	"github.com/gustavoeyros/message-broker-golang/internal/usecase"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306/prodcuts")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	repository := repository.NewProductRepositoryMySql(db)
	createProductUseCase := usecase.NewCreateProductUseCase(repository)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
		}

		_, err = createProductUseCase.Execute(dto)
	}
}
