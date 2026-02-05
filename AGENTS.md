# review-service

> AI Agent context for understanding this repository

## 📋 Overview

Product review microservice. Manages product reviews and ratings.

## 🏗️ Architecture

```
review-service/
├── cmd/main.go
├── config/config.go
├── db/migrations/sql/
├── internal/
│   ├── core/
│   │   ├── database.go
│   │   └── domain/
│   ├── logic/v1/service.go
│   └── web/v1/handler.go
├── middleware/
└── Dockerfile
```

## 🔌 API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/api/v1/reviews?product_id={id}` | Get reviews for product |
| `POST` | `/api/v1/reviews` | Create review (409 if duplicate) |

## 📐 3-Layer Architecture

| Layer | Location | Responsibility |
|-------|----------|----------------|
| **Web** | `internal/web/v1/handler.go` | HTTP, validation, error translation |
| **Logic** | `internal/logic/v1/service.go` | Business rules (❌ NO SQL) |
| **Core** | `internal/core/` | Domain models, repositories |

## 🗄️ Database

| Component | Value |
|-----------|-------|
| **Cluster** | review-db (Zalando Postgres Operator) |
| **PostgreSQL** | 16 |
| **HA** | Single instance (no HA) |
| **Pooler** | **None** (direct connection) |
| **Endpoint** | `review-db.review.svc.cluster.local:5432` |
| **Driver** | pgx/v5 |

**Why no pooler?**
- Low traffic service
- No connection pooler overhead
- Direct PostgreSQL connection is sufficient

## 🚀 Graceful Shutdown

**VictoriaMetrics Pattern:**
1. `/ready` → 503 when shutting down
2. Drain delay (5s)
3. Sequential: HTTP → Database → Tracer

## 🔧 Tech Stack

| Component | Technology |
|-----------|------------|
| **Framework** | Gin |
| **Database** | PostgreSQL 16 via pgx/v5 |
| **Tracing** | OpenTelemetry |

## 🛠️ Development

```bash
go mod download && go test ./... && go build ./cmd/main.go
```
