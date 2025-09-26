package core

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	erroTituloObrigatorio    = errors.New("titulo obrigatorio")
	erroDescricaoObrigatoria = errors.New("descricao obrigatoria")
	erroTarefaNaoEncontrada  = errors.New("tarefa nao encontrada")
)

// ServicoTarefas encapsula regras de negocio para tarefas.
type ServicoTarefas struct {
	repositorio RepositorioTarefas
}

// NovaEntradaTarefa representa dados para criar uma tarefa.
type NovaEntradaTarefa struct {
	Titulo    string
	Descricao string
}

// AtualizacaoTarefa representa dados para atualizar uma tarefa.
type AtualizacaoTarefa struct {
	Titulo    string
	Descricao string
	Concluida bool
}

// NovoServicoTarefas cria um novo servico com validacao de dependencias.
func NovoServicoTarefas(repositorio RepositorioTarefas) (*ServicoTarefas, error) {
	if repositorio == nil {
		return nil, errors.New("repositorio de tarefas obrigatorio")
	}

	return &ServicoTarefas{repositorio: repositorio}, nil
}

// Criar gera uma nova tarefa apos validar os dados.
func (s *ServicoTarefas) Criar(ctx context.Context, entrada NovaEntradaTarefa) (Tarefa, error) {
	if entrada.Titulo == "" {
		return Tarefa{}, erroTituloObrigatorio
	}
	if entrada.Descricao == "" {
		return Tarefa{}, erroDescricaoObrigatoria
	}

	agora := time.Now().UTC()
	tarefa := Tarefa{
		ID:         uuid.NewString(),
		Titulo:     entrada.Titulo,
		Descricao:  entrada.Descricao,
		Concluida:  false,
		CriadoEm:   agora,
		Atualizado: agora,
	}

	return s.repositorio.Criar(ctx, tarefa)
}

// BuscarTodas retorna a lista completa de tarefas.
func (s *ServicoTarefas) BuscarTodas(ctx context.Context) ([]Tarefa, error) {
	return s.repositorio.BuscarTodas(ctx)
}

// BuscarPorID retorna uma tarefa especifica.
func (s *ServicoTarefas) BuscarPorID(ctx context.Context, id string) (Tarefa, error) {
	tarefa, err := s.repositorio.BuscarPorID(ctx, id)
	if err != nil {
		return Tarefa{}, err
	}
	if tarefa.ID == "" {
		return Tarefa{}, erroTarefaNaoEncontrada
	}
	return tarefa, nil
}

// Atualizar aplica novos dados a uma tarefa existente.
func (s *ServicoTarefas) Atualizar(ctx context.Context, id string, dados AtualizacaoTarefa) (Tarefa, error) {
	if dados.Titulo == "" {
		return Tarefa{}, erroTituloObrigatorio
	}
	if dados.Descricao == "" {
		return Tarefa{}, erroDescricaoObrigatoria
	}

	tarefaExistente, err := s.repositorio.BuscarPorID(ctx, id)
	if err != nil {
		return Tarefa{}, err
	}
	if tarefaExistente.ID == "" {
		return Tarefa{}, erroTarefaNaoEncontrada
	}

	tarefaExistente.Titulo = dados.Titulo
	tarefaExistente.Descricao = dados.Descricao
	tarefaExistente.Concluida = dados.Concluida
	tarefaExistente.Atualizado = time.Now().UTC()

	return s.repositorio.Atualizar(ctx, tarefaExistente)
}

// Remover exclui uma tarefa por id.
func (s *ServicoTarefas) Remover(ctx context.Context, id string) error {
	tarefaExistente, err := s.repositorio.BuscarPorID(ctx, id)
	if err != nil {
		return err
	}
	if tarefaExistente.ID == "" {
		return erroTarefaNaoEncontrada
	}

	return s.repositorio.Remover(ctx, id)
}

// ErroTarefaNaoEncontrada expõe erro padrão de não encontrado.
func ErroTarefaNaoEncontrada() error {
	return erroTarefaNaoEncontrada
}

// ErroTituloObrigatorio expõe erro de titulo invalido.
func ErroTituloObrigatorio() error {
	return erroTituloObrigatorio
}

// ErroDescricaoObrigatoria expõe erro de descricao invalida.
func ErroDescricaoObrigatoria() error {
	return erroDescricaoObrigatoria
}
