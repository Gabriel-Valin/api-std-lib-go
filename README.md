# Course Roadmap

## Module 1 — Go Fundamentals
- Packages and project structure
- Variables and constants
- Functions
- Structs
- Methods and receivers
- Pointers
- Interfaces
- Error handling
- Dependency injection fundamentals

---

## Module 2 — Building HTTP APIs with the Standard Library
- net/http
- http.Handler
- http.HandlerFunc
- ServeMux
- Request and ResponseWriter
- Routing
- HTTP status codes
- Headers
- REST API fundamentals

---

## Module 3 — Working with JSON
- encoding/json
- JSON decoding
- JSON encoding
- Struct tags
- Streaming JSON
- Request validation

---

## Module 4 — Application Architecture
- Layered architecture
- Handlers
- Services
- Stores (Repositories)
- In-memory implementation
- PostgreSQL implementation
- Composition Root
- Dependency Injection
- Project structure

---

## Module 5 — Context
- context.Context
- Request-scoped values
- Deadlines
- Timeouts
- Cancellation
- Context propagation

---

## Module 6 — PostgreSQL with database/sql
- database/sql
- Connection pooling
- Queries
- Inserts
- Updates
- Deletes
- Scanning rows
- Prepared statements
- Error handling
- CRUD implementation

---

## Module 7 — Configuration Management
- Configuration architecture
- Typed configuration
- Environment variables
- .env files
- Validation
- Environment-specific configuration
- Application bootstrap

---

## Module 8 — HTTP Middleware
- Middleware fundamentals
- Middleware chaining
- Logging middleware
- Recovery middleware
- Timing middleware
- Request ID middleware
- Custom ResponseWriter
- Request pipeline

---

## Module 9 — Structured Logging
- log/slog
- Structured logging
- Context-aware logging
- Stack traces
- Request correlation
- Production-ready logging

---

## Module 10 — Graceful Shutdown
- http.Server
- Signal handling
- Graceful shutdown
- Resource cleanup
- Connection draining
- Production lifecycle

---

## Module 11 — Testing
- Unit testing
- Table-driven tests
- httptest
- Integration tests
- Fake implementations
- Test helpers
- Benchmarks
- Fuzz testing
- Race detector
- Coverage

---

## Module 12 — Database Transactions
- sql.Tx
- Transactions
- Commit
- Rollback
- Isolation levels
- Deadlocks
- Best practices

---

## Module 13 — Database Migrations
- golang-migrate
- Versioned migrations
- Rollbacks
- Seed data
- Migration workflow

---

## Module 14 — Observability
- Health checks
- Readiness probes
- Liveness probes
- pprof
- Metrics
- Performance profiling
- Application diagnostics

---

## Module 15 — Concurrency
- Goroutines
- Channels
- Select
- WaitGroup
- Mutexes
- RWMutex
- Atomic operations
- Worker pools
- Pipelines
- Fan-in / Fan-out
- Context with goroutines
- sync.Once
- sync.Cond
- sync.Map

---

## Module 16 — Caching
- Redis
- Cache-aside pattern
- TTL
- Cache invalidation
- Performance optimization

---

## Module 17 — Messaging
- RabbitMQ
- Publishers
- Consumers
- Retries
- Dead Letter Queues
- Idempotency
- Outbox pattern
- Graceful consumer shutdown

---

## Module 18 — Authentication & Authorization
- Password hashing
- JWT
- Refresh tokens
- Authentication middleware
- Authorization
- Role-Based Access Control (RBAC)

---

## Module 19 — File Uploads
- Multipart forms
- Streaming uploads
- File validation
- Local storage
- Cloud storage concepts

---

## Module 20 — Final Project
Build a production-ready REST API featuring:
- Authentication
- Users
- Products
- Categories
- PostgreSQL
- Redis
- RabbitMQ
- File uploads
- Migrations
- Testing
- Logging
- Observability
- Graceful shutdown
- Docker deployment

---

## Module 21 — Advanced Go (Bonus)
- Runtime internals
- Scheduler
- Garbage Collector
- Escape analysis
- Memory layout
- Reflection
- Generics
- Unsafe
- Interface internals
- Compiler optimizations

---

## Module 22 — Performance Engineering (Bonus)
- Benchmark-driven optimization
- CPU profiling
- Memory profiling
- Allocation analysis
- sync.Pool
- Performance tuning
- Real-world optimization techniques
