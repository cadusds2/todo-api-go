package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/estudante/todo-api-go/internal/core"
	"github.com/estudante/todo-api-go/internal/infra/memoria"
)

type envelopeTarefa struct {
	Dados core.Tarefa `json:"dados"`
}

type envelopeLista struct {
	Dados []core.Tarefa `json:"dados"`
}

type envelopePadrao struct {
	Dados map[string]interface{} `json:"dados"`
}

func TestManipuladorTarefasFluxoHTTP(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repositorio := memoria.NovoRepositorioTarefasMemoria()
	servico, err := core.NovoServicoTarefas(repositorio)
	if err != nil {
		t.Fatalf("nao esperava erro ao criar servico: %v", err)
	}

	roteador, err := NovoRoteador(servico)
	if err != nil {
		t.Fatalf("nao esperava erro ao criar roteador: %v", err)
	}

	requisicaoCriar := httptest.NewRequest(http.MethodPost, "/tarefas", strings.NewReader(`{"titulo":"Ler", "descricao":"Estudar Go"}`))
	requisicaoCriar.Header.Set("Content-Type", "application/json")
	gravadorCriar := httptest.NewRecorder()
	roteador.ServeHTTP(gravadorCriar, requisicaoCriar)

	if gravadorCriar.Code != http.StatusCreated {
		t.Fatalf("esperava status %d, recebeu %d", http.StatusCreated, gravadorCriar.Code)
	}

	var respostaCriar envelopeTarefa
	if err := json.Unmarshal(gravadorCriar.Body.Bytes(), &respostaCriar); err != nil {
		t.Fatalf("erro ao decodificar resposta de criacao: %v", err)
	}
	if respostaCriar.Dados.ID == "" {
		t.Fatalf("esperava id preenchido na resposta")
	}

	requisicaoListar := httptest.NewRequest(http.MethodGet, "/tarefas", nil)
	gravadorListar := httptest.NewRecorder()
	roteador.ServeHTTP(gravadorListar, requisicaoListar)

	if gravadorListar.Code != http.StatusOK {
		t.Fatalf("esperava status %d, recebeu %d", http.StatusOK, gravadorListar.Code)
	}

	var respostaLista envelopeLista
	if err := json.Unmarshal(gravadorListar.Body.Bytes(), &respostaLista); err != nil {
		t.Fatalf("erro ao decodificar lista: %v", err)
	}
	if len(respostaLista.Dados) != 1 {
		t.Fatalf("esperava 1 tarefa na lista, recebeu %d", len(respostaLista.Dados))
	}

	idTarefa := respostaCriar.Dados.ID

	requisicaoAtualizar := httptest.NewRequest(http.MethodPut, "/tarefas/"+idTarefa, strings.NewReader(`{"titulo":"Ler", "descricao":"Estudar testes", "concluida":true}`))
	requisicaoAtualizar.Header.Set("Content-Type", "application/json")
	gravadorAtualizar := httptest.NewRecorder()
	roteador.ServeHTTP(gravadorAtualizar, requisicaoAtualizar)

	if gravadorAtualizar.Code != http.StatusOK {
		t.Fatalf("esperava status %d na atualizacao, recebeu %d", http.StatusOK, gravadorAtualizar.Code)
	}

	var respostaAtualizada envelopeTarefa
	if err := json.Unmarshal(gravadorAtualizar.Body.Bytes(), &respostaAtualizada); err != nil {
		t.Fatalf("erro ao decodificar resposta de atualizacao: %v", err)
	}
	if !respostaAtualizada.Dados.Concluida {
		t.Fatalf("esperava tarefa marcada como concluida")
	}

	requisicaoRemover := httptest.NewRequest(http.MethodDelete, "/tarefas/"+idTarefa, nil)
	gravadorRemover := httptest.NewRecorder()
	roteador.ServeHTTP(gravadorRemover, requisicaoRemover)

	if gravadorRemover.Code != http.StatusOK {
		t.Fatalf("esperava status %d na remocao, recebeu %d", http.StatusOK, gravadorRemover.Code)
	}

	var respostaRemocao envelopePadrao
	if err := json.Unmarshal(gravadorRemover.Body.Bytes(), &respostaRemocao); err != nil {
		t.Fatalf("erro ao decodificar resposta de remocao: %v", err)
	}
	if respostaRemocao.Dados["idRemovido"] != idTarefa {
		t.Fatalf("esperava id removido %s, recebeu %v", idTarefa, respostaRemocao.Dados["idRemovido"])
	}

	requisicaoBuscar := httptest.NewRequest(http.MethodGet, "/tarefas/"+idTarefa, nil)
	gravadorBuscar := httptest.NewRecorder()
	roteador.ServeHTTP(gravadorBuscar, requisicaoBuscar)

	if gravadorBuscar.Code != http.StatusNotFound {
		t.Fatalf("esperava status %d ao buscar tarefa removida, recebeu %d", http.StatusNotFound, gravadorBuscar.Code)
	}
}
