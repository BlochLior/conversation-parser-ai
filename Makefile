# Project-root Makefile - Python & Go services

# Prevent caching issues and ensure all targets run fresh every time
.PHONY: health lint-python format-python test-python build-python run-docker-python build-go run-docker-go ci-python test-go logs gosec lint-go ci-go ci-all

# ✅ Ping /health endpoint to confirm service is running
health:
	./python-ai/health_check.sh

# 🧼 Lint Python code using ruff (report only)
lint-python:
	ruff check python-ai

# 🛠 Auto-fix style issues using ruff (safe fixes only)
format-python:
	ruff check --fix python-ai

# 🧪 Run Python unittests
test-python:
	PYTHONPATH=python-ai python3 -m unittest discover -s python-ai/tests -p "test_*.py"

# 🐳 Build Docker image for Python AI service
build-python:
	docker build -t python-ai-service ./python-ai

# 🐳 Run Python AI service in Docker
run-docker-python:
	docker run -p 8001:8001 --env-file=./python-ai/.env -e ENV=production python-ai-service

# 🐳 Build Docker image for Go backend
build-go:
	docker build -t go-backend-service ./go-backend

# 🐳 Run Go backend in Docker
run-docker-go:
	docker run -p 8000:8000 go-backend-service

# 🔁 Build both services
build-all: build-python build-go

# 📄 View recent logs from the Python AI service container
logs:
	docker ps -q --filter ancestor=python-ai-service | xargs -r docker logs --tail=50

# ✅ CI target — run lint, test, and health checks in sequence
ci-python: lint-python test-python health

# Run go tests
test-go:
	cd go-backend && go test ./...

# Go security check
gosec:
	cd go-backend && gosec ./...

# Go lint check
lint-go:
	cd go-backend && golangci-lint run ./...

# 🚦 Go CI target for lint + security
ci-go: lint-go gosec test-go

# 🚦 Full CI target for both services
ci-all: ci-python ci-go

# 🧪 Run both Go and Python services via Compose
dev:
	docker-compose up --build

# 🛑 Stop and clean up all Compose containers
dev-down:
	docker-compose down --remove-orphans --volumes

# 🔁 Rebuild and restart services with Docker Compose
rebuild:
	docker-compose down --remove-orphans --volumes && docker-compose up --build

# 📋 Show running containers in this Compose project
ps:
	docker-compose ps
