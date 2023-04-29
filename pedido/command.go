package pedido

type PedidoCommand interface {
	Execute() error
}
