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
			if len(pedido.FPedidos.Pedidos) > 0 {
				pedido.FPedidos.Expedir()
			} else {
				fmt.Println("Não há pedidos ativos.")
			}
			time.Sleep(5 * time.Second)
		}
	}
}