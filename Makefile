# Prevent caching issues and ensure all targets run fresh every time
.PHONY: health lint format test build run-docker ci logs

# ✅ Ping /health endpoint to confirm service is running
health:
	./python-ai/health_check.sh

# 🧼 Lint Python code using ruff (report only)
lint:
	ruff check python-ai

# 🛠 Auto-fix style issues using ruff (safe fixes only)
format:
	ruff check --fix python-ai

# 🧪 Run unit tests from root
test:
	PYTHONPATH=python-ai python3 -m unittest discover -s python-ai/tests -p "test_*.py"

# 🐳 Build Docker image for Python AI service
build:
	docker build -t python-ai-service ./python-ai

# 🐳 Run Python AI service in Docker
run-docker:
	docker run -p 8001:8001 --env-file=./python-ai/.env -e ENV=production python-ai-service

# 📄 View recent logs from the Python AI service container
logs:
	docker ps -q --filter ancestor=python-ai-service | xargs -r docker logs --tail=50

# ✅ CI target — run lint, test, and health checks in sequence
ci: lint test health