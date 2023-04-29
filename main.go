package main

import (
	"fmt"
	"logistica/gerenciamento"
	"logistica/rastreamento"
	"logistica/user"
)

func main() {
	tranporteMethod := gerenciamento.Aviao{}
	gerenciamentoFrota := gerenciamento.CreateGerenciamento(&tranporteMethod)
	gerenciamentoFrota.Transport()

	productNotifier := rastreamento.NewPublisher(10, "a enviar", gerenciamentoFrota.Strategy)

	clientOne := &user.Client{
		Id: "user@user.com",
	}
	productNotifier.Register(clientOne)

	productNotifier.UpdateEvent()
	fmt.Println("running")
}
