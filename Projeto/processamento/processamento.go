package processamento

import (
	"Projeto/handlers/loja"
	"Projeto/modelos/metricas"
	"Projeto/modelos/pedido"
	"fmt"
	"time"
)

func logMessage(message string) {
	Horario := time.Now().Format("02/01/2006 15:04:05")
	fmt.Printf("%s - %s\n", Horario, message)
}

func ProcessarPedidos() {
	for {
		if loja_handler.Aberta {
			if len(pedido.FPedidos.Pedidos) > 0 {
				pedido.FPedidos.Expedir()
			} else {
				logMessage("Não há pedidos ativos.")
			}
			time.Sleep(loja_handler.IntervaloExpedicao)
		}
	}
}

func ContarTempoDeFuncionamento(){
	inicio := time.Now()
	for {
		metricas.Metricas.TempoFuncionamento = int64(time.Since(inicio).Seconds())
		time.Sleep(1 * time.Second)
	}
}