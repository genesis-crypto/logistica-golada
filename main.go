package main

import (
	"fmt"
	"logistica/rastreamento"
	"logistica/user"
)

func main() {
	productNotifier := rastreamento.NewPublisher(10, "a enviar", "carro")

	clientOne := &user.Client{
		Id: "user@user.com",
	}
	productNotifier.Register(clientOne)

	productNotifier.UpdateEvent()
	fmt.Println("running")
}
