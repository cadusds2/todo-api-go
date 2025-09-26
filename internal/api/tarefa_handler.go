package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/estudante/todo-api-go/internal/core"
	"github.com/estudante/todo-api-go/pkg/resposta"
)

// ErrServicoObrigatorio indica que o servico de tarefas é obrigatorio.
var ErrServicoObrigatorio = errors.New("servico de tarefas obrigatorio")

// ManipuladorTarefas coordena as requisicoes HTTP relacionadas a tarefas.
type ManipuladorTarefas struct {
	servico *core.ServicoTarefas
}

// NovoManipuladorTarefas cria uma instancia validada de ManipuladorTarefas.
func NovoManipuladorTarefas(servico *core.ServicoTarefas) *ManipuladorTarefas {
	return &ManipuladorTarefas{servico: servico}
}

// requisicaoTarefa representa o payload esperado para criar ou atualizar uma tarefa.
type requisicaoTarefa struct {
	Titulo    string `json:"titulo" binding:"required"`
	Descricao string `json:"descricao" binding:"required"`
	Concluida bool   `json:"concluida"`
}

// Criar adiciona uma nova tarefa.
func (m *ManipuladorTarefas) Criar(ctx *gin.Context) {
	var corpo requisicaoTarefa
	if err := ctx.ShouldBindJSON(&corpo); err != nil {
		resposta.Erro(ctx, http.StatusBadRequest, "corpo invalido")
		return
	}

	tarefa, err := m.servico.Criar(ctx.Request.Context(), core.NovaEntradaTarefa{
		Titulo:    corpo.Titulo,
		Descricao: corpo.Descricao,
	})
	if err != nil {
		tratarErroDeDominio(ctx, err)
		return
	}

	resposta.Sucesso(ctx, http.StatusCreated, tarefa)
}

// Listar retorna todas as tarefas cadastradas.
func (m *ManipuladorTarefas) Listar(ctx *gin.Context) {
	tarefas, err := m.servico.BuscarTodas(ctx.Request.Context())
	if err != nil {
		resposta.Erro(ctx, http.StatusInternalServerError, "nao foi possivel listar tarefas")
		return
	}

	resposta.Sucesso(ctx, http.StatusOK, tarefas)
}

// BuscarPorID retorna uma tarefa especifica.
func (m *ManipuladorTarefas) BuscarPorID(ctx *gin.Context) {
	id := ctx.Param("id")
	tarefa, err := m.servico.BuscarPorID(ctx.Request.Context(), id)
	if err != nil {
		tratarErroDeDominio(ctx, err)
		return
	}

	resposta.Sucesso(ctx, http.StatusOK, tarefa)
}

// Atualizar modifica uma tarefa existente.
func (m *ManipuladorTarefas) Atualizar(ctx *gin.Context) {
	id := ctx.Param("id")
	var corpo requisicaoTarefa
	if err := ctx.ShouldBindJSON(&corpo); err != nil {
		resposta.Erro(ctx, http.StatusBadRequest, "corpo invalido")
		return
	}

	tarefa, err := m.servico.Atualizar(ctx.Request.Context(), id, core.AtualizacaoTarefa{
		Titulo:    corpo.Titulo,
		Descricao: corpo.Descricao,
		Concluida: corpo.Concluida,
	})
	if err != nil {
		tratarErroDeDominio(ctx, err)
		return
	}

	resposta.Sucesso(ctx, http.StatusOK, tarefa)
}

// Remover exclui uma tarefa pelo identificador.
func (m *ManipuladorTarefas) Remover(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := m.servico.Remover(ctx.Request.Context(), id); err != nil {
		tratarErroDeDominio(ctx, err)
		return
	}

	resposta.Sucesso(ctx, http.StatusOK, gin.H{"idRemovido": id})
}

func tratarErroDeDominio(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, core.ErroTarefaNaoEncontrada()):
		resposta.Erro(ctx, http.StatusNotFound, "tarefa nao encontrada")
	case errors.Is(err, core.ErroTituloObrigatorio()):
		resposta.Erro(ctx, http.StatusBadRequest, "titulo obrigatorio")
	case errors.Is(err, core.ErroDescricaoObrigatoria()):
		resposta.Erro(ctx, http.StatusBadRequest, "descricao obrigatoria")
	default:
		resposta.Erro(ctx, http.StatusInternalServerError, "erro interno")
	}
}
