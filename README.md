Módulo 1 — Fundamentos do Go ✅

Objetivo: aprender a linguagem.

* ✅ Packages
* ✅ Imports
* ✅ main
* ✅ Variáveis
* ✅ Constantes
* ✅ Structs
* ✅ Métodos
* ✅ Receivers
* ✅ Ponteiros
* ✅ Interfaces
* ✅ Errors
* ✅ Organização de packages

⸻

Módulo 2 — HTTP com a stdlib ✅

Objetivo: entender o net/http.

* ✅ http.ListenAndServe
* ✅ http.Handle
* ✅ http.HandleFunc
* ✅ http.Handler
* ✅ http.HandlerFunc
* ✅ Rotas
* ✅ Path Parameters
* ✅ Status Codes
* ✅ Headers (parcial)
* ✅ Request
* ✅ ResponseWriter

⸻

Módulo 3 — JSON ✅

* ✅ encoding/json
* ✅ json.Decoder
* ✅ json.Encoder
* ✅ Struct Tags
* ✅ Streaming de JSON
* ✅ Decode
* ✅ Encode

⸻

Módulo 4 — Organização da aplicação ✅

* ✅ internal/
* ✅ Separação por domínio
* ✅ Handler
* ✅ Store
* ✅ MemoryStore
* ✅ PostgresStore
* ✅ Composition Root
* ✅ Dependency Injection manual
* ✅ Constructors (NewXxx)
* ✅ Interfaces orientadas por comportamento

⸻

Módulo 5 — Context ✅

* ✅ context.Context
* ✅ r.Context()
* ✅ context.Background()
* ✅ context.WithTimeout
* ✅ Propagação do Context
* ✅ Cancelamento
* ✅ Deadlines

⸻

Módulo 6 — Concorrência (introdução) ✅

Ainda sem goroutines.

Mas já aprendemos:

* ✅ sync.RWMutex
* ✅ Lock
* ✅ Unlock
* ✅ RLock
* ✅ RUnlock
* ✅ Estado compartilhado
* ✅ Thread safety

Mais tarde estudaremos concorrência de verdade.

⸻

Módulo 7 — Banco de Dados (database/sql) ✅

Até agora:

* ✅ Driver PostgreSQL
* ✅ Blank Import
* ✅ sql.DB
* ✅ Pool de conexões
* ✅ sql.Open
* ✅ PingContext
* ✅ Configuração do pool
* ✅ QueryContext
* ✅ QueryRowContext
* ✅ ExecContext
* ✅ Rows
* ✅ Row
* ✅ Scan
* ✅ Placeholders ($1)
* ✅ RETURNING
* ✅ rows.Close
* ✅ rows.Err

⸻

Módulo 8 — Erros ✅

* ✅ errors.New
* ✅ errors.Is
* ✅ errors.As
* ✅ Erros de domínio
* ✅ ValidationError
* ✅ Separação entre erro técnico e erro de domínio

⸻

Módulo 9 — Validação ✅

* ✅ Validate()
* ✅ Validação sem bibliotecas
* ✅ strings.TrimSpace
* ✅ strings.Contains

⸻

Módulo 10 — Middleware (em andamento)

* ✅ http.Handler
* ✅ ServeHTTP
* ✅ Encadeamento de handlers

Ainda falta bastante.

⸻

Até aqui…

Já aprendemos aproximadamente:

Packages da stdlib

* ✅ net/http
* ✅ encoding/json
* ✅ database/sql
* ✅ context
* ✅ sync
* ✅ errors
* ✅ strings
* ✅ strconv
* ✅ time
* ✅ log

⸻

O que ainda vamos aprender

Agora começa a parte “profissional”.

⸻

Módulo 11 — Middleware (continuação)

* Logging
* Recovery
* Request ID
* CORS
* Timeout
* Compressão
* Rate Limiting
* Chain de middlewares

-----

Módulo 12 — Logging

* log
* slog
* Logger estruturado
* Levels
* Attributes
* Context
* Request ID
* JSON logs

Módulo 13 — Testes

Quero dedicar um módulo inteiro.

* testing
* httptest
* Table Driven Tests
* Subtests
* Benchmarks
* Fuzzing
* Coverage
* Integração com PostgreSQL
* Fake Store
* Test Helpers

Na minha opinião, será um dos módulos mais importantes.

⸻

Módulo 14 — Transações

Um dos assuntos mais importantes do backend.

Vamos aprender:

* sql.Tx
* BeginTx
* Commit
* Rollback
* defer Rollback
* Nested transactions (conceito)
* Isolation Levels
* Deadlocks

⸻

Módulo 15 — Prepared Statements

* PrepareContext
* Statement Cache
* Performance

⸻

Módulo 16 — Migrações

* golang-migrate
* Organização
* Rollback
* Seed

⸻

Módulo 17 — Configuração

Hoje temos:
postgres://...

Depois teremos:

* .env
* Variáveis de ambiente
* Config struct
* Inicialização

-----

Módulo 18 — Observabilidade

* log/slog
* pprof
* Metrics
* Health Check
* Readiness
* Liveness

⸻

Módulo 19 — Concorrência

Agora sim.

Na minha opinião este será outro grande módulo.

Vamos aprender:

* Goroutines
* Channels
* WaitGroup
* Context + Goroutines
* Worker Pool
* Fan In
* Fan Out
* Pipeline
* Select
* Mutex (mais profundo)
* Atomic
* Cond
* Once

Tudo aplicado na API.

⸻

Módulo 20 — Cache

* Redis
* Cache Aside
* TTL
* Invalidation

⸻

Módulo 21 — Autenticação

* JWT
* Refresh Token
* Middleware
* Hash de senha (crypto/bcrypt)
* Cookies
* Sessions

⸻

Módulo 22 — Uploads

* Multipart
* Arquivos
* Streaming
* Limites

⸻

Módulo 23 — RabbitMQ

Porque você já trabalha bastante com mensageria.

Vamos usar:

* Publisher
* Consumer
* Retry
* DLQ
* Context
* Graceful Shutdown

⸻

Módulo 24 — Graceful Shutdown

Outro assunto extremamente importante.

Aprenderemos:

* Signals
* os/signal
* context.WithCancel
* Fechar HTTP
* Fechar DB
* Fechar RabbitMQ

⸻

Módulo 25 — Performance

* pprof
* Benchmark
* Escape Analysis
* Heap
* Stack
* GC
* Allocation
* sync.Pool

⸻

Módulo 26 — Projeto Final

Aí sim.

Transformaremos esse CRUD em algo realmente interessante.

Minha ideia seria um backend com:

* Login
* JWT
* Produtos
* Usuários
* Categorias
* Upload
* RabbitMQ
* Cache
* PostgreSQL
* Docker
* Migrations
* Testes
* Observabilidade
* Graceful Shutdown

## E depois?

A partir daí eu partiria para um segundo curso.

Go Avançado

* Reflection
* Generics
* Unsafe
* Runtime
* Scheduler
* Garbage Collector
* Escape Analysis profunda
* Compiler
* Interfaces internamente
* Memory Layout
* Race Detector
* Profiling avançado
