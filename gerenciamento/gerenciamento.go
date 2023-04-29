package gerenciamento

type GerenciamentoFrotas struct {
	Strategy Transporte
}

func CreateGerenciamento(transporte Transporte) *GerenciamentoFrotas {
	return &GerenciamentoFrotas{
		Strategy: transporte,
	}
}

func (g *GerenciamentoFrotas) SetStrategy(transporte Transporte) {
	g.Strategy = transporte
}

func (g *GerenciamentoFrotas) Transport() {
	g.Strategy.Transport()
}
