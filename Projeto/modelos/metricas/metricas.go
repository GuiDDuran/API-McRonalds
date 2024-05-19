package metricas

type Sistema struct {
    TotalProdutos       int     `json:"totalprodutos"`
    PedidosEncerrados   int     `json:"pedidosencerrados"`
    PedidosAndamento    int     `json:"pedidosandamento"`
    FaturamentoTotal    float64 `json:"faturamentototal"`
    TicketMedio         float64 `json:"ticketmedio"`
    TempoFuncionamento  int64  `json:"tempofuncionamento"`
}

var Metricas Sistema