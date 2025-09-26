package api

import (
	"github.com/gin-gonic/gin"

	"github.com/estudante/todo-api-go/internal/core"
)

// NovoRoteador registra rotas e middlewares basicos.
func NovoRoteador(servico *core.ServicoTarefas) (*gin.Engine, error) {
	if servico == nil {
		return nil, ErrServicoObrigatorio
	}

	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())

	manipulador := NovoManipuladorTarefas(servico)

	r.GET("/saude", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	grupoTarefas := r.Group("/tarefas")
	{
		grupoTarefas.GET("", manipulador.Listar)
		grupoTarefas.POST("", manipulador.Criar)
		grupoTarefas.GET(":id", manipulador.BuscarPorID)
		grupoTarefas.PUT(":id", manipulador.Atualizar)
		grupoTarefas.DELETE(":id", manipulador.Remover)
	}

	return r, nil
}
