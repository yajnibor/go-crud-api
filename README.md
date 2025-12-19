# Go CRUD API

A simple RESTful CRUD API built with Go, using Gin web framework, PostgreSQL database, and sqlc for type-safe SQL queries.

## Features

- **RESTful API** endpoints for book management
- **Type-safe SQL queries** using sqlc
- **PostgreSQL** database integration with pgx connection pooling
- **Environment-based configuration** using .env files
- **Clean architecture** with separation of concerns

## Technologies

- **Go** 1.25.5
- **Gin** - HTTP web framework
- **PostgreSQL** - Database
- **pgx/v5** - PostgreSQL driver and connection pooling
- **sqlc** - Generate type-safe Go code from SQL
- **godotenv** - Environment variable management

## Project Structure

```
.
├── main.go           # Application entry point
├── go.mod            # Go module dependencies
├── .env              # Environment configuration
├── schema.sql        # Database schema
├── query.sql         # SQL queries
├── sqlc.yaml         # sqlc configuration
└── db/               # Generated database code
    ├── db.go
    ├── models.go
    └── query.sql.go
```

## Database Schema

The application uses a simple `books` table:

```sql
CREATE TABLE books (
    id     BIGSERIAL PRIMARY KEY,
    title  TEXT NOT NULL,
    author TEXT NOT NULL
);
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | `/books` | Get all books |
| POST   | `/books` | Create a new book |

## Setup

### Prerequisites

- Go 1.25.5 or higher
- PostgreSQL database
- sqlc (for code generation)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/yajnibor/go-crud-api.git
cd go-crud-api
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file in the project root:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=go_projects
```

4. Set up the database:
```bash
# Create the database
createdb go_projects

# Run the schema
psql -d go_projects -f schema.sql
```

5. Generate sqlc code (if needed):
```bash
sqlc generate
```

## Running the Application

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## Usage Examples

### Get all books
```bash
curl http://localhost:8080/books
```

### Create a new book
```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"title": "The Go Programming Language", "author": "Alan Donovan"}'
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| DB_HOST | Database host | localhost |
| DB_PORT | Database port | 5432 |
| DB_USER | Database user | postgres |
| DB_PASSWORD | Database password | - |
| DB_NAME | Database name | go_projects |

## Development

### Modifying SQL Queries

1. Edit `query.sql` to add/modify queries
2. Run `sqlc generate` to regenerate Go code
3. Use the generated functions in your handlers

## License

This project is open source and available under the MIT License.
