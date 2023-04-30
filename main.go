package main

import (
	"fmt"
	"logistica/sistema"
)

func main() {
	facade := sistema.NewSistemaFacade()
	pedidoFacade, err := facade.CreatePedido(1, "criando")
	if err != nil {
		fmt.Println("falha ao criar o pedido")
		return
	}

	userFacade, err := facade.CreateUser("user@user.com")
	if err != nil {
		fmt.Println("falha ao criar o usuário")
		return
	}

	facade.CreateTransporte()
	err = facade.CreateNotifier(pedidoFacade)
	if err != nil {
		fmt.Println(err)
		return
	}

	facade.RegistraUsuarioNotificacao(userFacade)
	err = facade.TransportaCarga()
	if err != nil {
		fmt.Println(err)
		return
	}
}
