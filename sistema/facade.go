package sistema

import (
	"errors"
	"fmt"
	"logistica/gerenciamento"
	"logistica/pedido"
	"logistica/rastreamento"
	"logistica/user"
)

type SistemaFacade struct {
	Rastreamento  *rastreamento.Publisher
	Pedido        *pedido.Pedido
	Gerenciamento *gerenciamento.GerenciamentoFrotas
	Usuario       *user.Client
}

func NewSistemaFacade() *SistemaFacade {
	return &SistemaFacade{
		Rastreamento:  nil,
		Pedido:        nil,
		Gerenciamento: nil,
		Usuario:       nil,
	}
}

func (s *SistemaFacade) CreatePedido(id int, status string) (*pedido.Pedido, error) {
	fmt.Println("(Command) - Criando Pedido")

	novoPedido := &pedido.Pedido{
		Id:     id,
		Status: status,
	}

	criarPedidoCommand := &pedido.CriarPedidoCommand{Pedido: novoPedido}
	invoker := &pedido.PedidoInvoker{}
	invoker.SetCommand(criarPedidoCommand)

	err := invoker.Run()

	s.Pedido = novoPedido
	return novoPedido, err
}

func (s *SistemaFacade) ExcluiPedido(id int) error {
	fmt.Println("(Command) - Excluindo Pedido")

	excluiPedidoCommand := &pedido.ExcluirPedidoCommand{PedidoID: id}
	invoker := &pedido.PedidoInvoker{}
	invoker.SetCommand(excluiPedidoCommand)

	err := invoker.Run()

	s.Pedido = nil
	return err
}

func (s *SistemaFacade) CreateUser(id string) (*user.Client, error) {
	fmt.Println("(Facade) - Criando Usuário")

	user := &user.Client{
		Id: id,
	}

	s.Usuario = user

	return user, nil
}

func (s *SistemaFacade) CreateNotifier(pedido *pedido.Pedido) error {
	fmt.Println("(Observer) - Cria Notificação")
	if s.Gerenciamento == nil {
		return errors.New("missing module Gerenciamento de Frotas")
	}

	notifier := rastreamento.NewPublisher(s.Pedido.Id, s.Pedido.Status, s.Gerenciamento.Strategy)
	s.Rastreamento = notifier

	return nil
}

func (s *SistemaFacade) RegistraUsuarioNotificacao(user *user.Client) error {
	fmt.Println("(Observer) - Registra Usuario Notificação")

	s.Rastreamento.Register(user)

	return nil
}

func (s *SistemaFacade) DesregistraNotificacao(user *user.Client) error {
	fmt.Println("(Observer) - Desregistra usuario notificação")

	s.Rastreamento.Register(user)

	return nil
}

func (s *SistemaFacade) CreateTransporte() error {
	fmt.Println("(Strategy) - Define método de transporte")

	tranporteMethod := gerenciamento.Aviao{}
	gerenciamentoFrota := gerenciamento.CreateGerenciamento(&tranporteMethod)

	s.Gerenciamento = gerenciamentoFrota

	return nil
}

func (s *SistemaFacade) TransportaCarga() error {
	fmt.Println("(Strategy) - Transporta carga")

	if s.Gerenciamento == nil {
		return errors.New("missing module Gerenciamento de Frotas")
	}

	s.Gerenciamento.Transport()

	if s.Rastreamento == nil {
		return errors.New("missing module Rastreamento")
	}
	s.Rastreamento.NotifyAll()

	return nil
}
