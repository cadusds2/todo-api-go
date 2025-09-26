package core_test

import (
	"context"
	"errors"
	"testing"

	"github.com/estudante/todo-api-go/internal/core"
)

type repositorioMemoria struct {
	dados map[string]core.Tarefa
}

func novoRepositorioMemoria() *repositorioMemoria {
	return &repositorioMemoria{dados: make(map[string]core.Tarefa)}
}

func (r *repositorioMemoria) Criar(ctx context.Context, tarefa core.Tarefa) (core.Tarefa, error) {
	if ctx.Err() != nil {
		return core.Tarefa{}, ctx.Err()
	}
	r.dados[tarefa.ID] = tarefa
	return tarefa, nil
}

func (r *repositorioMemoria) BuscarTodas(ctx context.Context) ([]core.Tarefa, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	resultado := make([]core.Tarefa, 0, len(r.dados))
	for _, tarefa := range r.dados {
		resultado = append(resultado, tarefa)
	}
	return resultado, nil
}

func (r *repositorioMemoria) BuscarPorID(ctx context.Context, id string) (core.Tarefa, error) {
	if ctx.Err() != nil {
		return core.Tarefa{}, ctx.Err()
	}
	return r.dados[id], nil
}

func (r *repositorioMemoria) Atualizar(ctx context.Context, tarefa core.Tarefa) (core.Tarefa, error) {
	if ctx.Err() != nil {
		return core.Tarefa{}, ctx.Err()
	}
	r.dados[tarefa.ID] = tarefa
	return tarefa, nil
}

func (r *repositorioMemoria) Remover(ctx context.Context, id string) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	delete(r.dados, id)
	return nil
}

func TestNovoServicoTarefasValidaDependencia(t *testing.T) {
	_, err := core.NovoServicoTarefas(nil)
	if err == nil {
		t.Fatalf("esperava erro ao criar servico sem repositorio")
	}
}

func TestServicoTarefasCriarValidaCamposObrigatorios(t *testing.T) {
	repositorio := novoRepositorioMemoria()
	servico, err := core.NovoServicoTarefas(repositorio)
	if err != nil {
		t.Fatalf("nao esperava erro ao criar servico: %v", err)
	}

	testeTabela := []struct {
		nome         string
		entrada      core.NovaEntradaTarefa
		erroEsperado error
	}{
		{
			nome: "titulo obrigatorio",
			entrada: core.NovaEntradaTarefa{
				Titulo:    "",
				Descricao: "descricao valida",
			},
			erroEsperado: core.ErroTituloObrigatorio(),
		},
		{
			nome: "descricao obrigatoria",
			entrada: core.NovaEntradaTarefa{
				Titulo:    "titulo valido",
				Descricao: "",
			},
			erroEsperado: core.ErroDescricaoObrigatoria(),
		},
	}

	for _, tt := range testeTabela {
		t.Run(tt.nome, func(t *testing.T) {
			_, err := servico.Criar(context.Background(), tt.entrada)
			if err == nil {
				t.Fatalf("esperava erro %v", tt.erroEsperado)
			}
			if !errors.Is(err, tt.erroEsperado) {
				t.Fatalf("esperava erro %v, recebeu %v", tt.erroEsperado, err)
			}
		})
	}
}

func TestServicoTarefasFluxoCompleto(t *testing.T) {
	repositorio := novoRepositorioMemoria()
	servico, err := core.NovoServicoTarefas(repositorio)
	if err != nil {
		t.Fatalf("nao esperava erro ao criar servico: %v", err)
	}

	ctx := context.Background()

	tarefa, err := servico.Criar(ctx, core.NovaEntradaTarefa{Titulo: "Estudar Go", Descricao: "Praticar testes"})
	if err != nil {
		t.Fatalf("nao esperava erro ao criar tarefa: %v", err)
	}

	tarefas, err := servico.BuscarTodas(ctx)
	if err != nil {
		t.Fatalf("nao esperava erro ao listar tarefas: %v", err)
	}
	if len(tarefas) != 1 {
		t.Fatalf("esperava 1 tarefa, recebeu %d", len(tarefas))
	}

	obteve, err := servico.BuscarPorID(ctx, tarefa.ID)
	if err != nil {
		t.Fatalf("nao esperava erro ao buscar por id: %v", err)
	}
	if obteve.ID != tarefa.ID {
		t.Fatalf("esperava id %s, recebeu %s", tarefa.ID, obteve.ID)
	}

	tempoAnterior := obteve.Atualizado

	atualizada, err := servico.Atualizar(ctx, tarefa.ID, core.AtualizacaoTarefa{
		Titulo:    "Estudar Go",
		Descricao: "Praticar testes com tabela",
		Concluida: true,
	})
	if err != nil {
		t.Fatalf("nao esperava erro ao atualizar tarefa: %v", err)
	}
	if !atualizada.Concluida {
		t.Fatalf("esperava tarefa concluida")
	}
	if atualizada.Atualizado.Before(tempoAnterior) {
		t.Fatalf("esperava campo atualizado com tempo igual ou posterior")
	}

	if err := servico.Remover(ctx, tarefa.ID); err != nil {
		t.Fatalf("nao esperava erro ao remover tarefa: %v", err)
	}

	_, err = servico.BuscarPorID(ctx, tarefa.ID)
	if !errors.Is(err, core.ErroTarefaNaoEncontrada()) {
		t.Fatalf("esperava erro de nao encontrado apos remocao, recebeu %v", err)
	}
}
