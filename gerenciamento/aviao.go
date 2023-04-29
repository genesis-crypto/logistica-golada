package gerenciamento

import "fmt"

type Aviao struct{}

func (a *Aviao) Transport() {
	fmt.Println("Transportando via avião")
}

func (a *Aviao) TranportMethodName() string {
	return "Avião"
}
