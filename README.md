## <div align="center"> 📋 Kanban Board </div>

<div align="center">
A complete task management application built with the GoFr framework. Create boards, organize tasks into lists, collaborate with team members, and track project progress with an intuitive Kanban-style interface.
</div>

<br>

## 🚀 Getting Started

### Prerequisites
- Go 1.21+
- PostgreSQL 12+
- Git

### 🛠️ Setup Instructions

1. **Clone the repository**
   ```bash
   git clone https://github.com/AvishkarPatil/GoFr-Kanban-Board.git
   cd GoFr-Kanban-Board
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Setup PostgreSQL database**
   ```sql
   CREATE DATABASE kanban_board;
   ```

4. **Configure environment**
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

5. **Run migrations**
   ```bash
   # GoFr will auto-run migrations on startup
   ```

6. **Start the server**
   ```bash
   make run
   # or
   go run cmd/server/main.go
   ```

### 🧪 Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run linter
make lint
```

### 📊 Database Schema

The application uses 5 core tables:
- `users` - User authentication and profiles
- `boards` - Project boards
- `lists` - Board columns (Todo, Doing, Done)
- `cards` - Tasks within lists
- `board_members` - User collaboration on boards

### 🔗 API Endpoints

#### Health Check
- `GET /health` - Service health status

### 📜 Development Status (PR1)

#### ✅ Completed
- [x] Project structure setup
- [x] Database schema design
- [x] Migration files
- [x] Environment configuration
- [x] Health check endpoint
- [x] Unit tests for database operations
- [x] Linting configuration
- [x] Development tooling (Makefile)

#### 🔄 Next (PR2)
- [ ] Authentication & Users
- [ ] Board Management
- [ ] Lists & Cards
- [ ] JWT Middleware
- [ ] API Documentation
