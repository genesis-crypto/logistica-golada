package sistema

import (
	"fmt"
	"logistica/gerenciamento"
	"logistica/pedido"
	"logistica/rastreamento"
)

type SistemaFacade struct {
	rastreamento  *rastreamento.Publisher
	pedido        *pedido.PedidoInvoker
	gerenciamento *gerenciamento.GerenciamentoFrotas
}

func NewSistemaFacade() *SistemaFacade {
	return &SistemaFacade{
		rastreamento:  nil,
		pedido:        nil,
		gerenciamento: nil,
	}
}

func (s *SistemaFacade) CreatePedido() error {
	fmt.Println("Criando Pedido")
	return nil
}

func (s *SistemaFacade) ExcluiPedido() error {
	fmt.Println("Excluindo Pedido")
	return nil
}

func (s *SistemaFacade) TransportaPedido() error {
	fmt.Println("Transportando Pedido")
	return nil
}
