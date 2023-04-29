package main

import (
	"fmt"
	"logistica/gerenciamento"
	"logistica/pedido"
	"logistica/rastreamento"
	"logistica/user"
)

func main() {
	novoPedido := &pedido.Pedido{
		Id:     1,
		Status: "pendente",
	}

	criarPedidoCommand := &pedido.CriarPedidoCommand{Pedido: novoPedido}
	invoker := &pedido.PedidoInvoker{}
	invoker.SetCommand(criarPedidoCommand)
	err := invoker.Run()
	if err != nil {
		fmt.Printf("Erro ao criar o pedido: %v\n", err)
	}

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
