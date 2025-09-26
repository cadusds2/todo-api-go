package memoria

import (
	"context"
	"errors"
	"sync"

	"github.com/estudante/todo-api-go/internal/core"
)

var erroOperacaoCancelada = errors.New("operacao cancelada pelo contexto")

// RepositorioTarefasMemoria implementa RepositorioTarefas usando um mapa em memoria.
type RepositorioTarefasMemoria struct {
	mutex *sync.RWMutex
	dados map[string]core.Tarefa
}

// NovoRepositorioTarefasMemoria cria uma instancia pronta para uso.
func NovoRepositorioTarefasMemoria() *RepositorioTarefasMemoria {
	return &RepositorioTarefasMemoria{
		mutex: &sync.RWMutex{},
		dados: make(map[string]core.Tarefa),
	}
}

func (r *RepositorioTarefasMemoria) Criar(ctx context.Context, tarefa core.Tarefa) (core.Tarefa, error) {
	if ctx.Err() != nil {
		return core.Tarefa{}, erroOperacaoCancelada
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.dados[tarefa.ID] = tarefa
	return tarefa, nil
}

func (r *RepositorioTarefasMemoria) BuscarTodas(ctx context.Context) ([]core.Tarefa, error) {
	if ctx.Err() != nil {
		return nil, erroOperacaoCancelada
	}

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	resultado := make([]core.Tarefa, 0, len(r.dados))
	for _, tarefa := range r.dados {
		resultado = append(resultado, tarefa)
	}

	return resultado, nil
}

func (r *RepositorioTarefasMemoria) BuscarPorID(ctx context.Context, id string) (core.Tarefa, error) {
	if ctx.Err() != nil {
		return core.Tarefa{}, erroOperacaoCancelada
	}

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	tarefa, ok := r.dados[id]
	if !ok {
		return core.Tarefa{}, nil
	}
	return tarefa, nil
}

func (r *RepositorioTarefasMemoria) Atualizar(ctx context.Context, tarefa core.Tarefa) (core.Tarefa, error) {
	if ctx.Err() != nil {
		return core.Tarefa{}, erroOperacaoCancelada
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.dados[tarefa.ID] = tarefa
	return tarefa, nil
}

func (r *RepositorioTarefasMemoria) Remover(ctx context.Context, id string) error {
	if ctx.Err() != nil {
		return erroOperacaoCancelada
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	delete(r.dados, id)
	return nil
}
