package pedido

import "fmt"

type CriarPedidoCommand struct {
	Pedido *Pedido
}

func (c *CriarPedidoCommand) Execute() error {
	fmt.Printf("Criando pedido com ID %d\n", c.Pedido.Id)
	return nil
}
