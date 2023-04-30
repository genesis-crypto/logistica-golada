# __Sistema de Logística__
Este projeto consiste em uma aplicação de logística, que permite realizar a gestão de pedidos em questão. O projeto foi desenvolvido como parte da disciplina de Design Patterns.

## __Requisitos Funcionais__
* Permitir a criação de usuários
* Permitir a criação e exclusão de pedidos
* Permitir gerenciar o método de envio de um pedido
* Receber notificações ao atualizar o status do pedido

## __Requisitos Não-Funcionais__
* Utilizar Design Patterns para garantir uma arquitetura flexível e extensível

## __Estrutura de Arquivos__

O projeto está estruturado da seguinte forma:

```
logistica-golada/
  ├── gerenciamento/
  │   ├── aviao.go
  │   ├── gerenciamento.go
  │   └── veiculo.go
  ├── pedido/
  │   ├── command.go
  │   ├── criaPedidoCommand.go
  │   ├── excluiPedidoCommand.go
  │   ├── invoker.go
  │   └── pedido.go
  ├── rastreamento/
  │   ├── rastreamento.go
  │   ├── subject.go
  │   └── subscriber.go
  ├── sistema/
  │   └── facade.go
  ├── user/
  │   └── user.go
  ├── go.mod
  ├── main.mod
  └── README.md
```

A pasta gerenciamento contém a implementação do gerenciamento de transporte, com o uso do Strategy Pattern. A pasta pedido contém a implementação dos pedidos, com o uso do Command Pattern. A pasta rastreamento contém a implementação do rastreamento de pedidos, com o uso do Observer Pattern. A pasta sistema contém a implementação de todos os módulos, com o uso do Facade Pattern.

## __Design Patterns utilizados__

### __Strategy Pattern__
O Strategy Pattern foi utilizado na implementação do gerenciamento de transporte do sistema de logística. Esse padrão de projeto permitiu que diferentes estratégias de transporte fossem definidas e selecionadas de forma flexível e fácil de manter.

Foram criadas diferentes structs de estratégia de transporte, como TransportePorAviao, que implementam uma interface comum Transporte. Cada struct de estratégia de transporte define o comportamento do objeto Pedido para a estratégia de transporte correspondente.

Esse padrão de projeto também permite a fácil inclusão de novas estratégias de transporte no futuro, tornando o sistema de logística mais extensível.

### __Command Pattern__
O Command Pattern foi utilizado na implementação dos pedidos no sistema de logística. Esse padrão de projeto permitiu que os pedidos fossem tratados como comandos, que podem ser facilmente criados, alterados e excluídos.

Foram criadas diferentes structs de comando, como CriarPedidoCommand e ExcluirPedidoCommand, que implementam uma interface comum PedidoCommand. Cada struct de comando define o comportamento do objeto Pedido para o comando correspondente.

Ao executar um comando, a referência para a struct de comando correspondente é atualizada, permitindo que o objeto Pedido execute o comportamento apropriado. Esse padrão de projeto também permite a fácil inclusão de novos comandos no futuro, tornando o sistema de logística mais extensível.

### __Observer Pattern__
O Observer Pattern foi utilizado na implementação do rastreamento de pedidos no sistema de logística. Esse padrão de projeto permitiu que os diferentes usuários do sistema fossem notificados automaticamente sobre as atualizações de status dos pedidos.

Ao atualizar o status de um pedido, a referência para a struct de pedido correspondente é atualizada, permitindo que o objeto Pedido notifique automaticamente todos os assinantes registrados sobre a atualização de status. Esse padrão de projeto também permite a fácil inclusão de novos tipos de assinantes no futuro, tornando o sistema de rastreamento de pedidos mais extensível.


### __Facade Pattern__
O Facade Pattern foi utilizado na implementação do sistema de logística como um todo. Esse padrão de projeto permitiu que a complexidade do sistema fosse reduzida, tornando mais simples a utilização dos diferentes módulos e funcionalidades.

Foi criada uma struct de fachada, chamada SistemaFacade, que encapsula a complexidade dos diferentes módulos do sistema de logística. Essa struct oferece uma interface simplificada para os usuários do sistema, que podem criar usuários, pedidos e gerenciar o método de envio dos pedidos, sem precisar conhecer a complexidade de cada módulo.

Ao utilizar o Facade Pattern, o sistema de logística fica mais fácil de utilizar e manter, além de permitir uma melhor organização do código, já que cada módulo pode ser implementado e testado separadamente. Esse padrão de projeto também permite a fácil inclusão de novos módulos no futuro, tornando o sistema de logística mais extensível.

## __Contato__
Pedro Cardozo - `p-cardozo@hotmail.com` ou `609455@univem.edu.br`