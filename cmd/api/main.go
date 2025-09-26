package main

import (
	"log"

	"github.com/estudante/todo-api-go/internal/api"
	"github.com/estudante/todo-api-go/internal/core"
	"github.com/estudante/todo-api-go/internal/infra/memoria"
)

func main() {
	repositorio := memoria.NovoRepositorioTarefasMemoria()

	servico, err := core.NovoServicoTarefas(repositorio)
	if err != nil {
		log.Fatalf("nao foi possivel iniciar o servico de tarefas: %v", err)
	}

	roteador, err := api.NovoRoteador(servico)
	if err != nil {
		log.Fatalf("nao foi possivel configurar o roteador: %v", err)
	}

	if err := roteador.Run(":8080"); err != nil {
		log.Fatalf("erro ao iniciar servidor HTTP: %v", err)
	}
}
