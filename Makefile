# Ensure reflex is installed
ensure-reflex:
	@if ! command -v reflex >/dev/null 2>&1; then \
		echo "🔍 Reflex not found. Installing..."; \
		go install github.com/cespare/reflex@latest; \
		echo "✅ Reflex installed."; \
	fi

# Ensure goose is installed
ensure-goose:
	@if ! command -v goose >/dev/null 2>&1; then \
		echo "🔍 Goose not found. Installing..."; \
		go install github.com/pressly/goose/v3/cmd/goose@latest; \
		echo "✅ Goose installed."; \
	fi

# Ensure swag is installed
ensure-swagger:
	@if ! command -v swag >/dev/null 2>&1; then \
		echo "🔍 Swag not found. Installing..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
		echo "✅ Swag installed."; \
	fi

# Ensure mockgen is installed
ensure-mockgen:
	@if ! command -v mockgen >/dev/null 2>&1; then \
		echo "🔍 Mockgen not found. Installing..."; \
		go install github.com/golang/mock/mockgen@v1.6.0; \
		echo "✅ Mockgen installed."; \
	fi

# Ensure gomock dependency
ensure-gomock: 
	@go get -d github.com/golang/mock/gomock@v1.6.0 >/dev/null 2>&1 || true

# Run go mod tidy
tidy:
	go mod tidy

# Generate mocks
mock: ensure-mockgen ensure-gomock
	go generate ./...

# Run unit test
unit-test:
	go test -cover $$(go list ./... | grep -v -E '/(mocks|vendor|docs)')

# Run test suite
test: mock unit-test

# Generate Swagger docs
api-docs: ensure-swagger
	swag init -g server/http_router.go --parseDependency true --parseInternal true

# Run migration up
migrate-up: ensure-goose
	goose -dir ./migrations -table migration_version postgres "postgres://attendanceuser:attendancepassword@attendance-postgresql:5432/attendance_db?sslmode=disable" up

# Run migration down
migrate-down: ensure-goose
	goose -dir ./migrations -table migration_version postgres "postgres://attendanceuser:attendancepassword@attendance-postgresql:5432/attendance_db?sslmode=disable" down

# Pre-run setup
pre-run: tidy mock unit-test api-docs

# Run HTTP server with reflex
run-http: ensure-reflex
	@echo "🚀 Running HTTP server..."
	@reflex -r '\.go$$' -s -- sh -c "go run main.go serve-http"

# Run worker with reflex
run-worker: ensure-reflex
	@echo "👷 Running worker..."
	@reflex -r '\.go$$' -s -- sh -c "go run main.go serve-worker"

# Run all services
run-all: pre-run migrate-up
	@echo "👷 Run http and worker 🚀"
	@reflex -r '\.go$$' -s -- sh -c "go run main.go serve-all"
