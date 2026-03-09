# Project Explanation: Go Project Structure Template

This project is a comprehensive **Go Project Structure Template** designed with **Domain-Driven Design (DDD)**, **Clean Architecture**, and **Hexagonal Architecture** principles. It serves as a robust starting point for building scalable and maintainable microservices, specifically exemplified by a CRM-like API for managing customer profiles.

---

## 🏗 Architectural Patterns

The project follows a layered architecture to ensure separation of concerns:

1.  **Domain Layer (`internal/domain`)**: The core of the application. It contains business logic, entities (`CustomerProfile`), value objects (`Email`, `DateOfBirth`), and repository interfaces. It has no dependencies on other layers.
2.  **Application Layer (`internal/application`)**: Orchestrates the business logic. It implements the **CQRS-lite** pattern using Command/Query handlers. It also defines the **Unit of Work** interface to manage transactions.
3.  **Infrastructure Layer (`internal/infrastructure`)**: Contains technical implementations.
    -   **HTTP**: Uses the **Gin** framework for routing and middleware.
    -   **Persistence**: Supports both **PostgreSQL** (via `sqlx`) and **MongoDB**.
    -   **Observability**: Integrated with **OpenTelemetry (OTel)** for distributed tracing, metrics, and structured logging.
4.  **CMD Layer (`cmd/`)**: Entry points for the CLI application using **Cobra**.

---

## 📂 Directory Structure

-   `cmd/`: CLI commands (`serve`, `migrate`).
-   `config/`: Configuration management using **Viper** (`config.yaml`).
-   `daemon/`: Bootstrapping logic to initialize the server, databases, and observability.
-   `internal/`: Private application code.
    -   `application/`: Command handlers, DTOs, and Unit of Work interfaces.
    -   `core/`: Shared domain errors and events.
    -   `domain/`: Business entities and logic.
    -   `infrastructure/`: Database, HTTP, and Observability implementations.
-   `migrate/`: Database migration logic and scripts (SQL and NoSQL).
-   `otel/`: Configuration for OpenTelemetry components (Alloy, Mimir, Tempo, Grafana).
-   `pkg/`: Shared utility packages (e.g., `array_utils`, `result`).

---

## 🛠 Technology Stack

-   **Language**: Go 1.25.2
-   **HTTP Framework**: [Gin Gonic](https://github.com/gin-gonic/gin)
-   **CLI Framework**: [Cobra](https://github.com/spf13/cobra)
-   **Configuration**: [Viper](https://github.com/spf13/viper)
-   **Databases**: PostgreSQL and MongoDB
-   **Migrations**: [Golang Migrate](https://github.com/golang-migrate/migrate)
-   **Observability**: [OpenTelemetry](https://opentelemetry.io/) (OTel)
-   **Testing**: [Testify](https://github.com/stretchr/testify), `sqlmock`

---

## 🚀 Key Features

### 1. Robust Observability
The project is pre-configured with a full observability stack:
-   **Tracing**: Exported to OTLP (e.g., Tempo).
-   **Metrics**: Custom metrics for HTTP requests and business operations (e.g., Mimir).
-   **Logging**: Structured logging via `slog` with OTel integration.

### 2. Multi-Database Support
It provides implementations for both SQL (Postgres) and NoSQL (MongoDB), demonstrating how to swap persistence layers while keeping the domain intact.

### 3. Graceful Shutdown
The server handles `SIGINT` and `SIGTERM` signals to shut down gracefully, ensuring all connections (DB, OTel) are closed properly.

### 4. Unit of Work Pattern
Ensures that multiple repository operations can be treated as a single transaction, maintaining data integrity.

### 5. Load Testing
Includes a `load_test.js` for **k6**, pre-configured to simulate traffic with a controlled error rate (10%) to test observability and error handling under load.

---

## 🚦 How to Run

### Prerequisite
- Go 1.25+
- Docker & Docker-Compose (for databases and observability stack)

### Steps
1.  **Start Infrastructure**:
    ```bash
    docker-compose up -d
    ```
2.  **Run Migrations**:
    ```bash
    go run main.go migrate sql up
    # or for mongo
    go run main.go migrate mongo up
    ```
3.  **Start the Server**:
    ```bash
    go run main.go serve
    ```

The API will be available at `http://localhost:8080/crm-api/v1`.
Check health at `GET /crm-api/v1/health`.
