package metricas

type Sistema struct {
    TotalProdutos     int     `json:"totalprodutos"`
    PedidosEncerrados int     `json:"pedidosencerrados"`
    PedidosAndamento  int     `json:"pedidosandamento"`
    FaturamentoTotal  float64 `json:"faturamentototal"`
}