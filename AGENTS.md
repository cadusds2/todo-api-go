# AGENTS.md

## Contexto
- Este repositório serve como ambiente de estudo para aprender Go construindo uma API de to-do list.
- O foco principal é praticar boas práticas idiomáticas da linguagem e registrar aprendizados.

## Fluxo de trabalho
- Criar branches nomeadas seguindo o padrão `feature/nome-descritivo`, `bugfix/nome-descritivo` ou `hotfix/nome-descritivo`.
- Manter mensagens de commit curtas em português no formato `tipo: descrição` (ex.: `feat: adicionar endpoint de criação de tarefa`).
- Sempre executar `gofmt ./...` e `go test ./...` antes de abrir um PR.
- Ao preparar um PR, descrever claramente o que foi feito e relacionar tarefas concluídas em uma única checklist quando possível.

## Estilo e organização
- Estruturar o código conforme práticas comuns de Go, usando pastas como `cmd/`, `internal/` e `pkg/` quando necessário.
- Centralizar handlers HTTP em `internal/api`, regras de negócio em `internal/core` (ou nome equivalente) e dependências compartilhadas em `pkg/`.
- Nomear variáveis, funções, arquivos e comentários em português para reforçar o aprendizado.
- Utilizar `context.Context` em operações de I/O e em handlers HTTP para permitir cancelamento e deadlines.
- Preferir validações explícitas e tratamento de erros detalhado, retornando mensagens claras.

## Testes
- Garantir que todo novo comportamento possua testes (`*_test.go`), priorizando testes orientados a tabelas.
- Escrever testes unitários para cada camada e, quando aplicável, adicionar testes de integração ou de contrato.
- Mirar cobertura mínima de 80% e documentar como executar testes adicionais se forem necessários (ex.: uso de banco de dados ou mocks).

## Checklist antes do PR
- `gofmt ./...`
- `go test ./...`
- Executar `golangci-lint run` (se configurado no projeto) e corrigir alertas.
- Atualizar documentação relevante (ex.: `README.md`, exemplos de requisições, coleções de testes) sempre que endpoints mudarem.
- Verificar se exemplos de uso ou scripts auxiliares continuam válidos.

## Recursos úteis
- [Tour of Go](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Documentação de net/http](https://pkg.go.dev/net/http)
- [Guia de módulos Go](https://go.dev/doc/modules)
- Materiais sobre Clean Architecture em Go e padrões de manipulação de erros para inspirar a organização do código.
