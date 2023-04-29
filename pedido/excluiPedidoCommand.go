package pedido

import "fmt"

type ExcluirPedidoCommand struct {
	PedidoID int
}

func (c *ExcluirPedidoCommand) Execute() error {
	fmt.Printf("Excluindo pedido com ID %d\n", c.PedidoID)
	return nil
}
