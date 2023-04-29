package pedido

import "fmt"

type PedidoInvoker struct {
	command PedidoCommand
}

func (i *PedidoInvoker) SetCommand(command PedidoCommand) {
	i.command = command
}

func (i *PedidoInvoker) Run() error {
	if i.command == nil {
		return fmt.Errorf("O comando não foi definido")
	}
	return i.command.Execute()
}
