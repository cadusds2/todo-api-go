# todo-api-go

## Contexto do projeto
Este repositório hospeda uma API de lista de tarefas escrita em Go, planejada para demonstrar boas práticas de arquitetura modular (com os diretórios `cmd/`, `internal/` e `pkg/`) e integrações futuras com ferramentas de observabilidade, autenticação e persistência de dados. A proposta é oferecer uma base didática que facilite a evolução incremental da aplicação, mantendo simplicidade para execução local e testes automatizados.

## Pré-requisitos
- Go 1.22 ou superior instalado e disponível no `PATH`.
- Ferramentas auxiliares sugeridas:
  - `golangci-lint` para linting.
  - `air` ou `reflex` para hot reload durante o desenvolvimento.
  - `make` (opcional) para padronizar execuções futuras via tarefas automatizadas.

## Instruções de instalação e execução locais
1. Clone o repositório:
   ```bash
   git clone https://github.com/seu-usuario/todo-api-go.git
   cd todo-api-go
   ```
2. Baixe as dependências do projeto:
   ```bash
   go mod tidy
   ```
3. Execute a aplicação localmente:
   ```bash
   go run ./cmd/api
   ```
4. Opcional: gere o binário para distribuição local:
   ```bash
   go build -o bin/todo-api ./cmd/api
   ./bin/todo-api
   ```

Consulte o arquivo `AGENTS.md` para orientações complementares sobre estilo de código e práticas de contribuição.

## Estrutura de diretórios prevista
- `cmd/`: pontos de entrada (como `cmd/api/main.go`) que inicializam servidores ou tarefas específicas.
- `internal/`: implementação interna não exportada para outros módulos, contendo casos de uso, serviços, adaptadores e repositórios.
- `pkg/`: bibliotecas reutilizáveis e exportáveis para outros projetos.

## Estratégia de testes
- Execute toda a suíte com:
  ```bash
  go test ./...
  ```
- Inclua testes unitários para casos de uso em `internal/` e testes de integração quando a camada de persistência estiver disponível.
- Utilize `-race` e `-cover` em execuções periódicas para garantir detecção de condições de corrida e visibilidade de cobertura.

## Exemplos de chamadas HTTP planejadas
- Criar uma nova tarefa:
  ```bash
  curl -X POST http://localhost:8080/v1/tarefas \
    -H "Content-Type: application/json" \
    -d '{"titulo":"Comprar café","descricao":"Passar no mercado","concluida":false}'
  ```
- Listar tarefas existentes:
  ```bash
  curl http://localhost:8080/v1/tarefas
  ```

Futuramente, uma coleção de requisições no formato `.http` ou Postman/Insomnia será disponibilizada em `docs/colecoes/` para facilitar o consumo da API.

## Roadmap e próximos passos
- Estruturar os diretórios `cmd/`, `internal/` e `pkg/` com esqueleto inicial.
- Implementar CRUD completo de tarefas com persistência em banco (ex.: PostgreSQL ou SQLite).
- Adicionar validações, paginação e filtros na listagem.
- Configurar observabilidade (logs estruturados, métricas e tracing).
- Publicar pipeline de CI com lint, testes e cobertura automatizados.
- Documentar coleções de requisições e exemplos de payloads adicionais.
