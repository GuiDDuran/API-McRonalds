package processamento

import (
	"Projeto/modelos/pedido"
	"Projeto/handlers/loja"
	"fmt"
	"time"
)

func ProcessarPedidos() {
	for {
		if loja_handler.Aberta {
			pedidosAtivos := pedido.FPedidos.ListarPedidos()

			if len(pedidosAtivos) > 0 {
				pedido.FPedidos.Expedir()
			} else {
				fmt.Println("Não há pedidos ativos.")
			}
			time.Sleep(30 * time.Second)
		} else {
			fmt.Println("A loja está fechada.")
			return
		}
	}
}