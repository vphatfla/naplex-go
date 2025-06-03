myapp/
├── cmd/
│   ├── api/
│   │   └── main.go              # API server entry point
│
├── internal/                    # Private application code
│   ├── user/                    # User feature/domain
│   │   ├── handler.go          # HTTP handlers
│   │   ├── service.go          # Business logic
│   │   ├── repository.go       # Data access interface
│   │   ├── repository_pg.go    # PostgreSQL implementation
│   │   ├── model.go            # Domain models
│   │   ├── dto.go              # Data transfer objects
│   │   ├── validation.go       # Input validation
│   │   ├── user_test.go        # Tests
│   │   ├── queries/            # SQLC queries
│   │   │   ├── user.sql        # User-specific SQL queries
│   │   │   └── sqlc.yaml       # SQLC config for this feature
│   │   └── db/                 # SQLC generated code
│   │       ├── user.sql.go     # Generated query functions
│   │       ├── models.go       # Generated models
│   │       └── db.go           # Generated interface
│   │
│   ├── auth/                    # Authentication feature
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   ├── repository_pg.go
│   │   ├── jwt.go              # JWT token handling
│   │   ├── middleware.go       # Auth middleware
│   │   ├── model.go
│   │   ├── auth_test.go
│   │   ├── queries/
│   │   │   └── auth.sql
│   │   └── db/
│   │       └── auth.sql.go
│   │
│   └── shared/                  # Shared internal packages
│       ├── database/
│       │   ├── postgres.go     # Database connection
│       │   └── transaction.go  # Transaction handling
│       ├── cache/
│       │   └── redis.go        # Redis client
│
│
├── api/                         # API definitions
│   ├── openapi/
│   │   └── swagger.yaml        # OpenAPI specification
│   └── grpc/
│       └── service.proto       # gRPC service definitions
│
├── web/                         # Web assets (if applicable)
│   ├── templates/
│   └── static/
│
├── migrations/                  # Database migrations (used by SQLC for schema)
│   ├── 001_create_users.up.sql
│   ├── 001_create_users.down.sql
│   ├── 002_create_products.up.sql
│   └── 002_create_products.down.sql
│
├── scripts/                     # Build and utility scripts
│   ├── build.sh
│   └── test.sh
│
├── deployments/                 # Deployment configurations
│   ├── docker/
│   │   ├── Dockerfile
│   │   └── docker-compose.yml
│   ├── kubernetes/
│   │   ├── deployment.yaml
│   │   └── service.yaml
│   └── terraform/
│       └── main.tf
│
├── configs/                     # Configuration files
│   ├── dev.yaml
│   ├── staging.yaml
│   └── prod.yaml
│
├── test/                        # Integration tests
│   ├── integration/
│   │   ├── user_test.go
│   │   └── order_test.go
│   └── e2e/
│       └── api_test.go
│
├── docs/                        # Documentation
│   ├── README.md
│   ├── architecture.md
│   └── api.md
│
├── .github/                     # GitHub specific files
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
│
├── go.mod                       # Go module file
├── go.sum
├── .gitignore
├── .golangci.yml               # Linter configuration
├── Makefile                    # Build automation
└── README.md
