# Root Makefile — unified build, lint, test, CI for Python & Go services

.PHONY: help build rebuild dev dev-down logs ps \
        test-python ci-python lint-python format-python \
        test-go ci-go lint-go gosec \
        build-python build-go run-docker-python run-docker-go \
        build-all ci-all health-all health-python health-go health-frontend

## 🔎 Help
help:
	@echo "\n📌 Makefile Targets"
	@echo "\n🚀 Local Dev"
	@echo "  make build             Build Go & Python services"
	@echo "  make rebuild           Rebuild everything via Docker Compose"
	@echo "  make dev               Run all services via Compose"
	@echo "  make dev-down          Stop and cleanup Compose"
	@echo "  make logs              Tail logs from Python container"
	@echo "  make ps                Show Compose container status"
	@echo "\n🧪 Python"
	@echo "  make lint-python       Ruff check"
	@echo "  make format-python     Ruff autofix"
	@echo "  make test-python       Run unittests"
	@echo "  make ci-python         Lint + test + health"
	@echo "\n🧪 Go"
	@echo "  make lint-go           GolangCI lint"
	@echo "  make test-go           Run Go tests"
	@echo "  make gosec             Run gosec security check"
	@echo "  make ci-go             Lint + test + security"
	@echo "\n✅ Health"
	@echo "  make health-all        Check all services"
	@echo "\n🐳 Docker"
	@echo "  make build-python      Build Python Docker image"
	@echo "  make build-go          Build Go Docker image"
	@echo "  make run-docker-python Run Python container manually"
	@echo "  make run-docker-go     Run Go container manually"
	@echo "\n🧪 CI"
	@echo "  make ci-all            Run all CI targets"

## 🚀 Local Dev

build: build-python build-go
rebuild:
	docker-compose down --remove-orphans --volumes && docker-compose up --build

dev:
	docker-compose up --build

dev-down:
	docker-compose down --remove-orphans --volumes

logs:
	docker ps -q --filter ancestor=python-ai-service | xargs -r docker logs --tail=50

ps:
	docker-compose ps

## 🧪 Python: Lint, Test, CI

lint-python:
	ruff check python-ai

format-python:
	ruff check --fix python-ai

test-python:
	PYTHONPATH=python-ai python3 -m unittest discover -s python-ai/tests -p "test_*.py"

ci-python: lint-python test-python health-python

## ✅ Python healthcheck

health-python:
	./python-ai/health_check.sh

## 🧪 Go: Lint, Test, CI

lint-go:
	cd go-backend && golangci-lint run ./...

test-go:
	cd go-backend && go test ./...

gosec:
	cd go-backend && gosec ./...

ci-go: lint-go gosec test-go

## ✅ Go healthcheck

health-go:
	curl -sf http://localhost:8000/health || (echo "❌ go-backend /health failed" && exit 1)

## ✅ Frontend healthcheck

health-frontend:
	curl -sf http://localhost:8080/health || (echo "❌ frontend /health failed" && exit 1)

## ✅ Check all /health endpoints

health-all: health-python health-go health-frontend

## 🐳 Build individual images

build-python:
	docker build -t python-ai-service ./python-ai

build-go:
	docker build -t go-backend-service ./go-backend

run-docker-python:
	docker run -p 8001:8001 --env-file=./python-ai/.env -e ENV=production python-ai-service

run-docker-go:
	docker run -p 8000:8000 go-backend-service

## 🔁 Aggregate CI

ci-all: ci-python ci-go
