# review-service

Product review microservice for ratings and comments.

## Features

- Create product reviews
- Get reviews by product
- Rating aggregation
- Duplicate prevention

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/api/v1/reviews?product_id={id}` | Get reviews |
| `POST` | `/api/v1/reviews` | Create review |

## Tech Stack

- Go + Gin framework
- PostgreSQL 16 (review-db cluster, single instance)
- Direct connection (no pooler)
- OpenTelemetry tracing

## Development

```bash
go mod download
go test ./...
go run cmd/main.go
```

## License

MIT
