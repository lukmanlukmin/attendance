# 🏦 Attendance Service

A scalable backend service for managing employee attendance, overtime, reimbursement, and payroll generation. Built with **Golang**, **PostgreSQL**, and **Kafka**, the system is designed with clean architecture and auditability in mind.

---

## 🚀 Tech Stack

- **Golang** 1.23+
- **Fiber** – HTTP framework
- **PostgreSQL** – database
- **sqlx** + **Squirrel** – query builder
- **Kafka** – async event handling
- **Docker Compose** – development environment
- **Testify** + **GoMock** – testing

---

## 📁 Project Structure

```
loan-service/
├── bootstrap/          # register all module
├── cmd/                # command Manager
├── config/             # configuration loader
├── constant/           # global constants
├── docs/               # swagger docs
├── handler/            # HTTP handlers (Fiber)
├── middleware/         # middleware functions
├── migrations/         # migration files
├── model/
│   ├── db/             # Database models
│   ├── event/          # Event models
│   └── payload/        # HTTP request/response payloads
├── repository/         # Data access layer using sqlx + squirrel
├── server/             # server setup
├── service/            # Business logic layer
├── utils/              # utility package
├── main.go             # Application entry point
├── Dockerfile.go.1.23
├── config.yaml
├── docker-compose.yml
└── README.md
```

---

## 🧩 Features Summary

- ⏱️ **Attendance Submission**  
  Employees can check in once per weekday. Weekends are disabled.

- 🕒 **Overtime Proposal**  
  Limited to 3 hours/day after working hours. Paid 2× normal hourly rate.

- 💸 **Reimbursement Requests**  
  With amount and description, included in payslip.

- 💼 **Payroll Execution by Admin**  
  Locked after processed. One-time only per period.

- 🧾 **Employee Payslip Generation**  
  Includes attendance, overtime, reimbursements, and total take-home pay.

- 📊 **Admin Payroll Summary**  
  Displays each employee’s take-home pay and total payroll.

---

## ⚙️ Makefile Commands

This project includes a `Makefile` for simplified development. Key commands include:

### 🧰 Dependency Installer
Automatically installs required CLI tools if not present:
```bash
make ensure-reflex     # Hot reload Go server during development
make ensure-goose      # Database migration tool
make ensure-swagger    # Swagger documentation generator
make ensure-mockgen    # GoMock interface generator
```

### 🔍 Testing & Mocking
```bash
make tidy              # Run go mod tidy
make mock              # Generate mocks from interfaces
make unit-test         # Run unit tests with coverage
make test              # Run mocks + tests together
```

### 🚀 Running the Server
```bash
make run-http          # Run HTTP server with hot reload
make run-worker        # Run Kafka worker with hot reload
make run-all           # Run both HTTP and worker
```

### 📄 API Documentation (Swagger)
```bash
make api-docs
```
Generates Swagger documentation in the `docs/` folder. Access via:

```
http://localhost:8081/swagger/index.html#/
```

### 🔧 Database Migration
```bash
make migrate-up        # Run all up migrations
make migrate-down      # Rollback last migration
```

To create a new migration:
```bash
goose -dir ./migrations create add_column_name sql
```

---

## 🐳 Dockerized Setup

This project is **preconfigured** in both `docker-compose.yml` and `config.yaml`. You don't need to modify anything to get started.

Run the full stack with:

```bash
docker-compose up --build
```

Accessible services:

- **Swagger API Docs**:  
  [http://localhost:8081/swagger/index.html#/](http://localhost:8081/swagger/index.html#/)

- **Kafka UI**:  
  [http://localhost:8082/](http://localhost:8082/)

## 🧪 Testing Strategy

- ✅ **Partial unit tests** implemented for key modules (handler, service, repository).
- ✅ Uses **GoMock** for mocking dependencies.
- 🔧 More coverage planned, especially for edge cases and full integration flow.
- 🚀 Easily runnable via:
  ```bash
  make test
  ```

---

## 🧠 Architecture Highlights

- 📦 Clean and modular: **Handler → Service → Repository**.
- ⚙️ Uses **Kafka** to handle payroll execution asynchronously and scalably.
- 🔄 Designed to be **idempotent** and safe for concurrent processing.
- 🧵 Uses **context propagation** to pass `request_id`, `user_id`, and IP address for tracing.

---

## 🔍 Auditability & Tracing

- ✅ Includes fields: `created_at`, `updated_at`, `created_by`, `updated_by`, `created_ip`, `updated_ip`.
- ✅ Middleware captures and logs **`request_id`** for end-to-end request tracing.
- ⚠️ Audit log table for critical record tracking is **planned but not yet implemented**.

---

## 📈 Scalability Goals

- ✅ Asynchronous **Kafka-based worker** architecture to support horizontal scaling.
- ✅ Modular, **testable components** for potential microservice extraction.
- 🛠️ Open for enhancements like **sharding**, **batch processing**, or queue backpressure handling.

---

## 🧊 Status

🚧 Actively developed  
✅ Core features implemented  
📌 Remaining work includes:  
- Increasing unit test coverage  
- Completing audit logging system  
- Improving API rate limiting and validation

---