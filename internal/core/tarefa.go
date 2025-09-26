package core

import "context"
import "time"

// Tarefa representa uma tarefa a ser executada.
type Tarefa struct {
	ID         string    `json:"id"`
	Titulo     string    `json:"titulo"`
	Descricao  string    `json:"descricao"`
	Concluida  bool      `json:"concluida"`
	CriadoEm   time.Time `json:"criadoEm"`
	Atualizado time.Time `json:"atualizado"`
}

// RepositorioTarefas define as operações necessárias para persistir tarefas.
type RepositorioTarefas interface {
	Criar(ctx context.Context, tarefa Tarefa) (Tarefa, error)
	BuscarTodas(ctx context.Context) ([]Tarefa, error)
	BuscarPorID(ctx context.Context, id string) (Tarefa, error)
	Atualizar(ctx context.Context, tarefa Tarefa) (Tarefa, error)
	Remover(ctx context.Context, id string) error
}
